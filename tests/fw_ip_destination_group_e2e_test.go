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

// TestFwIpDestinationGroupCRUDLifecycle exercises the FwIpDestinationGroup resource against the live ZIA API.
func TestFwIpDestinationGroupCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "fw-ip-destination-groups")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	name := shortUniqueName()
	desc := "this is an acceptance test"
	groupType := "DSTN_FQDN"

	projectKey := func(k string) string { return "fw-ip-destination-groups-e2e:" + k }
	pt.SetConfig(t, projectKey("name"), name)
	pt.SetConfig(t, projectKey("description"), desc)
	pt.SetConfig(t, projectKey("type"), groupType)

	pt.Preview(t)

	createResult := pt.Up(t)

	require.NotNil(t, createResult.Outputs["groupId"].Value, "groupId must be set")
	gotName, ok := createResult.Outputs["name"].Value.(string)
	require.True(t, ok && gotName != "", "name must be set")
	require.Equal(t, name, gotName, "name must match config")
	require.Equal(t, desc, createResult.Outputs["description"].Value, "description must match config")
	require.Equal(t, groupType, createResult.Outputs["type"].Value, "type must be DSTN_FQDN")

	addresses := createResult.Outputs["addresses"].Value
	require.NotNil(t, addresses, "addresses must be set")
	assert.Len(t, toSlice(addresses), 3, "addresses must have 3 items")

	_ = pt.Refresh(t)
}
