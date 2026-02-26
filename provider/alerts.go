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

// Package provider implements the ZIA Subscription Alerts resource.
// Adopted from terraform-provider-zia resource_zia_alerts.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/alerts"
)

// SubscriptionAlert implements the zia:index:SubscriptionAlert resource.
type SubscriptionAlert struct{}

// SubscriptionAlertArgs are the inputs.
type SubscriptionAlertArgs struct {
	Email            *string  `pulumi:"email,optional"`
	Description      *string  `pulumi:"description,optional"`
	Pt0Severities    []string `pulumi:"pt0Severities,optional"`
	SecureSeverities []string `pulumi:"secureSeverities,optional"`
	ManageSeverities []string `pulumi:"manageSeverities,optional"`
	ComplySeverities []string `pulumi:"complySeverities,optional"`
	SystemSeverities []string `pulumi:"systemSeverities,optional"`
}

// SubscriptionAlertState is the persisted state.
type SubscriptionAlertState struct {
	SubscriptionAlertArgs
	AlertId *int `pulumi:"alertId"`
}

func subscriptionAlertToAPI(args SubscriptionAlertArgs, id *int) alerts.AlertSubscriptions {
	alertID := 0
	if id != nil {
		alertID = *id
	}
	return alerts.AlertSubscriptions{
		ID:               alertID,
		Email:            ptrToString(args.Email),
		Description:      ptrToString(args.Description),
		Pt0Severities:    args.Pt0Severities,
		SecureSeverities: args.SecureSeverities,
		ManageSeverities: args.ManageSeverities,
		ComplySeverities: args.ComplySeverities,
		SystemSeverities: args.SystemSeverities,
	}
}

