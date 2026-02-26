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

// Package provider implements the getSandboxReport invoke (data source).
// Adopted from terraform-provider-zia data_source_zia_sandbox_report.go.

package provider

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/common"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/sandbox/sandbox_report"
)

type GetSandboxReportArgs struct {
	Md5Hash string  `pulumi:"md5Hash"`
	Details *string `pulumi:"details,optional"`
}

type SummaryDetailOutput struct {
	Status    string `pulumi:"status"`
	Category  string `pulumi:"category"`
	FileType  string `pulumi:"fileType"`
	StartTime int    `pulumi:"startTime"`
	Duration  int    `pulumi:"duration"`
}

type ClassificationOutput struct {
	Type            string `pulumi:"type"`
	Category        string `pulumi:"category"`
	Score           int    `pulumi:"score"`
	DetectedMalware string `pulumi:"detectedMalware"`
}

type FilePropertiesOutput struct {
	FileType           string `pulumi:"fileType"`
	FileSize           int    `pulumi:"fileSize"`
	Md5                string `pulumi:"md5"`
	Sha1               string `pulumi:"sha1"`
	Sha256             string `pulumi:"sha256"`
	Issuer             string `pulumi:"issuer"`
	DigitalCertificate string `pulumi:"digitalCertificate"`
	Ssdeep             string `pulumi:"ssdeep"`
	RootCa             string `pulumi:"rootCa"`
}

type OriginOutput struct {
	Risk     string `pulumi:"risk"`
	Language string `pulumi:"language"`
	Country  string `pulumi:"country"`
}

type SandboxRssOutput struct {
	Risk             string   `pulumi:"risk"`
	Signature        string   `pulumi:"signature"`
	SignatureSources []string `pulumi:"signatureSources"`
}

type GetSandboxReportResult struct {
	Md5Hash        string                `pulumi:"md5Hash"`
	Summary        *SummaryDetailOutput  `pulumi:"summary,optional"`
	Classification *ClassificationOutput `pulumi:"classification,optional"`
	FileProperties *FilePropertiesOutput `pulumi:"fileProperties,optional"`
	Origin         *OriginOutput         `pulumi:"origin,optional"`
	SystemSummary  []SandboxRssOutput    `pulumi:"systemSummary"`
	Spyware        []SandboxRssOutput    `pulumi:"spyware"`
	Networking     []SandboxRssOutput    `pulumi:"networking"`
	SecurityBypass []SandboxRssOutput    `pulumi:"securityBypass"`
	Exploit        []SandboxRssOutput    `pulumi:"exploit"`
	Stealth        []SandboxRssOutput    `pulumi:"stealth"`
	Persistence    []SandboxRssOutput    `pulumi:"persistence"`
}

type GetSandboxReport struct{}

func (f *GetSandboxReport) Annotate(a infer.Annotator) {
	a.Describe(f, "Use this data source to retrieve a sandbox report for a given MD5 hash.")
}

func (a *GetSandboxReportArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Md5Hash, "The MD5 hash of the file to retrieve the sandbox report for.")
	ann.Describe(&a.Details, "The level of detail for the report. Accepted values: 'summary' or 'full'. Defaults to 'summary'.")
}

func (r *GetSandboxReportResult) Annotate(a infer.Annotator) {
	a.Describe(&r.Md5Hash, "The MD5 hash of the file.")
	a.Describe(&r.Summary, "Summary details of the sandbox analysis.")
	a.Describe(&r.Classification, "Classification details of the analyzed file.")
	a.Describe(&r.FileProperties, "File properties from the sandbox analysis.")
	a.Describe(&r.Origin, "Origin information of the analyzed file.")
	a.Describe(&r.SystemSummary, "System summary RSS entries from the full report.")
	a.Describe(&r.Spyware, "Spyware RSS entries from the full report.")
	a.Describe(&r.Networking, "Networking RSS entries from the full report.")
	a.Describe(&r.SecurityBypass, "Security bypass RSS entries from the full report.")
	a.Describe(&r.Exploit, "Exploit RSS entries from the full report.")
	a.Describe(&r.Stealth, "Stealth RSS entries from the full report.")
	a.Describe(&r.Persistence, "Persistence RSS entries from the full report.")
}

