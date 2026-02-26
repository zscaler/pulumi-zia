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
	"github.com/stretchr/testify/require"
)

// TestWorkloadGroupsCRUDLifecycle exercises the WorkloadGroup resource against the live ZIA API.
// Mirrors Python test_workload_groups.py: create with name/description, update name/description, refresh.
func TestWorkloadGroupsCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "workload-groups")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	initialName := shortUniqueName()
	desc := "Test workload group for VCR testing"

	projectKey := func(k string) string { return "workload-groups-e2e:" + k }
	pt.SetConfig(t, projectKey("name"), initialName)
	pt.SetConfig(t, projectKey("description"), desc)

	pt.Preview(t)

	createResult := pt.Up(t)

	require.NotNil(t, createResult.Outputs["groupId"].Value, "groupId must be set")
	require.NotEqual(t, float64(0), createResult.Outputs["groupId"].Value, "groupId must not be zero")
	require.Equal(t, initialName, createResult.Outputs["name"].Value, "name must match config")
	require.Equal(t, desc, createResult.Outputs["description"].Value, "description must match config")

	assertrefresh.HasNoChanges(t, pt.Refresh(t))

	// Update: change name and description (same as Python TestWorkloadGroup_VCR_Updated)
	updatedName := initialName + "_Updated"
	updatedDesc := "Updated test workload group"
	pt.SetConfig(t, projectKey("name"), updatedName)
	pt.SetConfig(t, projectKey("description"), updatedDesc)

	updateResult := pt.Up(t)
	require.Equal(t, updatedName, updateResult.Outputs["name"].Value, "name must be updated")
	require.Equal(t, updatedDesc, updateResult.Outputs["description"].Value, "description must be updated")

	assertrefresh.HasNoChanges(t, pt.Refresh(t))
}
