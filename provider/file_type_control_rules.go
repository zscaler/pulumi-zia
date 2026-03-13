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

// Package provider implements the File Type Control Rules resource.
// Uses filetypecontrol package. Predefined rules cannot be deleted.

package provider

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/errorx"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/filetypecontrol"
)

const fileTypeControlResourceType = "file_type_control_rules"

// FileTypeControlRule implements the zia:index:FileTypeControlRule resource.
type FileTypeControlRule struct{}

// FileTypeControlRuleArgs are the inputs.
type FileTypeControlRuleArgs struct {
	Name                 string               `pulumi:"name"`
	Order                int                  `pulumi:"order"`
	Description          *string              `pulumi:"description,optional"`
	State                *string              `pulumi:"state,optional"`
	Rank                 *int                 `pulumi:"rank,optional"`
	FilteringAction      *string              `pulumi:"filteringAction,optional"`
	Operation            *string              `pulumi:"operation,optional"`
	TimeQuota            *int                 `pulumi:"timeQuota,optional"`
	SizeQuota            *int                 `pulumi:"sizeQuota,optional"`
	CapturePcap          *bool                `pulumi:"capturePcap,optional"`
	PasswordProtected    *bool                `pulumi:"passwordProtected,optional"`
	ActiveContent        *bool                `pulumi:"activeContent,optional"`
	Unscannable          *bool                `pulumi:"unscannable,optional"`
	FileTypes            []string             `pulumi:"fileTypes,optional"`
	URLCategories        []string             `pulumi:"urlCategories,optional"`
	Protocols            []string             `pulumi:"protocols,optional"`
	CloudApplications    []string             `pulumi:"cloudApplications,optional"`
	MinSize              *int                 `pulumi:"minSize,optional"`
	MaxSize              *int                 `pulumi:"maxSize,optional"`
	BrowserEunTemplateID *int                 `pulumi:"browserEunTemplateId,optional"`
	Locations            []int                `pulumi:"locations,optional"`
	LocationGroups       []int                `pulumi:"locationGroups,optional"`
	Departments          []int                `pulumi:"departments,optional"`
	Groups               []int                `pulumi:"groups,optional"`
	Users                []int                `pulumi:"users,optional"`
	TimeWindows          []int                `pulumi:"timeWindows,optional"`
	Labels               []int                `pulumi:"labels,optional"`
	DeviceGroups         []int                `pulumi:"deviceGroups,optional"`
	Devices              []int                `pulumi:"devices,optional"`
	DeviceTrustLevels    []string             `pulumi:"deviceTrustLevels,optional"`
	ZpaAppSegments       []ZPAAppSegmentInput `pulumi:"zpaAppSegments,optional"`
}

// FileTypeControlRuleState is the persisted state.
type FileTypeControlRuleState struct {
	FileTypeControlRuleArgs
	RuleID *int `pulumi:"ruleId"`
}

func fileTypeControlRuleArgsToAPI(args *FileTypeControlRuleArgs, id int) filetypecontrol.FileTypeRules {
	order := args.Order
	if order == 0 {
		order = 1
	}
	return filetypecontrol.FileTypeRules{
		ID:                   id,
		Name:                 args.Name,
		Order:                order,
		Description:          ptrToString(args.Description),
		State:                ptrToString(args.State),
		Rank:                 ptrToIntDefault(args.Rank, 7),
		FilteringAction:      ptrToString(args.FilteringAction),
		Operation:            ptrToString(args.Operation),
		TimeQuota:            ptrToIntDefault(args.TimeQuota, 0),
		SizeQuota:            ptrToIntDefault(args.SizeQuota, 0),
		CapturePCAP:          ptrToBool(args.CapturePcap),
		PasswordProtected:    ptrToBool(args.PasswordProtected),
		ActiveContent:        ptrToBool(args.ActiveContent),
		Unscannable:          ptrToBool(args.Unscannable),
		FileTypes:            args.FileTypes,
		URLCategories:        args.URLCategories,
		Protocols:            args.Protocols,
		CloudApplications:    args.CloudApplications,
		MinSize:              ptrToIntDefault(args.MinSize, 0),
		MaxSize:              ptrToIntDefault(args.MaxSize, 0),
		BrowserEunTemplateID: ptrToIntDefault(args.BrowserEunTemplateID, 0),
		Locations:            idsToIDNameExtensions(args.Locations),
		LocationGroups:       idsToIDNameExtensions(args.LocationGroups),
		Departments:          idsToIDNameExtensions(args.Departments),
		Groups:               idsToIDNameExtensions(args.Groups),
		Users:                idsToIDNameExtensions(args.Users),
		TimeWindows:          idsToIDNameExtensions(args.TimeWindows),
		Labels:               idsToIDNameExtensions(args.Labels),
		DeviceGroups:         idsToIDNameExtensions(args.DeviceGroups),
		Devices:              idsToIDNameExtensions(args.Devices),
		DeviceTrustLevels:    args.DeviceTrustLevels,
		ZPAAppSegments:       expandZPAAppSegments(args.ZpaAppSegments),
	}
}

