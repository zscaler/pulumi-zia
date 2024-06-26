// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use the **zia_dlp_engines** data source to get information about a ZIA DLP Engines in the Zscaler Internet Access cloud or via the API.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@pulumi/zia";
 *
 * const this = zia.getDLPEngines({
 *     name: "Example",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@pulumi/zia";
 *
 * const this = zia.getDLPEngines({
 *     id: 1234567890,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDLPEngines(args?: GetDLPEnginesArgs, opts?: pulumi.InvokeOptions): Promise<GetDLPEnginesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zia:index/getDLPEngines:getDLPEngines", {
        "id": args.id,
        "name": args.name,
        "predefinedEngineName": args.predefinedEngineName,
    }, opts);
}

/**
 * A collection of arguments for invoking getDLPEngines.
 */
export interface GetDLPEnginesArgs {
    /**
     * The unique identifier for the DLP engine.
     */
    id?: number;
    /**
     * The DLP engine name as configured by the admin. This attribute is required in POST and PUT requests for custom DLP engines.
     */
    name?: string;
    /**
     * The name of the predefined DLP engine.
     */
    predefinedEngineName?: string;
}

/**
 * A collection of values returned by getDLPEngines.
 */
export interface GetDLPEnginesResult {
    readonly customDlpEngine: boolean;
    readonly description: string;
    readonly engineExpression: string;
    readonly id?: number;
    readonly name?: string;
    readonly predefinedEngineName?: string;
}
/**
 * Use the **zia_dlp_engines** data source to get information about a ZIA DLP Engines in the Zscaler Internet Access cloud or via the API.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@pulumi/zia";
 *
 * const this = zia.getDLPEngines({
 *     name: "Example",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@pulumi/zia";
 *
 * const this = zia.getDLPEngines({
 *     id: 1234567890,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDLPEnginesOutput(args?: GetDLPEnginesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDLPEnginesResult> {
    return pulumi.output(args).apply((a: any) => getDLPEngines(a, opts))
}

/**
 * A collection of arguments for invoking getDLPEngines.
 */
export interface GetDLPEnginesOutputArgs {
    /**
     * The unique identifier for the DLP engine.
     */
    id?: pulumi.Input<number>;
    /**
     * The DLP engine name as configured by the admin. This attribute is required in POST and PUT requests for custom DLP engines.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the predefined DLP engine.
     */
    predefinedEngineName?: pulumi.Input<string>;
}
