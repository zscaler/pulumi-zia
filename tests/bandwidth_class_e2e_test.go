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

func TestBandwidthClassCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "bandwidth-classes")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	initialName := shortUniqueName()
	updatedName := shortUniqueName()

	projectKey := func(k string) string { return "bandwidth-classes-e2e:" + k }
	pt.SetConfig(t, projectKey("name"), initialName)

	pt.Preview(t)

	createResult := pt.Up(t)
	require.Equal(t, initialName, createResult.Outputs["name"].Value, "name must match initial config")
	classID := createResult.Outputs["classId"].Value
	require.NotNil(t, classID, "classId must be set after create")
	assert.NotEqual(t, float64(0), classID, "classId must not be zero")

	urls := createResult.Outputs["urls"].Value
	require.NotNil(t, urls, "urls must be set")
	assert.ElementsMatch(t, []interface{}{"chatgpt.com", "chatgpt1.com", "openai.com"}, toSlice(urls),
		"urls must match config")

	assertrefresh.HasNoChanges(t, pt.Refresh(t))

	pt.SetConfig(t, projectKey("name"), updatedName)
	updateResult := pt.Up(t)
	require.Equal(t, updatedName, updateResult.Outputs["name"].Value, "name must match updated config")
	assert.Equal(t, classID, updateResult.Outputs["classId"].Value, "classId must remain the same after update")

	assertrefresh.HasNoChanges(t, pt.Refresh(t))
}
