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

// Package provider implements the Security Policy Settings resource.
// Adopted from terraform-provider-zia resource_zia_security_policy_settings.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/security_policy_settings"
)

const securityPolicySettingsID = "all_urls"

// SecurityPolicySettings implements the zia:index:SecurityPolicySettings resource.
// Singleton: manages whitelist/blacklist URLs. Delete is a no-op.
type SecurityPolicySettings struct{}

// SecurityPolicySettingsArgs are the inputs.
type SecurityPolicySettingsArgs struct {
	WhitelistUrls []string `pulumi:"whitelistUrls,optional"`
	BlacklistUrls []string `pulumi:"blacklistUrls,optional"`
}

// SecurityPolicySettingsState is the persisted state.
type SecurityPolicySettingsState struct {
	SecurityPolicySettingsArgs
	Id string `pulumi:"resourceId"` // Pulumi reserves "id" for resource identifier
}

func (SecurityPolicySettings) Create(ctx context.Context, req infer.CreateRequest[SecurityPolicySettingsArgs]) (infer.CreateResponse[SecurityPolicySettingsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[SecurityPolicySettingsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	listUrls := security_policy_settings.ListUrls{
		White: req.Inputs.WhitelistUrls,
		Black: req.Inputs.BlacklistUrls,
	}
	if _, err := security_policy_settings.UpdateListUrls(ctx, service, listUrls); err != nil {
		return infer.CreateResponse[SecurityPolicySettingsState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[SecurityPolicySettingsState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	state := SecurityPolicySettingsState{
		Id: securityPolicySettingsID,
		SecurityPolicySettingsArgs: req.Inputs,
	}
	return infer.CreateResponse[SecurityPolicySettingsState]{
		ID:     securityPolicySettingsID,
		Output: state,
	}, nil
}

func (SecurityPolicySettings) Read(ctx context.Context, req infer.ReadRequest[SecurityPolicySettingsArgs, SecurityPolicySettingsState]) (infer.ReadResponse[SecurityPolicySettingsArgs, SecurityPolicySettingsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[SecurityPolicySettingsArgs, SecurityPolicySettingsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	resp, err := security_policy_settings.GetListUrls(ctx, service)
	if err != nil {
		return infer.ReadResponse[SecurityPolicySettingsArgs, SecurityPolicySettingsState]{}, err
	}
	if resp == nil {
		return infer.ReadResponse[SecurityPolicySettingsArgs, SecurityPolicySettingsState]{}, fmt.Errorf("couldn't read security policy settings")
	}

	state := SecurityPolicySettingsState{
		Id: securityPolicySettingsID,
		SecurityPolicySettingsArgs: SecurityPolicySettingsArgs{
			WhitelistUrls: resp.White,
			BlacklistUrls: resp.Black,
		},
	}
	return infer.ReadResponse[SecurityPolicySettingsArgs, SecurityPolicySettingsState]{
		ID:     securityPolicySettingsID,
		Inputs: state.SecurityPolicySettingsArgs,
		State:  state,
	}, nil
}

func (SecurityPolicySettings) Update(ctx context.Context, req infer.UpdateRequest[SecurityPolicySettingsArgs, SecurityPolicySettingsState]) (infer.UpdateResponse[SecurityPolicySettingsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[SecurityPolicySettingsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	listUrls := security_policy_settings.ListUrls{
		White: req.Inputs.WhitelistUrls,
		Black: req.Inputs.BlacklistUrls,
	}
	if _, err := security_policy_settings.UpdateListUrls(ctx, service, listUrls); err != nil {
		return infer.UpdateResponse[SecurityPolicySettingsState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[SecurityPolicySettingsState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	state := SecurityPolicySettingsState{
		Id: securityPolicySettingsID,
		SecurityPolicySettingsArgs: req.Inputs,
	}
	return infer.UpdateResponse[SecurityPolicySettingsState]{Output: state}, nil
}

func (SecurityPolicySettings) Delete(ctx context.Context, req infer.DeleteRequest[SecurityPolicySettingsState]) (infer.DeleteResponse, error) {
	// No-op: singleton cannot be deleted
	return infer.DeleteResponse{}, nil
}

func (SecurityPolicySettings) Annotate(a infer.Annotator) {
	describeResource(a, &SecurityPolicySettings{}, `The zia_security_policy_settings resource manages the whitelist and blacklist URL configuration for the ZIA security policy. This is a singleton resource that controls which URLs are always allowed (whitelisted) or always blocked (blacklisted) across the organization.

For more information, see the [ZIA Security Policy Settings documentation](https://help.zscaler.com/zia/configuring-security-policy).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Security Policy Settings

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.SecurityPolicySettings("example", {
    whitelistUrls: [
        "example.com",
        "trusted-site.org",
    ],
    blacklistUrls: [
        "malicious-site.com",
        "phishing-site.net",
    ],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.SecurityPolicySettings("example",
    whitelist_urls=[
        "example.com",
        "trusted-site.org",
    ],
    blacklist_urls=[
        "malicious-site.com",
        "phishing-site.net",
    ],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:SecurityPolicySettings
    properties:
      whitelistUrls:
        - example.com
        - trusted-site.org
      blacklistUrls:
        - malicious-site.com
        - phishing-site.net
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

> This is a singleton resource. Import is not applicable.
`)
}

func (a *SecurityPolicySettingsArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.WhitelistUrls, "List of URLs that are always allowed (whitelisted) by the security policy.")
	ann.Describe(&a.BlacklistUrls, "List of URLs that are always blocked (blacklisted) by the security policy.")
}

func (s *SecurityPolicySettingsState) Annotate(a infer.Annotator) {
	a.Describe(&s.Id, "The internal resource identifier for the security policy settings.")
}

var _ infer.CustomResource[SecurityPolicySettingsArgs, SecurityPolicySettingsState] = SecurityPolicySettings{}
