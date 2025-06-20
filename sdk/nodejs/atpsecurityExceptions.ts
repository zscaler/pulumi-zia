// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zia/configuring-advanced-threat-protection-policy)
 * * [API documentation](https://help.zscaler.com/zia/advanced-threat-protection-policy#/)
 *
 * The **zia_atp_security_exceptions** resource alows you to updates security exceptions for the ATP policy. To learn more see [Advanced Threat Protection](https://help.zscaler.com/unified/configuring-security-exceptions-advanced-threat-protection-policy)
 *
 * ## Example Usage
 *
 * ## Import
 *
 * Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.
 *
 * Visit
 *
 * **zia_atp_security_exceptions** can be imported by using `all_urls` as the import ID.
 *
 * For example:
 *
 * ```sh
 * $ pulumi import zia:index/aTPSecurityExceptions:ATPSecurityExceptions this all_urls
 * ```
 */
export class ATPSecurityExceptions extends pulumi.CustomResource {
    /**
     * Get an existing ATPSecurityExceptions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ATPSecurityExceptionsState, opts?: pulumi.CustomResourceOptions): ATPSecurityExceptions {
        return new ATPSecurityExceptions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zia:index/aTPSecurityExceptions:ATPSecurityExceptions';

    /**
     * Returns true if the given object is an instance of ATPSecurityExceptions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ATPSecurityExceptions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ATPSecurityExceptions.__pulumiType;
    }

    public readonly bypassUrls!: pulumi.Output<string[]>;

    /**
     * Create a ATPSecurityExceptions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ATPSecurityExceptionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ATPSecurityExceptionsArgs | ATPSecurityExceptionsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ATPSecurityExceptionsState | undefined;
            resourceInputs["bypassUrls"] = state ? state.bypassUrls : undefined;
        } else {
            const args = argsOrState as ATPSecurityExceptionsArgs | undefined;
            resourceInputs["bypassUrls"] = args ? args.bypassUrls : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ATPSecurityExceptions.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ATPSecurityExceptions resources.
 */
export interface ATPSecurityExceptionsState {
    bypassUrls?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a ATPSecurityExceptions resource.
 */
export interface ATPSecurityExceptionsArgs {
    bypassUrls?: pulumi.Input<pulumi.Input<string>[]>;
}
