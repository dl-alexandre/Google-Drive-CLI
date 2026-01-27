package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/dl-alexandre/gdrv/internal/types"
	"golang.org/x/oauth2"
)

// OAuthFlow handles the OAuth2 authentication flow
type OAuthFlow struct {
	config   *oauth2.Config
	listener net.Listener
	state    string
	codeChan chan string
	errChan  chan error
}

// NewOAuthFlow creates a new OAuth flow handler
func NewOAuthFlow(config *oauth2.Config) (*OAuthFlow, error) {
	listener, err := net.Listen("tcp", "localhost:8085")
	if err != nil {
		return nil, fmt.Errorf("failed to start local server: %w", err)
	}

	state, err := generateState()
	if err != nil {
		if closeErr := listener.Close(); closeErr != nil {
			return nil, fmt.Errorf("failed to close listener after state error: %w", closeErr)
		}
		return nil, fmt.Errorf("failed to generate state: %w", err)
	}

	return &OAuthFlow{
		config:   config,
		listener: listener,
		state:    state,
		codeChan: make(chan string, 1),
		errChan:  make(chan error, 1),
	}, nil
}

// GetAuthURL returns the URL to redirect user for authentication
func (f *OAuthFlow) GetAuthURL() string {
	return f.config.AuthCodeURL(f.state, oauth2.AccessTypeOffline)
}

// StartCallbackServer starts the callback server and waits for auth code
func (f *OAuthFlow) StartCallbackServer(ctx context.Context) {
	mux := http.NewServeMux()
	mux.HandleFunc("/callback", f.handleCallback)

	server := &http.Server{Handler: mux}
	go func() {
		if err := server.Serve(f.listener); err != http.ErrServerClosed {
			f.errChan <- err
		}
	}()

	go func() {
		<-ctx.Done()
		if err := server.Close(); err != nil && err != http.ErrServerClosed {
			f.errChan <- err
		}
	}()
}

// handleCallback processes the OAuth callback
func (f *OAuthFlow) handleCallback(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("state") != f.state {
		f.errChan <- fmt.Errorf("invalid state parameter")
		http.Error(w, "Invalid state", http.StatusBadRequest)
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		errMsg := r.URL.Query().Get("error")
		f.errChan <- fmt.Errorf("auth error: %s", errMsg)
		http.Error(w, "No code received", http.StatusBadRequest)
		return
	}

	f.codeChan <- code
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<html><body><h1>Authentication successful!</h1><p>You can close this window.</p></body></html>`)
}

// WaitForCode waits for the authorization code
func (f *OAuthFlow) WaitForCode(timeout time.Duration) (string, error) {
	select {
	case code := <-f.codeChan:
		return code, nil
	case err := <-f.errChan:
		return "", err
	case <-time.After(timeout):
		return "", fmt.Errorf("authentication timed out")
	}
}

// ExchangeCode exchanges auth code for tokens
func (f *OAuthFlow) ExchangeCode(ctx context.Context, code string) (*types.Credentials, error) {
	token, err := f.config.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code: %w", err)
	}

	return &types.Credentials{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiryDate:   token.Expiry,
		Scopes:       f.config.Scopes,
		Type:         types.AuthTypeOAuth,
	}, nil
}

// Close cleans up resources
func (f *OAuthFlow) Close() error {
	if f.listener != nil {
		if err := f.listener.Close(); err != nil {
			return err
		}
	}
	return nil
}

func generateState() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// Authenticate performs the full OAuth flow
func (m *Manager) Authenticate(ctx context.Context, profile string, openBrowser func(string) error) (*types.Credentials, error) {
	if m.oauthConfig == nil {
		return nil, fmt.Errorf("OAuth config not set")
	}

	flow, err := NewOAuthFlow(m.oauthConfig)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := flow.Close(); err != nil {
			fmt.Printf("Warning: failed to close OAuth listener: %v\n", err)
		}
	}()

	authURL := flow.GetAuthURL()
	fmt.Printf("Opening browser for authentication...\n")
	fmt.Printf("If browser doesn't open, visit: %s\n", authURL)

	flow.StartCallbackServer(ctx)

	if err := openBrowser(authURL); err != nil {
		fmt.Printf("Failed to open browser: %v\n", err)
	}

	code, err := flow.WaitForCode(5 * time.Minute)
	if err != nil {
		return nil, err
	}

	creds, err := flow.ExchangeCode(ctx, code)
	if err != nil {
		return nil, err
	}

	if err := m.SaveCredentials(profile, creds); err != nil {
		return nil, fmt.Errorf("failed to save credentials: %w", err)
	}

	return creds, nil
}
