package auth

import (
	"bufio"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/dl-alexandre/gdrv/internal/types"
	"golang.org/x/oauth2"
)

// OAuthFlow handles the OAuth2 authentication flow
type OAuthFlow struct {
	config       *oauth2.Config
	listener     net.Listener
	redirectURL  string
	state        string
	codeVerifier string
	codeChan     chan string
	errChan      chan error
}

// NewOAuthFlow creates a new OAuth flow handler
func NewOAuthFlow(config *oauth2.Config, listener net.Listener, redirectURL string) (*OAuthFlow, error) {
	if config == nil {
		return nil, fmt.Errorf("OAuth config not set")
	}

	state, err := generateState()
	if err != nil {
		return nil, fmt.Errorf("failed to generate state: %w", err)
	}

	verifier, err := generateCodeVerifier()
	if err != nil {
		return nil, fmt.Errorf("failed to generate code verifier: %w", err)
	}

	cfg := *config
	if redirectURL != "" {
		cfg.RedirectURL = redirectURL
	}
	if cfg.RedirectURL == "" {
		return nil, fmt.Errorf("redirect URL not set")
	}

	return &OAuthFlow{
		config:       &cfg,
		listener:     listener,
		redirectURL:  cfg.RedirectURL,
		state:        state,
		codeVerifier: verifier,
		codeChan:     make(chan string, 1),
		errChan:      make(chan error, 1),
	}, nil
}

// GetAuthURL returns the URL to redirect user for authentication
func (f *OAuthFlow) GetAuthURL() string {
	return f.config.AuthCodeURL(
		f.state,
		oauth2.AccessTypeOffline,
		oauth2.SetAuthURLParam("code_challenge", codeChallengeS256(f.codeVerifier)),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
	)
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
		server.Close()
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
	token, err := f.config.Exchange(
		ctx,
		code,
		oauth2.SetAuthURLParam("code_verifier", f.codeVerifier),
	)
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
func (f *OAuthFlow) Close() {
	if f.listener != nil {
		f.listener.Close()
	}
}

func generateState() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func generateCodeVerifier() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func codeChallengeS256(verifier string) string {
	sum := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(sum[:])
}

func promptForAuthCode(reader *bufio.Reader) (string, error) {
	fmt.Print("Paste the authorization code from the redirected URL: ")
	code, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(code), nil
}

// Authenticate performs the full OAuth flow
func (m *Manager) Authenticate(ctx context.Context, profile string, openBrowser func(string) error, opts OAuthAuthOptions) (*types.Credentials, error) {
	if m.oauthConfig == nil {
		return nil, fmt.Errorf("OAuth config not set")
	}

	var flow *OAuthFlow
	manualFallback := opts.NoBrowser || isHeadlessEnv()
	if !manualFallback {
		var err error
		flow, err = newLoopbackFlow(m.oauthConfig)
		if err != nil {
			manualFallback = true
		}
	}

	if flow != nil && !manualFallback {
		defer flow.Close()
	}

	var authURL string
	if flow != nil {
		authURL = flow.GetAuthURL()
	}

	if manualFallback {
		if flow != nil {
			flow.Close()
		}
		flow, err := newManualFlow(m.oauthConfig)
		if err != nil {
			return nil, err
		}
		authURL = flow.GetAuthURL()
		fmt.Printf("Manual authentication required.\n")
		fmt.Printf("Open this URL in a browser and approve access:\n%s\n", authURL)
		fmt.Printf("After approval, you will be redirected to a localhost URL.\n")
		fmt.Printf("Copy the `code` parameter from the address bar and paste it here.\n")
		code, err := promptForAuthCode(bufio.NewReader(os.Stdin))
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

	fmt.Printf("Opening browser for authentication...\n")
	fmt.Printf("If browser doesn't open, visit: %s\n", authURL)

	flow.StartCallbackServer(ctx)

	if err := openBrowser(authURL); err != nil {
		fmt.Printf("Failed to open browser: %v\n", err)
		fmt.Printf("Switching to manual authentication.\n")
		opts.NoBrowser = true
		return m.Authenticate(ctx, profile, openBrowser, opts)
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

// OAuthAuthOptions controls OAuth authentication behavior.
type OAuthAuthOptions struct {
	NoBrowser bool
}

func newLoopbackFlow(config *oauth2.Config) (*OAuthFlow, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, fmt.Errorf("failed to start local server: %w", err)
	}
	addr := listener.Addr().(*net.TCPAddr)
	redirectURL := fmt.Sprintf("http://127.0.0.1:%d/callback", addr.Port)
	return NewOAuthFlow(config, listener, redirectURL)
}

func newManualFlow(config *oauth2.Config) (*OAuthFlow, error) {
	port := pickManualPort()
	redirectURL := fmt.Sprintf("http://127.0.0.1:%d/callback", port)
	return NewOAuthFlow(config, nil, redirectURL)
}

func pickManualPort() int {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err == nil {
		addr := listener.Addr().(*net.TCPAddr)
		_ = listener.Close()
		return addr.Port
	}
	return 8765
}

func isHeadlessEnv() bool {
	if os.Getenv("GDRV_NO_BROWSER") != "" {
		return true
	}
	if os.Getenv("CI") != "" || os.Getenv("GITHUB_ACTIONS") != "" {
		return true
	}
	if runtime.GOOS != "windows" && os.Getenv("DISPLAY") == "" && os.Getenv("WAYLAND_DISPLAY") == "" {
		return true
	}
	if os.Getenv("SSH_CONNECTION") != "" || os.Getenv("SSH_TTY") != "" {
		return true
	}
	return false
}
