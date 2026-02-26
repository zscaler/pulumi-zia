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

// Package provider implements the Sandbox Submission resource.
// Adopted from terraform-provider-zia resource_zia_sandbox_submission.go.
// Submits files for sandbox analysis. Read has no API; Delete is a no-op.

package provider

import (
	"context"
	"fmt"
	"os"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/sandbox/sandbox_submission"
)

// SandboxSubmission implements the zia:index:SandboxSubmission resource.
type SandboxSubmission struct{}

// SandboxSubmissionArgs are the inputs.
type SandboxSubmissionArgs struct {
	FilePath         string `pulumi:"filePath"`
	Force            *bool  `pulumi:"force,optional"`
	SubmissionMethod string `pulumi:"submissionMethod"` // "submit" or "discan"
}

// SandboxSubmissionState is the persisted state.
type SandboxSubmissionState struct {
	SandboxSubmissionArgs
	Code              *int    `pulumi:"code,optional"`
	Message           *string `pulumi:"message,optional"`
	FileType          *string `pulumi:"fileType,optional"`
	Md5               *string `pulumi:"md5,optional"`
	SandboxSubmission *string `pulumi:"submissionResult,optional"` // Renamed from sandboxSubmission to avoid C# class name clash
	VirusName         *string `pulumi:"virusName,optional"`
	VirusType         *string `pulumi:"virusType,optional"`
}

