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

// Package provider implements the getCasbEmailLabel invoke (data source).
// Adopted from terraform-provider-zia data_source_zia_casb_email_label.go.

package provider

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/saas_security_api"
)

type GetCasbEmailLabelArgs struct {
	Id   *int    `pulumi:"resourceId,optional"`
	Name *string `pulumi:"name,optional"`
}

type GetCasbEmailLabelResult struct {
	Id           int    `pulumi:"resourceId"`
	Name         string `pulumi:"name"`
	LabelDeleted bool   `pulumi:"labelDeleted"`
}

type GetCasbEmailLabel struct{}

func (f *GetCasbEmailLabel) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to look up a CASB email label by ID or name.")
}

func (a *GetCasbEmailLabelArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Id, "The ID of the CASB email label to look up.")
	ann.Describe(&a.Name, "The name of the CASB email label to look up.")
}

func (r *GetCasbEmailLabelResult) Annotate(a infer.Annotator) {
	a.Describe(&r.Id, "The ID of the CASB email label.")
	a.Describe(&r.Name, "The name of the CASB email label.")
	a.Describe(&r.LabelDeleted, "Whether the email label has been deleted.")
}

func (*GetCasbEmailLabel) Invoke(ctx context.Context, req infer.FunctionRequest[GetCasbEmailLabelArgs]) (infer.FunctionResponse[GetCasbEmailLabelResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetCasbEmailLabelResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	labels, err := saas_security_api.GetCasbEmailLabelLite(ctx, svc)
	if err != nil {
		return infer.FunctionResponse[GetCasbEmailLabelResult]{}, err
	}

	id := ptrToIntDefault(req.Input.Id, 0)
	name := ptrToString(req.Input.Name)
	var matched *saas_security_api.CasbEmailLabel
	for i := range labels {
		if id != 0 && labels[i].ID == id {
			matched = &labels[i]
			break
		}
		if name != "" && labels[i].Name == name {
			matched = &labels[i]
			break
		}
	}
	if matched == nil {
		return infer.FunctionResponse[GetCasbEmailLabelResult]{}, fmt.Errorf("couldn't find any email label with name '%s' or id '%d'", name, id)
	}

	return infer.FunctionResponse[GetCasbEmailLabelResult]{Output: GetCasbEmailLabelResult{
		Id: matched.ID, Name: matched.Name, LabelDeleted: matched.LabelDeleted,
	}}, nil
}
