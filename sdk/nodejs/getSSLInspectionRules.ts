// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zia/about-ssl-inspection-policy)
 * * [API documentation](https://help.zscaler.com/zia/ssl-inspection-policy#/sslInspectionRules-get)
 *
 * Use the **zia_ssl_inspection_rules** data source to get information about a ssl inspection rule in the Zscaler Internet Access.
 *
 * ## Example Usage
 */
export function getSSLInspectionRules(args?: GetSSLInspectionRulesArgs, opts?: pulumi.InvokeOptions): Promise<GetSSLInspectionRulesResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zia:index/getSSLInspectionRules:getSSLInspectionRules", {
        "id": args.id,
        "name": args.name,
        "urlCategories": args.urlCategories,
    }, opts);
}

/**
 * A collection of arguments for invoking getSSLInspectionRules.
 */
export interface GetSSLInspectionRulesArgs {
    /**
     * Unique identifier for the SSL Inspection
     */
    id?: number;
    /**
     * Name of the SSL Inspection
     */
    name?: string;
    /**
     * The list of URL categories to which the DLP policy rule must be applied.
     */
    urlCategories?: string[];
}

/**
 * A collection of values returned by getSSLInspectionRules.
 */
export interface GetSSLInspectionRulesResult {
    /**
     * Action taken when the traffic matches policy
     */
    readonly actions: outputs.GetSSLInspectionRulesAction[];
    /**
     * The list of URL categories to which the DLP policy rule must be applied.
     */
    readonly cloudApplications: string[];
    /**
     * ID pairs of departments for which the rule is applied.
     */
    readonly departments: outputs.GetSSLInspectionRulesDepartment[];
    /**
     * The description of the workload group
     */
    readonly description: string;
    /**
     * ID pairs of destination IP address groups for which the rule is applied.
     */
    readonly destIpGroups: outputs.GetSSLInspectionRulesDestIpGroup[];
    /**
     * ID pairs of device groups for which the rule is applied.
     */
    readonly deviceGroups: outputs.GetSSLInspectionRulesDeviceGroup[];
    /**
     * Lists device trust levels for which the rule must be applied (for devices managed using Zscaler Client Connector).
     */
    readonly deviceTrustLevels: string[];
    /**
     * ID pairs of devices for which the rule is applied
     */
    readonly devices: outputs.GetSSLInspectionRulesDevice[];
    /**
     * ID pairs of groups for which the rule is applied. If not set, rule is applied for all groups.
     */
    readonly groups: outputs.GetSSLInspectionRulesGroup[];
    /**
     * A unique identifier assigned to the workload group
     */
    readonly id: number;
    /**
     * ID pairs of labels associated with the rule.
     */
    readonly labels: outputs.GetSSLInspectionRulesLabel[];
    /**
     * A nested block with details about who last modified the workload group.
     */
    readonly lastModifiedBies: outputs.GetSSLInspectionRulesLastModifiedBy[];
    /**
     * Timestamp when the workload group was last modified.
     */
    readonly lastModifiedTime: number;
    /**
     * ID pairs of location groups to which the rule is applied. When empty, it implies applying to all location groups.
     */
    readonly locationGroups: outputs.GetSSLInspectionRulesLocationGroup[];
    /**
     * ID pairs of locations to which the rule is applied. When empty, it implies applying to all locations.
     */
    readonly locations: outputs.GetSSLInspectionRulesLocation[];
    /**
     * The name of the workload group
     */
    readonly name: string;
    /**
     * Policy rules are evaluated in ascending numerical order (Rule 1 before Rule 2, and so on), and the Rule Order reflects this rule's place in the order.
     */
    readonly order: number;
    /**
     * Zscaler Client Connector device platforms for which this rule is applied. Supported Values: `SCAN_IOS`, `SCAN_ANDROID`, `SCAN_MACOS`, `SCAN_WINDOWS`, `NO_CLIENT_CONNECTOR`, `SCAN_LINUX`
     */
    readonly platforms: string[];
    /**
     * When using ZPA Gateway forwarding, name-ID pairs of ZPA Application Segments for which the rule is applicable.
     */
    readonly proxyGateways: outputs.GetSSLInspectionRulesProxyGateway[];
    /**
     * The admin rank specified for the rule based on your assigned admin rank. Admin rank determines the rule order that can be specified for the rule. Admin rank can be configured if it is enabled in the Advanced Settings.
     */
    readonly rank: number;
    /**
     * Indicates whether this rule is applied to remote users that use PAC with Kerberos authentication.
     */
    readonly roadWarriorForKerberos: boolean;
    /**
     * ID pairs of source IP address groups for which the rule is applied.
     */
    readonly sourceIpGroups: outputs.GetSSLInspectionRulesSourceIpGroup[];
    /**
     * The state of the rule indicating whether it is enabled or disabled. Supported values: `ENABLED` or `DISABLED`
     */
    readonly state: string;
    /**
     * The time intervals during which the rule applies
     */
    readonly timeWindows: outputs.GetSSLInspectionRulesTimeWindow[];
    /**
     * The list of URL categories to which the DLP policy rule must be applied.
     */
    readonly urlCategories?: string[];
    /**
     * A list of user agent types the rule applies to.
     */
    readonly userAgentTypes: string[];
    /**
     * The list of preconfigured workload groups to which the policy must be applied.
     */
    readonly users: outputs.GetSSLInspectionRulesUser[];
    /**
     * The list of preconfigured workload groups to which the policy must be applied.
     */
    readonly workloadGroups: outputs.GetSSLInspectionRulesWorkloadGroup[];
    /**
     * The list of ZPA Application Segments for which this rule is applicable (applicable only for ZPA Gateway forwarding).
     */
    readonly zpaAppSegments: outputs.GetSSLInspectionRulesZpaAppSegment[];
}
/**
 * * [Official documentation](https://help.zscaler.com/zia/about-ssl-inspection-policy)
 * * [API documentation](https://help.zscaler.com/zia/ssl-inspection-policy#/sslInspectionRules-get)
 *
 * Use the **zia_ssl_inspection_rules** data source to get information about a ssl inspection rule in the Zscaler Internet Access.
 *
 * ## Example Usage
 */
export function getSSLInspectionRulesOutput(args?: GetSSLInspectionRulesOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetSSLInspectionRulesResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("zia:index/getSSLInspectionRules:getSSLInspectionRules", {
        "id": args.id,
        "name": args.name,
        "urlCategories": args.urlCategories,
    }, opts);
}

/**
 * A collection of arguments for invoking getSSLInspectionRules.
 */
export interface GetSSLInspectionRulesOutputArgs {
    /**
     * Unique identifier for the SSL Inspection
     */
    id?: pulumi.Input<number>;
    /**
     * Name of the SSL Inspection
     */
    name?: pulumi.Input<string>;
    /**
     * The list of URL categories to which the DLP policy rule must be applied.
     */
    urlCategories?: pulumi.Input<pulumi.Input<string>[]>;
}
