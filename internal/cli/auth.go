package cli

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/dl-alexandre/gdrv/internal/auth"
	"github.com/dl-alexandre/gdrv/internal/config"
	"github.com/dl-alexandre/gdrv/internal/types"
	"github.com/dl-alexandre/gdrv/internal/utils"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authentication commands",
	Long:  "Manage authentication with Google Drive API",
}

var authLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate with Google Drive",
	Long:  "Initiate OAuth2 authentication flow to obtain credentials",
	RunE:  runAuthLogin,
}

var authLogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove stored credentials",
	Long:  "Delete stored credentials for the current or specified profile",
	RunE:  runAuthLogout,
}

var authServiceAccountCmd = &cobra.Command{
	Use:   "service-account",
	Short: "Authenticate with a service account",
	Long:  "Load service account credentials from a JSON key file",
	RunE:  runAuthServiceAccount,
}

var authStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show authentication status",
	Long:  "Display current authentication status and credential information",
	RunE:  runAuthStatus,
}

var authDeviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Authenticate with device code flow",
	Long:  "Initiate device code authentication flow to obtain credentials",
	RunE:  runAuthDevice,
}

var authProfilesCmd = &cobra.Command{
	Use:   "profiles",
	Short: "List credential profiles",
	Long:  "Display all stored credential profiles",
	RunE:  runAuthProfiles,
}

var authDiagnoseCmd = &cobra.Command{
	Use:   "diagnose",
	Short: "Diagnose authentication configuration",
	Long:  "Show detailed authentication diagnostics and token status",
	RunE:  runAuthDiagnose,
}

var (
	authScopes               []string
	authNoBrowser            bool
	authWide                 bool
	authPreset               string
	authKeyFile              string
	authImpersonateUser      string
	clientID                 string
	clientSecret             string
	authDiagnoseRefreshCheck bool
)

func init() {
	authLoginCmd.Flags().StringSliceVar(&authScopes, "scopes", []string{}, "OAuth scopes to request")
	authLoginCmd.Flags().BoolVar(&authNoBrowser, "no-browser", false, "Do not open a browser; use manual code entry")
	authLoginCmd.Flags().BoolVar(&authWide, "wide", false, "Request full Drive access scope")
	authLoginCmd.Flags().StringVar(&authPreset, "preset", "", "Scope preset: workspace-basic, workspace-full, admin, workspace-with-admin, workspace-activity, workspace-labels, workspace-sync, workspace-complete")
	authLoginCmd.Flags().StringVar(&clientID, "client-id", "", "OAuth client ID")
	authLoginCmd.Flags().StringVar(&clientSecret, "client-secret", "", "OAuth client secret")
	authDeviceCmd.Flags().BoolVar(&authWide, "wide", false, "Request full Drive access scope")
	authDeviceCmd.Flags().StringVar(&authPreset, "preset", "", "Scope preset: workspace-basic, workspace-full, admin, workspace-with-admin, workspace-activity, workspace-labels, workspace-sync, workspace-complete")
	authServiceAccountCmd.Flags().StringVar(&authKeyFile, "key-file", "", "Path to service account JSON key file (required)")
	authServiceAccountCmd.Flags().StringVar(&authImpersonateUser, "impersonate-user", "", "User email to impersonate (required for Admin SDK scopes)")
	authServiceAccountCmd.Flags().StringSliceVar(&authScopes, "scopes", []string{}, "OAuth scopes to request")
	authServiceAccountCmd.Flags().BoolVar(&authWide, "wide", false, "Request full Drive access scope")
	authServiceAccountCmd.Flags().StringVar(&authPreset, "preset", "", "Scope preset: workspace-basic, workspace-full, admin, workspace-with-admin, workspace-activity, workspace-labels, workspace-sync, workspace-complete")
	_ = authServiceAccountCmd.MarkFlagRequired("key-file")
	authDiagnoseCmd.Flags().BoolVar(&authDiagnoseRefreshCheck, "refresh-check", false, "Attempt a token refresh and report errors")

	authCmd.AddCommand(authLoginCmd)
	authCmd.AddCommand(authDeviceCmd)
	authCmd.AddCommand(authServiceAccountCmd)
	authCmd.AddCommand(authLogoutCmd)
	authCmd.AddCommand(authStatusCmd)
	authCmd.AddCommand(authProfilesCmd)
	authCmd.AddCommand(authDiagnoseCmd)
	rootCmd.AddCommand(authCmd)
}

