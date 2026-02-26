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

// Package provider implements the URL Filtering and Cloud App Settings resource.
// Adopted from terraform-provider-zia resource_zia_url_filtering_and_cloud_app_settings.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/urlfilteringpolicies"
)

const urlFilteringCloudAppSettingsID = "app_setting"

// UrlFilteringCloudAppSettings implements the zia:index:UrlFilteringCloudAppSettings resource.
// It is a singleton: there is only one global configuration. Delete is a no-op.
type UrlFilteringCloudAppSettings struct{}

// UrlFilteringCloudAppSettingsArgs are the inputs.
type UrlFilteringCloudAppSettingsArgs struct {
	EnableDynamicContentCat          *bool    `pulumi:"enableDynamicContentCat,optional"`
	ConsiderEmbeddedSites             *bool    `pulumi:"considerEmbeddedSites,optional"`
	EnforceSafeSearch                *bool    `pulumi:"enforceSafeSearch,optional"`
	SafeSearchApps                   []string `pulumi:"safeSearchApps,optional"`
	EnableOffice365                  *bool    `pulumi:"enableOffice365,optional"`
	EnableMsftO365                    *bool    `pulumi:"enableMsftO365,optional"`
	EnableUcaasZoom                   *bool    `pulumi:"enableUcaasZoom,optional"`
	EnableUcaasLogmein                *bool    `pulumi:"enableUcaasLogmein,optional"`
	EnableUcaasRingCentral            *bool    `pulumi:"enableUcaasRingCentral,optional"`
	EnableUcaasWebex                  *bool    `pulumi:"enableUcaasWebex,optional"`
	EnableUcaasTalkdesk               *bool    `pulumi:"enableUcaasTalkdesk,optional"`
	EnableChatgptPrompt               *bool    `pulumi:"enableChatgptPrompt,optional"`
	EnableMicrosoftCopilotPrompt      *bool    `pulumi:"enableMicrosoftCopilotPrompt,optional"`
	EnableGeminiPrompt                *bool    `pulumi:"enableGeminiPrompt,optional"`
	EnablePoepPrompt                  *bool    `pulumi:"enablePoepPrompt,optional"`
	EnableMetaPrompt                  *bool    `pulumi:"enableMetaPrompt,optional"`
	EnablePerPlexityPrompt            *bool    `pulumi:"enablePerPlexityPrompt,optional"`
	BlockSkype                       *bool    `pulumi:"blockSkype,optional"`
	EnableNewlyRegisteredDomains     *bool    `pulumi:"enableNewlyRegisteredDomains,optional"`
	EnableBlockOverrideForNonAuthUser *bool    `pulumi:"enableBlockOverrideForNonAuthUser,optional"`
	EnableCipaCompliance              *bool    `pulumi:"enableCipaCompliance,optional"`
}

// UrlFilteringCloudAppSettingsState is the persisted state.
type UrlFilteringCloudAppSettingsState struct {
	Id string `pulumi:"resourceId"` // Pulumi reserves "id" for resource identifier
	UrlFilteringCloudAppSettingsArgs
}

