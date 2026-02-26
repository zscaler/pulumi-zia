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

// Package provider implements the Sandbox Behavioral Analysis Advanced Settings V2 resource (singleton).
// Adopted from terraform-provider-zia resource_zia_sandbox_behavioral_analysis_advanced_settings_v2.go.
// Uses Getv2/Updatev2 with md5HashValueList (url, urlComment, type). Delete is a no-op.

package provider

import (
	"context"
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/sandbox/sandbox_settings"
)

const sandboxSettingsV2ID = "sandbox_settings"

// SandboxBehavioralAnalysisAdvancedSettingsV2 implements the zia:index:SandboxBehavioralAnalysisAdvancedSettingsV2 resource.
type SandboxBehavioralAnalysisAdvancedSettingsV2 struct{}

// Md5HashValueInput is a single entry in the MD5 hash value list.
type Md5HashValueInput struct {
	Url        *string `pulumi:"url,optional"`
	UrlComment *string `pulumi:"urlComment,optional"`
	Type       *string `pulumi:"type,optional"` // CUSTOM_FILEHASH_ALLOW, CUSTOM_FILEHASH_DENY, MALWARE
}

// SandboxBehavioralAnalysisAdvancedSettingsV2Args are the inputs.
type SandboxBehavioralAnalysisAdvancedSettingsV2Args struct {
	Md5HashValueList []Md5HashValueInput `pulumi:"md5HashValueList,optional"`
}

// SandboxBehavioralAnalysisAdvancedSettingsV2State is the persisted state.
type SandboxBehavioralAnalysisAdvancedSettingsV2State struct {
	SandboxBehavioralAnalysisAdvancedSettingsV2Args
}

func md5HashValueInputsToAPI(list []Md5HashValueInput) []sandbox_settings.Md5HashValue {
	out := make([]sandbox_settings.Md5HashValue, 0, len(list))
	for _, v := range list {
		url := ptrToString(v.Url)
		typ := ptrToString(v.Type)
		// Skip empty blocks (same as Terraform expandMd5HashValueList)
		if url == "" && typ == "" {
			continue
		}
		out = append(out, sandbox_settings.Md5HashValue{
			URL:        url,
			URLComment: ptrToString(v.UrlComment),
			Type:       typ,
		})
	}
	return out
}

func md5HashValueListsEqual(a, b []sandbox_settings.Md5HashValue) bool {
	if len(a) != len(b) {
		return false
	}
	sortedA := make([]sandbox_settings.Md5HashValue, len(a))
	copy(sortedA, a)
	sort.Slice(sortedA, func(i, j int) bool {
		keyI := sortedA[i].URL + "|" + sortedA[i].Type
		keyJ := sortedA[j].URL + "|" + sortedA[j].Type
		return keyI < keyJ
	})
	sortedB := make([]sandbox_settings.Md5HashValue, len(b))
	copy(sortedB, b)
	sort.Slice(sortedB, func(i, j int) bool {
		keyI := sortedB[i].URL + "|" + sortedB[i].Type
		keyJ := sortedB[j].URL + "|" + sortedB[j].Type
		return keyI < keyJ
	})
	for i := range sortedA {
		if sortedA[i].URL != sortedB[i].URL ||
			sortedA[i].URLComment != sortedB[i].URLComment ||
			sortedA[i].Type != sortedB[i].Type {
			return false
		}
	}
	return true
}

func (SandboxBehavioralAnalysisAdvancedSettingsV2) Create(ctx context.Context, req infer.CreateRequest[SandboxBehavioralAnalysisAdvancedSettingsV2Args]) (infer.CreateResponse[SandboxBehavioralAnalysisAdvancedSettingsV2State], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[SandboxBehavioralAnalysisAdvancedSettingsV2State]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	md5List := md5HashValueInputsToAPI(req.Inputs.Md5HashValueList)
	payload := sandbox_settings.Md5HashValueListPayload{
		Md5HashValueList: md5List,
	}
	if _, err := sandbox_settings.Updatev2(ctx, service, payload); err != nil {
		return infer.CreateResponse[SandboxBehavioralAnalysisAdvancedSettingsV2State]{}, err
	}
	log.Printf("[INFO] Created/Updated ZIA Sandbox Behavioral Analysis Advanced Settings V2")

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[SandboxBehavioralAnalysisAdvancedSettingsV2State]{}, activationErr
		}
	}

	state := SandboxBehavioralAnalysisAdvancedSettingsV2State{SandboxBehavioralAnalysisAdvancedSettingsV2Args: req.Inputs}
	return infer.CreateResponse[SandboxBehavioralAnalysisAdvancedSettingsV2State]{
		ID:     sandboxSettingsV2ID,
		Output: state,
	}, nil
}

