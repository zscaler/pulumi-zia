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

// Package provider implements the DLP Dictionary resource.
// Adopted from terraform-provider-zia resource_zia_dlp_dictionaries.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/dlp/dlpdictionaries"
)

// DlpDictionary implements the zia:index:DlpDictionary resource.
type DlpDictionary struct{}

// DlpDictionaryPhraseInput is a phrase entry for a custom DLP dictionary.
type DlpDictionaryPhraseInput struct {
	Action *string `pulumi:"action,optional"`
	Phrase *string `pulumi:"phrase,optional"`
}

// DlpDictionaryPatternInput is a pattern entry for a custom DLP dictionary.
type DlpDictionaryPatternInput struct {
	Action  *string `pulumi:"action,optional"`
	Pattern *string `pulumi:"pattern,optional"`
}

// DlpDictionaryArgs are the inputs.
type DlpDictionaryArgs struct {
	Name                                *string                     `pulumi:"name,optional"`
	Description                         *string                     `pulumi:"description,optional"`
	Custom                              *bool                       `pulumi:"custom,optional"`
	ConfidenceThreshold                 *string                     `pulumi:"confidenceThreshold,optional"`
	CustomPhraseMatchType               *string                     `pulumi:"customPhraseMatchType,optional"`
	DictionaryType                      *string                     `pulumi:"dictionaryType,optional"`
	Phrases                             []DlpDictionaryPhraseInput  `pulumi:"phrases,optional"`
	Patterns                            []DlpDictionaryPatternInput `pulumi:"patterns,optional"`
	HierarchicalIdentifiers             []string                    `pulumi:"hierarchicalIdentifiers,optional"`
	Proximity                           *int                        `pulumi:"proximity,optional"`
	ProximityEnabledForCustomDictionary *bool                       `pulumi:"proximityEnabledForCustomDictionary,optional"`
}

// DlpDictionaryState is the persisted state.
type DlpDictionaryState struct {
	DlpDictionaryArgs
	DictionaryId *int `pulumi:"dictionaryId"`
}

func dlpDictionaryToAPI(args DlpDictionaryArgs, id int) dlpdictionaries.DlpDictionary {
	out := dlpdictionaries.DlpDictionary{
		ID:                                  id,
		Name:                                ptrToString(args.Name),
		Description:                         ptrToString(args.Description),
		Custom:                              ptrToBool(args.Custom),
		ConfidenceThreshold:                 ptrToString(args.ConfidenceThreshold),
		CustomPhraseMatchType:               ptrToString(args.CustomPhraseMatchType),
		DictionaryType:                      ptrToString(args.DictionaryType),
		Proximity:                           ptrToIntDefault(args.Proximity, 0),
		ProximityEnabledForCustomDictionary: ptrToBool(args.ProximityEnabledForCustomDictionary),
		HierarchicalIdentifiers:             args.HierarchicalIdentifiers,
	}
	if len(args.Phrases) > 0 {
		out.Phrases = make([]dlpdictionaries.Phrases, 0, len(args.Phrases))
		for _, p := range args.Phrases {
			out.Phrases = append(out.Phrases, dlpdictionaries.Phrases{
				Action: ptrToString(p.Action),
				Phrase: ptrToString(p.Phrase),
			})
		}
	}
	if len(args.Patterns) > 0 {
		out.Patterns = make([]dlpdictionaries.Patterns, 0, len(args.Patterns))
		for _, p := range args.Patterns {
			out.Patterns = append(out.Patterns, dlpdictionaries.Patterns{
				Action:  ptrToString(p.Action),
				Pattern: ptrToString(p.Pattern),
			})
		}
	}
	return out
}

