package auth

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/dl-alexandre/gdrv/internal/types"
	"github.com/dl-alexandre/gdrv/internal/utils"
)

// withEnv is a test helper that sets environment variables and restores them after the test.
// It uses t.Cleanup to ensure restoration even if the test panics.
func withEnv(t *testing.T, vars map[string]string, fn func()) {
	// Snapshot current values
	oldValues := make(map[string]string)
	for key := range vars {
		oldValues[key] = os.Getenv(key)
	}

	// Set new values
	for key, value := range vars {
		if value == "" {
			_ = os.Unsetenv(key)
		} else {
			_ = os.Setenv(key, value)
		}
	}

	// Register cleanup to restore original values
	t.Cleanup(func() {
		for key, oldValue := range oldValues {
			if oldValue == "" {
				_ = os.Unsetenv(key)
			} else {
				_ = os.Setenv(key, oldValue)
			}
		}
	})

	// Execute the test function
	fn()
}

// createTempCredentialsFile creates a temporary credentials file for testing
// Returns the file path. The file is automatically cleaned up when the test ends.
func createTempCredentialsFile(t *testing.T, content string) string {
	t.Helper()
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "credentials.json")

	if err := os.WriteFile(tmpFile, []byte(content), 0600); err != nil {
		t.Fatalf("Failed to create temp credentials file: %v", err)
	}

	return tmpFile
}

// TestResolver_TokenBeatsProfile verifies that GDRV_TOKEN takes precedence over stored profile
func TestResolver_TokenBeatsProfile(t *testing.T) {
	configDir := t.TempDir()

	// Create a valid JWT token (expiring in 1 hour)
	// This is a mock JWT: header.payload.signature
	// Payload contains exp claim
	mockJWT := createMockJWT(time.Now().Add(1 * time.Hour))

	withEnv(t, map[string]string{
		"GDRV_TOKEN": mockJWT,
	}, func() {
		resolver := NewResolver(configDir, false)
		resolution, creds, err := resolver.Resolve(context.Background(), ResolveOptions{Profile: "default"})

		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		if resolution.Source != AuthSourceToken {
			t.Errorf("Expected source 'token', got: %s", resolution.Source)
		}

		if resolution.Reason != "GDRV_TOKEN set" {
			t.Errorf("Expected reason 'GDRV_TOKEN set', got: %s", resolution.Reason)
		}

		if creds == nil {
			t.Fatal("Expected credentials, got nil")
		}

		if creds.AccessToken != mockJWT {
			t.Error("Credentials access token mismatch")
		}

		if resolution.Refreshable {
			t.Error("Token-based auth should not be refreshable")
		}
	})
}

// TestResolver_FileBeatsProfile verifies that GDRV_CREDENTIALS_FILE takes precedence over profile
func TestResolver_FileBeatsProfile(t *testing.T) {
	configDir := t.TempDir()

	// Create a mock service account file (with invalid key that will fail to load)
	saContent := `{
		"type": "service_account",
		"client_email": "test@example.com",
		"private_key": "-----BEGIN RSA PRIVATE KEY-----\ntest\n-----END RSA PRIVATE KEY-----"
	}`
	tmpFile := createTempCredentialsFile(t, saContent)

	// Ensure no token is set
	withEnv(t, map[string]string{
		"GDRV_TOKEN":            "",
		"GDRV_CREDENTIALS_FILE": tmpFile,
	}, func() {
		resolver := NewResolver(configDir, false)
		resolution, _, err := resolver.Resolve(context.Background(), ResolveOptions{Profile: "default"})

		// In Phase 2, SA file is detected and loading is attempted
		// This should return an error since the key is invalid, but resolution should show file was tried
		if err == nil {
			t.Fatal("Expected error for invalid SA key")
		}

		if resolution.Source != AuthSourceCredentialsFile {
			t.Errorf("Expected source 'credentials_file', got: %s", resolution.Source)
		}

		// In Phase 2, when a service account file is detected and loading fails,
		// the reason should indicate the failure
		if !stringsContains(resolution.Reason, "Failed to load service account") &&
			!stringsContains(resolution.Reason, "credentials file") {
			t.Errorf("Expected reason to indicate credentials file loading was attempted, got: %s", resolution.Reason)
		}

		if resolution.Path != tmpFile {
			t.Errorf("Expected path '%s', got: %s", tmpFile, resolution.Path)
		}
	})
}

