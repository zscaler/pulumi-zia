// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zia/sandbox-policy-settings#/behavioralAnalysisAdvancedSettings-get)
 * * [API documentation](https://help.zscaler.com/zia/sandbox-policy-settings#/behavioralAnalysisAdvancedSettings-get)
 *
 * The **zia_sandbox_behavioral_analysis** resource updates the custom list of MD5 file hashes that are blocked by Sandbox. This overwrites a previously generated blocklist. If you need to completely erase the blocklist, submit an empty list.
 *
 * **Note**: Only the file types that are supported by Sandbox analysis can be blocked using MD5 hashes.
 *
 * ## Example Usage
 *
 * ### Add MD5 Hashes To Sandbox
 *
 * ### Remove All MD5 Hashes To Sandbox
 *
 * ## Import
 *
 * Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.
 *
 * Visit
 *
 * **zia_sandbox_behavioral_analysis** can be imported by using `sandbox_settings` as the import ID.
 *
 * For example:
 *
 * ```sh
 * $ pulumi import zia:index/sandboxBehavioralAnalysis:SandboxBehavioralAnalysis example sandbox_settings
 * ```
 */
export class SandboxBehavioralAnalysis extends pulumi.CustomResource {
    /**
     * Get an existing SandboxBehavioralAnalysis resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SandboxBehavioralAnalysisState, opts?: pulumi.CustomResourceOptions): SandboxBehavioralAnalysis {
        return new SandboxBehavioralAnalysis(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zia:index/sandboxBehavioralAnalysis:SandboxBehavioralAnalysis';

    /**
     * Returns true if the given object is an instance of SandboxBehavioralAnalysis.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SandboxBehavioralAnalysis {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SandboxBehavioralAnalysis.__pulumiType;
    }

    /**
     * A custom list of unique MD5 file hashes that must be blocked by Sandbox. A maximum of 10000 MD5 file hashes can be
     * blocked
     */
    public readonly fileHashesToBeBlockeds!: pulumi.Output<string[] | undefined>;

    /**
     * Create a SandboxBehavioralAnalysis resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SandboxBehavioralAnalysisArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SandboxBehavioralAnalysisArgs | SandboxBehavioralAnalysisState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SandboxBehavioralAnalysisState | undefined;
            resourceInputs["fileHashesToBeBlockeds"] = state ? state.fileHashesToBeBlockeds : undefined;
        } else {
            const args = argsOrState as SandboxBehavioralAnalysisArgs | undefined;
            resourceInputs["fileHashesToBeBlockeds"] = args ? args.fileHashesToBeBlockeds : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SandboxBehavioralAnalysis.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SandboxBehavioralAnalysis resources.
 */
export interface SandboxBehavioralAnalysisState {
    /**
     * A custom list of unique MD5 file hashes that must be blocked by Sandbox. A maximum of 10000 MD5 file hashes can be
     * blocked
     */
    fileHashesToBeBlockeds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a SandboxBehavioralAnalysis resource.
 */
export interface SandboxBehavioralAnalysisArgs {
    /**
     * A custom list of unique MD5 file hashes that must be blocked by Sandbox. A maximum of 10000 MD5 file hashes can be
     * blocked
     */
    fileHashesToBeBlockeds?: pulumi.Input<pulumi.Input<string>[]>;
}
