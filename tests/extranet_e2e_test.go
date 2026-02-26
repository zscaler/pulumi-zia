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

// TestExtranetCRUDLifecycle exercises Extranet against the live ZIA API.
// Mirrors Terraform test: same input values (DNS list, IP pool list), create, update name, refresh.
func TestExtranetCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "extranet")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	initialName := shortUniqueName()
	desc := "testAcc_extranet"

	// Clean up any stale tf-acc-test-* extranets from previous failed runs.
	svc := newTestZIAService(t)
	sweepStaleExtranets(t, svc)

	projectKey := func(k string) string { return "extranet-e2e:" + k }
	pt.SetConfig(t, projectKey("name"), initialName)
	pt.SetConfig(t, projectKey("description"), desc)

	pt.Preview(t)

	createResult := pt.Up(t)

	require.NotNil(t, createResult.Outputs["extranetId"].Value, "extranetId must be set")
	assert.NotEqual(t, float64(0), createResult.Outputs["extranetId"].Value, "extranetId must not be zero")
	require.Equal(t, initialName, createResult.Outputs["name"].Value, "name must match config")
	require.Equal(t, desc, createResult.Outputs["description"].Value, "description must match config")

	dnsList := toSlice(createResult.Outputs["extranetDnsList"].Value)
	require.Len(t, dnsList, 2, "extranetDnsList must have 2 entries")

	ipPoolList := toSlice(createResult.Outputs["extranetIpPoolList"].Value)
	require.Len(t, ipPoolList, 2, "extranetIpPoolList must have 2 entries")

	assertrefresh.HasNoChanges(t, pt.Refresh(t))

	// Update: change name (same as Terraform tf-updated- + same suffix)
	suffix := initialName[len("tf-acc-test-"):]
	updatedName := "tf-updated-" + suffix
	pt.SetConfig(t, projectKey("name"), updatedName)

	updateResult := pt.Up(t)
	require.Equal(t, updatedName, updateResult.Outputs["name"].Value, "name must be updated")

	assertrefresh.HasNoChanges(t, pt.Refresh(t))
}
