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

//go:build !unit

package tests

import (
	"testing"

	"github.com/pulumi/providertest/pulumitest/assertrefresh"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Terraform variable values from zia/common/testing/variable/variable.go
const (
	vzenStatus         = "ENABLED"
	vzenType           = "VIP"
	vzenIPAddress      = "10.0.0.2"
	vzenSubnetMask     = "255.255.255.0"
	vzenDefaultGateway = "10.0.0.3"
	vzenIpSecEnabled   = true
)

// TestVzenClusterCRUDLifecycle exercises the VzenCluster resource against the live ZIA API.
// Mirrors Terraform TestAccResourceVZENCluster_Basic: same input values, create, update name, refresh.
func TestVzenClusterCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "vzen-cluster")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	initialName := shortUniqueName()

	projectKey := func(k string) string { return "vzen-cluster-e2e:" + k }
	pt.SetConfig(t, projectKey("name"), initialName)
	pt.SetConfig(t, projectKey("status"), vzenStatus)
	pt.SetConfig(t, projectKey("type"), vzenType)
	pt.SetConfig(t, projectKey("ipAddress"), vzenIPAddress)
	pt.SetConfig(t, projectKey("subnetMask"), vzenSubnetMask)
	pt.SetConfig(t, projectKey("defaultGateway"), vzenDefaultGateway)
	pt.SetConfig(t, projectKey("ipSecEnabled"), "true")

	pt.Preview(t)

	createResult := pt.Up(t)

	require.NotNil(t, createResult.Outputs["clusterId"].Value, "clusterId must be set")
	assert.NotEqual(t, float64(0), createResult.Outputs["clusterId"].Value, "clusterId must not be zero")
	require.Equal(t, initialName, createResult.Outputs["name"].Value, "name must match config")
	require.Equal(t, vzenStatus, createResult.Outputs["status"].Value, "status must match config")
	require.Equal(t, vzenType, createResult.Outputs["type"].Value, "type must match config")
	require.Equal(t, vzenIPAddress, createResult.Outputs["ipAddress"].Value, "ipAddress must match config")
	require.Equal(t, vzenSubnetMask, createResult.Outputs["subnetMask"].Value, "subnetMask must match config")
	require.Equal(t, vzenDefaultGateway, createResult.Outputs["defaultGateway"].Value, "defaultGateway must match config")
	require.Equal(t, vzenIpSecEnabled, createResult.Outputs["ipSecEnabled"].Value, "ipSecEnabled must match config")

	assertrefresh.HasNoChanges(t, pt.Refresh(t))

	// Update: change name (same as Terraform tf-updated- + same suffix)
	suffix := initialName[len("tf-acc-test-"):]
	updatedName := "tf-updated-" + suffix
	pt.SetConfig(t, projectKey("name"), updatedName)

	updateResult := pt.Up(t)
	require.Equal(t, updatedName, updateResult.Outputs["name"].Value, "name must be updated")
	require.Equal(t, vzenStatus, updateResult.Outputs["status"].Value, "status must still match")
	require.Equal(t, vzenType, updateResult.Outputs["type"].Value, "type must still match")
	require.Equal(t, vzenIPAddress, updateResult.Outputs["ipAddress"].Value, "ipAddress must still match")
	require.Equal(t, vzenSubnetMask, updateResult.Outputs["subnetMask"].Value, "subnetMask must still match")
	require.Equal(t, vzenDefaultGateway, updateResult.Outputs["defaultGateway"].Value, "defaultGateway must still match")
	require.Equal(t, vzenIpSecEnabled, updateResult.Outputs["ipSecEnabled"].Value, "ipSecEnabled must still match")

	assertrefresh.HasNoChanges(t, pt.Refresh(t))
}
