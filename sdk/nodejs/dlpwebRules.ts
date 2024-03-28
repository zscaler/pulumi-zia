// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The **zia_dlp_web_rules** resource allows the creation and management of ZIA DLP Web Rules in the Zscaler Internet Access cloud or via the API.
 *
 * ⚠️ **WARNING:** Zscaler Internet Access DLP supports a maximum of 127 Web DLP Rules to be created via API.
 *
 * ## Example Usage
 *
 * ### OCR ENABLED
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@bdzscaler/pulumi-zia";
 *
 * const test = new zia.DLPWebRules("test", {
 *     action: "ALLOW",
 *     cloudApplications: [
 *         "ZENDESK",
 *         "LUCKY_ORANGE",
 *         "MICROSOFT_POWERAPPS",
 *         "MICROSOFTLIVEMEETING",
 *     ],
 *     description: "Test",
 *     fileTypes: [
 *         "BITMAP",
 *         "JPEG",
 *         "PNG",
 *         "TIFF",
 *     ],
 *     matchOnly: false,
 *     minSize: 20,
 *     ocrEnabled: true,
 *     order: 1,
 *     protocols: [
 *         "FTP_RULE",
 *         "HTTPS_RULE",
 *         "HTTP_RULE",
 *     ],
 *     rank: 7,
 *     state: "ENABLED",
 *     withoutContentInspection: false,
 *     zscalerIncidentReceiver: true,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### "ALL_OUTBOUND" File Type
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@bdzscaler/pulumi-zia";
 * import * as zia from "@pulumi/zia";
 *
 * const thisDLPEngines = zia.getDLPEngines({
 *     predefinedEngineName: "EXTERNAL",
 * });
 * const thisDLPWebRules = new zia.DLPWebRules("thisDLPWebRules", {
 *     description: "Example",
 *     action: "BLOCK",
 *     order: 1,
 *     rank: 7,
 *     state: "ENABLED",
 *     protocols: [
 *         "FTP_RULE",
 *         "HTTPS_RULE",
 *         "HTTP_RULE",
 *     ],
 *     fileTypes: ["ALL_OUTBOUND"],
 *     zscalerIncidentReceiver: true,
 *     withoutContentInspection: false,
 *     userRiskScoreLevels: [
 *         "LOW",
 *         "MEDIUM",
 *         "HIGH",
 *         "CRITICAL",
 *     ],
 *     severity: "RULE_SEVERITY_HIGH",
 *     dlpEngines: {
 *         ids: [thisDLPEngines.then(thisDLPEngines => thisDLPEngines.id)],
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.
 *
 * Visit
 *
 * **zia_dlp_web_rules** can be imported by using `<RULE ID>` or `<RULE NAME>` as the import ID.
 *
 * For example:
 *
 * ```sh
 * $ pulumi import zia:index/dLPWebRules:DLPWebRules example <rule_id>
 * ```
 *
 * or
 *
 * ```sh
 * $ pulumi import zia:index/dLPWebRules:DLPWebRules example <rule_name>
 * ```
 */
export class DLPWebRules extends pulumi.CustomResource {
    /**
     * Get an existing DLPWebRules resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DLPWebRulesState, opts?: pulumi.CustomResourceOptions): DLPWebRules {
        return new DLPWebRules(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zia:index/dLPWebRules:DLPWebRules';

    /**
     * Returns true if the given object is an instance of DLPWebRules.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DLPWebRules {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DLPWebRules.__pulumiType;
    }

    /**
     * The action taken when traffic matches the DLP policy rule criteria. The supported values are:
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * The auditor to which the DLP policy rule must be applied.
     */
    public readonly auditor!: pulumi.Output<outputs.DLPWebRulesAuditor>;
    /**
     * The list of cloud applications to which the DLP policy rule must be applied.
     */
    public readonly cloudApplications!: pulumi.Output<string[]>;
    /**
     * The name-ID pairs of the departments that are excluded from the DLP policy rule.
     */
    public readonly departments!: pulumi.Output<outputs.DLPWebRulesDepartments>;
    /**
     * The description of the DLP policy rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Indicates a DLP policy rule without content inspection, when the value is set to true.
     */
    public readonly dlpDownloadScanEnabled!: pulumi.Output<boolean>;
    /**
     * The list of DLP engines to which the DLP policy rule must be applied.
     */
    public readonly dlpEngines!: pulumi.Output<outputs.DLPWebRulesDlpEngines>;
    /**
     * The name-ID pairs of the groups that are excluded from the DLP policy rule. Maximum of up to `256` departments.
     */
    public readonly excludedDepartments!: pulumi.Output<outputs.DLPWebRulesExcludedDepartments>;
    /**
     * The name-ID pairs of the groups that are excluded from the DLP policy rule. Maximum of up to `256` groups.
     */
    public readonly excludedGroups!: pulumi.Output<outputs.DLPWebRulesExcludedGroups>;
    /**
     * The name-ID pairs of the users that are excluded from the DLP policy rule. Maximum of up to `256` users.
     */
    public readonly excludedUsers!: pulumi.Output<outputs.DLPWebRulesExcludedUsers>;
    /**
     * The email address of an external auditor to whom DLP email notifications are sent.
     */
    public readonly externalAuditorEmail!: pulumi.Output<string>;
    /**
     * The list of file types to which the DLP policy rule must be applied. For the complete list of supported file types refer to the  [ZIA API documentation](https://help.zscaler.com/zia/data-loss-prevention#/webDlpRules-post)
     *
     * * > Note: `BITMAP`, `JPEG`, `PNG`, and `TIFF` file types are exclusively supported when optical character recognition `ocrEnabled` is set to `true` for DLP rules with content inspection.
     *
     * * > Note: `ALL_OUTBOUND` file type is applicable only when the predefined DLP engine called `EXTERNAL` is used and when the attribute `withoutContentInspection` is set to `false`.
     *
     * * > Note: `ALL_OUTBOUND` file type cannot be used alongside any any other file type.
     */
    public readonly fileTypes!: pulumi.Output<string[]>;
    /**
     * The Name-ID pairs of groups to which the DLP policy rule must be applied. Maximum of up to `8` groups. When not used it implies `Any` to apply the rule to all groups.
     */
    public readonly groups!: pulumi.Output<outputs.DLPWebRulesGroups>;
    /**
     * The DLP server, using ICAP, to which the transaction content is forwarded.
     */
    public readonly icapServer!: pulumi.Output<outputs.DLPWebRulesIcapServer>;
    /**
     * The Name-ID pairs of rule labels associated to the DLP policy rule.
     */
    public readonly labels!: pulumi.Output<outputs.DLPWebRulesLabels>;
    /**
     * The Name-ID pairs of locations groups to which the DLP policy rule must be applied. Maximum of up to `32` location groups. When not used it implies `Any` to apply the rule to all location groups.
     */
    public readonly locationGroups!: pulumi.Output<outputs.DLPWebRulesLocationGroups>;
    /**
     * The Name-ID pairs of locations to which the DLP policy rule must be applied. Maximum of up to `8` locations. When not used it implies `Any` to apply the rule to all locations.
     */
    public readonly locations!: pulumi.Output<outputs.DLPWebRulesLocations>;
    /**
     * The match only criteria for DLP engines.
     */
    public readonly matchOnly!: pulumi.Output<boolean>;
    /**
     * The minimum file size (in KB) used for evaluation of the DLP policy rule.
     */
    public readonly minSize!: pulumi.Output<number>;
    /**
     * The name of the workload group
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The template used for DLP notification emails.
     */
    public readonly notificationTemplate!: pulumi.Output<outputs.DLPWebRulesNotificationTemplate>;
    /**
     * Enables or disables image file scanning. When OCR is enabled only the following ``fileTypes`` are supported: ``WINDOWS_META_FORMAT``, ``BITMAP``, ``JPEG``, ``PNG``, ``TIFF``
     */
    public readonly ocrEnabled!: pulumi.Output<boolean>;
    /**
     * The rule order of execution for the DLP policy rule with respect to other rules.
     */
    public readonly order!: pulumi.Output<number>;
    /**
     * The unique identifier of the parent rule under which an exception rule is added.
     * > Note: Exception rules can be configured only when the inline DLP rule evaluation type is set to evaluate all DLP rules in the DLP Advanced Settings.
     */
    public readonly parentRule!: pulumi.Output<number>;
    /**
     * The protocol criteria specified for the DLP policy rule.
     */
    public readonly protocols!: pulumi.Output<string[]>;
    /**
     * Admin rank of the admin who creates this rule
     */
    public readonly rank!: pulumi.Output<number | undefined>;
    public /*out*/ readonly ruleId!: pulumi.Output<number>;
    /**
     * Indicates the severity selected for the DLP rule violation: Returned values are:  `RULE_SEVERITY_HIGH`, `RULE_SEVERITY_MEDIUM`, `RULE_SEVERITY_LOW`, `RULE_SEVERITY_INFO`
     */
    public readonly severity!: pulumi.Output<string>;
    /**
     * Enables or disables the DLP policy rule.. The supported values are:
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * The list of exception rules added to a parent rule.
     * > Note: All attributes within the WebDlpRule model are applicable to the sub-rules. Values for each rule are specified by using the WebDlpRule object Exception rules can be configured only when the inline DLP rule evaluation type is set to evaluate all DLP rules in the DLP Advanced Settings.
     */
    public readonly subRules!: pulumi.Output<string[]>;
    /**
     * The Name-ID pairs of time windows to which the DLP policy rule must be applied. Maximum of up to `2` time intervals. When not used it implies `always` to apply the rule to all time intervals.
     */
    public readonly timeWindows!: pulumi.Output<outputs.DLPWebRulesTimeWindows>;
    /**
     * The list of URL categories to which the DLP policy rule must be applied.
     */
    public readonly urlCategories!: pulumi.Output<outputs.DLPWebRulesUrlCategories>;
    /**
     * Indicates the user risk score level selectedd for the DLP rule violation: Returned values are: `LOW`, `MEDIUM`, `HIGH`, `CRITICAL`
     */
    public readonly userRiskScoreLevels!: pulumi.Output<string[]>;
    /**
     * The Name-ID pairs of users to which the DLP policy rule must be applied. Maximum of up to `4` users. When not used it implies `Any` to apply the rule to all users.
     */
    public readonly users!: pulumi.Output<outputs.DLPWebRulesUsers>;
    /**
     * must be set to false if `fileTypes` is not defined.
     */
    public readonly withoutContentInspection!: pulumi.Output<boolean>;
    /**
     * The list of preconfigured workload groups to which the policy must be applied
     */
    public readonly workloadGroups!: pulumi.Output<outputs.DLPWebRulesWorkloadGroup[]>;
    /**
     * Indicates a DLP policy rule without content inspection, when the value is set to true.
     */
    public readonly zccNotificationsEnabled!: pulumi.Output<boolean>;
    /**
     * Indicates whether a Zscaler Incident Receiver is associated to the DLP policy rule.
     */
    public readonly zscalerIncidentReceiver!: pulumi.Output<boolean>;

    /**
     * Create a DLPWebRules resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DLPWebRulesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DLPWebRulesArgs | DLPWebRulesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DLPWebRulesState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["auditor"] = state ? state.auditor : undefined;
            resourceInputs["cloudApplications"] = state ? state.cloudApplications : undefined;
            resourceInputs["departments"] = state ? state.departments : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dlpDownloadScanEnabled"] = state ? state.dlpDownloadScanEnabled : undefined;
            resourceInputs["dlpEngines"] = state ? state.dlpEngines : undefined;
            resourceInputs["excludedDepartments"] = state ? state.excludedDepartments : undefined;
            resourceInputs["excludedGroups"] = state ? state.excludedGroups : undefined;
            resourceInputs["excludedUsers"] = state ? state.excludedUsers : undefined;
            resourceInputs["externalAuditorEmail"] = state ? state.externalAuditorEmail : undefined;
            resourceInputs["fileTypes"] = state ? state.fileTypes : undefined;
            resourceInputs["groups"] = state ? state.groups : undefined;
            resourceInputs["icapServer"] = state ? state.icapServer : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["locationGroups"] = state ? state.locationGroups : undefined;
            resourceInputs["locations"] = state ? state.locations : undefined;
            resourceInputs["matchOnly"] = state ? state.matchOnly : undefined;
            resourceInputs["minSize"] = state ? state.minSize : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["notificationTemplate"] = state ? state.notificationTemplate : undefined;
            resourceInputs["ocrEnabled"] = state ? state.ocrEnabled : undefined;
            resourceInputs["order"] = state ? state.order : undefined;
            resourceInputs["parentRule"] = state ? state.parentRule : undefined;
            resourceInputs["protocols"] = state ? state.protocols : undefined;
            resourceInputs["rank"] = state ? state.rank : undefined;
            resourceInputs["ruleId"] = state ? state.ruleId : undefined;
            resourceInputs["severity"] = state ? state.severity : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["subRules"] = state ? state.subRules : undefined;
            resourceInputs["timeWindows"] = state ? state.timeWindows : undefined;
            resourceInputs["urlCategories"] = state ? state.urlCategories : undefined;
            resourceInputs["userRiskScoreLevels"] = state ? state.userRiskScoreLevels : undefined;
            resourceInputs["users"] = state ? state.users : undefined;
            resourceInputs["withoutContentInspection"] = state ? state.withoutContentInspection : undefined;
            resourceInputs["workloadGroups"] = state ? state.workloadGroups : undefined;
            resourceInputs["zccNotificationsEnabled"] = state ? state.zccNotificationsEnabled : undefined;
            resourceInputs["zscalerIncidentReceiver"] = state ? state.zscalerIncidentReceiver : undefined;
        } else {
            const args = argsOrState as DLPWebRulesArgs | undefined;
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["auditor"] = args ? args.auditor : undefined;
            resourceInputs["cloudApplications"] = args ? args.cloudApplications : undefined;
            resourceInputs["departments"] = args ? args.departments : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dlpDownloadScanEnabled"] = args ? args.dlpDownloadScanEnabled : undefined;
            resourceInputs["dlpEngines"] = args ? args.dlpEngines : undefined;
            resourceInputs["excludedDepartments"] = args ? args.excludedDepartments : undefined;
            resourceInputs["excludedGroups"] = args ? args.excludedGroups : undefined;
            resourceInputs["excludedUsers"] = args ? args.excludedUsers : undefined;
            resourceInputs["externalAuditorEmail"] = args ? args.externalAuditorEmail : undefined;
            resourceInputs["fileTypes"] = args ? args.fileTypes : undefined;
            resourceInputs["groups"] = args ? args.groups : undefined;
            resourceInputs["icapServer"] = args ? args.icapServer : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["locationGroups"] = args ? args.locationGroups : undefined;
            resourceInputs["locations"] = args ? args.locations : undefined;
            resourceInputs["matchOnly"] = args ? args.matchOnly : undefined;
            resourceInputs["minSize"] = args ? args.minSize : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notificationTemplate"] = args ? args.notificationTemplate : undefined;
            resourceInputs["ocrEnabled"] = args ? args.ocrEnabled : undefined;
            resourceInputs["order"] = args ? args.order : undefined;
            resourceInputs["parentRule"] = args ? args.parentRule : undefined;
            resourceInputs["protocols"] = args ? args.protocols : undefined;
            resourceInputs["rank"] = args ? args.rank : undefined;
            resourceInputs["severity"] = args ? args.severity : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["subRules"] = args ? args.subRules : undefined;
            resourceInputs["timeWindows"] = args ? args.timeWindows : undefined;
            resourceInputs["urlCategories"] = args ? args.urlCategories : undefined;
            resourceInputs["userRiskScoreLevels"] = args ? args.userRiskScoreLevels : undefined;
            resourceInputs["users"] = args ? args.users : undefined;
            resourceInputs["withoutContentInspection"] = args ? args.withoutContentInspection : undefined;
            resourceInputs["workloadGroups"] = args ? args.workloadGroups : undefined;
            resourceInputs["zccNotificationsEnabled"] = args ? args.zccNotificationsEnabled : undefined;
            resourceInputs["zscalerIncidentReceiver"] = args ? args.zscalerIncidentReceiver : undefined;
            resourceInputs["ruleId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DLPWebRules.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DLPWebRules resources.
 */
export interface DLPWebRulesState {
    /**
     * The action taken when traffic matches the DLP policy rule criteria. The supported values are:
     */
    action?: pulumi.Input<string>;
    /**
     * The auditor to which the DLP policy rule must be applied.
     */
    auditor?: pulumi.Input<inputs.DLPWebRulesAuditor>;
    /**
     * The list of cloud applications to which the DLP policy rule must be applied.
     */
    cloudApplications?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name-ID pairs of the departments that are excluded from the DLP policy rule.
     */
    departments?: pulumi.Input<inputs.DLPWebRulesDepartments>;
    /**
     * The description of the DLP policy rule.
     */
    description?: pulumi.Input<string>;
    /**
     * Indicates a DLP policy rule without content inspection, when the value is set to true.
     */
    dlpDownloadScanEnabled?: pulumi.Input<boolean>;
    /**
     * The list of DLP engines to which the DLP policy rule must be applied.
     */
    dlpEngines?: pulumi.Input<inputs.DLPWebRulesDlpEngines>;
    /**
     * The name-ID pairs of the groups that are excluded from the DLP policy rule. Maximum of up to `256` departments.
     */
    excludedDepartments?: pulumi.Input<inputs.DLPWebRulesExcludedDepartments>;
    /**
     * The name-ID pairs of the groups that are excluded from the DLP policy rule. Maximum of up to `256` groups.
     */
    excludedGroups?: pulumi.Input<inputs.DLPWebRulesExcludedGroups>;
    /**
     * The name-ID pairs of the users that are excluded from the DLP policy rule. Maximum of up to `256` users.
     */
    excludedUsers?: pulumi.Input<inputs.DLPWebRulesExcludedUsers>;
    /**
     * The email address of an external auditor to whom DLP email notifications are sent.
     */
    externalAuditorEmail?: pulumi.Input<string>;
    /**
     * The list of file types to which the DLP policy rule must be applied. For the complete list of supported file types refer to the  [ZIA API documentation](https://help.zscaler.com/zia/data-loss-prevention#/webDlpRules-post)
     *
     * * > Note: `BITMAP`, `JPEG`, `PNG`, and `TIFF` file types are exclusively supported when optical character recognition `ocrEnabled` is set to `true` for DLP rules with content inspection.
     *
     * * > Note: `ALL_OUTBOUND` file type is applicable only when the predefined DLP engine called `EXTERNAL` is used and when the attribute `withoutContentInspection` is set to `false`.
     *
     * * > Note: `ALL_OUTBOUND` file type cannot be used alongside any any other file type.
     */
    fileTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Name-ID pairs of groups to which the DLP policy rule must be applied. Maximum of up to `8` groups. When not used it implies `Any` to apply the rule to all groups.
     */
    groups?: pulumi.Input<inputs.DLPWebRulesGroups>;
    /**
     * The DLP server, using ICAP, to which the transaction content is forwarded.
     */
    icapServer?: pulumi.Input<inputs.DLPWebRulesIcapServer>;
    /**
     * The Name-ID pairs of rule labels associated to the DLP policy rule.
     */
    labels?: pulumi.Input<inputs.DLPWebRulesLabels>;
    /**
     * The Name-ID pairs of locations groups to which the DLP policy rule must be applied. Maximum of up to `32` location groups. When not used it implies `Any` to apply the rule to all location groups.
     */
    locationGroups?: pulumi.Input<inputs.DLPWebRulesLocationGroups>;
    /**
     * The Name-ID pairs of locations to which the DLP policy rule must be applied. Maximum of up to `8` locations. When not used it implies `Any` to apply the rule to all locations.
     */
    locations?: pulumi.Input<inputs.DLPWebRulesLocations>;
    /**
     * The match only criteria for DLP engines.
     */
    matchOnly?: pulumi.Input<boolean>;
    /**
     * The minimum file size (in KB) used for evaluation of the DLP policy rule.
     */
    minSize?: pulumi.Input<number>;
    /**
     * The name of the workload group
     */
    name?: pulumi.Input<string>;
    /**
     * The template used for DLP notification emails.
     */
    notificationTemplate?: pulumi.Input<inputs.DLPWebRulesNotificationTemplate>;
    /**
     * Enables or disables image file scanning. When OCR is enabled only the following ``fileTypes`` are supported: ``WINDOWS_META_FORMAT``, ``BITMAP``, ``JPEG``, ``PNG``, ``TIFF``
     */
    ocrEnabled?: pulumi.Input<boolean>;
    /**
     * The rule order of execution for the DLP policy rule with respect to other rules.
     */
    order?: pulumi.Input<number>;
    /**
     * The unique identifier of the parent rule under which an exception rule is added.
     * > Note: Exception rules can be configured only when the inline DLP rule evaluation type is set to evaluate all DLP rules in the DLP Advanced Settings.
     */
    parentRule?: pulumi.Input<number>;
    /**
     * The protocol criteria specified for the DLP policy rule.
     */
    protocols?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Admin rank of the admin who creates this rule
     */
    rank?: pulumi.Input<number>;
    ruleId?: pulumi.Input<number>;
    /**
     * Indicates the severity selected for the DLP rule violation: Returned values are:  `RULE_SEVERITY_HIGH`, `RULE_SEVERITY_MEDIUM`, `RULE_SEVERITY_LOW`, `RULE_SEVERITY_INFO`
     */
    severity?: pulumi.Input<string>;
    /**
     * Enables or disables the DLP policy rule.. The supported values are:
     */
    state?: pulumi.Input<string>;
    /**
     * The list of exception rules added to a parent rule.
     * > Note: All attributes within the WebDlpRule model are applicable to the sub-rules. Values for each rule are specified by using the WebDlpRule object Exception rules can be configured only when the inline DLP rule evaluation type is set to evaluate all DLP rules in the DLP Advanced Settings.
     */
    subRules?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Name-ID pairs of time windows to which the DLP policy rule must be applied. Maximum of up to `2` time intervals. When not used it implies `always` to apply the rule to all time intervals.
     */
    timeWindows?: pulumi.Input<inputs.DLPWebRulesTimeWindows>;
    /**
     * The list of URL categories to which the DLP policy rule must be applied.
     */
    urlCategories?: pulumi.Input<inputs.DLPWebRulesUrlCategories>;
    /**
     * Indicates the user risk score level selectedd for the DLP rule violation: Returned values are: `LOW`, `MEDIUM`, `HIGH`, `CRITICAL`
     */
    userRiskScoreLevels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Name-ID pairs of users to which the DLP policy rule must be applied. Maximum of up to `4` users. When not used it implies `Any` to apply the rule to all users.
     */
    users?: pulumi.Input<inputs.DLPWebRulesUsers>;
    /**
     * must be set to false if `fileTypes` is not defined.
     */
    withoutContentInspection?: pulumi.Input<boolean>;
    /**
     * The list of preconfigured workload groups to which the policy must be applied
     */
    workloadGroups?: pulumi.Input<pulumi.Input<inputs.DLPWebRulesWorkloadGroup>[]>;
    /**
     * Indicates a DLP policy rule without content inspection, when the value is set to true.
     */
    zccNotificationsEnabled?: pulumi.Input<boolean>;
    /**
     * Indicates whether a Zscaler Incident Receiver is associated to the DLP policy rule.
     */
    zscalerIncidentReceiver?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a DLPWebRules resource.
 */
export interface DLPWebRulesArgs {
    /**
     * The action taken when traffic matches the DLP policy rule criteria. The supported values are:
     */
    action?: pulumi.Input<string>;
    /**
     * The auditor to which the DLP policy rule must be applied.
     */
    auditor?: pulumi.Input<inputs.DLPWebRulesAuditor>;
    /**
     * The list of cloud applications to which the DLP policy rule must be applied.
     */
    cloudApplications?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name-ID pairs of the departments that are excluded from the DLP policy rule.
     */
    departments?: pulumi.Input<inputs.DLPWebRulesDepartments>;
    /**
     * The description of the DLP policy rule.
     */
    description?: pulumi.Input<string>;
    /**
     * Indicates a DLP policy rule without content inspection, when the value is set to true.
     */
    dlpDownloadScanEnabled?: pulumi.Input<boolean>;
    /**
     * The list of DLP engines to which the DLP policy rule must be applied.
     */
    dlpEngines?: pulumi.Input<inputs.DLPWebRulesDlpEngines>;
    /**
     * The name-ID pairs of the groups that are excluded from the DLP policy rule. Maximum of up to `256` departments.
     */
    excludedDepartments?: pulumi.Input<inputs.DLPWebRulesExcludedDepartments>;
    /**
     * The name-ID pairs of the groups that are excluded from the DLP policy rule. Maximum of up to `256` groups.
     */
    excludedGroups?: pulumi.Input<inputs.DLPWebRulesExcludedGroups>;
    /**
     * The name-ID pairs of the users that are excluded from the DLP policy rule. Maximum of up to `256` users.
     */
    excludedUsers?: pulumi.Input<inputs.DLPWebRulesExcludedUsers>;
    /**
     * The email address of an external auditor to whom DLP email notifications are sent.
     */
    externalAuditorEmail?: pulumi.Input<string>;
    /**
     * The list of file types to which the DLP policy rule must be applied. For the complete list of supported file types refer to the  [ZIA API documentation](https://help.zscaler.com/zia/data-loss-prevention#/webDlpRules-post)
     *
     * * > Note: `BITMAP`, `JPEG`, `PNG`, and `TIFF` file types are exclusively supported when optical character recognition `ocrEnabled` is set to `true` for DLP rules with content inspection.
     *
     * * > Note: `ALL_OUTBOUND` file type is applicable only when the predefined DLP engine called `EXTERNAL` is used and when the attribute `withoutContentInspection` is set to `false`.
     *
     * * > Note: `ALL_OUTBOUND` file type cannot be used alongside any any other file type.
     */
    fileTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Name-ID pairs of groups to which the DLP policy rule must be applied. Maximum of up to `8` groups. When not used it implies `Any` to apply the rule to all groups.
     */
    groups?: pulumi.Input<inputs.DLPWebRulesGroups>;
    /**
     * The DLP server, using ICAP, to which the transaction content is forwarded.
     */
    icapServer?: pulumi.Input<inputs.DLPWebRulesIcapServer>;
    /**
     * The Name-ID pairs of rule labels associated to the DLP policy rule.
     */
    labels?: pulumi.Input<inputs.DLPWebRulesLabels>;
    /**
     * The Name-ID pairs of locations groups to which the DLP policy rule must be applied. Maximum of up to `32` location groups. When not used it implies `Any` to apply the rule to all location groups.
     */
    locationGroups?: pulumi.Input<inputs.DLPWebRulesLocationGroups>;
    /**
     * The Name-ID pairs of locations to which the DLP policy rule must be applied. Maximum of up to `8` locations. When not used it implies `Any` to apply the rule to all locations.
     */
    locations?: pulumi.Input<inputs.DLPWebRulesLocations>;
    /**
     * The match only criteria for DLP engines.
     */
    matchOnly?: pulumi.Input<boolean>;
    /**
     * The minimum file size (in KB) used for evaluation of the DLP policy rule.
     */
    minSize?: pulumi.Input<number>;
    /**
     * The name of the workload group
     */
    name?: pulumi.Input<string>;
    /**
     * The template used for DLP notification emails.
     */
    notificationTemplate?: pulumi.Input<inputs.DLPWebRulesNotificationTemplate>;
    /**
     * Enables or disables image file scanning. When OCR is enabled only the following ``fileTypes`` are supported: ``WINDOWS_META_FORMAT``, ``BITMAP``, ``JPEG``, ``PNG``, ``TIFF``
     */
    ocrEnabled?: pulumi.Input<boolean>;
    /**
     * The rule order of execution for the DLP policy rule with respect to other rules.
     */
    order?: pulumi.Input<number>;
    /**
     * The unique identifier of the parent rule under which an exception rule is added.
     * > Note: Exception rules can be configured only when the inline DLP rule evaluation type is set to evaluate all DLP rules in the DLP Advanced Settings.
     */
    parentRule?: pulumi.Input<number>;
    /**
     * The protocol criteria specified for the DLP policy rule.
     */
    protocols?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Admin rank of the admin who creates this rule
     */
    rank?: pulumi.Input<number>;
    /**
     * Indicates the severity selected for the DLP rule violation: Returned values are:  `RULE_SEVERITY_HIGH`, `RULE_SEVERITY_MEDIUM`, `RULE_SEVERITY_LOW`, `RULE_SEVERITY_INFO`
     */
    severity?: pulumi.Input<string>;
    /**
     * Enables or disables the DLP policy rule.. The supported values are:
     */
    state?: pulumi.Input<string>;
    /**
     * The list of exception rules added to a parent rule.
     * > Note: All attributes within the WebDlpRule model are applicable to the sub-rules. Values for each rule are specified by using the WebDlpRule object Exception rules can be configured only when the inline DLP rule evaluation type is set to evaluate all DLP rules in the DLP Advanced Settings.
     */
    subRules?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Name-ID pairs of time windows to which the DLP policy rule must be applied. Maximum of up to `2` time intervals. When not used it implies `always` to apply the rule to all time intervals.
     */
    timeWindows?: pulumi.Input<inputs.DLPWebRulesTimeWindows>;
    /**
     * The list of URL categories to which the DLP policy rule must be applied.
     */
    urlCategories?: pulumi.Input<inputs.DLPWebRulesUrlCategories>;
    /**
     * Indicates the user risk score level selectedd for the DLP rule violation: Returned values are: `LOW`, `MEDIUM`, `HIGH`, `CRITICAL`
     */
    userRiskScoreLevels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Name-ID pairs of users to which the DLP policy rule must be applied. Maximum of up to `4` users. When not used it implies `Any` to apply the rule to all users.
     */
    users?: pulumi.Input<inputs.DLPWebRulesUsers>;
    /**
     * must be set to false if `fileTypes` is not defined.
     */
    withoutContentInspection?: pulumi.Input<boolean>;
    /**
     * The list of preconfigured workload groups to which the policy must be applied
     */
    workloadGroups?: pulumi.Input<pulumi.Input<inputs.DLPWebRulesWorkloadGroup>[]>;
    /**
     * Indicates a DLP policy rule without content inspection, when the value is set to true.
     */
    zccNotificationsEnabled?: pulumi.Input<boolean>;
    /**
     * Indicates whether a Zscaler Incident Receiver is associated to the DLP policy rule.
     */
    zscalerIncidentReceiver?: pulumi.Input<boolean>;
}
