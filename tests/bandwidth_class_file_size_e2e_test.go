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

	"github.com/stretchr/testify/require"
)

func TestBandwidthClassFileSizeCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "bandwidth-classes-file-size")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	// Terraform test: Step 1 - Create with FILE_250MB
	projectKey := func(k string) string { return "bandwidth-classes-file-size-e2e:" + k }
	pt.SetConfig(t, projectKey("fileSize"), "FILE_250MB")
	pt.Preview(t)
	createResult := pt.Up(t)
	require.NotNil(t, createResult.Outputs["classId"].Value, "classId must be set")
	require.Equal(t, "BANDWIDTH_CAT_LARGE_FILE", createResult.Outputs["name"].Value, "name must be BANDWIDTH_CAT_LARGE_FILE")
	require.Equal(t, "BANDWIDTH_CAT_LARGE_FILE", createResult.Outputs["type"].Value, "type must be BANDWIDTH_CAT_LARGE_FILE")
	require.Equal(t, "FILE_250MB", createResult.Outputs["fileSize"].Value, "fileSize must match initial config")

	// Refresh verifies Read works. Skip HasNoChanges: updates predefined class, can show drift.
	_ = pt.Refresh(t)

	// Terraform test: Step 2 - Update to FILE_1GB
	pt.SetConfig(t, projectKey("fileSize"), "FILE_1GB")
	updateResult := pt.Up(t)
	require.Equal(t, "FILE_1GB", updateResult.Outputs["fileSize"].Value, "fileSize must match updated config")

	_ = pt.Refresh(t)
}
