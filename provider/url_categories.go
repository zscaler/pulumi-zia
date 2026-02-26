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

// Package provider implements the URL Categories resource.
// Adopted from terraform-provider-zia resource_zia_url_categories.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/urlcategories"
)

// UrlCategory implements the zia:index:UrlCategory resource.
type UrlCategory struct{}

// UrlCategoryScopeInput is a scope block for URL categories.
type UrlCategoryScopeInput struct {
	Type                     *string `pulumi:"type,optional"`
	ScopeGroupMemberEntities []int   `pulumi:"scopeGroupMemberEntities,optional"`
	ScopeEntities            []int   `pulumi:"scopeEntities,optional"`
}

// UrlCategoryUrlKeywordCountsInput is the url_keyword_counts block.
type UrlCategoryUrlKeywordCountsInput struct {
	TotalURLCount            *int `pulumi:"totalUrlCount,optional"`
	RetainParentURLCount     *int `pulumi:"retainParentUrlCount,optional"`
	TotalKeywordCount        *int `pulumi:"totalKeywordCount,optional"`
	RetainParentKeywordCount *int `pulumi:"retainParentKeywordCount,optional"`
}

// UrlCategoryArgs are the inputs.
type UrlCategoryArgs struct {
	ConfiguredName                   *string                        `pulumi:"configuredName,optional"`
	Description                     *string                        `pulumi:"description,optional"`
	Urls                            []string                      `pulumi:"urls,optional"`
	Keywords                        []string                      `pulumi:"keywords,optional"`
	KeywordsRetainingParentCategory []string                      `pulumi:"keywordsRetainingParentCategory,optional"`
	DbCategorizedUrls               []string                      `pulumi:"dbCategorizedUrls,optional"`
	CustomCategory                  *bool                         `pulumi:"customCategory,optional"`
	Scopes                          []UrlCategoryScopeInput       `pulumi:"scopes,optional"`
	Type                            *string                       `pulumi:"type,optional"`
	UrlType                         *string                       `pulumi:"urlType,optional"`
	SuperCategory                   *string                       `pulumi:"superCategory,optional"`
	IpRanges                        []string                      `pulumi:"ipRanges,optional"`
	IpRangesRetainingParentCategory []string                      `pulumi:"ipRangesRetainingParentCategory,optional"`
	RegexPatterns                   []string                      `pulumi:"regexPatterns,optional"`
	RegexPatternsRetainingParentCategory []string                  `pulumi:"regexPatternsRetainingParentCategory,optional"`
	UrlKeywordCounts                *UrlCategoryUrlKeywordCountsInput `pulumi:"urlKeywordCounts,optional"`
}

// UrlCategoryState is the persisted state.
type UrlCategoryState struct {
	UrlCategoryArgs
	CategoryId                         *string `pulumi:"categoryId"`
	Val                                *int    `pulumi:"val,optional"`
	Editable                           *bool   `pulumi:"editable,optional"`
	CustomUrlsCount                    *int    `pulumi:"customUrlsCount,optional"`
	UrlsRetainingParentCategoryCount   *int    `pulumi:"urlsRetainingParentCategoryCount,optional"`
	CustomIpRangesCount                *int    `pulumi:"customIpRangesCount,optional"`
	IpRangesRetainingParentCategoryCount *int   `pulumi:"ipRangesRetainingParentCategoryCount,optional"`
}

