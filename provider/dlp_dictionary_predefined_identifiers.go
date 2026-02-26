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

// Package provider implements the getDlpDictionaryPredefinedIdentifiers invoke (data source).
// Adopted from terraform-provider-zia data_source_zia_dlp_dictionary_predefined_identifiers.go.

package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/dlp/dlpdictionaries"
)

type GetDlpDictionaryPredefinedIdentifiersArgs struct {
	Name string `pulumi:"name"`
}

type GetDlpDictionaryPredefinedIdentifiersResult struct {
	Id                    string   `pulumi:"resourceId"`
	PredefinedIdentifiers []string `pulumi:"predefinedIdentifiers"`
}

type GetDlpDictionaryPredefinedIdentifiers struct{}

func (f *GetDlpDictionaryPredefinedIdentifiers) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to look up predefined DLP dictionary identifiers by dictionary name.")
}

func (a *GetDlpDictionaryPredefinedIdentifiersArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the predefined DLP dictionary to look up.")
}

func (r *GetDlpDictionaryPredefinedIdentifiersResult) Annotate(a infer.Annotator) {
	a.Describe(&r.Id, "The ID of the predefined DLP dictionary.")
	a.Describe(&r.PredefinedIdentifiers, "The list of predefined identifiers for the dictionary.")
}

func (*GetDlpDictionaryPredefinedIdentifiers) Invoke(ctx context.Context, req infer.FunctionRequest[GetDlpDictionaryPredefinedIdentifiersArgs]) (infer.FunctionResponse[GetDlpDictionaryPredefinedIdentifiersResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetDlpDictionaryPredefinedIdentifiersResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	if req.Input.Name == "" {
		return infer.FunctionResponse[GetDlpDictionaryPredefinedIdentifiersResult]{}, fmt.Errorf("name is required")
	}

	identifiers, dictionaryID, err := dlpdictionaries.GetPredefinedIdentifiers(ctx, svc, req.Input.Name)
	if err != nil {
		return infer.FunctionResponse[GetDlpDictionaryPredefinedIdentifiersResult]{}, err
	}

	return infer.FunctionResponse[GetDlpDictionaryPredefinedIdentifiersResult]{Output: GetDlpDictionaryPredefinedIdentifiersResult{
		Id:                    strconv.Itoa(dictionaryID),
		PredefinedIdentifiers: identifiers,
	}}, nil
}
