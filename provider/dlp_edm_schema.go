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

// Package provider implements the getDlpEdmSchema invoke (data source).
// Adopted from terraform-provider-zia data_source_zia_dlp_edm_schema.go.

package provider

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/dlp/dlp_exact_data_match"
)

type GetDlpEdmSchemaArgs struct {
	SchemaId    *int    `pulumi:"schemaId,optional"`
	ProjectName *string `pulumi:"projectName,optional"`
}

type GetDlpEdmSchemaResult struct {
	SchemaId         int    `pulumi:"schemaId"`
	ProjectName      string `pulumi:"projectName"`
	Revision         int    `pulumi:"revision"`
	FileName         string `pulumi:"fileName"`
	OriginalFileName string `pulumi:"originalFileName"`
	FileUploadStatus string `pulumi:"fileUploadStatus"`
	OrigColCount     int    `pulumi:"origColCount"`
	LastModifiedTime int    `pulumi:"lastModifiedTime"`
	CellsUsed        int    `pulumi:"cellsUsed"`
	SchemaActive     bool   `pulumi:"schemaActive"`
	SchedulePresent  bool   `pulumi:"schedulePresent"`
}

type GetDlpEdmSchema struct{}

func (f *GetDlpEdmSchema) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to look up a DLP Exact Data Match (EDM) schema by ID or project name.")
}

func (a *GetDlpEdmSchemaArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.SchemaId, "The ID of the DLP EDM schema to look up.")
	ann.Describe(&a.ProjectName, "The project name of the DLP EDM schema to look up.")
}

func (r *GetDlpEdmSchemaResult) Annotate(a infer.Annotator) {
	a.Describe(&r.SchemaId, "The ID of the DLP EDM schema.")
	a.Describe(&r.ProjectName, "The project name of the DLP EDM schema.")
	a.Describe(&r.Revision, "The revision number of the schema.")
	a.Describe(&r.FileName, "The file name of the schema.")
	a.Describe(&r.OriginalFileName, "The original file name of the schema.")
	a.Describe(&r.FileUploadStatus, "The file upload status.")
	a.Describe(&r.OrigColCount, "The original column count.")
	a.Describe(&r.LastModifiedTime, "The last modification time (epoch).")
	a.Describe(&r.CellsUsed, "The number of cells used.")
	a.Describe(&r.SchemaActive, "Whether the schema is active.")
	a.Describe(&r.SchedulePresent, "Whether a schedule is present.")
}

func (*GetDlpEdmSchema) Invoke(ctx context.Context, req infer.FunctionRequest[GetDlpEdmSchemaArgs]) (infer.FunctionResponse[GetDlpEdmSchemaResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetDlpEdmSchemaResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	var resp *dlp_exact_data_match.DLPEDMSchema
	if req.Input.SchemaId != nil && *req.Input.SchemaId != 0 {
		r, err := dlp_exact_data_match.GetDLPEDMSchemaID(ctx, svc, *req.Input.SchemaId)
		if err != nil {
			return infer.FunctionResponse[GetDlpEdmSchemaResult]{}, err
		}
		resp = r
	}
	if resp == nil && req.Input.ProjectName != nil && *req.Input.ProjectName != "" {
		r, err := dlp_exact_data_match.GetDLPEDMByName(ctx, svc, *req.Input.ProjectName)
		if err != nil {
			return infer.FunctionResponse[GetDlpEdmSchemaResult]{}, err
		}
		resp = r
	}
	if resp == nil {
		return infer.FunctionResponse[GetDlpEdmSchemaResult]{}, fmt.Errorf("couldn't find any DLP EDM schema with id %v or name %s", req.Input.SchemaId, ptrToString(req.Input.ProjectName))
	}

	return infer.FunctionResponse[GetDlpEdmSchemaResult]{Output: GetDlpEdmSchemaResult{
		SchemaId:         resp.SchemaID,
		ProjectName:      resp.ProjectName,
		Revision:         resp.Revision,
		FileName:         resp.Filename,
		OriginalFileName: resp.OriginalFileName,
		FileUploadStatus: resp.FileUploadStatus,
		OrigColCount:     resp.OrigColCount,
		LastModifiedTime: resp.LastModifiedTime,
		CellsUsed:        resp.CellsUsed,
		SchemaActive:     resp.SchemaActive,
		SchedulePresent:  resp.SchedulePresent,
	}}, nil
}
