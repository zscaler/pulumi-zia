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

// Package provider implements the FTP Control Policy resource.
// Adopted from terraform-provider-zia resource_zia_ftp_control_policy.go.
// Singleton resource (one per tenant). Delete is a no-op.

package provider

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/ftp_control_policy"
)

// FtpControlPolicy implements the zia:index:FtpControlPolicy resource.
type FtpControlPolicy struct{}

// FtpControlPolicyArgs are the inputs.
type FtpControlPolicyArgs struct {
	FtpOverHttpEnabled *bool    `pulumi:"ftpOverHttpEnabled,optional"`
	FtpEnabled         *bool    `pulumi:"ftpEnabled,optional"`
	Urls               []string `pulumi:"urls,optional"`
	UrlCategories      []string `pulumi:"urlCategories,optional"`
}

// FtpControlPolicyState is the persisted state.
type FtpControlPolicyState struct {
	FtpControlPolicyArgs
}

const ftpControlPolicyID = "ftp_control"

func ftpControlPolicyToAPI(args FtpControlPolicyArgs) ftp_control_policy.FTPControlPolicy {
	return ftp_control_policy.FTPControlPolicy{
		FtpOverHttpEnabled: ptrToBool(args.FtpOverHttpEnabled),
		FtpEnabled:         ptrToBool(args.FtpEnabled),
		Urls:               args.Urls,
		UrlCategories:      args.UrlCategories,
	}
}

func (FtpControlPolicy) Create(ctx context.Context, req infer.CreateRequest[FtpControlPolicyArgs]) (infer.CreateResponse[FtpControlPolicyState], error) {
	if req.DryRun {
		s := FtpControlPolicyState{FtpControlPolicyArgs: req.Inputs}
		return infer.CreateResponse[FtpControlPolicyState]{ID: "preview", Output: s}, nil
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[FtpControlPolicyState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	policy := ftpControlPolicyToAPI(req.Inputs)
	if _, _, err := ftp_control_policy.UpdateFTPControlPolicy(ctx, service, &policy); err != nil {
		return infer.CreateResponse[FtpControlPolicyState]{}, err
	}

	time.Sleep(1 * time.Second)
	if shouldActivate() {
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[FtpControlPolicyState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	return infer.CreateResponse[FtpControlPolicyState]{
		ID:     ftpControlPolicyID,
		Output: FtpControlPolicyState{FtpControlPolicyArgs: req.Inputs},
	}, nil
}

func (FtpControlPolicy) Read(ctx context.Context, req infer.ReadRequest[FtpControlPolicyArgs, FtpControlPolicyState]) (infer.ReadResponse[FtpControlPolicyArgs, FtpControlPolicyState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[FtpControlPolicyArgs, FtpControlPolicyState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	resp, err := ftp_control_policy.GetFTPControlPolicy(ctx, service)
	if err != nil {
		return infer.ReadResponse[FtpControlPolicyArgs, FtpControlPolicyState]{}, err
	}
	if resp == nil {
		return infer.ReadResponse[FtpControlPolicyArgs, FtpControlPolicyState]{}, fmt.Errorf("could not read FTP control policy")
	}

	urlCategories := resp.UrlCategories
	if len(urlCategories) == 1 && urlCategories[0] == "ANY" {
		urlCategories = nil
	}
	args := FtpControlPolicyArgs{
		FtpOverHttpEnabled: boolPtr(resp.FtpOverHttpEnabled),
		FtpEnabled:         boolPtr(resp.FtpEnabled),
		Urls:               resp.Urls,
		UrlCategories:      urlCategories,
	}
	state := FtpControlPolicyState{FtpControlPolicyArgs: args}
	return infer.ReadResponse[FtpControlPolicyArgs, FtpControlPolicyState]{
		ID:     ftpControlPolicyID,
		Inputs: args,
		State:  state,
	}, nil
}

func (FtpControlPolicy) Update(ctx context.Context, req infer.UpdateRequest[FtpControlPolicyArgs, FtpControlPolicyState]) (infer.UpdateResponse[FtpControlPolicyState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[FtpControlPolicyState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	policy := ftpControlPolicyToAPI(req.Inputs)
	if _, _, err := ftp_control_policy.UpdateFTPControlPolicy(ctx, service, &policy); err != nil {
		return infer.UpdateResponse[FtpControlPolicyState]{}, err
	}

	time.Sleep(1 * time.Second)
	if shouldActivate() {
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[FtpControlPolicyState]{}, activationErr
		}
	}

	return infer.UpdateResponse[FtpControlPolicyState]{
		Output: FtpControlPolicyState{FtpControlPolicyArgs: req.Inputs},
	}, nil
}

func (FtpControlPolicy) Delete(ctx context.Context, req infer.DeleteRequest[FtpControlPolicyState]) (infer.DeleteResponse, error) {
	// No-op: singleton policy; deleting the Pulumi resource does not remove the underlying settings
	return infer.DeleteResponse{}, nil
}

func (FtpControlPolicy) Diff(ctx context.Context, req infer.DiffRequest[FtpControlPolicyArgs, FtpControlPolicyState]) (infer.DiffResponse, error) {
	diff, hasChanges := diffSkippingUnsetInputs(req.State.FtpControlPolicyArgs, req.Inputs)
	return infer.DiffResponse{HasChanges: hasChanges, DetailedDiff: diff}, nil
}

func (FtpControlPolicy) Annotate(a infer.Annotator) {
	describeResource(a, &FtpControlPolicy{}, `The zia.FtpControlPolicy resource manages the FTP control policy settings in the Zscaler Internet Access (ZIA) cloud.
This is a singleton resource (one per tenant). The policy controls whether FTP and FTP-over-HTTP traffic is allowed
or blocked, and which URLs and URL categories are subject to FTP controls. Deleting the Pulumi resource does not
remove the underlying settings.

{{% examples %}}
## Example Usage

{{% example %}}
### Basic FTP Control Policy

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.FtpControlPolicy("example", {
    ftpEnabled: true,
    ftpOverHttpEnabled: false,
    urls: ["example.com"],
    urlCategories: ["OTHER_ADULT_MATERIAL"],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.FtpControlPolicy("example",
    ftp_enabled=True,
    ftp_over_http_enabled=False,
    urls=["example.com"],
    url_categories=["OTHER_ADULT_MATERIAL"],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:FtpControlPolicy
    properties:
      ftpEnabled: true
      ftpOverHttpEnabled: false
      urls:
        - example.com
      urlCategories:
        - OTHER_ADULT_MATERIAL
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

This is a singleton resource and import is not supported. The resource is managed by creating it in your Pulumi program.
`)
}

func (a *FtpControlPolicyArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.FtpOverHttpEnabled, "Whether FTP-over-HTTP traffic is enabled.")
	ann.Describe(&a.FtpEnabled, "Whether native FTP traffic is enabled.")
	ann.Describe(&a.Urls, "List of URLs subject to the FTP control policy.")
	ann.Describe(&a.UrlCategories, "List of URL categories subject to the FTP control policy.")
}

var _ infer.CustomResource[FtpControlPolicyArgs, FtpControlPolicyState] = FtpControlPolicy{}
