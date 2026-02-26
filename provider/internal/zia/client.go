// Package zia provides the Zscaler ZIA SDK client wrapper for the Pulumi provider.
package zia

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/zscaler/pulumi-zia/provider/version"
	"github.com/zscaler/zscaler-sdk-go/v3/logger"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler"
)

// Client wraps the Zscaler SDK Service for ZIA API operations.
type Client struct {
	Service *zscaler.Service
}

// Config holds the configuration needed to create a Zscaler ZIA client.
type Config struct {
	// OAuth2 / API Key auth
	ClientID     string
	ClientSecret string
	PrivateKey   string
	VanityDomain string
	Cloud        string

	// Legacy auth (username/password/apiKey + zia_cloud)
	Username   string
	Password   string
	APIKey     string
	ZIABaseURL string

	// Sandbox
	SandboxToken string
	SandboxCloud string

	// Optional
	HTTPProxy       string
	RetryCount      int
	RequestTimeout  int
	UseLegacyClient bool
	Debug           bool

	// Internal - set during configure
	terraformVersion string
}

// NewConfigFromEnv populates config from environment variables as fallback.
func NewConfigFromEnv(c *Config) {
	if c.ClientID == "" && os.Getenv("ZSCALER_CLIENT_ID") != "" {
		c.ClientID = os.Getenv("ZSCALER_CLIENT_ID")
	}
	if c.ClientSecret == "" && os.Getenv("ZSCALER_CLIENT_SECRET") != "" {
		c.ClientSecret = os.Getenv("ZSCALER_CLIENT_SECRET")
	}
	if c.PrivateKey == "" && os.Getenv("ZSCALER_PRIVATE_KEY") != "" {
		c.PrivateKey = os.Getenv("ZSCALER_PRIVATE_KEY")
	}
	if c.VanityDomain == "" && os.Getenv("ZSCALER_VANITY_DOMAIN") != "" {
		c.VanityDomain = os.Getenv("ZSCALER_VANITY_DOMAIN")
	}
	if c.Cloud == "" && os.Getenv("ZSCALER_CLOUD") != "" {
		c.Cloud = os.Getenv("ZSCALER_CLOUD")
	}
	if c.SandboxToken == "" && os.Getenv("ZSCALER_SANDBOX_TOKEN") != "" {
		c.SandboxToken = os.Getenv("ZSCALER_SANDBOX_TOKEN")
	}
	if c.SandboxCloud == "" && os.Getenv("ZSCALER_SANDBOX_CLOUD") != "" {
		c.SandboxCloud = os.Getenv("ZSCALER_SANDBOX_CLOUD")
	}
	if c.Username == "" {
		c.Username = os.Getenv("ZIA_USERNAME")
	}
	if c.Password == "" {
		c.Password = os.Getenv("ZIA_PASSWORD")
	}
	if c.APIKey == "" {
		c.APIKey = os.Getenv("ZIA_API_KEY")
	}
	if c.ZIABaseURL == "" {
		c.ZIABaseURL = os.Getenv("ZIA_CLOUD")
	}
	if c.HTTPProxy == "" && os.Getenv("ZSCALER_HTTP_PROXY") != "" {
		c.HTTPProxy = os.Getenv("ZSCALER_HTTP_PROXY")
	}
	if strings.ToLower(os.Getenv("ZSCALER_USE_LEGACY_CLIENT")) == "true" {
		c.UseLegacyClient = true
	}
}

// UserAgent generates a user agent string in the form:
//
//	(darwin arm64) Pulumi/v3.212.0 ZIAProvider/v1.2.0
//
// The Pulumi SDK version is read from the Go module build info at runtime.
// In test binaries where ReadBuildInfo().Deps is empty, the fallback
// constant from provider/version is used instead.
func UserAgent(providerVersion string) string {
	if providerVersion == "" {
		providerVersion = "0.0.0-dev"
	}

	pulumiVersion := ""
	if bi, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range bi.Deps {
			if dep.Path == "github.com/pulumi/pulumi/sdk/v3" {
				pulumiVersion = dep.Version
				break
			}
		}
	}
	if pulumiVersion == "" {
		pulumiVersion = version.PulumiSDKVersion
	}

	return fmt.Sprintf("(%s %s) Pulumi/%s ZIAProvider/%s",
		runtime.GOOS, runtime.GOARCH, pulumiVersion, providerVersion)
}

