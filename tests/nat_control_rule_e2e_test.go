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

// TestNatControlRuleCRUDLifecycle exercises the NatControlRule resource against the live ZIA API.
// Uses RuleLabel and GetFwNetworkService (ZSCALER_PROXY_NW_SERVICES) as dependencies.
func TestNatControlRuleCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "nat-control-rules")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	ruleName := shortUniqueName()
	labelName := shortUniqueName()
	desc := "this is an acceptance test"
	state := "ENABLED"

	projectKey := func(k string) string { return "nat-control-rules-e2e:" + k }
	pt.SetConfig(t, projectKey("ruleName"), ruleName)
	pt.SetConfig(t, projectKey("labelName"), labelName)
	pt.SetConfig(t, projectKey("description"), desc)
	pt.SetConfig(t, projectKey("state"), state)

	pt.Preview(t)

	createResult := pt.Up(t)

	require.NotNil(t, createResult.Outputs["ruleId"].Value, "ruleId must be set")
	gotName, ok := createResult.Outputs["name"].Value.(string)
	require.True(t, ok && gotName != "", "name must be set")
	require.Equal(t, ruleName, gotName, "name must match config")
	require.Equal(t, desc, createResult.Outputs["description"].Value, "description must match config")
	require.Equal(t, state, createResult.Outputs["state"].Value, "state must be ENABLED")

	order := createResult.Outputs["order"].Value
	require.NotNil(t, order, "order must be set")
	assert.Equal(t, float64(1), order, "order must be 1")

	redirectPort := createResult.Outputs["redirectPort"].Value
	require.NotNil(t, redirectPort, "redirectPort must be set")
	assert.Equal(t, float64(5000), redirectPort, "redirectPort must be 5000")

	redirectIp := createResult.Outputs["redirectIp"].Value
	require.NotNil(t, redirectIp, "redirectIp must be set")
	assert.Equal(t, "192.168.100.150", redirectIp, "redirectIp must be 192.168.100.150")

	labels := createResult.Outputs["labels"].Value
	require.NotNil(t, labels, "labels must be set")
	assert.NotEmpty(t, toSlice(labels), "labels must not be empty")

	nwServices := createResult.Outputs["nwServices"].Value
	require.NotNil(t, nwServices, "nwServices must be set")
	assert.NotEmpty(t, toSlice(nwServices), "nwServices must not be empty")

	_ = pt.Refresh(t)
}
