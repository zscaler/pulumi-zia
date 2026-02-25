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

// Package provider implements the Risk Profiles resource.
// Adopted from terraform-provider-zia resource_zia_risk_profiles.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/cloudapplications/risk_profiles"
)

// RiskProfile implements the zia:index:RiskProfile resource.
type RiskProfile struct{}

// RiskProfileArgs are the inputs.
type RiskProfileArgs struct {
	ProfileName               string   `pulumi:"profileName"`
	ProfileType               *string  `pulumi:"profileType,optional"`
	Status                    *string  `pulumi:"status,optional"`
	ExcludeCertificates       *int     `pulumi:"excludeCertificates,optional"`
	PoorItemsOfService        *string  `pulumi:"poorItemsOfService,optional"`
	AdminAuditLogs            *string  `pulumi:"adminAuditLogs,optional"`
	DataBreach                *string  `pulumi:"dataBreach,optional"`
	SourceIpRestrictions      *string  `pulumi:"sourceIpRestrictions,optional"`
	MfaSupport                *string  `pulumi:"mfaSupport,optional"`
	SslPinned                 *string  `pulumi:"sslPinned,optional"`
	HttpSecurityHeaders       *string  `pulumi:"httpSecurityHeaders,optional"`
	Evasive                   *string  `pulumi:"evasive,optional"`
	DnsCaaPolicy              *string  `pulumi:"dnsCaaPolicy,optional"`
	WeakCipherSupport         *string  `pulumi:"weakCipherSupport,optional"`
	PasswordStrength          *string  `pulumi:"passwordStrength,optional"`
	SslCertValidity           *string  `pulumi:"sslCertValidity,optional"`
	Vulnerability             *string  `pulumi:"vulnerability,optional"`
	MalwareScanningForContent *string  `pulumi:"malwareScanningForContent,optional"`
	FileSharing               *string  `pulumi:"fileSharing,optional"`
	SslCertKeySize            *string  `pulumi:"sslCertKeySize,optional"`
	VulnerableToHeartBleed    *string  `pulumi:"vulnerableToHeartBleed,optional"`
	VulnerableToLogJam        *string  `pulumi:"vulnerableToLogJam,optional"`
	VulnerableToPoodle        *string  `pulumi:"vulnerableToPoodle,optional"`
	VulnerabilityDisclosure   *string  `pulumi:"vulnerabilityDisclosure,optional"`
	SupportForWaf             *string  `pulumi:"supportForWaf,optional"`
	RemoteScreenSharing       *string  `pulumi:"remoteScreenSharing,optional"`
	SenderPolicyFramework     *string  `pulumi:"senderPolicyFramework,optional"`
	DomainKeysIdentifiedMail  *string  `pulumi:"domainKeysIdentifiedMail,optional"`
	DomainBasedMessageAuth    *string  `pulumi:"domainBasedMessageAuth,optional"`
	RiskIndex                 []int    `pulumi:"riskIndex,optional"`
	Certifications            []string `pulumi:"certifications,optional"`
	DataEncryptionInTransit   []string `pulumi:"dataEncryptionInTransit,optional"`
	CustomTags                []int   `pulumi:"customTags,optional"`
}

// RiskProfileState is the persisted state.
type RiskProfileState struct {
	RiskProfileArgs
	ProfileId *int `pulumi:"profileId"`
}