// NewClient creates a ZIA client from config.
// If sdkLogger is non-nil it replaces the SDK's built-in stdout logger,
// which is critical in Pulumi where stdout is the gRPC channel.
func NewClient(cfg *Config, providerVersion string, sdkLogger logger.Logger) (*Client, error) {
	NewConfigFromEnv(cfg)

	// Defaults
	if cfg.RetryCount == 0 {
		cfg.RetryCount = 100
	}
	if cfg.RequestTimeout == 0 {
		cfg.RequestTimeout = 1800 // 30 min for large GetAll operations
	}

	userAgent := UserAgent(providerVersion)

	if cfg.UseLegacyClient {
		// Legacy V2 SDK - would require zia package from zscaler-sdk-go
		return nil, fmt.Errorf("legacy ZIA client (use_legacy_client) is not yet supported in the native Pulumi provider; use OAuth2 or API key auth")
	}

	// V3 SDK setup
	setters := []zscaler.ConfigSetter{
		zscaler.WithCache(true),
		zscaler.WithCacheTtl(10 * time.Minute),
	}
	if cfg.Debug {
		setters = append(setters, zscaler.WithDebug(true))
	}
	setters = append(setters,
		zscaler.WithCacheTti(8*time.Minute),
		zscaler.WithHttpClientPtr(http.DefaultClient),
		zscaler.WithRateLimitMaxRetries(int32(cfg.RetryCount)),
		zscaler.WithRequestTimeout(time.Duration(cfg.RequestTimeout)*time.Second),
		zscaler.WithRateLimitMinWait(2*time.Second),
		zscaler.WithRateLimitMaxWait(10*time.Second),
		zscaler.WithUserAgentExtra(userAgent),
	)

	if cfg.HTTPProxy != "" {
		u, err := url.Parse(cfg.HTTPProxy)
		if err != nil {
			return nil, fmt.Errorf("invalid proxy URL: %w", err)
		}
		setters = append(setters, zscaler.WithProxyHost(u.Hostname()))
		portStr := u.Port()
		if portStr == "" {
			portStr = "80"
		}
		port, err := strconv.ParseInt(portStr, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid proxy port: %w", err)
		}
		if port < 1 || port > 65535 {
			return nil, fmt.Errorf("invalid port: must be 1-65535")
		}
		setters = append(setters, zscaler.WithProxyPort(int32(port)))
	}

	// Sandbox-only auth
	if cfg.SandboxToken != "" && cfg.SandboxCloud != "" && cfg.ClientID == "" && cfg.ClientSecret == "" && cfg.PrivateKey == "" {
		setters = append(setters,
			zscaler.WithSandboxToken(cfg.SandboxToken),
			zscaler.WithSandboxCloud(cfg.SandboxCloud),
		)
		config, err := zscaler.NewConfiguration(setters...)
		if err != nil {
			return nil, fmt.Errorf("failed to create SDK config for Sandbox: %w", err)
		}
		config.UserAgent = userAgent
		if sdkLogger != nil {
			config.Logger = sdkLogger
		}
		v3Client, err := zscaler.NewOneAPIClient(config)
		if err != nil {
			return nil, fmt.Errorf("failed to create Sandbox client: %w", err)
		}
		return &Client{Service: zscaler.NewService(v3Client.Client, nil)}, nil
	}

	// OAuth2 auth
	switch {
	case cfg.ClientID != "" && cfg.ClientSecret != "" && cfg.VanityDomain != "":
		setters = append(setters,
			zscaler.WithClientID(cfg.ClientID),
			zscaler.WithClientSecret(cfg.ClientSecret),
			zscaler.WithVanityDomain(cfg.VanityDomain),
			zscaler.WithSandboxToken(cfg.SandboxToken),
			zscaler.WithSandboxCloud(cfg.SandboxCloud),
		)
		if cfg.Cloud != "" {
			setters = append(setters, zscaler.WithZscalerCloud(cfg.Cloud))
		}
	case cfg.ClientID != "" && cfg.PrivateKey != "" && cfg.VanityDomain != "":
		setters = append(setters,
			zscaler.WithClientID(cfg.ClientID),
			zscaler.WithPrivateKey(cfg.PrivateKey),
			zscaler.WithVanityDomain(cfg.VanityDomain),
			zscaler.WithSandboxToken(cfg.SandboxToken),
			zscaler.WithSandboxCloud(cfg.SandboxCloud),
		)
		if cfg.Cloud != "" {
			setters = append(setters, zscaler.WithZscalerCloud(cfg.Cloud))
		}
	default:
		return nil, fmt.Errorf("invalid authentication: provide (client_id + client_secret + vanity_domain) or (client_id + private_key + vanity_domain), or set ZSCALER_* env vars")
	}

	config, err := zscaler.NewConfiguration(setters...)
	if err != nil {
		return nil, fmt.Errorf("failed to create SDK config: %w", err)
	}
	config.UserAgent = userAgent
	if sdkLogger != nil {
		config.Logger = sdkLogger
	}
	v3Client, err := zscaler.NewOneAPIClient(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create ZIA client: %w", err)
	}
	return &Client{Service: zscaler.NewService(v3Client.Client, nil)}, nil
}
