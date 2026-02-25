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

// Package provider implements the getFileTypeCategories invoke (data source).
// Adopted from terraform-provider-zia data_source_zia_file_type_categories.go.

package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/filetypecontrol"
)

type GetFileTypeCategoriesArgs struct {
	Id                    *int    `pulumi:"resourceId,optional"`
	Name                  *string `pulumi:"name,optional"`
	Enums                 *string `pulumi:"enums,optional"`
	ExcludeCustomFileTypes *bool   `pulumi:"excludeCustomFileTypes,optional"`
}

type FileTypeCategoryItem struct {
	Id     int    `pulumi:"resourceId"`
	Name   string `pulumi:"name"`
	Parent string `pulumi:"parent"`
}

type GetFileTypeCategoriesResult struct {
	Id         *int                  `pulumi:"resourceId,optional"`
	Name       *string               `pulumi:"name,optional"`
	Parent     *string               `pulumi:"parent,optional"`
	Categories []FileTypeCategoryItem `pulumi:"categories"`
}

type GetFileTypeCategories struct{}

func (f *GetFileTypeCategories) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to look up file type categories, optionally filtered by ID, name, or enum type.")
}

func (a *GetFileTypeCategoriesArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Id, "The ID of the file type category to look up.")
	ann.Describe(&a.Name, "The name of the file type category to look up.")
	ann.Describe(&a.Enums, "The enum type to filter file type categories.")
	ann.Describe(&a.ExcludeCustomFileTypes, "If true, exclude custom file types from the results.")
}

func (r *GetFileTypeCategoriesResult) Annotate(a infer.Annotator) {
	a.Describe(&r.Id, "The ID of the matched file type category.")
	a.Describe(&r.Name, "The name of the matched file type category.")
	a.Describe(&r.Parent, "The parent category of the matched file type category.")
	a.Describe(&r.Categories, "The list of file type categories returned.")
}

func (*GetFileTypeCategories) Invoke(ctx context.Context, req infer.FunctionRequest[GetFileTypeCategoriesArgs]) (infer.FunctionResponse[GetFileTypeCategoriesResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetFileTypeCategoriesResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	var opts *filetypecontrol.GetFileTypeCategoriesFilterOptions
	if req.Input.Enums != nil && *req.Input.Enums != "" {
		if opts == nil {
			opts = &filetypecontrol.GetFileTypeCategoriesFilterOptions{}
		}
		opts.Enums = []string{*req.Input.Enums}
	}
	if req.Input.ExcludeCustomFileTypes != nil {
		if opts == nil {
			opts = &filetypecontrol.GetFileTypeCategoriesFilterOptions{}
		}
		opts.ExcludeCustomFileTypes = req.Input.ExcludeCustomFileTypes
	}

	cats, err := filetypecontrol.GetFileTypeCategories(ctx, svc, opts)
	if err != nil {
		return infer.FunctionResponse[GetFileTypeCategoriesResult]{}, err
	}
	if len(cats) == 0 {
		return infer.FunctionResponse[GetFileTypeCategoriesResult]{}, fmt.Errorf("no file type categories found")
	}

	if (req.Input.Id != nil && *req.Input.Id != 0) || (req.Input.Name != nil && *req.Input.Name != "") {
		var found *filetypecontrol.FileTypeCategory
		for i := range cats {
			if req.Input.Id != nil && cats[i].ID == *req.Input.Id {
				found = &cats[i]
				break
			}
			if req.Input.Name != nil && strings.EqualFold(cats[i].Name, *req.Input.Name) {
				found = &cats[i]
				break
			}
		}
		if found == nil {
			return infer.FunctionResponse[GetFileTypeCategoriesResult]{}, fmt.Errorf("couldn't find file type category with id %v or name %s", req.Input.Id, ptrToString(req.Input.Name))
		}
		return infer.FunctionResponse[GetFileTypeCategoriesResult]{Output: GetFileTypeCategoriesResult{
			Id:         intPtr(found.ID),
			Name:       stringPtr(found.Name),
			Parent:     stringPtr(found.Parent),
			Categories: []FileTypeCategoryItem{},
		}}, nil
	}

	list := make([]FileTypeCategoryItem, len(cats))
	for i, c := range cats {
		list[i] = FileTypeCategoryItem{Id: c.ID, Name: c.Name, Parent: c.Parent}
	}
	return infer.FunctionResponse[GetFileTypeCategoriesResult]{Output: GetFileTypeCategoriesResult{
		Categories: list,
	}}, nil
}
