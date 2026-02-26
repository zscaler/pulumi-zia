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

// TestRuleLabelCRUDLifecycle exercises every CRUD function of the RuleLabel
// resource against the live ZIA API, mirroring the Terraform acceptance test:
//
//  1. Preview  — dry-run validates inputs without creating anything
//  2. Create   — creates the rule label; asserts exact name, description, and id
//  3. Read     — refresh re-reads from the API; asserts zero drift
//  4. Update   — changes both name and description; asserts new values, same id
//  5. Read     — refresh again; asserts zero drift after update
//  6. Destroy  — deletes the resource; verifies no error
func TestRuleLabelCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "rule-labels")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	initialName := shortUniqueName()
	initialDesc := "Initial description for CRUD lifecycle test"

	projectKey := func(k string) string { return "rule-labels-e2e:" + k }
	pt.SetConfig(t, projectKey("labelName"), initialName)
	pt.SetConfig(t, projectKey("description"), initialDesc)

	// ---- Preview (dry-run) ----
	pt.Preview(t)

	// ---- Create ----
	createResult := pt.Up(t)

	require.Equal(t, initialName, createResult.Outputs["labelName"].Value,
		"created label name must match input")
	require.Equal(t, initialDesc, createResult.Outputs["labelDescription"].Value,
		"created label description must match input")
	labelID := createResult.Outputs["labelId"].Value
	require.NotNil(t, labelID, "label id must be set after create")
	assert.NotEqual(t, float64(0), labelID, "label id must not be zero")

	// ---- Read (Refresh) ----
	assertrefresh.HasNoChanges(t, pt.Refresh(t))

	// ---- Update ----
	updatedName := shortUniqueName()
	updatedDesc := "Updated description for CRUD lifecycle test"

	pt.SetConfig(t, projectKey("labelName"), updatedName)
	pt.SetConfig(t, projectKey("description"), updatedDesc)

	updateResult := pt.Up(t)

	require.Equal(t, updatedName, updateResult.Outputs["labelName"].Value,
		"updated label name must match new input")
	require.Equal(t, updatedDesc, updateResult.Outputs["labelDescription"].Value,
		"updated label description must match new input")
	assert.Equal(t, labelID, updateResult.Outputs["labelId"].Value,
		"label id must remain the same after update")

	// ---- Read after Update ----
	assertrefresh.HasNoChanges(t, pt.Refresh(t))
}
