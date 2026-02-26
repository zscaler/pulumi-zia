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

// Package provider implements the Cloud NSS Feed resource.
// Adopted from terraform-provider-zia resource_zia_cloud_nss_server.go (resourceCloudNSSFeed).

package provider

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/cloudnss/cloudnss"
)

// CloudNssFeed implements the zia:index:CloudNssFeed resource.
type CloudNssFeed struct{}

// CloudNssFeedArgs are the inputs.
type CloudNssFeedArgs struct {
	Name                     *string  `pulumi:"name,optional"`
	FeedStatus               *string  `pulumi:"feedStatus,optional"`
	NssLogType               *string  `pulumi:"nssLogType,optional"`
	NssFeedType              *string  `pulumi:"nssFeedType,optional"`
	FeedOutputFormat         *string  `pulumi:"feedOutputFormat,optional"`
	TimeZone                 *string  `pulumi:"timeZone,optional"`
	CustomEscapedCharacter    []string `pulumi:"customEscapedCharacter,optional"`
	EpsRateLimit             *int    `pulumi:"epsRateLimit,optional"`
	JsonArrayToggle          *bool    `pulumi:"jsonArrayToggle,optional"`
	SiemType                 *string  `pulumi:"siemType,optional"`
	MaxBatchSize             *int    `pulumi:"maxBatchSize,optional"`
	ConnectionURL             *string  `pulumi:"connectionUrl,optional"`
	AuthenticationToken      *string  `pulumi:"authenticationToken,optional"`
	ConnectionHeaders        []string `pulumi:"connectionHeaders,optional"`
	Base64EncodedCertificate *string  `pulumi:"base64EncodedCertificate,optional"`
	NssType                  *string  `pulumi:"nssType,optional"`
	ClientID                 *string  `pulumi:"clientId,optional"`
	ClientSecret             *string  `pulumi:"clientSecret,optional"`
	AuthenticationUrl        *string  `pulumi:"authenticationUrl,optional"`
	GrantType                *string  `pulumi:"grantType,optional"`
	Scope                    *string  `pulumi:"scope,optional"`
	OauthAuthentication      *bool    `pulumi:"oauthAuthentication,optional"`
	ServerIps                []string `pulumi:"serverIps,optional"`
	ClientIps                []string `pulumi:"clientIps,optional"`
	Domains                  []string `pulumi:"domains,optional"`
	Locations                []int    `pulumi:"locations,optional"`
	LocationGroups           []int    `pulumi:"locationGroups,optional"`
	Departments              []int    `pulumi:"departments,optional"`
	Users                    []int    `pulumi:"users,optional"`
	CasbTenant               []int    `pulumi:"casbTenant,optional"`
	Buckets                  []int    `pulumi:"buckets,optional"`
	VpnCredentials           []int    `pulumi:"vpnCredentials,optional"`
	DlpEngines               []int    `pulumi:"dlpEngines,optional"`
	DlpDictionaries          []int    `pulumi:"dlpDictionaries,optional"`
	ExternalOwners            []int    `pulumi:"externalOwners,optional"`
	ExternalCollaborators     []int    `pulumi:"externalCollaborators,optional"`
	InternalCollaborators    []int    `pulumi:"internalCollaborators,optional"`
	ItsmObjectType           []int    `pulumi:"itsmObjectType,optional"`
	UrlCategories             []int    `pulumi:"urlCategories,optional"`
	Rules                    []int    `pulumi:"rules,optional"`
	NwServices               []int    `pulumi:"nwServices,optional"`
	SenderName               []int    `pulumi:"senderName,optional"`
}

// CloudNssFeedState is the persisted state.
type CloudNssFeedState struct {
	CloudNssFeedArgs
	NssId *int `pulumi:"nssId"`
}

