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

// Package provider implements the ATP Security Exceptions resource.
// Adopted from terraform-provider-zia resource_zia_atp_security_exceptions.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/advancedthreatsettings"
)

const atpSecurityExceptionsID = "bypass_url"

// AtpSecurityExceptions implements the zia:index:AtpSecurityExceptions resource.
// Manages ATP bypass URLs. Singleton. Delete is a no-op.
type AtpSecurityExceptions struct{}

// AtpSecurityExceptionsArgs are the inputs.
type AtpSecurityExceptionsArgs struct {
	BypassUrls []string `pulumi:"bypassUrls,optional"`
}

// AtpSecurityExceptionsState is the persisted state.
type AtpSecurityExceptionsState struct {
	AtpSecurityExceptionsArgs
	ResourceId string `pulumi:"resourceId"`
}

func (AtpSecurityExceptions) Create(ctx context.Context, req infer.CreateRequest[AtpSecurityExceptionsArgs]) (infer.CreateResponse[AtpSecurityExceptionsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[AtpSecurityExceptionsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	reqAPI := advancedthreatsettings.SecurityExceptions{BypassUrls: req.Inputs.BypassUrls}
	if _, err := advancedthreatsettings.UpdateSecurityExceptions(ctx, service, reqAPI); err != nil {
		return infer.CreateResponse[AtpSecurityExceptionsState]{}, err
	}

	time.Sleep(1 * time.Second)
	if shouldActivate() {
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[AtpSecurityExceptionsState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	state := AtpSecurityExceptionsState{
		AtpSecurityExceptionsArgs: req.Inputs,
		ResourceId:                atpSecurityExceptionsID,
	}
	return infer.CreateResponse[AtpSecurityExceptionsState]{
		ID:     atpSecurityExceptionsID,
		Output: state,
	}, nil
}

func (AtpSecurityExceptions) Read(ctx context.Context, req infer.ReadRequest[AtpSecurityExceptionsArgs, AtpSecurityExceptionsState]) (infer.ReadResponse[AtpSecurityExceptionsArgs, AtpSecurityExceptionsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[AtpSecurityExceptionsArgs, AtpSecurityExceptionsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	resp, err := advancedthreatsettings.GetSecurityExceptions(ctx, service)
	if err != nil {
		return infer.ReadResponse[AtpSecurityExceptionsArgs, AtpSecurityExceptionsState]{}, err
	}
	if resp == nil {
		return infer.ReadResponse[AtpSecurityExceptionsArgs, AtpSecurityExceptionsState]{}, fmt.Errorf("couldn't read bypass URLs")
	}

	state := AtpSecurityExceptionsState{
		AtpSecurityExceptionsArgs: AtpSecurityExceptionsArgs{BypassUrls: resp.BypassUrls},
		ResourceId:                atpSecurityExceptionsID,
	}
	return infer.ReadResponse[AtpSecurityExceptionsArgs, AtpSecurityExceptionsState]{
		ID:     atpSecurityExceptionsID,
		Inputs: state.AtpSecurityExceptionsArgs,
		State:  state,
	}, nil
}

func (AtpSecurityExceptions) Update(ctx context.Context, req infer.UpdateRequest[AtpSecurityExceptionsArgs, AtpSecurityExceptionsState]) (infer.UpdateResponse[AtpSecurityExceptionsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[AtpSecurityExceptionsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	reqAPI := advancedthreatsettings.SecurityExceptions{BypassUrls: req.Inputs.BypassUrls}
	if _, err := advancedthreatsettings.UpdateSecurityExceptions(ctx, service, reqAPI); err != nil {
		return infer.UpdateResponse[AtpSecurityExceptionsState]{}, err
	}

	time.Sleep(1 * time.Second)
	if shouldActivate() {
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[AtpSecurityExceptionsState]{}, activationErr
		}
	}

	state := AtpSecurityExceptionsState{
		AtpSecurityExceptionsArgs: req.Inputs,
		ResourceId:                atpSecurityExceptionsID,
	}
	return infer.UpdateResponse[AtpSecurityExceptionsState]{Output: state}, nil
}

func (AtpSecurityExceptions) Delete(ctx context.Context, req infer.DeleteRequest[AtpSecurityExceptionsState]) (infer.DeleteResponse, error) {
	return infer.DeleteResponse{}, nil
}

func (AtpSecurityExceptions) Annotate(a infer.Annotator) {
	describeResource(a, &AtpSecurityExceptions{}, `The zia_atp_security_exceptions resource manages the list of bypass URLs for Advanced Threat Protection (ATP) in the Zscaler Internet Access (ZIA) cloud service. URLs added to this list are excluded from ATP scanning. This is a singleton resource.

For more information, see the [ZIA Advanced Threat Protection documentation](https://help.zscaler.com/zia/about-advanced-threat-protection-policy).

{{% examples %}}
## Example Usage

{{% example %}}
### Configure ATP Security Exceptions

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.AtpSecurityExceptions("example", {
    bypassUrls: [
        "trusted-partner.com",
        "internal-app.example.org",
    ],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.AtpSecurityExceptions("example",
    bypass_urls=[
        "trusted-partner.com",
        "internal-app.example.org",
    ],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:AtpSecurityExceptions
    properties:
      bypassUrls:
        - trusted-partner.com
        - internal-app.example.org
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

> This is a singleton resource. Import is not applicable.
`)
}

func (a *AtpSecurityExceptionsArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.BypassUrls, "List of URLs to be excluded (bypassed) from Advanced Threat Protection scanning.")
}

func (s *AtpSecurityExceptionsState) Annotate(a infer.Annotator) {
	a.Describe(&s.ResourceId, "The internal resource identifier for the ATP security exceptions.")
}

var _ infer.CustomResource[AtpSecurityExceptionsArgs, AtpSecurityExceptionsState] = AtpSecurityExceptions{}
