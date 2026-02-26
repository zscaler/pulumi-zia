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

// Package provider implements utility functions for the ZIA native provider.
// Adopted from terraform-provider-zia utils.go.

package provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/pulumi-zia/provider/internal/zia"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/activation"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/common"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/firewallpolicies/filteringrules"
)

// ConvertRFC1123ToEpoch parses an RFC1123 time string and returns Unix epoch seconds.
func ConvertRFC1123ToEpoch(timeStr string) (int, error) {
	t, err := time.Parse(time.RFC1123, timeStr)
	if err != nil {
		return 0, fmt.Errorf("invalid time format: %v. Expected format: RFC1123 (Mon, 02 Jan 2006 15:04:05 MST)", err)
	}
	return int(t.Unix()), nil
}

// ExclusionTimeUTCLayout is the layout for subcloud exclusion times (UTC): "MM/DD/YYYY HH:MM:SS am/pm"
const ExclusionTimeUTCLayout = "01/02/2006 03:04:05 pm"

// ParseExclusionTimeUTC parses a string in "MM/DD/YYYY HH:MM:SS am/pm" (UTC) and returns Unix epoch seconds.
func ParseExclusionTimeUTC(timeStr string) (int, error) {
	s := strings.TrimSpace(timeStr)
	if s == "" {
		return 0, fmt.Errorf("exclusion time string is empty")
	}
	t, err := time.ParseInLocation(ExclusionTimeUTCLayout, s, time.UTC)
	if err != nil {
		return 0, fmt.Errorf("invalid time format: %v. Expected format: MM/DD/YYYY HH:MM:SS am/pm (e.g. 02/19/2026 11:59:00 pm)", err)
	}
	return int(t.Unix()), nil
}

// FormatExclusionTimeUTC formats Unix epoch seconds as "MM/DD/YYYY HH:MM:SS am/pm" in UTC.
func FormatExclusionTimeUTC(epoch int) string {
	return time.Unix(int64(epoch), 0).UTC().Format(ExclusionTimeUTCLayout)
}

// ResolveExclusionTimeEpoch returns Unix epoch from either the UTC string or the epoch int.
func ResolveExclusionTimeEpoch(hasEpoch bool, epochVal int, utcStr string) (int, error) {
	s := strings.TrimSpace(utcStr)
	if s != "" {
		return ParseExclusionTimeUTC(s)
	}
	if hasEpoch {
		return epochVal, nil
	}
	return 0, fmt.Errorf("either epoch or UTC string must be set")
}

// convertAndValidateSizeQuota validates size_quota in MB (10-100000) and returns KB.
func convertAndValidateSizeQuota(sizeQuotaMB int) (int, error) {
	const (
		minMB = 10
		maxMB = 100000
	)
	if sizeQuotaMB < minMB || sizeQuotaMB > maxMB {
		return 0, fmt.Errorf("size_quota must be between %d MB and %d MB", minMB, maxMB)
	}
	return sizeQuotaMB * 1024, nil
}

// isSingleDigitDay returns true if the time string has a single-digit day (e.g., "2" instead of "02").
func isSingleDigitDay(timeStr string) bool {
	parts := strings.Split(timeStr, " ")
	if len(parts) < 2 {
		return false
	}
	day := parts[1]
	return len(day) == 1
}

// processCountries prefixes 2-letter ISO country codes with "COUNTRY_".
func processCountries(countries []string) []string {
	processed := make([]string, len(countries))
	for i, country := range countries {
		if country != "ANY" && country != "NONE" && len(country) == 2 {
			processed[i] = "COUNTRY_" + country
		} else {
			processed[i] = country
		}
	}
	return processed
}

var failFastErrorCodes = []string{
	"INVALID_INPUT_ARGUMENT",
	"TRIAL_EXPIRED",
	"EDIT_LOCK_NOT_AVAILABLE",
	"DUPLICATE_ITEM",
}

// failFastOnErrorCodes detects known fatal API error codes and returns the original error to fail immediately.
func failFastOnErrorCodes(err error) error {
	if err == nil {
		return nil
	}
	var apiErr *errorx.ErrorResponse
	if errors.As(err, &apiErr) {
		code := extractErrorCodeFromBody(apiErr.Message)
		for _, c := range failFastErrorCodes {
			if code == c {
				log.Printf("[ERROR] Failing immediately due to API error code '%s': %s", c, apiErr.Message)
				return err
			}
		}
	}
	errMsg := err.Error()
	for _, code := range failFastErrorCodes {
		match := fmt.Sprintf(`"code":"%s"`, code)
		if strings.Contains(errMsg, match) {
			log.Printf("[WARN] Failing due to fallback match for code '%s': %s", code, errMsg)
			return err
		}
	}
	return nil
}

func extractErrorCodeFromBody(body string) string {
	var parsed struct {
		Code string `json:"code"`
	}
	if err := json.Unmarshal([]byte(body), &parsed); err == nil {
		return parsed.Code
	}
	return ""
}

// triggerActivation triggers ZIA configuration activation.
func triggerActivation(ctx context.Context, client *zia.Client) error {
	req := activation.Activation{Status: "ACTIVE"}
	log.Printf("[INFO] Triggering configuration activation\n%+v\n", req)
	_, err := activation.CreateActivation(ctx, client.Service, req)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Configuration activation triggered successfully.")
	return nil
}

// ptrToString returns the string value or "" if nil.
func ptrToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// ptrToStringDefault returns the string value or d if nil/empty.
func ptrToStringDefault(s *string, d string) string {
	if s == nil || *s == "" {
		return d
	}
	return *s
}

