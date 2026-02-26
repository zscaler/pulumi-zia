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

// Package provider implements the Location Management resource.
// Adopted from terraform-provider-zia resource_zia_location_management.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/common"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/firewallpolicies/filteringrules"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/location/locationmanagement"
)

// VpnCredentialInput is a nested input for VPN credentials.
type VpnCredentialInput struct {
	Id           *int    `pulumi:"id,optional"`
	Type         *string `pulumi:"type,optional"`
	Fqdn         *string `pulumi:"fqdn,optional"`
	IpAddress    *string `pulumi:"ipAddress,optional"`
	PreSharedKey *string `pulumi:"preSharedKey,optional"`
}

// LocationManagementArgs are the inputs.
type LocationManagementArgs struct {
	Name                                *string              `pulumi:"name,optional"`
	Description                         *string              `pulumi:"description,optional"`
	ParentId                            *int                 `pulumi:"parentId,optional"`
	UpBandwidth                         *int                 `pulumi:"upBandwidth,optional"`
	DnBandwidth                         *int                 `pulumi:"dnBandwidth,optional"`
	Country                             *string              `pulumi:"country,optional"`
	Tz                                  *string              `pulumi:"tz,optional"`
	State                               *string              `pulumi:"state,optional"`
	IpAddresses                         []string             `pulumi:"ipAddresses,optional"`
	Ports                               []int                `pulumi:"ports,optional"`
	VpnCredentials                      []VpnCredentialInput `pulumi:"vpnCredentials,optional"`
	SslScanEnabled                      *bool                `pulumi:"sslScanEnabled,optional"`
	ZappSslScanEnabled                  *bool                `pulumi:"zappSslScanEnabled,optional"`
	XffForwardEnabled                   *bool                `pulumi:"xffForwardEnabled,optional"`
	AuthRequired                        *bool                `pulumi:"authRequired,optional"`
	BasicAuthEnabled                    *bool                `pulumi:"basicAuthEnabled,optional"`
	DigestAuthEnabled                   *bool                `pulumi:"digestAuthEnabled,optional"`
	KerberosAuth                        *bool                `pulumi:"kerberosAuth,optional"`
	SurrogateIp                         *bool                `pulumi:"surrogateIp,optional"`
	IdleTimeInMinutes                   *int                 `pulumi:"idleTimeInMinutes,optional"`
	DisplayTimeUnit                     *string              `pulumi:"displayTimeUnit,optional"`
	SurrogateRefreshTimeInMinutes       *int                 `pulumi:"surrogateRefreshTimeInMinutes,optional"`
	SurrogateRefreshTimeUnit            *string              `pulumi:"surrogateRefreshTimeUnit,optional"`
	SurrogateIpEnforcedForKnownBrowsers *bool                `pulumi:"surrogateIpEnforcedForKnownBrowsers,optional"`
	OfwEnabled                          *bool                `pulumi:"ofwEnabled,optional"`
	IpsControl                          *bool                `pulumi:"ipsControl,optional"`
	Profile                             *string              `pulumi:"profile,optional"`
	StaticLocationGroups                []int                `pulumi:"staticLocationGroups,optional"`
}

// LocationManagementState is the persisted state.
type LocationManagementState struct {
	LocationManagementArgs
	LocationId *int `pulumi:"locationId"`
}

// LocationManagement implements the zia:index:LocationManagement resource.
type LocationManagement struct{}

func vpnCredentialsToAPI(list []VpnCredentialInput) []locationmanagement.VPNCredentials {
	if len(list) == 0 {
		return nil
	}
	result := make([]locationmanagement.VPNCredentials, len(list))
	for i, v := range list {
		cred := locationmanagement.VPNCredentials{
			ID:   ptrToIntDefault(v.Id, 0),
			Type: ptrToString(v.Type),
			FQDN: ptrToString(v.Fqdn),
		}
		// Only set IPAddress when non-empty; the SDK field lacks omitempty
		// and the API rejects UFQDN credentials that include "ipAddress":"".
		if v.IpAddress != nil && *v.IpAddress != "" {
			cred.IPAddress = *v.IpAddress
		}
		if v.PreSharedKey != nil && *v.PreSharedKey != "" {
			cred.PreSharedKey = *v.PreSharedKey
		}
		result[i] = cred
	}
	return result
}

func vpnCredentialsFromAPI(list []locationmanagement.VPNCredentials) []VpnCredentialInput {
	if len(list) == 0 {
		return nil
	}
	result := make([]VpnCredentialInput, len(list))
	for i, v := range list {
		result[i] = VpnCredentialInput{
			Id:           intPtr(v.ID),
			Type:         stringPtr(v.Type),
			Fqdn:         stringPtr(v.FQDN),
			IpAddress:    stringPtr(v.IPAddress),
			PreSharedKey: stringPtr(v.PreSharedKey),
		}
	}
	return result
}