func (SandboxBehavioralAnalysisAdvancedSettingsV2) Read(ctx context.Context, req infer.ReadRequest[SandboxBehavioralAnalysisAdvancedSettingsV2Args, SandboxBehavioralAnalysisAdvancedSettingsV2State]) (infer.ReadResponse[SandboxBehavioralAnalysisAdvancedSettingsV2Args, SandboxBehavioralAnalysisAdvancedSettingsV2State], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[SandboxBehavioralAnalysisAdvancedSettingsV2Args, SandboxBehavioralAnalysisAdvancedSettingsV2State]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	resp, err := sandbox_settings.Getv2(ctx, service)
	if err != nil {
		return infer.ReadResponse[SandboxBehavioralAnalysisAdvancedSettingsV2Args, SandboxBehavioralAnalysisAdvancedSettingsV2State]{}, err
	}

	md5List := []Md5HashValueInput{}
	if resp != nil && len(resp.Md5HashValueList) > 0 {
		for _, v := range resp.Md5HashValueList {
			md5List = append(md5List, Md5HashValueInput{
				Url:        stringPtr(v.URL),
				UrlComment: stringPtr(v.URLComment),
				Type:       stringPtr(v.Type),
			})
		}
	}

	args := SandboxBehavioralAnalysisAdvancedSettingsV2Args{Md5HashValueList: md5List}
	state := SandboxBehavioralAnalysisAdvancedSettingsV2State{SandboxBehavioralAnalysisAdvancedSettingsV2Args: args}
	return infer.ReadResponse[SandboxBehavioralAnalysisAdvancedSettingsV2Args, SandboxBehavioralAnalysisAdvancedSettingsV2State]{
		ID:     sandboxSettingsV2ID,
		Inputs: args,
		State:  state,
	}, nil
}

func (SandboxBehavioralAnalysisAdvancedSettingsV2) Update(ctx context.Context, req infer.UpdateRequest[SandboxBehavioralAnalysisAdvancedSettingsV2Args, SandboxBehavioralAnalysisAdvancedSettingsV2State]) (infer.UpdateResponse[SandboxBehavioralAnalysisAdvancedSettingsV2State], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[SandboxBehavioralAnalysisAdvancedSettingsV2State]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	statePayload := md5HashValueInputsToAPI(req.Inputs.Md5HashValueList)
	payload := sandbox_settings.Md5HashValueListPayload{
		Md5HashValueList: statePayload,
	}

	current, err := sandbox_settings.Getv2(ctx, service)
	if err != nil {
		return infer.UpdateResponse[SandboxBehavioralAnalysisAdvancedSettingsV2State]{}, err
	}
	currentList := []sandbox_settings.Md5HashValue{}
	if current != nil {
		currentList = current.Md5HashValueList
	}

	if !md5HashValueListsEqual(statePayload, currentList) {
		if _, err := sandbox_settings.Updatev2(ctx, service, payload); err != nil {
			return infer.UpdateResponse[SandboxBehavioralAnalysisAdvancedSettingsV2State]{}, err
		}
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[SandboxBehavioralAnalysisAdvancedSettingsV2State]{}, activationErr
		}
	}

	state := SandboxBehavioralAnalysisAdvancedSettingsV2State{SandboxBehavioralAnalysisAdvancedSettingsV2Args: req.Inputs}
	return infer.UpdateResponse[SandboxBehavioralAnalysisAdvancedSettingsV2State]{Output: state}, nil
}

func (SandboxBehavioralAnalysisAdvancedSettingsV2) Delete(ctx context.Context, req infer.DeleteRequest[SandboxBehavioralAnalysisAdvancedSettingsV2State]) (infer.DeleteResponse, error) {
	log.Printf("[INFO] Sandbox Behavioral Analysis Advanced Settings V2: delete is a no-op for singleton resource")
	return infer.DeleteResponse{}, nil
}

func (SandboxBehavioralAnalysisAdvancedSettingsV2) Annotate(a infer.Annotator) {
	describeResource(a, &SandboxBehavioralAnalysisAdvancedSettingsV2{}, `The zia_sandbox_behavioral_analysis_advanced_settings_v2 resource manages the V2 MD5 hash value list for the ZIA sandbox behavioral analysis engine. Each entry includes the hash URL, an optional comment, and a type (allow, deny, or malware). This is a singleton resource.

For more information, see the [ZIA Cloud Sandbox documentation](https://help.zscaler.com/zia/about-cloud-sandbox-policies).

{{% examples %}}
## Example Usage

{{% example %}}
### Manage MD5 Hash Value List

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.SandboxBehavioralAnalysisAdvancedSettingsV2("example", {
    md5HashValueList: [
        {
            url: "42914d6d213a20a2684064be5c80ffa9",
            urlComment: "Known safe file hash",
            type: "CUSTOM_FILEHASH_ALLOW",
        },
        {
            url: "c0202cf6aeab8437c638533d14563d35",
            urlComment: "Known malicious file hash",
            type: "CUSTOM_FILEHASH_DENY",
        },
    ],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.SandboxBehavioralAnalysisAdvancedSettingsV2("example",
    md5_hash_value_list=[
        {
            "url": "42914d6d213a20a2684064be5c80ffa9",
            "url_comment": "Known safe file hash",
            "type": "CUSTOM_FILEHASH_ALLOW",
        },
        {
            "url": "c0202cf6aeab8437c638533d14563d35",
            "url_comment": "Known malicious file hash",
            "type": "CUSTOM_FILEHASH_DENY",
        },
    ],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:SandboxBehavioralAnalysisAdvancedSettingsV2
    properties:
      md5HashValueList:
        - url: 42914d6d213a20a2684064be5c80ffa9
          urlComment: Known safe file hash
          type: CUSTOM_FILEHASH_ALLOW
        - url: c0202cf6aeab8437c638533d14563d35
          urlComment: Known malicious file hash
          type: CUSTOM_FILEHASH_DENY
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

> This is a singleton resource. Import is not applicable.
`)
}

func (a *SandboxBehavioralAnalysisAdvancedSettingsV2Args) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Md5HashValueList, "List of MD5 hash value entries. Each entry contains a hash URL, optional comment, and type.")
}

var _ infer.CustomResource[SandboxBehavioralAnalysisAdvancedSettingsV2Args, SandboxBehavioralAnalysisAdvancedSettingsV2State] = SandboxBehavioralAnalysisAdvancedSettingsV2{}
