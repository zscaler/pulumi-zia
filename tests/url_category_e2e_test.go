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

// TestUrlCategoryCRUDLifecycle exercises the UrlCategory resource against the live ZIA API.
func TestUrlCategoryCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "url-categories")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	configuredName := shortUniqueName()
	description := shortUniqueName()

	projectKey := func(k string) string { return "url-categories-e2e:" + k }
	pt.SetConfig(t, projectKey("configuredName"), configuredName)
	pt.SetConfig(t, projectKey("description"), description)

	pt.Preview(t)

	createResult := pt.Up(t)

	require.NotNil(t, createResult.Outputs["categoryId"].Value, "categoryId must be set")
	gotName, ok := createResult.Outputs["configuredName"].Value.(string)
	require.True(t, ok && gotName != "", "configuredName must be set")
	require.Equal(t, configuredName, gotName, "configuredName must match config")
	require.Equal(t, description, createResult.Outputs["description"].Value, "description must match config")
	require.Equal(t, true, createResult.Outputs["customCategory"].Value, "customCategory must be true")
	require.Equal(t, "URL_CATEGORY", createResult.Outputs["type"].Value, "type must be URL_CATEGORY")
	require.NotNil(t, createResult.Outputs["val"].Value, "val (category value) must be set")

	keywords := createResult.Outputs["keywords"].Value
	require.NotNil(t, keywords, "keywords must be set")
	assert.ElementsMatch(t, []interface{}{"microsoft"}, toSlice(keywords), "keywords must match config")

	dbUrls := createResult.Outputs["dbCategorizedUrls"].Value
	require.NotNil(t, dbUrls, "dbCategorizedUrls must be set")
	assert.ElementsMatch(t, []interface{}{".creditkarma.com", ".youku.com"}, toSlice(dbUrls), "dbCategorizedUrls must match config")

	_ = pt.Refresh(t)
}