// TestResolver_Precedence_TokenBeatsFile verifies full precedence: token > file > profile
func TestResolver_Precedence_TokenBeatsFile(t *testing.T) {
	configDir := t.TempDir()

	mockJWT := createMockJWT(time.Now().Add(1 * time.Hour))
	saContent := `{"type": "service_account", "client_email": "test@example.com", "private_key": "test"}`
	tmpFile := createTempCredentialsFile(t, saContent)

	// Set both TOKEN and CREDENTIALS_FILE - TOKEN should win
	withEnv(t, map[string]string{
		"GDRV_TOKEN":            mockJWT,
		"GDRV_CREDENTIALS_FILE": tmpFile,
	}, func() {
		resolver := NewResolver(configDir, false)
		resolution, _, err := resolver.Resolve(context.Background(), ResolveOptions{})

		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		if resolution.Source != AuthSourceToken {
			t.Errorf("Expected token to beat file, got source: %s", resolution.Source)
		}
	})
}

// TestResolver_FileBeatsProfile verifies that file takes precedence over stored profile
func TestResolver_FileBeatsProfile_Phase3(t *testing.T) {
	// Skip until Phase 3 is implemented
	t.Skip("Skipping until Phase 3 implements credentials file support")

	configDir := t.TempDir()

	// Create a valid service account file
	saContent := `{
		"type": "service_account",
		"project_id": "test-project",
		"private_key_id": "key-id",
		"private_key": "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAxG...\n-----END RSA PRIVATE KEY-----",
		"client_email": "test@test-project.iam.gserviceaccount.com",
		"client_id": "123456789",
		"auth_uri": "https://accounts.google.com/o/oauth2/auth",
		"token_uri": "https://oauth2.googleapis.com/token"
	}`
	tmpFile := createTempCredentialsFile(t, saContent)

	// Ensure no token, but set credentials file
	withEnv(t, map[string]string{
		"GDRV_TOKEN":            "",
		"GDRV_CREDENTIALS_FILE": tmpFile,
	}, func() {
		// Create a stored profile to verify it doesn't get used
		mgr := NewManager(configDir)
		_ = mgr.SaveCredentials("default", &types.Credentials{
			AccessToken: "stored-token",
			ExpiryDate:  time.Now().Add(1 * time.Hour),
			Scopes:      []string{"drive"},
			Type:        types.AuthTypeOAuth,
		})

		resolver := NewResolver(configDir, false)
		resolution, _, err := resolver.Resolve(context.Background(), ResolveOptions{Profile: "default"})

		if err != nil {
			// Expected in Phase 1 since credentials file loading isn't implemented
			t.Logf("Expected error in Phase 1: %v", err)
		}

		// The resolution should still indicate credentials_file was attempted
		if resolution.Source != AuthSourceCredentialsFile {
			t.Errorf("Expected credentials_file source, got: %s", resolution.Source)
		}
	})
}

// TestResolver_MissingSources verifies behavior when no auth sources are available
func TestResolver_MissingSources(t *testing.T) {
	configDir := t.TempDir()

	// Clear all env vars
	withEnv(t, map[string]string{
		"GDRV_TOKEN":            "",
		"GDRV_CREDENTIALS_FILE": "",
	}, func() {
		resolver := NewResolver(configDir, false)
		resolution, creds, err := resolver.Resolve(context.Background(), ResolveOptions{Profile: "nonexistent"})

		if err == nil {
			t.Fatal("Expected error when no sources available")
		}

		if creds != nil {
			t.Error("Expected nil credentials when resolution fails")
		}

		// Resolution record should still be returned
		if resolution == nil {
			t.Fatal("Expected resolution record even on failure")
		}

		if resolution.Source != AuthSourceProfile {
			t.Errorf("Expected profile source (as fallback attempt), got: %s", resolution.Source)
		}
	})
}

