package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	discoverygen "github.com/dl-alexandre/gdrv/cmd/discovery-checker/generator"
	"github.com/dl-alexandre/gdrv/internal/discovery"
	"github.com/dl-alexandre/gdrv/internal/logging"
	"gopkg.in/yaml.v3"
)

// Config represents the apis.yaml configuration
type Config struct {
	APIs       []APIConfig      `yaml:"apis"`
	RiskRules  RiskRules        `yaml:"risk_rules"`
	Generation GenerationConfig `yaml:"generation"`
	CI         CIConfig         `yaml:"ci"`
}

type APIConfig struct {
	Service      string `yaml:"service"`
	Version      string `yaml:"version"`
	DiscoveryURL string `yaml:"discovery_url"`
	Owner        string `yaml:"owner"`
	Priority     int    `yaml:"priority"`
	Description  string `yaml:"description"`
}

type RiskRules struct {
	Additive []string `yaml:"additive"`
	Risky    []string `yaml:"risky"`
	Breaking []string `yaml:"breaking"`
}

type GenerationConfig struct {
	SnapshotDir             string `yaml:"snapshot_dir"`
	GeneratedTypesDir       string `yaml:"generated_types_dir"`
	GeneratedDescriptorsDir string `yaml:"generated_descriptors_dir"`
	GenerateTypes           bool   `yaml:"generate_types"`
	GenerateDescriptors     bool   `yaml:"generate_descriptors"`
	GenerateMethods         bool   `yaml:"generate_methods"`
	GenerateCLI             bool   `yaml:"generate_cli"`
	TemplateDir             string `yaml:"template_dir"`
	GoPackagePrefix         string `yaml:"go_package_prefix"`
}

type CIConfig struct {
	Schedule                  string `yaml:"schedule"`
	AutoMergeAdditive         bool   `yaml:"auto_merge_additive"`
	RequireOwnerApprovalRisky bool   `yaml:"require_owner_approval_risky"`
	BlockOnBreaking           bool   `yaml:"block_on_breaking"`
	CreateIssuesForBreaking   bool   `yaml:"create_issues_for_breaking"`
	SlackChannel              string `yaml:"slack_channel"`
	RunTestsAfterGeneration   bool   `yaml:"run_tests_after_generation"`
	RequireGoFmt              bool   `yaml:"require_gofmt"`
	RequireGoImports          bool   `yaml:"require_goimports"`
}

// Command flags
var (
	configFile = flag.String("config", "apis.yaml", "Path to apis.yaml configuration")
	checkOnly  = flag.Bool("check", false, "Check for drift without generating code")
	generate   = flag.Bool("generate", false, "Generate code from discovery docs")
	verbose    = flag.Bool("verbose", false, "Enable verbose logging")
	service    = flag.String("service", "", "Specific service to check (empty = all)")
	outputDir  = flag.String("output", ".", "Output directory for generated code")
)

func main() {
	flag.Parse()

	logger := logging.NewNoOpLogger()
	if *verbose {
		logger = logging.NewNoOpLogger()
	}

	// Load configuration
	config, err := loadConfig(*configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	// Initialize discovery client
	client := discovery.NewClient(discovery.ClientOptions{
		Logger: logger,
	})

	ctx := context.Background()

	switch {
	case *checkOnly:
		driftDetected, err := checkDrift(ctx, client, config, *service)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error checking drift: %v\n", err)
			os.Exit(1)
		}
		if driftDetected {
			os.Exit(2) // Special exit code to indicate drift
		}
		fmt.Println("No drift detected.")

	case *generate:
		if err := generateCode(ctx, client, config, *service, *outputDir); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating code: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Code generation complete.")

	default:
		fmt.Println("Usage: discovery-checker [-config path] -check|-generate [-service name] [-verbose]")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("parsing config file: %w", err)
	}

	return &config, nil
}

func checkDrift(ctx context.Context, client *discovery.Client, config *Config, specificService string) (bool, error) {
	fmt.Println("Checking for API drift...")

	var driftDetected bool
	var allReports []*ChangeReport

	for _, api := range config.APIs {
		if specificService != "" && api.Service != specificService {
			continue
		}

		fmt.Printf("  Checking %s %s...\n", api.Service, api.Version)

		// Fetch current discovery document
		var doc *discovery.DiscoveryDocument
		var err error

		// Check if a custom discovery URL is provided (not the standard www.googleapis.com)
		if api.DiscoveryURL != "" && !strings.Contains(api.DiscoveryURL, "www.googleapis.com") {
			doc, err = client.GetDiscoveryDocumentFromURL(ctx, api.DiscoveryURL)
		} else {
			doc, err = client.GetDiscoveryDocument(ctx, api.Service, api.Version)
		}

		if err != nil {
			return false, fmt.Errorf("fetching discovery for %s: %w", api.Service, err)
		}

		// Load committed snapshot
		snapshotPath := filepath.Join(config.Generation.SnapshotDir, api.Service, api.Version+".json")
		existingDoc, err := loadSnapshot(snapshotPath)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("    ✗ No existing snapshot found - initial import needed\n")
				driftDetected = true
				// Create a report for this initial import case
				allReports = append(allReports, &ChangeReport{
					API:     api.Service,
					Version: api.Version,
					Changes: []Change{{
						Severity:    SeverityAdditive,
						Category:    CategoryType,
						Description: "Initial import - no previous snapshot",
						API:         api.Service,
					}},
					TotalChanges:    1,
					AdditiveChanges: 1,
				})
				continue
			}
			return false, fmt.Errorf("loading snapshot for %s: %w", api.Service, err)
		}

		// Compare and classify changes using refined classifier
		classifier := NewRefinedClassifier()
		report := classifier.ClassifyChanges(api.Service, existingDoc, doc)
		allReports = append(allReports, report)

		if report.TotalChanges == 0 {
			fmt.Printf("    ✓ No changes detected\n")
			continue
		}

		fmt.Printf("    ✗ %d change(s) detected (%d additive, %d risky, %d breaking):\n",
			report.TotalChanges, report.AdditiveChanges, report.RiskyChanges, report.BreakingChanges)
		for _, change := range report.Changes {
			fmt.Printf("      - %s (%s): %s\n", change.Severity, change.Category, change.Description)
		}
		driftDetected = true
	}

	// Write machine-readable report
	if err := writeJSONReport(allReports); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to write JSON report: %v\n", err)
	}

	return driftDetected, nil
}

