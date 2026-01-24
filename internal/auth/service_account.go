package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/dl-alexandre/gdrive/internal/types"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

// ServiceAccountKey represents the JSON structure of a service account key file
type ServiceAccountKey struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
}

// LoadServiceAccount loads credentials from service account key file
func (m *Manager) LoadServiceAccount(ctx context.Context, keyFilePath string, scopes []string, impersonateUser string) (*types.Credentials, error) {
	if keyFilePath == "" {
		return nil, fmt.Errorf("service account key file required")
	}
	if _, err := os.Stat(keyFilePath); err != nil {
		return nil, fmt.Errorf("service account key file not found: %s", keyFilePath)
	}
	if len(scopes) == 0 {
		return nil, fmt.Errorf("at least one scope required")
	}
	if impersonateUser != "" && !strings.Contains(impersonateUser, "@") {
		return nil, fmt.Errorf("impersonate user must be an email address")
	}
	keyData, err := os.ReadFile(keyFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read service account key: %w", err)
	}

	// Parse service account key to extract email
	var saKey ServiceAccountKey
	if err := json.Unmarshal(keyData, &saKey); err != nil {
		return nil, fmt.Errorf("failed to parse service account key: %w", err)
	}
	if saKey.Type != "service_account" {
		return nil, fmt.Errorf("invalid service account key type: %s", saKey.Type)
	}
	if saKey.ClientEmail == "" {
		return nil, fmt.Errorf("missing client_email in service account key")
	}
	if saKey.PrivateKey == "" {
		return nil, fmt.Errorf("missing private_key in service account key")
	}

	var config *google.Credentials
	if impersonateUser != "" {
		config, err = google.CredentialsFromJSONWithParams(ctx, keyData, google.CredentialsParams{
			Scopes:  scopes,
			Subject: impersonateUser,
		})
	} else {
		config, err = google.CredentialsFromJSON(ctx, keyData, scopes...)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to parse service account key: %w", err)
	}

	token, err := config.TokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	authType := types.AuthTypeServiceAccount
	if impersonateUser != "" {
		authType = types.AuthTypeImpersonated
	}

	return &types.Credentials{
		AccessToken:         token.AccessToken,
		ExpiryDate:          token.Expiry,
		Scopes:              scopes,
		Type:                authType,
		ServiceAccountEmail: saKey.ClientEmail,
		ImpersonatedUser:    impersonateUser,
	}, nil
}

// GetDriveService creates a Drive API service from credentials
func (m *Manager) GetDriveService(ctx context.Context, creds *types.Credentials) (*drive.Service, error) {
	client := m.GetHTTPClient(ctx, creds)
	return drive.NewService(ctx, option.WithHTTPClient(client))
}