func (SubscriptionAlert) Create(ctx context.Context, req infer.CreateRequest[SubscriptionAlertArgs]) (infer.CreateResponse[SubscriptionAlertState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[SubscriptionAlertState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := subscriptionAlertToAPI(req.Inputs, nil)
	log.Printf("[INFO] Creating ZIA subscription alert")
	resp, _, err := alerts.Create(ctx, service, &apiReq)
	if err != nil {
		return infer.CreateResponse[SubscriptionAlertState]{}, err
	}
	log.Printf("[INFO] Created ZIA subscription alert. ID: %v", resp.ID)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[SubscriptionAlertState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	state := SubscriptionAlertState{
		SubscriptionAlertArgs: req.Inputs,
		AlertId:               &resp.ID,
	}
	return infer.CreateResponse[SubscriptionAlertState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (SubscriptionAlert) Read(ctx context.Context, req infer.ReadRequest[SubscriptionAlertArgs, SubscriptionAlertState]) (infer.ReadResponse[SubscriptionAlertArgs, SubscriptionAlertState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[SubscriptionAlertArgs, SubscriptionAlertState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	alertID := 0
	if req.State.AlertId != nil {
		alertID = *req.State.AlertId
	}
	if alertID == 0 {
		return infer.ReadResponse[SubscriptionAlertArgs, SubscriptionAlertState]{}, fmt.Errorf("no alert ID in state")
	}

	resp, err := alerts.Get(ctx, service, alertID)
	if err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.ReadResponse[SubscriptionAlertArgs, SubscriptionAlertState]{
				ID: "",
			}, nil
		}
		return infer.ReadResponse[SubscriptionAlertArgs, SubscriptionAlertState]{}, err
	}

	args := SubscriptionAlertArgs{
		Email:            stringPtr(resp.Email),
		Description:      stringPtr(resp.Description),
		Pt0Severities:    resp.Pt0Severities,
		SecureSeverities: resp.SecureSeverities,
		ManageSeverities: resp.ManageSeverities,
		ComplySeverities: resp.ComplySeverities,
		SystemSeverities: resp.SystemSeverities,
	}
	state := SubscriptionAlertState{
		SubscriptionAlertArgs: args,
		AlertId:               &resp.ID,
	}
	return infer.ReadResponse[SubscriptionAlertArgs, SubscriptionAlertState]{
		ID:     strconv.Itoa(resp.ID),
		Inputs: args,
		State:  state,
	}, nil
}

func (SubscriptionAlert) Update(ctx context.Context, req infer.UpdateRequest[SubscriptionAlertArgs, SubscriptionAlertState]) (infer.UpdateResponse[SubscriptionAlertState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[SubscriptionAlertState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	alertID := 0
	if req.State.AlertId != nil {
		alertID = *req.State.AlertId
	}
	if alertID == 0 {
		return infer.UpdateResponse[SubscriptionAlertState]{}, fmt.Errorf("no alert ID in state")
	}

	// Verify resource still exists
	if _, err := alerts.Get(ctx, service, alertID); err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.UpdateResponse[SubscriptionAlertState]{}, nil
		}
		return infer.UpdateResponse[SubscriptionAlertState]{}, err
	}

	apiReq := subscriptionAlertToAPI(req.Inputs, &alertID)
	if _, _, err := alerts.Update(ctx, service, alertID, &apiReq); err != nil {
		return infer.UpdateResponse[SubscriptionAlertState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[SubscriptionAlertState]{}, activationErr
		}
	}

	state := SubscriptionAlertState{
		SubscriptionAlertArgs: req.Inputs,
		AlertId:               &alertID,
	}
	return infer.UpdateResponse[SubscriptionAlertState]{Output: state}, nil
}

func (SubscriptionAlert) Delete(ctx context.Context, req infer.DeleteRequest[SubscriptionAlertState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	alertID := 0
	if req.State.AlertId != nil {
		alertID = *req.State.AlertId
	}
	if alertID != 0 {
		if _, err := alerts.Delete(ctx, service, alertID); err != nil {
			return infer.DeleteResponse{}, err
		}
		log.Printf("[INFO] ZIA subscription alert deleted")
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}

	return infer.DeleteResponse{}, nil
}

func (SubscriptionAlert) Annotate(a infer.Annotator) {
	describeResource(a, &SubscriptionAlert{}, `The zia_subscription_alert resource manages subscription alert configurations in the Zscaler Internet Access (ZIA) cloud service. Subscription alerts notify administrators about various system events with configurable severity levels across different categories including security, management, compliance, and system alerts.

For more information, see the [ZIA Subscription Alerts documentation](https://help.zscaler.com/zia/subscription-alerts).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Subscription Alert

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.SubscriptionAlert("example", {
    email: "admin@example.com",
    description: "Critical security alerts",
    secureSeverities: ["CRITICAL", "HIGH"],
    systemSeverities: ["CRITICAL"],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.SubscriptionAlert("example",
    email="admin@example.com",
    description="Critical security alerts",
    secure_severities=["CRITICAL", "HIGH"],
    system_severities=["CRITICAL"],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:SubscriptionAlert
    properties:
      email: admin@example.com
      description: Critical security alerts
      secureSeverities:
        - CRITICAL
        - HIGH
      systemSeverities:
        - CRITICAL
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing Subscription Alert can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:SubscriptionAlert example 12345
`+tripleBacktick("")+`
`)
}

func (a *SubscriptionAlertArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Email, "The email address to which alerts are sent.")
	ann.Describe(&a.Description, "Additional information about the subscription alert.")
	ann.Describe(&a.Pt0Severities, "Severity levels for Pt0 alerts.")
	ann.Describe(&a.SecureSeverities, "Severity levels for security alerts.")
	ann.Describe(&a.ManageSeverities, "Severity levels for management alerts.")
	ann.Describe(&a.ComplySeverities, "Severity levels for compliance alerts.")
	ann.Describe(&a.SystemSeverities, "Severity levels for system alerts.")
}

func (s *SubscriptionAlertState) Annotate(a infer.Annotator) {
	a.Describe(&s.AlertId, "The system-generated ID of the subscription alert.")
}

var _ infer.CustomResource[SubscriptionAlertArgs, SubscriptionAlertState] = SubscriptionAlert{}
