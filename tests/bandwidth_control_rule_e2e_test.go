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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestBandwidthControlRuleCRUDLifecycle exercises the BandwidthControlRule
// resource against the live ZIA API, mirroring the Terraform acceptance test.
// Uses BandwidthClassWebConferencing (BANDWIDTH_CAT_WEBCONF) and RuleLabel
// as secondary resources, without time_windows.
func TestBandwidthControlRuleCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "bandwidth-control-rules")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	ruleName := shortUniqueName()
	desc := "tf-acc-test-cloud-app-control"
	labelName := shortUniqueName()

	projectKey := func(k string) string { return "bandwidth-control-rules-e2e:" + k }
	pt.SetConfig(t, projectKey("ruleName"), ruleName)
	pt.SetConfig(t, projectKey("description"), desc)
	pt.SetConfig(t, projectKey("state"), "ENABLED")
	pt.SetConfig(t, projectKey("labelName"), labelName)

	pt.Preview(t)

	createResult := pt.Up(t)

	require.NotNil(t, createResult.Outputs["ruleId"].Value, "ruleId must be set")
	require.Equal(t, ruleName, createResult.Outputs["name"].Value, "name must match config")
	require.Equal(t, desc, createResult.Outputs["description"].Value, "description must match config")
	require.Equal(t, "ENABLED", createResult.Outputs["state"].Value, "state must be ENABLED")
	require.Equal(t, float64(1), createResult.Outputs["order"].Value, "order must be 1")

	bandwidthClasses := createResult.Outputs["bandwidthClasses"].Value
	require.NotNil(t, bandwidthClasses, "bandwidthClasses must be set")
	assert.NotEmpty(t, toSlice(bandwidthClasses), "bandwidthClasses must not be empty")

	labels := createResult.Outputs["labels"].Value
	require.NotNil(t, labels, "labels must be set")
	assert.NotEmpty(t, toSlice(labels), "labels must not be empty")

	// Refresh verifies Read works. Skip HasNoChanges: rule ordering can show drift.
	_ = pt.Refresh(t)

	// Update description
	updatedDesc := "tf-acc-test-cloud-app-control-updated"
	pt.SetConfig(t, projectKey("description"), updatedDesc)
	updateResult := pt.Up(t)
	require.Equal(t, updatedDesc, updateResult.Outputs["description"].Value, "description must match updated config")
	assert.Equal(t, createResult.Outputs["ruleId"].Value, updateResult.Outputs["ruleId"].Value,
		"ruleId must remain the same after update")

	_ = pt.Refresh(t)
}
