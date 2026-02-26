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
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAtpSettingsCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "atp-settings")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	initialRisk := 50
	updatedRisk := 75

	projectKey := func(k string) string { return "atp-settings-e2e:" + k }
	pt.SetConfig(t, projectKey("riskTolerance"), strconv.Itoa(initialRisk))

	pt.Preview(t)

	createResult := pt.Up(t)
	require.NotNil(t, createResult.Outputs["resourceId"].Value, "resourceId must be set")
	require.Equal(t, float64(initialRisk), createResult.Outputs["riskTolerance"].Value,
		"riskTolerance must match initial config")
	require.Equal(t, false, createResult.Outputs["cmdCtlServerBlocked"].Value, "cmdCtlServerBlocked must be false")
	require.Equal(t, false, createResult.Outputs["malwareSitesBlocked"].Value, "malwareSitesBlocked must be false")

	// Refresh verifies Read works. Skip HasNoChanges: ATP settings has many Capture fields
	// returned by the API that can cause drift when only a subset is configured.
	_ = pt.Refresh(t)

	pt.SetConfig(t, projectKey("riskTolerance"), strconv.Itoa(updatedRisk))
	updateResult := pt.Up(t)
	require.Equal(t, float64(updatedRisk), updateResult.Outputs["riskTolerance"].Value,
		"riskTolerance must match updated config")

	_ = pt.Refresh(t)
}
