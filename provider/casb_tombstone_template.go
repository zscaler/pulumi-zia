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

// Package provider implements the getCasbTombstoneTemplate invoke (data source).
// Adopted from terraform-provider-zia data_source_zia_casb_tombstone_template.go.

package provider

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/saas_security_api"
)

type GetCasbTombstoneTemplateArgs struct {
	Id   *int    `pulumi:"resourceId,optional"`
	Name *string `pulumi:"name,optional"`
}

type GetCasbTombstoneTemplateResult struct {
	Id          int    `pulumi:"resourceId"`
	Name        string `pulumi:"name"`
	Description string `pulumi:"description"`
}

type GetCasbTombstoneTemplate struct{}

func (f *GetCasbTombstoneTemplate) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to look up a CASB quarantine tombstone template by ID or name.")
}

func (a *GetCasbTombstoneTemplateArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Id, "The ID of the tombstone template to look up.")
	ann.Describe(&a.Name, "The name of the tombstone template to look up.")
}

func (r *GetCasbTombstoneTemplateResult) Annotate(a infer.Annotator) {
	a.Describe(&r.Id, "The ID of the tombstone template.")
	a.Describe(&r.Name, "The name of the tombstone template.")
	a.Describe(&r.Description, "The description of the tombstone template.")
}

func (*GetCasbTombstoneTemplate) Invoke(ctx context.Context, req infer.FunctionRequest[GetCasbTombstoneTemplateArgs]) (infer.FunctionResponse[GetCasbTombstoneTemplateResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetCasbTombstoneTemplateResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	templates, err := saas_security_api.GetQuarantineTombstoneLite(ctx, svc)
	if err != nil {
		return infer.FunctionResponse[GetCasbTombstoneTemplateResult]{}, err
	}

	id := ptrToIntDefault(req.Input.Id, 0)
	name := ptrToString(req.Input.Name)
	var matched *saas_security_api.QuarantineTombstoneLite
	for i := range templates {
		if id != 0 && templates[i].ID == id {
			matched = &templates[i]
			break
		}
		if name != "" && templates[i].Name == name {
			matched = &templates[i]
			break
		}
	}
	if matched == nil {
		return infer.FunctionResponse[GetCasbTombstoneTemplateResult]{}, fmt.Errorf("couldn't find any quarantine tombstone template with name '%s' or id '%d'", name, id)
	}

	return infer.FunctionResponse[GetCasbTombstoneTemplateResult]{Output: GetCasbTombstoneTemplateResult{
		Id: matched.ID, Name: matched.Name, Description: matched.Description,
	}}, nil
}
