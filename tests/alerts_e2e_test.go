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

func TestSubscriptionAlertCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "alerts")
	if pt == nil {
		return
	}
	defer func() { pt.Destroy(t) }()

	initialEmail := shortUniqueName() + "@acme.com"
	updatedEmail := shortUniqueName() + "@acme.com"
	desc := "testAcc_Alert_Subscription"

	projectKey := func(k string) string { return "alerts-e2e:" + k }
	pt.SetConfig(t, projectKey("email"), initialEmail)
	pt.SetConfig(t, projectKey("description"), desc)

	pt.Preview(t)

	createResult := pt.Up(t)
	require.Equal(t, initialEmail, createResult.Outputs["email"].Value, "created alert email must match input")
	require.Equal(t, desc, createResult.Outputs["description"].Value, "created alert description must match input")
	alertID := createResult.Outputs["alertId"].Value
	require.NotNil(t, alertID, "alert id must be set after create")
	assert.NotEqual(t, float64(0), alertID, "alert id must not be zero")

	assertrefresh.HasNoChanges(t, pt.Refresh(t))

	pt.SetConfig(t, projectKey("email"), updatedEmail)
	updateResult := pt.Up(t)
	require.Equal(t, updatedEmail, updateResult.Outputs["email"].Value, "updated alert email must match new input")
	assert.Equal(t, alertID, updateResult.Outputs["alertId"].Value, "alert id must remain the same after update")

	assertrefresh.HasNoChanges(t, pt.Refresh(t))
}