// TestResolver_Token_ExpiredJWT verifies token expiration detection
func TestResolver_Token_ExpiredJWT(t *testing.T) {
	configDir := t.TempDir()

	// Create an expired JWT token
	expiredJWT := createMockJWT(time.Now().Add(-1 * time.Hour))

	withEnv(t, map[string]string{
		"GDRV_TOKEN": expiredJWT,
	}, func() {
		resolver := NewResolver(configDir, false)
		resolution, creds, err := resolver.Resolve(context.Background(), ResolveOptions{})

		if err == nil {
			t.Fatal("Expected error for expired token")
		}

		if creds != nil {
			t.Error("Expected nil credentials for expired token")
		}

		if resolution == nil {
			t.Fatal("Expected resolution record even on error")
		}

		if resolution.Source != AuthSourceToken {
			t.Errorf("Expected token source, got: %s", resolution.Source)
		}

		if resolution.ExpiresAt == nil {
			t.Error("Expected expiry time in resolution")
		}

		// Suppress unused variable warning by logging
		t.Logf("Got credentials: %v", creds)

		// Verify error message mentions token expiration
		if err.Error() == "" || !stringsContains(err.Error(), "expired") {
			t.Errorf("Expected error message to mention expiration, got: %v", err)
		}
	})
}

// TestResolver_Token_ValidOpaqueToken verifies handling of non-JWT access tokens
func TestResolver_Token_ValidOpaqueToken(t *testing.T) {
	configDir := t.TempDir()

	// Opaque token (not JWT format)
	opaqueToken := "ya29.a0AfH6SMBxGK2A0tNq5kP3..."

	withEnv(t, map[string]string{
		"GDRV_TOKEN": opaqueToken,
	}, func() {
		resolver := NewResolver(configDir, false)
		resolution, creds, err := resolver.Resolve(context.Background(), ResolveOptions{})

		if err != nil {
			t.Fatalf("Expected no error for opaque token, got: %v", err)
		}

		if creds == nil {
			t.Fatal("Expected credentials for opaque token")
		}

		if creds.AccessToken != opaqueToken {
			t.Error("Access token mismatch")
		}

		if resolution.ExpiresAt != nil {
			t.Error("Expected nil expiry for opaque token")
		}

		if len(creds.Scopes) != 1 || creds.Scopes[0] != "unknown" {
			t.Error("Expected 'unknown' scope for opaque token")
		}
	})
}

// TestResolver_Token_Malformed verifies handling of malformed tokens
func TestResolver_Token_Malformed(t *testing.T) {
	configDir := t.TempDir()

	testCases := []struct {
		name  string
		token string
	}{
		{"empty", ""},
		{"garbage", "not-a-token"},
		{"invalid_jwt_segments", "header.payload"}, // Missing signature
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			withEnv(t, map[string]string{
				"GDRV_TOKEN": tc.token,
			}, func() {
				resolver := NewResolver(configDir, false)
				resolution, _, err := resolver.Resolve(context.Background(), ResolveOptions{})

				if tc.token == "" {
					// Empty token should fall through to next source (profile)
					if resolution.Source == AuthSourceToken {
						t.Error("Empty token should not select token source")
					}
					return
				}

				// Non-empty but invalid token should try to use it as opaque
				if err != nil {
					t.Logf("Got error (may be acceptable): %v", err)
				}

				if resolution == nil {
					t.Fatal("Expected resolution record")
				}
			})
		})
	}
}