func fileTypeControlRuleAPIToState(api *filetypecontrol.FileTypeRules) FileTypeControlRuleState {
	return FileTypeControlRuleState{
		FileTypeControlRuleArgs: FileTypeControlRuleArgs{
			Name:                 api.Name,
			Order:                api.Order,
			Description:          stringPtr(api.Description),
			State:                stringPtr(api.State),
			Rank:                 intPtr(api.Rank),
			FilteringAction:      stringPtr(api.FilteringAction),
			Operation:            stringPtr(api.Operation),
			TimeQuota:            intPtr(api.TimeQuota),
			SizeQuota:            intPtr(api.SizeQuota),
			CapturePcap:          boolPtr(api.CapturePCAP),
			PasswordProtected:    boolPtr(api.PasswordProtected),
			ActiveContent:        boolPtr(api.ActiveContent),
			Unscannable:          boolPtr(api.Unscannable),
			FileTypes:            api.FileTypes,
			URLCategories:        api.URLCategories,
			Protocols:            api.Protocols,
			CloudApplications:    api.CloudApplications,
			MinSize:              intPtr(api.MinSize),
			MaxSize:              intPtr(api.MaxSize),
			BrowserEunTemplateID: intPtr(api.BrowserEunTemplateID),
			Locations:            idNameExtensionsToIDs(api.Locations),
			LocationGroups:       idNameExtensionsToIDs(api.LocationGroups),
			Departments:          idNameExtensionsToIDs(api.Departments),
			Groups:               idNameExtensionsToIDs(api.Groups),
			Users:                idNameExtensionsToIDs(api.Users),
			TimeWindows:          idNameExtensionsToIDs(api.TimeWindows),
			Labels:               idNameExtensionsToIDs(api.Labels),
			DeviceGroups:         idNameExtensionsToIDs(api.DeviceGroups),
			Devices:              idNameExtensionsToIDs(api.Devices),
			DeviceTrustLevels:    api.DeviceTrustLevels,
			ZpaAppSegments:       flattenZPAAppSegments(api.ZPAAppSegments),
		},
		RuleID: intPtr(api.ID),
	}
}

