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

// Package provider implements the Predefined URL Categories resource.
// Adopted from terraform-provider-zia resource_zia_url_categories_predefined.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/urlcategories"
)

// UrlCategoryPredefined implements the zia:index:UrlCategoryPredefined resource.
// Manages custom URLs/keywords/IPs added to a predefined (built-in) URL category.
// Delete is a no-op (predefined categories cannot be deleted).
type UrlCategoryPredefined struct{}

// UrlCategoryPredefinedArgs are the inputs.
type UrlCategoryPredefinedArgs struct {
	Name                           string   `pulumi:"name"`
	Urls                           []string `pulumi:"urls,optional"`
	Keywords                       []string `pulumi:"keywords,optional"`
	KeywordsRetainingParentCategory []string `pulumi:"keywordsRetainingParentCategory,optional"`
	IpRanges                       []string `pulumi:"ipRanges,optional"`
	IpRangesRetainingParentCategory []string `pulumi:"ipRangesRetainingParentCategory,optional"`
}

// UrlCategoryPredefinedState is the persisted state.
type UrlCategoryPredefinedState struct {
	UrlCategoryPredefinedArgs
	CategoryId                         *string `pulumi:"categoryId"`
	ConfiguredName                     *string `pulumi:"configuredName,optional"`
	SuperCategory                      *string `pulumi:"superCategory,optional"`
	UrlType                            *string `pulumi:"urlType,optional"`
	Type                               *string `pulumi:"type,optional"`
	Val                                *int    `pulumi:"val,optional"`
	Editable                           *bool   `pulumi:"editable,optional"`
	DbCategorizedUrls                   []string `pulumi:"dbCategorizedUrls,optional"`
	CustomUrlsCount                    *int    `pulumi:"customUrlsCount,optional"`
	UrlsRetainingParentCategoryCount   *int    `pulumi:"urlsRetainingParentCategoryCount,optional"`
	CustomIpRangesCount                *int    `pulumi:"customIpRangesCount,optional"`
	IpRangesRetainingParentCategoryCount *int   `pulumi:"ipRangesRetainingParentCategoryCount,optional"`
}

func resolvePredefinedCategory(allCategories []urlcategories.URLCategory, identifier string) *urlcategories.URLCategory {
	for i := range allCategories {
		if allCategories[i].CustomCategory {
			continue
		}
		if strings.EqualFold(allCategories[i].ID, identifier) || strings.EqualFold(allCategories[i].ConfiguredName, identifier) {
			return &allCategories[i]
		}
	}
	return nil
}

