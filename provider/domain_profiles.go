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

// Package provider implements the getDomainProfile invoke (data source).
// Adopted from terraform-provider-zia data_source_zia_domain_profiles.go.

package provider

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/saas_security_api"
)

type GetDomainProfileArgs struct {
	ProfileId   *int    `pulumi:"profileId,optional"`
	ProfileName *string `pulumi:"profileName,optional"`
}

type GetDomainProfileResult struct {
	ProfileId              int      `pulumi:"profileId"`
	ProfileName            string   `pulumi:"profileName"`
	Description            string   `pulumi:"description"`
	IncludeCompanyDomains  bool     `pulumi:"includeCompanyDomains"`
	IncludeSubdomains      bool     `pulumi:"includeSubdomains"`
	CustomDomains          []string `pulumi:"customDomains"`
	PredefinedEmailDomains []string `pulumi:"predefinedEmailDomains"`
}

type GetDomainProfile struct{}

func (f *GetDomainProfile) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to look up a domain profile by ID or profile name.")
}

func (a *GetDomainProfileArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.ProfileId, "The ID of the domain profile to look up.")
	ann.Describe(&a.ProfileName, "The name of the domain profile to look up.")
}

func (r *GetDomainProfileResult) Annotate(a infer.Annotator) {
	a.Describe(&r.ProfileId, "The ID of the domain profile.")
	a.Describe(&r.ProfileName, "The name of the domain profile.")
	a.Describe(&r.Description, "The description of the domain profile.")
	a.Describe(&r.IncludeCompanyDomains, "Whether company domains are included.")
	a.Describe(&r.IncludeSubdomains, "Whether subdomains are included.")
	a.Describe(&r.CustomDomains, "The list of custom domains in the profile.")
	a.Describe(&r.PredefinedEmailDomains, "The list of predefined email domains in the profile.")
}

func (*GetDomainProfile) Invoke(ctx context.Context, req infer.FunctionRequest[GetDomainProfileArgs]) (infer.FunctionResponse[GetDomainProfileResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetDomainProfileResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	profiles, err := saas_security_api.GetDomainProfiles(ctx, svc)
	if err != nil {
		return infer.FunctionResponse[GetDomainProfileResult]{}, fmt.Errorf("failed to retrieve domain profiles: %w", err)
	}

	id := ptrToIntDefault(req.Input.ProfileId, 0)
	name := ptrToString(req.Input.ProfileName)
	var matched *saas_security_api.DomainProfiles
	for i := range profiles {
		if id != 0 && profiles[i].ProfileID == id {
			matched = &profiles[i]
			break
		}
		if name != "" && profiles[i].ProfileName == name {
			matched = &profiles[i]
			break
		}
	}
	if matched == nil {
		return infer.FunctionResponse[GetDomainProfileResult]{}, fmt.Errorf("couldn't find any domain profile with name '%s' or id '%d'", name, id)
	}

	return infer.FunctionResponse[GetDomainProfileResult]{Output: GetDomainProfileResult{
		ProfileId:              matched.ProfileID,
		ProfileName:            matched.ProfileName,
		Description:            matched.Description,
		IncludeCompanyDomains:  matched.IncludeCompanyDomains,
		IncludeSubdomains:      matched.IncludeSubdomains,
		CustomDomains:          matched.CustomDomains,
		PredefinedEmailDomains: matched.PredefinedEmailDomains,
	}}, nil
}
