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

// Package provider implements the DC Exclusion resource.
// Adopted from terraform-provider-zia resource_zia_dc_exclusions.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/trafficforwarding/dc_exclusions"
)

// DcExclusionArgs are the inputs.
type DcExclusionArgs struct {
	DatacenterId *int    `pulumi:"datacenterId,optional"`
	StartTime    *int    `pulumi:"startTime,optional"`
	StartTimeUtc *string `pulumi:"startTimeUtc,optional"`
	EndTime      *int    `pulumi:"endTime,optional"`
	EndTimeUtc   *string `pulumi:"endTimeUtc,optional"`
	Description  *string `pulumi:"description,optional"`
}

// DcExclusionState is the persisted state.
type DcExclusionState struct {
	DcExclusionArgs
	Expired *bool `pulumi:"expired"`
}

// DcExclusion implements the zia:index:DcExclusion resource.
type DcExclusion struct{}

func resolveExclusionTime(hasEpoch bool, epochVal int, utcStr string) (int, error) {
	return ResolveExclusionTimeEpoch(hasEpoch, epochVal, utcStr)
}

func dcExclusionToAPI(args DcExclusionArgs, dcID int) (dc_exclusions.DCExclusions, error) {
	hasStartEpoch := args.StartTime != nil && *args.StartTime != 0
	startEpoch := 0
	if args.StartTime != nil {
		startEpoch = *args.StartTime
	}
	startUtc := ""
	if args.StartTimeUtc != nil {
		startUtc = *args.StartTimeUtc
	}
	startTime, err := resolveExclusionTime(hasStartEpoch, startEpoch, startUtc)
	if err != nil {
		return dc_exclusions.DCExclusions{}, fmt.Errorf("start time: %w", err)
	}

	hasEndEpoch := args.EndTime != nil && *args.EndTime != 0
	endEpoch := 0
	if args.EndTime != nil {
		endEpoch = *args.EndTime
	}
	endUtc := ""
	if args.EndTimeUtc != nil {
		endUtc = *args.EndTimeUtc
	}
	endTime, err := resolveExclusionTime(hasEndEpoch, endEpoch, endUtc)
	if err != nil {
		return dc_exclusions.DCExclusions{}, fmt.Errorf("end time: %w", err)
	}

	return dc_exclusions.DCExclusions{
		DcID:        dcID,
		StartTime:   startTime,
		EndTime:     endTime,
		Description: ptrToString(args.Description),
	}, nil
}

func (DcExclusion) Create(ctx context.Context, req infer.CreateRequest[DcExclusionArgs]) (infer.CreateResponse[DcExclusionState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[DcExclusionState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	dcID := 0
	if req.Inputs.DatacenterId != nil {
		dcID = *req.Inputs.DatacenterId
	}
	if dcID == 0 {
		return infer.CreateResponse[DcExclusionState]{}, fmt.Errorf("datacenterId is required on create")
	}

	apiReq, err := dcExclusionToAPI(req.Inputs, dcID)
	if err != nil {
		return infer.CreateResponse[DcExclusionState]{}, err
	}

	resp, _, err := dc_exclusions.Create(ctx, service, &apiReq)
	if err != nil {
		return infer.CreateResponse[DcExclusionState]{}, err
	}
	log.Printf("[INFO] Created ZIA DC exclusion. DcID: %v", resp.DcID)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[DcExclusionState]{}, activationErr
		}
	}

	state := DcExclusionState{
		DcExclusionArgs: req.Inputs,
		Expired:         boolPtr(resp.Expired),
	}
	// Set computed start/end from API response
	state.StartTime = intPtr(resp.StartTime)
	state.EndTime = intPtr(resp.EndTime)
	utcStart := FormatExclusionTimeUTC(resp.StartTime)
	utcEnd := FormatExclusionTimeUTC(resp.EndTime)
	state.StartTimeUtc = &utcStart
	state.EndTimeUtc = &utcEnd
	return infer.CreateResponse[DcExclusionState]{
		ID:     strconv.Itoa(resp.DcID),
		Output: state,
	}, nil
}

