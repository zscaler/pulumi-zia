// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The **zia_firewall_filtering_destination_groups** resource allows the creation and management of ZIA Cloud Firewall IP destination groups in the Zscaler Internet Access. This resource can then be associated with a ZIA cloud firewall filtering rule.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@bdzscaler/pulumi-zia";
 *
 * // IP Destination Group of Type DSTN_FQDN
 * const dstnFqdn = new zia.FirewallFilteringDestinationGroups("dstnFqdn", {
 *     addresses: [
 *         "test1.acme.com",
 *         "test2.acme.com",
 *         "test3.acme.com",
 *     ],
 *     description: "Example Destination FQDN",
 *     type: "DSTN_FQDN",
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@bdzscaler/pulumi-zia";
 *
 * // IP Destination Group of Type DSTN_IP
 * const exampleDstnIp = new zia.FirewallFilteringDestinationGroups("exampleDstnIp", {
 *     addresses: [
 *         "3.217.228.0-3.217.231.255",
 *         "3.235.112.0-3.235.119.255",
 *         "52.23.61.0-52.23.62.25",
 *         "35.80.88.0-35.80.95.255",
 *     ],
 *     description: "Example Destination IP",
 *     type: "DSTN_IP",
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@bdzscaler/pulumi-zia";
 *
 * // IP Destination Group of Type DSTN_DOMAIN
 * const exampleDstnDomain = new zia.FirewallFilteringDestinationGroups("exampleDstnDomain", {
 *     addresses: [
 *         "acme.com",
 *         "acme1.com",
 *     ],
 *     description: "Example Destination Domain",
 *     type: "DSTN_DOMAIN",
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@bdzscaler/pulumi-zia";
 *
 * // IP Destination Group of Type DSTN_OTHER
 * const exampleDstnOther = new zia.FirewallFilteringDestinationGroups("exampleDstnOther", {
 *     countries: ["COUNTRY_CA"],
 *     description: "Example Destination Other",
 *     ipCategories: [
 *         "CUSTOM_01",
 *         "CUSTOM_02",
 *     ],
 *     type: "DSTN_OTHER",
 * });
 * ```
 *
 * ## Import
 *
 * Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.
 *
 * Visit
 *
 * **zia_firewall_filtering_destination_groups** can be imported by using `<GROUP_ID>` or `<GROUP_NAME>` as the import ID.
 *
 * For example:
 *
 * ```sh
 * $ pulumi import zia:index/firewallFilteringDestinationGroups:FirewallFilteringDestinationGroups example <group_id>
 * ```
 *
 * or
 *
 * ```sh
 * $ pulumi import zia:index/firewallFilteringDestinationGroups:FirewallFilteringDestinationGroups example <group_name>
 * ```
 */
export class FirewallFilteringDestinationGroups extends pulumi.CustomResource {
    /**
     * Get an existing FirewallFilteringDestinationGroups resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallFilteringDestinationGroupsState, opts?: pulumi.CustomResourceOptions): FirewallFilteringDestinationGroups {
        return new FirewallFilteringDestinationGroups(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zia:index/firewallFilteringDestinationGroups:FirewallFilteringDestinationGroups';

    /**
     * Returns true if the given object is an instance of FirewallFilteringDestinationGroups.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallFilteringDestinationGroups {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallFilteringDestinationGroups.__pulumiType;
    }

    /**
     * Destination IP addresses within the group
     */
    public readonly addresses!: pulumi.Output<string[]>;
    /**
     * Destination countries for which the rule is applicable. If not set, the rule is not restricted to specific destination
     * countries.
     */
    public readonly countries!: pulumi.Output<string[]>;
    /**
     * Additional information about the destination IP group
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Unique identifer for the destination IP group
     */
    public /*out*/ readonly groupId!: pulumi.Output<number>;
    /**
     * List of URL categories for which rule must be applied
     */
    public readonly ipCategories!: pulumi.Output<string[] | undefined>;
    /**
     * Destination IP group name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Destination IP group type (i.e., the group can contain destination IP addresses or FQDNs)
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a FirewallFilteringDestinationGroups resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallFilteringDestinationGroupsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallFilteringDestinationGroupsArgs | FirewallFilteringDestinationGroupsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallFilteringDestinationGroupsState | undefined;
            resourceInputs["addresses"] = state ? state.addresses : undefined;
            resourceInputs["countries"] = state ? state.countries : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["ipCategories"] = state ? state.ipCategories : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as FirewallFilteringDestinationGroupsArgs | undefined;
            resourceInputs["addresses"] = args ? args.addresses : undefined;
            resourceInputs["countries"] = args ? args.countries : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ipCategories"] = args ? args.ipCategories : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["groupId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallFilteringDestinationGroups.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallFilteringDestinationGroups resources.
 */
export interface FirewallFilteringDestinationGroupsState {
    /**
     * Destination IP addresses within the group
     */
    addresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Destination countries for which the rule is applicable. If not set, the rule is not restricted to specific destination
     * countries.
     */
    countries?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Additional information about the destination IP group
     */
    description?: pulumi.Input<string>;
    /**
     * Unique identifer for the destination IP group
     */
    groupId?: pulumi.Input<number>;
    /**
     * List of URL categories for which rule must be applied
     */
    ipCategories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Destination IP group name
     */
    name?: pulumi.Input<string>;
    /**
     * Destination IP group type (i.e., the group can contain destination IP addresses or FQDNs)
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallFilteringDestinationGroups resource.
 */
export interface FirewallFilteringDestinationGroupsArgs {
    /**
     * Destination IP addresses within the group
     */
    addresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Destination countries for which the rule is applicable. If not set, the rule is not restricted to specific destination
     * countries.
     */
    countries?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Additional information about the destination IP group
     */
    description?: pulumi.Input<string>;
    /**
     * List of URL categories for which rule must be applied
     */
    ipCategories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Destination IP group name
     */
    name?: pulumi.Input<string>;
    /**
     * Destination IP group type (i.e., the group can contain destination IP addresses or FQDNs)
     */
    type?: pulumi.Input<string>;
}
