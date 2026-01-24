package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dl-alexandre/gdrive/internal/types"
	"github.com/dl-alexandre/gdrive/internal/utils"
	"github.com/zalando/go-keyring"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/docs/v1"
	"google.golang.org/api/sheets/v4"
	"google.golang.org/api/slides/v1"
)

const (
	serviceName        = "gdrive-cli"
	tokenRefreshBuffer = 5 * time.Minute
)

// Manager handles authentication operations
type Manager struct {
	configDir      string
	useKeyring     bool
	useEncryption  bool
	storage        StorageBackend
	oauthConfig    *oauth2.Config
	storageWarning string
}

// NewManager creates a new auth manager
func NewManager(configDir string) *Manager {
	return NewManagerWithOptions(configDir, ManagerOptions{})
}

// ManagerOptions configures the auth manager
type ManagerOptions struct {
	ForceEncryptedFile bool // Force use of encrypted file storage
	ForcePlainFile     bool // Force use of plain file storage (insecure, dev only)
}

// NewManagerWithOptions creates a new auth manager with specific options
func NewManagerWithOptions(configDir string, opts ManagerOptions) *Manager {
	mgr := &Manager{
		configDir: configDir,
	}

	// Determine storage backend
	if opts.ForcePlainFile {
		// Plain file storage (insecure, development only)
		mgr.storage = NewPlainFileStorage(configDir)
		mgr.useKeyring = false
		mgr.useEncryption = false
		mgr.storageWarning = "WARNING: Using unencrypted file storage. Credentials are stored in plain text."
	} else if opts.ForceEncryptedFile || !checkKeyringAvailable() {
		// Encrypted file storage
		storage, err := NewEncryptedFileStorage(configDir)
		if err != nil {
			// Fallback to plain file if encryption setup fails
			mgr.storage = NewPlainFileStorage(configDir)
			mgr.useEncryption = false
			mgr.storageWarning = fmt.Sprintf("WARNING: Encryption setup failed (%v). Using plain file storage.", err)
		} else {
			mgr.storage = storage
			mgr.useEncryption = true
			if !opts.ForceEncryptedFile {
				mgr.storageWarning = "INFO: System keyring not available. Using encrypted file storage."
			}
		}
		mgr.useKeyring = false
	} else {
		// System keyring (preferred)
		mgr.storage = NewKeyringStorage(serviceName)
		mgr.useKeyring = true
		mgr.useEncryption = false
	}

	return mgr
}

// checkKeyringAvailable tests if system keyring is available
func checkKeyringAvailable() bool {
	testKey := "gdrive-cli-test"
	err := keyring.Set(serviceName, testKey, "test")
	if err != nil {
		return false
	}
	_ = keyring.Delete(serviceName, testKey)
	return true
}

// SetOAuthConfig sets the OAuth2 configuration
func (m *Manager) SetOAuthConfig(clientID, clientSecret string, scopes []string) {
	m.oauthConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8085/callback",
	}
}

// GetOAuthConfig returns the current OAuth2 configuration
func (m *Manager) GetOAuthConfig() *oauth2.Config {
	return m.oauthConfig
}

// LoadCredentials loads stored credentials for a profile
func (m *Manager) LoadCredentials(profile string) (*types.Credentials, error) {
	stored, err := m.loadStoredCredentials(profile)
	if err != nil {
		return nil, err
	}

	expiryDate, err := time.Parse(time.RFC3339, stored.ExpiryDate)
	if err != nil {
		return nil, fmt.Errorf("invalid expiry date: %w", err)
	}

	return &types.Credentials{
		AccessToken:         stored.AccessToken,
		RefreshToken:        stored.RefreshToken,
		ExpiryDate:          expiryDate,
		Scopes:              stored.Scopes,
		Type:                stored.Type,
		ServiceAccountEmail: stored.ServiceAccountEmail,
		ImpersonatedUser:    stored.ImpersonatedUser,
	}, nil
}

