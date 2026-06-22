package auth

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dl-alexandre/gdrv/internal/types"
	"github.com/dl-alexandre/gdrv/internal/utils"
)

// Resolver handles deterministic auth source resolution with explicit precedence:
// 1. GDRV_TOKEN (pre-obtained access token)
// 2. GDRV_CREDENTIALS_FILE (service account or export bundle)
// 3. Profile store (--profile or default)
// 4. Plaintext fallback (legacy)
type Resolver struct {
	configDir string
	verbose   bool
	manager   *Manager
	opts      ResolveOptions // Store options for impersonation and scope handling
}

// NewResolver creates a new auth resolver
func NewResolver(configDir string, verbose bool) *Resolver {
	return &Resolver{
		configDir: configDir,
		verbose:   verbose,
		manager:   NewManagerWithOptions(configDir, ManagerOptions{}),
	}
}

// NewResolverWithOptions creates a resolver with specific options
func NewResolverWithOptions(configDir string, verbose bool, opts ResolveOptions) *Resolver {
	r := NewResolver(configDir, verbose)
	r.opts = opts
	return r
}

// Resolve determines credentials source using deterministic precedence.
// Always returns an AuthResolution record, even on failure.
func (r *Resolver) Resolve(ctx context.Context, opts ResolveOptions) (*AuthResolution, *types.Credentials, error) {
	r.opts = opts
	now := time.Now()

	// Precedence 1: GDRV_TOKEN env var
	if token := os.Getenv("GDRV_TOKEN"); token != "" {
		resolution, creds, err := r.resolveFromToken(token, now)
		if err == nil {
			return resolution, creds, nil
		}
		// Return resolution even on failure for transparency
		return resolution, nil, err
	}

	// Precedence 2: GDRV_CREDENTIALS_FILE env var
	if credFile := os.Getenv("GDRV_CREDENTIALS_FILE"); credFile != "" {
		resolution, creds, err := r.resolveFromCredentialsFile(ctx, credFile, now)
		if err == nil {
			return resolution, creds, nil
		}
		// Return resolution even on failure
		return resolution, nil, err
	}

	// Precedence 3: Profile store
	profile := opts.Profile
	if profile == "" {
		profile = "default"
	}
	resolution, creds, err := r.resolveFromProfile(ctx, profile, now)
	if err == nil {
		return resolution, creds, nil
	}

	// Precedence 4: Plaintext fallback (handled by manager internally)
	// If profile resolution failed, return the error with resolution info
	return resolution, nil, err
}

// resolveFromToken handles GDRV_TOKEN env var
// Token is treated as non-refreshable bearer token
func (r *Resolver) resolveFromToken(token string, now time.Time) (*AuthResolution, *types.Credentials, error) {
	resolution := &AuthResolution{
		Source:    AuthSourceToken,
		Reason:    "GDRV_TOKEN set",
		Subject:   "token",
		Timestamp: now,
		Path:      "env://GDRV_TOKEN",
	}

	// Attempt to parse as JWT to extract expiry and scopes
	expiry, scopes, err := parseJWTToken(token)
	if err == nil && expiry != nil {
		resolution.ExpiresAt = expiry
		resolution.Scopes = scopes

		// Check if expired
		if now.After(*expiry) {
			return resolution, nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeAuthExpired,
				fmt.Sprintf("GDRV_TOKEN expired at %s. Re-export the token or use credentials file/profile auth.",
					expiry.Format(time.RFC3339))).Build())
		}
	}

	// Build credentials from token
	creds := &types.Credentials{
		AccessToken: token,
		ExpiryDate:  time.Now().Add(1 * time.Hour), // Default 1 hour if unknown
		Scopes:      scopes,
		Type:        types.AuthTypeOAuth,
	}

	if expiry != nil {
		creds.ExpiryDate = *expiry
	}

	if len(scopes) == 0 {
		creds.Scopes = []string{"unknown"}
		resolution.Scopes = []string{"unknown"}
	}

	resolution.Refreshable = false
	return resolution, creds, nil
}