// TestResolver_Profile_SpecificProfile verifies resolving a specific named profile
func TestResolver_Profile_SpecificProfile(t *testing.T) {
	configDir := t.TempDir()

	// Save credentials for a named profile
	mgr := NewManager(configDir)
	_ = mgr.SaveCredentials("work", &types.Credentials{
		AccessToken:  "work-token",
		ExpiryDate:   time.Now().Add(1 * time.Hour),
		Scopes:       []string{"drive", "admin"},
		Type:         types.AuthTypeOAuth,
		RefreshToken: "refresh-token",
	})

	withEnv(t, map[string]string{
		"GDRV_TOKEN":            "",
		"GDRV_CREDENTIALS_FILE": "",
	}, func() {
		resolver := NewResolver(configDir, false)
		resolution, creds, err := resolver.Resolve(context.Background(), ResolveOptions{Profile: "work"})

		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		if resolution.Source != AuthSourceProfile {
			t.Errorf("Expected profile source, got: %s", resolution.Source)
		}

		if resolution.Subject != "work" {
			t.Errorf("Expected subject 'work', got: %s", resolution.Subject)
		}

		if creds.AccessToken != "work-token" {
			t.Error("Access token mismatch")
		}

		if !resolution.Refreshable {
			t.Error("Expected refreshable=true for OAuth with refresh token")
		}

		if len(resolution.Scopes) != 2 {
			t.Errorf("Expected 2 scopes, got: %d", len(resolution.Scopes))
		}
	})
}

// TestResolver_ResolutionRecord_AlwaysReturned verifies resolution is returned even on failure
func TestResolver_ResolutionRecord_AlwaysReturned(t *testing.T) {
	configDir := t.TempDir()

	testCases := []struct {
		name           string
		env            map[string]string
		expectedSource AuthSource
	}{
		{
			name: "expired_token",
			env: map[string]string{
				"GDRV_TOKEN": createMockJWT(time.Now().Add(-1 * time.Hour)),
			},
			expectedSource: AuthSourceToken,
		},
		{
			name: "missing_file",
			env: map[string]string{
				"GDRV_TOKEN":            "",
				"GDRV_CREDENTIALS_FILE": "/nonexistent/path/creds.json",
			},
			expectedSource: AuthSourceCredentialsFile,
		},
		{
			name: "missing_profile",
			env: map[string]string{
				"GDRV_TOKEN":            "",
				"GDRV_CREDENTIALS_FILE": "",
			},
			expectedSource: AuthSourceProfile,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			withEnv(t, tc.env, func() {
				resolver := NewResolver(configDir, false)
				resolution, _, err := resolver.Resolve(context.Background(), ResolveOptions{})

				if err == nil && tc.name != "valid_scenario" {
					// Expected error for these failure cases
					t.Logf("Expected error: %v", err)
				}

				if resolution == nil {
					t.Fatal("Resolution record should always be returned, even on failure")
				}

				if resolution.Source != tc.expectedSource {
					t.Errorf("Expected source %s, got: %s", tc.expectedSource, resolution.Source)
				}

				if resolution.Timestamp.IsZero() {
					t.Error("Expected non-zero timestamp in resolution")
				}

				// Verify JSON serializability (machine-readable shape)
				_, jsonErr := json.Marshal(resolution)
				if jsonErr != nil {
					t.Errorf("Resolution record should be JSON serializable: %v", jsonErr)
				}
			})
		})
	}
}

// TestResolver_FileFormat_ServiceAccount verifies service account file detection
func TestResolver_FileFormat_ServiceAccount(t *testing.T) {
	validSA := `{
		"type": "service_account",
		"project_id": "test",
		"private_key_id": "key123",
		"private_key": "-----BEGIN RSA PRIVATE KEY-----\ntest\n-----END RSA PRIVATE KEY-----",
		"client_email": "test@test.iam.gserviceaccount.com",
		"client_id": "123",
		"auth_uri": "https://accounts.google.com/o/oauth2/auth",
		"token_uri": "https://oauth2.googleapis.com/token"
	}`

	format, err := detectFileFormat([]byte(validSA))
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if format != formatServiceAccount {
		t.Errorf("Expected service_account format, got: %s", format)
	}
}

