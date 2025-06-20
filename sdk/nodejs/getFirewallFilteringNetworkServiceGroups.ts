// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zia/firewall-policies#/networkServiceGroups-get)
 * * [API documentation](https://help.zscaler.com/zia/firewall-policies#/networkServiceGroups-get)
 *
 * Use the **zia_firewall_filtering_network_service_groups** data source to get information about a network service groups available in the Zscaler Internet Access cloud firewall. This data source can then be associated with a ZIA firewall filtering network service rule.
 *
 * ## Example Usage
 */
export function getFirewallFilteringNetworkServiceGroups(args?: GetFirewallFilteringNetworkServiceGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallFilteringNetworkServiceGroupsResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zia:index/getFirewallFilteringNetworkServiceGroups:getFirewallFilteringNetworkServiceGroups", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getFirewallFilteringNetworkServiceGroups.
 */
export interface GetFirewallFilteringNetworkServiceGroupsArgs {
    /**
     * The ID of the ip source group to be exported.
     */
    id?: number;
    /**
     * The name of the ip source group to be exported.
     */
    name?: string;
}

/**
 * A collection of values returned by getFirewallFilteringNetworkServiceGroups.
 */
export interface GetFirewallFilteringNetworkServiceGroupsResult {
    /**
     * (String)
     */
    readonly description: string;
    /**
     * (Number)
     */
    readonly id: number;
    /**
     * (String)
     */
    readonly name: string;
    /**
     * (Number) The ID of this resource.
     */
    readonly services: outputs.GetFirewallFilteringNetworkServiceGroupsService[];
}
/**
 * * [Official documentation](https://help.zscaler.com/zia/firewall-policies#/networkServiceGroups-get)
 * * [API documentation](https://help.zscaler.com/zia/firewall-policies#/networkServiceGroups-get)
 *
 * Use the **zia_firewall_filtering_network_service_groups** data source to get information about a network service groups available in the Zscaler Internet Access cloud firewall. This data source can then be associated with a ZIA firewall filtering network service rule.
 *
 * ## Example Usage
 */
export function getFirewallFilteringNetworkServiceGroupsOutput(args?: GetFirewallFilteringNetworkServiceGroupsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetFirewallFilteringNetworkServiceGroupsResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("zia:index/getFirewallFilteringNetworkServiceGroups:getFirewallFilteringNetworkServiceGroups", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getFirewallFilteringNetworkServiceGroups.
 */
export interface GetFirewallFilteringNetworkServiceGroupsOutputArgs {
    /**
     * The ID of the ip source group to be exported.
     */
    id?: pulumi.Input<number>;
    /**
     * The name of the ip source group to be exported.
     */
    name?: pulumi.Input<string>;
}
