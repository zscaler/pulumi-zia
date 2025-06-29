// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zia/about-rule-labels)
 * * [API documentation](https://help.zscaler.com/zia/rule-labels#/ruleLabels-get)
 *
 * The **zia_rule_labels** resource allows the creation and management of rule labels in the Zscaler Internet Access cloud or via the API. This resource can then be associated with resources such as: Firewall Rules and URL filtering rules
 *
 * ## Example Usage
 *
 * ## Import
 *
 * Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.
 *
 * Visit
 *
 * **zia_rule_labels** can be imported by using `<LABEL_ID>` or `<LABEL_NAME>` as the import ID.
 *
 * For example:
 *
 * ```sh
 * $ pulumi import zia:index/ruleLabels:RuleLabels example <label_id>
 * ```
 *
 * or
 *
 * ```sh
 * $ pulumi import zia:index/ruleLabels:RuleLabels example <label_name>
 * ```
 */
export class RuleLabels extends pulumi.CustomResource {
    /**
     * Get an existing RuleLabels resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleLabelsState, opts?: pulumi.CustomResourceOptions): RuleLabels {
        return new RuleLabels(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zia:index/ruleLabels:RuleLabels';

    /**
     * Returns true if the given object is an instance of RuleLabels.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RuleLabels {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RuleLabels.__pulumiType;
    }

    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the devices to be created.
     */
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly ruleLabelId!: pulumi.Output<number>;

    /**
     * Create a RuleLabels resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RuleLabelsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleLabelsArgs | RuleLabelsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RuleLabelsState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["ruleLabelId"] = state ? state.ruleLabelId : undefined;
        } else {
            const args = argsOrState as RuleLabelsArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ruleLabelId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RuleLabels.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RuleLabels resources.
 */
export interface RuleLabelsState {
    description?: pulumi.Input<string>;
    /**
     * The name of the devices to be created.
     */
    name?: pulumi.Input<string>;
    ruleLabelId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RuleLabels resource.
 */
export interface RuleLabelsArgs {
    description?: pulumi.Input<string>;
    /**
     * The name of the devices to be created.
     */
    name?: pulumi.Input<string>;
}
