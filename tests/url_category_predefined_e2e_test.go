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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestUrlCategoryPredefinedCRUDLifecycle exercises the UrlCategoryPredefined resource against the live ZIA API.
// Manages custom URLs/keywords/IPs added to the predefined FINANCE category. Delete is a no-op.
func TestUrlCategoryPredefinedCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "url-categories-predefined")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	pt.Preview(t)

	createResult := pt.Up(t)

	require.NotNil(t, createResult.Outputs["categoryId"].Value, "categoryId must be set")
	// configuredName is read-only from API; may be "FINANCE" or "Finance" depending on ZIA response
	gotName, ok := createResult.Outputs["configuredName"].Value.(string)
	require.True(t, ok && gotName != "", "configuredName must be set")
	assert.True(t, strings.EqualFold("FINANCE", gotName), "configuredName must match FINANCE (case-insensitive)")
	require.Equal(t, "URL_CATEGORY", createResult.Outputs["type"].Value, "type must be URL_CATEGORY")

	// Verify urls, keywords, ipRanges are set (from config defaults)
	if ov := createResult.Outputs["urls"]; ov.Value != nil {
		assert.Len(t, toSlice(ov.Value), 2, "urls should have 2 items")
	}
	if ov := createResult.Outputs["keywords"]; ov.Value != nil {
		assert.Len(t, toSlice(ov.Value), 2, "keywords should have 2 items")
	}
	if ov := createResult.Outputs["ipRanges"]; ov.Value != nil {
		assert.Len(t, toSlice(ov.Value), 1, "ipRanges should have 1 item")
	}

	_ = pt.Refresh(t)
}
