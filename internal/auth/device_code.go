package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/dl-alexandre/gdrv/internal/types"
	"golang.org/x/oauth2"
)

// DeviceCodeResponse represents the response from device authorization endpoint
type DeviceCodeResponse struct {
	DeviceCode              string `json:"device_code"`
	UserCode                string `json:"user_code"`
	VerificationURL         string `json:"verification_url"`
	VerificationURLComplete string `json:"verification_uri_complete"`
	ExpiresIn               int    `json:"expires_in"`
	Interval                int    `json:"interval"`
}

// TokenResponse represents the token response from polling
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	Error        string `json:"error,omitempty"`
}

const (
	deviceCodeEndpoint = "https://oauth2.googleapis.com/device/code"
	tokenEndpoint      = "https://oauth2.googleapis.com/token"
)

// DeviceCodeFlow handles device code authentication flow
type DeviceCodeFlow struct {
	config   *oauth2.Config
	response *DeviceCodeResponse
}

// NewDeviceCodeFlow creates a new device code flow handler
func NewDeviceCodeFlow(config *oauth2.Config) *DeviceCodeFlow {
	return &DeviceCodeFlow{
		config: config,
	}
}

// RequestDeviceCode requests a device code from Google
func (f *DeviceCodeFlow) RequestDeviceCode(ctx context.Context) (response *DeviceCodeResponse, err error) {
	data := url.Values{}
	data.Set("client_id", f.config.ClientID)
	data.Set("scope", strings.Join(f.config.Scopes, " "))

	req, err := http.NewRequestWithContext(ctx, "POST", deviceCodeEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to request device code: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("failed to close device code response body: %w", closeErr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("device code request failed: %s - %s", resp.Status, string(body))
	}

	var deviceResp DeviceCodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&deviceResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	f.response = &deviceResp
	return &deviceResp, nil
}

// PollForToken polls the token endpoint until user completes authorization
func (f *DeviceCodeFlow) PollForToken(ctx context.Context) (*types.Credentials, error) {
	if f.response == nil {
		return nil, fmt.Errorf("device code not requested yet")
	}

	interval := time.Duration(f.response.Interval) * time.Second
	if interval < 5*time.Second {
		interval = 5 * time.Second
	}

	expiry := time.Now().Add(time.Duration(f.response.ExpiresIn) * time.Second)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	client := &http.Client{Timeout: 30 * time.Second}

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			if time.Now().After(expiry) {
				return nil, fmt.Errorf("device code expired")
			}

			token, err := f.pollOnce(ctx, client)
			if err != nil {
				// Continue polling on specific errors
				if strings.Contains(err.Error(), "authorization_pending") || 
				   strings.Contains(err.Error(), "slow_down") {
					continue
				}
				return nil, err
			}

			if token != nil {
				return token, nil
			}
		}
	}
}

// pollOnce performs a single poll attempt
func (f *DeviceCodeFlow) pollOnce(ctx context.Context, client *http.Client) (creds *types.Credentials, err error) {
	data := url.Values{}
	data.Set("client_id", f.config.ClientID)
	data.Set("client_secret", f.config.ClientSecret)
	data.Set("device_code", f.response.DeviceCode)
	data.Set("grant_type", "urn:ietf:params:oauth:grant-type:device_code")

	req, err := http.NewRequestWithContext(ctx, "POST", tokenEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to poll for token: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("failed to close token response body: %w", closeErr)
		}
	}()

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Handle errors
	if tokenResp.Error != "" {
		switch tokenResp.Error {
		case "authorization_pending":
			return nil, fmt.Errorf("authorization_pending")
		case "slow_down":
			return nil, fmt.Errorf("slow_down")
		case "expired_token":
			return nil, fmt.Errorf("device code has expired")
		case "access_denied":
			return nil, fmt.Errorf("user denied authorization")
		default:
			return nil, fmt.Errorf("token error: %s", tokenResp.Error)
		}
	}

	// Success - we have a token
	if tokenResp.AccessToken != "" {
		scopes := strings.Split(tokenResp.Scope, " ")
		expiryDate := time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)

		return &types.Credentials{
			AccessToken:  tokenResp.AccessToken,
			RefreshToken: tokenResp.RefreshToken,
			ExpiryDate:   expiryDate,
			Scopes:       scopes,
			Type:         types.AuthTypeOAuth,
		}, nil
	}

	return nil, nil
}

// AuthenticateWithDeviceCode performs device code authentication flow
func (m *Manager) AuthenticateWithDeviceCode(ctx context.Context, profile string) (*types.Credentials, error) {
	if m.oauthConfig == nil {
		return nil, fmt.Errorf("OAuth config not set")
	}

	flow := NewDeviceCodeFlow(m.oauthConfig)

	// Request device code
	deviceResp, err := flow.RequestDeviceCode(ctx)
	if err != nil {
		return nil, err
	}

	// Display instructions to user
	fmt.Printf("\nDevice Code Authentication\n")
	fmt.Printf("==========================\n\n")
	fmt.Printf("Please visit the following URL and enter the code:\n\n")
	fmt.Printf("URL:  %s\n", deviceResp.VerificationURL)
	fmt.Printf("Code: %s\n\n", deviceResp.UserCode)
	
	if deviceResp.VerificationURLComplete != "" {
		fmt.Printf("Or visit this URL to auto-fill the code:\n")
		fmt.Printf("%s\n\n", deviceResp.VerificationURLComplete)
	}
	
	fmt.Printf("Waiting for authorization (expires in %d seconds)...\n", deviceResp.ExpiresIn)

	// Poll for token
	creds, err := flow.PollForToken(ctx)
	if err != nil {
		return nil, err
	}

	// Save credentials
	if err := m.SaveCredentials(profile, creds); err != nil {
		return nil, fmt.Errorf("failed to save credentials: %w", err)
	}

	return creds, nil
}
