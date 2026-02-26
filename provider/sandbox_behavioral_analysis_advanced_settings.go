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

// Package provider implements the Sandbox Behavioral Analysis Advanced Settings resource.
// Adopted from terraform-provider-zia resource_zia_sandbox_behavioral_analysis_advanced_settings.go.
// Singleton: file hashes to block. Delete is a no-op.

package provider

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/sandbox/sandbox_settings"
)

const sandboxSettingsID = "sandbox_settings"

// SandboxBehavioralAnalysisAdvancedSettings implements the zia:index:SandboxBehavioralAnalysisAdvancedSettings resource.
type SandboxBehavioralAnalysisAdvancedSettings struct{}

// SandboxBehavioralAnalysisAdvancedSettingsArgs are the inputs.
type SandboxBehavioralAnalysisAdvancedSettingsArgs struct {
	FileHashesToBeBlocked []string `pulumi:"fileHashesToBeBlocked,optional"`
}

// SandboxBehavioralAnalysisAdvancedSettingsState is the persisted state.
type SandboxBehavioralAnalysisAdvancedSettingsState struct {
	SandboxBehavioralAnalysisAdvancedSettingsArgs
	Id string `pulumi:"resourceId"` // Pulumi reserves "id" for resource identifier
}

func validateSandboxHashes(hashes []string) error {
	for _, hash := range hashes {
		switch len(hash) {
		case 32:
			// MD5
		case 40:
			return fmt.Errorf("hash '%s' is SHA1; sandbox only supports MD5", hash)
		case 64:
			return fmt.Errorf("hash '%s' is SHA256; sandbox only supports MD5", hash)
		default:
			return fmt.Errorf("hash '%s' has invalid length; MD5 must be 32 chars", hash)
		}
	}
	return nil
}

func sortStringSlice(s []string) []string {
	if len(s) == 0 {
		return nil
	}
	out := make([]string, len(s))
	copy(out, s)
	sort.Strings(out)
	return out
}

