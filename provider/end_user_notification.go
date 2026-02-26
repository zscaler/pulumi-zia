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

// Package provider implements the End User Notification resource (singleton).
// Adopted from terraform-provider-zia resource_zia_end_user_notification.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/end_user_notification"
)

// EndUserNotification implements the zia:index:EndUserNotification resource (singleton).
type EndUserNotification struct{}

// EndUserNotificationArgs are the inputs.
type EndUserNotificationArgs struct {
	AupFrequency                        *string `pulumi:"aupFrequency,optional"`
	AupCustomFrequency                  *int    `pulumi:"aupCustomFrequency,optional"`
	AupDayOffset                        *int    `pulumi:"aupDayOffset,optional"`
	AupMessage                          *string `pulumi:"aupMessage,optional"`
	NotificationType                    *string `pulumi:"notificationType,optional"`
	DisplayReason                       *bool   `pulumi:"displayReason,optional"`
	DisplayCompanyName                  *bool   `pulumi:"displayCompanyName,optional"`
	DisplayCompanyLogo                  *bool   `pulumi:"displayCompanyLogo,optional"`
	CustomText                          *string `pulumi:"customText,optional"`
	UrlCatReviewEnabled                 *bool   `pulumi:"urlCatReviewEnabled,optional"`
	UrlCatReviewSubmitToSecurityCloud    *bool   `pulumi:"urlCatReviewSubmitToSecurityCloud,optional"`
	UrlCatReviewCustomLocation          *string `pulumi:"urlCatReviewCustomLocation,optional"`
	UrlCatReviewText                    *string `pulumi:"urlCatReviewText,optional"`
	SecurityReviewEnabled               *bool   `pulumi:"securityReviewEnabled,optional"`
	SecurityReviewSubmitToSecurityCloud *bool   `pulumi:"securityReviewSubmitToSecurityCloud,optional"`
	SecurityReviewCustomLocation        *string `pulumi:"securityReviewCustomLocation,optional"`
	SecurityReviewText                  *string `pulumi:"securityReviewText,optional"`
	WebDlpReviewEnabled                 *bool   `pulumi:"webDlpReviewEnabled,optional"`
	WebDlpReviewSubmitToSecurityCloud   *bool   `pulumi:"webDlpReviewSubmitToSecurityCloud,optional"`
	WebDlpReviewCustomLocation          *string `pulumi:"webDlpReviewCustomLocation,optional"`
	WebDlpReviewText                    *string `pulumi:"webDlpReviewText,optional"`
	RedirectUrl                          *string `pulumi:"redirectUrl,optional"`
	SupportEmail                         *string `pulumi:"supportEmail,optional"`
	SupportPhone                         *string `pulumi:"supportPhone,optional"`
	OrgPolicyLink                        *string `pulumi:"orgPolicyLink,optional"`
	CautionAgainAfter                    *int    `pulumi:"cautionAgainAfter,optional"`
	CautionPerDomain                     *bool   `pulumi:"cautionPerDomain,optional"`
	CautionCustomText                    *string `pulumi:"cautionCustomText,optional"`
	IdpProxyNotificationText            *string `pulumi:"idpProxyNotificationText,optional"`
	QuarantineCustomNotificationText     *string `pulumi:"quarantineCustomNotificationText,optional"`
}

// EndUserNotificationState is the persisted state.
type EndUserNotificationState struct {
	EndUserNotificationArgs
}

const endUserNotificationID = "enduser_notification"