// TestResolver_FileFormat_ExportBundle verifies export bundle detection
func TestResolver_FileFormat_ExportBundle(t *testing.T) {
	validBundle := `{
		"version": "1.0",
		"profile": "default",
		"created_at": "2024-01-01T00:00:00Z",
		"encrypted_data": "base64encrypted...",
		"salt": "base64salt..."
	}`

	format, err := detectFileFormat([]byte(validBundle))
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if format != formatExportBundle {
		t.Errorf("Expected export_bundle format, got: %s", format)
	}
}

// TestResolver_FileFormat_Unknown verifies error for unrecognized formats
func TestResolver_FileFormat_Unknown(t *testing.T) {
	unknownJSON := `{"random": "data"}`

	format, err := detectFileFormat([]byte(unknownJSON))
	if err == nil {
		t.Error("Expected error for unknown format")
	}

	if format != formatUnknown {
		t.Errorf("Expected unknown format, got: %s", format)
	}
}

func TestResolver_ExportBundle_ResolvesBase64Credentials(t *testing.T) {
	configDir := t.TempDir()
	expiry := time.Now().Add(1 * time.Hour).UTC().Truncate(time.Second)
	payload, err := json.Marshal(types.StoredCredentials{
		Profile:      "work",
		AccessToken:  "bundle-token",
		RefreshToken: "bundle-refresh",
		ExpiryDate:   expiry.Format(time.RFC3339),
		Scopes:       []string{"https://www.googleapis.com/auth/drive"},
		Type:         types.AuthTypeOAuth,
		ClientID:     "client-123",
	})
	if err != nil {
		t.Fatal(err)
	}

	bundle := map[string]any{
		"version":        "1.0",
		"profile":        "work",
		"created_at":     "2024-01-01T00:00:00Z",
		"encrypted_data": base64.StdEncoding.EncodeToString(payload),
	}
	bundleData, err := json.Marshal(bundle)
	if err != nil {
		t.Fatal(err)
	}
	tmpFile := createTempCredentialsFile(t, string(bundleData))

	withEnv(t, map[string]string{
		"GDRV_TOKEN":            "",
		"GDRV_CREDENTIALS_FILE": tmpFile,
	}, func() {
		resolver := NewResolver(configDir, false)
		resolution, creds, err := resolver.Resolve(context.Background(), ResolveOptions{Profile: "default"})
		if err != nil {
			t.Fatalf("expected export bundle to resolve, got %v", err)
		}
		if resolution.Source != AuthSourceCredentialsFile {
			t.Fatalf("expected credentials file source, got %s", resolution.Source)
		}
		if resolution.Subject != "work" {
			t.Fatalf("expected profile subject, got %q", resolution.Subject)
		}
		if !resolution.Refreshable {
			t.Fatal("expected refreshable OAuth bundle")
		}
		if creds.AccessToken != "bundle-token" || creds.RefreshToken != "bundle-refresh" {
			t.Fatalf("unexpected credentials: %+v", creds)
		}
		if !creds.ExpiryDate.Equal(expiry) {
			t.Fatalf("expected expiry %s, got %s", expiry, creds.ExpiryDate)
		}
	})
}

func TestResolver_ExportBundle_RejectsUnsupportedEncryptedPayload(t *testing.T) {
	configDir := t.TempDir()
	tmpFile := createTempCredentialsFile(t, `{
		"version": "1.0",
		"profile": "default",
		"encrypted_data": "not-json-ciphertext",
		"nonce": "nonce",
		"key_hint": "host"
	}`)

	withEnv(t, map[string]string{
		"GDRV_TOKEN":            "",
		"GDRV_CREDENTIALS_FILE": tmpFile,
	}, func() {
		resolver := NewResolver(configDir, false)
		_, _, err := resolver.Resolve(context.Background(), ResolveOptions{})
		if err == nil {
			t.Fatal("expected unsupported encrypted bundle error")
		}

		var appErr *utils.AppError
		if !errors.As(err, &appErr) {
			t.Fatalf("expected AppError, got %T: %v", err, err)
		}
		if appErr.CLIError.Code != utils.ErrCodeAuthRequired {
			t.Fatalf("expected %s, got %s", utils.ErrCodeAuthRequired, appErr.CLIError.Code)
		}
		if !strings.Contains(appErr.CLIError.Message, "encrypted export bundles are not supported") {
			t.Fatalf("expected actionable encrypted bundle message, got %q", appErr.CLIError.Message)
		}
	})
}

