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

// Terraform variable values from zia/common/testing/variable/variable.go
const (
	vzenNodeStatus = "ENABLED"
	// vzenNodeType               = "VZEN"
	vzenNodeIPAddress          = "10.0.0.10"
	vzenNodeSubnetMask         = "255.255.255.0"
	vzenNodeDefaultGateway     = "10.0.0.20"
	vzenNodeInProduction       = true
	vzenNodeDeploymentMode     = "STANDALONE"
	vzenSkuType                = "LARGE"
	vzenOnDemandSupportTunnel  = true
	vzenEstablishSupportTunnel = false
	vzenNodeIpSecEnabled       = false
)

// TestVzenNodeCRUDLifecycle exercises the VzenNode resource against the live ZIA API.
// Mirrors Terraform TestAccResourceVZENNodeBasic and HCL local_dev/vzen_node/main.tf.
func TestVzenNodeCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "vzen-node")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	initialName := shortUniqueName()

	projectKey := func(k string) string { return "vzen-node-e2e:" + k }
	pt.SetConfig(t, projectKey("name"), initialName)
	pt.SetConfig(t, projectKey("status"), vzenNodeStatus)
	// pt.SetConfig(t, projectKey("type"), vzenNodeType)
	pt.SetConfig(t, projectKey("ipAddress"), vzenNodeIPAddress)
	pt.SetConfig(t, projectKey("subnetMask"), vzenNodeSubnetMask)
	pt.SetConfig(t, projectKey("defaultGateway"), vzenNodeDefaultGateway)
	pt.SetConfig(t, projectKey("inProduction"), "true")
	pt.SetConfig(t, projectKey("deploymentMode"), vzenNodeDeploymentMode)
	pt.SetConfig(t, projectKey("vzenSkuType"), vzenSkuType)
	pt.SetConfig(t, projectKey("onDemandSupportTunnelEnabled"), "true")
	pt.SetConfig(t, projectKey("establishSupportTunnelEnabled"), "false")
	pt.SetConfig(t, projectKey("ipSecEnabled"), "false")

	pt.Preview(t)

	createResult := pt.Up(t)

	require.NotNil(t, createResult.Outputs["nodeId"].Value, "nodeId must be set")
	assert.NotEqual(t, float64(0), createResult.Outputs["nodeId"].Value, "nodeId must not be zero")
	require.Equal(t, initialName, createResult.Outputs["name"].Value, "name must match config")
	require.Equal(t, vzenNodeStatus, createResult.Outputs["status"].Value, "status must match config")
	//require.Equal(t, vzenNodeType, createResult.Outputs["type"].Value, "type must match config")
	require.Equal(t, vzenNodeIPAddress, createResult.Outputs["ipAddress"].Value, "ipAddress must match config")
	require.Equal(t, vzenNodeIpSecEnabled, createResult.Outputs["ipSecEnabled"].Value, "ipSecEnabled must match config")
	require.Equal(t, vzenNodeSubnetMask, createResult.Outputs["subnetMask"].Value, "subnetMask must match config")
	require.Equal(t, vzenNodeDefaultGateway, createResult.Outputs["defaultGateway"].Value, "defaultGateway must match config")
	require.Equal(t, vzenNodeInProduction, createResult.Outputs["inProduction"].Value, "inProduction must match config")
	require.Equal(t, vzenNodeDeploymentMode, createResult.Outputs["deploymentMode"].Value, "deploymentMode must match config")
	require.Equal(t, vzenSkuType, createResult.Outputs["vzenSkuType"].Value, "vzenSkuType must match config")
	require.Equal(t, vzenOnDemandSupportTunnel, createResult.Outputs["onDemandSupportTunnelEnabled"].Value, "onDemandSupportTunnelEnabled must match config")
	require.Equal(t, vzenEstablishSupportTunnel, createResult.Outputs["establishSupportTunnelEnabled"].Value, "establishSupportTunnelEnabled must match config")

	assertrefresh.HasNoChanges(t, pt.Refresh(t))

	// Update: change name (same as Terraform tf-updated- + same suffix)
	suffix := initialName[len("tf-acc-test-"):]
	updatedName := "tf-updated-" + suffix
	pt.SetConfig(t, projectKey("name"), updatedName)

	updateResult := pt.Up(t)
	require.Equal(t, updatedName, updateResult.Outputs["name"].Value, "name must be updated")
	require.Equal(t, vzenNodeStatus, updateResult.Outputs["status"].Value, "status must still match")
	// require.Equal(t, vzenNodeType, updateResult.Outputs["type"].Value, "type must still match")
	require.Equal(t, vzenNodeIPAddress, updateResult.Outputs["ipAddress"].Value, "ipAddress must still match")
	require.Equal(t, vzenNodeIpSecEnabled, updateResult.Outputs["ipSecEnabled"].Value, "ipSecEnabled must still match")
	require.Equal(t, vzenNodeSubnetMask, updateResult.Outputs["subnetMask"].Value, "subnetMask must still match")
	require.Equal(t, vzenNodeDefaultGateway, updateResult.Outputs["defaultGateway"].Value, "defaultGateway must still match")
	require.Equal(t, vzenNodeInProduction, updateResult.Outputs["inProduction"].Value, "inProduction must still match")
	require.Equal(t, vzenNodeDeploymentMode, updateResult.Outputs["deploymentMode"].Value, "deploymentMode must still match")
	require.Equal(t, vzenSkuType, updateResult.Outputs["vzenSkuType"].Value, "vzenSkuType must still match")
	require.Equal(t, vzenOnDemandSupportTunnel, updateResult.Outputs["onDemandSupportTunnelEnabled"].Value, "onDemandSupportTunnelEnabled must still match")
	require.Equal(t, vzenEstablishSupportTunnel, updateResult.Outputs["establishSupportTunnelEnabled"].Value, "establishSupportTunnelEnabled must still match")

	assertrefresh.HasNoChanges(t, pt.Refresh(t))
}