func endUserNotificationToAPI(args EndUserNotificationArgs) end_user_notification.UserNotificationSettings {
	return end_user_notification.UserNotificationSettings{
		AUPFrequency:                        ptrToString(args.AupFrequency),
		AUPCustomFrequency:                  ptrToIntDefault(args.AupCustomFrequency, 0),
		AUPDayOffset:                        ptrToIntDefault(args.AupDayOffset, 0),
		AUPMessage:                          ptrToString(args.AupMessage),
		NotificationType:                    ptrToString(args.NotificationType),
		DisplayReason:                       ptrToBool(args.DisplayReason),
		DisplayCompName:                     ptrToBool(args.DisplayCompanyName),
		DisplayCompLogo:                     ptrToBool(args.DisplayCompanyLogo),
		CustomText:                          ptrToString(args.CustomText),
		URLCatReviewEnabled:                 ptrToBool(args.UrlCatReviewEnabled),
		URLCatReviewSubmitToSecurityCloud:   ptrToBool(args.UrlCatReviewSubmitToSecurityCloud),
		URLCatReviewCustomLocation:          ptrToString(args.UrlCatReviewCustomLocation),
		URLCatReviewText:                    ptrToString(args.UrlCatReviewText),
		SecurityReviewEnabled:               ptrToBool(args.SecurityReviewEnabled),
		SecurityReviewSubmitToSecurityCloud: ptrToBool(args.SecurityReviewSubmitToSecurityCloud),
		SecurityReviewCustomLocation:        ptrToString(args.SecurityReviewCustomLocation),
		SecurityReviewText:                  ptrToString(args.SecurityReviewText),
		WebDLPReviewEnabled:                 ptrToBool(args.WebDlpReviewEnabled),
		WebDLPReviewSubmitToSecurityCloud:   ptrToBool(args.WebDlpReviewSubmitToSecurityCloud),
		WebDLPReviewCustomLocation:          ptrToString(args.WebDlpReviewCustomLocation),
		WebDLPReviewText:                    ptrToString(args.WebDlpReviewText),
		RedirectURL:                         ptrToString(args.RedirectUrl),
		SupportEmail:                        ptrToString(args.SupportEmail),
		SupportPhone:                        ptrToString(args.SupportPhone),
		OrgPolicyLink:                       ptrToString(args.OrgPolicyLink),
		CautionAgainAfter:                   ptrToIntDefault(args.CautionAgainAfter, 0),
		CautionPerDomain:                    ptrToBool(args.CautionPerDomain),
		CautionCustomText:                   ptrToString(args.CautionCustomText),
		IDPProxyNotificationText:            ptrToString(args.IdpProxyNotificationText),
		QuarantineCustomNotificationText:    ptrToString(args.QuarantineCustomNotificationText),
	}
}