func (UrlCategory) Check(ctx context.Context, req infer.CheckRequest) (infer.CheckResponse[UrlCategoryArgs], error) {
	inputs, failures, err := infer.DefaultCheck[UrlCategoryArgs](ctx, req.NewInputs)
	if err != nil {
		return infer.CheckResponse[UrlCategoryArgs]{}, err
	}
	if len(failures) > 0 {
		return infer.CheckResponse[UrlCategoryArgs]{Failures: failures}, nil
	}
	if inputs.Description != nil && len(*inputs.Description) > 256 {
		return infer.CheckResponse[UrlCategoryArgs]{Failures: []p.CheckFailure{{
			Property: "description",
			Reason:   "description must be at most 256 characters",
		}}}, nil
	}
	if inputs.Type != nil {
		t := strings.ToUpper(*inputs.Type)
		if t != "URL_CATEGORY" && t != "TLD_CATEGORY" && t != "ALL" {
			return infer.CheckResponse[UrlCategoryArgs]{Failures: []p.CheckFailure{{
				Property: "type",
				Reason:   "type must be URL_CATEGORY, TLD_CATEGORY, or ALL",
			}}}, nil
		}
	}
	if inputs.UrlType != nil {
		ut := strings.ToUpper(*inputs.UrlType)
		if ut != "EXACT" && ut != "REGEX" {
			return infer.CheckResponse[UrlCategoryArgs]{Failures: []p.CheckFailure{{
				Property: "urlType",
				Reason:   "urlType must be EXACT or REGEX",
			}}}, nil
		}
	}
	for i, s := range inputs.Scopes {
		if s.Type != nil {
			st := strings.ToUpper(*s.Type)
			if st != "ORGANIZATION" && st != "DEPARTMENT" && st != "LOCATION" && st != "LOCATION_GROUP" {
				return infer.CheckResponse[UrlCategoryArgs]{Failures: []p.CheckFailure{{
					Property: fmt.Sprintf("scopes[%d].type", i),
					Reason:   "scope type must be ORGANIZATION, DEPARTMENT, LOCATION, or LOCATION_GROUP",
				}}}, nil
			}
		}
	}
	return infer.CheckResponse[UrlCategoryArgs]{Inputs: inputs}, nil
}

