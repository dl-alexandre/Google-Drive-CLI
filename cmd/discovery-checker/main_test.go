package main

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dl-alexandre/gdrv/internal/discovery"
)

func TestGenerateTypesAndDescriptors(t *testing.T) {
	doc := representativeDiscoveryDocument()
	tmpDir := t.TempDir()

	typesDir := filepath.Join(tmpDir, "types", "drive")
	descPath := filepath.Join(tmpDir, "descriptors", "drive.go")

	if err := generateTypes(doc, typesDir, "github.com/dl-alexandre/gdrv/internal/google/generated/drive"); err != nil {
		t.Fatalf("generateTypes failed: %v", err)
	}
	if err := generateDescriptors(doc, descPath, "github.com/dl-alexandre/gdrv/internal/google/generated/descriptors"); err != nil {
		t.Fatalf("generateDescriptors failed: %v", err)
	}

	typesPath := filepath.Join(typesDir, "types.go")
	firstTypes := readGeneratedFile(t, typesPath)
	firstDescriptors := readGeneratedFile(t, descPath)
	assertParsableGo(t, typesPath, firstTypes)
	assertParsableGo(t, descPath, firstDescriptors)

	if !strings.Contains(firstTypes, "package drive") {
		t.Fatalf("generated types used unexpected package:\n%s", firstTypes)
	}
	if !strings.Contains(firstTypes, `import "time"`) {
		t.Fatalf("generated types should import time for date-time fields:\n%s", firstTypes)
	}
	if !strings.Contains(firstTypes, "ModifiedTime time.Time") {
		t.Fatalf("generated types missing mapped date-time field:\n%s", firstTypes)
	}
	if !strings.Contains(firstTypes, `RoleOwner Role = "owner"`) {
		t.Fatalf("generated types missing enum constant:\n%s", firstTypes)
	}

	if !strings.Contains(firstDescriptors, "package descriptors") {
		t.Fatalf("generated descriptors used unexpected package:\n%s", firstDescriptors)
	}
	if !strings.Contains(firstDescriptors, "var Files = ResourceDescriptor") {
		t.Fatalf("generated descriptors missing resource descriptor:\n%s", firstDescriptors)
	}
	if !strings.Contains(firstDescriptors, `"drive.files.get": {`) {
		t.Fatalf("generated descriptors missing method descriptor:\n%s", firstDescriptors)
	}
	if !strings.Contains(firstDescriptors, `Name: "fileId"`) {
		t.Fatalf("generated descriptors missing parameter descriptor:\n%s", firstDescriptors)
	}

	if err := generateTypes(doc, typesDir, "github.com/dl-alexandre/gdrv/internal/google/generated/drive"); err != nil {
		t.Fatalf("second generateTypes failed: %v", err)
	}
	if err := generateDescriptors(doc, descPath, "github.com/dl-alexandre/gdrv/internal/google/generated/descriptors"); err != nil {
		t.Fatalf("second generateDescriptors failed: %v", err)
	}

	if secondTypes := readGeneratedFile(t, typesPath); secondTypes != firstTypes {
		t.Fatal("generated types are not stable across repeated runs")
	}
	if secondDescriptors := readGeneratedFile(t, descPath); secondDescriptors != firstDescriptors {
		t.Fatal("generated descriptors are not stable across repeated runs")
	}
}

func representativeDiscoveryDocument() *discovery.DiscoveryDocument {
	return &discovery.DiscoveryDocument{
		Name:        "drive",
		Version:     "v3",
		Title:       "Drive API",
		BaseURL:     "https://www.googleapis.com/drive/v3/",
		RootURL:     "https://www.googleapis.com/",
		ServicePath: "drive/v3/",
		Schemas: map[string]discovery.Schema{
			"File": {
				Type:        "object",
				Description: "A Drive file.",
				Required:    []string{"id"},
				Properties: map[string]discovery.Schema{
					"id": {
						Type:        "string",
						Description: "The file ID.",
					},
					"modifiedTime": {
						Type:        "string",
						Format:      "date-time",
						Description: "The last modification time.",
					},
					"parents": {
						Type:        "array",
						Description: "Parent IDs.",
						Items:       &discovery.Schema{Type: "string"},
					},
					"size": {
						Type:   "integer",
						Format: "int64",
					},
				},
			},
			"Role": {
				Type:        "string",
				Description: "Permission role.",
				Enum:        []string{"owner", "writer"},
			},
		},
		Resources: map[string]discovery.Resource{
			"files": {
				Methods: map[string]discovery.Method{
					"list": {
						ID:          "drive.files.list",
						HTTPMethod:  "GET",
						Path:        "files",
						Description: "Lists files.",
						Parameters: map[string]discovery.Parameter{
							"pageSize": {Type: "integer", Location: "query"},
							"fields":   {Type: "string", Location: "query"},
						},
					},
					"get": {
						ID:          "drive.files.get",
						HTTPMethod:  "GET",
						Path:        "files/{fileId}",
						Description: "Gets a file.",
						Response:    &discovery.TypeRef{Ref: "File"},
						Parameters: map[string]discovery.Parameter{
							"fileId": {Type: "string", Location: "path", Required: true},
							"fields": {Type: "string", Location: "query"},
						},
					},
				},
			},
		},
	}
}

func readGeneratedFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read %s: %v", path, err)
	}
	return string(data)
}

func assertParsableGo(t *testing.T, path, source string) {
	t.Helper()
	if _, err := parser.ParseFile(token.NewFileSet(), path, source, parser.ParseComments); err != nil {
		t.Fatalf("generated Go does not parse for %s: %v\n%s", path, err, source)
	}
}
