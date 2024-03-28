// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use the **zia_rule_labels** data source to get information about a rule label resource in the Zscaler Internet Access cloud or via the API. This data source can then be associated with resources such as: Firewall Rules and URL filtering rules
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@pulumi/zia";
 *
 * const example = zia.getRuleLabels({
 *     name: "Example",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getRuleLabels(args?: GetRuleLabelsArgs, opts?: pulumi.InvokeOptions): Promise<GetRuleLabelsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zia:index/getRuleLabels:getRuleLabels", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getRuleLabels.
 */
export interface GetRuleLabelsArgs {
    /**
     * The unique identifer for the device group.
     */
    id?: number;
    /**
     * The name of the rule label to be exported.
     */
    name?: string;
}

/**
 * A collection of values returned by getRuleLabels.
 */
export interface GetRuleLabelsResult {
    /**
     * (String) The admin that created the rule label. This is a read-only field. Ignored by PUT requests.
     */
    readonly createdBies: outputs.GetRuleLabelsCreatedBy[];
    /**
     * (String) The rule label description.
     */
    readonly description: string;
    readonly id: number;
    /**
     * (String) The admin that modified the rule label last. This is a read-only field. Ignored by PUT requests.
     */
    readonly lastModifiedBies: outputs.GetRuleLabelsLastModifiedBy[];
    /**
     * (String) Timestamp when the rule lable was last modified. This is a read-only field. Ignored by PUT and DELETE requests.
     */
    readonly lastModifiedTime: number;
    readonly name: string;
    /**
     * (int) The number of rules that reference the label.
     */
    readonly referencedRuleCount: number;
}
/**
 * Use the **zia_rule_labels** data source to get information about a rule label resource in the Zscaler Internet Access cloud or via the API. This data source can then be associated with resources such as: Firewall Rules and URL filtering rules
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@pulumi/zia";
 *
 * const example = zia.getRuleLabels({
 *     name: "Example",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getRuleLabelsOutput(args?: GetRuleLabelsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRuleLabelsResult> {
    return pulumi.output(args).apply((a: any) => getRuleLabels(a, opts))
}

/**
 * A collection of arguments for invoking getRuleLabels.
 */
export interface GetRuleLabelsOutputArgs {
    /**
     * The unique identifer for the device group.
     */
    id?: pulumi.Input<number>;
    /**
     * The name of the rule label to be exported.
     */
    name?: pulumi.Input<string>;
}
