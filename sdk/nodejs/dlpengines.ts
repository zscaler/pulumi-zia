// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zia/about-dlp-engines)
 * * [API documentation](https://help.zscaler.com/zia/data-loss-prevention#/dlpEngines-get)
 *
 * Use the **zia_dlp_engines** resource allows the creation and management of ZIA DLP Engines in the Zscaler Internet Access cloud or via the API.
 *
 * ⚠️ **WARNING:** "Before using the new ``zia.DLPEngines`` resource contact [Zscaler Support](https://help.zscaler.com/login-tickets)." and request the following API methods ``POST``, ``PUT``, and ``DELETE`` to be enabled for your organization.
 *
 * ## Example Usage
 *
 * ## Import
 *
 * Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.
 *
 * Visit
 *
 * **zia_dlp_engines** can be imported by using `<ENGINE_ID>` or `<ENGINE_NAME>` as the import ID.
 *
 * For example:
 *
 * ```sh
 * $ pulumi import zia:index/dLPEngines:DLPEngines example <engine_id>
 * ```
 *
 * or
 *
 * ```sh
 * $ pulumi import zia:index/dLPEngines:DLPEngines example <engine_name>
 * ```
 */
export class DLPEngines extends pulumi.CustomResource {
    /**
     * Get an existing DLPEngines resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DLPEnginesState, opts?: pulumi.CustomResourceOptions): DLPEngines {
        return new DLPEngines(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zia:index/dLPEngines:DLPEngines';

    /**
     * Returns true if the given object is an instance of DLPEngines.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DLPEngines {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DLPEngines.__pulumiType;
    }

    /**
     * Indicates whether this is a custom DLP engine. If this value is set to true, the engine is custom.
     */
    public readonly customDlpEngine!: pulumi.Output<boolean | undefined>;
    /**
     * The DLP engine's description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The boolean logical operator in which various DLP dictionaries are combined within a DLP engine's expression.
     */
    public readonly engineExpression!: pulumi.Output<string | undefined>;
    public /*out*/ readonly engineId!: pulumi.Output<number>;
    /**
     * The DLP engine name as configured by the admin.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a DLPEngines resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DLPEnginesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DLPEnginesArgs | DLPEnginesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DLPEnginesState | undefined;
            resourceInputs["customDlpEngine"] = state ? state.customDlpEngine : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["engineExpression"] = state ? state.engineExpression : undefined;
            resourceInputs["engineId"] = state ? state.engineId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as DLPEnginesArgs | undefined;
            resourceInputs["customDlpEngine"] = args ? args.customDlpEngine : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["engineExpression"] = args ? args.engineExpression : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["engineId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DLPEngines.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DLPEngines resources.
 */
export interface DLPEnginesState {
    /**
     * Indicates whether this is a custom DLP engine. If this value is set to true, the engine is custom.
     */
    customDlpEngine?: pulumi.Input<boolean>;
    /**
     * The DLP engine's description.
     */
    description?: pulumi.Input<string>;
    /**
     * The boolean logical operator in which various DLP dictionaries are combined within a DLP engine's expression.
     */
    engineExpression?: pulumi.Input<string>;
    engineId?: pulumi.Input<number>;
    /**
     * The DLP engine name as configured by the admin.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DLPEngines resource.
 */
export interface DLPEnginesArgs {
    /**
     * Indicates whether this is a custom DLP engine. If this value is set to true, the engine is custom.
     */
    customDlpEngine?: pulumi.Input<boolean>;
    /**
     * The DLP engine's description.
     */
    description?: pulumi.Input<string>;
    /**
     * The boolean logical operator in which various DLP dictionaries are combined within a DLP engine's expression.
     */
    engineExpression?: pulumi.Input<string>;
    /**
     * The DLP engine name as configured by the admin.
     */
    name?: pulumi.Input<string>;
}
