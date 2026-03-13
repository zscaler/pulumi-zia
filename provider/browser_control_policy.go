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

// Package provider implements the Browser Control Policy resource.
// Adopted from terraform-provider-zia resource_zia_browser_control_policy.go.
// Singleton resource (one per tenant). Delete is a no-op.

package provider

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/browser_control_settings"
)

// BrowserControlPolicy implements the zia:index:BrowserControlPolicy resource.
type BrowserControlPolicy struct{}

// SmartIsolationProfileInput is a nested input for the isolation profile (single object with id).
type SmartIsolationProfileInput struct {
	Id *string `pulumi:"id,optional"`
}

// BrowserControlPolicyArgs are the inputs.
type BrowserControlPolicyArgs struct {
	PluginCheckFrequency            *string                     `pulumi:"pluginCheckFrequency,optional"`
	BypassPlugins                   []string                    `pulumi:"bypassPlugins,optional"`
	BypassApplications              []string                    `pulumi:"bypassApplications,optional"`
	BlockedInternetExplorerVersions []string                    `pulumi:"blockedInternetExplorerVersions,optional"`
	BlockedChromeVersions           []string                    `pulumi:"blockedChromeVersions,optional"`
	BlockedFirefoxVersions          []string                    `pulumi:"blockedFirefoxVersions,optional"`
	BlockedSafariVersions           []string                    `pulumi:"blockedSafariVersions,optional"`
	BlockedOperaVersions            []string                    `pulumi:"blockedOperaVersions,optional"`
	BypassAllBrowsers               *bool                       `pulumi:"bypassAllBrowsers,optional"`
	AllowAllBrowsers                *bool                       `pulumi:"allowAllBrowsers,optional"`
	EnableWarnings                  *bool                       `pulumi:"enableWarnings,optional"`
	EnableSmartBrowserIsolation     *bool                       `pulumi:"enableSmartBrowserIsolation,optional"`
	SmartIsolationProfile           *SmartIsolationProfileInput `pulumi:"smartIsolationProfile,optional"`
	SmartIsolationGroups            []int                       `pulumi:"smartIsolationGroups,optional"`
	SmartIsolationUsers             []int                       `pulumi:"smartIsolationUsers,optional"`
}

// BrowserControlPolicyState is the persisted state.
type BrowserControlPolicyState struct {
	BrowserControlPolicyArgs
}

const browserControlPolicyID = "browser_settings"

func browserControlPolicyToAPI(args BrowserControlPolicyArgs) browser_control_settings.BrowserControlSettings {
	profile := browser_control_settings.SmartIsolationProfile{}
	if args.SmartIsolationProfile != nil && args.SmartIsolationProfile.Id != nil {
		profile.ID = *args.SmartIsolationProfile.Id
	}
	return browser_control_settings.BrowserControlSettings{
		PluginCheckFrequency:            ptrToString(args.PluginCheckFrequency),
		BypassPlugins:                   args.BypassPlugins,
		BypassApplications:              args.BypassApplications,
		BlockedInternetExplorerVersions: args.BlockedInternetExplorerVersions,
		BlockedChromeVersions:           args.BlockedChromeVersions,
		BlockedFirefoxVersions:          args.BlockedFirefoxVersions,
		BlockedSafariVersions:           args.BlockedSafariVersions,
		BlockedOperaVersions:            args.BlockedOperaVersions,
		BypassAllBrowsers:               ptrToBool(args.BypassAllBrowsers),
		AllowAllBrowsers:                ptrToBool(args.AllowAllBrowsers),
		EnableWarnings:                  ptrToBool(args.EnableWarnings),
		EnableSmartBrowserIsolation:     ptrToBool(args.EnableSmartBrowserIsolation),
		SmartIsolationProfile:           profile,
		SmartIsolationGroups:            idsToIDNameExtensions(args.SmartIsolationGroups),
		SmartIsolationUsers:             idsToIDNameExtensions(args.SmartIsolationUsers),
	}
}