func (FileTypeControlRule) Create(ctx context.Context, req infer.CreateRequest[FileTypeControlRuleArgs]) (infer.CreateResponse[FileTypeControlRuleState], error) {
	if req.DryRun {
		s := FileTypeControlRuleState{FileTypeControlRuleArgs: req.Inputs, RuleID: intPtr(0)}
		return infer.CreateResponse[FileTypeControlRuleState]{ID: "preview", Output: s}, nil
	}
	if req.Inputs.Order < 1 {
		return infer.CreateResponse[FileTypeControlRuleState]{}, fmt.Errorf("order must be a positive whole number (>= 1), got %d", req.Inputs.Order)
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[FileTypeControlRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	apiReq := fileTypeControlRuleArgsToAPI(&req.Inputs, 0)
	timeout := 60 * time.Minute
	start := time.Now()

	intendedOrder := apiReq.Order
	intendedRank := ptrToIntDefault(req.Inputs.Rank, 7)

	for {
		select {
		case fileTypeSem <- struct{}{}:
		case <-ctx.Done():
			return infer.CreateResponse[FileTypeControlRuleState]{}, fmt.Errorf("creating file type control rule: %w", ctx.Err())
		}

		fileTypeOrderMu.Lock()
		if fileTypeStartingOrder == 0 {
			list, _ := filetypecontrol.GetAll(ctx, svc)
			for _, r := range list {
				if r.Order > fileTypeStartingOrder {
					fileTypeStartingOrder = r.Order
				}
			}
			if fileTypeStartingOrder == 0 {
				fileTypeStartingOrder = 1
			}
		}

		if intendedRank < 7 {
			apiReq.Rank = 7
		}
		apiReq.Order = fileTypeStartingOrder
		fileTypeOrderMu.Unlock()

		resp, err := filetypecontrol.Create(ctx, svc, &apiReq)

		if err == nil {
			fileTypeOrderMu.Lock()
			fileTypeStartingOrder++
			fileTypeOrderMu.Unlock()
		}

		<-fileTypeSem

		if customErr := failFastOnErrorCodes(err); customErr != nil {
			return infer.CreateResponse[FileTypeControlRuleState]{}, customErr
		}
		if err != nil {
			reg := regexp.MustCompile("Rule with rank [0-9]+ is not allowed at order [0-9]+")
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") && reg.MatchString(err.Error()) {
				return infer.CreateResponse[FileTypeControlRuleState]{}, fmt.Errorf("error creating file type control rule: %s, check order %d vs rank %d, err:%s",
					apiReq.Name, intendedOrder, intendedRank, err)
			}
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") {
				log.Printf("[INFO] Creating file type control rule name: %v, got INVALID_INPUT_ARGUMENT\n", apiReq.Name)
				fileTypeOrderMu.Lock()
				fileTypeStartingOrder = 0
				fileTypeOrderMu.Unlock()
				if time.Since(start) < timeout {
					select {
					case <-time.After(10 * time.Second):
					case <-ctx.Done():
						return infer.CreateResponse[FileTypeControlRuleState]{}, fmt.Errorf("creating file type control rule: %w", ctx.Err())
					}
					continue
				}
			}
			return infer.CreateResponse[FileTypeControlRuleState]{}, fmt.Errorf("creating file type control rule: %w", err)
		}

		reorderWithBeforeReorder(
			OrderRule{Order: intendedOrder, Rank: intendedRank},
			resp.ID,
			fileTypeControlResourceType,
			func() (int, error) {
				allRules, err := filetypecontrol.GetAll(ctx, svc)
				if err != nil {
					return 0, err
				}
				return len(allRules), nil
			},
			func(id int, order OrderRule) error {
				rule, err := filetypecontrol.Get(ctx, svc, id)
				if err != nil {
					return err
				}
				// to avoid the STALE_CONFIGURATION_ERROR
				rule.LastModifiedTime = 0
				rule.LastModifiedBy = nil
				// Strip read-only fields that cause "Request body is invalid" for predefined rules
				rule.AccessControl = ""
				rule.Order = order.Order
				rule.Rank = order.Rank
				_, err = filetypecontrol.Update(ctx, svc, id, rule)
				return err
			},
			nil,
		)

		markOrderRuleAsDone(resp.ID, fileTypeControlResourceType)
		waitForReorder(fileTypeControlResourceType)

		if shouldActivate() {
			time.Sleep(2 * time.Second)
			if activationErr := triggerActivation(ctx, client); activationErr != nil {
				return infer.CreateResponse[FileTypeControlRuleState]{}, activationErr
			}
		} else {
			log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
		}

		rule, err := filetypecontrol.Get(ctx, svc, resp.ID)
		if err != nil {
			state := FileTypeControlRuleState{FileTypeControlRuleArgs: req.Inputs, RuleID: intPtr(resp.ID)}
			return infer.CreateResponse[FileTypeControlRuleState]{ID: strconv.Itoa(resp.ID), Output: state}, nil
		}
		return infer.CreateResponse[FileTypeControlRuleState]{
			ID:     strconv.Itoa(resp.ID),
			Output: fileTypeControlRuleAPIToState(rule),
		}, nil
	}
}

func (FileTypeControlRule) Read(ctx context.Context, req infer.ReadRequest[FileTypeControlRuleArgs, FileTypeControlRuleState]) (infer.ReadResponse[FileTypeControlRuleArgs, FileTypeControlRuleState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[FileTypeControlRuleArgs, FileTypeControlRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	svc := cfg.Client().Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		rule, lookupErr := filetypecontrol.GetByName(ctx, svc, req.ID)
		if lookupErr != nil {
			if respErr, ok := lookupErr.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
				return infer.ReadResponse[FileTypeControlRuleArgs, FileTypeControlRuleState]{}, fmt.Errorf("file type control rule not found")
			}
			return infer.ReadResponse[FileTypeControlRuleArgs, FileTypeControlRuleState]{}, lookupErr
		}
		id = rule.ID
	}

	rule, err := filetypecontrol.Get(ctx, svc, id)
	if err != nil {
		if respErr, ok := err.(*errorx.ErrorResponse); ok && respErr.IsObjectNotFound() {
			return infer.ReadResponse[FileTypeControlRuleArgs, FileTypeControlRuleState]{}, fmt.Errorf("file type control rule not found")
		}
		return infer.ReadResponse[FileTypeControlRuleArgs, FileTypeControlRuleState]{}, err
	}

	state := fileTypeControlRuleAPIToState(rule)
	return infer.ReadResponse[FileTypeControlRuleArgs, FileTypeControlRuleState]{
		ID:     req.ID,
		Inputs: state.FileTypeControlRuleArgs,
		State:  state,
	}, nil
}