func riskProfileArgsToAPI(args *RiskProfileArgs, id int) risk_profiles.RiskProfiles {
	api := risk_profiles.RiskProfiles{
		ID:                        id,
		ProfileName:               args.ProfileName,
		ProfileType:               ptrToString(args.ProfileType),
		Status:                    ptrToString(args.Status),
		ExcludeCertificates:       ptrToIntDefault(args.ExcludeCertificates, 0),
		PoorItemsOfService:        ptrToString(args.PoorItemsOfService),
		AdminAuditLogs:            ptrToString(args.AdminAuditLogs),
		DataBreach:                ptrToString(args.DataBreach),
		SourceIpRestrictions:      ptrToString(args.SourceIpRestrictions),
		MfaSupport:                ptrToString(args.MfaSupport),
		SslPinned:                 ptrToString(args.SslPinned),
		HttpSecurityHeaders:       ptrToString(args.HttpSecurityHeaders),
		Evasive:                   ptrToString(args.Evasive),
		DnsCaaPolicy:              ptrToString(args.DnsCaaPolicy),
		WeakCipherSupport:         ptrToString(args.WeakCipherSupport),
		PasswordStrength:          ptrToString(args.PasswordStrength),
		SslCertValidity:           ptrToString(args.SslCertValidity),
		Vulnerability:             ptrToString(args.Vulnerability),
		MalwareScanningForContent: ptrToString(args.MalwareScanningForContent),
		FileSharing:               ptrToString(args.FileSharing),
		SslCertKeySize:            ptrToString(args.SslCertKeySize),
		VulnerableToHeartBleed:    ptrToString(args.VulnerableToHeartBleed),
		VulnerableToLogJam:        ptrToString(args.VulnerableToLogJam),
		VulnerableToPoodle:        ptrToString(args.VulnerableToPoodle),
		VulnerabilityDisclosure:   ptrToString(args.VulnerabilityDisclosure),
		SupportForWaf:             ptrToString(args.SupportForWaf),
		RemoteScreenSharing:       ptrToString(args.RemoteScreenSharing),
		SenderPolicyFramework:     ptrToString(args.SenderPolicyFramework),
		DomainKeysIdentifiedMail:  ptrToString(args.DomainKeysIdentifiedMail),
		DomainBasedMessageAuth:    ptrToString(args.DomainBasedMessageAuth),
		RiskIndex:                 args.RiskIndex,
		Certifications:            args.Certifications,
		DataEncryptionInTransit:   args.DataEncryptionInTransit,
		CustomTags:                idsToIDNameExternalIDs(args.CustomTags),
	}
	return api
}

func riskProfileAPIToState(api *risk_profiles.RiskProfiles) RiskProfileState {
	state := RiskProfileState{
		RiskProfileArgs: RiskProfileArgs{
			ProfileName:               api.ProfileName,
			ProfileType:               stringPtr(api.ProfileType),
			Status:                    stringPtr(api.Status),
			ExcludeCertificates:       intPtr(api.ExcludeCertificates),
			PoorItemsOfService:        stringPtr(api.PoorItemsOfService),
			AdminAuditLogs:            stringPtr(api.AdminAuditLogs),
			DataBreach:                stringPtr(api.DataBreach),
			SourceIpRestrictions:      stringPtr(api.SourceIpRestrictions),
			MfaSupport:                stringPtr(api.MfaSupport),
			SslPinned:                 stringPtr(api.SslPinned),
			HttpSecurityHeaders:       stringPtr(api.HttpSecurityHeaders),
			Evasive:                   stringPtr(api.Evasive),
			DnsCaaPolicy:              stringPtr(api.DnsCaaPolicy),
			WeakCipherSupport:         stringPtr(api.WeakCipherSupport),
			PasswordStrength:          stringPtr(api.PasswordStrength),
			SslCertValidity:           stringPtr(api.SslCertValidity),
			Vulnerability:             stringPtr(api.Vulnerability),
			MalwareScanningForContent: stringPtr(api.MalwareScanningForContent),
			FileSharing:               stringPtr(api.FileSharing),
			SslCertKeySize:            stringPtr(api.SslCertKeySize),
			VulnerableToHeartBleed:    stringPtr(api.VulnerableToHeartBleed),
			VulnerableToLogJam:        stringPtr(api.VulnerableToLogJam),
			VulnerableToPoodle:        stringPtr(api.VulnerableToPoodle),
			VulnerabilityDisclosure:   stringPtr(api.VulnerabilityDisclosure),
			SupportForWaf:             stringPtr(api.SupportForWaf),
			RemoteScreenSharing:       stringPtr(api.RemoteScreenSharing),
			SenderPolicyFramework:     stringPtr(api.SenderPolicyFramework),
			DomainKeysIdentifiedMail:  stringPtr(api.DomainKeysIdentifiedMail),
			DomainBasedMessageAuth:    stringPtr(api.DomainBasedMessageAuth),
			RiskIndex:                 api.RiskIndex,
			Certifications:            api.Certifications,
			DataEncryptionInTransit:   api.DataEncryptionInTransit,
			CustomTags:                idNameExternalIDsToIDs(api.CustomTags),
		},
		ProfileId: intPtr(api.ID),
	}
	return state
}

