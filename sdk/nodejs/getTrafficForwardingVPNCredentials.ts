// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use the **zia_traffic_forwarding_vpn_credentials** data source to get information about VPN credentials that can be associated to locations. VPN is one way to route traffic from customer locations to the cloud. Site-to-Site IPSec VPN credentials can be identified by the cloud through one of the following methods:
 *
 * * Common Name (CN) of IPSec Certificate
 * * VPN User FQDN - requires VPN_SITE_TO_SITE subscription
 * * VPN IP Address - requires VPN_SITE_TO_SITE subscription
 * * Extended Authentication (XAUTH) or hosted mobile UserID - requires VPN_MOBILE subscription
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@pulumi/zia";
 *
 * const example = zia.getTrafficForwardingVPNCredentials({
 *     fqdn: "sjc-1-37@acme.com",
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@pulumi/zia";
 *
 * const example = zia.getTrafficForwardingVPNCredentials({
 *     ipAddress: "1.1.1.1",
 * });
 * ```
 */
export function getTrafficForwardingVPNCredentials(args?: GetTrafficForwardingVPNCredentialsArgs, opts?: pulumi.InvokeOptions): Promise<GetTrafficForwardingVPNCredentialsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zia:index/getTrafficForwardingVPNCredentials:getTrafficForwardingVPNCredentials", {
        "fqdn": args.fqdn,
        "id": args.id,
        "ipAddress": args.ipAddress,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getTrafficForwardingVPNCredentials.
 */
export interface GetTrafficForwardingVPNCredentialsArgs {
    /**
     * (String) Fully Qualified Domain Name. Applicable only to `UFQDN` or `XAUTH` (or `HOSTED_MOBILE_USERS`) auth type.
     */
    fqdn?: string;
    /**
     * Unique identifer of the GRE virtual IP address (VIP)
     */
    id?: number;
    /**
     * Filter based on an IP address range.
     */
    ipAddress?: string;
    /**
     * (String) VPN authentication type (i.e., how the VPN credential is sent to the server). It is not modifiable after VpnCredential is created.
     */
    type?: string;
}

/**
 * A collection of values returned by getTrafficForwardingVPNCredentials.
 */
export interface GetTrafficForwardingVPNCredentialsResult {
    /**
     * (String) Additional information about this VPN credential.
     */
    readonly comments: string;
    /**
     * (String) Fully Qualified Domain Name. Applicable only to `UFQDN` or `XAUTH` (or `HOSTED_MOBILE_USERS`) auth type.
     */
    readonly fqdn?: string;
    /**
     * (Number) Identifier that uniquely identifies an entity
     */
    readonly id: number;
    readonly ipAddress?: string;
    /**
     * (Set of Object) Location that is associated to this VPN credential. Non-existence means not associated to any location.
     */
    readonly locations: outputs.GetTrafficForwardingVPNCredentialsLocation[];
    /**
     * (Set of Object) SD-WAN Partner that manages the location. If a partner does not manage the locaton, this is set to Self.
     */
    readonly managedBies: outputs.GetTrafficForwardingVPNCredentialsManagedBy[];
    /**
     * (String) Pre-shared key. This is a required field for `UFQDN` and `IP` auth type.
     */
    readonly preSharedKey: string;
    /**
     * (String) VPN authentication type (i.e., how the VPN credential is sent to the server). It is not modifiable after VpnCredential is created.
     */
    readonly type: string;
}
/**
 * Use the **zia_traffic_forwarding_vpn_credentials** data source to get information about VPN credentials that can be associated to locations. VPN is one way to route traffic from customer locations to the cloud. Site-to-Site IPSec VPN credentials can be identified by the cloud through one of the following methods:
 *
 * * Common Name (CN) of IPSec Certificate
 * * VPN User FQDN - requires VPN_SITE_TO_SITE subscription
 * * VPN IP Address - requires VPN_SITE_TO_SITE subscription
 * * Extended Authentication (XAUTH) or hosted mobile UserID - requires VPN_MOBILE subscription
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@pulumi/zia";
 *
 * const example = zia.getTrafficForwardingVPNCredentials({
 *     fqdn: "sjc-1-37@acme.com",
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@pulumi/zia";
 *
 * const example = zia.getTrafficForwardingVPNCredentials({
 *     ipAddress: "1.1.1.1",
 * });
 * ```
 */
export function getTrafficForwardingVPNCredentialsOutput(args?: GetTrafficForwardingVPNCredentialsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTrafficForwardingVPNCredentialsResult> {
    return pulumi.output(args).apply((a: any) => getTrafficForwardingVPNCredentials(a, opts))
}

/**
 * A collection of arguments for invoking getTrafficForwardingVPNCredentials.
 */
export interface GetTrafficForwardingVPNCredentialsOutputArgs {
    /**
     * (String) Fully Qualified Domain Name. Applicable only to `UFQDN` or `XAUTH` (or `HOSTED_MOBILE_USERS`) auth type.
     */
    fqdn?: pulumi.Input<string>;
    /**
     * Unique identifer of the GRE virtual IP address (VIP)
     */
    id?: pulumi.Input<number>;
    /**
     * Filter based on an IP address range.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * (String) VPN authentication type (i.e., how the VPN credential is sent to the server). It is not modifiable after VpnCredential is created.
     */
    type?: pulumi.Input<string>;
}
