// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use the **zia_dlp_notification_templates** data source to get information about a ZIA DLP Notification Templates in the Zscaler Internet Access cloud or via the API.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@pulumi/zia";
 *
 * const example = zia.getDLPNotificationTemplates({
 *     name: "DLP Auditor Template Test",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDLPNotificationTemplates(args?: GetDLPNotificationTemplatesArgs, opts?: pulumi.InvokeOptions): Promise<GetDLPNotificationTemplatesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zia:index/getDLPNotificationTemplates:getDLPNotificationTemplates", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getDLPNotificationTemplates.
 */
export interface GetDLPNotificationTemplatesArgs {
    /**
     * The unique identifier for a DLP notification template.
     */
    id?: number;
    /**
     * The DLP policy rule name.
     */
    name?: string;
}

/**
 * A collection of values returned by getDLPNotificationTemplates.
 */
export interface GetDLPNotificationTemplatesResult {
    readonly attachContent: boolean;
    readonly htmlMessage: string;
    readonly id: number;
    readonly name: string;
    readonly plainTextMessage: string;
    readonly subject: string;
    readonly tlsEnabled: boolean;
}
/**
 * Use the **zia_dlp_notification_templates** data source to get information about a ZIA DLP Notification Templates in the Zscaler Internet Access cloud or via the API.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@pulumi/zia";
 *
 * const example = zia.getDLPNotificationTemplates({
 *     name: "DLP Auditor Template Test",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDLPNotificationTemplatesOutput(args?: GetDLPNotificationTemplatesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDLPNotificationTemplatesResult> {
    return pulumi.output(args).apply((a: any) => getDLPNotificationTemplates(a, opts))
}

/**
 * A collection of arguments for invoking getDLPNotificationTemplates.
 */
export interface GetDLPNotificationTemplatesOutputArgs {
    /**
     * The unique identifier for a DLP notification template.
     */
    id?: pulumi.Input<number>;
    /**
     * The DLP policy rule name.
     */
    name?: pulumi.Input<string>;
}