func (SandboxBehavioralAnalysisAdvancedSettings) Create(ctx context.Context, req infer.CreateRequest[SandboxBehavioralAnalysisAdvancedSettingsArgs]) (infer.CreateResponse[SandboxBehavioralAnalysisAdvancedSettingsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[SandboxBehavioralAnalysisAdvancedSettingsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	service := client.Service

	if err := validateSandboxHashes(req.Inputs.FileHashesToBeBlocked); err != nil {
		return infer.CreateResponse[SandboxBehavioralAnalysisAdvancedSettingsState]{}, err
	}

	apiReq := sandbox_settings.BaAdvancedSettings{
		FileHashesToBeBlocked: sortStringSlice(req.Inputs.FileHashesToBeBlocked),
	}
	if _, err := sandbox_settings.Update(ctx, service, apiReq); err != nil {
		return infer.CreateResponse[SandboxBehavioralAnalysisAdvancedSettingsState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.CreateResponse[SandboxBehavioralAnalysisAdvancedSettingsState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	state := SandboxBehavioralAnalysisAdvancedSettingsState{
		Id:   sandboxSettingsID,
		SandboxBehavioralAnalysisAdvancedSettingsArgs: req.Inputs,
	}
	return infer.CreateResponse[SandboxBehavioralAnalysisAdvancedSettingsState]{
		ID:     sandboxSettingsID,
		Output: state,
	}, nil
}

func (SandboxBehavioralAnalysisAdvancedSettings) Read(ctx context.Context, req infer.ReadRequest[SandboxBehavioralAnalysisAdvancedSettingsArgs, SandboxBehavioralAnalysisAdvancedSettingsState]) (infer.ReadResponse[SandboxBehavioralAnalysisAdvancedSettingsArgs, SandboxBehavioralAnalysisAdvancedSettingsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[SandboxBehavioralAnalysisAdvancedSettingsArgs, SandboxBehavioralAnalysisAdvancedSettingsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	resp, err := sandbox_settings.Get(ctx, service)
	if err != nil {
		return infer.ReadResponse[SandboxBehavioralAnalysisAdvancedSettingsArgs, SandboxBehavioralAnalysisAdvancedSettingsState]{}, err
	}
	hashes := []string{}
	if resp != nil && resp.FileHashesToBeBlocked != nil {
		hashes = sortStringSlice(resp.FileHashesToBeBlocked)
	}

	state := SandboxBehavioralAnalysisAdvancedSettingsState{
		Id: sandboxSettingsID,
		SandboxBehavioralAnalysisAdvancedSettingsArgs: SandboxBehavioralAnalysisAdvancedSettingsArgs{
			FileHashesToBeBlocked: hashes,
		},
	}
	return infer.ReadResponse[SandboxBehavioralAnalysisAdvancedSettingsArgs, SandboxBehavioralAnalysisAdvancedSettingsState]{
		ID:     sandboxSettingsID,
		Inputs: state.SandboxBehavioralAnalysisAdvancedSettingsArgs,
		State:  state,
	}, nil
}

func (SandboxBehavioralAnalysisAdvancedSettings) Update(ctx context.Context, req infer.UpdateRequest[SandboxBehavioralAnalysisAdvancedSettingsArgs, SandboxBehavioralAnalysisAdvancedSettingsState]) (infer.UpdateResponse[SandboxBehavioralAnalysisAdvancedSettingsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[SandboxBehavioralAnalysisAdvancedSettingsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	service := cfg.Client().Service

	if err := validateSandboxHashes(req.Inputs.FileHashesToBeBlocked); err != nil {
		return infer.UpdateResponse[SandboxBehavioralAnalysisAdvancedSettingsState]{}, err
	}

	sortedNew := sortStringSlice(req.Inputs.FileHashesToBeBlocked)
	current, err := sandbox_settings.Get(ctx, service)
	if err != nil {
		return infer.UpdateResponse[SandboxBehavioralAnalysisAdvancedSettingsState]{}, err
	}
	currentSorted := []string{}
	if current != nil && current.FileHashesToBeBlocked != nil {
		currentSorted = sortStringSlice(current.FileHashesToBeBlocked)
	}
	if strings.Join(sortedNew, ",") != strings.Join(currentSorted, ",") {
		if _, err := sandbox_settings.Update(ctx, service, sandbox_settings.BaAdvancedSettings{FileHashesToBeBlocked: sortedNew}); err != nil {
			return infer.UpdateResponse[SandboxBehavioralAnalysisAdvancedSettingsState]{}, err
		}
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[SandboxBehavioralAnalysisAdvancedSettingsState]{}, activationErr
		}
	}

	state := SandboxBehavioralAnalysisAdvancedSettingsState{
		Id:   sandboxSettingsID,
		SandboxBehavioralAnalysisAdvancedSettingsArgs: req.Inputs,
	}
	return infer.UpdateResponse[SandboxBehavioralAnalysisAdvancedSettingsState]{Output: state}, nil
}

func (SandboxBehavioralAnalysisAdvancedSettings) Delete(ctx context.Context, req infer.DeleteRequest[SandboxBehavioralAnalysisAdvancedSettingsState]) (infer.DeleteResponse, error) {
	return infer.DeleteResponse{}, nil
}

func (SandboxBehavioralAnalysisAdvancedSettings) Annotate(a infer.Annotator) {
	describeResource(a, &SandboxBehavioralAnalysisAdvancedSettings{}, `The zia_sandbox_behavioral_analysis_advanced_settings resource manages the list of MD5 file hashes that are blocked by the ZIA sandbox behavioral analysis engine. This is a singleton resource. Only MD5 hashes (32 characters) are supported.

For more information, see the [ZIA Cloud Sandbox documentation](https://help.zscaler.com/zia/about-cloud-sandbox-policies).

{{% examples %}}
## Example Usage

{{% example %}}
### Block File Hashes via Sandbox Settings

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.SandboxBehavioralAnalysisAdvancedSettings("example", {
    fileHashesToBeBlocked: [
        "42914d6d213a20a2684064be5c80ffa9",
        "c0202cf6aeab8437c638533d14563d35",
    ],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.SandboxBehavioralAnalysisAdvancedSettings("example",
    file_hashes_to_be_blocked=[
        "42914d6d213a20a2684064be5c80ffa9",
        "c0202cf6aeab8437c638533d14563d35",
    ],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:SandboxBehavioralAnalysisAdvancedSettings
    properties:
      fileHashesToBeBlocked:
        - 42914d6d213a20a2684064be5c80ffa9
        - c0202cf6aeab8437c638533d14563d35
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

> This is a singleton resource. Import is not applicable.
`)
}

func (a *SandboxBehavioralAnalysisAdvancedSettingsArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.FileHashesToBeBlocked, "List of MD5 file hashes to be blocked. Each hash must be exactly 32 characters (MD5 format). SHA1 and SHA256 are not supported.")
}

func (s *SandboxBehavioralAnalysisAdvancedSettingsState) Annotate(a infer.Annotator) {
	a.Describe(&s.Id, "The internal resource identifier for the sandbox settings.")
}

var _ infer.CustomResource[SandboxBehavioralAnalysisAdvancedSettingsArgs, SandboxBehavioralAnalysisAdvancedSettingsState] = SandboxBehavioralAnalysisAdvancedSettings{}