func locationManagementToAPI(args LocationManagementArgs, id int) locationmanagement.Locations {
	loc := locationmanagement.Locations{
		ID:             id,
		Name:           ptrToString(args.Name),
		Description:    ptrToString(args.Description),
		Country:        ptrToString(args.Country),
		TZ:             ptrToString(args.Tz),
		State:          ptrToString(args.State),
		Profile:        ptrToString(args.Profile),
		IPAddresses:    args.IpAddresses,
		Ports:          args.Ports,
		VPNCredentials: vpnCredentialsToAPI(args.VpnCredentials),
	}
	// Only set numeric fields when explicitly provided (avoid sending zero
	// values for fields that the API treats differently from absent).
	if args.ParentId != nil {
		loc.ParentID = *args.ParentId
	}
	if args.UpBandwidth != nil {
		loc.UpBandwidth = *args.UpBandwidth
	}
	if args.DnBandwidth != nil {
		loc.DnBandwidth = *args.DnBandwidth
	}
	// Bool fields: only set when user explicitly provided them
	if args.SslScanEnabled != nil {
		loc.SSLScanEnabled = *args.SslScanEnabled
	}
	if args.ZappSslScanEnabled != nil {
		loc.ZappSSLScanEnabled = *args.ZappSslScanEnabled
	}
	if args.XffForwardEnabled != nil {
		loc.XFFForwardEnabled = *args.XffForwardEnabled
	}
	if args.AuthRequired != nil {
		loc.AuthRequired = *args.AuthRequired
	}
	if args.BasicAuthEnabled != nil {
		loc.BasicAuthEnabled = *args.BasicAuthEnabled
	}
	if args.DigestAuthEnabled != nil {
		loc.DigestAuthEnabled = *args.DigestAuthEnabled
	}
	if args.KerberosAuth != nil {
		loc.KerberosAuth = *args.KerberosAuth
	}
	if args.SurrogateIp != nil {
		loc.SurrogateIP = *args.SurrogateIp
	}
	if args.IdleTimeInMinutes != nil {
		loc.IdleTimeInMinutes = *args.IdleTimeInMinutes
	}
	if args.DisplayTimeUnit != nil {
		loc.DisplayTimeUnit = *args.DisplayTimeUnit
	}
	if args.SurrogateRefreshTimeInMinutes != nil {
		loc.SurrogateRefreshTimeInMinutes = *args.SurrogateRefreshTimeInMinutes
	}
	if args.SurrogateRefreshTimeUnit != nil {
		loc.SurrogateRefreshTimeUnit = *args.SurrogateRefreshTimeUnit
	}
	if args.SurrogateIpEnforcedForKnownBrowsers != nil {
		loc.SurrogateIPEnforcedForKnownBrowsers = *args.SurrogateIpEnforcedForKnownBrowsers
	}
	if args.OfwEnabled != nil {
		loc.OFWEnabled = *args.OfwEnabled
	}
	if args.IpsControl != nil {
		loc.IPSControl = *args.IpsControl
	}
	// Slice fields without omitempty: only set when non-empty to avoid
	// sending "staticLocationGroups":[] which the API may reject.
	if len(args.StaticLocationGroups) > 0 {
		loc.StaticLocationGroups = idsToIDNameExtensions(args.StaticLocationGroups)
	}
	return loc
}

