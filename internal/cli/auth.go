package cli

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
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

var (
	authScopes    []string
	authNoBrowser bool
	authWide      bool
	authPreset    string
	authKeyFile   string
	authImpersonateUser string
	clientID      string
	clientSecret  string
)

func init() {
	authLoginCmd.Flags().StringSliceVar(&authScopes, "scopes", []string{}, "OAuth scopes to request")
	authLoginCmd.Flags().BoolVar(&authNoBrowser, "no-browser", false, "Use device code flow (limited scopes)")
	authLoginCmd.Flags().BoolVar(&authWide, "wide", false, "Request full Drive access scope")
	authLoginCmd.Flags().StringVar(&authPreset, "preset", "", "Scope preset: workspace-basic, workspace-full, admin, workspace-with-admin")
	authLoginCmd.Flags().StringVar(&clientID, "client-id", "", "OAuth client ID")
	authLoginCmd.Flags().StringVar(&clientSecret, "client-secret", "", "OAuth client secret")
	authDeviceCmd.Flags().BoolVar(&authWide, "wide", false, "Request full Drive access scope")
	authDeviceCmd.Flags().StringVar(&authPreset, "preset", "", "Scope preset: workspace-basic, workspace-full, admin, workspace-with-admin")
	authServiceAccountCmd.Flags().StringVar(&authKeyFile, "key-file", "", "Path to service account JSON key file (required)")
	authServiceAccountCmd.Flags().StringVar(&authImpersonateUser, "impersonate-user", "", "User email to impersonate (required for Admin SDK scopes)")
	authServiceAccountCmd.Flags().StringSliceVar(&authScopes, "scopes", []string{}, "OAuth scopes to request")
	authServiceAccountCmd.Flags().BoolVar(&authWide, "wide", false, "Request full Drive access scope")
	authServiceAccountCmd.Flags().StringVar(&authPreset, "preset", "", "Scope preset: workspace-basic, workspace-full, admin, workspace-with-admin")
	_ = authServiceAccountCmd.MarkFlagRequired("key-file")

	authCmd.AddCommand(authLoginCmd)
	authCmd.AddCommand(authDeviceCmd)
	authCmd.AddCommand(authServiceAccountCmd)
	authCmd.AddCommand(authLogoutCmd)
	authCmd.AddCommand(authStatusCmd)
	authCmd.AddCommand(authProfilesCmd)
	rootCmd.AddCommand(authCmd)
}

func runAuthLogin(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	if clientID == "" || clientSecret == "" {
		clientID = os.Getenv("GDRV_CLIENT_ID")
		clientSecret = os.Getenv("GDRV_CLIENT_SECRET")
		if clientID == "" || clientSecret == "" {
			return fmt.Errorf("OAuth client ID and secret required. Set via --client-id/--client-secret or GDRV_CLIENT_ID/GDRV_CLIENT_SECRET")
		}
	}

	configDir := getConfigDir()
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
	if authNoBrowser {
		out.Log("Using device code authentication flow...")
		creds, err = mgr.AuthenticateWithDeviceCode(ctx, flags.Profile)
	} else {
		creds, err = mgr.Authenticate(ctx, flags.Profile, openBrowser)
	}

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

	if clientID == "" || clientSecret == "" {
		clientID = os.Getenv("GDRV_CLIENT_ID")
		clientSecret = os.Getenv("GDRV_CLIENT_SECRET")
		if clientID == "" || clientSecret == "" {
			return fmt.Errorf("OAuth client ID and secret required. Set via --client-id/--client-secret or GDRV_CLIENT_ID/GDRV_CLIENT_SECRET")
		}
	}

	configDir := getConfigDir()
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
	authenticated := true
	if expired && (creds.Type == types.AuthTypeServiceAccount || creds.Type == types.AuthTypeImpersonated) {
		authenticated = false
	}

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
		return fmt.Errorf("Admin SDK scopes require --impersonate-user")
	}

	return nil
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