func (UrlFilteringCloudAppSettings) Create(ctx context.Context, req infer.CreateRequest[UrlFilteringCloudAppSettingsArgs]) (infer.CreateResponse[UrlFilteringCloudAppSettingsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[UrlFilteringCloudAppSettingsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	service := client.Service

	settings := argsToUrlAdvancedPolicySettings(req.Inputs)
	_, _, err := urlfilteringpolicies.UpdateUrlAndAppSettings(ctx, service, settings)
	if err != nil {
		return infer.CreateResponse[UrlFilteringCloudAppSettingsState]{}, err
	}

	if shouldActivate() {
		time.Sleep(1 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.CreateResponse[UrlFilteringCloudAppSettingsState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	state := UrlFilteringCloudAppSettingsState{Id: urlFilteringCloudAppSettingsID, UrlFilteringCloudAppSettingsArgs: req.Inputs}
	return infer.CreateResponse[UrlFilteringCloudAppSettingsState]{
		ID:     urlFilteringCloudAppSettingsID,
		Output: state,
	}, nil
}

func (UrlFilteringCloudAppSettings) Read(ctx context.Context, req infer.ReadRequest[UrlFilteringCloudAppSettingsArgs, UrlFilteringCloudAppSettingsState]) (infer.ReadResponse[UrlFilteringCloudAppSettingsArgs, UrlFilteringCloudAppSettingsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[UrlFilteringCloudAppSettingsArgs, UrlFilteringCloudAppSettingsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	resp, err := urlfilteringpolicies.GetUrlAndAppSettings(ctx, service)
	if err != nil {
		return infer.ReadResponse[UrlFilteringCloudAppSettingsArgs, UrlFilteringCloudAppSettingsState]{}, err
	}
	if resp == nil {
		return infer.ReadResponse[UrlFilteringCloudAppSettingsArgs, UrlFilteringCloudAppSettingsState]{}, nil
	}

	state := urlAdvancedPolicySettingsToState(resp)
	state.Id = urlFilteringCloudAppSettingsID
	return infer.ReadResponse[UrlFilteringCloudAppSettingsArgs, UrlFilteringCloudAppSettingsState]{
		ID:     urlFilteringCloudAppSettingsID,
		Inputs: state.UrlFilteringCloudAppSettingsArgs,
		State:  state,
	}, nil
}

func (UrlFilteringCloudAppSettings) Update(ctx context.Context, req infer.UpdateRequest[UrlFilteringCloudAppSettingsArgs, UrlFilteringCloudAppSettingsState]) (infer.UpdateResponse[UrlFilteringCloudAppSettingsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[UrlFilteringCloudAppSettingsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	service := client.Service

	settings := argsToUrlAdvancedPolicySettings(req.Inputs)
	_, _, err := urlfilteringpolicies.UpdateUrlAndAppSettings(ctx, service, settings)
	if err != nil {
		return infer.UpdateResponse[UrlFilteringCloudAppSettingsState]{}, err
	}

	time.Sleep(1 * time.Second)
	if shouldActivate() {
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[UrlFilteringCloudAppSettingsState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	state := UrlFilteringCloudAppSettingsState{Id: urlFilteringCloudAppSettingsID, UrlFilteringCloudAppSettingsArgs: req.Inputs}
	return infer.UpdateResponse[UrlFilteringCloudAppSettingsState]{Output: state}, nil
}

func (UrlFilteringCloudAppSettings) Delete(ctx context.Context, req infer.DeleteRequest[UrlFilteringCloudAppSettingsState]) (infer.DeleteResponse, error) {
	// No-op: singleton resource cannot be deleted
	return infer.DeleteResponse{}, nil
}

func (UrlFilteringCloudAppSettings) Diff(ctx context.Context, req infer.DiffRequest[UrlFilteringCloudAppSettingsArgs, UrlFilteringCloudAppSettingsState]) (infer.DiffResponse, error) {
	return infer.DiffResponse{}, nil
}

func argsToUrlAdvancedPolicySettings(args UrlFilteringCloudAppSettingsArgs) urlfilteringpolicies.URLAdvancedPolicySettings {
	return urlfilteringpolicies.URLAdvancedPolicySettings{
		EnableDynamicContentCat:          ptrToBool(args.EnableDynamicContentCat),
		ConsiderEmbeddedSites:            ptrToBool(args.ConsiderEmbeddedSites),
		EnforceSafeSearch:                ptrToBool(args.EnforceSafeSearch),
		SafeSearchApps:                   args.SafeSearchApps,
		EnableOffice365:                  ptrToBool(args.EnableOffice365),
		EnableMsftO365:                   ptrToBool(args.EnableMsftO365),
		EnableUcaasZoom:                  ptrToBool(args.EnableUcaasZoom),
		EnableUcaasLogMeIn:               ptrToBool(args.EnableUcaasLogmein),
		EnableUcaasRingCentral:           ptrToBool(args.EnableUcaasRingCentral),
		EnableUcaasWebex:                 ptrToBool(args.EnableUcaasWebex),
		EnableUcaasTalkdesk:              ptrToBool(args.EnableUcaasTalkdesk),
		EnableChatGptPrompt:              ptrToBool(args.EnableChatgptPrompt),
		EnableMicrosoftCoPilotPrompt:     ptrToBool(args.EnableMicrosoftCopilotPrompt),
		EnableGeminiPrompt:               ptrToBool(args.EnableGeminiPrompt),
		EnablePOEPrompt:                  ptrToBool(args.EnablePoepPrompt),
		EnableMetaPrompt:                 ptrToBool(args.EnableMetaPrompt),
		EnablePerPlexityPrompt:           ptrToBool(args.EnablePerPlexityPrompt),
		BlockSkype:                       ptrToBool(args.BlockSkype),
		EnableNewlyRegisteredDomains:     ptrToBool(args.EnableNewlyRegisteredDomains),
		EnableBlockOverrideForNonAuthUser: ptrToBool(args.EnableBlockOverrideForNonAuthUser),
		EnableCIPACompliance:             ptrToBool(args.EnableCipaCompliance),
	}
}

func urlAdvancedPolicySettingsToState(r *urlfilteringpolicies.URLAdvancedPolicySettings) UrlFilteringCloudAppSettingsState {
	return UrlFilteringCloudAppSettingsState{
		Id: urlFilteringCloudAppSettingsID,
		UrlFilteringCloudAppSettingsArgs: UrlFilteringCloudAppSettingsArgs{
			EnableDynamicContentCat:          boolPtr(r.EnableDynamicContentCat),
			ConsiderEmbeddedSites:            boolPtr(r.ConsiderEmbeddedSites),
			EnforceSafeSearch:                boolPtr(r.EnforceSafeSearch),
			SafeSearchApps:                   r.SafeSearchApps,
			EnableOffice365:                  boolPtr(r.EnableOffice365),
			EnableMsftO365:                   boolPtr(r.EnableMsftO365),
			EnableUcaasZoom:                  boolPtr(r.EnableUcaasZoom),
			EnableUcaasLogmein:               boolPtr(r.EnableUcaasLogMeIn),
			EnableUcaasRingCentral:           boolPtr(r.EnableUcaasRingCentral),
			EnableUcaasWebex:                 boolPtr(r.EnableUcaasWebex),
			EnableUcaasTalkdesk:              boolPtr(r.EnableUcaasTalkdesk),
			EnableChatgptPrompt:              boolPtr(r.EnableChatGptPrompt),
			EnableMicrosoftCopilotPrompt:     boolPtr(r.EnableMicrosoftCoPilotPrompt),
			EnableGeminiPrompt:               boolPtr(r.EnableGeminiPrompt),
			EnablePoepPrompt:                 boolPtr(r.EnablePOEPrompt),
			EnableMetaPrompt:                 boolPtr(r.EnableMetaPrompt),
			EnablePerPlexityPrompt:           boolPtr(r.EnablePerPlexityPrompt),
			BlockSkype:                       boolPtr(r.BlockSkype),
			EnableNewlyRegisteredDomains:     boolPtr(r.EnableNewlyRegisteredDomains),
			EnableBlockOverrideForNonAuthUser: boolPtr(r.EnableBlockOverrideForNonAuthUser),
			EnableCipaCompliance:             boolPtr(r.EnableCIPACompliance),
		},
	}
}

func (UrlFilteringCloudAppSettings) Annotate(a infer.Annotator) {
	describeResource(a, &UrlFilteringCloudAppSettings{}, `The zia_url_filtering_and_cloud_app_settings resource manages URL filtering and cloud application settings in the Zscaler Internet Access (ZIA) cloud service. This is a singleton resource that controls global settings for URL filtering features such as safe search enforcement, UCaaS application controls, AI/ML prompt visibility, and CIPA compliance.

For more information, see the [ZIA URL Filtering documentation](https://help.zscaler.com/zia/url-filtering).

{{% examples %}}
## Example Usage

{{% example %}}
### Configure URL Filtering and Cloud App Settings

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.UrlFilteringCloudAppSettings("example", {
    enableDynamicContentCat: true,
    enforceSafeSearch: true,
    enableOffice365: true,
    enableChatgptPrompt: true,
    enableCipaCompliance: false,
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.UrlFilteringCloudAppSettings("example",
    enable_dynamic_content_cat=True,
    enforce_safe_search=True,
    enable_office365=True,
    enable_chatgpt_prompt=True,
    enable_cipa_compliance=False,
)
`+tripleBacktick("")+`

`+tripleBacktick("go")+`
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	zia "github.com/zscaler/pulumi-zia/sdk/go/pulumi-zia"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := zia.NewUrlFilteringCloudAppSettings(ctx, "example", &zia.UrlFilteringCloudAppSettingsArgs{
			EnableDynamicContentCat: pulumi.BoolRef(true),
			EnforceSafeSearch:       pulumi.BoolRef(true),
			EnableOffice365:         pulumi.BoolRef(true),
			EnableChatgptPrompt:     pulumi.BoolRef(true),
			EnableCipaCompliance:    pulumi.BoolRef(false),
		})
		return err
	})
}
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:UrlFilteringCloudAppSettings
    properties:
      enableDynamicContentCat: true
      enforceSafeSearch: true
      enableOffice365: true
      enableChatgptPrompt: true
      enableCipaCompliance: false
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

The singleton URL Filtering Cloud App Settings resource can be imported using its fixed ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:UrlFilteringCloudAppSettings example app_setting
`+tripleBacktick("")+`
`)
}

func (a *UrlFilteringCloudAppSettingsArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.EnableDynamicContentCat, "If true, dynamic content categorization is enabled.")
	ann.Describe(&a.ConsiderEmbeddedSites, "If true, embedded sites within web pages are considered for URL filtering.")
	ann.Describe(&a.EnforceSafeSearch, "If true, safe search is enforced for supported search engines.")
	ann.Describe(&a.SafeSearchApps, "List of application names for which safe search is enforced.")
	ann.Describe(&a.EnableOffice365, "If true, Office 365 one-click configuration is enabled.")
	ann.Describe(&a.EnableMsftO365, "If true, Microsoft Office 365 optimization is enabled.")
	ann.Describe(&a.EnableUcaasZoom, "If true, UCaaS controls for Zoom are enabled.")
	ann.Describe(&a.EnableUcaasLogmein, "If true, UCaaS controls for LogMeIn are enabled.")
	ann.Describe(&a.EnableUcaasRingCentral, "If true, UCaaS controls for RingCentral are enabled.")
	ann.Describe(&a.EnableUcaasWebex, "If true, UCaaS controls for Webex are enabled.")
	ann.Describe(&a.EnableUcaasTalkdesk, "If true, UCaaS controls for Talkdesk are enabled.")
	ann.Describe(&a.EnableChatgptPrompt, "If true, ChatGPT prompt visibility and logging is enabled.")
	ann.Describe(&a.EnableMicrosoftCopilotPrompt, "If true, Microsoft Copilot prompt visibility and logging is enabled.")
	ann.Describe(&a.EnableGeminiPrompt, "If true, Google Gemini prompt visibility and logging is enabled.")
	ann.Describe(&a.EnablePoepPrompt, "If true, POE prompt visibility and logging is enabled.")
	ann.Describe(&a.EnableMetaPrompt, "If true, Meta AI prompt visibility and logging is enabled.")
	ann.Describe(&a.EnablePerPlexityPrompt, "If true, Perplexity AI prompt visibility and logging is enabled.")
	ann.Describe(&a.BlockSkype, "If true, Skype is blocked.")
	ann.Describe(&a.EnableNewlyRegisteredDomains, "If true, newly registered domains detection is enabled.")
	ann.Describe(&a.EnableBlockOverrideForNonAuthUser, "If true, block override is enabled for non-authenticated users.")
	ann.Describe(&a.EnableCipaCompliance, "If true, CIPA (Children's Internet Protection Act) compliance mode is enabled.")
}

func (s *UrlFilteringCloudAppSettingsState) Annotate(a infer.Annotator) {
	a.Describe(&s.Id, "The fixed resource ID of the singleton URL filtering cloud app settings.")
}

var _ infer.CustomResource[UrlFilteringCloudAppSettingsArgs, UrlFilteringCloudAppSettingsState] = UrlFilteringCloudAppSettings{}
