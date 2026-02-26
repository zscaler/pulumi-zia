// Copyright (c) 2023 Zscaler Technology Alliances, <devrel@zscaler.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Integration / E2E tests that exercise real ZIA API calls.
// Excluded from unit-test runs via the build tag below.

//go:build !unit

package tests

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pulumi/providertest/providers"
	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/providertest/pulumitest/opttest"
	goprovider "github.com/pulumi/pulumi-go-provider"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"

	"github.com/zscaler/pulumi-zia/provider"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/trafficforwarding/extranet"
)

// ziaProviderFactory creates an in-process ZIA provider server for pulumitest.
var ziaProviderFactory = func(_ providers.PulumiTest) (pulumirpc.ResourceProviderServer, error) {
	return goprovider.RawServer(provider.Name, "1.0.0-dev", provider.Provider())(nil)
}

// newZIATest creates a PulumiTest instance pointing at a test-program directory
// with the ZIA provider attached in-process. Credentials come from environment
// variables (ZSCALER_CLIENT_ID, ZSCALER_CLIENT_SECRET, ZSCALER_VANITY_DOMAIN).
//
// Tests are skipped automatically when:
//   - testing.Short() is set (CI runs without cloud credentials)
//   - Required environment variables are missing
func newZIATest(t *testing.T, testProgramDir string, opts ...opttest.Option) *pulumitest.PulumiTest {
	t.Helper()

	if testing.Short() {
		t.Skipf("Skipping E2E test in short mode (no cloud credentials)")
		return nil
	}

	requiredEnvVars := []string{
		"ZSCALER_CLIENT_ID",
		"ZSCALER_CLIENT_SECRET",
		"ZSCALER_VANITY_DOMAIN",
	}
	for _, env := range requiredEnvVars {
		if os.Getenv(env) == "" {
			t.Skipf("Skipping E2E test: %s not set", env)
			return nil
		}
	}

	dir := filepath.Join("test-programs", testProgramDir)
	attachOpt := opttest.AttachProviderServer(provider.Name, ziaProviderFactory)

	allOpts := append([]opttest.Option{attachOpt, opttest.SkipInstall()}, opts...)
	pt := pulumitest.NewPulumiTest(t, dir, allOpts...)

	return pt
}

// uniqueName generates a test-specific resource name to avoid collisions when tests
// run in parallel or when leftover resources exist from failed runs.
// Format: test-<8 random alphanumeric>-<resource>, e.g. test-k7j2m9ab-bandwidth-class
func uniqueName(prefix string) string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		panic(fmt.Sprintf("uniqueName rand.Read: %v", err))
	}
	suffix := make([]byte, 8)
	for i := range b {
		suffix[i] = letters[int(b[i])%len(letters)]
	}
	return fmt.Sprintf("test-%s-%s", string(suffix), prefix)
}

// shortUniqueName returns a random name under 31 chars for ZIA resources with strict length limits.
// Same pattern as Terraform: tf-acc-test- + 10 alpha chars (22 chars total).
func shortUniqueName() string {
	const alpha = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, 10)
	if _, err := rand.Read(b); err != nil {
		panic(fmt.Sprintf("shortUniqueName rand.Read: %v", err))
	}
	suffix := make([]byte, 10)
	for i := range b {
		suffix[i] = alpha[int(b[i])%len(alpha)]
	}
	return fmt.Sprintf("tf-acc-test-%s", string(suffix))
}

// randomIPFromCIDR generates a random IP address within the given CIDR block.
// Used for TrafficForwardingStaticIp tests (avoid DUPLICATE_ITEM on IP).
func randomIPFromCIDR(prefix string) string {
	// prefix like "104.238.235" - we randomize the last octet (1-254)
	b := make([]byte, 1)
	if _, err := rand.Read(b); err != nil {
		panic(fmt.Sprintf("randomIPFromCIDR: %v", err))
	}
	last := 1 + int(b[0])%254 // 1-254
	return fmt.Sprintf("%s.%d", prefix, last)
}

// toSlice converts a Pulumi output value (often []interface{}) to []interface{} for assertion.
func toSlice(v interface{}) []interface{} {
	if v == nil {
		return nil
	}
	if s, ok := v.([]interface{}); ok {
		return s
	}
	switch val := v.(type) {
	case []string:
		out := make([]interface{}, len(val))
		for i, x := range val {
			out[i] = x
		}
		return out
	default:
		return []interface{}{v}
	}
}

// newTestZIAService creates a one-off ZIA SDK service from environment variables.
// Used for pre-test cleanup of stale resources.
func newTestZIAService(t *testing.T) *zscaler.Service {
	t.Helper()
	config, err := zscaler.NewConfiguration(
		zscaler.WithClientID(os.Getenv("ZSCALER_CLIENT_ID")),
		zscaler.WithClientSecret(os.Getenv("ZSCALER_CLIENT_SECRET")),
		zscaler.WithVanityDomain(os.Getenv("ZSCALER_VANITY_DOMAIN")),
	)
	if err != nil {
		t.Fatalf("newTestZIAService: config: %v", err)
	}
	client, err := zscaler.NewOneAPIClient(config)
	if err != nil {
		t.Fatalf("newTestZIAService: client: %v", err)
	}
	return zscaler.NewService(client.Client, nil)
}

// sweepStaleExtranets deletes all extranets whose name starts with "tf-acc-test-" or
// "tf-updated-", cleaning up leftovers from previous failed test runs.
func sweepStaleExtranets(t *testing.T, service *zscaler.Service) {
	t.Helper()
	ctx := context.Background()
	all, err := extranet.GetAll(ctx, service, nil)
	if err != nil {
		t.Logf("[SWEEP] Warning: could not list extranets: %v", err)
		return
	}
	for _, e := range all {
		if strings.HasPrefix(e.Name, "tf-acc-test-") || strings.HasPrefix(e.Name, "tf-updated-") {
			log.Printf("[SWEEP] Deleting stale extranet %q (id=%d)", e.Name, e.ID)
			if _, err := extranet.Delete(ctx, service, e.ID); err != nil {
				t.Logf("[SWEEP] Warning: failed to delete stale extranet %q: %v", e.Name, err)
			}
		}
	}
}
