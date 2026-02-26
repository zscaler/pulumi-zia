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

// Package provider implements the ATP Malicious URLs resource.
// Adopted from terraform-provider-zia resource_zia_atp_malicious_urls.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/advancedthreatsettings"
)

const atpMaliciousUrlsID = "all_urls"

// AtpMaliciousUrls implements the zia:index:AtpMaliciousUrls resource.
// Manages ATP malicious URL list. Singleton. Delete is a no-op.
type AtpMaliciousUrls struct{}

// AtpMaliciousUrlsArgs are the inputs.
type AtpMaliciousUrlsArgs struct {
	MaliciousUrls []string `pulumi:"maliciousUrls,optional"`
}

// AtpMaliciousUrlsState is the persisted state.
type AtpMaliciousUrlsState struct {
	AtpMaliciousUrlsArgs
	ResourceId string `pulumi:"resourceId"`
}

func (AtpMaliciousUrls) Create(ctx context.Context, req infer.CreateRequest[AtpMaliciousUrlsArgs]) (infer.CreateResponse[AtpMaliciousUrlsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[AtpMaliciousUrlsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	reqAPI := advancedthreatsettings.MaliciousURLs{MaliciousUrls: req.Inputs.MaliciousUrls}
	if _, err := advancedthreatsettings.UpdateMaliciousURLs(ctx, service, reqAPI); err != nil {
		return infer.CreateResponse[AtpMaliciousUrlsState]{}, err
	}

	time.Sleep(1 * time.Second)
	if shouldActivate() {
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[AtpMaliciousUrlsState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	state := AtpMaliciousUrlsState{
		AtpMaliciousUrlsArgs: req.Inputs,
		ResourceId:          atpMaliciousUrlsID,
	}
	return infer.CreateResponse[AtpMaliciousUrlsState]{
		ID:     atpMaliciousUrlsID,
		Output: state,
	}, nil
}

func (AtpMaliciousUrls) Read(ctx context.Context, req infer.ReadRequest[AtpMaliciousUrlsArgs, AtpMaliciousUrlsState]) (infer.ReadResponse[AtpMaliciousUrlsArgs, AtpMaliciousUrlsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[AtpMaliciousUrlsArgs, AtpMaliciousUrlsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	resp, err := advancedthreatsettings.GetMaliciousURLs(ctx, service)
	if err != nil {
		return infer.ReadResponse[AtpMaliciousUrlsArgs, AtpMaliciousUrlsState]{}, err
	}
	if resp == nil {
		return infer.ReadResponse[AtpMaliciousUrlsArgs, AtpMaliciousUrlsState]{}, fmt.Errorf("couldn't read malicious URLs")
	}

	state := AtpMaliciousUrlsState{
		AtpMaliciousUrlsArgs: AtpMaliciousUrlsArgs{MaliciousUrls: resp.MaliciousUrls},
		ResourceId:           atpMaliciousUrlsID,
	}
	return infer.ReadResponse[AtpMaliciousUrlsArgs, AtpMaliciousUrlsState]{
		ID:     atpMaliciousUrlsID,
		Inputs: state.AtpMaliciousUrlsArgs,
		State:  state,
	}, nil
}

func (AtpMaliciousUrls) Update(ctx context.Context, req infer.UpdateRequest[AtpMaliciousUrlsArgs, AtpMaliciousUrlsState]) (infer.UpdateResponse[AtpMaliciousUrlsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[AtpMaliciousUrlsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	reqAPI := advancedthreatsettings.MaliciousURLs{MaliciousUrls: req.Inputs.MaliciousUrls}
	if _, err := advancedthreatsettings.UpdateMaliciousURLs(ctx, service, reqAPI); err != nil {
		return infer.UpdateResponse[AtpMaliciousUrlsState]{}, err
	}

	time.Sleep(1 * time.Second)
	if shouldActivate() {
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[AtpMaliciousUrlsState]{}, activationErr
		}
	}

	state := AtpMaliciousUrlsState{
		AtpMaliciousUrlsArgs: req.Inputs,
		ResourceId:           atpMaliciousUrlsID,
	}
	return infer.UpdateResponse[AtpMaliciousUrlsState]{Output: state}, nil
}

func (AtpMaliciousUrls) Delete(ctx context.Context, req infer.DeleteRequest[AtpMaliciousUrlsState]) (infer.DeleteResponse, error) {
	return infer.DeleteResponse{}, nil
}

func (AtpMaliciousUrls) Annotate(a infer.Annotator) {
	describeResource(a, &AtpMaliciousUrls{}, `The zia_atp_malicious_urls resource manages the list of malicious URL exceptions for Advanced Threat Protection (ATP) in the Zscaler Internet Access (ZIA) cloud service. URLs added to this list are treated as known malicious and will be blocked. This is a singleton resource.

For more information, see the [ZIA Advanced Threat Protection documentation](https://help.zscaler.com/zia/about-advanced-threat-protection-policy).

{{% examples %}}
## Example Usage

{{% example %}}
### Configure ATP Malicious URL Exceptions

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.AtpMaliciousUrls("example", {
    maliciousUrls: [
        "malicious-site.com",
        "phishing-example.net",
        "bad-domain.org",
    ],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.AtpMaliciousUrls("example",
    malicious_urls=[
        "malicious-site.com",
        "phishing-example.net",
        "bad-domain.org",
    ],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:AtpMaliciousUrls
    properties:
      maliciousUrls:
        - malicious-site.com
        - phishing-example.net
        - bad-domain.org
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

> This is a singleton resource. Import is not applicable.
`)
}

func (a *AtpMaliciousUrlsArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.MaliciousUrls, "List of URLs to be treated as malicious by Advanced Threat Protection.")
}

func (s *AtpMaliciousUrlsState) Annotate(a infer.Annotator) {
	a.Describe(&s.ResourceId, "The internal resource identifier for the ATP malicious URLs.")
}

var _ infer.CustomResource[AtpMaliciousUrlsArgs, AtpMaliciousUrlsState] = AtpMaliciousUrls{}
