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

// Package provider implements the getDlpIdmProfile invoke (data source).
// Adopted from terraform-provider-zia data_source_zia_dlp_idm_profiles.go.

package provider

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/dlp/dlp_idm_profiles"
)

type GetDlpIdmProfileArgs struct {
	ProfileId   *int    `pulumi:"profileId,optional"`
	ProfileName *string `pulumi:"profileName,optional"`
}

type GetDlpIdmProfileResult struct {
	ProfileId          int      `pulumi:"profileId"`
	ProfileName        string   `pulumi:"profileName"`
	ProfileDesc        string   `pulumi:"profileDesc"`
	ProfileType        string   `pulumi:"profileType"`
	Host               string   `pulumi:"host"`
	Port               int      `pulumi:"port"`
	ProfileDirPath     string   `pulumi:"profileDirPath"`
	ScheduleType       string   `pulumi:"scheduleType"`
	ScheduleDay        int      `pulumi:"scheduleDay"`
	ScheduleTime       int      `pulumi:"scheduleTime"`
	ScheduleDisabled   bool     `pulumi:"scheduleDisabled"`
	UploadStatus       string   `pulumi:"uploadStatus"`
	UserName           string   `pulumi:"userName"`
	Version            int      `pulumi:"version"`
	VolumeOfDocuments  int      `pulumi:"volumeOfDocuments"`
	NumDocuments       int      `pulumi:"numDocuments"`
	LastModifiedTime   int      `pulumi:"lastModifiedTime"`
	ScheduleDayOfMonth []string `pulumi:"scheduleDayOfMonth"`
	ScheduleDayOfWeek  []string `pulumi:"scheduleDayOfWeek"`
}

type GetDlpIdmProfile struct{}

func (f *GetDlpIdmProfile) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to look up a DLP IDM profile by ID or name.")
}

func (a *GetDlpIdmProfileArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.ProfileId, "The ID of the DLP IDM profile to look up.")
	ann.Describe(&a.ProfileName, "The name of the DLP IDM profile to look up.")
}

func (r *GetDlpIdmProfileResult) Annotate(a infer.Annotator) {
	a.Describe(&r.ProfileId, "The ID of the DLP IDM profile.")
	a.Describe(&r.ProfileName, "The name of the DLP IDM profile.")
	a.Describe(&r.ProfileDesc, "The description of the DLP IDM profile.")
	a.Describe(&r.ProfileType, "The type of the DLP IDM profile.")
	a.Describe(&r.Host, "The host for the IDM profile.")
	a.Describe(&r.Port, "The port for the IDM profile.")
	a.Describe(&r.ProfileDirPath, "The directory path of the profile.")
	a.Describe(&r.ScheduleType, "The schedule type for the profile.")
	a.Describe(&r.ScheduleDay, "The schedule day for the profile.")
	a.Describe(&r.ScheduleTime, "The schedule time for the profile.")
	a.Describe(&r.ScheduleDisabled, "Whether the schedule is disabled.")
	a.Describe(&r.UploadStatus, "The upload status of the profile.")
	a.Describe(&r.UserName, "The username associated with the profile.")
	a.Describe(&r.Version, "The version of the profile.")
	a.Describe(&r.VolumeOfDocuments, "The volume of documents in the profile.")
	a.Describe(&r.NumDocuments, "The number of documents in the profile.")
	a.Describe(&r.LastModifiedTime, "The last modification time (epoch).")
	a.Describe(&r.ScheduleDayOfMonth, "The schedule days of the month.")
	a.Describe(&r.ScheduleDayOfWeek, "The schedule days of the week.")
}

func (*GetDlpIdmProfile) Invoke(ctx context.Context, req infer.FunctionRequest[GetDlpIdmProfileArgs]) (infer.FunctionResponse[GetDlpIdmProfileResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetDlpIdmProfileResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	var resp *dlp_idm_profiles.DLPIDMProfile
	if req.Input.ProfileId != nil && *req.Input.ProfileId != 0 {
		r, err := dlp_idm_profiles.Get(ctx, svc, *req.Input.ProfileId)
		if err != nil {
			return infer.FunctionResponse[GetDlpIdmProfileResult]{}, err
		}
		resp = r
	}
	if resp == nil && req.Input.ProfileName != nil && *req.Input.ProfileName != "" {
		r, err := dlp_idm_profiles.GetByName(ctx, svc, *req.Input.ProfileName)
		if err != nil {
			return infer.FunctionResponse[GetDlpIdmProfileResult]{}, err
		}
		resp = r
	}
	if resp == nil {
		return infer.FunctionResponse[GetDlpIdmProfileResult]{}, fmt.Errorf("couldn't find any DLP IDM profile with id %v or name %s", req.Input.ProfileId, ptrToString(req.Input.ProfileName))
	}

	return infer.FunctionResponse[GetDlpIdmProfileResult]{Output: GetDlpIdmProfileResult{
		ProfileId:          resp.ProfileID,
		ProfileName:        resp.ProfileName,
		ProfileDesc:        resp.ProfileDesc,
		ProfileType:        resp.ProfileType,
		Host:               resp.Host,
		Port:               resp.Port,
		ProfileDirPath:     resp.ProfileDirPath,
		ScheduleType:       resp.ScheduleType,
		ScheduleDay:        resp.ScheduleDay,
		ScheduleTime:       resp.ScheduleTime,
		ScheduleDisabled:   resp.ScheduleDisabled,
		UploadStatus:       resp.UploadStatus,
		UserName:           resp.UserName,
		Version:            resp.Version,
		VolumeOfDocuments:  resp.VolumeOfDocuments,
		NumDocuments:       resp.NumDocuments,
		LastModifiedTime:   resp.LastModifiedTime,
		ScheduleDayOfMonth: resp.ScheduleDayOfMonth,
		ScheduleDayOfWeek:  resp.ScheduleDayOfWeek,
	}}, nil
}