// resolveFromCredentialsFile handles GDRV_CREDENTIALS_FILE env var
func (r *Resolver) resolveFromCredentialsFile(ctx context.Context, filePath string, now time.Time) (*AuthResolution, *types.Credentials, error) {
	resolution := &AuthResolution{
		Source:    AuthSourceCredentialsFile,
		Reason:    "GDRV_CREDENTIALS_FILE set",
		Path:      filePath,
		Timestamp: now,
	}

	// Check if file exists
	if _, err := os.Stat(filePath); err != nil {
		resolution.Reason = fmt.Sprintf("GDRV_CREDENTIALS_FILE set but file not found: %s", filePath)
		return resolution, nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeAuthRequired,
			fmt.Sprintf("Credentials file not found: %s", filePath)).Build())
	}

	// Read file
	data, err := os.ReadFile(filePath)
	if err != nil {
		resolution.Reason = fmt.Sprintf("Failed to read credentials file: %v", err)
		return resolution, nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeAuthRequired,
			fmt.Sprintf("Failed to read credentials file: %v", err)).Build())
	}

	// Detect format
	format, err := detectFileFormat(data)
	if err != nil {
		resolution.Reason = fmt.Sprintf("Invalid credentials file format: %v", err)
		return resolution, nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeAuthRequired,
			fmt.Sprintf("Invalid credentials file format: %v", err)).Build())
	}

	switch format {
	case formatServiceAccount:
		// Determine impersonation user using precedence:
		// 1. CLI flag --impersonate-user (from opts.ImpersonateUser)
		// 2. GDRV_IMPERSONATE_USER env var
		// 3. none
		impersonateUser := r.opts.ImpersonateUser
		if impersonateUser == "" {
			impersonateUser = os.Getenv("GDRV_IMPERSONATE_USER")
		}
		if impersonateUser != "" {
			resolution.Subject = impersonateUser
		}

		// Use default scopes if none specified
		scopes := r.opts.RequiredScopes
		if len(scopes) == 0 {
			// Default to workspace-full scopes
			scopes = []string{"https://www.googleapis.com/auth/drive"}
		}

		// Load service account credentials
		creds, err := r.manager.LoadServiceAccount(ctx, filePath, scopes, impersonateUser)
		if err != nil {
			resolution.Reason = fmt.Sprintf("Failed to load service account: %v", err)
			return resolution, nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeAuthRequired,
				fmt.Sprintf("Failed to load service account: %v", err)).Build())
		}

		// Update resolution with loaded credential info
		resolution.Scopes = creds.Scopes
		resolution.Refreshable = false // SA tokens are refreshed via JWT exchange, not refresh token
		if creds.ServiceAccountEmail != "" {
			resolution.Subject = creds.ServiceAccountEmail
		}
		if creds.ImpersonatedUser != "" {
			resolution.Subject = creds.ImpersonatedUser
		}
		if !creds.ExpiryDate.IsZero() {
			resolution.ExpiresAt = &creds.ExpiryDate
		}

		return resolution, creds, nil

	case formatExportBundle:
		creds, profile, err := parseExportBundleCredentials(data)
		if err != nil {
			resolution.Reason = fmt.Sprintf("Invalid export bundle in %s: %v", filePath, err)
			return resolution, nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeAuthRequired,
				fmt.Sprintf("Invalid export bundle: %v", err)).
				WithContext("credentialsFile", filePath).
				WithContext("suggestedAction", "export credentials again or provide a service account JSON file").
				Build())
		}

		if !creds.ExpiryDate.IsZero() && now.After(creds.ExpiryDate) {
			resolution.Reason = fmt.Sprintf("Export bundle token expired at %s", creds.ExpiryDate.Format(time.RFC3339))
			return resolution, nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeAuthExpired,
				fmt.Sprintf("Export bundle token expired at %s. Export fresh credentials or run 'gdrv auth login'.",
					creds.ExpiryDate.Format(time.RFC3339))).
				WithContext("credentialsFile", filePath).
				Build())
		}

		resolution.Reason = "GDRV_CREDENTIALS_FILE export bundle"
		resolution.Subject = profile
		resolution.Scopes = creds.Scopes
		resolution.Refreshable = creds.RefreshToken != "" && creds.Type == types.AuthTypeOAuth
		if creds.ServiceAccountEmail != "" {
			resolution.Subject = creds.ServiceAccountEmail
		}
		if creds.ImpersonatedUser != "" {
			resolution.Subject = creds.ImpersonatedUser
		}
		if !creds.ExpiryDate.IsZero() {
			resolution.ExpiresAt = &creds.ExpiryDate
		}

		return resolution, creds, nil

	default:
		resolution.Reason = "Unrecognized credentials file format"
		return resolution, nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeAuthRequired,
			"Unrecognized credentials file format. Expected service account JSON or gdrv export bundle.").Build())
	}
}

