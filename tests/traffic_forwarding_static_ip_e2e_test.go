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

	"github.com/stretchr/testify/require"
)

// TestTrafficForwardingStaticIpCRUDLifecycle exercises the TrafficForwardingStaticIp resource against the live ZIA API.
func TestTrafficForwardingStaticIpCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "traffic-forwarding-static-ips")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	ipAddress := randomIPFromCIDR("104.238.235")
	comment := "tf-acc-test-" + shortUniqueName()

	projectKey := func(k string) string { return "traffic-forwarding-static-ips-e2e:" + k }
	pt.SetConfig(t, projectKey("ipAddress"), ipAddress)
	pt.SetConfig(t, projectKey("comment"), comment)
	pt.SetConfig(t, projectKey("routableIp"), "true")
	pt.SetConfig(t, projectKey("geoOverride"), "true")

	pt.Preview(t)

	createResult := pt.Up(t)

	require.NotNil(t, createResult.Outputs["staticIpId"].Value, "staticIpId must be set")
	require.Equal(t, ipAddress, createResult.Outputs["ipAddress"].Value, "ipAddress must match config")
	require.Equal(t, comment, createResult.Outputs["comment"].Value, "comment must match config")
	require.Equal(t, true, createResult.Outputs["routableIp"].Value, "routableIp must be true")
	require.Equal(t, true, createResult.Outputs["geoOverride"].Value, "geoOverride must be true")

	_ = pt.Refresh(t)
}