func (DlpDictionary) Create(ctx context.Context, req infer.CreateRequest[DlpDictionaryArgs]) (infer.CreateResponse[DlpDictionaryState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[DlpDictionaryState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := dlpDictionaryToAPI(req.Inputs, 0)
	resp, _, err := dlpdictionaries.Create(ctx, service, &apiReq)
	if err != nil {
		return infer.CreateResponse[DlpDictionaryState]{}, err
	}
	log.Printf("[INFO] Created ZIA DLP dictionary. ID: %v", resp.ID)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[DlpDictionaryState]{}, activationErr
		}
	}

	state := DlpDictionaryState{
		DlpDictionaryArgs: req.Inputs,
		DictionaryId:      &resp.ID,
	}
	return infer.CreateResponse[DlpDictionaryState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (DlpDictionary) Read(ctx context.Context, req infer.ReadRequest[DlpDictionaryArgs, DlpDictionaryState]) (infer.ReadResponse[DlpDictionaryArgs, DlpDictionaryState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[DlpDictionaryArgs, DlpDictionaryState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.DictionaryId != nil {
		id = *req.State.DictionaryId
	}
	if id == 0 {
		return infer.ReadResponse[DlpDictionaryArgs, DlpDictionaryState]{}, fmt.Errorf("no DLP dictionary id in state")
	}

	resp, err := dlpdictionaries.Get(ctx, service, id)
	if err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.ReadResponse[DlpDictionaryArgs, DlpDictionaryState]{ID: ""}, nil
		}
		return infer.ReadResponse[DlpDictionaryArgs, DlpDictionaryState]{}, err
	}

	phrases := make([]DlpDictionaryPhraseInput, 0, len(resp.Phrases))
	for _, p := range resp.Phrases {
		phrases = append(phrases, DlpDictionaryPhraseInput{
			Action: stringPtr(p.Action),
			Phrase: stringPtr(p.Phrase),
		})
	}
	patterns := make([]DlpDictionaryPatternInput, 0, len(resp.Patterns))
	for _, p := range resp.Patterns {
		patterns = append(patterns, DlpDictionaryPatternInput{
			Action:  stringPtr(p.Action),
			Pattern: stringPtr(p.Pattern),
		})
	}

	args := DlpDictionaryArgs{
		Name:                                stringPtr(resp.Name),
		Description:                         stringPtr(resp.Description),
		Custom:                              boolPtr(resp.Custom),
		ConfidenceThreshold:                 stringPtr(resp.ConfidenceThreshold),
		CustomPhraseMatchType:               stringPtr(resp.CustomPhraseMatchType),
		DictionaryType:                      stringPtr(resp.DictionaryType),
		Phrases:                             phrases,
		Patterns:                            patterns,
		HierarchicalIdentifiers:             resp.HierarchicalIdentifiers,
		Proximity:                           intPtr(resp.Proximity),
		ProximityEnabledForCustomDictionary: boolPtr(resp.ProximityEnabledForCustomDictionary),
	}
	state := DlpDictionaryState{
		DlpDictionaryArgs: args,
		DictionaryId:      &resp.ID,
	}
	return infer.ReadResponse[DlpDictionaryArgs, DlpDictionaryState]{
		ID:     strconv.Itoa(resp.ID),
		Inputs: args,
		State:  state,
	}, nil
}

func (DlpDictionary) Update(ctx context.Context, req infer.UpdateRequest[DlpDictionaryArgs, DlpDictionaryState]) (infer.UpdateResponse[DlpDictionaryState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[DlpDictionaryState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.DictionaryId != nil {
		id = *req.State.DictionaryId
	}
	if id == 0 {
		return infer.UpdateResponse[DlpDictionaryState]{}, fmt.Errorf("no DLP dictionary id in state")
	}

	if _, err := dlpdictionaries.Get(ctx, service, id); err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.UpdateResponse[DlpDictionaryState]{}, nil
		}
		return infer.UpdateResponse[DlpDictionaryState]{}, err
	}

	apiReq := dlpDictionaryToAPI(req.Inputs, id)
	if _, _, err := dlpdictionaries.Update(ctx, service, id, &apiReq); err != nil {
		return infer.UpdateResponse[DlpDictionaryState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[DlpDictionaryState]{}, activationErr
		}
	}

	state := DlpDictionaryState{
		DlpDictionaryArgs: req.Inputs,
		DictionaryId:      &id,
	}
	return infer.UpdateResponse[DlpDictionaryState]{Output: state}, nil
}

func (DlpDictionary) Delete(ctx context.Context, req infer.DeleteRequest[DlpDictionaryState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.DictionaryId != nil {
		id = *req.State.DictionaryId
	}
	if id != 0 {
		if _, err := dlpdictionaries.DeleteDlpDictionary(ctx, service, id); err != nil {
			if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.DeleteResponse{}, nil
			}
			return infer.DeleteResponse{}, err
		}
		log.Printf("[INFO] ZIA DLP dictionary deleted")
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}

	return infer.DeleteResponse{}, nil
}

func (DlpDictionary) Annotate(a infer.Annotator) {
	describeResource(a, &DlpDictionary{}, `The zia_dlp_dictionaries resource manages DLP (Data Loss Prevention) dictionaries in the Zscaler Internet Access (ZIA) cloud service. DLP dictionaries are used to define custom or predefined patterns and phrases that identify sensitive data for DLP policy enforcement.

For more information, see the [ZIA Data Loss Prevention documentation](https://help.zscaler.com/zia/data-loss-prevention).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic DLP Dictionary

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.DlpDictionary("example", {
    name: "Example DLP Dictionary",
    description: "Custom DLP dictionary for detecting sensitive patterns",
    dictionaryType: "PATTERNS_AND_PHRASES",
    customPhraseMatchType: "MATCH_ALL_CUSTOM_PHRASE_PATTERN_DICTIONARY",
    phrases: [{
        action: "PHRASE_COUNT_TYPE_ALL",
        phrase: "confidential",
    }],
    patterns: [{
        action: "PATTERN_COUNT_TYPE_ALL",
        pattern: "\\b\\d{3}-\\d{2}-\\d{4}\\b",
    }],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.DlpDictionary("example",
    name="Example DLP Dictionary",
    description="Custom DLP dictionary for detecting sensitive patterns",
    dictionary_type="PATTERNS_AND_PHRASES",
    custom_phrase_match_type="MATCH_ALL_CUSTOM_PHRASE_PATTERN_DICTIONARY",
    phrases=[{
        "action": "PHRASE_COUNT_TYPE_ALL",
        "phrase": "confidential",
    }],
    patterns=[{
        "action": "PATTERN_COUNT_TYPE_ALL",
        "pattern": "\\b\\d{3}-\\d{2}-\\d{4}\\b",
    }],
)
`+tripleBacktick("")+`

`+tripleBacktick("go")+`
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	zia "github.com/zscaler/pulumi-zia/sdk/go/pulumi-zia"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := zia.NewDlpDictionary(ctx, "example", &zia.DlpDictionaryArgs{
			Name:                  pulumi.StringRef("Example DLP Dictionary"),
			Description:           pulumi.StringRef("Custom DLP dictionary for detecting sensitive patterns"),
			DictionaryType:        pulumi.StringRef("PATTERNS_AND_PHRASES"),
			CustomPhraseMatchType: pulumi.StringRef("MATCH_ALL_CUSTOM_PHRASE_PATTERN_DICTIONARY"),
		})
		return err
	})
}
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:DlpDictionary
    properties:
      name: Example DLP Dictionary
      description: Custom DLP dictionary for detecting sensitive patterns
      dictionaryType: PATTERNS_AND_PHRASES
      customPhraseMatchType: MATCH_ALL_CUSTOM_PHRASE_PATTERN_DICTIONARY
      phrases:
        - action: PHRASE_COUNT_TYPE_ALL
          phrase: confidential
      patterns:
        - action: PATTERN_COUNT_TYPE_ALL
          pattern: "\\b\\d{3}-\\d{2}-\\d{4}\\b"
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing DLP Dictionary can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:DlpDictionary example 12345
`+tripleBacktick("")+`
`)
}

func (a *DlpDictionaryArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the DLP dictionary. Must be unique.")
	ann.Describe(&a.Description, "A description of the DLP dictionary.")
	ann.Describe(&a.Custom, "If true, this is a custom DLP dictionary; false indicates a predefined dictionary.")
	ann.Describe(&a.ConfidenceThreshold, "The DLP confidence threshold. Valid values: `CONFIDENCE_LEVEL_LOW`, `CONFIDENCE_LEVEL_MEDIUM`, `CONFIDENCE_LEVEL_HIGH`.")
	ann.Describe(&a.CustomPhraseMatchType, "The match type for custom phrases. Valid values: `MATCH_ALL_CUSTOM_PHRASE_PATTERN_DICTIONARY`, `MATCH_ANY_CUSTOM_PHRASE_PATTERN_DICTIONARY`.")
	ann.Describe(&a.DictionaryType, "The type of DLP dictionary. Valid values: `PATTERNS_AND_PHRASES`, `EXACT_DATA_MATCH`, `INDEXED_DATA_MATCH`.")
	ann.Describe(&a.Phrases, "List of DLP dictionary phrases with their match actions.")
	ann.Describe(&a.Patterns, "List of DLP dictionary patterns with their match actions.")
	ann.Describe(&a.HierarchicalIdentifiers, "List of hierarchical identifiers for the DLP dictionary.")
	ann.Describe(&a.Proximity, "The proximity length for dictionary matching. Specifies the distance between phrases/patterns for a match.")
	ann.Describe(&a.ProximityEnabledForCustomDictionary, "If true, proximity matching is enabled for this custom DLP dictionary.")
}

func (s *DlpDictionaryState) Annotate(a infer.Annotator) {
	a.Describe(&s.DictionaryId, "The system-generated ID of the DLP dictionary.")
}

var _ infer.CustomResource[DlpDictionaryArgs, DlpDictionaryState] = DlpDictionary{}