// resolveFromProfile handles stored profile credentials
func (r *Resolver) resolveFromProfile(ctx context.Context, profile string, now time.Time) (*AuthResolution, *types.Credentials, error) {
	resolution := &AuthResolution{
		Source:    AuthSourceProfile,
		Reason:    fmt.Sprintf("Profile '%s'", profile),
		Subject:   profile,
		Timestamp: now,
	}

	// Use the existing manager to load credentials
	creds, err := r.manager.GetValidCredentials(ctx, profile)
	if err != nil {
		// Enhance error with resolution info
		resolution.Reason = fmt.Sprintf("Profile '%s' not found or invalid", profile)
		return resolution, nil, err
	}

	// Populate resolution from loaded credentials
	resolution.Scopes = creds.Scopes
	resolution.Refreshable = creds.RefreshToken != "" && creds.Type == types.AuthTypeOAuth

	if creds.ServiceAccountEmail != "" {
		resolution.Subject = creds.ServiceAccountEmail
	} else if creds.ImpersonatedUser != "" {
		resolution.Subject = creds.ImpersonatedUser
	}

	if !creds.ExpiryDate.IsZero() {
		resolution.ExpiresAt = &creds.ExpiryDate
	}

	return resolution, creds, nil
}

// parseJWTToken attempts to parse a JWT access token and extract expiry and scopes
// Returns: expiry (may be nil), scopes, error
// Manual base64url decoding - no external JWT library needed
func parseJWTToken(token string) (*time.Time, []string, error) {
	// JWT format: header.payload.signature (3 segments separated by '.')
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		// Not a JWT, treat as opaque token
		return nil, nil, fmt.Errorf("not a JWT format")
	}

	// Decode payload (middle segment)
	payload := parts[1]

	// Base64url decode (with padding handling)
	// JWT uses base64url encoding without padding, but Go's decoder needs padding
	padding := 4 - len(payload)%4
	if padding != 4 {
		payload += strings.Repeat("=", padding)
	}

	decoded, err := base64.URLEncoding.DecodeString(payload)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode JWT payload: %w", err)
	}

	// Parse JSON
	var claims struct {
		Exp   interface{} `json:"exp"`
		Scope string      `json:"scope"`
	}

	if err := json.Unmarshal(decoded, &claims); err != nil {
		return nil, nil, fmt.Errorf("failed to parse JWT claims: %w", err)
	}

	// Extract expiry
	var expiry *time.Time
	if claims.Exp != nil {
		var expTime time.Time
		switch v := claims.Exp.(type) {
		case float64:
			expTime = time.Unix(int64(v), 0)
			expiry = &expTime
		case string:
			// Try to parse as Unix timestamp string
			if ts, err := parseUnixTimestamp(v); err == nil {
				expTime = time.Unix(ts, 0)
				expiry = &expTime
			}
		}
	}

	// Extract scopes
	var scopes []string
	if claims.Scope != "" {
		scopes = strings.Split(claims.Scope, " ")
	}

	return expiry, scopes, nil
}

func parseUnixTimestamp(s string) (int64, error) {
	// Try direct parse
	if ts, err := json.Number(s).Int64(); err == nil {
		return ts, nil
	}
	return 0, fmt.Errorf("not a valid timestamp")
}