func TestResolver_ExportBundle_RejectsExpiredCredentials(t *testing.T) {
	configDir := t.TempDir()
	payload, err := json.Marshal(types.StoredCredentials{
		Profile:     "expired",
		AccessToken: "expired-token",
		ExpiryDate:  time.Now().Add(-1 * time.Hour).UTC().Format(time.RFC3339),
		Scopes:      []string{"https://www.googleapis.com/auth/drive"},
		Type:        types.AuthTypeOAuth,
	})
	if err != nil {
		t.Fatal(err)
	}
	bundleData, err := json.Marshal(map[string]any{
		"version":        "1.0",
		"profile":        "expired",
		"encrypted_data": base64.StdEncoding.EncodeToString(payload),
	})
	if err != nil {
		t.Fatal(err)
	}
	tmpFile := createTempCredentialsFile(t, string(bundleData))

	withEnv(t, map[string]string{
		"GDRV_TOKEN":            "",
		"GDRV_CREDENTIALS_FILE": tmpFile,
	}, func() {
		resolver := NewResolver(configDir, false)
		_, _, err := resolver.Resolve(context.Background(), ResolveOptions{})
		if err == nil {
			t.Fatal("expected expired bundle error")
		}

		var appErr *utils.AppError
		if !errors.As(err, &appErr) {
			t.Fatalf("expected AppError, got %T: %v", err, err)
		}
		if appErr.CLIError.Code != utils.ErrCodeAuthExpired {
			t.Fatalf("expected %s, got %s", utils.ErrCodeAuthExpired, appErr.CLIError.Code)
		}
	})
}

// TestResolver_JWTParsing_ScopesExtracted verifies scope extraction from JWT
func TestResolver_JWTParsing_ScopesExtracted(t *testing.T) {
	// Create JWT with scope claim
	scopes := []string{"https://www.googleapis.com/auth/drive", "https://www.googleapis.com/auth/admin"}
	jwt := createMockJWTWithScopes(time.Now().Add(1*time.Hour), scopes)

	expiry, extractedScopes, err := parseJWTToken(jwt)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if expiry == nil {
		t.Error("Expected expiry from JWT")
	}

	if len(extractedScopes) != len(scopes) {
		t.Errorf("Expected %d scopes, got: %d", len(scopes), len(extractedScopes))
	}

	for i, scope := range scopes {
		if extractedScopes[i] != scope {
			t.Errorf("Scope mismatch at %d: expected %s, got %s", i, scope, extractedScopes[i])
		}
	}
}

// TestResolver_Impersonation_EnvVar verifies GDRV_IMPERSONATE_USER is captured in resolution
func TestResolver_Impersonation_EnvVar(t *testing.T) {
	configDir := t.TempDir()

	saContent := `{
		"type": "service_account",
		"client_email": "sa@test.iam.gserviceaccount.com",
		"private_key": "-----BEGIN RSA PRIVATE KEY-----\ntest\n-----END RSA PRIVATE KEY-----"
	}`
	tmpFile := createTempCredentialsFile(t, saContent)

	withEnv(t, map[string]string{
		"GDRV_TOKEN":            "",
		"GDRV_CREDENTIALS_FILE": tmpFile,
		"GDRV_IMPERSONATE_USER": "admin@example.com",
	}, func() {
		resolver := NewResolver(configDir, false)
		resolution, _, err := resolver.Resolve(context.Background(), ResolveOptions{})

		// In Phase 2, SA file is detected and impersonation is captured
		// but loading will fail due to invalid key - that's expected
		if err == nil {
			t.Fatal("Expected error when loading SA with invalid key")
		}

		// Resolution should include impersonated user
		if resolution.Subject != "admin@example.com" {
			t.Errorf("Expected subject 'admin@example.com', got: %s", resolution.Subject)
		}

		if resolution.Source != AuthSourceCredentialsFile {
			t.Errorf("Expected source credentials_file, got: %s", resolution.Source)
		}
	})
}

