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

// TestRiskProfileCRUDLifecycle exercises the RiskProfile resource against the live ZIA API.
func TestRiskProfileCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "risk-profiles")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	initialName := shortUniqueName()
	updatedName := shortUniqueName()

	projectKey := func(k string) string { return "risk-profiles-e2e:" + k }
	pt.SetConfig(t, projectKey("profileName"), initialName)

	pt.Preview(t)

	createResult := pt.Up(t)

	require.NotNil(t, createResult.Outputs["profileId"].Value, "profileId must be set")
	assert.NotEqual(t, float64(0), createResult.Outputs["profileId"].Value, "profileId must not be zero")
	require.Equal(t, initialName, createResult.Outputs["profileName"].Value, "profileName must match config")
	require.Equal(t, "CLOUD_APPLICATIONS", createResult.Outputs["profileType"].Value, "profileType must be CLOUD_APPLICATIONS")
	require.Equal(t, "ANY", createResult.Outputs["status"].Value, "status must be ANY")

	assertrefresh.HasNoChanges(t, pt.Refresh(t))

	pt.SetConfig(t, projectKey("profileName"), updatedName)
	updateResult := pt.Up(t)
	require.Equal(t, updatedName, updateResult.Outputs["profileName"].Value, "profileName must match updated config")
	assert.Equal(t, createResult.Outputs["profileId"].Value, updateResult.Outputs["profileId"].Value, "profileId must remain the same after update")

	assertrefresh.HasNoChanges(t, pt.Refresh(t))
}
