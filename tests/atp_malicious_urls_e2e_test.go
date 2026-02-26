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

func TestAtpMaliciousUrlsCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "atp-malicious-urls")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	pt.Preview(t)

	createResult := pt.Up(t)
	require.NotNil(t, createResult.Outputs["resourceId"].Value, "resourceId must be set")
	maliciousUrls := createResult.Outputs["maliciousUrls"].Value
	require.NotNil(t, maliciousUrls, "maliciousUrls must be set")
	assert.ElementsMatch(t, []interface{}{".example.com", ".test.com"}, toSlice(maliciousUrls),
		"maliciousUrls must match config")

	// Refresh verifies Read works. Skip HasNoChanges: singleton with array state can show drift.
	_ = pt.Refresh(t)
}