// stringPtr returns a pointer to s, or nil if s is empty.
func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// ptrOrDefault returns ptr if non-nil, otherwise a pointer to def.
func ptrOrDefault(ptr *string, def string) *string {
	if ptr != nil {
		return ptr
	}
	return stringPtr(def)
}

// intPtr returns a pointer to i.
func intPtr(i int) *int {
	return &i
}

// ptrToBool returns the bool value or false if nil.
func ptrToBool(p *bool) bool {
	if p == nil {
		return false
	}
	return *p
}

// boolPtr returns a pointer to b.
func boolPtr(b bool) *bool {
	return &b
}

// omitFalseBool returns nil when b is false (the zero value), avoiding
// spurious diffs for optional Computed fields the user did not set.
func omitFalseBool(b bool) *bool {
	if !b {
		return nil
	}
	return &b
}

// omitZeroInt returns nil when i is 0, avoiding spurious diffs for
// optional Computed int fields the user did not set.
func omitZeroInt(i int) *int {
	if i == 0 {
		return nil
	}
	return &i
}

// nilIfEmpty returns nil when the slice is empty, avoiding spurious
// diffs between nil (user omitted) and [] (API returned empty).
func nilIfEmpty(s []string) []string {
	if len(s) == 0 {
		return nil
	}
	return s
}

// processCountriesFromAPI strips the "COUNTRY_" prefix from API country codes
// and returns nil if the resulting slice is empty.
func processCountriesFromAPI(countries []string) []string {
	if len(countries) == 0 {
		return nil
	}
	result := make([]string, len(countries))
	for i, c := range countries {
		result[i] = strings.TrimPrefix(c, "COUNTRY_")
	}
	return result
}

// idsFromIDNameExtensions extracts IDs from a slice of common.IDNameExtensions.
func idsFromIDNameExtensions(list []common.IDNameExtensions) []int {
	out := make([]int, 0, len(list))
	for _, v := range list {
		if v.ID != 0 || v.Name != "" {
			out = append(out, v.ID)
		}
	}
	return out
}

// detachUserFromFilteringRules detaches a user ID from all firewall filtering rules before delete.
// Adopted from terraform-provider-zia utils.go DetachRuleIDNameExtensions.
func detachUserFromFilteringRules(ctx context.Context, client *zia.Client, userID int) error {
	service := client.Service
	log.Printf("[INFO] Detaching user %d from firewall filtering rules", userID)
	rules, err := filteringrules.GetAll(ctx, service, nil)
	if err != nil {
		return err
	}
	for _, rule := range rules {
		var ids []common.IDNameExtensions
		shouldUpdate := false
		for _, u := range rule.Users {
			if u.ID != userID {
				ids = append(ids, u)
			} else {
				shouldUpdate = true
			}
		}
		if shouldUpdate {
			rule.Users = ids
			time.Sleep(5 * time.Second)
			if _, err := filteringrules.Get(ctx, service, rule.ID); err == nil {
				if _, err := filteringrules.Update(ctx, service, rule.ID, &rule); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// detachFromFilteringRules detaches an ID from a specific field in all firewall filtering rules before delete.
// Adopted from terraform-provider-zia utils.go DetachRuleIDNameExtensions.
func detachFromFilteringRules(ctx context.Context, client *zia.Client, id int, resourceName string,
	getter func(*filteringrules.FirewallFilteringRules) []common.IDNameExtensions,
	setter func(*filteringrules.FirewallFilteringRules, []common.IDNameExtensions)) error {
	service := client.Service
	log.Printf("[INFO] Detaching %s %d from firewall filtering rules", resourceName, id)
	rules, err := filteringrules.GetAll(ctx, service, nil)
	if err != nil {
		return err
	}
	for _, rule := range rules {
		var ids []common.IDNameExtensions
		shouldUpdate := false
		for _, item := range getter(&rule) {
			if item.ID != id {
				ids = append(ids, item)
			} else {
				shouldUpdate = true
			}
		}
		if shouldUpdate {
			setter(&rule, ids)
			time.Sleep(5 * time.Second)
			if _, err := filteringrules.Get(ctx, service, rule.ID); err == nil {
				if _, err := filteringrules.Update(ctx, service, rule.ID, &rule); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// tripleBacktick returns a fenced code block delimiter with an optional language tag.
// Used by Annotate() methods to embed code examples in resource descriptions.
func tripleBacktick(lang string) string {
	return "```" + lang
}

// describeResource safely sets the resource-level description on an Annotator.
// This wraps the call in a recover because the GetToken code path in
// pulumi-go-provider creates a non-addressable FieldMatcher value, causing
// a.Describe(&T{}, ...) to panic. The schema generation path (getAnnotated)
// handles this correctly by wrapping in reflect.New, so the description IS
// set for schema output despite the recovered panic during GetToken.
func describeResource(a infer.Annotator, resource any, description string) {
	defer func() { recover() }() //nolint:errcheck
	a.Describe(resource, description)
}

// shouldActivate returns true if ZIA_ACTIVATION env var is set to a truthy value.
func shouldActivate() bool {
	activationEnv, exists := os.LookupEnv("ZIA_ACTIVATION")
	if !exists {
		return false
	}
	activationBool, err := strconv.ParseBool(activationEnv)
	if err != nil {
		log.Printf("[WARN] Error parsing ZIA_ACTIVATION env var as bool: %v", err)
		return false
	}
	return activationBool
}
