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
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestURLFilteringRuleCRUDLifecycle exercises the URLFilteringRule
// resource against the live ZIA API, mirroring the Terraform acceptance test.
// Uses RuleLabel as secondary resource. Same attribute values as Terraform test.
func TestURLFilteringRuleCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "url-filtering-rules")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	ruleName := shortUniqueName()
	desc := "testAcc_url_filtering_rule"
	action := "ALLOW"
	state := "ENABLED"
	labelName := shortUniqueName()

	// Validity times: same as Terraform - start 5 min from now, end 365 days + 5 min
	validityStartTime := time.Now().Add(5 * time.Minute).UTC().Format(time.RFC1123)
	validityEndTime := time.Now().AddDate(1, 0, 0).Add(5 * time.Minute).UTC().Format(time.RFC1123)

	projectKey := func(k string) string { return "url-filtering-rules-e2e:" + k }
	pt.SetConfig(t, projectKey("ruleName"), ruleName)
	pt.SetConfig(t, projectKey("description"), desc)
	pt.SetConfig(t, projectKey("action"), action)
	pt.SetConfig(t, projectKey("state"), state)
	pt.SetConfig(t, projectKey("labelName"), labelName)
	pt.SetConfig(t, projectKey("validityStartTime"), validityStartTime)
	pt.SetConfig(t, projectKey("validityEndTime"), validityEndTime)

	pt.Preview(t)

	createResult := pt.Up(t)

	require.NotNil(t, createResult.Outputs["ruleId"].Value, "ruleId must be set")
	require.Equal(t, ruleName, createResult.Outputs["name"].Value, "name must match config")
	require.Equal(t, desc, createResult.Outputs["description"].Value, "description must match config")
	require.Equal(t, action, createResult.Outputs["action"].Value, "action must be ALLOW")
	require.Equal(t, state, createResult.Outputs["state"].Value, "state must be ENABLED")

	urlCategories := createResult.Outputs["urlCategories"].Value
	require.NotNil(t, urlCategories, "urlCategories must be set")
	assert.ElementsMatch(t, []interface{}{"ANY"}, toSlice(urlCategories), "urlCategories must be [ANY]")

	protocols := createResult.Outputs["protocols"].Value
	require.NotNil(t, protocols, "protocols must be set")
	assert.ElementsMatch(t, []interface{}{"ANY_RULE"}, toSlice(protocols), "protocols must be [ANY_RULE]")

	requestMethods := createResult.Outputs["requestMethods"].Value
	require.NotNil(t, requestMethods, "requestMethods must be set")
	assert.Len(t, toSlice(requestMethods), 9, "requestMethods must have 9 elements")

	labels := createResult.Outputs["labels"].Value
	require.NotNil(t, labels, "labels must be set")
	assert.NotEmpty(t, toSlice(labels), "labels must not be empty")

	// Refresh verifies Read works. Skip HasNoChanges: rule ordering can show drift.
	_ = pt.Refresh(t)

	// Update: change description (same as Terraform - FWRuleResourceDescription)
	updatedDesc := "this is an acceptance test"
	pt.SetConfig(t, projectKey("description"), updatedDesc)
	updateResult := pt.Up(t)
	require.Equal(t, updatedDesc, updateResult.Outputs["description"].Value,
		"description must match updated config")
	assert.Equal(t, createResult.Outputs["ruleId"].Value, updateResult.Outputs["ruleId"].Value,
		"ruleId must remain the same after update")
	require.Equal(t, action, updateResult.Outputs["action"].Value, "action must remain ALLOW")
	require.Equal(t, state, updateResult.Outputs["state"].Value, "state must remain ENABLED")

	_ = pt.Refresh(t)
}
