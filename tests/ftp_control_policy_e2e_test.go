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

	"github.com/pulumi/providertest/pulumitest/assertrefresh"
	"github.com/stretchr/testify/require"
)

// TestFtpControlPolicyCRUDLifecycle exercises the FtpControlPolicy resource (singleton) against the live ZIA API.
// Mirrors Terraform test: create with disabled, update with enabled + url_categories + urls.
func TestFtpControlPolicyCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "ftp-control-policy")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	projectKey := func(k string) string { return "ftp-control-policy-e2e:" + k }

	// Step 1: Create with both disabled (no categories/urls)
	pt.SetConfig(t, projectKey("ftpEnabled"), "false")
	pt.SetConfig(t, projectKey("ftpOverHttpEnabled"), "false")

	pt.Preview(t)
	createResult := pt.Up(t)

	require.Equal(t, false, createResult.Outputs["ftpEnabled"].Value, "ftpEnabled must be false")
	require.Equal(t, false, createResult.Outputs["ftpOverHttpEnabled"].Value, "ftpOverHttpEnabled must be false")

	_ = pt.Refresh(t)

	// Step 2: Update with enabled
	pt.SetConfig(t, projectKey("ftpEnabled"), "true")
	pt.SetConfig(t, projectKey("ftpOverHttpEnabled"), "true")

	updateResult := pt.Up(t)

	require.Equal(t, true, updateResult.Outputs["ftpEnabled"].Value, "ftpEnabled must be true")
	require.Equal(t, true, updateResult.Outputs["ftpOverHttpEnabled"].Value, "ftpOverHttpEnabled must be true")

	// Delete is a no-op for this singleton
	assertrefresh.HasNoChanges(t, pt.Refresh(t))
}
