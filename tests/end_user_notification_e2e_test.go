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

// TestEndUserNotificationCRUDLifecycle exercises End User Notification against the live ZIA API.
// Mirrors Terraform testAccResourceEndUserNotificationConfig: same input values, create, refresh.
// Delete is a no-op for this singleton resource.
func TestEndUserNotificationCRUDLifecycle(t *testing.T) {
	t.Parallel()

	pt := newZIATest(t, "end-user-notification")
	if pt == nil {
		return
	}
	// No defer Destroy: End User Notification is a singleton; Delete is no-op

	pt.Preview(t)

	createResult := pt.Up(t)

	require.Equal(t, "NEVER", createResult.Outputs["aupFrequency"].Value, "aupFrequency must match")
	require.Equal(t, float64(0), createResult.Outputs["aupCustomFrequency"].Value, "aupCustomFrequency must be 0")
	require.Equal(t, float64(0), createResult.Outputs["aupDayOffset"].Value, "aupDayOffset must be 0")
	require.Equal(t, float64(300), createResult.Outputs["cautionAgainAfter"].Value, "cautionAgainAfter must be 300")
	require.Equal(t, true, createResult.Outputs["cautionPerDomain"].Value, "cautionPerDomain must be true")
	require.Equal(t, "Website blocked", createResult.Outputs["customText"].Value, "customText must match")
	require.Equal(t, true, createResult.Outputs["displayCompanyLogo"].Value, "displayCompanyLogo must be true")
	require.Equal(t, true, createResult.Outputs["displayCompanyName"].Value, "displayCompanyName must be true")
	require.Equal(t, true, createResult.Outputs["displayReason"].Value, "displayReason must be true")
	require.Equal(t, "DEFAULT", createResult.Outputs["notificationType"].Value, "notificationType must be DEFAULT")
	require.Equal(t, "http://24326813.zscalerthree.net/policy.html", createResult.Outputs["orgPolicyLink"].Value, "orgPolicyLink must match")
	require.Equal(t, "https://redirect.acme.com", createResult.Outputs["redirectUrl"].Value, "redirectUrl must match")
	require.Equal(t, true, createResult.Outputs["securityReviewEnabled"].Value, "securityReviewEnabled must be true")
	require.Equal(t, true, createResult.Outputs["securityReviewSubmitToSecurityCloud"].Value, "securityReviewSubmitToSecurityCloud must be true")
	require.Equal(t, "Click to request security review.", createResult.Outputs["securityReviewText"].Value, "securityReviewText must match")
	require.Equal(t, "support@24326813.zscalerthree.net", createResult.Outputs["supportEmail"].Value, "supportEmail must match")
	require.Equal(t, "+91-9000000000", createResult.Outputs["supportPhone"].Value, "supportPhone must match")
	require.Equal(t, true, createResult.Outputs["urlCatReviewEnabled"].Value, "urlCatReviewEnabled must be true")
	require.Equal(t, false, createResult.Outputs["urlCatReviewSubmitToSecurityCloud"].Value, "urlCatReviewSubmitToSecurityCloud must be false")
	require.Equal(t, "If you believe you received this message in error, please click here to request a review of this site.", createResult.Outputs["urlCatReviewText"].Value, "urlCatReviewText must match")
	require.Equal(t, true, createResult.Outputs["webDlpReviewEnabled"].Value, "webDlpReviewEnabled must be true")
	require.Equal(t, false, createResult.Outputs["webDlpReviewSubmitToSecurityCloud"].Value, "webDlpReviewSubmitToSecurityCloud must be false")
	require.Equal(t, "Click to request policy review.", createResult.Outputs["webDlpReviewText"].Value, "webDlpReviewText must match")

	assertrefresh.HasNoChanges(t, pt.Refresh(t))
}