func (RiskProfile) Create(ctx context.Context, req infer.CreateRequest[RiskProfileArgs]) (infer.CreateResponse[RiskProfileState], error) {
	if req.DryRun {
		s := RiskProfileState{RiskProfileArgs: req.Inputs, ProfileId: intPtr(0)}
		return infer.CreateResponse[RiskProfileState]{ID: "preview", Output: s}, nil
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[RiskProfileState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	service := client.Service

	apiReq := riskProfileArgsToAPI(&req.Inputs, 0)
	resp, _, err := risk_profiles.Create(ctx, service, &apiReq)
	if err != nil {
		return infer.CreateResponse[RiskProfileState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.CreateResponse[RiskProfileState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	state := riskProfileAPIToState(resp)
	return infer.CreateResponse[RiskProfileState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (RiskProfile) Read(ctx context.Context, req infer.ReadRequest[RiskProfileArgs, RiskProfileState]) (infer.ReadResponse[RiskProfileArgs, RiskProfileState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[RiskProfileArgs, RiskProfileState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		profile, lookupErr := risk_profiles.GetByName(ctx, service, req.ID)
		if lookupErr != nil {
			if respErr, ok := lookupErr.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.ReadResponse[RiskProfileArgs, RiskProfileState]{}, fmt.Errorf("risk profile not found")
			}
			return infer.ReadResponse[RiskProfileArgs, RiskProfileState]{}, lookupErr
		}
		id = profile.ID
	}

	resp, err := risk_profiles.Get(ctx, service, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.ReadResponse[RiskProfileArgs, RiskProfileState]{}, fmt.Errorf("risk profile not found")
		}
		return infer.ReadResponse[RiskProfileArgs, RiskProfileState]{}, err
	}

	state := riskProfileAPIToState(resp)
	return infer.ReadResponse[RiskProfileArgs, RiskProfileState]{
		ID:     req.ID,
		Inputs: state.RiskProfileArgs,
		State:  state,
	}, nil
}

func (RiskProfile) Update(ctx context.Context, req infer.UpdateRequest[RiskProfileArgs, RiskProfileState]) (infer.UpdateResponse[RiskProfileState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[RiskProfileState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	service := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[RiskProfileState]{}, fmt.Errorf("invalid risk profile ID: %s", req.ID)
	}
	apiReq := riskProfileArgsToAPI(&req.Inputs, id)

	if _, _, err := risk_profiles.Update(ctx, service, id, &apiReq); err != nil {
		return infer.UpdateResponse[RiskProfileState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[RiskProfileState]{}, activationErr
		}
	}

	updated, err := risk_profiles.Get(ctx, service, id)
	if err != nil {
		return infer.UpdateResponse[RiskProfileState]{Output: RiskProfileState{
			RiskProfileArgs: req.Inputs,
			ProfileId:       intPtr(id),
		}}, nil
	}
	return infer.UpdateResponse[RiskProfileState]{Output: riskProfileAPIToState(updated)}, nil
}

func (RiskProfile) Delete(ctx context.Context, req infer.DeleteRequest[RiskProfileState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	service := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid risk profile ID: %s", req.ID)
	}
	if _, err := risk_profiles.Delete(ctx, service, id); err != nil {
		return infer.DeleteResponse{}, err
	}
	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}
	return infer.DeleteResponse{}, nil
}

func (RiskProfile) Diff(ctx context.Context, req infer.DiffRequest[RiskProfileArgs, RiskProfileState]) (infer.DiffResponse, error) {
	diff, hasChanges := diffSkippingUnsetInputs(req.State.RiskProfileArgs, req.Inputs)
	return infer.DiffResponse{HasChanges: hasChanges, DetailedDiff: diff}, nil
}

func (RiskProfile) Annotate(a infer.Annotator) {
	describeResource(a, &RiskProfile{}, `The zia.RiskProfile resource manages cloud application risk profiles in the Zscaler Internet Access (ZIA) cloud.
Risk profiles define criteria for evaluating the security posture of cloud applications based on factors such as
certifications, encryption, vulnerability disclosure, and more.

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Risk Profile

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.RiskProfile("example", {
    profileName: "Example Risk Profile",
    profileType: "PREDEFINED",
    status: "ENABLED",
    riskIndex: [1, 2, 3],
    certifications: ["CSA_STAR", "ISO_27001"],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.RiskProfile("example",
    profile_name="Example Risk Profile",
    profile_type="PREDEFINED",
    status="ENABLED",
    risk_index=[1, 2, 3],
    certifications=["CSA_STAR", "ISO_27001"],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:RiskProfile
    properties:
      profileName: Example Risk Profile
      profileType: PREDEFINED
      status: ENABLED
      riskIndex:
        - 1
        - 2
        - 3
      certifications:
        - CSA_STAR
        - ISO_27001
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing risk profile can be imported using its ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:RiskProfile example 12345
`+tripleBacktick("")+`
`)
}

func (a *RiskProfileArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.ProfileName, "Name of the risk profile.")
	ann.Describe(&a.ProfileType, "The profile type (e.g., 'PREDEFINED', 'CUSTOM').")
	ann.Describe(&a.Status, "The status of the risk profile (e.g., 'ENABLED', 'DISABLED').")
	ann.Describe(&a.ExcludeCertificates, "Number of certifications to exclude.")
	ann.Describe(&a.PoorItemsOfService, "Risk level for poor items of service.")
	ann.Describe(&a.AdminAuditLogs, "Risk level for admin audit log support.")
	ann.Describe(&a.DataBreach, "Risk level for data breach history.")
	ann.Describe(&a.SourceIpRestrictions, "Risk level for source IP restrictions.")
	ann.Describe(&a.MfaSupport, "Risk level for MFA support.")
	ann.Describe(&a.SslPinned, "Risk level for SSL pinning.")
	ann.Describe(&a.HttpSecurityHeaders, "Risk level for HTTP security headers.")
	ann.Describe(&a.Evasive, "Risk level for evasive behavior.")
	ann.Describe(&a.DnsCaaPolicy, "Risk level for DNS CAA policy.")
	ann.Describe(&a.WeakCipherSupport, "Risk level for weak cipher support.")
	ann.Describe(&a.PasswordStrength, "Risk level for password strength enforcement.")
	ann.Describe(&a.SslCertValidity, "Risk level for SSL certificate validity.")
	ann.Describe(&a.Vulnerability, "Risk level for known vulnerabilities.")
	ann.Describe(&a.MalwareScanningForContent, "Risk level for malware scanning.")
	ann.Describe(&a.FileSharing, "Risk level for file sharing support.")
	ann.Describe(&a.SslCertKeySize, "Risk level for SSL certificate key size.")
	ann.Describe(&a.VulnerableToHeartBleed, "Risk level for HeartBleed vulnerability.")
	ann.Describe(&a.VulnerableToLogJam, "Risk level for LogJam vulnerability.")
	ann.Describe(&a.VulnerableToPoodle, "Risk level for POODLE vulnerability.")
	ann.Describe(&a.VulnerabilityDisclosure, "Risk level for vulnerability disclosure policy.")
	ann.Describe(&a.SupportForWaf, "Risk level for WAF support.")
	ann.Describe(&a.RemoteScreenSharing, "Risk level for remote screen sharing support.")
	ann.Describe(&a.SenderPolicyFramework, "Risk level for SPF support.")
	ann.Describe(&a.DomainKeysIdentifiedMail, "Risk level for DKIM support.")
	ann.Describe(&a.DomainBasedMessageAuth, "Risk level for DMARC support.")
	ann.Describe(&a.RiskIndex, "List of risk index values.")
	ann.Describe(&a.Certifications, "List of required certifications (e.g., 'CSA_STAR', 'ISO_27001').")
	ann.Describe(&a.DataEncryptionInTransit, "List of data encryption in transit protocols.")
	ann.Describe(&a.CustomTags, "List of custom tag IDs associated with the profile.")
}

func (s *RiskProfileState) Annotate(a infer.Annotator) {
	a.Describe(&s.ProfileId, "The unique identifier for the risk profile assigned by the ZIA cloud.")
}

var _ infer.CustomResource[RiskProfileArgs, RiskProfileState] = RiskProfile{}
