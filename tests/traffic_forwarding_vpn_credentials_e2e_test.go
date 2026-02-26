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
	"crypto/rand"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

// randomPSK generates a random pre-shared key for VPN credentials (min 8 chars).
func randomPSK() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		panic("randomPSK: " + err.Error())
	}
	return hex.EncodeToString(b)
}

// TestTrafficForwardingVpnCredentialsCRUDLifecycle exercises the TrafficForwardingVpnCredentials resource (UFQDN type).
func TestTrafficForwardingVpnCredentialsCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "traffic-forwarding-vpn-credentials")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	name := shortUniqueName()
	fqdn := name + "@securitygeek.io"
	psk := randomPSK()
	comments := "tf-acc-test-" + name

	projectKey := func(k string) string { return "traffic-forwarding-vpn-credentials-e2e:" + k }
	pt.SetConfig(t, projectKey("fqdn"), fqdn)
	pt.SetConfig(t, projectKey("preSharedKey"), psk)
	pt.SetConfig(t, projectKey("comments"), comments)

	pt.Preview(t)

	createResult := pt.Up(t)

	require.NotNil(t, createResult.Outputs["vpnId"].Value, "vpnId must be set")
	require.Equal(t, "UFQDN", createResult.Outputs["type"].Value, "type must be UFQDN")
	require.Equal(t, fqdn, createResult.Outputs["fqdn"].Value, "fqdn must match config")
	require.Equal(t, comments, createResult.Outputs["comments"].Value, "comments must match config")

	_ = pt.Refresh(t)
}