func (DcExclusion) Read(ctx context.Context, req infer.ReadRequest[DcExclusionArgs, DcExclusionState]) (infer.ReadResponse[DcExclusionArgs, DcExclusionState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[DcExclusionArgs, DcExclusionState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	dcID := 0
	if req.State.DatacenterId != nil {
		dcID = *req.State.DatacenterId
	}
	if dcID == 0 && req.ID != "" {
		parsed, err := strconv.Atoi(req.ID)
		if err == nil {
			dcID = parsed
		}
	}
	if dcID == 0 {
		return infer.ReadResponse[DcExclusionArgs, DcExclusionState]{}, fmt.Errorf("no DC exclusion id in state")
	}

	all, err := dc_exclusions.GetAll(ctx, service)
	if err != nil {
		return infer.ReadResponse[DcExclusionArgs, DcExclusionState]{}, err
	}

	var found *dc_exclusions.DCExclusions
	for i := range all {
		if all[i].DcID == dcID {
			found = &all[i]
			break
		}
	}
	if found == nil {
		return infer.ReadResponse[DcExclusionArgs, DcExclusionState]{ID: ""}, nil
	}

	utcStart := FormatExclusionTimeUTC(found.StartTime)
	utcEnd := FormatExclusionTimeUTC(found.EndTime)
	args := DcExclusionArgs{
		DatacenterId: &found.DcID,
		StartTime:    intPtr(found.StartTime),
		StartTimeUtc: &utcStart,
		EndTime:      intPtr(found.EndTime),
		EndTimeUtc:   &utcEnd,
		Description:  stringPtr(found.Description),
	}
	state := DcExclusionState{
		DcExclusionArgs: args,
		Expired:         boolPtr(found.Expired),
	}
	return infer.ReadResponse[DcExclusionArgs, DcExclusionState]{
		ID:     strconv.Itoa(found.DcID),
		Inputs: args,
		State:  state,
	}, nil
}

func (DcExclusion) Update(ctx context.Context, req infer.UpdateRequest[DcExclusionArgs, DcExclusionState]) (infer.UpdateResponse[DcExclusionState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[DcExclusionState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	dcID := 0
	if req.State.DatacenterId != nil {
		dcID = *req.State.DatacenterId
	}
	if dcID == 0 {
		return infer.UpdateResponse[DcExclusionState]{}, fmt.Errorf("no DC exclusion id in state")
	}

	all, err := dc_exclusions.GetAll(ctx, service)
	if err != nil {
		return infer.UpdateResponse[DcExclusionState]{}, err
	}
	var existing bool
	for i := range all {
		if all[i].DcID == dcID {
			existing = true
			break
		}
	}
	if !existing {
		return infer.UpdateResponse[DcExclusionState]{}, nil
	}

	apiReq, err := dcExclusionToAPI(req.Inputs, dcID)
	if err != nil {
		return infer.UpdateResponse[DcExclusionState]{}, err
	}
	if _, _, err := dc_exclusions.Update(ctx, service, &apiReq); err != nil {
		return infer.UpdateResponse[DcExclusionState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[DcExclusionState]{}, activationErr
		}
	}

	// Re-read to get updated state
	all, _ = dc_exclusions.GetAll(ctx, service)
	var updated *dc_exclusions.DCExclusions
	for i := range all {
		if all[i].DcID == dcID {
			updated = &all[i]
			break
		}
	}
	state := DcExclusionState{DcExclusionArgs: req.Inputs}
	if updated != nil {
		state.DatacenterId = &updated.DcID
		state.StartTime = intPtr(updated.StartTime)
		state.EndTime = intPtr(updated.EndTime)
		utcStart := FormatExclusionTimeUTC(updated.StartTime)
		utcEnd := FormatExclusionTimeUTC(updated.EndTime)
		state.StartTimeUtc = &utcStart
		state.EndTimeUtc = &utcEnd
		state.Expired = boolPtr(updated.Expired)
	}
	return infer.UpdateResponse[DcExclusionState]{Output: state}, nil
}

func (DcExclusion) Delete(ctx context.Context, req infer.DeleteRequest[DcExclusionState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	dcID := 0
	if req.State.DatacenterId != nil {
		dcID = *req.State.DatacenterId
	}
	if dcID != 0 {
		if _, err := dc_exclusions.Delete(ctx, service, dcID); err != nil {
			if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
				return infer.DeleteResponse{}, nil
			}
			return infer.DeleteResponse{}, err
		}
		log.Printf("[INFO] ZIA DC exclusion deleted")
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}

	return infer.DeleteResponse{}, nil
}

func (DcExclusion) Annotate(a infer.Annotator) {
	describeResource(a, &DcExclusion{}, `The zia_dc_exclusion resource manages data center exclusions in the Zscaler Internet Access (ZIA) cloud service. Data center exclusions allow you to temporarily exclude specific data centers from traffic forwarding during maintenance windows or other planned events.

{{% examples %}}
## Example Usage

{{% example %}}
### Data Center Exclusion

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.DcExclusion("example", {
    datacenterId: 12345,
    startTimeUtc: "02/25/2026 08:00:00 am",
    endTimeUtc: "02/25/2026 10:00:00 pm",
    description: "Maintenance window exclusion",
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.DcExclusion("example",
    datacenter_id=12345,
    start_time_utc="02/25/2026 08:00:00 am",
    end_time_utc="02/25/2026 10:00:00 pm",
    description="Maintenance window exclusion",
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:DcExclusion
    properties:
      datacenterId: 12345
      startTimeUtc: "02/25/2026 08:00:00 am"
      endTimeUtc: "02/25/2026 10:00:00 pm"
      description: Maintenance window exclusion
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing Data Center Exclusion can be imported using its datacenter ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:DcExclusion example 12345
`+tripleBacktick("")+`
`)
}

func (a *DcExclusionArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.DatacenterId, "The ID of the data center to exclude.")
	ann.Describe(&a.StartTime, "Exclusion start time as Unix epoch seconds.")
	ann.Describe(&a.StartTimeUtc, "Exclusion start time in UTC format: `MM/DD/YYYY HH:MM:SS am/pm`.")
	ann.Describe(&a.EndTime, "Exclusion end time as Unix epoch seconds.")
	ann.Describe(&a.EndTimeUtc, "Exclusion end time in UTC format: `MM/DD/YYYY HH:MM:SS am/pm`.")
	ann.Describe(&a.Description, "A description of the data center exclusion.")
}

func (s *DcExclusionState) Annotate(a infer.Annotator) {
	a.Describe(&s.Expired, "Whether the data center exclusion has expired.")
}

var _ infer.CustomResource[DcExclusionArgs, DcExclusionState] = DcExclusion{}
