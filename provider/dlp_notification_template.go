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

// Package provider implements the DLP Notification Template resource.
// Adopted from terraform-provider-zia resource_zia_dlp_notification_templates.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/dlp/dlp_notification_templates"
)

// DlpNotificationTemplate implements the zia:index:DlpNotificationTemplate resource.
type DlpNotificationTemplate struct{}

// DlpNotificationTemplateArgs are the inputs.
type DlpNotificationTemplateArgs struct {
	Name             *string `pulumi:"name,optional"`
	Subject          *string `pulumi:"subject,optional"`
	AttachContent    *bool   `pulumi:"attachContent,optional"`
	PlainTextMessage *string `pulumi:"plainTextMessage,optional"`
	HtmlMessage      *string `pulumi:"htmlMessage,optional"`
	TlsEnabled       *bool   `pulumi:"tlsEnabled,optional"`
}

// DlpNotificationTemplateState is the persisted state.
type DlpNotificationTemplateState struct {
	DlpNotificationTemplateArgs
	TemplateId *int `pulumi:"templateId"`
}

func dlpNotificationTemplateToAPI(args DlpNotificationTemplateArgs, id int) dlp_notification_templates.DlpNotificationTemplates {
	return dlp_notification_templates.DlpNotificationTemplates{
		ID:               id,
		Name:             ptrToString(args.Name),
		Subject:          ptrToString(args.Subject),
		AttachContent:    ptrToBool(args.AttachContent),
		PlainTextMessage: ptrToString(args.PlainTextMessage),
		HtmlMessage:      ptrToString(args.HtmlMessage),
		TLSEnabled:       ptrToBool(args.TlsEnabled),
	}
}