func cloudNssFeedToAPI(args CloudNssFeedArgs, id int) cloudnss.NSSFeed {
	return cloudnss.NSSFeed{
		ID:                          id,
		Name:                        ptrToString(args.Name),
		FeedStatus:                  ptrToString(args.FeedStatus),
		NssLogType:                  ptrToString(args.NssLogType),
		NssFeedType:                 ptrToString(args.NssFeedType),
		FeedOutputFormat:            ptrToString(args.FeedOutputFormat),
		TimeZone:                    ptrToString(args.TimeZone),
		CustomEscapedCharacter:      args.CustomEscapedCharacter,
		EpsRateLimit:                ptrToIntDefault(args.EpsRateLimit, 0),
		JsonArrayToggle:             ptrToBool(args.JsonArrayToggle),
		SiemType:                    ptrToString(args.SiemType),
		MaxBatchSize:                ptrToIntDefault(args.MaxBatchSize, 0),
		ConnectionURL:               ptrToString(args.ConnectionURL),
		AuthenticationToken:         ptrToString(args.AuthenticationToken),
		ConnectionHeaders:           args.ConnectionHeaders,
		Base64EncodedCertificate:    ptrToString(args.Base64EncodedCertificate),
		NssType:                     ptrToString(args.NssType),
		ClientID:                    ptrToString(args.ClientID),
		ClientSecret:                ptrToString(args.ClientSecret),
		AuthenticationUrl:           ptrToString(args.AuthenticationUrl),
		GrantType:                   ptrToString(args.GrantType),
		Scope:                       ptrToString(args.Scope),
		OauthAuthentication:         ptrToBool(args.OauthAuthentication),
		ServerIps:                   args.ServerIps,
		ClientIps:                   args.ClientIps,
		Domains:                     args.Domains,
		Locations:                   idsToCommonNSS(args.Locations),
		LocationGroups:              idsToCommonNSS(args.LocationGroups),
		Departments:                 idsToCommonNSS(args.Departments),
		Users:                       idsToCommonNSS(args.Users),
		CasbTenant:                  idsToCommonNSS(args.CasbTenant),
		Buckets:                     idsToCommonNSS(args.Buckets),
		VPNCredentials:              idsToCommonNSS(args.VpnCredentials),
		DLPEngines:                  idsToIDNameExtensions(args.DlpEngines),
		DLPDictionaries:             idsToIDNameExtensions(args.DlpDictionaries),
		ExternalOwners:              idsToIDNameExtensions(args.ExternalOwners),
		ExternalCollaborators:       idsToIDNameExtensions(args.ExternalCollaborators),
		InternalCollaborators:       idsToIDNameExtensions(args.InternalCollaborators),
		ItsmObjectType:              idsToIDNameExtensions(args.ItsmObjectType),
		URLCategories:               idsToIDNameExtensions(args.UrlCategories),
		Rules:                       idsToIDNameExtensions(args.Rules),
		NwServices:                  idsToIDNameExtensions(args.NwServices),
		SenderName:                  idsToCommonNSS(args.SenderName),
	}
}

