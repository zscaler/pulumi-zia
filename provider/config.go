package provider

import (
	"context"
	"os"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/pulumi-zia/provider/internal/zia"
	"github.com/zscaler/pulumi-zia/provider/version"
)

// Config defines the ZIA provider configuration.
// Implements infer.CustomConfigure to initialize the Zscaler client.
type Config struct {
	// OAuth2
	ClientID     *string `pulumi:"clientId,optional"`
	ClientSecret *string `pulumi:"clientSecret,optional" provider:"secret"`
	PrivateKey   *string `pulumi:"privateKey,optional" provider:"secret"`
	VanityDomain *string `pulumi:"vanityDomain,optional" provider:"secret"`
	Cloud        *string `pulumi:"cloud,optional" provider:"secret"`

	// Legacy auth
	Username *string `pulumi:"username,optional"`
	Password *string `pulumi:"password,optional" provider:"secret"`
	APIKey   *string `pulumi:"apiKey,optional" provider:"secret"`
	ZIACloud *string `pulumi:"ziaCloud,optional"`

	// Sandbox
	SandboxToken *string `pulumi:"sandboxToken,optional" provider:"secret"`
	SandboxCloud *string `pulumi:"sandboxCloud,optional" provider:"secret"`

	// Optional
	HTTPProxy       *string `pulumi:"httpProxy,optional"`
	MaxRetries      *int    `pulumi:"maxRetries,optional"`
	RequestTimeout  *int    `pulumi:"requestTimeout,optional"`
	UseLegacyClient *bool   `pulumi:"useLegacyClient,optional"`

	// Debug: when true, enables Zscaler SDK logging (API requests/responses).
	// Logs go to stderr and optionally to ZSCALER_SDK_LOG_FILE.
	// Equivalent to setting ZSCALER_SDK_LOG=true + ZSCALER_SDK_VERBOSE=true env vars.
	Debug *bool `pulumi:"debug,optional"`

	// Internal: set during Configure, not serialized
	client *zia.Client
}

func (c *Config) Annotate(a infer.Annotator) {
	a.Describe(&c.ClientID, "The OAuth2 client ID for authenticating with the Zscaler API. Can also be set via the `ZSCALER_CLIENT_ID` environment variable.")
	a.Describe(&c.ClientSecret, "The OAuth2 client secret for authenticating with the Zscaler API. Can also be set via the `ZSCALER_CLIENT_SECRET` environment variable.")
	a.Describe(&c.PrivateKey, "The private key for service principal authentication. Can also be set via the `ZSCALER_PRIVATE_KEY` environment variable.")
	a.Describe(&c.VanityDomain, "The vanity domain for your Zscaler organization. Can also be set via the `ZSCALER_VANITY_DOMAIN` environment variable.")
	a.Describe(&c.Cloud, "The Zscaler cloud name (e.g. 'zscaler', 'zscalerone', 'zscalertwo', 'zscalerthree', 'zscloud', 'zscalerbeta', 'zscalergov'). Can also be set via the `ZSCALER_CLOUD` environment variable.")
	a.Describe(&c.Username, "(Legacy) The admin username for ZIA. Can also be set via the `ZIA_USERNAME` environment variable. Prefer OAuth2 credentials instead.")
	a.Describe(&c.Password, "(Legacy) The admin password for ZIA. Can also be set via the `ZIA_PASSWORD` environment variable. Prefer OAuth2 credentials instead.")
	a.Describe(&c.APIKey, "(Legacy) The API key for ZIA. Can also be set via the `ZIA_API_KEY` environment variable. Prefer OAuth2 credentials instead.")
	a.Describe(&c.ZIACloud, "(Legacy) The ZIA cloud name. Can also be set via the `ZIA_CLOUD` environment variable. Prefer the 'cloud' parameter instead.")
	a.Describe(&c.SandboxToken, "The API token for Zscaler Sandbox. Can also be set via the `ZIA_SANDBOX_TOKEN` environment variable.")
	a.Describe(&c.SandboxCloud, "The Zscaler Sandbox cloud name. Can also be set via the `ZIA_SANDBOX_CLOUD` environment variable.")
	a.Describe(&c.HTTPProxy, "HTTP proxy URL for API requests (e.g. 'http://proxy.example.com:8080'). Can also be set via the `ZSCALER_HTTP_PROXY` environment variable.")
	a.Describe(&c.MaxRetries, "Maximum number of retries for API requests. Default is determined by the SDK.")
	a.Describe(&c.RequestTimeout, "Timeout in seconds for API requests.")
	a.Describe(&c.UseLegacyClient, "If true, use the legacy ZIA client authentication instead of OAuth2.")
	a.Describe(&c.Debug, "If true, enables verbose Zscaler SDK logging (API requests/responses). Logs are written to stderr and optionally to the file specified by the `ZSCALER_SDK_LOG_FILE` environment variable.")
}

// Configure initializes the Zscaler ZIA client from provider config.
// If credentials are missing, client remains nil (operations will fail with a clear error).
func (c *Config) Configure(ctx context.Context) error {
	cfg := &zia.Config{}
	if c.ClientID != nil {
		cfg.ClientID = *c.ClientID
	}
	if c.ClientSecret != nil {
		cfg.ClientSecret = *c.ClientSecret
	}
	if c.PrivateKey != nil {
		cfg.PrivateKey = *c.PrivateKey
	}
	if c.VanityDomain != nil {
		cfg.VanityDomain = *c.VanityDomain
	}
	if c.Cloud != nil {
		cfg.Cloud = *c.Cloud
	}
	if c.Username != nil {
		cfg.Username = *c.Username
	}
	if c.Password != nil {
		cfg.Password = *c.Password
	}
	if c.APIKey != nil {
		cfg.APIKey = *c.APIKey
	}
	if c.ZIACloud != nil {
		cfg.ZIABaseURL = *c.ZIACloud
	}
	if c.SandboxToken != nil {
		cfg.SandboxToken = *c.SandboxToken
	}
	if c.SandboxCloud != nil {
		cfg.SandboxCloud = *c.SandboxCloud
	}
	if c.HTTPProxy != nil {
		cfg.HTTPProxy = *c.HTTPProxy
	}
	if c.MaxRetries != nil {
		cfg.RetryCount = *c.MaxRetries
	}
	if c.RequestTimeout != nil {
		cfg.RequestTimeout = *c.RequestTimeout
	}
	if c.UseLegacyClient != nil {
		cfg.UseLegacyClient = *c.UseLegacyClient
	}
	if c.Debug != nil && *c.Debug {
		cfg.Debug = true
		_ = os.Setenv("ZSCALER_SDK_LOG", "true")
		_ = os.Setenv("ZSCALER_SDK_VERBOSE", "true")
	}

	sdkLogger := SetupProviderLogging()
	client, err := zia.NewClient(cfg, version.Version, sdkLogger)
	if err != nil {
		// Allow missing credentials for dry-run/preview; operations will fail with clear error when needed
		c.client = nil
		return nil
	}
	c.client = client
	return nil
}

// Diff prevents provider version upgrades from triggering resource replacement.
// A version-only change in the provider config is never a breaking change.
func (c *Config) Diff(_ context.Context, _ infer.DiffRequest[Config, Config]) (p.DiffResponse, error) {
	return p.DiffResponse{
		HasChanges:          false,
		DeleteBeforeReplace: false,
	}, nil
}

// Client returns the configured ZIA client. Must be called after Configure.
func (c *Config) Client() *zia.Client {
	return c.client
}
