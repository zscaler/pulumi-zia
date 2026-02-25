// Copyright (c) 2023 Zscaler Technology Alliances, <zscaler-partner-labs@z-bd.com>
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

package tests

import (
	"context"
	"testing"

	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zscaler/pulumi-zia/provider"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/property"
)

func TestRuleLabelCreateDryRun(t *testing.T) {
	t.Parallel()

	prov := providerServer(t)

	// Dry-run create: should succeed without real ZIA credentials
	response, err := prov.Create(p.CreateRequest{
		Urn: urn("RuleLabel"),
		Properties: property.NewMap(map[string]property.Value{
			"name": property.New("test-label"),
		}),
		DryRun: true,
	})

	require.NoError(t, err)
	assert.Equal(t, "preview", response.ID)
	_, ok := response.Properties.GetOk("ruleLabelId")
	assert.True(t, ok)
}

// urn builds a URN for integration tests.
func urn(typ string) resource.URN {
	return resource.NewURN("stack", "proj", "",
		tokens.Type("zia:index:"+typ), "name")
}

func providerServer(t *testing.T) integration.Server {
	s, err := integration.NewServer(
		context.Background(),
		provider.Name,
		semver.MustParse("1.0.0"),
		integration.WithProvider(provider.Provider()),
	)
	require.NoError(t, err)
	return s
}