func (CloudNssFeed) Create(ctx context.Context, req infer.CreateRequest[CloudNssFeedArgs]) (infer.CreateResponse[CloudNssFeedState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[CloudNssFeedState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := cloudNssFeedToAPI(req.Inputs, 0)
	resp, err := cloudnss.Create(ctx, service, &apiReq)
	if err != nil {
		return infer.CreateResponse[CloudNssFeedState]{}, err
	}
	log.Printf("[INFO] Created ZIA cloud NSS feed. ID: %v", resp.ID)

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[CloudNssFeedState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	state := CloudNssFeedState{
		CloudNssFeedArgs: req.Inputs,
		NssId:            &resp.ID,
	}
	return infer.CreateResponse[CloudNssFeedState]{
		ID:     strconv.Itoa(resp.ID),
		Output: state,
	}, nil
}

func (CloudNssFeed) Read(ctx context.Context, req infer.ReadRequest[CloudNssFeedArgs, CloudNssFeedState]) (infer.ReadResponse[CloudNssFeedArgs, CloudNssFeedState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[CloudNssFeedArgs, CloudNssFeedState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	nssID := 0
	if req.State.NssId != nil {
		nssID = *req.State.NssId
	}
	if nssID == 0 {
		return infer.ReadResponse[CloudNssFeedArgs, CloudNssFeedState]{}, fmt.Errorf("no cloud NSS feed id in state")
	}

	resp, err := cloudnss.Get(ctx, service, nssID)
	if err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.ReadResponse[CloudNssFeedArgs, CloudNssFeedState]{ID: ""}, nil
		}
		return infer.ReadResponse[CloudNssFeedArgs, CloudNssFeedState]{}, err
	}

	args := CloudNssFeedArgs{
		Name:                     stringPtr(resp.Name),
		FeedStatus:               stringPtr(resp.FeedStatus),
		NssLogType:               stringPtr(resp.NssLogType),
		NssFeedType:              stringPtr(resp.NssFeedType),
		FeedOutputFormat:         stringPtr(resp.FeedOutputFormat),
		TimeZone:                 stringPtr(resp.TimeZone),
		CustomEscapedCharacter:   resp.CustomEscapedCharacter,
		EpsRateLimit:             intPtr(resp.EpsRateLimit),
		JsonArrayToggle:          boolPtr(resp.JsonArrayToggle),
		SiemType:                 stringPtr(resp.SiemType),
		MaxBatchSize:             intPtr(resp.MaxBatchSize),
		ConnectionURL:             stringPtr(resp.ConnectionURL),
		AuthenticationToken:      stringPtr(resp.AuthenticationToken),
		ConnectionHeaders:        resp.ConnectionHeaders,
		Base64EncodedCertificate: stringPtr(resp.Base64EncodedCertificate),
		NssType:                  stringPtr(resp.NssType),
		ClientID:                 stringPtr(resp.ClientID),
		ClientSecret:             stringPtr(resp.ClientSecret),
		AuthenticationUrl:        stringPtr(resp.AuthenticationUrl),
		GrantType:                stringPtr(resp.GrantType),
		Scope:                    stringPtr(resp.Scope),
		OauthAuthentication:      boolPtr(resp.OauthAuthentication),
		ServerIps:                resp.ServerIps,
		ClientIps:                resp.ClientIps,
		Domains:                  resp.Domains,
		Locations:                commonNSSToIDs(resp.Locations),
		LocationGroups:           commonNSSToIDs(resp.LocationGroups),
		Departments:              commonNSSToIDs(resp.Departments),
		Users:                    commonNSSToIDs(resp.Users),
		CasbTenant:               commonNSSToIDs(resp.CasbTenant),
		Buckets:                  commonNSSToIDs(resp.Buckets),
		VpnCredentials:           commonNSSToIDs(resp.VPNCredentials),
		DlpEngines:               idNameExtensionsToIDs(resp.DLPEngines),
		DlpDictionaries:          idNameExtensionsToIDs(resp.DLPDictionaries),
		ExternalOwners:           idNameExtensionsToIDs(resp.ExternalOwners),
		ExternalCollaborators:    idNameExtensionsToIDs(resp.ExternalCollaborators),
		InternalCollaborators:    idNameExtensionsToIDs(resp.InternalCollaborators),
		ItsmObjectType:           idNameExtensionsToIDs(resp.ItsmObjectType),
		UrlCategories:            idNameExtensionsToIDs(resp.URLCategories),
		Rules:                    idNameExtensionsToIDs(resp.Rules),
		NwServices:               idNameExtensionsToIDs(resp.NwServices),
		SenderName:               commonNSSToIDs(resp.SenderName),
	}
	state := CloudNssFeedState{
		CloudNssFeedArgs: args,
		NssId:            &resp.ID,
	}
	return infer.ReadResponse[CloudNssFeedArgs, CloudNssFeedState]{
		ID:     strconv.Itoa(resp.ID),
		Inputs: args,
		State:  state,
	}, nil
}

func (CloudNssFeed) Update(ctx context.Context, req infer.UpdateRequest[CloudNssFeedArgs, CloudNssFeedState]) (infer.UpdateResponse[CloudNssFeedState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[CloudNssFeedState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	nssID := 0
	if req.State.NssId != nil {
		nssID = *req.State.NssId
	}
	if nssID == 0 {
		return infer.UpdateResponse[CloudNssFeedState]{}, fmt.Errorf("no cloud NSS feed id in state")
	}

	if _, err := cloudnss.Get(ctx, service, nssID); err != nil {
		if apiErr, ok := err.(*errorx.ErrorResponse); ok && apiErr.IsObjectNotFound() {
			return infer.UpdateResponse[CloudNssFeedState]{}, nil
		}
		return infer.UpdateResponse[CloudNssFeedState]{}, err
	}

	apiReq := cloudNssFeedToAPI(req.Inputs, nssID)
	if _, err := cloudnss.Update(ctx, service, nssID, &apiReq); err != nil {
		return infer.UpdateResponse[CloudNssFeedState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[CloudNssFeedState]{}, activationErr
		}
	}

	state := CloudNssFeedState{
		CloudNssFeedArgs: req.Inputs,
		NssId:            &nssID,
	}
	return infer.UpdateResponse[CloudNssFeedState]{Output: state}, nil
}

func (CloudNssFeed) Delete(ctx context.Context, req infer.DeleteRequest[CloudNssFeedState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	nssID := 0
	if req.State.NssId != nil {
		nssID = *req.State.NssId
	}
	if nssID != 0 {
		if _, err := cloudnss.Delete(ctx, service, nssID); err != nil {
			return infer.DeleteResponse{}, err
		}
		log.Printf("[INFO] ZIA cloud NSS feed deleted")
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.DeleteResponse{}, activationErr
		}
	}

	return infer.DeleteResponse{}, nil
}

func (CloudNssFeed) Annotate(a infer.Annotator) {
	describeResource(a, &CloudNssFeed{}, `The zia_cloud_nss_feed resource manages Cloud NSS (Nanolog Streaming Service) feeds in the Zscaler Internet Access (ZIA) cloud service. Cloud NSS feeds allow you to stream logs from ZIA to external SIEM or log management systems via HTTPS.

{{% examples %}}
## Example Usage

{{% example %}}
### Cloud NSS Feed

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.CloudNssFeed("example", {
    name: "Example NSS Feed",
    feedStatus: "ENABLED",
    nssLogType: "weblog",
    nssFeedType: "NSS_FOR_WEB",
    siemType: "SPLUNK",
    connectionUrl: "https://splunk.example.com:8088/services/collector",
    authenticationToken: "your-auth-token",
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.CloudNssFeed("example",
    name="Example NSS Feed",
    feed_status="ENABLED",
    nss_log_type="weblog",
    nss_feed_type="NSS_FOR_WEB",
    siem_type="SPLUNK",
    connection_url="https://splunk.example.com:8088/services/collector",
    authentication_token="your-auth-token",
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:CloudNssFeed
    properties:
      name: Example NSS Feed
      feedStatus: ENABLED
      nssLogType: weblog
      nssFeedType: NSS_FOR_WEB
      siemType: SPLUNK
      connectionUrl: https://splunk.example.com:8088/services/collector
      authenticationToken: your-auth-token
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing Cloud NSS Feed can be imported using its resource ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:CloudNssFeed example 12345
`+tripleBacktick("")+`
`)
}

func (a *CloudNssFeedArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "The name of the Cloud NSS feed.")
	ann.Describe(&a.FeedStatus, "Status of the feed. Valid values: `ENABLED`, `DISABLED`.")
	ann.Describe(&a.NssLogType, "The NSS log type (e.g. `weblog`, `firewalllog`, `dnslog`).")
	ann.Describe(&a.NssFeedType, "The NSS feed type (e.g. `NSS_FOR_WEB`, `NSS_FOR_FIREWALL`).")
	ann.Describe(&a.FeedOutputFormat, "The output format for the feed.")
	ann.Describe(&a.TimeZone, "The timezone for log timestamps.")
	ann.Describe(&a.CustomEscapedCharacter, "Custom characters to escape in feed output.")
	ann.Describe(&a.EpsRateLimit, "Events per second rate limit.")
	ann.Describe(&a.JsonArrayToggle, "Whether to output logs as a JSON array.")
	ann.Describe(&a.SiemType, "The SIEM type (e.g. `SPLUNK`, `QRADAR`, `SENTINEL`).")
	ann.Describe(&a.MaxBatchSize, "Maximum batch size for log delivery.")
	ann.Describe(&a.ConnectionURL, "The HTTPS connection URL for the SIEM endpoint.")
	ann.Describe(&a.AuthenticationToken, "Authentication token for the SIEM connection.")
	ann.Describe(&a.ConnectionHeaders, "Custom HTTP headers for the connection.")
	ann.Describe(&a.Base64EncodedCertificate, "Base64-encoded certificate for TLS mutual authentication.")
	ann.Describe(&a.NssType, "The NSS type.")
	ann.Describe(&a.ClientID, "OAuth client ID for the SIEM connection.")
	ann.Describe(&a.ClientSecret, "OAuth client secret for the SIEM connection.")
	ann.Describe(&a.AuthenticationUrl, "OAuth authentication URL for the SIEM connection.")
	ann.Describe(&a.GrantType, "OAuth grant type.")
	ann.Describe(&a.Scope, "OAuth scope.")
	ann.Describe(&a.OauthAuthentication, "Whether to use OAuth authentication for the SIEM connection.")
	ann.Describe(&a.ServerIps, "Filter: server IP addresses.")
	ann.Describe(&a.ClientIps, "Filter: client IP addresses.")
	ann.Describe(&a.Domains, "Filter: domain names.")
	ann.Describe(&a.Locations, "Filter: IDs of locations.")
	ann.Describe(&a.LocationGroups, "Filter: IDs of location groups.")
	ann.Describe(&a.Departments, "Filter: IDs of departments.")
	ann.Describe(&a.Users, "Filter: IDs of users.")
	ann.Describe(&a.CasbTenant, "Filter: IDs of CASB tenants.")
	ann.Describe(&a.Buckets, "Filter: IDs of buckets.")
	ann.Describe(&a.VpnCredentials, "Filter: IDs of VPN credentials.")
	ann.Describe(&a.DlpEngines, "Filter: IDs of DLP engines.")
	ann.Describe(&a.DlpDictionaries, "Filter: IDs of DLP dictionaries.")
	ann.Describe(&a.ExternalOwners, "Filter: IDs of external owners.")
	ann.Describe(&a.ExternalCollaborators, "Filter: IDs of external collaborators.")
	ann.Describe(&a.InternalCollaborators, "Filter: IDs of internal collaborators.")
	ann.Describe(&a.ItsmObjectType, "Filter: IDs of ITSM object types.")
	ann.Describe(&a.UrlCategories, "Filter: IDs of URL categories.")
	ann.Describe(&a.Rules, "Filter: IDs of rules.")
	ann.Describe(&a.NwServices, "Filter: IDs of network services.")
	ann.Describe(&a.SenderName, "Filter: IDs of sender names.")
}

func (s *CloudNssFeedState) Annotate(a infer.Annotator) {
	a.Describe(&s.NssId, "The system-generated ID of the Cloud NSS feed.")
}

var _ infer.CustomResource[CloudNssFeedArgs, CloudNssFeedState] = CloudNssFeed{}