func (LocationManagement) Create(ctx context.Context, req infer.CreateRequest[LocationManagementArgs]) (infer.CreateResponse[LocationManagementState], error) {
	if req.DryRun {
		s := LocationManagementState{LocationManagementArgs: req.Inputs, LocationId: intPtr(0)}
		return infer.CreateResponse[LocationManagementState]{ID: "preview", Output: s}, nil
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[LocationManagementState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := locationManagementToAPI(req.Inputs, 0)
	resp, err := locationmanagement.Create(ctx, service, &apiReq)
	if err != nil {
		return infer.CreateResponse[LocationManagementState]{}, err
	}
	log.Printf("[INFO] Created ZIA location. ID: %v", resp.ID)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[LocationManagementState]{}, activationErr
		}
	}

	state := LocationManagementState{
		LocationManagementArgs: req.Inputs,
		LocationId:             &resp.ID,
	}
	return infer.CreateResponse[LocationManagementState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (LocationManagement) Read(ctx context.Context, req infer.ReadRequest[LocationManagementArgs, LocationManagementState]) (infer.ReadResponse[LocationManagementArgs, LocationManagementState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[LocationManagementArgs, LocationManagementState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.LocationId != nil {
		id = *req.State.LocationId
	}
	if id == 0 {
		return infer.ReadResponse[LocationManagementArgs, LocationManagementState]{}, fmt.Errorf("no location id in state")
	}

	resp, err := locationmanagement.GetLocationOrSublocationByID(ctx, service, id)
	if err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.ReadResponse[LocationManagementArgs, LocationManagementState]{ID: ""}, nil
		}
		return infer.ReadResponse[LocationManagementArgs, LocationManagementState]{}, err
	}

	args := LocationManagementArgs{
		Name:                                stringPtr(resp.Name),
		Description:                         stringPtr(resp.Description),
		ParentId:                            intPtr(resp.ParentID),
		UpBandwidth:                         intPtr(resp.UpBandwidth),
		DnBandwidth:                         intPtr(resp.DnBandwidth),
		Country:                             stringPtr(resp.Country),
		Tz:                                  stringPtr(resp.TZ),
		State:                               stringPtr(resp.State),
		IpAddresses:                         resp.IPAddresses,
		Ports:                               resp.Ports,
		VpnCredentials:                      vpnCredentialsFromAPI(resp.VPNCredentials),
		SslScanEnabled:                      boolPtr(resp.SSLScanEnabled),
		ZappSslScanEnabled:                  boolPtr(resp.ZappSSLScanEnabled),
		XffForwardEnabled:                   boolPtr(resp.XFFForwardEnabled),
		AuthRequired:                        boolPtr(resp.AuthRequired),
		BasicAuthEnabled:                    boolPtr(resp.BasicAuthEnabled),
		DigestAuthEnabled:                   boolPtr(resp.DigestAuthEnabled),
		KerberosAuth:                        boolPtr(resp.KerberosAuth),
		SurrogateIp:                         boolPtr(resp.SurrogateIP),
		IdleTimeInMinutes:                   intPtr(resp.IdleTimeInMinutes),
		DisplayTimeUnit:                     stringPtr(resp.DisplayTimeUnit),
		SurrogateRefreshTimeInMinutes:       intPtr(resp.SurrogateRefreshTimeInMinutes),
		SurrogateRefreshTimeUnit:            stringPtr(resp.SurrogateRefreshTimeUnit),
		SurrogateIpEnforcedForKnownBrowsers: boolPtr(resp.SurrogateIPEnforcedForKnownBrowsers),
		OfwEnabled:                          boolPtr(resp.OFWEnabled),
		IpsControl:                          boolPtr(resp.IPSControl),
		Profile:                             stringPtr(resp.Profile),
		StaticLocationGroups:                idNameExtensionsToIDs(resp.StaticLocationGroups),
	}
	state := LocationManagementState{
		LocationManagementArgs: args,
		LocationId:             &resp.ID,
	}
	return infer.ReadResponse[LocationManagementArgs, LocationManagementState]{
		ID:     strconv.Itoa(resp.ID),
		Inputs: args,
		State:  state,
	}, nil
}

func (LocationManagement) Update(ctx context.Context, req infer.UpdateRequest[LocationManagementArgs, LocationManagementState]) (infer.UpdateResponse[LocationManagementState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[LocationManagementState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	id := 0
	if req.State.LocationId != nil {
		id = *req.State.LocationId
	}
	if id == 0 {
		return infer.UpdateResponse[LocationManagementState]{}, fmt.Errorf("no location id in state")
	}

	if _, err := locationmanagement.GetLocationOrSublocationByID(ctx, service, id); err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.UpdateResponse[LocationManagementState]{}, nil
		}
		return infer.UpdateResponse[LocationManagementState]{}, err
	}

	apiReq := locationManagementToAPI(req.Inputs, id)
	if _, _, err := locationmanagement.Update(ctx, service, id, &apiReq); err != nil {
		return infer.UpdateResponse[LocationManagementState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[LocationManagementState]{}, activationErr
		}
	}

	state := LocationManagementState{
		LocationManagementArgs: req.Inputs,
		LocationId:             &id,
	}
	return infer.UpdateResponse[LocationManagementState]{Output: state}, nil
}

func (LocationManagement) Delete(ctx context.Context, req infer.DeleteRequest[LocationManagementState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	service := client.Service

	id := 0
	if req.State.LocationId != nil {
		id = *req.State.LocationId
	}
	if id != 0 {
		if err := detachFromFilteringRules(ctx, client, id, "Users",
			func(r *filteringrules.FirewallFilteringRules) []common.IDNameExtensions { return r.Users },
			func(r *filteringrules.FirewallFilteringRules, ids []common.IDNameExtensions) { r.Users = ids }); err != nil {
			return infer.DeleteResponse{}, err
		}
		if _, err := locationmanagement.Delete(ctx, service, id); err != nil {
			return infer.DeleteResponse{}, err
		}
		log.Printf("[INFO] ZIA location deleted")
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}

	return infer.DeleteResponse{}, nil
}

func (LocationManagement) Diff(ctx context.Context, req infer.DiffRequest[LocationManagementArgs, LocationManagementState]) (infer.DiffResponse, error) {
	diff, hasChanges := diffSkippingUnsetInputs(req.State.LocationManagementArgs, req.Inputs)
	return infer.DiffResponse{HasChanges: hasChanges, DetailedDiff: diff}, nil
}

func (LocationManagement) Annotate(a infer.Annotator) {
	describeResource(a, &LocationManagement{}, `The zia_location_management resource manages locations in the Zscaler Internet Access (ZIA) cloud service. Locations represent offices, branches, or data centers with specific traffic forwarding and policy settings.

For more information, see the [ZIA Location Management documentation](https://help.zscaler.com/zia/location-management).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Location

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.LocationManagement("example", {
    name: "Example Location",
    description: "Branch office location",
    country: "UNITED_STATES",
    tz: "UNITED_STATES_AMERICA_LOS_ANGELES",
    authRequired: true,
    ofwEnabled: true,
    ipsControl: true,
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.LocationManagement("example",
    name="Example Location",
    description="Branch office location",
    country="UNITED_STATES",
    tz="UNITED_STATES_AMERICA_LOS_ANGELES",
    auth_required=True,
    ofw_enabled=True,
    ips_control=True,
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:LocationManagement
    properties:
      name: Example Location
      description: Branch office location
      country: UNITED_STATES
      tz: UNITED_STATES_AMERICA_LOS_ANGELES
      authRequired: true
      ofwEnabled: true
      ipsControl: true
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing location can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:LocationManagement example 12345
`+tripleBacktick("")+`
`)
}

func (a *LocationManagementArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the location.")
	ann.Describe(&a.Description, "Additional information about the location.")
	ann.Describe(&a.ParentId, "The parent location ID. If this ID does not exist or is 0, it is implied that it is a parent location.")
	ann.Describe(&a.UpBandwidth, "Upload bandwidth in Mbps. If set to 0, the default value is unlimited.")
	ann.Describe(&a.DnBandwidth, "Download bandwidth in Mbps. If set to 0, the default value is unlimited.")
	ann.Describe(&a.Country, "The country in which the location is located.")
	ann.Describe(&a.Tz, "The timezone of the location. Uses IANA-style identifiers.")
	ann.Describe(&a.State, "The state or province in which the location is located.")
	ann.Describe(&a.IpAddresses, "IP addresses associated with this location.")
	ann.Describe(&a.Ports, "IP ports allowed to send traffic to the location.")
	ann.Describe(&a.VpnCredentials, "VPN credentials associated with this location.")
	ann.Describe(&a.SslScanEnabled, "Enable SSL Inspection for the location.")
	ann.Describe(&a.ZappSslScanEnabled, "Enable Zscaler App SSL Setting for the location.")
	ann.Describe(&a.XffForwardEnabled, "Enable XFF Forwarding for the location.")
	ann.Describe(&a.AuthRequired, "Whether authentication is required for this location.")
	ann.Describe(&a.BasicAuthEnabled, "Enable Basic authentication for the location.")
	ann.Describe(&a.DigestAuthEnabled, "Enable Digest authentication for the location.")
	ann.Describe(&a.KerberosAuth, "Enable Kerberos authentication for the location.")
	ann.Describe(&a.SurrogateIp, "Enable surrogate IP enforcement for known browsers.")
	ann.Describe(&a.IdleTimeInMinutes, "Idle time in minutes to disassociate a surrogate IP from the user.")
	ann.Describe(&a.DisplayTimeUnit, "Display time unit. Valid values: `MINUTE`, `HOUR`, `DAY`.")
	ann.Describe(&a.SurrogateRefreshTimeInMinutes, "Refresh time in minutes for re-validating the surrogacy.")
	ann.Describe(&a.SurrogateRefreshTimeUnit, "Display refresh time unit. Valid values: `MINUTE`, `HOUR`, `DAY`.")
	ann.Describe(&a.SurrogateIpEnforcedForKnownBrowsers, "Enforce surrogate IP for known browsers.")
	ann.Describe(&a.OfwEnabled, "Enable firewall for the location.")
	ann.Describe(&a.IpsControl, "Enable IPS control for the location.")
	ann.Describe(&a.Profile, "Profile tag that specifies the location traffic type. Valid values: `NONE`, `CORPORATE`, `SERVER`, `GUESTWIFI`, `IOT`, `WORKLOAD`.")
	ann.Describe(&a.StaticLocationGroups, "IDs of static location groups to which this location belongs.")
}

func (s *LocationManagementState) Annotate(a infer.Annotator) {
	a.Describe(&s.LocationId, "The system-generated ID of the location.")
}

var _ infer.CustomResource[LocationManagementArgs, LocationManagementState] = LocationManagement{}
