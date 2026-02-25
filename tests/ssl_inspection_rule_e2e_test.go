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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSslInspectionRuleCRUDLifecycle exercises the SslInspectionRule resource against the live ZIA API,
// mirroring the Terraform acceptance test. Uses RuleLabel as secondary resource.
func TestSslInspectionRuleCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "ssl-inspection-rules")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	ruleName := shortUniqueName()
	desc := "testAcc_ssl_rule"
	state := "ENABLED"
	labelName := shortUniqueName()

	projectKey := func(k string) string { return "ssl-inspection-rules-e2e:" + k }
	pt.SetConfig(t, projectKey("ruleName"), ruleName)
	pt.SetConfig(t, projectKey("description"), desc)
	pt.SetConfig(t, projectKey("state"), state)
	pt.SetConfig(t, projectKey("labelName"), labelName)

	pt.Preview(t)

	createResult := pt.Up(t)

	require.NotNil(t, createResult.Outputs["ruleId"].Value, "ruleId must be set")
	require.Equal(t, ruleName, createResult.Outputs["name"].Value, "name must match config")
	require.Equal(t, desc, createResult.Outputs["description"].Value, "description must match config")
	require.Equal(t, state, createResult.Outputs["state"].Value, "state must be ENABLED")
	require.Equal(t, true, createResult.Outputs["roadWarriorForKerberos"].Value, "roadWarriorForKerberos must be true")

	cloudApplications := createResult.Outputs["cloudApplications"].Value
	require.NotNil(t, cloudApplications, "cloudApplications must be set")
	assert.ElementsMatch(t, []interface{}{"CHATGPT_AI", "ANDI"}, toSlice(cloudApplications),
		"cloudApplications must match config")

	platforms := createResult.Outputs["platforms"].Value
	require.NotNil(t, platforms, "platforms must be set")
	assert.Len(t, toSlice(platforms), 6, "platforms must have 6 elements")

	labels := createResult.Outputs["labels"].Value
	require.NotNil(t, labels, "labels must be set")
	assert.NotEmpty(t, toSlice(labels), "labels must not be empty")

	_ = pt.Refresh(t)
}
