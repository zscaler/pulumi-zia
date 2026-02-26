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

// TestSandboxRuleCRUDLifecycle exercises the SandboxRule resource against the live ZIA API,
// mirroring the Terraform acceptance test. Uses RuleLabel as secondary resource.
func TestSandboxRuleCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "sandbox-rules")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	ruleName := shortUniqueName()
	desc := "testAcc_sandbox_rule"
	state := "ENABLED"
	action := "ALLOW"
	labelName := shortUniqueName()

	projectKey := func(k string) string { return "sandbox-rules-e2e:" + k }
	pt.SetConfig(t, projectKey("ruleName"), ruleName)
	pt.SetConfig(t, projectKey("description"), desc)
	pt.SetConfig(t, projectKey("state"), state)
	pt.SetConfig(t, projectKey("action"), action)
	pt.SetConfig(t, projectKey("labelName"), labelName)

	pt.Preview(t)

	createResult := pt.Up(t)

	require.NotNil(t, createResult.Outputs["ruleId"].Value, "ruleId must be set")
	require.Equal(t, ruleName, createResult.Outputs["name"].Value, "name must match config")
	require.Equal(t, desc, createResult.Outputs["description"].Value, "description must match config")
	require.Equal(t, state, createResult.Outputs["state"].Value, "state must be ENABLED")
	require.Equal(t, action, createResult.Outputs["baRuleAction"].Value, "baRuleAction must be ALLOW")

	baPolicyCategories := createResult.Outputs["baPolicyCategories"].Value
	require.NotNil(t, baPolicyCategories, "baPolicyCategories must be set")
	assert.Len(t, toSlice(baPolicyCategories), 4, "baPolicyCategories must have 4 elements")

	fileTypes := createResult.Outputs["fileTypes"].Value
	require.NotNil(t, fileTypes, "fileTypes must be set")
	assert.Len(t, toSlice(fileTypes), 19, "fileTypes must have 19 elements")

	protocols := createResult.Outputs["protocols"].Value
	require.NotNil(t, protocols, "protocols must be set")
	assert.Len(t, toSlice(protocols), 4, "protocols must have 4 elements")

	labels := createResult.Outputs["labels"].Value
	require.NotNil(t, labels, "labels must be set")
	assert.NotEmpty(t, toSlice(labels), "labels must not be empty")

	_ = pt.Refresh(t)
}
