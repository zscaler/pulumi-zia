// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use the **zia_cloud_browser_isolation_profile** data source to get information about an isolation profile in the Zscaler Internet Access cloud. This data source is required when configuring URL filtering rule where the action is set to `ISOLATE`
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@pulumi/zia";
 *
 * const this = zia.getCbiProfile({
 *     name: "ZS_CBI_Profile1",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCbiProfile(args?: GetCbiProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetCbiProfileResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zia:index/getCbiProfile:getCbiProfile", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getCbiProfile.
 */
export interface GetCbiProfileArgs {
    /**
     * (string) The universally unique identifier (UUID) for the browser isolation profile.
     */
    id?: string;
    /**
     * This field defines the name of the isolation profile.
     */
    name?: string;
}

/**
 * A collection of values returned by getCbiProfile.
 */
export interface GetCbiProfileResult {
    /**
     * (Optional) Indicates whether this is a default browser isolation profile. Zscaler sets this field
     */
    readonly defaultProfile: boolean;
    /**
     * (string) The universally unique identifier (UUID) for the browser isolation profile.
     */
    readonly id?: string;
    readonly name?: string;
    /**
     * (string) The browser isolation profile URL
     */
    readonly url: string;
}
/**
 * Use the **zia_cloud_browser_isolation_profile** data source to get information about an isolation profile in the Zscaler Internet Access cloud. This data source is required when configuring URL filtering rule where the action is set to `ISOLATE`
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@pulumi/zia";
 *
 * const this = zia.getCbiProfile({
 *     name: "ZS_CBI_Profile1",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCbiProfileOutput(args?: GetCbiProfileOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCbiProfileResult> {
    return pulumi.output(args).apply((a: any) => getCbiProfile(a, opts))
}

/**
 * A collection of arguments for invoking getCbiProfile.
 */
export interface GetCbiProfileOutputArgs {
    /**
     * (string) The universally unique identifier (UUID) for the browser isolation profile.
     */
    id?: pulumi.Input<string>;
    /**
     * This field defines the name of the isolation profile.
     */
    name?: pulumi.Input<string>;
}