func (SandboxSubmission) Create(ctx context.Context, req infer.CreateRequest[SandboxSubmissionArgs]) (infer.CreateResponse[SandboxSubmissionState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[SandboxSubmissionState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	force := ptrToBool(req.Inputs.Force)
	if req.Inputs.SubmissionMethod == "discan" && force {
		return infer.CreateResponse[SandboxSubmissionState]{}, fmt.Errorf("'force' attribute is not applicable for 'discan' submission method")
	}

	file, err := os.Open(req.Inputs.FilePath)
	if err != nil {
		return infer.CreateResponse[SandboxSubmissionState]{}, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var result *sandbox_submission.ScanResult
	if req.Inputs.SubmissionMethod == "submit" {
		forceStr := "0"
		if force {
			forceStr = "1"
		}
		result, err = sandbox_submission.SubmitFile(ctx, service, req.Inputs.FilePath, file, forceStr)
	} else if req.Inputs.SubmissionMethod == "discan" {
		result, err = sandbox_submission.Discan(ctx, service, req.Inputs.FilePath, file)
	} else {
		return infer.CreateResponse[SandboxSubmissionState]{}, fmt.Errorf("invalid submission method: %s", req.Inputs.SubmissionMethod)
	}

	if err != nil {
		return infer.CreateResponse[SandboxSubmissionState]{}, fmt.Errorf("error submitting file to Sandbox: %w", err)
	}

	state := SandboxSubmissionState{
		SandboxSubmissionArgs: req.Inputs,
		Code:                  intPtr(result.Code),
		Message:               stringPtr(result.Message),
		FileType:              stringPtr(result.FileType),
		Md5:                   stringPtr(result.Md5),
		SandboxSubmission:     stringPtr(result.SandboxSubmission),
		VirusName:             stringPtr(result.VirusName),
		VirusType:             stringPtr(result.VirusType),
	}
	return infer.CreateResponse[SandboxSubmissionState]{
		ID:     result.Md5,
		Output: state,
	}, nil
}

func (SandboxSubmission) Read(ctx context.Context, req infer.ReadRequest[SandboxSubmissionArgs, SandboxSubmissionState]) (infer.ReadResponse[SandboxSubmissionArgs, SandboxSubmissionState], error) {
	// No GET API; return state as-is
	return infer.ReadResponse[SandboxSubmissionArgs, SandboxSubmissionState]{
		ID:     req.ID,
		Inputs: req.State.SandboxSubmissionArgs,
		State:  req.State,
	}, nil
}

func (SandboxSubmission) Update(ctx context.Context, req infer.UpdateRequest[SandboxSubmissionArgs, SandboxSubmissionState]) (infer.UpdateResponse[SandboxSubmissionState], error) {
	// Re-create on file_path or force change
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[SandboxSubmissionState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	force := ptrToBool(req.Inputs.Force)
	if req.Inputs.SubmissionMethod == "discan" && force {
		return infer.UpdateResponse[SandboxSubmissionState]{}, fmt.Errorf("'force' attribute is not applicable for 'discan' submission method")
	}

	file, err := os.Open(req.Inputs.FilePath)
	if err != nil {
		return infer.UpdateResponse[SandboxSubmissionState]{}, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var result *sandbox_submission.ScanResult
	if req.Inputs.SubmissionMethod == "submit" {
		forceStr := "0"
		if force {
			forceStr = "1"
		}
		result, err = sandbox_submission.SubmitFile(ctx, service, req.Inputs.FilePath, file, forceStr)
	} else if req.Inputs.SubmissionMethod == "discan" {
		result, err = sandbox_submission.Discan(ctx, service, req.Inputs.FilePath, file)
	} else {
		return infer.UpdateResponse[SandboxSubmissionState]{}, fmt.Errorf("invalid submission method: %s", req.Inputs.SubmissionMethod)
	}

	if err != nil {
		return infer.UpdateResponse[SandboxSubmissionState]{}, fmt.Errorf("error submitting file to Sandbox: %w", err)
	}

	state := SandboxSubmissionState{
		SandboxSubmissionArgs: req.Inputs,
		Code:                  intPtr(result.Code),
		Message:               stringPtr(result.Message),
		FileType:              stringPtr(result.FileType),
		Md5:                   stringPtr(result.Md5),
		SandboxSubmission:     stringPtr(result.SandboxSubmission),
		VirusName:             stringPtr(result.VirusName),
		VirusType:             stringPtr(result.VirusType),
	}
	return infer.UpdateResponse[SandboxSubmissionState]{Output: state}, nil
}

func (SandboxSubmission) Delete(ctx context.Context, req infer.DeleteRequest[SandboxSubmissionState]) (infer.DeleteResponse, error) {
	// No DELETE API
	return infer.DeleteResponse{}, nil
}

func (SandboxSubmission) Annotate(a infer.Annotator) {
	describeResource(a, &SandboxSubmission{}, `The zia_sandbox_submission resource submits files to the Zscaler cloud sandbox for analysis. Files can be submitted for full analysis or a quick discan (distributed scan). This resource is create-only; there is no remote GET API, and delete is a no-op.

For more information, see the [ZIA Cloud Sandbox Submission documentation](https://help.zscaler.com/zia/about-sandbox-analysis).

{{% examples %}}
## Example Usage

{{% example %}}
### Submit a File for Sandbox Analysis

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.SandboxSubmission("example", {
    filePath: "/tmp/suspicious-file.exe",
    submissionMethod: "submit",
    force: true,
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.SandboxSubmission("example",
    file_path="/tmp/suspicious-file.exe",
    submission_method="submit",
    force=True,
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:SandboxSubmission
    properties:
      filePath: /tmp/suspicious-file.exe
      submissionMethod: submit
      force: true
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

> Import is not supported for this resource.
`)
}

func (a *SandboxSubmissionArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.FilePath, "The local file path of the file to submit for sandbox analysis.")
	ann.Describe(&a.Force, "Force re-analysis of a previously submitted file. Only applicable for 'submit' method. Not applicable for 'discan'.")
	ann.Describe(&a.SubmissionMethod, "The submission method. Valid values: `submit` (full analysis) or `discan` (distributed scan).")
}

func (s *SandboxSubmissionState) Annotate(a infer.Annotator) {
	a.Describe(&s.Code, "The response status code from the sandbox submission.")
	a.Describe(&s.Message, "The response message from the sandbox submission.")
	a.Describe(&s.FileType, "The detected file type of the submitted file.")
	a.Describe(&s.Md5, "The MD5 hash of the submitted file.")
	a.Describe(&s.SandboxSubmission, "The sandbox submission result string.")
	a.Describe(&s.VirusName, "The virus name if the file is detected as malicious.")
	a.Describe(&s.VirusType, "The virus type if the file is detected as malicious.")
}

func (SandboxSubmission) Check(ctx context.Context, req infer.CheckRequest) (infer.CheckResponse[SandboxSubmissionArgs], error) {
	inputs, failures, err := infer.DefaultCheck[SandboxSubmissionArgs](ctx, req.NewInputs)
	if err != nil {
		return infer.CheckResponse[SandboxSubmissionArgs]{}, err
	}
	if inputs.SubmissionMethod != "submit" && inputs.SubmissionMethod != "discan" {
		failures = append(failures, p.CheckFailure{Property: "submissionMethod", Reason: "must be 'submit' or 'discan'"})
	}
	if len(failures) > 0 {
		return infer.CheckResponse[SandboxSubmissionArgs]{Failures: failures}, nil
	}
	return infer.CheckResponse[SandboxSubmissionArgs]{Inputs: inputs}, nil
}

var _ infer.CustomResource[SandboxSubmissionArgs, SandboxSubmissionState] = SandboxSubmission{}