func (FileTypeControlRule) Update(ctx context.Context, req infer.UpdateRequest[FileTypeControlRuleArgs, FileTypeControlRuleState]) (infer.UpdateResponse[FileTypeControlRuleState], error) {
	if req.Inputs.Order < 1 {
		return infer.UpdateResponse[FileTypeControlRuleState]{}, fmt.Errorf("order must be a positive whole number (>= 1), got %d", req.Inputs.Order)
	}
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[FileTypeControlRuleState]{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[FileTypeControlRuleState]{}, fmt.Errorf("invalid file type control rule ID: %s", req.ID)
	}
	apiReq := fileTypeControlRuleArgsToAPI(&req.Inputs, id)

	existingRules, err := filetypecontrol.GetAll(ctx, svc)
	if err == nil && len(existingRules) > 0 {
		sort.Slice(existingRules, func(i, j int) bool {
			return existingRules[i].Rank < existingRules[j].Rank || (existingRules[i].Rank == existingRules[j].Rank && existingRules[i].Order < existingRules[j].Order)
		})
		apiReq.Rank = 7
		apiReq.Order = existingRules[len(existingRules)-1].Order
	}

	timeout := 60 * time.Minute
	start := time.Now()
	for {
		if _, err := filetypecontrol.Update(ctx, svc, id, &apiReq); err != nil {
			if customErr := failFastOnErrorCodes(err); customErr != nil {
				return infer.UpdateResponse[FileTypeControlRuleState]{}, customErr
			}
			if strings.Contains(err.Error(), "INVALID_INPUT_ARGUMENT") {
				log.Printf("[INFO] Updating file type control rule ID: %v, got INVALID_INPUT_ARGUMENT\n", id)
				if time.Since(start) < timeout {
					time.Sleep(10 * time.Second)
					continue
				}
			}
			return infer.UpdateResponse[FileTypeControlRuleState]{}, err
		}

		intendedOrder := req.Inputs.Order
		intendedRank := ptrToIntDefault(req.Inputs.Rank, 7)
		reorderWithBeforeReorder(
			OrderRule{Order: intendedOrder, Rank: intendedRank},
			id,
			fileTypeControlResourceType,
			func() (int, error) {
				allRules, err := filetypecontrol.GetAll(ctx, svc)
				if err != nil {
					return 0, err
				}
				return len(allRules), nil
			},
			func(ruleID int, order OrderRule) error {
				rule, err := filetypecontrol.Get(ctx, svc, ruleID)
				if err != nil {
					return err
				}
				// to avoid the STALE_CONFIGURATION_ERROR
				rule.LastModifiedTime = 0
				rule.LastModifiedBy = nil
				// Strip read-only fields that cause "Request body is invalid" for predefined rules
				rule.AccessControl = ""
				rule.Order = order.Order
				rule.Rank = order.Rank
				_, err = filetypecontrol.Update(ctx, svc, ruleID, rule)
				return err
			},
			nil,
		)

		markOrderRuleAsDone(id, fileTypeControlResourceType)
		waitForReorder(fileTypeControlResourceType)
		break
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, client); activationErr != nil {
			return infer.UpdateResponse[FileTypeControlRuleState]{}, activationErr
		}
	}

	updated, err := filetypecontrol.Get(ctx, svc, id)
	if err != nil {
		return infer.UpdateResponse[FileTypeControlRuleState]{Output: FileTypeControlRuleState{
			FileTypeControlRuleArgs: req.Inputs,
			RuleID:                  intPtr(id),
		}}, nil
	}
	return infer.UpdateResponse[FileTypeControlRuleState]{Output: fileTypeControlRuleAPIToState(updated)}, nil
}