// TestResolver_ServiceAccount_FileBeatsProfile_Phase2 verifies SA file takes precedence over profile
func TestResolver_ServiceAccount_FileBeatsProfile_Phase2(t *testing.T) {
	configDir := t.TempDir()

	// Create a valid service account file structure (but with invalid key for testing)
	saContent := `{
		"type": "service_account",
		"project_id": "test-project",
		"private_key_id": "key-id",
		"private_key": "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAxG...invalid...\n-----END RSA PRIVATE KEY-----",
		"client_email": "test@test-project.iam.gserviceaccount.com",
		"client_id": "123456789",
		"auth_uri": "https://accounts.google.com/o/oauth2/auth",
		"token_uri": "https://oauth2.googleapis.com/token"
	}`
	tmpFile := createTempCredentialsFile(t, saContent)

	// Create a stored profile to verify it doesn't get used
	mgr := NewManager(configDir)
	_ = mgr.SaveCredentials("default", &types.Credentials{
		AccessToken: "stored-token",
		ExpiryDate:  time.Now().Add(1 * time.Hour),
		Scopes:      []string{"drive"},
		Type:        types.AuthTypeOAuth,
	})

	// Ensure no token, but set credentials file
	withEnv(t, map[string]string{
		"GDRV_TOKEN":            "",
		"GDRV_CREDENTIALS_FILE": tmpFile,
	}, func() {
		resolver := NewResolver(configDir, false)
		resolution, _, err := resolver.Resolve(context.Background(), ResolveOptions{Profile: "default"})

		// Should detect SA file (will fail to load due to invalid key, but resolution shows it was tried)
		if resolution.Source != AuthSourceCredentialsFile {
			t.Errorf("Expected credentials_file source, got: %s", resolution.Source)
		}

		if resolution.Path != tmpFile {
			t.Errorf("Expected path '%s', got: %s", tmpFile, resolution.Path)
		}

		// Should have error since key is invalid
		if err == nil {
			t.Error("Expected error loading SA with invalid key")
		}
	})
}

// TestResolver_ServiceAccount_FlagBeatsEnv verifies --impersonate-user flag beats GDRV_IMPERSONATE_USER
func TestResolver_ServiceAccount_FlagBeatsEnv(t *testing.T) {
	configDir := t.TempDir()

	saContent := `{
		"type": "service_account",
		"client_email": "sa@test.iam.gserviceaccount.com",
		"private_key": "-----BEGIN RSA PRIVATE KEY-----\ntest\n-----END RSA PRIVATE KEY-----"
	}`
	tmpFile := createTempCredentialsFile(t, saContent)

	withEnv(t, map[string]string{
		"GDRV_TOKEN":            "",
		"GDRV_CREDENTIALS_FILE": tmpFile,
		"GDRV_IMPERSONATE_USER": "env-user@example.com",
	}, func() {
		// Pass flag value via ResolveOptions
		resolver := NewResolver(configDir, false)
		opts := ResolveOptions{
			Profile:         "default",
			ImpersonateUser: "flag-user@example.com", // CLI flag takes precedence
		}
		resolution, _, err := resolver.Resolve(context.Background(), opts)

		// Should fail to load due to invalid key, but verify impersonation was set
		if resolution.Subject != "flag-user@example.com" {
			t.Errorf("Expected flag to beat env (subject='flag-user@example.com'), got: %s", resolution.Subject)
		}

		if err == nil {
			t.Error("Expected error loading SA with invalid key")
		}
	})
}

