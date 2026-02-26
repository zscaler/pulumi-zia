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

//go:build !unit

package tests

import (
	"testing"

	"github.com/pulumi/providertest/pulumitest/assertrefresh"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNssServerCRUDLifecycle exercises the NssServer resource against the live ZIA API.
func TestNssServerCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "nss-servers")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	name := shortUniqueName()
	status := "ENABLED"
	serverType := "NSS_FOR_FIREWALL"

	projectKey := func(k string) string { return "nss-servers-e2e:" + k }
	pt.SetConfig(t, projectKey("name"), name)
	pt.SetConfig(t, projectKey("status"), status)
	pt.SetConfig(t, projectKey("type"), serverType)

	pt.Preview(t)

	createResult := pt.Up(t)

	require.NotNil(t, createResult.Outputs["nssId"].Value, "nssId must be set")
	assert.NotEqual(t, float64(0), createResult.Outputs["nssId"].Value, "nssId must not be zero")
	require.Equal(t, name, createResult.Outputs["name"].Value, "name must match config")
	require.Equal(t, status, createResult.Outputs["status"].Value, "status must be ENABLED")
	require.Equal(t, serverType, createResult.Outputs["type"].Value, "type must be NSS_FOR_FIREWALL")

	assertrefresh.HasNoChanges(t, pt.Refresh(t))
}