func (UrlCategoryPredefined) Create(ctx context.Context, req infer.CreateRequest[UrlCategoryPredefinedArgs]) (infer.CreateResponse[UrlCategoryPredefinedState], error) {
	if req.DryRun {
		s := UrlCategoryPredefinedState{UrlCategoryPredefinedArgs: req.Inputs, CategoryId: stringPtr("preview")}
		return infer.CreateResponse[UrlCategoryPredefinedState]{ID: "preview", Output: s}, nil
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[UrlCategoryPredefinedState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	allCategories, err := urlcategories.GetAll(ctx, service, false, false, "ALL")
	if err != nil {
		return infer.CreateResponse[UrlCategoryPredefinedState]{}, fmt.Errorf("failed retrieving URL categories: %w", err)
	}

	existing := resolvePredefinedCategory(allCategories, req.Inputs.Name)
	if existing == nil {
		return infer.CreateResponse[UrlCategoryPredefinedState]{}, fmt.Errorf("predefined URL category %q not found", req.Inputs.Name)
	}

	resolvedID := existing.ID
	desired := urlCategoryPredefinedArgsToAPI(req.Inputs, existing)

	if err := applyPredefinedCategoryDiff(ctx, service, resolvedID, existing, &desired); err != nil {
		return infer.CreateResponse[UrlCategoryPredefinedState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[UrlCategoryPredefinedState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	// Prefer GetAll over Get: the individual Get(id) endpoint may return incomplete data for
	// predefined categories (e.g. empty ConfiguredName). GetAll returns full category data.
	allCategoriesAfter, err := urlcategories.GetAll(ctx, service, false, false, "ALL")
	if err != nil {
		return infer.CreateResponse[UrlCategoryPredefinedState]{}, fmt.Errorf("failed retrieving URL categories after update: %w", err)
	}
	updated := resolvePredefinedCategory(allCategoriesAfter, resolvedID)
	if updated == nil {
		// Fallback to Get if not found in GetAll (e.g. eventual consistency)
		var getErr error
		updated, getErr = urlcategories.Get(ctx, service, resolvedID)
		if getErr != nil {
			return infer.CreateResponse[UrlCategoryPredefinedState]{}, getErr
		}
		// Merge read-only fields from existing when Get returns empty (API quirk)
		if updated.ConfiguredName == "" {
			updated.ConfiguredName = existing.ConfiguredName
		}
		if updated.SuperCategory == "" {
			updated.SuperCategory = existing.SuperCategory
		}
	}
	state := urlCategoryPredefinedToState(updated)
	return infer.CreateResponse[UrlCategoryPredefinedState]{
		ID:     resolvedID,
		Output: state,
	}, nil
}

func (UrlCategoryPredefined) Read(ctx context.Context, req infer.ReadRequest[UrlCategoryPredefinedArgs, UrlCategoryPredefinedState]) (infer.ReadResponse[UrlCategoryPredefinedArgs, UrlCategoryPredefinedState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[UrlCategoryPredefinedArgs, UrlCategoryPredefinedState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	categoryID := req.ID
	if categoryID == "" {
		return infer.ReadResponse[UrlCategoryPredefinedArgs, UrlCategoryPredefinedState]{}, nil
	}

	allCategories, err := urlcategories.GetAll(ctx, service, false, false, "ALL")
	if err != nil {
		return infer.ReadResponse[UrlCategoryPredefinedArgs, UrlCategoryPredefinedState]{}, err
	}

	resp := resolvePredefinedCategory(allCategories, categoryID)
	if resp == nil {
		individualResp, err := urlcategories.Get(ctx, service, categoryID)
		if err != nil {
			return infer.ReadResponse[UrlCategoryPredefinedArgs, UrlCategoryPredefinedState]{}, nil
		}
		if individualResp.CustomCategory {
			return infer.ReadResponse[UrlCategoryPredefinedArgs, UrlCategoryPredefinedState]{}, nil
		}
		resp = individualResp
	}

	state := urlCategoryPredefinedToState(resp)
	return infer.ReadResponse[UrlCategoryPredefinedArgs, UrlCategoryPredefinedState]{
		ID:     resp.ID,
		Inputs: state.UrlCategoryPredefinedArgs,
		State:  state,
	}, nil
}

func (UrlCategoryPredefined) Update(ctx context.Context, req infer.UpdateRequest[UrlCategoryPredefinedArgs, UrlCategoryPredefinedState]) (infer.UpdateResponse[UrlCategoryPredefinedState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[UrlCategoryPredefinedState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	categoryID := req.ID
	if categoryID == "" {
		return infer.UpdateResponse[UrlCategoryPredefinedState]{}, fmt.Errorf("category ID not set")
	}

	currentCategory, err := urlcategories.Get(ctx, service, categoryID)
	if err != nil {
		return infer.UpdateResponse[UrlCategoryPredefinedState]{}, fmt.Errorf("failed to read predefined URL category: %w", err)
	}

	desired := urlCategoryPredefinedArgsToAPI(req.Inputs, currentCategory)

	if err := applyPredefinedCategoryDiff(ctx, service, categoryID, currentCategory, &desired); err != nil {
		return infer.UpdateResponse[UrlCategoryPredefinedState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[UrlCategoryPredefinedState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	// Prefer GetAll over Get: individual Get(id) may return incomplete data for predefined categories.
	allCategories, err := urlcategories.GetAll(ctx, service, false, false, "ALL")
	if err != nil {
		return infer.UpdateResponse[UrlCategoryPredefinedState]{}, fmt.Errorf("failed retrieving URL categories after update: %w", err)
	}
	updated := resolvePredefinedCategory(allCategories, categoryID)
	if updated == nil {
		updated, err = urlcategories.Get(ctx, service, categoryID)
		if err != nil {
			return infer.UpdateResponse[UrlCategoryPredefinedState]{}, err
		}
		if updated.ConfiguredName == "" {
			updated.ConfiguredName = currentCategory.ConfiguredName
		}
		if updated.SuperCategory == "" {
			updated.SuperCategory = currentCategory.SuperCategory
		}
	}
	state := urlCategoryPredefinedToState(updated)
	return infer.UpdateResponse[UrlCategoryPredefinedState]{Output: state}, nil
}

func (UrlCategoryPredefined) Delete(ctx context.Context, req infer.DeleteRequest[UrlCategoryPredefinedState]) (infer.DeleteResponse, error) {
	// No-op: predefined categories cannot be deleted
	return infer.DeleteResponse{}, nil
}

func (UrlCategoryPredefined) Annotate(a infer.Annotator) {
	describeResource(a, &UrlCategoryPredefined{}, `The zia_url_categories_predefined resource manages predefined URL category overrides in the Zscaler Internet Access (ZIA) cloud service. This resource allows administrators to add custom URLs, keywords, and IP ranges to existing predefined (built-in) URL categories. Predefined categories cannot be deleted; the delete operation is a no-op.

For more information, see the [ZIA URL Categories documentation](https://help.zscaler.com/zia/url-categories).

{{% examples %}}
## Example Usage

{{% example %}}
### Override a Predefined URL Category

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.UrlCategoryPredefined("example", {
    name: "FINANCE",
    urls: [
        "finance.example.com",
        "banking.example.com",
    ],
    keywords: ["financial-portal"],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.UrlCategoryPredefined("example",
    name="FINANCE",
    urls=[
        "finance.example.com",
        "banking.example.com",
    ],
    keywords=["financial-portal"],
)
`+tripleBacktick("")+`

`+tripleBacktick("go")+`
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	zia "github.com/zscaler/pulumi-zia/sdk/go/pulumi-zia"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := zia.NewUrlCategoryPredefined(ctx, "example", &zia.UrlCategoryPredefinedArgs{
			Name:     pulumi.String("FINANCE"),
			Urls:     pulumi.ToStringArray([]string{"finance.example.com", "banking.example.com"}),
			Keywords: pulumi.ToStringArray([]string{"financial-portal"}),
		})
		return err
	})
}
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:UrlCategoryPredefined
    properties:
      name: FINANCE
      urls:
        - finance.example.com
        - banking.example.com
      keywords:
        - financial-portal
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing Predefined URL Category override can be imported using its category ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:UrlCategoryPredefined example FINANCE
`+tripleBacktick("")+`
`)
}

func (a *UrlCategoryPredefinedArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name or ID of the predefined URL category to override (e.g., `FINANCE`, `SOCIAL_NETWORKING`).")
	ann.Describe(&a.Urls, "List of custom URLs to add to the predefined category.")
	ann.Describe(&a.Keywords, "List of custom keywords to add to the predefined category.")
	ann.Describe(&a.KeywordsRetainingParentCategory, "List of keywords that retain their parent category classification.")
	ann.Describe(&a.IpRanges, "List of custom IP address ranges to add to the predefined category.")
	ann.Describe(&a.IpRangesRetainingParentCategory, "List of IP ranges that retain their parent category classification.")
}

func (s *UrlCategoryPredefinedState) Annotate(a infer.Annotator) {
	a.Describe(&s.CategoryId, "The system-generated ID of the predefined URL category.")
	a.Describe(&s.ConfiguredName, "The configured display name of the predefined URL category.")
	a.Describe(&s.SuperCategory, "The super category of the predefined URL category.")
	a.Describe(&s.UrlType, "The URL type of the predefined category.")
	a.Describe(&s.Type, "The type of the URL category.")
	a.Describe(&s.Val, "The internal numeric value of the URL category.")
	a.Describe(&s.Editable, "Whether the predefined URL category is editable.")
	a.Describe(&s.DbCategorizedUrls, "URLs in this category that have been categorized by the Zscaler database.")
	a.Describe(&s.CustomUrlsCount, "The number of custom URLs in the category.")
	a.Describe(&s.UrlsRetainingParentCategoryCount, "The number of URLs retaining parent category.")
	a.Describe(&s.CustomIpRangesCount, "The number of custom IP ranges in the category.")
	a.Describe(&s.IpRangesRetainingParentCategoryCount, "The number of IP ranges retaining parent category.")
}

func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	aMap := stringSliceToMap(a)
	bMap := stringSliceToMap(b)
	for k := range aMap {
		if !bMap[k] {
			return false
		}
	}
	return true
}

func computeSliceDiff(current, desired []string) (toAdd, toRemove []string) {
	currentMap := stringSliceToMap(current)
	desiredMap := stringSliceToMap(desired)
	for item := range desiredMap {
		if !currentMap[item] {
			toAdd = append(toAdd, item)
		}
	}
	for item := range currentMap {
		if !desiredMap[item] {
			toRemove = append(toRemove, item)
		}
	}
	return
}

func buildImmutableCategory(existing *urlcategories.URLCategory) urlcategories.URLCategory {
	return urlcategories.URLCategory{
		ID:             existing.ID,
		ConfiguredName: existing.ConfiguredName,
		SuperCategory:  existing.SuperCategory,
		UrlType:        existing.UrlType,
		Type:           existing.Type,
		Val:            existing.Val,
		CustomCategory: existing.CustomCategory,
		Editable:       existing.Editable,
	}
}

func applyPredefinedCategoryDiff(ctx context.Context, service *zscaler.Service, categoryID string, current *urlcategories.URLCategory, desired *urlcategories.URLCategory) error {
	urlsToAdd, urlsToRemove := computeSliceDiff(current.Urls, desired.Urls)
	ipToAdd, ipToRemove := computeSliceDiff(current.IPRanges, desired.IPRanges)
	ipRetToAdd, ipRetToRemove := computeSliceDiff(current.IPRangesRetainingParentCategory, desired.IPRangesRetainingParentCategory)

	incrementalRemoves := len(urlsToRemove) + len(ipToRemove) + len(ipRetToRemove)
	incrementalAdds := len(urlsToAdd) + len(ipToAdd) + len(ipRetToAdd)
	keywordsChanged := !stringSlicesEqual(current.Keywords, desired.Keywords)
	kwRetChanged := !stringSlicesEqual(current.KeywordsRetainingParentCategory, desired.KeywordsRetainingParentCategory)

	if incrementalRemoves > 0 {
		removeCategory := buildImmutableCategory(current)
		removeCategory.Urls = urlsToRemove
		removeCategory.IPRanges = ipToRemove
		removeCategory.IPRangesRetainingParentCategory = ipRetToRemove
		if _, _, err := urlcategories.UpdateURLCategories(ctx, service, categoryID, &removeCategory, "REMOVE_FROM_LIST"); err != nil {
			return fmt.Errorf("failed to remove items from predefined category: %w", err)
		}
	}

	if incrementalAdds > 0 {
		addCategory := buildImmutableCategory(current)
		addCategory.Urls = urlsToAdd
		addCategory.IPRanges = ipToAdd
		addCategory.IPRangesRetainingParentCategory = ipRetToAdd
		if _, _, err := urlcategories.UpdateURLCategories(ctx, service, categoryID, &addCategory, "ADD_TO_LIST"); err != nil {
			return fmt.Errorf("failed to add items to predefined category: %w", err)
		}
	}

	if keywordsChanged || kwRetChanged {
		kwCategory := buildImmutableCategory(current)
		kwCategory.Keywords = desired.Keywords
		kwCategory.KeywordsRetainingParentCategory = desired.KeywordsRetainingParentCategory
		if _, _, err := urlcategories.UpdateURLCategories(ctx, service, categoryID, &kwCategory, ""); err != nil {
			return fmt.Errorf("failed to update keywords for predefined category: %w", err)
		}
	}

	return nil
}

func urlCategoryPredefinedArgsToAPI(args UrlCategoryPredefinedArgs, existing *urlcategories.URLCategory) urlcategories.URLCategory {
	return urlcategories.URLCategory{
		ID:                              existing.ID,
		ConfiguredName:                  existing.ConfiguredName,
		SuperCategory:                   existing.SuperCategory,
		UrlType:                         existing.UrlType,
		Type:                            existing.Type,
		Val:                             existing.Val,
		CustomCategory:                  existing.CustomCategory,
		Editable:                        existing.Editable,
		Urls:                            args.Urls,
		Keywords:                        args.Keywords,
		KeywordsRetainingParentCategory: args.KeywordsRetainingParentCategory,
		IPRanges:                        args.IpRanges,
		IPRangesRetainingParentCategory: args.IpRangesRetainingParentCategory,
	}
}

func urlCategoryPredefinedToState(r *urlcategories.URLCategory) UrlCategoryPredefinedState {
	configuredName := r.ConfiguredName
	if configuredName == "" && r.ID != "" {
		// Get-by-ID may return empty ConfiguredName for predefined categories; ID often matches (e.g. "FINANCE")
		configuredName = r.ID
	}
	return UrlCategoryPredefinedState{
		UrlCategoryPredefinedArgs: UrlCategoryPredefinedArgs{
			Name:                           r.ID,
			Urls:                           r.Urls,
			Keywords:                       r.Keywords,
			KeywordsRetainingParentCategory: r.KeywordsRetainingParentCategory,
			IpRanges:                       r.IPRanges,
			IpRangesRetainingParentCategory: r.IPRangesRetainingParentCategory,
		},
		CategoryId:                         stringPtr(r.ID),
		ConfiguredName:                     stringPtr(configuredName),
		SuperCategory:                      stringPtr(r.SuperCategory),
		UrlType:                            stringPtr(r.UrlType),
		Type:                               stringPtr(r.Type),
		Val:                                intPtr(r.Val),
		Editable:                           boolPtr(r.Editable),
		DbCategorizedUrls:                  r.DBCategorizedUrls,
		CustomUrlsCount:                    intPtr(r.CustomUrlsCount),
		UrlsRetainingParentCategoryCount:   intPtr(r.UrlsRetainingParentCategoryCount),
		CustomIpRangesCount:                intPtr(r.CustomIpRangesCount),
		IpRangesRetainingParentCategoryCount: intPtr(r.IPRangesRetainingParentCategoryCount),
	}
}

var _ infer.CustomResource[UrlCategoryPredefinedArgs, UrlCategoryPredefinedState] = UrlCategoryPredefined{}