// SaveCredentials saves credentials for a profile
func (m *Manager) SaveCredentials(profile string, creds *types.Credentials) error {
	stored := types.StoredCredentials{
		Profile:             profile,
		AccessToken:         creds.AccessToken,
		RefreshToken:        creds.RefreshToken,
		ExpiryDate:          creds.ExpiryDate.Format(time.RFC3339),
		Scopes:              creds.Scopes,
		Type:                creds.Type,
		ServiceAccountEmail: creds.ServiceAccountEmail,
		ImpersonatedUser:    creds.ImpersonatedUser,
	}

	data, err := json.Marshal(stored)
	if err != nil {
		return fmt.Errorf("failed to marshal credentials: %w", err)
	}

	if err := m.storage.Save(profile, data); err != nil {
		return err
	}

	// Track profile for keyring storage
	if err := m.addProfileToList(profile); err != nil {
		// Non-fatal error, just log it
		fmt.Fprintf(os.Stderr, "Warning: failed to update profile list: %v\n", err)
	}

	return nil
}

// DeleteCredentials removes credentials for a profile
func (m *Manager) DeleteCredentials(profile string) error {
	if err := m.storage.Delete(profile); err != nil {
		return err
	}

	// Remove from profile list
	if err := m.removeProfileFromList(profile); err != nil {
		// Non-fatal error, just log it
		fmt.Fprintf(os.Stderr, "Warning: failed to update profile list: %v\n", err)
	}

	return nil
}

// NeedsRefresh checks if credentials need refreshing
func (m *Manager) NeedsRefresh(creds *types.Credentials) bool {
	return time.Now().Add(tokenRefreshBuffer).After(creds.ExpiryDate)
}

// RefreshCredentials refreshes OAuth2 tokens
func (m *Manager) RefreshCredentials(ctx context.Context, creds *types.Credentials) (*types.Credentials, error) {
	if creds.Type != types.AuthTypeOAuth {
		return nil, fmt.Errorf("refresh only supported for OAuth credentials")
	}
	if m.oauthConfig == nil {
		return nil, fmt.Errorf("OAuth config not set")
	}

	token := &oauth2.Token{
		AccessToken:  creds.AccessToken,
		RefreshToken: creds.RefreshToken,
		Expiry:       creds.ExpiryDate,
	}

	tokenSource := m.oauthConfig.TokenSource(ctx, token)
	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %w", err)
	}

	return &types.Credentials{
		AccessToken:  newToken.AccessToken,
		RefreshToken: newToken.RefreshToken,
		ExpiryDate:   newToken.Expiry,
		Scopes:       creds.Scopes,
		Type:         types.AuthTypeOAuth,
	}, nil
}

// GetValidCredentials returns valid credentials, refreshing if necessary
func (m *Manager) GetValidCredentials(ctx context.Context, profile string) (*types.Credentials, error) {
	creds, err := m.LoadCredentials(profile)
	if err != nil {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeAuthRequired,
			"No credentials found. Run 'gdrive auth login' first.").Build())
	}

	if creds.Type == types.AuthTypeServiceAccount || creds.Type == types.AuthTypeImpersonated {
		if time.Now().After(creds.ExpiryDate) {
			return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeAuthExpired,
				"Service account token expired. Run 'gdrive auth service-account' to re-authenticate.").Build())
		}
		return creds, nil
	}

	if m.NeedsRefresh(creds) {
		newCreds, err := m.RefreshCredentials(ctx, creds)
		if err != nil {
			return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeAuthExpired,
				"Token refresh failed. Run 'gdrive auth login' to re-authenticate.").Build())
		}
		if err := m.SaveCredentials(profile, newCreds); err != nil {
			return nil, fmt.Errorf("failed to save refreshed credentials: %w", err)
		}
		return newCreds, nil
	}

	return creds, nil
}

// GetHTTPClient returns an authenticated HTTP client
func (m *Manager) GetHTTPClient(ctx context.Context, creds *types.Credentials) *http.Client {
	token := &oauth2.Token{
		AccessToken:  creds.AccessToken,
		RefreshToken: creds.RefreshToken,
		Expiry:       creds.ExpiryDate,
	}
	if m.oauthConfig == nil {
		return oauth2.NewClient(ctx, oauth2.StaticTokenSource(token))
	}
	if creds.Type != types.AuthTypeOAuth {
		return oauth2.NewClient(ctx, oauth2.StaticTokenSource(token))
	}
	return m.oauthConfig.Client(ctx, token)
}