func (BrowserControlPolicy) Create(ctx context.Context, req infer.CreateRequest[BrowserControlPolicyArgs]) (infer.CreateResponse[BrowserControlPolicyState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[BrowserControlPolicyState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	policy := browserControlPolicyToAPI(req.Inputs)
	if _, _, err := browser_control_settings.UpdateBrowserControlSettings(ctx, service, policy); err != nil {
		return infer.CreateResponse[BrowserControlPolicyState]{}, err
	}

	time.Sleep(1 * time.Second)
	if shouldActivate() {
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[BrowserControlPolicyState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	return infer.CreateResponse[BrowserControlPolicyState]{
		ID:     browserControlPolicyID,
		Output: BrowserControlPolicyState{BrowserControlPolicyArgs: req.Inputs},
	}, nil
}

func (BrowserControlPolicy) Read(ctx context.Context, req infer.ReadRequest[BrowserControlPolicyArgs, BrowserControlPolicyState]) (infer.ReadResponse[BrowserControlPolicyArgs, BrowserControlPolicyState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[BrowserControlPolicyArgs, BrowserControlPolicyState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	resp, err := browser_control_settings.GetBrowserControlSettings(ctx, service)
	if err != nil {
		return infer.ReadResponse[BrowserControlPolicyArgs, BrowserControlPolicyState]{}, err
	}
	if resp == nil {
		return infer.ReadResponse[BrowserControlPolicyArgs, BrowserControlPolicyState]{}, nil
	}

	var smartProfile *SmartIsolationProfileInput
	if resp.SmartIsolationProfile.ID != "" {
		smartProfile = &SmartIsolationProfileInput{Id: stringPtr(resp.SmartIsolationProfile.ID)}
	}
	args := BrowserControlPolicyArgs{
		PluginCheckFrequency:            stringPtr(resp.PluginCheckFrequency),
		BypassPlugins:                   resp.BypassPlugins,
		BypassApplications:              resp.BypassApplications,
		BlockedInternetExplorerVersions: resp.BlockedInternetExplorerVersions,
		BlockedChromeVersions:           resp.BlockedChromeVersions,
		BlockedFirefoxVersions:          resp.BlockedFirefoxVersions,
		BlockedSafariVersions:           resp.BlockedSafariVersions,
		BlockedOperaVersions:            resp.BlockedOperaVersions,
		BypassAllBrowsers:               boolPtr(resp.BypassAllBrowsers),
		AllowAllBrowsers:                boolPtr(resp.AllowAllBrowsers),
		EnableWarnings:                  boolPtr(resp.EnableWarnings),
		EnableSmartBrowserIsolation:     boolPtr(resp.EnableSmartBrowserIsolation),
		SmartIsolationProfile:           smartProfile,
		SmartIsolationGroups:            idNameExtensionsToIDs(resp.SmartIsolationGroups),
		SmartIsolationUsers:             idNameExtensionsToIDs(resp.SmartIsolationUsers),
	}
	state := BrowserControlPolicyState{BrowserControlPolicyArgs: args}
	return infer.ReadResponse[BrowserControlPolicyArgs, BrowserControlPolicyState]{
		ID:     browserControlPolicyID,
		Inputs: args,
		State:  state,
	}, nil
}

func (BrowserControlPolicy) Update(ctx context.Context, req infer.UpdateRequest[BrowserControlPolicyArgs, BrowserControlPolicyState]) (infer.UpdateResponse[BrowserControlPolicyState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[BrowserControlPolicyState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	policy := browserControlPolicyToAPI(req.Inputs)
	if _, _, err := browser_control_settings.UpdateBrowserControlSettings(ctx, service, policy); err != nil {
		return infer.UpdateResponse[BrowserControlPolicyState]{}, err
	}

	time.Sleep(1 * time.Second)
	if shouldActivate() {
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[BrowserControlPolicyState]{}, activationErr
		}
	}

	return infer.UpdateResponse[BrowserControlPolicyState]{
		Output: BrowserControlPolicyState{BrowserControlPolicyArgs: req.Inputs},
	}, nil
}

func (BrowserControlPolicy) Delete(ctx context.Context, req infer.DeleteRequest[BrowserControlPolicyState]) (infer.DeleteResponse, error) {
	// No-op: singleton policy; deleting the Pulumi resource does not remove the underlying settings
	return infer.DeleteResponse{}, nil
}

func (BrowserControlPolicy) Annotate(a infer.Annotator) {
	describeResource(a, &BrowserControlPolicy{}, `The zia_browser_control_policy resource manages browser control policy settings in the Zscaler Internet Access (ZIA) cloud service. This is a singleton resource — only one browser control policy exists per tenant. Deleting the Pulumi resource does not remove the underlying settings.

{{% examples %}}
## Example Usage

{{% example %}}
### Browser Control Policy

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.BrowserControlPolicy("example", {
    allowAllBrowsers: true,
    enableWarnings: true,
    enableSmartBrowserIsolation: false,
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.BrowserControlPolicy("example",
    allow_all_browsers=True,
    enable_warnings=True,
    enable_smart_browser_isolation=False,
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:BrowserControlPolicy
    properties:
      allowAllBrowsers: true
      enableWarnings: true
      enableSmartBrowserIsolation: false
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

This is a singleton resource. Import is not applicable because there is no unique API identifier.
`)
}

func (a *BrowserControlPolicyArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.PluginCheckFrequency, "How often to check for browser plugins. Valid values: `NEVER`, `ALWAYS`, `DAILY`, `WEEKLY`, `MONTHLY`.")
	ann.Describe(&a.BypassPlugins, "List of plugins to bypass.")
	ann.Describe(&a.BypassApplications, "List of applications to bypass.")
	ann.Describe(&a.BlockedInternetExplorerVersions, "List of blocked Internet Explorer versions.")
	ann.Describe(&a.BlockedChromeVersions, "List of blocked Chrome versions.")
	ann.Describe(&a.BlockedFirefoxVersions, "List of blocked Firefox versions.")
	ann.Describe(&a.BlockedSafariVersions, "List of blocked Safari versions.")
	ann.Describe(&a.BlockedOperaVersions, "List of blocked Opera versions.")
	ann.Describe(&a.BypassAllBrowsers, "Whether to bypass all browsers.")
	ann.Describe(&a.AllowAllBrowsers, "Whether to allow all browsers.")
	ann.Describe(&a.EnableWarnings, "Whether to enable browser warnings for unsupported browsers.")
	ann.Describe(&a.EnableSmartBrowserIsolation, "Whether to enable Smart Browser Isolation.")
	ann.Describe(&a.SmartIsolationProfile, "The Cloud Browser Isolation profile to use for Smart Browser Isolation.")
	ann.Describe(&a.SmartIsolationGroups, "IDs of groups for Smart Browser Isolation.")
	ann.Describe(&a.SmartIsolationUsers, "IDs of users for Smart Browser Isolation.")
}

var _ infer.CustomResource[BrowserControlPolicyArgs, BrowserControlPolicyState] = BrowserControlPolicy{}
