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

// Package provider implements validation for ZIA resources.
// Adopted exactly from terraform-provider-zia validator.go and resource CustomizeDiff.

package provider

import (
	"errors"
	"fmt"
	"strings"

	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/urlfilteringpolicies"
)

// validateURLFilteringActions validates the URL filtering rule based on action type.
func validateURLFilteringActions(rule urlfilteringpolicies.URLFilteringRule) error {
	switch rule.Action {
	case "ISOLATE":
		if rule.CBIProfile == nil {
			return errors.New("cbi_profile attribute is required when action is ISOLATE")
		}
		if rule.CBIProfile.ID == "" && rule.CBIProfile.Name == "" && rule.CBIProfile.URL == "" {
			return errors.New("cbi_profile attribute is required when action is ISOLATE")
		}
		for _, userAgent := range rule.UserAgentTypes {
			if userAgent == "OTHER" {
				return errors.New("user_agent_types should not contain 'OTHER' when action is ISOLATE. Valid options are: CHROME, FIREFOX, MSIE, MSEDGE, MSCHREDGE, OPERA, SAFARI")
			}
		}
		validProtocols := map[string]bool{"HTTPS_RULE": true, "HTTP_RULE": true}
		for _, protocol := range rule.Protocols {
			if !validProtocols[strings.ToUpper(protocol)] {
				return errors.New("when action is ISOLATE, valid options for protocols are: HTTP and/or HTTPS")
			}
		}

	case "CAUTION":
		validMethods := map[string]bool{"CONNECT": true, "GET": true, "HEAD": true}
		for _, method := range rule.RequestMethods {
			if !validMethods[strings.ToUpper(method)] {
				return errors.New("'CAUTION' action is allowed only for CONNECT/GET/HEAD request methods")
			}
		}
	}
	return nil
}

// validateURLFilteringCheck runs the full CustomizeDiff-style validation for URL filtering rule inputs.
// Called from Check.
func validateURLFilteringCheck(action *string, blockOverride *bool, overrideUsers, overrideGroups []int,
	cbiProfile *CBIProfileInput, userAgentTypes, protocols, requestMethods []string,
	enforceTimeValidity *bool, validityStartTime, validityEndTime, validityTimeZoneID *string,
	browserEunTemplateID *int) error {

	act := ptrToString(action)
	blockOv := blockOverride != nil && *blockOverride

	// Common validation for actions other than BLOCK
	if act != "" && act != "BLOCK" {
		if blockOv {
			return errors.New("block_override can only be set when action is BLOCK")
		}
		if len(overrideUsers) > 0 {
			return errors.New("override_users can only be set when action is BLOCK and block_override is true")
		}
		if len(overrideGroups) > 0 {
			return errors.New("override_groups can only be set when action is BLOCK and block_override is true")
		}
	}

	switch act {
	case "ISOLATE":
		if cbiProfile == nil {
			return errors.New("cbi_profile attribute is required when action is ISOLATE")
		}
		if ptrToString(cbiProfile.ID) == "" && ptrToString(cbiProfile.Name) == "" && ptrToString(cbiProfile.URL) == "" {
			return errors.New("cbi_profile attribute is required when action is ISOLATE")
		}
		for _, ua := range userAgentTypes {
			if ua == "OTHER" {
				return errors.New("user_agent_types should not contain 'OTHER' when action is ISOLATE. Valid options are: CHROME, FIREFOX, MSIE, MSEDGE, MSCHREDGE, OPERA, SAFARI")
			}
		}
		validProtocols := map[string]bool{"HTTPS_RULE": true, "HTTP_RULE": true}
		for _, p := range protocols {
			if !validProtocols[strings.ToUpper(p)] {
				return errors.New("when action is ISOLATE, valid options for protocols are: HTTP and/or HTTPS")
			}
		}

	case "CAUTION":
		if len(requestMethods) == 0 {
			return errors.New("request_methods attribute is required when action is CAUTION")
		}
		validMethods := map[string]bool{"CONNECT": true, "GET": true, "HEAD": true}
		for _, m := range requestMethods {
			if !validMethods[strings.ToUpper(m)] {
				return errors.New("'CAUTION' action is allowed only for CONNECT/GET/HEAD request methods")
			}
		}

	case "BLOCK":
		if !blockOv {
			if len(overrideUsers) > 0 {
				return errors.New("override_users can only be set when block_override is true")
			}
			if len(overrideGroups) > 0 {
				return errors.New("override_groups can only be set when block_override is true")
			}
		}
	}

	// Validate enforce_time_validity, validity_start_time, validity_end_time, validity_time_zone_id
	enforce := enforceTimeValidity != nil && *enforceTimeValidity
	if enforce {
		if validityStartTime == nil || *validityStartTime == "" {
			return errors.New("validity_start_time must be set when enforce_time_validity is true")
		}
		if isSingleDigitDay(*validityStartTime) {
			return errors.New("validity_start_time must have a two-digit day (e.g., 02 instead of 2)")
		}
		if validityEndTime == nil || *validityEndTime == "" {
			return errors.New("validity_end_time must be set when enforce_time_validity is true")
		}
		if isSingleDigitDay(*validityEndTime) {
			return errors.New("validity_end_time must have a two-digit day (e.g., 02 instead of 2)")
		}
		if validityTimeZoneID == nil || *validityTimeZoneID == "" {
			return errors.New("validity_time_zone_id must be set when enforce_time_validity is true")
		}
	} else {
		if validityStartTime != nil && *validityStartTime != "" {
			return errors.New("validity_start_time can only be set when enforce_time_validity is true")
		}
		if validityEndTime != nil && *validityEndTime != "" {
			return errors.New("validity_end_time can only be set when enforce_time_validity is true")
		}
		if validityTimeZoneID != nil && *validityTimeZoneID != "" {
			return errors.New("validity_time_zone_id can only be set when enforce_time_validity is true")
		}
	}

	// Validate browser_eun_template_id
	if browserEunTemplateID != nil && *browserEunTemplateID != 0 {
		if act != "BLOCK" && act != "CAUTION" {
			return fmt.Errorf("browser_eun_template_id can only be set when action is BLOCK or CAUTION")
		}
	}
	return nil
}