func (DlpNotificationTemplate) Create(ctx context.Context, req infer.CreateRequest[DlpNotificationTemplateArgs]) (infer.CreateResponse[DlpNotificationTemplateState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[DlpNotificationTemplateState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := dlpNotificationTemplateToAPI(req.Inputs, 0)
	resp, _, err := dlp_notification_templates.Create(ctx, service, &apiReq)
	if err != nil {
		return infer.CreateResponse[DlpNotificationTemplateState]{}, err
	}
	log.Printf("[INFO] Created ZIA DLP notification template. ID: %v", resp.ID)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[DlpNotificationTemplateState]{}, activationErr
		}
	}

	state := DlpNotificationTemplateState{
		DlpNotificationTemplateArgs: req.Inputs,
		TemplateId:                  &resp.ID,
	}
	return infer.CreateResponse[DlpNotificationTemplateState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (DlpNotificationTemplate) Read(ctx context.Context, req infer.ReadRequest[DlpNotificationTemplateArgs, DlpNotificationTemplateState]) (infer.ReadResponse[DlpNotificationTemplateArgs, DlpNotificationTemplateState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[DlpNotificationTemplateArgs, DlpNotificationTemplateState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.TemplateId != nil {
		id = *req.State.TemplateId
	}
	if id == 0 {
		return infer.ReadResponse[DlpNotificationTemplateArgs, DlpNotificationTemplateState]{}, fmt.Errorf("no DLP notification template id in state")
	}

	resp, err := dlp_notification_templates.Get(ctx, service, id)
	if err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.ReadResponse[DlpNotificationTemplateArgs, DlpNotificationTemplateState]{ID: ""}, nil
		}
		return infer.ReadResponse[DlpNotificationTemplateArgs, DlpNotificationTemplateState]{}, err
	}

	args := DlpNotificationTemplateArgs{
		Name:             stringPtr(resp.Name),
		Subject:          stringPtr(resp.Subject),
		AttachContent:    boolPtr(resp.AttachContent),
		PlainTextMessage: stringPtr(resp.PlainTextMessage),
		HtmlMessage:      stringPtr(resp.HtmlMessage),
		TlsEnabled:       boolPtr(resp.TLSEnabled),
	}
	state := DlpNotificationTemplateState{
		DlpNotificationTemplateArgs: args,
		TemplateId:                  &resp.ID,
	}
	return infer.ReadResponse[DlpNotificationTemplateArgs, DlpNotificationTemplateState]{
		ID:     strconv.Itoa(resp.ID),
		Inputs: args,
		State:  state,
	}, nil
}

func (DlpNotificationTemplate) Update(ctx context.Context, req infer.UpdateRequest[DlpNotificationTemplateArgs, DlpNotificationTemplateState]) (infer.UpdateResponse[DlpNotificationTemplateState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[DlpNotificationTemplateState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.TemplateId != nil {
		id = *req.State.TemplateId
	}
	if id == 0 {
		return infer.UpdateResponse[DlpNotificationTemplateState]{}, fmt.Errorf("no DLP notification template id in state")
	}

	if _, err := dlp_notification_templates.Get(ctx, service, id); err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.UpdateResponse[DlpNotificationTemplateState]{}, nil
		}
		return infer.UpdateResponse[DlpNotificationTemplateState]{}, err
	}

	apiReq := dlpNotificationTemplateToAPI(req.Inputs, id)
	if _, _, err := dlp_notification_templates.Update(ctx, service, id, &apiReq); err != nil {
		return infer.UpdateResponse[DlpNotificationTemplateState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[DlpNotificationTemplateState]{}, activationErr
		}
	}

	state := DlpNotificationTemplateState{
		DlpNotificationTemplateArgs: req.Inputs,
		TemplateId:                  &id,
	}
	return infer.UpdateResponse[DlpNotificationTemplateState]{Output: state}, nil
}

func (DlpNotificationTemplate) Delete(ctx context.Context, req infer.DeleteRequest[DlpNotificationTemplateState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.TemplateId != nil {
		id = *req.State.TemplateId
	}
	if id != 0 {
		if _, err := dlp_notification_templates.Delete(ctx, service, id); err != nil {
			return infer.DeleteResponse{}, err
		}
		log.Printf("[INFO] ZIA DLP notification template deleted")
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}

	return infer.DeleteResponse{}, nil
}

func (DlpNotificationTemplate) Annotate(a infer.Annotator) {
	describeResource(a, &DlpNotificationTemplate{}, `The zia_dlp_notification_templates resource manages DLP (Data Loss Prevention) notification templates in the Zscaler Internet Access (ZIA) cloud service. DLP notification templates define the email notifications sent to users or auditors when a DLP policy rule is triggered.

For more information, see the [ZIA Data Loss Prevention documentation](https://help.zscaler.com/zia/data-loss-prevention).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic DLP Notification Template

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.DlpNotificationTemplate("example", {
    name: "Example DLP Notification",
    subject: "DLP Policy Violation Detected",
    attachContent: false,
    plainTextMessage: "A DLP policy violation was detected.",
    tlsEnabled: true,
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.DlpNotificationTemplate("example",
    name="Example DLP Notification",
    subject="DLP Policy Violation Detected",
    attach_content=False,
    plain_text_message="A DLP policy violation was detected.",
    tls_enabled=True,
)
`+tripleBacktick("")+`

`+tripleBacktick("go")+`
import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	zia "github.com/zscaler/pulumi-zia/sdk/go/pulumi-zia"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := zia.NewDlpNotificationTemplate(ctx, "example", &zia.DlpNotificationTemplateArgs{
			Name:             pulumi.StringRef("Example DLP Notification"),
			Subject:          pulumi.StringRef("DLP Policy Violation Detected"),
			AttachContent:    pulumi.BoolRef(false),
			PlainTextMessage: pulumi.StringRef("A DLP policy violation was detected."),
			TlsEnabled:       pulumi.BoolRef(true),
		})
		return err
	})
}
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:DlpNotificationTemplate
    properties:
      name: Example DLP Notification
      subject: DLP Policy Violation Detected
      attachContent: false
      plainTextMessage: A DLP policy violation was detected.
      tlsEnabled: true
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing DLP Notification Template can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:DlpNotificationTemplate example 12345
`+tripleBacktick("")+`
`)
}

func (a *DlpNotificationTemplateArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the DLP notification template. Must be unique.")
	ann.Describe(&a.Subject, "The subject line of the DLP notification email.")
	ann.Describe(&a.AttachContent, "If true, the content that triggered the DLP violation is attached to the notification email.")
	ann.Describe(&a.PlainTextMessage, "The plain text message body of the DLP notification email.")
	ann.Describe(&a.HtmlMessage, "The HTML message body of the DLP notification email.")
	ann.Describe(&a.TlsEnabled, "If true, TLS is enabled for delivering the DLP notification email.")
}

func (s *DlpNotificationTemplateState) Annotate(a infer.Annotator) {
	a.Describe(&s.TemplateId, "The system-generated ID of the DLP notification template.")
}

var _ infer.CustomResource[DlpNotificationTemplateArgs, DlpNotificationTemplateState] = DlpNotificationTemplate{}
