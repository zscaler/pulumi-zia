// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ### Submit Raw Or Archive Files
 *
 * ### Submits Raw Or Archive For Out-Of-Band File Inspection
 */
export class SandboxFileSubmission extends pulumi.CustomResource {
    /**
     * Get an existing SandboxFileSubmission resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SandboxFileSubmissionState, opts?: pulumi.CustomResourceOptions): SandboxFileSubmission {
        return new SandboxFileSubmission(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zia:index/sandboxFileSubmission:SandboxFileSubmission';

    /**
     * Returns true if the given object is an instance of SandboxFileSubmission.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SandboxFileSubmission {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SandboxFileSubmission.__pulumiType;
    }

    public /*out*/ readonly code!: pulumi.Output<number>;
    /**
     * (Required) The path where the raw or archive files for submission are located.
     */
    public readonly filePath!: pulumi.Output<string>;
    public /*out*/ readonly fileType!: pulumi.Output<string>;
    /**
     * (Optional) Submit file to sandbox even if found malicious during AV scan and a verdict already exists. Supported values are `true` or `false`
     */
    public readonly force!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly md5!: pulumi.Output<string>;
    public /*out*/ readonly message!: pulumi.Output<string>;
    public /*out*/ readonly sandboxSubmission!: pulumi.Output<string>;
    /**
     * (Required) The submission method to be used. Supportedd values are: `submit` and `discan`
     */
    public readonly submissionMethod!: pulumi.Output<string>;
    public /*out*/ readonly virusName!: pulumi.Output<string>;
    public /*out*/ readonly virusType!: pulumi.Output<string>;

    /**
     * Create a SandboxFileSubmission resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SandboxFileSubmissionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SandboxFileSubmissionArgs | SandboxFileSubmissionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SandboxFileSubmissionState | undefined;
            resourceInputs["code"] = state ? state.code : undefined;
            resourceInputs["filePath"] = state ? state.filePath : undefined;
            resourceInputs["fileType"] = state ? state.fileType : undefined;
            resourceInputs["force"] = state ? state.force : undefined;
            resourceInputs["md5"] = state ? state.md5 : undefined;
            resourceInputs["message"] = state ? state.message : undefined;
            resourceInputs["sandboxSubmission"] = state ? state.sandboxSubmission : undefined;
            resourceInputs["submissionMethod"] = state ? state.submissionMethod : undefined;
            resourceInputs["virusName"] = state ? state.virusName : undefined;
            resourceInputs["virusType"] = state ? state.virusType : undefined;
        } else {
            const args = argsOrState as SandboxFileSubmissionArgs | undefined;
            if ((!args || args.filePath === undefined) && !opts.urn) {
                throw new Error("Missing required property 'filePath'");
            }
            if ((!args || args.submissionMethod === undefined) && !opts.urn) {
                throw new Error("Missing required property 'submissionMethod'");
            }
            resourceInputs["filePath"] = args ? args.filePath : undefined;
            resourceInputs["force"] = args ? args.force : undefined;
            resourceInputs["submissionMethod"] = args ? args.submissionMethod : undefined;
            resourceInputs["code"] = undefined /*out*/;
            resourceInputs["fileType"] = undefined /*out*/;
            resourceInputs["md5"] = undefined /*out*/;
            resourceInputs["message"] = undefined /*out*/;
            resourceInputs["sandboxSubmission"] = undefined /*out*/;
            resourceInputs["virusName"] = undefined /*out*/;
            resourceInputs["virusType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SandboxFileSubmission.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SandboxFileSubmission resources.
 */
export interface SandboxFileSubmissionState {
    code?: pulumi.Input<number>;
    /**
     * (Required) The path where the raw or archive files for submission are located.
     */
    filePath?: pulumi.Input<string>;
    fileType?: pulumi.Input<string>;
    /**
     * (Optional) Submit file to sandbox even if found malicious during AV scan and a verdict already exists. Supported values are `true` or `false`
     */
    force?: pulumi.Input<boolean>;
    md5?: pulumi.Input<string>;
    message?: pulumi.Input<string>;
    sandboxSubmission?: pulumi.Input<string>;
    /**
     * (Required) The submission method to be used. Supportedd values are: `submit` and `discan`
     */
    submissionMethod?: pulumi.Input<string>;
    virusName?: pulumi.Input<string>;
    virusType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SandboxFileSubmission resource.
 */
export interface SandboxFileSubmissionArgs {
    /**
     * (Required) The path where the raw or archive files for submission are located.
     */
    filePath: pulumi.Input<string>;
    /**
     * (Optional) Submit file to sandbox even if found malicious during AV scan and a verdict already exists. Supported values are `true` or `false`
     */
    force?: pulumi.Input<boolean>;
    /**
     * (Required) The submission method to be used. Supportedd values are: `submit` and `discan`
     */
    submissionMethod: pulumi.Input<string>;
}