// fileFormat represents the detected format of a credentials file
type fileFormat string

const (
	formatServiceAccount fileFormat = "service_account"
	formatExportBundle   fileFormat = "export_bundle"
	formatUnknown        fileFormat = "unknown"
)

type exportBundle struct {
	Version       string          `json:"version"`
	Profile       string          `json:"profile"`
	ProfileName   string          `json:"profile_name"`
	EncryptedData json.RawMessage `json:"encrypted_data"`
	Salt          string          `json:"salt"`
	Nonce         string          `json:"nonce"`
	KeyHint       string          `json:"key_hint"`
	Token         *exportToken    `json:"token"`
	Credentials   json.RawMessage `json:"credentials"`
}

type exportToken struct {
	EncryptedData string `json:"encrypted_data"`
	Nonce         string `json:"nonce"`
	KeyHint       string `json:"key_hint"`
}

// detectFileFormat determines the credential file type by inspecting JSON structure
func detectFileFormat(data []byte) (fileFormat, error) {
	// Try to parse as generic JSON first
	var generic map[string]interface{}
	if err := json.Unmarshal(data, &generic); err != nil {
		return formatUnknown, fmt.Errorf("not valid JSON: %w", err)
	}

	// Check for service account format
	if typeVal, ok := generic["type"].(string); ok && typeVal == "service_account" {
		// Validate required fields
		required := []string{"client_email", "private_key"}
		for _, field := range required {
			if _, ok := generic[field]; !ok {
				return formatUnknown, fmt.Errorf("service account JSON missing required field: %s", field)
			}
		}
		return formatServiceAccount, nil
	}

	// Check for export bundle format (version + encrypted_data)
	if _, hasVersion := generic["version"]; hasVersion {
		if _, hasEncrypted := generic["encrypted_data"]; hasEncrypted {
			return formatExportBundle, nil
		}
		if token, ok := generic["token"].(map[string]interface{}); ok {
			if _, hasEncrypted := token["encrypted_data"]; hasEncrypted {
				return formatExportBundle, nil
			}
		}
	}

	return formatUnknown, fmt.Errorf("unrecognized format. Expected service account JSON (with type='service_account') or gdrv export bundle (with version and encrypted_data)")
}

func parseExportBundleCredentials(data []byte) (*types.Credentials, string, error) {
	var bundle exportBundle
	if err := json.Unmarshal(data, &bundle); err != nil {
		return nil, "", fmt.Errorf("parse bundle JSON: %w", err)
	}
	if bundle.Version == "" {
		return nil, "", fmt.Errorf("missing version")
	}

	profile := bundle.Profile
	if profile == "" {
		profile = bundle.ProfileName
	}

	if len(bundle.Credentials) > 0 {
		creds, err := credentialsFromJSON(bundle.Credentials)
		if err != nil {
			return nil, "", fmt.Errorf("invalid credentials payload: %w", err)
		}
		return creds, profile, nil
	}

	if credentialsFieldPresent(data) {
		creds, err := credentialsFromJSON(data)
		if err != nil {
			return nil, "", fmt.Errorf("invalid credentials payload: %w", err)
		}
		return creds, profile, nil
	}

	if len(bundle.EncryptedData) > 0 {
		creds, err := credentialsFromRawBundleData(bundle.EncryptedData)
		if err == nil {
			return creds, profile, nil
		}
		if bundle.Salt != "" || bundle.Nonce != "" || bundle.KeyHint != "" {
			return nil, "", fmt.Errorf("encrypted export bundles are not supported by this resolver: %w", err)
		}
		return nil, "", fmt.Errorf("invalid encrypted_data payload: %w", err)
	}

	if bundle.Token != nil && bundle.Token.EncryptedData != "" {
		creds, err := credentialsFromEncodedBundleData(bundle.Token.EncryptedData)
		if err == nil {
			return creds, profile, nil
		}
		if bundle.Token.Nonce != "" || bundle.Token.KeyHint != "" {
			return nil, "", fmt.Errorf("encrypted token export bundles are not supported by this resolver: %w", err)
		}
		return nil, "", fmt.Errorf("invalid token encrypted_data payload: %w", err)
	}

	return nil, "", fmt.Errorf("missing credentials payload")
}