func (UrlCategory) Create(ctx context.Context, req infer.CreateRequest[UrlCategoryArgs]) (infer.CreateResponse[UrlCategoryState], error) {
	if req.DryRun {
		s := UrlCategoryState{UrlCategoryArgs: req.Inputs, CategoryId: stringPtr("preview"), Val: intPtr(0)}
		return infer.CreateResponse[UrlCategoryState]{ID: "preview", Output: s}, nil
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[UrlCategoryState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := urlCategoryArgsToAPI(req.Inputs, "", 0)
	resp, err := urlcategories.CreateURLCategories(ctx, service, &apiReq)
	if err != nil {
		return infer.CreateResponse[UrlCategoryState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[UrlCategoryState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	state := urlCategoryToState(resp, &req.Inputs)
	return infer.CreateResponse[UrlCategoryState]{
		ID:     resp.ID,
		Output: state,
	}, nil
}

func (UrlCategory) Read(ctx context.Context, req infer.ReadRequest[UrlCategoryArgs, UrlCategoryState]) (infer.ReadResponse[UrlCategoryArgs, UrlCategoryState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[UrlCategoryArgs, UrlCategoryState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := req.ID
	if id == "" {
		return infer.ReadResponse[UrlCategoryArgs, UrlCategoryState]{}, nil
	}

	allCategories, err := urlcategories.GetAll(ctx, service, true, false, "ALL")
	if err != nil {
		return infer.ReadResponse[UrlCategoryArgs, UrlCategoryState]{}, err
	}

	var resp *urlcategories.URLCategory
	for i := range allCategories {
		if allCategories[i].ID == id {
			resp = &allCategories[i]
			break
		}
	}

	if resp == nil {
		individualResp, err := urlcategories.Get(ctx, service, id)
		if err != nil {
			if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.ReadResponse[UrlCategoryArgs, UrlCategoryState]{}, nil
			}
			return infer.ReadResponse[UrlCategoryArgs, UrlCategoryState]{}, err
		}
		resp = individualResp
	}

	state := urlCategoryToState(resp, nil)
	return infer.ReadResponse[UrlCategoryArgs, UrlCategoryState]{
		ID:     resp.ID,
		Inputs: state.UrlCategoryArgs,
		State:  state,
	}, nil
}

func (UrlCategory) Update(ctx context.Context, req infer.UpdateRequest[UrlCategoryArgs, UrlCategoryState]) (infer.UpdateResponse[UrlCategoryState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[UrlCategoryState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := req.ID
	if id == "" {
		return infer.UpdateResponse[UrlCategoryState]{}, fmt.Errorf("url category ID not set")
	}

	currentCategory, err := urlcategories.Get(ctx, service, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.UpdateResponse[UrlCategoryState]{}, nil
		}
		return infer.UpdateResponse[UrlCategoryState]{}, fmt.Errorf("failed to get current url category: %w", err)
	}

	desiredCategory := urlCategoryArgsToAPI(req.Inputs, id, currentCategory.Val)

	currentUrls := stringSliceToMap(currentCategory.Urls)
	desiredUrls := stringSliceToMap(desiredCategory.Urls)

	var urlsToAdd, urlsToRemove []string
	for url := range desiredUrls {
		if !currentUrls[url] {
			urlsToAdd = append(urlsToAdd, url)
		}
	}
	for url := range currentUrls {
		if !desiredUrls[url] {
			urlsToRemove = append(urlsToRemove, url)
		}
	}

	hasAdds := len(urlsToAdd) > 0
	hasRemoves := len(urlsToRemove) > 0
	hasBoth := hasAdds && hasRemoves

	if hasBoth {
		removeCategory := desiredCategory
		removeCategory.Urls = urlsToRemove
		if _, _, err := urlcategories.UpdateURLCategories(ctx, service, id, &removeCategory, "REMOVE_FROM_LIST"); err != nil {
			return infer.UpdateResponse[UrlCategoryState]{}, fmt.Errorf("failed to remove URLs: %w", err)
		}
		if len(urlsToAdd) > 0 {
			addCategory := desiredCategory
			addCategory.Urls = urlsToAdd
			if _, _, err := urlcategories.UpdateURLCategories(ctx, service, id, &addCategory, "ADD_TO_LIST"); err != nil {
				return infer.UpdateResponse[UrlCategoryState]{}, fmt.Errorf("failed to add URLs: %w", err)
			}
		}
	} else if hasAdds {
		addCategory := desiredCategory
		addCategory.Urls = urlsToAdd
		if _, _, err := urlcategories.UpdateURLCategories(ctx, service, id, &addCategory, "ADD_TO_LIST"); err != nil {
			return infer.UpdateResponse[UrlCategoryState]{}, fmt.Errorf("failed to add URLs: %w", err)
		}
	} else if hasRemoves {
		removeCategory := desiredCategory
		removeCategory.Urls = urlsToRemove
		if _, _, err := urlcategories.UpdateURLCategories(ctx, service, id, &removeCategory, "REMOVE_FROM_LIST"); err != nil {
			return infer.UpdateResponse[UrlCategoryState]{}, fmt.Errorf("failed to remove URLs: %w", err)
		}
	} else {
		if _, _, err := urlcategories.UpdateURLCategories(ctx, service, id, &desiredCategory, ""); err != nil {
			return infer.UpdateResponse[UrlCategoryState]{}, err
		}
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[UrlCategoryState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	updated, err := urlcategories.Get(ctx, service, id)
	if err != nil {
		return infer.UpdateResponse[UrlCategoryState]{}, err
	}
	state := urlCategoryToState(updated, &req.Inputs)
	return infer.UpdateResponse[UrlCategoryState]{Output: state}, nil
}

func (UrlCategory) Delete(ctx context.Context, req infer.DeleteRequest[UrlCategoryState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := req.ID
	if id == "" {
		return infer.DeleteResponse{}, nil
	}

	if _, err := urlcategories.DeleteURLCategories(ctx, service, id); err != nil {
		return infer.DeleteResponse{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	return infer.DeleteResponse{}, nil
}

func (UrlCategory) Annotate(a infer.Annotator) {
	describeResource(a, &UrlCategory{}, `The zia_url_categories resource manages custom URL categories in the Zscaler Internet Access (ZIA) cloud service. Custom URL categories allow administrators to define their own groupings of URLs, keywords, and IP ranges for use in URL filtering policies.

For more information, see the [ZIA URL Categories documentation](https://help.zscaler.com/zia/url-categories).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Custom URL Category

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.UrlCategory("example", {
    configuredName: "Example Custom Category",
    description: "Custom category for internal tools",
    superCategory: "USER_DEFINED",
    type: "URL_CATEGORY",
    urls: [
        "internal.example.com",
        "tools.example.com",
    ],
    customCategory: true,
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.UrlCategory("example",
    configured_name="Example Custom Category",
    description="Custom category for internal tools",
    super_category="USER_DEFINED",
    type="URL_CATEGORY",
    urls=[
        "internal.example.com",
        "tools.example.com",
    ],
    custom_category=True,
)
`+tripleBacktick("")+`

`+tripleBacktick("go")+`
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	zia "github.com/zscaler/pulumi-zia/sdk/go/pulumi-zia"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := zia.NewUrlCategory(ctx, "example", &zia.UrlCategoryArgs{
			ConfiguredName: pulumi.StringRef("Example Custom Category"),
			Description:    pulumi.StringRef("Custom category for internal tools"),
			SuperCategory:  pulumi.StringRef("USER_DEFINED"),
			Type:           pulumi.StringRef("URL_CATEGORY"),
			Urls:           pulumi.ToStringArray([]string{"internal.example.com", "tools.example.com"}),
			CustomCategory: pulumi.BoolRef(true),
		})
		return err
	})
}
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:UrlCategory
    properties:
      configuredName: Example Custom Category
      description: Custom category for internal tools
      superCategory: USER_DEFINED
      type: URL_CATEGORY
      urls:
        - internal.example.com
        - tools.example.com
      customCategory: true
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing URL Category can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:UrlCategory example CUSTOM_01
`+tripleBacktick("")+`
`)
}

func (a *UrlCategoryArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.ConfiguredName, "The name of the URL category. Must be unique.")
	ann.Describe(&a.Description, "A description of the URL category. Maximum 256 characters.")
	ann.Describe(&a.Urls, "List of custom URLs to add to the category.")
	ann.Describe(&a.Keywords, "List of custom keywords associated with the URL category.")
	ann.Describe(&a.KeywordsRetainingParentCategory, "List of keywords that retain their parent category classification.")
	ann.Describe(&a.DbCategorizedUrls, "URLs added to a custom URL category that have been categorized by the Zscaler database.")
	ann.Describe(&a.CustomCategory, "If true, this is a custom URL category. Set to true for custom categories.")
	ann.Describe(&a.Scopes, "Scopes for the custom URL category, defining location or department restrictions.")
	ann.Describe(&a.Type, "The type of URL category. Valid values: `URL_CATEGORY`, `TLD_CATEGORY`, `ALL`.")
	ann.Describe(&a.UrlType, "The URL type. Valid values: `EXACT`, `REGEX`.")
	ann.Describe(&a.SuperCategory, "The super category for the URL category (e.g., `USER_DEFINED`).")
	ann.Describe(&a.IpRanges, "List of custom IP address ranges associated with the URL category.")
	ann.Describe(&a.IpRangesRetainingParentCategory, "List of IP ranges that retain their parent category classification.")
	ann.Describe(&a.RegexPatterns, "List of regex-based patterns for URL matching.")
	ann.Describe(&a.RegexPatternsRetainingParentCategory, "List of regex patterns that retain their parent category classification.")
	ann.Describe(&a.UrlKeywordCounts, "URL and keyword count statistics for the category.")
}

func (s *UrlCategoryState) Annotate(a infer.Annotator) {
	a.Describe(&s.CategoryId, "The system-generated ID of the URL category.")
	a.Describe(&s.Val, "The internal numeric value of the URL category.")
	a.Describe(&s.Editable, "Whether the URL category is editable.")
	a.Describe(&s.CustomUrlsCount, "The number of custom URLs in the category.")
	a.Describe(&s.UrlsRetainingParentCategoryCount, "The number of URLs retaining parent category.")
	a.Describe(&s.CustomIpRangesCount, "The number of custom IP ranges in the category.")
	a.Describe(&s.IpRangesRetainingParentCategoryCount, "The number of IP ranges retaining parent category.")
}

func stringSliceToMap(slice []string) map[string]bool {
	result := make(map[string]bool, len(slice))
	for _, s := range slice {
		result[s] = true
	}
	return result
}

func urlCategoryArgsToAPI(args UrlCategoryArgs, id string, val int) urlcategories.URLCategory {
	cat := urlcategories.URLCategory{
		ID:                                   id,
		Val:                                  val,
		ConfiguredName:                       ptrToString(args.ConfiguredName),
		Keywords:                             args.Keywords,
		KeywordsRetainingParentCategory:      args.KeywordsRetainingParentCategory,
		Urls:                                 args.Urls,
		DBCategorizedUrls:                    args.DbCategorizedUrls,
		IPRanges:                             args.IpRanges,
		IPRangesRetainingParentCategory:      args.IpRangesRetainingParentCategory,
		CustomCategory:                       ptrToBool(args.CustomCategory),
		SuperCategory:                        ptrToString(args.SuperCategory),
		Description:                          ptrToString(args.Description),
		Type:                                 ptrToString(args.Type),
		RegexPatterns:                        args.RegexPatterns,
		RegexPatternsRetainingParentCategory: args.RegexPatternsRetainingParentCategory,
		UrlType:                              ptrToString(args.UrlType),
	}
	for _, s := range args.Scopes {
		scope := urlcategories.Scopes{
			Type:                     ptrToString(s.Type),
			ScopeGroupMemberEntities: idsToIDNameExtensions(s.ScopeGroupMemberEntities),
			ScopeEntities:            idsToIDNameExtensions(s.ScopeEntities),
		}
		cat.Scopes = append(cat.Scopes, scope)
	}
	if args.UrlKeywordCounts != nil {
		cat.URLKeywordCounts = &urlcategories.URLKeywordCounts{
			TotalURLCount:            ptrToInt(args.UrlKeywordCounts.TotalURLCount),
			RetainParentURLCount:     ptrToInt(args.UrlKeywordCounts.RetainParentURLCount),
			TotalKeywordCount:        ptrToInt(args.UrlKeywordCounts.TotalKeywordCount),
			RetainParentKeywordCount: ptrToInt(args.UrlKeywordCounts.RetainParentKeywordCount),
		}
	}
	return cat
}

func urlCategoryToState(r *urlcategories.URLCategory, args *UrlCategoryArgs) UrlCategoryState {
	state := UrlCategoryState{
		CategoryId:                         stringPtr(r.ID),
		Val:                                intPtr(r.Val),
		Editable:                           boolPtr(r.Editable),
		CustomUrlsCount:                    intPtr(r.CustomUrlsCount),
		UrlsRetainingParentCategoryCount:   intPtr(r.UrlsRetainingParentCategoryCount),
		CustomIpRangesCount:                intPtr(r.CustomIpRangesCount),
		IpRangesRetainingParentCategoryCount: intPtr(r.IPRangesRetainingParentCategoryCount),
	}
	if args != nil {
		state.UrlCategoryArgs = *args
	} else {
		state.UrlCategoryArgs = UrlCategoryArgs{
			ConfiguredName:                   stringPtr(r.ConfiguredName),
			Description:                     stringPtr(r.Description),
			Urls:                            r.Urls,
			Keywords:                        r.Keywords,
			KeywordsRetainingParentCategory: r.KeywordsRetainingParentCategory,
			DbCategorizedUrls:               r.DBCategorizedUrls,
			CustomCategory:                   boolPtr(r.CustomCategory),
			Type:                            stringPtr(r.Type),
			UrlType:                         stringPtr(r.UrlType),
			SuperCategory:                   stringPtr(r.SuperCategory),
			IpRanges:                        r.IPRanges,
			IpRangesRetainingParentCategory: r.IPRangesRetainingParentCategory,
			RegexPatterns:                   r.RegexPatterns,
			RegexPatternsRetainingParentCategory: r.RegexPatternsRetainingParentCategory,
		}
		for _, s := range r.Scopes {
			state.Scopes = append(state.Scopes, UrlCategoryScopeInput{
				Type:                     stringPtr(s.Type),
				ScopeGroupMemberEntities: idNameExtensionsToIDs(s.ScopeGroupMemberEntities),
				ScopeEntities:            idNameExtensionsToIDs(s.ScopeEntities),
			})
		}
		if r.URLKeywordCounts != nil {
			state.UrlKeywordCounts = &UrlCategoryUrlKeywordCountsInput{
				TotalURLCount:            intPtr(r.URLKeywordCounts.TotalURLCount),
				RetainParentURLCount:     intPtr(r.URLKeywordCounts.RetainParentURLCount),
				TotalKeywordCount:        intPtr(r.URLKeywordCounts.TotalKeywordCount),
				RetainParentKeywordCount: intPtr(r.URLKeywordCounts.RetainParentKeywordCount),
			}
		}
	}
	return state
}

func ptrToInt(p *int) int {
	if p == nil {
		return 0
	}
	return *p
}

var _ infer.CustomResource[UrlCategoryArgs, UrlCategoryState] = UrlCategory{}
