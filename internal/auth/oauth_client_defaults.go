package auth

// BundledOAuthClientID and BundledOAuthClientSecret can be set at build time
// via -ldflags. If unset, the CLI requires a custom OAuth client.
var (
	BundledOAuthClientID     string
	BundledOAuthClientSecret string
)

// GetBundledOAuthClient returns the bundled OAuth client credentials.
func GetBundledOAuthClient() (string, string, bool) {
	if BundledOAuthClientID == "" {
		return "", "", false
	}
	return BundledOAuthClientID, BundledOAuthClientSecret, true
}
