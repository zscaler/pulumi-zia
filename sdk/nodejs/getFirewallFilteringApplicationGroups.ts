// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use the **zia_firewall_filtering_network_application_groups** data source to get information about a network application group available in the Zscaler Internet Access cloud firewall. This data source can then be associated with a ZIA firewall filtering network application rule.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@pulumi/zia";
 *
 * const example = zia.getFirewallFilteringApplicationGroups({
 *     name: "example",
 * });
 * ```
 */
export function getFirewallFilteringApplicationGroups(args?: GetFirewallFilteringApplicationGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallFilteringApplicationGroupsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zia:index/getFirewallFilteringApplicationGroups:getFirewallFilteringApplicationGroups", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getFirewallFilteringApplicationGroups.
 */
export interface GetFirewallFilteringApplicationGroupsArgs {
    /**
     * The ID of the ip source group resource.
     */
    id?: number;
    /**
     * The name of the ip source group to be exported.
     */
    name?: string;
}

/**
 * A collection of values returned by getFirewallFilteringApplicationGroups.
 */
export interface GetFirewallFilteringApplicationGroupsResult {
    /**
     * (String)
     */
    readonly description: string;
    /**
     * The ID of this resource.
     */
    readonly id: number;
    readonly name: string;
    /**
     * (List of String)
     */
    readonly networkApplications: string[];
}
/**
 * Use the **zia_firewall_filtering_network_application_groups** data source to get information about a network application group available in the Zscaler Internet Access cloud firewall. This data source can then be associated with a ZIA firewall filtering network application rule.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@pulumi/zia";
 *
 * const example = zia.getFirewallFilteringApplicationGroups({
 *     name: "example",
 * });
 * ```
 */
export function getFirewallFilteringApplicationGroupsOutput(args?: GetFirewallFilteringApplicationGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallFilteringApplicationGroupsResult> {
    return pulumi.output(args).apply((a: any) => getFirewallFilteringApplicationGroups(a, opts))
}

/**
 * A collection of arguments for invoking getFirewallFilteringApplicationGroups.
 */
export interface GetFirewallFilteringApplicationGroupsOutputArgs {
    /**
     * The ID of the ip source group resource.
     */
    id?: pulumi.Input<number>;
    /**
     * The name of the ip source group to be exported.
     */
    name?: pulumi.Input<string>;
}
