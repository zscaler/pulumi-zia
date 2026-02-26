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

	"github.com/pulumi/providertest/pulumitest/assertrefresh"
	"github.com/stretchr/testify/require"
)

func TestAdvancedSettingsCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "advanced-settings")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	initialTimeout := 300
	updatedTimeout := 600

	projectKey := func(k string) string { return "advanced-settings-e2e:" + k }
	pt.SetConfig(t, projectKey("uiSessionTimeout"), strconv.Itoa(initialTimeout))

	pt.Preview(t)

	createResult := pt.Up(t)
	require.NotNil(t, createResult.Outputs["resourceId"].Value, "resourceId must be set")
	require.Equal(t, float64(initialTimeout), createResult.Outputs["uiSessionTimeout"].Value,
		"uiSessionTimeout must match initial config")

	assertrefresh.HasNoChanges(t, pt.Refresh(t))

	pt.SetConfig(t, projectKey("uiSessionTimeout"), strconv.Itoa(updatedTimeout))
	updateResult := pt.Up(t)
	require.Equal(t, float64(updatedTimeout), updateResult.Outputs["uiSessionTimeout"].Value,
		"uiSessionTimeout must match updated config")

	assertrefresh.HasNoChanges(t, pt.Refresh(t))
}
