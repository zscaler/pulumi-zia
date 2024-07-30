// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The **zia_firewall_filtering_ip_source_groups** resource allows the creation and management of ZIA Cloud Firewall IP source groups in the Zscaler Internet Access. This resource can then be associated with a ZIA cloud firewall filtering rule.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@bdzscaler/pulumi-zia";
 *
 * // Add an IP address or addresses to a new IP Source Group
 * const example = new zia.FirewallFilteringSourceGroups("example", {
 *     description: "Example",
 *     ipAddresses: [
 *         "192.168.100.1",
 *         "192.168.100.2",
 *         "192.168.100.3",
 *     ],
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@bdzscaler/pulumi-zia";
 *
 * // Add an IP address range(s) to a new IP Source Group
 * const example = new zia.FirewallFilteringSourceGroups("example", {
 *     description: "Example",
 *     ipAddresses: ["192.0.2.1-192.0.2.10"],
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@bdzscaler/pulumi-zia";
 *
 * // Add subnet to a new IP Source Group
 * const example = new zia.FirewallFilteringSourceGroups("example", {
 *     description: "Example",
 *     ipAddresses: ["203.0.113.0/24"],
 * });
 * ```
 *
 * ## Import
 *
 * Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.
 *
 * Visit
 *
 * **zia_firewall_filtering_ip_source_groups** can be imported by using `<GROUP_ID>` or `<GROUP_NAME>` as the import ID.
 *
 * For example:
 *
 * ```sh
 * $ pulumi import zia:index/firewallFilteringSourceGroups:FirewallFilteringSourceGroups example <group_id>
 * ```
 *
 * or
 *
 * ```sh
 * $ pulumi import zia:index/firewallFilteringSourceGroups:FirewallFilteringSourceGroups example <group_name>
 * ```
 */
export class FirewallFilteringSourceGroups extends pulumi.CustomResource {
    /**
     * Get an existing FirewallFilteringSourceGroups resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallFilteringSourceGroupsState, opts?: pulumi.CustomResourceOptions): FirewallFilteringSourceGroups {
        return new FirewallFilteringSourceGroups(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zia:index/firewallFilteringSourceGroups:FirewallFilteringSourceGroups';

    /**
     * Returns true if the given object is an instance of FirewallFilteringSourceGroups.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallFilteringSourceGroups {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallFilteringSourceGroups.__pulumiType;
    }

    public readonly description!: pulumi.Output<string | undefined>;
    public /*out*/ readonly groupId!: pulumi.Output<number>;
    public readonly ipAddresses!: pulumi.Output<string[]>;
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a FirewallFilteringSourceGroups resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallFilteringSourceGroupsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallFilteringSourceGroupsArgs | FirewallFilteringSourceGroupsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallFilteringSourceGroupsState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["ipAddresses"] = state ? state.ipAddresses : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as FirewallFilteringSourceGroupsArgs | undefined;
            if ((!args || args.ipAddresses === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipAddresses'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ipAddresses"] = args ? args.ipAddresses : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["groupId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallFilteringSourceGroups.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallFilteringSourceGroups resources.
 */
export interface FirewallFilteringSourceGroupsState {
    description?: pulumi.Input<string>;
    groupId?: pulumi.Input<number>;
    ipAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallFilteringSourceGroups resource.
 */
export interface FirewallFilteringSourceGroupsArgs {
    description?: pulumi.Input<string>;
    ipAddresses: pulumi.Input<pulumi.Input<string>[]>;
    name?: pulumi.Input<string>;
}
