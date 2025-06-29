// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zia/about-email-labels)
 * * [API documentation](https://help.zscaler.com/zia/saas-security-api#/casbEmailLabel/lite-get)
 *
 * Use the **zia_casb_email_label** data source to get information about email labels generated for the SaaS Security API policies in a user's email account
 *
 * ## Example Usage
 *
 * ### By Name
 *
 * ### By ID
 */
export function getCasbEmailLabel(args?: GetCasbEmailLabelArgs, opts?: pulumi.InvokeOptions): Promise<GetCasbEmailLabelResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zia:index/getCasbEmailLabel:getCasbEmailLabel", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getCasbEmailLabel.
 */
export interface GetCasbEmailLabelArgs {
    /**
     * SaaS Security API email label ID
     */
    id?: number;
    /**
     * SaaS Security API email label name
     */
    name?: string;
}

/**
 * A collection of values returned by getCasbEmailLabel.
 */
export interface GetCasbEmailLabelResult {
    readonly id: number;
    /**
     * (Boolean) A Boolean value that indicates whether or not the email label is deleted
     */
    readonly labelDeleted: boolean;
    readonly name: string;
}
/**
 * * [Official documentation](https://help.zscaler.com/zia/about-email-labels)
 * * [API documentation](https://help.zscaler.com/zia/saas-security-api#/casbEmailLabel/lite-get)
 *
 * Use the **zia_casb_email_label** data source to get information about email labels generated for the SaaS Security API policies in a user's email account
 *
 * ## Example Usage
 *
 * ### By Name
 *
 * ### By ID
 */
export function getCasbEmailLabelOutput(args?: GetCasbEmailLabelOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetCasbEmailLabelResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("zia:index/getCasbEmailLabel:getCasbEmailLabel", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getCasbEmailLabel.
 */
export interface GetCasbEmailLabelOutputArgs {
    /**
     * SaaS Security API email label ID
     */
    id?: pulumi.Input<number>;
    /**
     * SaaS Security API email label name
     */
    name?: pulumi.Input<string>;
}
