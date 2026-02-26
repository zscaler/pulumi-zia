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

// TestLocationManagementCRUDLifecycle exercises LocationManagement against the live ZIA API.
// Mirrors Terraform test: creates VPN credentials (UFQDN), then a location referencing them.
func TestLocationManagementCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "location-management")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	name := shortUniqueName()
	desc := "tf-acc-test-" + name
	fqdn := name + "@securitygeek.io"
	psk := randomPSK()

	projectKey := func(k string) string { return "location-management-e2e:" + k }
	pt.SetConfig(t, projectKey("name"), name)
	pt.SetConfig(t, projectKey("description"), desc)
	pt.SetConfig(t, projectKey("fqdn"), fqdn)
	pt.SetConfig(t, projectKey("preSharedKey"), psk)

	pt.Preview(t)

	createResult := pt.Up(t)

	require.NotNil(t, createResult.Outputs["locationId"].Value, "locationId must be set")
	assert.NotEqual(t, float64(0), createResult.Outputs["locationId"].Value, "locationId must not be zero")
	require.Equal(t, name, createResult.Outputs["name"].Value, "name must match config")
	require.Equal(t, desc, createResult.Outputs["description"].Value, "description must match config")
	require.Equal(t, "UNITED_STATES", createResult.Outputs["country"].Value, "country must be UNITED_STATES")
	require.Equal(t, true, createResult.Outputs["authRequired"].Value, "authRequired must be true")
	require.Equal(t, true, createResult.Outputs["surrogateIp"].Value, "surrogateIp must be true")

	assertrefresh.HasNoChanges(t, pt.Refresh(t))

	// Update: change description
	updatedDesc := "updated-" + name
	pt.SetConfig(t, projectKey("description"), updatedDesc)

	updateResult := pt.Up(t)
	require.Equal(t, updatedDesc, updateResult.Outputs["description"].Value, "description must be updated")

	assertrefresh.HasNoChanges(t, pt.Refresh(t))
}