func credentialsFieldPresent(data []byte) bool {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return false
	}
	_, ok := fields["access_token"]
	return ok
}

func credentialsFromRawBundleData(raw json.RawMessage) (*types.Credentials, error) {
	trimmed := strings.TrimSpace(string(raw))
	if trimmed == "" || trimmed == "null" {
		return nil, fmt.Errorf("empty encrypted_data")
	}
	if strings.HasPrefix(trimmed, "{") {
		return credentialsFromJSON(raw)
	}

	var encoded string
	if err := json.Unmarshal(raw, &encoded); err != nil {
		return nil, fmt.Errorf("encrypted_data must be a JSON object or string: %w", err)
	}
	return credentialsFromEncodedBundleData(encoded)
}

func credentialsFromEncodedBundleData(encoded string) (*types.Credentials, error) {
	if strings.TrimSpace(encoded) == "" {
		return nil, fmt.Errorf("empty encrypted_data")
	}

	if strings.HasPrefix(strings.TrimSpace(encoded), "{") {
		return credentialsFromJSON([]byte(encoded))
	}

	decoders := []*base64.Encoding{
		base64.StdEncoding,
		base64.RawStdEncoding,
		base64.URLEncoding,
		base64.RawURLEncoding,
	}

	var lastErr error
	for _, decoder := range decoders {
		decoded, err := decoder.DecodeString(encoded)
		if err != nil {
			lastErr = err
			continue
		}
		creds, err := credentialsFromJSON(decoded)
		if err == nil {
			return creds, nil
		}
		lastErr = err
	}

	if lastErr == nil {
		lastErr = fmt.Errorf("payload is not base64 encoded JSON")
	}
	return nil, lastErr
}

func credentialsFromJSON(data []byte) (*types.Credentials, error) {
	var creds types.Credentials
	if err := json.Unmarshal(data, &creds); err == nil && creds.AccessToken != "" {
		normalizeBundleCredentials(&creds)
		if err := validateBundleCredentials(&creds); err != nil {
			return nil, err
		}
		return &creds, nil
	}

	var stored types.StoredCredentials
	if err := json.Unmarshal(data, &stored); err != nil {
		return nil, err
	}
	if stored.AccessToken == "" {
		return nil, fmt.Errorf("missing access_token")
	}
	if stored.ExpiryDate == "" {
		return nil, fmt.Errorf("missing expiry_date")
	}
	expiry, err := time.Parse(time.RFC3339, stored.ExpiryDate)
	if err != nil {
		return nil, fmt.Errorf("invalid expiry_date: %w", err)
	}

	storedCreds := &types.Credentials{
		AccessToken:         stored.AccessToken,
		RefreshToken:        stored.RefreshToken,
		ExpiryDate:          expiry,
		Scopes:              stored.Scopes,
		Type:                stored.Type,
		ClientID:            stored.ClientID,
		ServiceAccountEmail: stored.ServiceAccountEmail,
		ImpersonatedUser:    stored.ImpersonatedUser,
	}
	normalizeBundleCredentials(storedCreds)
	if err := validateBundleCredentials(storedCreds); err != nil {
		return nil, err
	}
	return storedCreds, nil
}

func normalizeBundleCredentials(creds *types.Credentials) {
	if creds.Type == "" {
		creds.Type = types.AuthTypeOAuth
	}
	if len(creds.Scopes) == 0 {
		creds.Scopes = []string{"unknown"}
	}
}

func validateBundleCredentials(creds *types.Credentials) error {
	if creds.AccessToken == "" {
		return fmt.Errorf("missing access_token")
	}
	if creds.ExpiryDate.IsZero() {
		return fmt.Errorf("missing expiry_date")
	}
	return nil
}
