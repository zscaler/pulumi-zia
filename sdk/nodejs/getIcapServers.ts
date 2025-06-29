// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zia/about-icap-communication-between-zscaler-and-dlp-servers)
 * * [API documentation](https://help.zscaler.com/zia/data-loss-prevention#/icapServers/lite-get)
 *
 * Use the **zia_dlp_engines** data source to get information about a the list of DLP servers using ICAP in the Zscaler Internet Access cloud or via the API.
 *
 * ## Example Usage
 */
export function getIcapServers(args?: GetIcapServersArgs, opts?: pulumi.InvokeOptions): Promise<GetIcapServersResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zia:index/getIcapServers:getIcapServers", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getIcapServers.
 */
export interface GetIcapServersArgs {
    name?: string;
}

/**
 * A collection of values returned by getIcapServers.
 */
export interface GetIcapServersResult {
    readonly id: number;
    readonly name?: string;
    readonly status: string;
    readonly url: string;
}
/**
 * * [Official documentation](https://help.zscaler.com/zia/about-icap-communication-between-zscaler-and-dlp-servers)
 * * [API documentation](https://help.zscaler.com/zia/data-loss-prevention#/icapServers/lite-get)
 *
 * Use the **zia_dlp_engines** data source to get information about a the list of DLP servers using ICAP in the Zscaler Internet Access cloud or via the API.
 *
 * ## Example Usage
 */
export function getIcapServersOutput(args?: GetIcapServersOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetIcapServersResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("zia:index/getIcapServers:getIcapServers", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getIcapServers.
 */
export interface GetIcapServersOutputArgs {
    name?: pulumi.Input<string>;
}