func flattenSandboxRSS(items []*common.SandboxRSS) []SandboxRssOutput {
	if items == nil {
		return nil
	}
	out := make([]SandboxRssOutput, len(items))
	for i, item := range items {
		sigSources := []string{}
		if item.SignatureSources != "" {
			sigSources = []string{item.SignatureSources}
		}
		out[i] = SandboxRssOutput{
			Risk:             item.Risk,
			Signature:        item.Signature,
			SignatureSources: sigSources,
		}
	}
	return out
}

func (*GetSandboxReport) Invoke(ctx context.Context, req infer.FunctionRequest[GetSandboxReportArgs]) (infer.FunctionResponse[GetSandboxReportResult], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.FunctionResponse[GetSandboxReportResult]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	if req.Input.Md5Hash == "" {
		return infer.FunctionResponse[GetSandboxReportResult]{}, fmt.Errorf("md5Hash is required")
	}
	details := "summary"
	if req.Input.Details != nil && *req.Input.Details != "" {
		details = *req.Input.Details
		if details != "full" && details != "summary" {
			return infer.FunctionResponse[GetSandboxReportResult]{}, fmt.Errorf("details must be 'full' or 'summary'")
		}
	}

	resp, err := sandbox_report.GetReportMD5Hash(ctx, svc, req.Input.Md5Hash, details)
	if err != nil {
		return infer.FunctionResponse[GetSandboxReportResult]{}, err
	}
	if resp == nil || resp.Details == nil {
		return infer.FunctionResponse[GetSandboxReportResult]{}, fmt.Errorf("couldn't find any reports for MD5 hash: %s", req.Input.Md5Hash)
	}

	d := resp.Details
	result := GetSandboxReportResult{
		Md5Hash: req.Input.Md5Hash,
		Summary: &SummaryDetailOutput{
			Status: d.Summary.Status, Category: d.Summary.Category, FileType: d.Summary.FileType,
			StartTime: d.Summary.StartTime, Duration: d.Summary.Duration,
		},
		Classification: &ClassificationOutput{
			Type: d.Classification.Type, Category: d.Classification.Category,
			Score: d.Classification.Score, DetectedMalware: d.Classification.DetectedMalware,
		},
		FileProperties: &FilePropertiesOutput{
			FileType: d.FileProperties.FileType, FileSize: d.FileProperties.FileSize,
			Md5: d.FileProperties.MD5, Sha1: d.FileProperties.SHA1, Sha256: d.FileProperties.SHA256,
			Issuer: d.FileProperties.Issuer, DigitalCertificate: d.FileProperties.DigitalCerificate,
			Ssdeep: d.FileProperties.SSDeep, RootCa: d.FileProperties.RootCA,
		},
	}

	if details == "full" {
		if d.Origin != nil {
			result.Origin = &OriginOutput{
				Risk: d.Origin.Risk, Language: d.Origin.Language, Country: d.Origin.Country,
			}
		}
		if len(d.SystemSummary) > 0 {
			out := make([]SandboxRssOutput, len(d.SystemSummary))
			for i, ss := range d.SystemSummary {
				out[i] = SandboxRssOutput{Risk: ss.Risk, Signature: ss.Signature, SignatureSources: ss.SignatureSources}
			}
			result.SystemSummary = out
		}
		result.Spyware = flattenSandboxRSS(d.Spyware)
		result.Networking = flattenSandboxRSS(d.Networking)
		result.SecurityBypass = flattenSandboxRSS(d.SecurityBypass)
		result.Exploit = flattenSandboxRSS(d.Exploit)
		result.Stealth = flattenSandboxRSS(d.Stealth)
		result.Persistence = flattenSandboxRSS(d.Persistence)
	}

	return infer.FunctionResponse[GetSandboxReportResult]{Output: result}, nil
}