// TestResolver_ServiceAccount_MissingFile verifies error handling for missing file
func TestResolver_ServiceAccount_MissingFile(t *testing.T) {
	configDir := t.TempDir()

	withEnv(t, map[string]string{
		"GDRV_TOKEN":            "",
		"GDRV_CREDENTIALS_FILE": "/nonexistent/path/sa.json",
	}, func() {
		resolver := NewResolver(configDir, false)
		resolution, _, err := resolver.Resolve(context.Background(), ResolveOptions{})

		if err == nil {
			t.Fatal("Expected error for missing file")
		}

		if resolution.Source != AuthSourceCredentialsFile {
			t.Errorf("Expected source credentials_file, got: %s", resolution.Source)
		}

		if resolution.Path != "/nonexistent/path/sa.json" {
			t.Errorf("Expected path '/nonexistent/path/sa.json', got: %s", resolution.Path)
		}

		if !stringsContains(err.Error(), "not found") && !stringsContains(err.Error(), "unreadable") {
			t.Errorf("Expected error to mention file not found, got: %v", err)
		}
	})
}

// TestResolver_ServiceAccount_InvalidJSON verifies error handling for invalid JSON
func TestResolver_ServiceAccount_InvalidJSON(t *testing.T) {
	configDir := t.TempDir()

	invalidContent := `{invalid json}`
	tmpFile := createTempCredentialsFile(t, invalidContent)

	withEnv(t, map[string]string{
		"GDRV_TOKEN":            "",
		"GDRV_CREDENTIALS_FILE": tmpFile,
	}, func() {
		resolver := NewResolver(configDir, false)
		resolution, _, err := resolver.Resolve(context.Background(), ResolveOptions{})

		if err == nil {
			t.Fatal("Expected error for invalid JSON")
		}

		if resolution.Source != AuthSourceCredentialsFile {
			t.Errorf("Expected source credentials_file, got: %s", resolution.Source)
		}
	})
}

// TestResolver_ServiceAccount_NoImpersonation verifies SA without impersonation
func TestResolver_ServiceAccount_NoImpersonation(t *testing.T) {
	configDir := t.TempDir()

	saContent := `{
		"type": "service_account",
		"client_email": "sa@test.iam.gserviceaccount.com",
		"private_key": "-----BEGIN RSA PRIVATE KEY-----\ntest\n-----END RSA PRIVATE KEY-----"
	}`
	tmpFile := createTempCredentialsFile(t, saContent)

	withEnv(t, map[string]string{
		"GDRV_TOKEN":            "",
		"GDRV_CREDENTIALS_FILE": tmpFile,
		"GDRV_IMPERSONATE_USER": "", // No impersonation
	}, func() {
		resolver := NewResolver(configDir, false)
		opts := ResolveOptions{
			Profile:         "default",
			ImpersonateUser: "", // No flag either
		}
		resolution, _, err := resolver.Resolve(context.Background(), opts)

		// Should attempt SA without impersonation
		if resolution.Source != AuthSourceCredentialsFile {
			t.Errorf("Expected source credentials_file, got: %s", resolution.Source)
		}

		// Subject should be empty or the SA email (depends on implementation)
		// When loading fails early, subject may be empty
		t.Logf("Subject without impersonation: %s", resolution.Subject)

		if err == nil {
			t.Error("Expected error loading SA with invalid key")
		}
	})
}

// Helper functions

func createMockJWT(expiry time.Time) string {
	return createMockJWTWithScopes(expiry, []string{"https://www.googleapis.com/auth/drive"})
}

func createMockJWTWithScopes(expiry time.Time, scopes []string) string {
	// Create minimal JWT structure
	header := base64.URLEncoding.EncodeToString([]byte(`{"alg":"none","typ":"JWT"}`))

	claims := map[string]interface{}{
		"exp":   expiry.Unix(),
		"scope": strings.Join(scopes, " "),
	}

	claimsJSON, _ := json.Marshal(claims)
	payload := base64.URLEncoding.EncodeToString(claimsJSON)

	// JWT without valid signature (we don't validate signatures)
	return header + "." + payload + "."
}

func stringsContains(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}
