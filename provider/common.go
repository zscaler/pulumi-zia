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

// Package provider implements shared logic for the ZIA native provider.
// Adopted exactly from terraform-provider-zia common.go (reorder, expand/flatten helpers).

package provider

import (
	"log"
	"sort"
	"sync"
	"time"

	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/common"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/urlfilteringpolicies"
)

// --- Reorder logic (from terraform-provider-zia common.go) ---

var (
	urlFilteringSem           = make(chan struct{}, 1)
	urlFilteringStartingOrder int
	urlFilteringOrderMu       sync.Mutex

	trafficCaptureSem           = make(chan struct{}, 1)
	trafficCaptureStartingOrder int
	trafficCaptureOrderMu       sync.Mutex

	sslInspectionSem           = make(chan struct{}, 1)
	sslInspectionStartingOrder int
	sslInspectionOrderMu       sync.Mutex

	sandboxSem           = make(chan struct{}, 1)
	sandboxStartingOrder int
	sandboxOrderMu       sync.Mutex

	natControlRuleSem           = make(chan struct{}, 1)
	natControlRuleStartingOrder int
	natControlRuleOrderMu       sync.Mutex

	forwardingControlSem           = make(chan struct{}, 1)
	forwardingControlStartingOrder int
	forwardingControlOrderMu       sync.Mutex

	firewallIPSSem           = make(chan struct{}, 1)
	firewallIPSStartingOrder int
	firewallIPSOrderMu       sync.Mutex

	firewallFilteringSem           = make(chan struct{}, 1)
	firewallFilteringStartingOrder int
	firewallFilteringOrderMu       sync.Mutex

	firewallDNSSem           = make(chan struct{}, 1)
	firewallDNSStartingOrder int
	firewallDNSOrderMu       sync.Mutex

	fileTypeSem           = make(chan struct{}, 1)
	fileTypeStartingOrder int
	fileTypeOrderMu       sync.Mutex

	dlpWebRulesSem      = make(chan struct{}, 1)
	dlpWebStartingOrder int
	dlpWebOrderMu       sync.Mutex

	cloudAppRuleSem           = make(chan struct{}, 1)
	cloudAppRuleStartingOrder int
	cloudAppRuleOrderMu       sync.Mutex

	cloudCasbDlpRuleSem           = make(chan struct{}, 1)
	cloudCasbDlpRuleStartingOrder int
	cloudCasbDlpRuleOrderMu       sync.Mutex

	cloudCasbMalwareRuleSem           = make(chan struct{}, 1)
	cloudCasbMalwareRuleStartingOrder int
	cloudCasbMalwareRuleOrderMu       sync.Mutex

	bandwidthControlSem           = make(chan struct{}, 1)
	bandwidthControlStartingOrder int
	bandwidthControlOrderMu       sync.Mutex
)

// OrderRule holds the intended order and rank for a rule.
type OrderRule struct {
	Order int
	Rank  int
}

type orderWithState struct {
	order OrderRule
	done  bool
}

type listrules struct {
	orders      map[string]map[int]orderWithState
	orderer     map[string]int
	reorderDone map[string]chan struct{}
	sync.Mutex
}

var rules = listrules{
	orders:      make(map[string]map[int]orderWithState),
	reorderDone: make(map[string]chan struct{}),
}

// RuleIDOrderPair pairs a rule ID with its order.
type RuleIDOrderPair struct {
	ID    int
	Order OrderRule
}

// RuleIDOrderPairList is a sortable list of rule ID/order pairs.
type RuleIDOrderPairList []RuleIDOrderPair