func (EndUserNotification) Create(ctx context.Context, req infer.CreateRequest[EndUserNotificationArgs]) (infer.CreateResponse[EndUserNotificationState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[EndUserNotificationState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := endUserNotificationToAPI(req.Inputs)
	if _, _, err := end_user_notification.UpdateUserNotificationSettings(ctx, service, apiReq); err != nil {
		return infer.CreateResponse[EndUserNotificationState]{}, err
	}
	log.Printf("[INFO] Created/Updated ZIA End User Notification settings")

	time.Sleep(1 * time.Second)
	if shouldActivate() {
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[EndUserNotificationState]{}, activationErr
		}
	}

	state := EndUserNotificationState{EndUserNotificationArgs: req.Inputs}
	return infer.CreateResponse[EndUserNotificationState]{
		ID:     endUserNotificationID,
		Output: state,
	}, nil
}

func (EndUserNotification) Read(ctx context.Context, req infer.ReadRequest[EndUserNotificationArgs, EndUserNotificationState]) (infer.ReadResponse[EndUserNotificationArgs, EndUserNotificationState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[EndUserNotificationArgs, EndUserNotificationState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	resp, err := end_user_notification.GetUserNotificationSettings(ctx, service)
	if err != nil {
		return infer.ReadResponse[EndUserNotificationArgs, EndUserNotificationState]{}, err
	}

	args := EndUserNotificationArgs{
		AupFrequency:                        stringPtr(resp.AUPFrequency),
		AupCustomFrequency:                  intPtr(resp.AUPCustomFrequency),
		AupDayOffset:                        intPtr(resp.AUPDayOffset),
		AupMessage:                          stringPtr(resp.AUPMessage),
		NotificationType:                    stringPtr(resp.NotificationType),
		DisplayReason:                       boolPtr(resp.DisplayReason),
		DisplayCompanyName:                  boolPtr(resp.DisplayCompName),
		DisplayCompanyLogo:                  boolPtr(resp.DisplayCompLogo),
		CustomText:                          stringPtr(resp.CustomText),
		UrlCatReviewEnabled:                 boolPtr(resp.URLCatReviewEnabled),
		UrlCatReviewSubmitToSecurityCloud:   boolPtr(resp.URLCatReviewSubmitToSecurityCloud),
		UrlCatReviewCustomLocation:          stringPtr(resp.URLCatReviewCustomLocation),
		UrlCatReviewText:                    stringPtr(resp.URLCatReviewText),
		SecurityReviewEnabled:               boolPtr(resp.SecurityReviewEnabled),
		SecurityReviewSubmitToSecurityCloud: boolPtr(resp.SecurityReviewSubmitToSecurityCloud),
		SecurityReviewCustomLocation:        stringPtr(resp.SecurityReviewCustomLocation),
		SecurityReviewText:                  stringPtr(resp.SecurityReviewText),
		WebDlpReviewEnabled:                 boolPtr(resp.WebDLPReviewEnabled),
		WebDlpReviewSubmitToSecurityCloud:   boolPtr(resp.WebDLPReviewSubmitToSecurityCloud),
		WebDlpReviewCustomLocation:          stringPtr(resp.WebDLPReviewCustomLocation),
		WebDlpReviewText:                    stringPtr(resp.WebDLPReviewText),
		RedirectUrl:                         stringPtr(resp.RedirectURL),
		SupportEmail:                        stringPtr(resp.SupportEmail),
		SupportPhone:                        stringPtr(resp.SupportPhone),
		OrgPolicyLink:                       stringPtr(resp.OrgPolicyLink),
		CautionAgainAfter:                   intPtr(resp.CautionAgainAfter),
		CautionPerDomain:                    boolPtr(resp.CautionPerDomain),
		CautionCustomText:                   stringPtr(resp.CautionCustomText),
		IdpProxyNotificationText:            stringPtr(resp.IDPProxyNotificationText),
		QuarantineCustomNotificationText:    stringPtr(resp.QuarantineCustomNotificationText),
	}
	state := EndUserNotificationState{EndUserNotificationArgs: args}
	return infer.ReadResponse[EndUserNotificationArgs, EndUserNotificationState]{
		ID:     endUserNotificationID,
		Inputs: args,
		State:  state,
	}, nil
}

func (EndUserNotification) Update(ctx context.Context, req infer.UpdateRequest[EndUserNotificationArgs, EndUserNotificationState]) (infer.UpdateResponse[EndUserNotificationState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[EndUserNotificationState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := endUserNotificationToAPI(req.Inputs)
	if _, _, err := end_user_notification.UpdateUserNotificationSettings(ctx, service, apiReq); err != nil {
		return infer.UpdateResponse[EndUserNotificationState]{}, err
	}

	time.Sleep(1 * time.Second)
	if shouldActivate() {
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[EndUserNotificationState]{}, activationErr
		}
	}

	state := EndUserNotificationState{EndUserNotificationArgs: req.Inputs}
	return infer.UpdateResponse[EndUserNotificationState]{Output: state}, nil
}

func (EndUserNotification) Delete(ctx context.Context, req infer.DeleteRequest[EndUserNotificationState]) (infer.DeleteResponse, error) {
	log.Printf("[INFO] End User Notification: delete is a no-op for singleton resource")
	return infer.DeleteResponse{}, nil
}

func (EndUserNotification) Annotate(a infer.Annotator) {
	describeResource(a, &EndUserNotification{}, `The zia_end_user_notification resource manages end user notification settings in the Zscaler Internet Access (ZIA) cloud service. This is a singleton resource — only one end user notification configuration exists per tenant. It controls the messages and settings displayed to end users when their traffic is blocked, cautioned, or requires an Acceptable Use Policy (AUP) acknowledgment. Deleting the Pulumi resource does not remove the underlying settings.

{{% examples %}}
## Example Usage

{{% example %}}
### End User Notification

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.EndUserNotification("example", {
    aupFrequency: "ON_EVERY_LOGIN",
    notificationType: "CUSTOM",
    displayReason: true,
    displayCompanyName: true,
    supportEmail: "support@example.com",
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.EndUserNotification("example",
    aup_frequency="ON_EVERY_LOGIN",
    notification_type="CUSTOM",
    display_reason=True,
    display_company_name=True,
    support_email="support@example.com",
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:EndUserNotification
    properties:
      aupFrequency: ON_EVERY_LOGIN
      notificationType: CUSTOM
      displayReason: true
      displayCompanyName: true
      supportEmail: support@example.com
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

This is a singleton resource. Import is not applicable because there is no unique API identifier.
`)
}

func (a *EndUserNotificationArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.AupFrequency, "AUP display frequency. Valid values: `NEVER`, `ON_EVERY_LOGIN`, `ONCE`, `CUSTOM`.")
	ann.Describe(&a.AupCustomFrequency, "Custom frequency in days for displaying the AUP. Used when aupFrequency is `CUSTOM`.")
	ann.Describe(&a.AupDayOffset, "Day offset for AUP display.")
	ann.Describe(&a.AupMessage, "The Acceptable Use Policy message displayed to users.")
	ann.Describe(&a.NotificationType, "Notification type. Valid values: `DEFAULT`, `CUSTOM`.")
	ann.Describe(&a.DisplayReason, "Whether to display the reason for blocking or cautioning traffic.")
	ann.Describe(&a.DisplayCompanyName, "Whether to display the company name in the notification.")
	ann.Describe(&a.DisplayCompanyLogo, "Whether to display the company logo in the notification.")
	ann.Describe(&a.CustomText, "Custom text displayed in the notification page.")
	ann.Describe(&a.UrlCatReviewEnabled, "Whether URL category review requests are enabled.")
	ann.Describe(&a.UrlCatReviewSubmitToSecurityCloud, "Whether URL category review requests are submitted to the Zscaler security cloud.")
	ann.Describe(&a.UrlCatReviewCustomLocation, "Custom URL for URL category review submissions.")
	ann.Describe(&a.UrlCatReviewText, "Custom text for URL category review notifications.")
	ann.Describe(&a.SecurityReviewEnabled, "Whether security review requests are enabled.")
	ann.Describe(&a.SecurityReviewSubmitToSecurityCloud, "Whether security review requests are submitted to the Zscaler security cloud.")
	ann.Describe(&a.SecurityReviewCustomLocation, "Custom URL for security review submissions.")
	ann.Describe(&a.SecurityReviewText, "Custom text for security review notifications.")
	ann.Describe(&a.WebDlpReviewEnabled, "Whether Web DLP review requests are enabled.")
	ann.Describe(&a.WebDlpReviewSubmitToSecurityCloud, "Whether Web DLP review requests are submitted to the Zscaler security cloud.")
	ann.Describe(&a.WebDlpReviewCustomLocation, "Custom URL for Web DLP review submissions.")
	ann.Describe(&a.WebDlpReviewText, "Custom text for Web DLP review notifications.")
	ann.Describe(&a.RedirectUrl, "Redirect URL for the notification page.")
	ann.Describe(&a.SupportEmail, "Support email address displayed in notifications.")
	ann.Describe(&a.SupportPhone, "Support phone number displayed in notifications.")
	ann.Describe(&a.OrgPolicyLink, "Link to the organization's policy page.")
	ann.Describe(&a.CautionAgainAfter, "Time in minutes before showing the caution notification again.")
	ann.Describe(&a.CautionPerDomain, "Whether to show caution notifications per domain.")
	ann.Describe(&a.CautionCustomText, "Custom text for caution notifications.")
	ann.Describe(&a.IdpProxyNotificationText, "Custom text for IDP proxy notifications.")
	ann.Describe(&a.QuarantineCustomNotificationText, "Custom text for quarantine notifications.")
}

var _ infer.CustomResource[EndUserNotificationArgs, EndUserNotificationState] = EndUserNotification{}