// writeJSONReport writes a machine-readable JSON report for tooling integration
func writeJSONReport(reports []*ChangeReport) error {
	summary := map[string]interface{}{
		"timestamp":     time.Now().UTC().Format(time.RFC3339),
		"total_apis":    len(reports),
		"total_changes": 0,
		"additive":      0,
		"risky":         0,
		"breaking":      0,
		"apis":          reports,
	}

	for _, r := range reports {
		summary["total_changes"] = summary["total_changes"].(int) + r.TotalChanges
		summary["additive"] = summary["additive"].(int) + r.AdditiveChanges
		summary["risky"] = summary["risky"].(int) + r.RiskyChanges
		summary["breaking"] = summary["breaking"].(int) + r.BreakingChanges
	}

	data, err := json.MarshalIndent(summary, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling report: %w", err)
	}

	return os.WriteFile("drift_report.json", data, 0644)
}

func generateCode(ctx context.Context, client *discovery.Client, config *Config, specificService, outputDir string) error {
	fmt.Println("Generating code from discovery documents...")

	for _, api := range config.APIs {
		if specificService != "" && api.Service != specificService {
			continue
		}

		fmt.Printf("  Processing %s %s...\n", api.Service, api.Version)

		// Fetch discovery document
		var doc *discovery.DiscoveryDocument
		var err error

		// Check if a custom discovery URL is provided (not the standard www.googleapis.com)
		if api.DiscoveryURL != "" && !strings.Contains(api.DiscoveryURL, "www.googleapis.com") {
			doc, err = client.GetDiscoveryDocumentFromURL(ctx, api.DiscoveryURL)
		} else {
			doc, err = client.GetDiscoveryDocument(ctx, api.Service, api.Version)
		}

		if err != nil {
			return fmt.Errorf("fetching discovery for %s: %w", api.Service, err)
		}

		// Save snapshot
		snapshotPath := filepath.Join(outputDir, config.Generation.SnapshotDir, api.Service, api.Version+".json")
		if err := saveSnapshot(snapshotPath, doc); err != nil {
			return fmt.Errorf("saving snapshot for %s: %w", api.Service, err)
		}

		// Generate types
		if config.Generation.GenerateTypes {
			typesPath := filepath.Join(outputDir, config.Generation.GeneratedTypesDir, api.Service)
			if err := generateTypes(doc, typesPath, config.Generation.GoPackagePrefix+"/"+api.Service); err != nil {
				return fmt.Errorf("generating types for %s: %w", api.Service, err)
			}
		}

		// Generate descriptors
		if config.Generation.GenerateDescriptors {
			descPath := filepath.Join(outputDir, config.Generation.GeneratedDescriptorsDir, api.Service+".go")
			if err := generateDescriptors(doc, descPath, config.Generation.GoPackagePrefix+"/descriptors"); err != nil {
				return fmt.Errorf("generating descriptors for %s: %w", api.Service, err)
			}
		}
	}

	return nil
}

func loadSnapshot(path string) (*discovery.DiscoveryDocument, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var doc discovery.DiscoveryDocument
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("parsing snapshot: %w", err)
	}

	return &doc, nil
}

func saveSnapshot(path string, doc *discovery.DiscoveryDocument) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return fmt.Errorf("serializing document: %w", err)
	}

	return os.WriteFile(path, data, 0644)
}

func generateTypes(doc *discovery.DiscoveryDocument, outputPath, packageName string) error {
	gen, err := discoverygen.NewTypeGenerator(discoverygen.GeneratorConfig{
		PackageName: goPackageName(packageName, doc.Name),
	})
	if err != nil {
		return err
	}

	content, err := gen.Generate(doc, doc.Name)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(outputPath, 0755); err != nil {
		return fmt.Errorf("creating types output directory: %w", err)
	}
	if err := os.WriteFile(filepath.Join(outputPath, "types.go"), content, 0644); err != nil {
		return fmt.Errorf("writing generated types: %w", err)
	}
	return nil
}

func generateDescriptors(doc *discovery.DiscoveryDocument, outputPath, packageName string) error {
	gen, err := discoverygen.NewDescriptorGenerator(discoverygen.GeneratorConfig{
		PackageName: goPackageName(packageName, "descriptors"),
	})
	if err != nil {
		return err
	}

	content, err := gen.Generate(doc, doc.Name)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("creating descriptors output directory: %w", err)
	}
	if err := os.WriteFile(outputPath, content, 0644); err != nil {
		return fmt.Errorf("writing generated descriptors: %w", err)
	}
	return nil
}

func goPackageName(importPath, fallback string) string {
	importPath = strings.Trim(importPath, "/")
	if importPath == "" {
		return fallback
	}
	if idx := strings.LastIndex(importPath, "/"); idx >= 0 {
		importPath = importPath[idx+1:]
	}
	if importPath == "" {
		return fallback
	}
	return importPath
}

// Required imports
var (
	_ = context.Background
	_ = time.Now
)