func (p RuleIDOrderPairList) Len() int { return len(p) }
func (p RuleIDOrderPairList) Less(i, j int) bool {
	if p[i].Order == p[j].Order {
		return p[i].ID < p[j].ID
	}
	return p[i].Order.Rank < p[j].Order.Rank || p[i].Order.Rank == p[j].Order.Rank && p[i].Order.Order < p[j].Order.Order
}
func (p RuleIDOrderPairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

var reorderTickInterval = 30 * time.Second

func sortOrders(ruleOrderMap map[int]orderWithState) RuleIDOrderPairList {
	pl := make(RuleIDOrderPairList, len(ruleOrderMap))
	i := 0
	for k, v := range ruleOrderMap {
		pl[i] = RuleIDOrderPair{k, v.order}
		i++
	}
	sort.Sort(pl)
	return pl
}

func reorderAll(resourceType string, getCount func() (int, error), updateOrder func(id int, order OrderRule) error, beforeReorder func()) {
	ticker := time.NewTicker(reorderTickInterval)
	defer ticker.Stop()
	lastSeenSize := 0
	sizeStableTicks := 0
	for {
		select {
		case <-ticker.C:
			rules.Lock()
			size := len(rules.orders[resourceType])
			allDone := true
			for _, v := range rules.orders[resourceType] {
				if !v.done {
					allDone = false
					break
				}
			}

			if allDone && size > 0 {
				if size != lastSeenSize {
					log.Printf("[INFO] reorder: %s rule count changed from %d to %d, waiting for stability", resourceType, lastSeenSize, size)
					lastSeenSize = size
					sizeStableTicks = 0
				} else {
					sizeStableTicks++
					if sizeStableTicks == 1 {
						sorted := sortOrders(rules.orders[resourceType])
						rules.Unlock()

						count, _ := getCount()
						log.Printf("[INFO] reorder: performing reorder for %s (%d rules, api count=%d), sorted: %v", resourceType, size, count, sorted)
						if beforeReorder != nil {
							beforeReorder()
						}
						for _, v := range sorted {
							if v.Order.Order <= count {
								if err := updateOrder(v.ID, v.Order); err != nil {
									log.Printf("[ERROR] couldn't reorder rule %d for %s: %v", v.ID, resourceType, err)
								}
							}
						}

						rules.Lock()
					}
					log.Printf("[INFO] reorder stable tick %d/3 for %s (%d rules)", sizeStableTicks, resourceType, size)
				}

				if sizeStableTicks >= 3 {
					log.Printf("[INFO] reorder complete for %s: %d rules, stable for 3 ticks", resourceType, size)
					rules.Unlock()
					return
				}
			}
			rules.Unlock()
		default:
			time.Sleep(reorderTickInterval / 2)
		}
	}
}

// markOrderRuleAsDone marks a rule as done in the reorder state.
func markOrderRuleAsDone(id int, resourceType string) {
	rules.Lock()
	r := rules.orders[resourceType][id]
	r.done = true
	rules.orders[resourceType][id] = r
	rules.Unlock()
}

// reorderWithBeforeReorder registers a rule for reordering and optionally starts a reorder cycle.
func reorderWithBeforeReorder(order OrderRule, id int, resourceType string, getCount func() (int, error), updateOrder func(id int, order OrderRule) error, beforeReorder func()) {
	rules.Lock()
	shouldCallReorder := false
	if rules.orderer == nil {
		rules.orderer = map[string]int{}
		rules.reorderDone = map[string]chan struct{}{}
	}
	if rules.orders == nil {
		rules.orders = map[string]map[int]orderWithState{}
	}
	if _, ok := rules.orderer[resourceType]; ok {
		select {
		case <-rules.reorderDone[resourceType]:
			log.Printf("[INFO] previous reorder for %s completed, starting new cycle for rule:%d", resourceType, id)
			rules.reorderDone[resourceType] = make(chan struct{})
			shouldCallReorder = true
		default:
			shouldCallReorder = false
		}
	} else {
		rules.orderer[resourceType] = id
		shouldCallReorder = true
		rules.reorderDone[resourceType] = make(chan struct{})
	}
	if rules.orders[resourceType] == nil {
		rules.orders[resourceType] = map[int]orderWithState{}
	}
	rules.orders[resourceType][id] = orderWithState{order, false}
	rules.Unlock()
	if shouldCallReorder {
		log.Printf("[INFO] starting to reorder the rules, delegating to rule:%d, order:%d", id, order)
		doneCh := rules.reorderDone[resourceType]
		go func() {
			reorderAll(resourceType, getCount, updateOrder, beforeReorder)
			close(doneCh)
		}()
	}
}

// waitForReorder blocks until the reorder goroutine for the given resource type has completed.
func waitForReorder(resourceType string) {
	rules.Lock()
	ch := rules.reorderDone[resourceType]
	rules.Unlock()
	if ch != nil {
		<-ch
	}
}

// --- Expand/Flatten helpers (from terraform-provider-zia common.go) ---

// idsToIDNameExtensions converts a slice of IDs to []common.IDNameExtensions.
func idsToIDNameExtensions(ids []int) []common.IDNameExtensions {
	if len(ids) == 0 {
		return []common.IDNameExtensions{}
	}
	result := make([]common.IDNameExtensions, len(ids))
	for i, id := range ids {
		result[i] = common.IDNameExtensions{ID: id}
	}
	return result
}

// idsToCommonNSS converts a slice of IDs to []common.CommonNSS (used by cloudnss NSSFeed).
func idsToCommonNSS(ids []int) []common.CommonNSS {
	if len(ids) == 0 {
		return nil
	}
	result := make([]common.CommonNSS, len(ids))
	for i, id := range ids {
		result[i] = common.CommonNSS{ID: id}
	}
	return result
}

// commonNSSToIDs extracts IDs from []common.CommonNSS.
func commonNSSToIDs(list []common.CommonNSS) []int {
	if len(list) == 0 {
		return nil
	}
	ids := make([]int, 0, len(list))
	for _, item := range list {
		if item.ID != 0 || item.Name != "" {
			ids = append(ids, item.ID)
		}
	}
	return ids
}

// idToOptionalIDName converts an ID to *common.IDName. Returns nil if id is 0.
func idToOptionalIDName(id int) *common.IDName {
	if id == 0 {
		return nil
	}
	return &common.IDName{ID: id}
}

// idNameToOptionalID extracts ID from *common.IDName. Returns nil if api is nil or ID is 0.
func idNameToOptionalID(api *common.IDName) *int {
	if api == nil || api.ID == 0 {
		return nil
	}
	return intPtr(api.ID)
}

// idNameExtensionsToIDs extracts IDs from []common.IDNameExtensions.
func idNameExtensionsToIDs(list []common.IDNameExtensions) []int {
	if len(list) == 0 {
		return nil
	}
	ids := make([]int, 0, len(list))
	for _, item := range list {
		if item.ID != 0 || item.Name != "" {
			ids = append(ids, item.ID)
		}
	}
	return ids
}

// WorkloadGroupInput is input for workload groups (id+name pairs).
type WorkloadGroupInput struct {
	ID   int     `pulumi:"resourceId"`
	Name *string `pulumi:"name,optional"`
}

// ZPAAppSegmentInput is input for ZPA app segments (name+externalId).
type ZPAAppSegmentInput struct {
	Name       string `pulumi:"name"`
	ExternalID string `pulumi:"externalId"`
}

// expandZPAAppSegments converts []ZPAAppSegmentInput to []common.ZPAAppSegments.
func expandZPAAppSegments(in []ZPAAppSegmentInput) []common.ZPAAppSegments {
	if len(in) == 0 {
		return nil
	}
	result := make([]common.ZPAAppSegments, len(in))
	for i, s := range in {
		result[i] = common.ZPAAppSegments{Name: s.Name, ExternalID: s.ExternalID}
	}
	return result
}

// flattenZPAAppSegments converts []common.ZPAAppSegments to []ZPAAppSegmentInput.
func flattenZPAAppSegments(list []common.ZPAAppSegments) []ZPAAppSegmentInput {
	if len(list) == 0 {
		return nil
	}
	result := make([]ZPAAppSegmentInput, len(list))
	for i, s := range list {
		result[i] = ZPAAppSegmentInput{Name: s.Name, ExternalID: s.ExternalID}
	}
	return result
}

// WorkloadGroupOutput is output for workload groups.
type WorkloadGroupOutput struct {
	ID   int    `pulumi:"resourceId"`
	Name string `pulumi:"name"`
}

// CBIProfileInput is input for the CBI profile (ISOLATE action).
type CBIProfileInput struct {
	ProfileSeq *int    `pulumi:"profileSeq,optional"`
	ID         *string `pulumi:"resourceId,optional"`
	Name       *string `pulumi:"name,optional"`
	URL        *string `pulumi:"url,optional"`
}

// CBIProfileOutput is output for the CBI profile.
type CBIProfileOutput struct {
	ProfileSeq int    `pulumi:"profileSeq"`
	ID         string `pulumi:"resourceId"`
	Name       string `pulumi:"name"`
	URL        string `pulumi:"url"`
}

// flattenWorkloadGroups converts []common.IDName to []WorkloadGroupOutput.
func flattenWorkloadGroups(workloadGroups []common.IDName) []WorkloadGroupOutput {
	if workloadGroups == nil {
		return nil
	}
	result := make([]WorkloadGroupOutput, len(workloadGroups))
	for i, wg := range workloadGroups {
		result[i] = WorkloadGroupOutput{ID: wg.ID, Name: wg.Name}
	}
	return result
}

// workloadGroupOutputsToInputs converts []WorkloadGroupOutput to []WorkloadGroupInput.
func workloadGroupOutputsToInputs(out []WorkloadGroupOutput) []WorkloadGroupInput {
	if len(out) == 0 {
		return nil
	}
	result := make([]WorkloadGroupInput, len(out))
	for i, wg := range out {
		result[i] = WorkloadGroupInput{ID: wg.ID, Name: stringPtr(wg.Name)}
	}
	return result
}

// expandWorkloadGroups converts []WorkloadGroupInput to []common.IDName.
func expandWorkloadGroups(in []WorkloadGroupInput) []common.IDName {
	if len(in) == 0 {
		return []common.IDName{}
	}
	result := make([]common.IDName, len(in))
	for i, wg := range in {
		result[i] = common.IDName{ID: wg.ID, Name: ptrToString(wg.Name)}
	}
	return result
}

// flattenCBIProfileSimple converts *urlfilteringpolicies.CBIProfile to *CBIProfileOutput.
func flattenCBIProfileSimple(cbiProfile *urlfilteringpolicies.CBIProfile) *CBIProfileOutput {
	if cbiProfile == nil || (cbiProfile.ID == "" && cbiProfile.Name == "" && cbiProfile.URL == "") {
		return nil
	}
	return &CBIProfileOutput{
		ProfileSeq: cbiProfile.ProfileSeq,
		ID:         cbiProfile.ID,
		Name:       cbiProfile.Name,
		URL:        cbiProfile.URL,
	}
}

// expandCBIProfile converts *CBIProfileInput to *urlfilteringpolicies.CBIProfile.
func expandCBIProfile(in *CBIProfileInput) *urlfilteringpolicies.CBIProfile {
	if in == nil || (ptrToString(in.ID) == "" && ptrToString(in.Name) == "" && ptrToString(in.URL) == "") {
		return nil
	}
	return &urlfilteringpolicies.CBIProfile{
		ProfileSeq: ptrToIntDefault(in.ProfileSeq, 0),
		ID:         ptrToString(in.ID),
		Name:       ptrToString(in.Name),
		URL:        ptrToString(in.URL),
	}
}

func ptrToIntDefault(p *int, d int) int {
	if p == nil {
		return d
	}
	return *p
}

// idsToIDNameExternalIDs converts []int to []common.IDNameExternalID (for virtual_zen_nodes etc.).
func idsToIDNameExternalIDs(ids []int) []common.IDNameExternalID {
	if len(ids) == 0 {
		return nil
	}
	result := make([]common.IDNameExternalID, len(ids))
	for i, id := range ids {
		result[i] = common.IDNameExternalID{ID: id}
	}
	return result
}

// idNameExternalIDsToIDs extracts IDs from []common.IDNameExternalID.
func idNameExternalIDsToIDs(list []common.IDNameExternalID) []int {
	if len(list) == 0 {
		return nil
	}
	ids := make([]int, len(list))
	for i, item := range list {
		ids[i] = item.ID
	}
	return ids
}