func (FileTypeControlRule) Delete(ctx context.Context, req infer.DeleteRequest[FileTypeControlRuleState]) (infer.DeleteResponse, error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.DeleteResponse{}, fmt.Errorf("ZIA provider not configured")
	}
	client := cfg.Client()
	svc := client.Service

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("invalid file type control rule ID: %s", req.ID)
	}
	// Predefined rules cannot be deleted; API will reject if applicable.
	if _, err := filetypecontrol.Delete(ctx, svc, id); err != nil {
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

func (FileTypeControlRule) Annotate(a infer.Annotator) {
	describeResource(a, &FileTypeControlRule{}, `The zia.FileTypeControlRule resource manages file type control rules in the Zscaler Internet Access (ZIA) cloud.
File type control rules allow you to block, caution, or allow file downloads and uploads based on file types,
protocols, URL categories, and other criteria. Predefined rules cannot be deleted.

{{% examples %}}
## Example Usage

{{% example %}}
### Basic File Type Control Rule

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.FileTypeControlRule("example", {
    name: "Example File Type Rule",
    order: 1,
    description: "Managed by Pulumi",
    state: "ENABLED",
    filteringAction: "BLOCK",
    fileTypes: ["EXE", "DLL"],
    protocols: ["FTP_RULE", "HTTPS_RULE", "HTTP_PROXY"],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.FileTypeControlRule("example",
    name="Example File Type Rule",
    order=1,
    description="Managed by Pulumi",
    state="ENABLED",
    filtering_action="BLOCK",
    file_types=["EXE", "DLL"],
    protocols=["FTP_RULE", "HTTPS_RULE", "HTTP_PROXY"],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:FileTypeControlRule
    properties:
      name: Example File Type Rule
      order: 1
      description: Managed by Pulumi
      state: ENABLED
      filteringAction: BLOCK
      fileTypes:
        - EXE
        - DLL
      protocols:
        - FTP_RULE
        - HTTPS_RULE
        - HTTP_PROXY
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

An existing file type control rule can be imported using its ID, e.g.

`+tripleBacktick("sh")+`
$ pulumi import zia:index:FileTypeControlRule example 12345
`+tripleBacktick("")+`
`)
}

func (a *FileTypeControlRuleArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Name, "Name of the file type control rule.")
	ann.Describe(&a.Order, "The rule order of execution for the file type control rule.")
	ann.Describe(&a.Description, "Description of the file type control rule.")
	ann.Describe(&a.State, "The rule state. Accepted values: 'ENABLED' or 'DISABLED'.")
	ann.Describe(&a.Rank, "The admin rank of the rule. Default is 7.")
	ann.Describe(&a.FilteringAction, "The action taken when traffic matches the rule (e.g., 'BLOCK', 'CAUTION', 'ALLOW').")
	ann.Describe(&a.Operation, "The type of file operation (e.g., 'DOWNLOAD', 'UPLOAD').")
	ann.Describe(&a.TimeQuota, "Time quota in minutes after which the URL filtering rule is applied.")
	ann.Describe(&a.SizeQuota, "Size quota in KB beyond which the URL filtering rule is applied.")
	ann.Describe(&a.CapturePcap, "Whether to capture PCAP data for the rule.")
	ann.Describe(&a.PasswordProtected, "Whether the rule applies to password-protected files.")
	ann.Describe(&a.ActiveContent, "Whether the rule applies to files with active content.")
	ann.Describe(&a.Unscannable, "Whether the rule applies to unscannable files.")
	ann.Describe(&a.FileTypes, "List of file types to which the rule applies (e.g., 'EXE', 'DLL').")
	ann.Describe(&a.URLCategories, "List of URL categories to which the rule applies.")
	ann.Describe(&a.Protocols, "List of protocols to which the rule applies (e.g., 'FTP_RULE', 'HTTPS_RULE').")
	ann.Describe(&a.CloudApplications, "List of cloud applications to which the rule applies.")
	ann.Describe(&a.MinSize, "Minimum file size in bytes for the rule to apply.")
	ann.Describe(&a.MaxSize, "Maximum file size in bytes for the rule to apply.")
	ann.Describe(&a.BrowserEunTemplateID, "The browser end-user notification template ID.")
	ann.Describe(&a.Locations, "List of location IDs to which the rule applies.")
	ann.Describe(&a.LocationGroups, "List of location group IDs to which the rule applies.")
	ann.Describe(&a.Departments, "List of department IDs to which the rule applies.")
	ann.Describe(&a.Groups, "List of group IDs to which the rule applies.")
	ann.Describe(&a.Users, "List of user IDs to which the rule applies.")
	ann.Describe(&a.TimeWindows, "List of time window IDs during which the rule is active.")
	ann.Describe(&a.Labels, "List of label IDs associated with the rule.")
	ann.Describe(&a.DeviceGroups, "List of device group IDs to which the rule applies.")
	ann.Describe(&a.Devices, "List of device IDs to which the rule applies.")
	ann.Describe(&a.DeviceTrustLevels, "List of device trust levels for the rule.")
	ann.Describe(&a.ZpaAppSegments, "List of ZPA application segments for the rule.")
}

func (s *FileTypeControlRuleState) Annotate(a infer.Annotator) {
	a.Describe(&s.RuleID, "The unique identifier for the file type control rule assigned by the ZIA cloud.")
}

var _ infer.CustomResource[FileTypeControlRuleArgs, FileTypeControlRuleState] = FileTypeControlRule{}