// loadStoredCredentials loads credentials from storage
func (m *Manager) loadStoredCredentials(profile string) (*types.StoredCredentials, error) {
	data, err := m.storage.Load(profile)
	if err != nil {
		return nil, err
	}

	var stored types.StoredCredentials
	if err := json.Unmarshal(data, &stored); err != nil {
		return nil, fmt.Errorf("failed to parse credentials: %w", err)
	}

	return &stored, nil
}



// ValidateScopes checks if credentials have required scopes
func (m *Manager) ValidateScopes(creds *types.Credentials, required []string) error {
	scopeSet := make(map[string]bool)
	for _, s := range creds.Scopes {
		scopeSet[s] = true
	}
	for _, req := range required {
		if !scopeSet[req] {
			return utils.NewAppError(utils.NewCLIError(utils.ErrCodeScopeInsufficient,
				fmt.Sprintf("Missing required scope: %s. Re-authenticate with 'gdrive auth login --preset workspace-full' or 'gdrive auth login --scopes %s'", req, req)).Build())
		}
	}
	return nil
}

// UseKeyring returns whether the manager is using the system keyring
func (m *Manager) UseKeyring() bool {
	return m.useKeyring
}

// ConfigDir returns the configuration directory
func (m *Manager) ConfigDir() string {
	return m.configDir
}

// GetStorageBackend returns the name of the storage backend being used
func (m *Manager) GetStorageBackend() string {
	return m.storage.Name()
}

// GetStorageWarning returns any warning message about the storage backend
func (m *Manager) GetStorageWarning() string {
	return m.storageWarning
}

// GetScopesForCommand returns the required scopes for a command
func (m *Manager) GetScopesForCommand(command string) []string {
	// Map commands to required scopes
	scopeMap := map[string][]string{
		"upload":      {utils.ScopeFile},
		"download":    {utils.ScopeReadonly},
		"delete":      {utils.ScopeFull},
		"share":       {utils.ScopeFull},
		"list":        {utils.ScopeReadonly},
		"search":      {utils.ScopeReadonly},
		"mkdir":       {utils.ScopeFile},
		"copy":        {utils.ScopeFile},
		"move":        {utils.ScopeFull},
		"permissions": {utils.ScopeFull},
		"about":       {utils.ScopeMetadataReadonly},
	}

	if scopes, ok := scopeMap[command]; ok {
		return scopes
	}
	// Default to file scope
	return []string{utils.ScopeFile}
}

// ValidateScopesForCommand validates that credentials have the required scopes for a command
func (m *Manager) ValidateScopesForCommand(creds *types.Credentials, command string) error {
	required := m.GetScopesForCommand(command)
	return m.ValidateScopes(creds, required)
}

func (m *Manager) GetServiceFactory() *ServiceFactory {
	return NewServiceFactory(m)
}

func (m *Manager) GetSheetsService(ctx context.Context, creds *types.Credentials) (*sheets.Service, error) {
	return m.GetServiceFactory().CreateSheetsService(ctx, creds)
}

func (m *Manager) GetDocsService(ctx context.Context, creds *types.Credentials) (*docs.Service, error) {
	return m.GetServiceFactory().CreateDocsService(ctx, creds)
}

func (m *Manager) GetSlidesService(ctx context.Context, creds *types.Credentials) (*slides.Service, error) {
	return m.GetServiceFactory().CreateSlidesService(ctx, creds)
}

func (m *Manager) GetAdminService(ctx context.Context, creds *types.Credentials) (*admin.Service, error) {
	return m.GetServiceFactory().CreateAdminService(ctx, creds)
}

func RequiredScopesForService(svcType ServiceType) []string {
	switch svcType {
	case ServiceDrive:
		return []string{utils.ScopeFile}
	case ServiceSheets:
		return []string{utils.ScopeSheets}
	case ServiceDocs:
		return []string{utils.ScopeDocs}
	case ServiceSlides:
		return []string{utils.ScopeSlides}
	case ServiceAdminDir:
		return []string{utils.ScopeAdminDirectoryUser, utils.ScopeAdminDirectoryGroup}
	default:
		return nil
	}
}

func (m *Manager) ValidateServiceScopes(creds *types.Credentials, svcType ServiceType) error {
	required := RequiredScopesForService(svcType)
	if len(required) == 0 {
		return nil
	}
	return m.ValidateScopes(creds, required)
}
