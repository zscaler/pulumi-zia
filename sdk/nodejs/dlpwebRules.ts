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
 * ### "ALL_OUTBOUND" File Type"
 *
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
 *     zscalerIncidentReceiver: false,
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
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@bdzscaler/pulumi-zia";
 * import * as zia from "@pulumi/zia";
 *
 * const thisURLCategories = zia.getURLCategories({
 *     configuredName: "Example",
 * });
 * const thisIcapServers = zia.getIcapServers({
 *     name: "ZS_ICAP_01",
 * });
 * const thisDLPWebRules = new zia.DLPWebRules("thisDLPWebRules", {
 *     description: "Terraform_Test",
 *     action: "BLOCK",
 *     order: 1,
 *     protocols: [
 *         "FTP_RULE",
 *         "HTTPS_RULE",
 *         "HTTP_RULE",
 *     ],
 *     rank: 7,
 *     state: "ENABLED",
 *     zscalerIncidentReceiver: true,
 *     withoutContentInspection: false,
 *     urlCategories: {
 *         ids: [thisURLCategories.then(thisURLCategories => thisURLCategories.val)],
 *     },
 *     icapServers: [{
 *         id: thisIcapServers.then(thisIcapServers => thisIcapServers.id),
 *     }],
 * });
 * ```
 *
 * ### "Specify Incident Receiver Setting"
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@bdzscaler/pulumi-zia";
 * import * as zia from "@pulumi/zia";
 *
 * const thisURLCategories = zia.getURLCategories({
 *     configuredName: "Example",
 * });
 * const thisDLPIncidentReceiverServers = zia.getDLPIncidentReceiverServers({
 *     name: "ZS_INC_RECEIVER_01",
 * });
 * const thisDLPWebRules = new zia.DLPWebRules("thisDLPWebRules", {
 *     description: "Terraform_Test",
 *     action: "BLOCK",
 *     order: 1,
 *     protocols: [
 *         "FTP_RULE",
 *         "HTTPS_RULE",
 *         "HTTP_RULE",
 *     ],
 *     rank: 7,
 *     state: "ENABLED",
 *     zscalerIncidentReceiver: true,
 *     withoutContentInspection: false,
 *     urlCategories: {
 *         ids: [thisURLCategories.then(thisURLCategories => thisURLCategories.val)],
 *     },
 *     icapServers: [{
 *         id: thisDLPIncidentReceiverServers.then(thisDLPIncidentReceiverServers => thisDLPIncidentReceiverServers.id),
 *     }],
 * });
 * ```
 *
 * ### "Creating Parent Rules And SubRules"
 *
 * ⚠️ **WARNING:** Destroying a parent rule will also destroy all subrules
 *
 *  **NOTE** Exception rules can be configured only when the inline DLP rule evaluation type is set
 *  to evaluate all DLP rules in the DLP Advanced Settings.
 *  To learn more, see [Configuring DLP Advanced Settings](https://help.zscaler.com/%22/zia/configuring-dlp-advanced-settings/%22)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@bdzscaler/pulumi-zia";
 *
 * const parentRule = new zia.DLPWebRules("parentRule", {
 *     description: "ParentRule1",
 *     action: "ALLOW",
 *     state: "ENABLED",
 *     order: 1,
 *     rank: 0,
 *     protocols: [
 *         "FTP_RULE",
 *         "HTTPS_RULE",
 *         "HTTP_RULE",
 *     ],
 *     cloudApplications: [
 *         "GOOGLE_WEBMAIL",
 *         "WINDOWS_LIVE_HOTMAIL",
 *     ],
 *     withoutContentInspection: false,
 *     matchOnly: false,
 *     minSize: 20,
 *     zscalerIncidentReceiver: true,
 * });
 * const subrule1 = new zia.DLPWebRules("subrule1", {
 *     description: "SubRule1",
 *     action: "ALLOW",
 *     state: "ENABLED",
 *     order: 1,
 *     rank: 0,
 *     protocols: [
 *         "FTP_RULE",
 *         "HTTPS_RULE",
 *         "HTTP_RULE",
 *     ],
 *     cloudApplications: [
 *         "GOOGLE_WEBMAIL",
 *         "WINDOWS_LIVE_HOTMAIL",
 *     ],
 *     withoutContentInspection: false,
 *     matchOnly: false,
 *     parentRule: parentRule.id,
 * });
 * ```
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
     * The action taken when traffic matches the DLP policy rule criteria.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * The auditor to which the DLP policy rule must be applied.
     */
    public readonly auditors!: pulumi.Output<outputs.DLPWebRulesAuditor[]>;
    /**
     * The list of cloud applications to which the DLP policy rule must be applied.
     */
    public readonly cloudApplications!: pulumi.Output<string[]>;
    /**
     * The Name-ID pairs of departments to which the DLP policy rule must be applied.
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
     * The Name-ID pairs of users to which the DLP policy rule must be applied.
     */
    public readonly excludedDepartments!: pulumi.Output<outputs.DLPWebRulesExcludedDepartments>;
    /**
     * The Name-ID pairs of users to which the DLP policy rule must be applied.
     */
    public readonly excludedDomainProfiles!: pulumi.Output<outputs.DLPWebRulesExcludedDomainProfiles>;
    /**
     * The Name-ID pairs of users to which the DLP policy rule must be applied.
     */
    public readonly excludedGroups!: pulumi.Output<outputs.DLPWebRulesExcludedGroups>;
    /**
     * The Name-ID pairs of users to which the DLP policy rule must be applied.
     */
    public readonly excludedUsers!: pulumi.Output<outputs.DLPWebRulesExcludedUsers>;
    /**
     * The email address of an external auditor to whom DLP email notifications are sent.
     */
    public readonly externalAuditorEmail!: pulumi.Output<string>;
    /**
     * The list of file types for which the DLP policy rule must be applied.
     */
    public readonly fileTypes!: pulumi.Output<string[]>;
    /**
     * The Name-ID pairs of groups to which the DLP policy rule must be applied.
     */
    public readonly groups!: pulumi.Output<outputs.DLPWebRulesGroups>;
    /**
     * The DLP server, using ICAP, to which the transaction content is forwarded.
     */
    public readonly icapServers!: pulumi.Output<outputs.DLPWebRulesIcapServer[]>;
    /**
     * The Name-ID pairs of users to which the DLP policy rule must be applied.
     */
    public readonly includedDomainProfiles!: pulumi.Output<outputs.DLPWebRulesIncludedDomainProfiles>;
    /**
     * list of Labels that are applicable to the rule.
     */
    public readonly labels!: pulumi.Output<outputs.DLPWebRulesLabels>;
    /**
     * The Name-ID pairs of locations groups to which the DLP policy rule must be applied.
     */
    public readonly locationGroups!: pulumi.Output<outputs.DLPWebRulesLocationGroups>;
    /**
     * The Name-ID pairs of locations to which the DLP policy rule must be applied.
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
     * The DLP policy rule name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The template used for DLP notification emails.
     */
    public readonly notificationTemplates!: pulumi.Output<outputs.DLPWebRulesNotificationTemplate[]>;
    /**
     * The rule order of execution for the DLP policy rule with respect to other rules.
     */
    public readonly order!: pulumi.Output<number>;
    /**
     * The unique identifier of the parent rule under which an exception rule is added
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
     * Indicates the severity selected for the DLP rule violation
     */
    public readonly severity!: pulumi.Output<string>;
    /**
     * list of source ip groups
     */
    public readonly sourceIpGroups!: pulumi.Output<outputs.DLPWebRulesSourceIpGroups>;
    /**
     * Enables or disables the DLP policy rule.
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * The list of exception rules added to a parent rule
     */
    public readonly subRules!: pulumi.Output<string[]>;
    /**
     * list of time interval during which rule must be enforced.
     */
    public readonly timeWindows!: pulumi.Output<outputs.DLPWebRulesTimeWindows>;
    /**
     * The list of URL categories to which the DLP policy rule must be applied.
     */
    public readonly urlCategories!: pulumi.Output<outputs.DLPWebRulesUrlCategories>;
    public readonly userRiskScoreLevels!: pulumi.Output<string[]>;
    /**
     * The Name-ID pairs of users to which the DLP policy rule must be applied.
     */
    public readonly users!: pulumi.Output<outputs.DLPWebRulesUsers>;
    /**
     * Indicates a DLP policy rule without content inspection, when the value is set to true.
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
    public readonly zscalerIncidentReceiver!: pulumi.Output<boolean | undefined>;

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
            resourceInputs["auditors"] = state ? state.auditors : undefined;
            resourceInputs["cloudApplications"] = state ? state.cloudApplications : undefined;
            resourceInputs["departments"] = state ? state.departments : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dlpDownloadScanEnabled"] = state ? state.dlpDownloadScanEnabled : undefined;
            resourceInputs["dlpEngines"] = state ? state.dlpEngines : undefined;
            resourceInputs["excludedDepartments"] = state ? state.excludedDepartments : undefined;
            resourceInputs["excludedDomainProfiles"] = state ? state.excludedDomainProfiles : undefined;
            resourceInputs["excludedGroups"] = state ? state.excludedGroups : undefined;
            resourceInputs["excludedUsers"] = state ? state.excludedUsers : undefined;
            resourceInputs["externalAuditorEmail"] = state ? state.externalAuditorEmail : undefined;
            resourceInputs["fileTypes"] = state ? state.fileTypes : undefined;
            resourceInputs["groups"] = state ? state.groups : undefined;
            resourceInputs["icapServers"] = state ? state.icapServers : undefined;
            resourceInputs["includedDomainProfiles"] = state ? state.includedDomainProfiles : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["locationGroups"] = state ? state.locationGroups : undefined;
            resourceInputs["locations"] = state ? state.locations : undefined;
            resourceInputs["matchOnly"] = state ? state.matchOnly : undefined;
            resourceInputs["minSize"] = state ? state.minSize : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["notificationTemplates"] = state ? state.notificationTemplates : undefined;
            resourceInputs["order"] = state ? state.order : undefined;
            resourceInputs["parentRule"] = state ? state.parentRule : undefined;
            resourceInputs["protocols"] = state ? state.protocols : undefined;
            resourceInputs["rank"] = state ? state.rank : undefined;
            resourceInputs["ruleId"] = state ? state.ruleId : undefined;
            resourceInputs["severity"] = state ? state.severity : undefined;
            resourceInputs["sourceIpGroups"] = state ? state.sourceIpGroups : undefined;
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
            resourceInputs["auditors"] = args ? args.auditors : undefined;
            resourceInputs["cloudApplications"] = args ? args.cloudApplications : undefined;
            resourceInputs["departments"] = args ? args.departments : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dlpDownloadScanEnabled"] = args ? args.dlpDownloadScanEnabled : undefined;
            resourceInputs["dlpEngines"] = args ? args.dlpEngines : undefined;
            resourceInputs["excludedDepartments"] = args ? args.excludedDepartments : undefined;
            resourceInputs["excludedDomainProfiles"] = args ? args.excludedDomainProfiles : undefined;
            resourceInputs["excludedGroups"] = args ? args.excludedGroups : undefined;
            resourceInputs["excludedUsers"] = args ? args.excludedUsers : undefined;
            resourceInputs["externalAuditorEmail"] = args ? args.externalAuditorEmail : undefined;
            resourceInputs["fileTypes"] = args ? args.fileTypes : undefined;
            resourceInputs["groups"] = args ? args.groups : undefined;
            resourceInputs["icapServers"] = args ? args.icapServers : undefined;
            resourceInputs["includedDomainProfiles"] = args ? args.includedDomainProfiles : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["locationGroups"] = args ? args.locationGroups : undefined;
            resourceInputs["locations"] = args ? args.locations : undefined;
            resourceInputs["matchOnly"] = args ? args.matchOnly : undefined;
            resourceInputs["minSize"] = args ? args.minSize : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notificationTemplates"] = args ? args.notificationTemplates : undefined;
            resourceInputs["order"] = args ? args.order : undefined;
            resourceInputs["parentRule"] = args ? args.parentRule : undefined;
            resourceInputs["protocols"] = args ? args.protocols : undefined;
            resourceInputs["rank"] = args ? args.rank : undefined;
            resourceInputs["severity"] = args ? args.severity : undefined;
            resourceInputs["sourceIpGroups"] = args ? args.sourceIpGroups : undefined;
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
     * The action taken when traffic matches the DLP policy rule criteria.
     */
    action?: pulumi.Input<string>;
    /**
     * The auditor to which the DLP policy rule must be applied.
     */
    auditors?: pulumi.Input<pulumi.Input<inputs.DLPWebRulesAuditor>[]>;
    /**
     * The list of cloud applications to which the DLP policy rule must be applied.
     */
    cloudApplications?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Name-ID pairs of departments to which the DLP policy rule must be applied.
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
     * The Name-ID pairs of users to which the DLP policy rule must be applied.
     */
    excludedDepartments?: pulumi.Input<inputs.DLPWebRulesExcludedDepartments>;
    /**
     * The Name-ID pairs of users to which the DLP policy rule must be applied.
     */
    excludedDomainProfiles?: pulumi.Input<inputs.DLPWebRulesExcludedDomainProfiles>;
    /**
     * The Name-ID pairs of users to which the DLP policy rule must be applied.
     */
    excludedGroups?: pulumi.Input<inputs.DLPWebRulesExcludedGroups>;
    /**
     * The Name-ID pairs of users to which the DLP policy rule must be applied.
     */
    excludedUsers?: pulumi.Input<inputs.DLPWebRulesExcludedUsers>;
    /**
     * The email address of an external auditor to whom DLP email notifications are sent.
     */
    externalAuditorEmail?: pulumi.Input<string>;
    /**
     * The list of file types for which the DLP policy rule must be applied.
     */
    fileTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Name-ID pairs of groups to which the DLP policy rule must be applied.
     */
    groups?: pulumi.Input<inputs.DLPWebRulesGroups>;
    /**
     * The DLP server, using ICAP, to which the transaction content is forwarded.
     */
    icapServers?: pulumi.Input<pulumi.Input<inputs.DLPWebRulesIcapServer>[]>;
    /**
     * The Name-ID pairs of users to which the DLP policy rule must be applied.
     */
    includedDomainProfiles?: pulumi.Input<inputs.DLPWebRulesIncludedDomainProfiles>;
    /**
     * list of Labels that are applicable to the rule.
     */
    labels?: pulumi.Input<inputs.DLPWebRulesLabels>;
    /**
     * The Name-ID pairs of locations groups to which the DLP policy rule must be applied.
     */
    locationGroups?: pulumi.Input<inputs.DLPWebRulesLocationGroups>;
    /**
     * The Name-ID pairs of locations to which the DLP policy rule must be applied.
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
     * The DLP policy rule name.
     */
    name?: pulumi.Input<string>;
    /**
     * The template used for DLP notification emails.
     */
    notificationTemplates?: pulumi.Input<pulumi.Input<inputs.DLPWebRulesNotificationTemplate>[]>;
    /**
     * The rule order of execution for the DLP policy rule with respect to other rules.
     */
    order?: pulumi.Input<number>;
    /**
     * The unique identifier of the parent rule under which an exception rule is added
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
     * Indicates the severity selected for the DLP rule violation
     */
    severity?: pulumi.Input<string>;
    /**
     * list of source ip groups
     */
    sourceIpGroups?: pulumi.Input<inputs.DLPWebRulesSourceIpGroups>;
    /**
     * Enables or disables the DLP policy rule.
     */
    state?: pulumi.Input<string>;
    /**
     * The list of exception rules added to a parent rule
     */
    subRules?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * list of time interval during which rule must be enforced.
     */
    timeWindows?: pulumi.Input<inputs.DLPWebRulesTimeWindows>;
    /**
     * The list of URL categories to which the DLP policy rule must be applied.
     */
    urlCategories?: pulumi.Input<inputs.DLPWebRulesUrlCategories>;
    userRiskScoreLevels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Name-ID pairs of users to which the DLP policy rule must be applied.
     */
    users?: pulumi.Input<inputs.DLPWebRulesUsers>;
    /**
     * Indicates a DLP policy rule without content inspection, when the value is set to true.
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
     * The action taken when traffic matches the DLP policy rule criteria.
     */
    action?: pulumi.Input<string>;
    /**
     * The auditor to which the DLP policy rule must be applied.
     */
    auditors?: pulumi.Input<pulumi.Input<inputs.DLPWebRulesAuditor>[]>;
    /**
     * The list of cloud applications to which the DLP policy rule must be applied.
     */
    cloudApplications?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Name-ID pairs of departments to which the DLP policy rule must be applied.
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
     * The Name-ID pairs of users to which the DLP policy rule must be applied.
     */
    excludedDepartments?: pulumi.Input<inputs.DLPWebRulesExcludedDepartments>;
    /**
     * The Name-ID pairs of users to which the DLP policy rule must be applied.
     */
    excludedDomainProfiles?: pulumi.Input<inputs.DLPWebRulesExcludedDomainProfiles>;
    /**
     * The Name-ID pairs of users to which the DLP policy rule must be applied.
     */
    excludedGroups?: pulumi.Input<inputs.DLPWebRulesExcludedGroups>;
    /**
     * The Name-ID pairs of users to which the DLP policy rule must be applied.
     */
    excludedUsers?: pulumi.Input<inputs.DLPWebRulesExcludedUsers>;
    /**
     * The email address of an external auditor to whom DLP email notifications are sent.
     */
    externalAuditorEmail?: pulumi.Input<string>;
    /**
     * The list of file types for which the DLP policy rule must be applied.
     */
    fileTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Name-ID pairs of groups to which the DLP policy rule must be applied.
     */
    groups?: pulumi.Input<inputs.DLPWebRulesGroups>;
    /**
     * The DLP server, using ICAP, to which the transaction content is forwarded.
     */
    icapServers?: pulumi.Input<pulumi.Input<inputs.DLPWebRulesIcapServer>[]>;
    /**
     * The Name-ID pairs of users to which the DLP policy rule must be applied.
     */
    includedDomainProfiles?: pulumi.Input<inputs.DLPWebRulesIncludedDomainProfiles>;
    /**
     * list of Labels that are applicable to the rule.
     */
    labels?: pulumi.Input<inputs.DLPWebRulesLabels>;
    /**
     * The Name-ID pairs of locations groups to which the DLP policy rule must be applied.
     */
    locationGroups?: pulumi.Input<inputs.DLPWebRulesLocationGroups>;
    /**
     * The Name-ID pairs of locations to which the DLP policy rule must be applied.
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
     * The DLP policy rule name.
     */
    name?: pulumi.Input<string>;
    /**
     * The template used for DLP notification emails.
     */
    notificationTemplates?: pulumi.Input<pulumi.Input<inputs.DLPWebRulesNotificationTemplate>[]>;
    /**
     * The rule order of execution for the DLP policy rule with respect to other rules.
     */
    order?: pulumi.Input<number>;
    /**
     * The unique identifier of the parent rule under which an exception rule is added
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
     * Indicates the severity selected for the DLP rule violation
     */
    severity?: pulumi.Input<string>;
    /**
     * list of source ip groups
     */
    sourceIpGroups?: pulumi.Input<inputs.DLPWebRulesSourceIpGroups>;
    /**
     * Enables or disables the DLP policy rule.
     */
    state?: pulumi.Input<string>;
    /**
     * The list of exception rules added to a parent rule
     */
    subRules?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * list of time interval during which rule must be enforced.
     */
    timeWindows?: pulumi.Input<inputs.DLPWebRulesTimeWindows>;
    /**
     * The list of URL categories to which the DLP policy rule must be applied.
     */
    urlCategories?: pulumi.Input<inputs.DLPWebRulesUrlCategories>;
    userRiskScoreLevels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Name-ID pairs of users to which the DLP policy rule must be applied.
     */
    users?: pulumi.Input<inputs.DLPWebRulesUsers>;
    /**
     * Indicates a DLP policy rule without content inspection, when the value is set to true.
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