func runAuthLogin(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	configDir := getConfigDir()
	resolvedID, resolvedSecret, source, cliErr := resolveOAuthClient(cmd, configDir, false)
	if cliErr != nil {
		return out.WriteError("auth.login", cliErr.Build())
	}
	clientID = resolvedID
	clientSecret = resolvedSecret

	if source == oauthClientSourceBundled {
		out.Log("Using bundled OAuth client credentials.")
	}

	mgr := auth.NewManager(configDir)

	// Display storage warning if any
	if warning := mgr.GetStorageWarning(); warning != "" {
		out.Log("%s", warning)
	}

	scopes, err := resolveAuthScopes(out)
	if err != nil {
		return err
	}
	mgr.SetOAuthConfig(clientID, clientSecret, scopes)

	ctx := context.Background()
	var creds *types.Credentials
	creds, err = mgr.Authenticate(ctx, flags.Profile, openBrowser, auth.OAuthAuthOptions{
		NoBrowser: authNoBrowser,
	})

	if err != nil {
		return out.WriteError("auth.login", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	out.Log("Successfully authenticated!")
	return out.WriteSuccess("auth.login", map[string]interface{}{
		"profile":        flags.Profile,
		"scopes":         creds.Scopes,
		"expiry":         creds.ExpiryDate.Format("2006-01-02T15:04:05Z07:00"),
		"storageBackend": mgr.GetStorageBackend(),
	})
}

func runAuthDevice(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	configDir := getConfigDir()
	resolvedID, resolvedSecret, source, cliErr := resolveOAuthClient(cmd, configDir, false)
	if cliErr != nil {
		return out.WriteError("auth.device", cliErr.Build())
	}
	clientID = resolvedID
	clientSecret = resolvedSecret
	if source == oauthClientSourceBundled {
		out.Log("Using bundled OAuth client credentials.")
	}

	mgr := auth.NewManager(configDir)

	// Display storage warning if any
	if warning := mgr.GetStorageWarning(); warning != "" {
		out.Log("%s", warning)
	}

	scopes, err := resolveAuthScopes(out)
	if err != nil {
		return err
	}

	mgr.SetOAuthConfig(clientID, clientSecret, scopes)

	ctx := context.Background()
	out.Log("Using device code authentication flow...")
	creds, err := mgr.AuthenticateWithDeviceCode(ctx, flags.Profile)

	if err != nil {
		return out.WriteError("auth.device", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	out.Log("Successfully authenticated!")
	return out.WriteSuccess("auth.device", map[string]interface{}{
		"profile":        flags.Profile,
		"scopes":         creds.Scopes,
		"expiry":         creds.ExpiryDate.Format("2006-01-02T15:04:05Z07:00"),
		"storageBackend": mgr.GetStorageBackend(),
	})
}

func runAuthLogout(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	configDir := getConfigDir()
	mgr := auth.NewManager(configDir)

	if err := mgr.DeleteCredentials(flags.Profile); err != nil {
		return out.WriteError("auth.logout", utils.NewCLIError(utils.ErrCodeAuthRequired,
			fmt.Sprintf("No credentials found for profile '%s'", flags.Profile)).Build())
	}

	out.Log("Credentials removed for profile: %s", flags.Profile)
	return out.WriteSuccess("auth.logout", map[string]interface{}{
		"profile": flags.Profile,
		"status":  "logged_out",
	})
}

func runAuthStatus(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	configDir := getConfigDir()
	mgr := auth.NewManager(configDir)

	// Show storage backend info
	if warning := mgr.GetStorageWarning(); warning != "" && flags.Verbose {
		out.Log("%s", warning)
	}

	creds, err := mgr.LoadCredentials(flags.Profile)
	if err != nil {
		return out.WriteSuccess("auth.status", map[string]interface{}{
			"profile":        flags.Profile,
			"authenticated":  false,
			"storageBackend": mgr.GetStorageBackend(),
		})
	}

	expired := time.Now().After(creds.ExpiryDate)
	authenticated := !expired || (creds.Type != types.AuthTypeServiceAccount && creds.Type != types.AuthTypeImpersonated)

	return out.WriteSuccess("auth.status", map[string]interface{}{
		"profile":        flags.Profile,
		"authenticated":  authenticated,
		"scopes":         creds.Scopes,
		"expiry":         creds.ExpiryDate.Format("2006-01-02T15:04:05Z07:00"),
		"type":           creds.Type,
		"needsRefresh":   mgr.NeedsRefresh(creds),
		"expired":        expired,
		"serviceAccount": creds.ServiceAccountEmail,
		"impersonated":   creds.ImpersonatedUser,
		"storageBackend": mgr.GetStorageBackend(),
	})
}

func getConfigDir() string {
	dir, err := config.GetConfigDir()
	if err == nil {
		return dir
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "gdrv")
}

func runAuthProfiles(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	configDir := getConfigDir()
	mgr := auth.NewManager(configDir)

	profiles, err := mgr.ListProfiles()
	if err != nil {
		return out.WriteError("auth.profiles", utils.NewCLIError(utils.ErrCodeUnknown,
			fmt.Sprintf("Failed to list profiles: %v", err)).Build())
	}

	// Get detailed info for each profile
	var profileDetails []map[string]interface{}
	for _, profile := range profiles {
		detail := map[string]interface{}{
			"profile": profile,
		}

		creds, err := mgr.LoadCredentials(profile)
		if err == nil {
			detail["authenticated"] = true
			detail["type"] = creds.Type
			detail["expiry"] = creds.ExpiryDate.Format("2006-01-02T15:04:05Z07:00")
			detail["needsRefresh"] = mgr.NeedsRefresh(creds)
			detail["scopes"] = creds.Scopes
		} else {
			detail["authenticated"] = false
			detail["error"] = err.Error()
		}

		profileDetails = append(profileDetails, detail)
	}

	return out.WriteSuccess("auth.profiles", map[string]interface{}{
		"profiles":       profileDetails,
		"count":          len(profiles),
		"storageBackend": mgr.GetStorageBackend(),
	})
}

func runAuthServiceAccount(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	if authKeyFile == "" {
		return fmt.Errorf("service account key file required via --key-file")
	}

	scopes, err := resolveAuthScopes(out)
	if err != nil {
		return err
	}
	if err := validateAdminScopesRequireImpersonation(scopes, authImpersonateUser); err != nil {
		return out.WriteError("auth.service-account", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	configDir := getConfigDir()
	mgr := auth.NewManager(configDir)

	creds, err := mgr.LoadServiceAccount(context.Background(), authKeyFile, scopes, authImpersonateUser)
	if err != nil {
		return out.WriteError("auth.service-account", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	if err := mgr.SaveCredentials(flags.Profile, creds); err != nil {
		return out.WriteError("auth.service-account", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	out.Log("Service account loaded")
	return out.WriteSuccess("auth.service-account", map[string]interface{}{
		"profile":        flags.Profile,
		"scopes":         creds.Scopes,
		"type":           creds.Type,
		"serviceAccount": creds.ServiceAccountEmail,
		"impersonated":   creds.ImpersonatedUser,
		"storageBackend": mgr.GetStorageBackend(),
	})
}

func resolveAuthScopes(out *OutputWriter) ([]string, error) {
	if authPreset != "" {
		scopes, err := scopesForPreset(authPreset)
		if err != nil {
			return nil, err
		}
		out.Log("Using scope preset: %s", authPreset)
		return scopes, nil
	}
	if authWide {
		out.Log("Using full Drive scope (%s)", utils.ScopeFull)
		return []string{utils.ScopeFull}, nil
	}
	if len(authScopes) == 0 {
		out.Log("Using default scope preset: workspace-basic")
		return utils.ScopesWorkspaceBasic, nil
	}
	return authScopes, nil
}

func scopesForPreset(preset string) ([]string, error) {
	switch preset {
	case "workspace-basic":
		return utils.ScopesWorkspaceBasic, nil
	case "workspace-full":
		return utils.ScopesWorkspaceFull, nil
	case "admin":
		return utils.ScopesAdmin, nil
	case "workspace-with-admin":
		return utils.ScopesWorkspaceWithAdmin, nil
	case "workspace-activity":
		return utils.ScopesWorkspaceActivity, nil
	case "workspace-labels":
		return utils.ScopesWorkspaceLabels, nil
	case "workspace-sync":
		return utils.ScopesWorkspaceSync, nil
	case "workspace-complete":
		return utils.ScopesWorkspaceComplete, nil
	default:
		return nil, fmt.Errorf("unknown preset: %s", preset)
	}
}

func validateAdminScopesRequireImpersonation(scopes []string, impersonateUser string) error {
	adminScopes := []string{
		utils.ScopeAdminDirectoryUser,
		utils.ScopeAdminDirectoryUserReadonly,
		utils.ScopeAdminDirectoryGroup,
		utils.ScopeAdminDirectoryGroupReadonly,
	}

	hasAdminScope := false
	for _, scope := range scopes {
		for _, adminScope := range adminScopes {
			if scope == adminScope {
				hasAdminScope = true
				break
			}
		}
		if hasAdminScope {
			break
		}
	}

	if hasAdminScope && impersonateUser == "" {
		return fmt.Errorf("admin SDK scopes require --impersonate-user")
	}

	return nil
}

func runAuthDiagnose(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	configDir := getConfigDir()
	mgr := auth.NewManager(configDir)

	resolvedID, resolvedSecret, source, cliErr := resolveOAuthClient(cmd, configDir, !authDiagnoseRefreshCheck)
	if cliErr != nil {
		return out.WriteError("auth.diagnose", cliErr.Build())
	}
	clientID = resolvedID
	clientSecret = resolvedSecret
	if source == oauthClientSourceBundled {
		out.Log("Using bundled OAuth client credentials.")
	}
	if clientID != "" {
		mgr.SetOAuthConfig(clientID, clientSecret, []string{})
	}

	creds, err := mgr.LoadCredentials(flags.Profile)
	if err != nil {
		return out.WriteError("auth.diagnose", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	location, _ := mgr.CredentialLocation(flags.Profile)
	metadata, _ := mgr.LoadAuthMetadata(flags.Profile)

	clientHash := ""
	clientFingerprint := ""
	if metadata != nil {
		clientHash = metadata.ClientIDHash
		clientFingerprint = metadata.ClientIDLast4
	}

	diagnostics := map[string]interface{}{
		"profile":          flags.Profile,
		"storageBackend":   mgr.GetStorageBackend(),
		"tokenLocation":    location,
		"clientIdHash":     clientHash,
		"clientIdLast4":    clientFingerprint,
		"scopes":           creds.Scopes,
		"expiry":           creds.ExpiryDate.Format(time.RFC3339),
		"refreshToken":     creds.RefreshToken != "",
		"type":             creds.Type,
		"serviceAccount":   creds.ServiceAccountEmail,
		"impersonatedUser": creds.ImpersonatedUser,
	}

	if authDiagnoseRefreshCheck && creds.Type == types.AuthTypeOAuth {
		if mgr.GetOAuthConfig() == nil {
			return out.WriteError("auth.diagnose", utils.NewCLIError(utils.ErrCodeAuthClientMissing,
				"OAuth client credentials required for refresh check. Set GDRV_CLIENT_ID (and GDRV_CLIENT_SECRET if required) or pass --client-id/--client-secret.").Build())
		}
		_, refreshErr := mgr.RefreshCredentials(context.Background(), creds)
		if refreshErr != nil {
			if appErr, ok := refreshErr.(*utils.AppError); ok {
				diagnostics["refreshCheck"] = map[string]interface{}{
					"success": false,
					"error":   appErr.CLIError,
				}
			} else {
				diagnostics["refreshCheck"] = map[string]interface{}{
					"success": false,
					"error":   refreshErr.Error(),
				}
			}
		} else {
			diagnostics["refreshCheck"] = map[string]interface{}{
				"success": true,
			}
		}
	}

	return out.WriteSuccess("auth.diagnose", diagnostics)
}

func openBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	default:
		return fmt.Errorf("unsupported platform")
	}
	return cmd.Start()
}

type oauthClientSource string

const (
	oauthClientSourceFlags   oauthClientSource = "flags"
	oauthClientSourceEnv     oauthClientSource = "env"
	oauthClientSourceConfig  oauthClientSource = "config"
	oauthClientSourceBundled oauthClientSource = "bundled"
)

func resolveOAuthClient(cmd *cobra.Command, configDir string, allowMissing bool) (string, string, oauthClientSource, *utils.CLIErrorBuilder) {
	requireCustom := isTruthyEnv("GDRV_REQUIRE_CUSTOM_OAUTH")
	requireSecret := false

	flagIDSet := cmd.Flags().Changed("client-id")
	flagSecretSet := cmd.Flags().Changed("client-secret")
	if flagIDSet || flagSecretSet {
		if clientID == "" || (requireSecret && clientSecret == "") {
			return "", "", "", buildOAuthClientError(utils.ErrCodeAuthClientPartial, configDir,
				"Partial OAuth client override not allowed. Set all required client fields via flags, or clear them to use the default/bundled client if available.")
		}
		return clientID, clientSecret, oauthClientSourceFlags, nil
	}

	envID := strings.TrimSpace(os.Getenv("GDRV_CLIENT_ID"))
	envSecret := strings.TrimSpace(os.Getenv("GDRV_CLIENT_SECRET"))
	if envID != "" || envSecret != "" {
		if envID == "" || (requireSecret && envSecret == "") {
			return "", "", "", buildOAuthClientError(utils.ErrCodeAuthClientPartial, configDir,
				"Partial OAuth client override not allowed. Set all required client fields via environment variables, or clear them to use the default/bundled client if available.")
		}
		return envID, envSecret, oauthClientSourceEnv, nil
	}

	cfg, cfgErr := config.Load()
	if cfgErr != nil {
		return "", "", "", utils.NewCLIError(utils.ErrCodeInvalidArgument, fmt.Sprintf("Failed to load config: %v", cfgErr))
	}
	if cfg.OAuthClientID != "" || cfg.OAuthClientSecret != "" {
		if cfg.OAuthClientID == "" || (requireSecret && cfg.OAuthClientSecret == "") {
			return "", "", "", buildOAuthClientError(utils.ErrCodeAuthClientPartial, configDir,
				"Partial OAuth client override not allowed. Set all required client fields in config, or remove them to use the default/bundled client if available.")
		}
		return cfg.OAuthClientID, cfg.OAuthClientSecret, oauthClientSourceConfig, nil
	}

	if requireCustom && !allowMissing {
		return "", "", "", buildOAuthClientError(utils.ErrCodeAuthClientMissing, configDir,
			"Custom OAuth client required. Set GDRV_CLIENT_ID (and GDRV_CLIENT_SECRET if required) or configure the client in the config file. Bundled credentials are disabled by GDRV_REQUIRE_CUSTOM_OAUTH.")
	}

	if bundledID, bundledSecret, ok := auth.GetBundledOAuthClient(); ok {
		if requireCustom {
			return "", "", "", buildOAuthClientError(utils.ErrCodeAuthClientMissing, configDir,
				"Custom OAuth client required. Set GDRV_CLIENT_ID (and GDRV_CLIENT_SECRET if required) or configure the client in the config file. Bundled credentials are disabled by GDRV_REQUIRE_CUSTOM_OAUTH.")
		}
		return bundledID, bundledSecret, oauthClientSourceBundled, nil
	}

	if allowMissing {
		return "", "", "", nil
	}

	return "", "", "", buildOAuthClientError(utils.ErrCodeAuthClientMissing, configDir,
		"OAuth client credentials missing. Bundled credentials are not available in this build. Provide a custom client via environment variables or config.")
}

func buildOAuthClientError(code, configDir, message string) *utils.CLIErrorBuilder {
	configPath, err := config.GetConfigPath()
	if err != nil {
		configPath = filepath.Join(configDir, config.ConfigFileName)
	}
	tokenHint := filepath.Join(configDir, "credentials")

	fullMessage := fmt.Sprintf(
		"%s\nConfig path: %s\nToken storage: system keyring (preferred) or %s\nUse --no-browser for manual login when running headless.",
		message,
		configPath,
		tokenHint,
	)

	return utils.NewCLIError(code, fullMessage).
		WithContext("configPath", configPath).
		WithContext("tokenLocation", tokenHint)
}

func isTruthyEnv(key string) bool {
	value := strings.ToLower(strings.TrimSpace(os.Getenv(key)))
	return value == "1" || value == "true" || value == "yes" || value == "on"
}
