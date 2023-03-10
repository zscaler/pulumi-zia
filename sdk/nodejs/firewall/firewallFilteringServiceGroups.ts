// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The **zia_firewall_filtering_network_service_groups** resource allows the creation and management of ZIA Cloud Firewall IP network service groups in the Zscaler Internet Access. This resource can then be associated with a ZIA cloud firewall filtering rule.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@pulumi/zia";
 * import * as zia from "@zscaler/pulumi-zia";
 *
 * const example1 = zia.Firewall.getFirewallFilteringNetworkServices({
 *     name: "FTP",
 * });
 * const example2 = zia.Firewall.getFirewallFilteringNetworkServices({
 *     name: "NETBIOS",
 * });
 * const example3 = zia.Firewall.getFirewallFilteringNetworkServices({
 *     name: "DNS",
 * });
 * // Add network services to a network services group
 * const example = new zia.firewall.FirewallFilteringServiceGroups("example", {
 *     description: "example",
 *     services: [{
 *         ids: [
 *             example1.then(example1 => example1.id),
 *             example2.then(example2 => example2.id),
 *             example3.then(example3 => example3.id),
 *         ],
 *     }],
 * });
 * ```
 */
export class FirewallFilteringServiceGroups extends pulumi.CustomResource {
    /**
     * Get an existing FirewallFilteringServiceGroups resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallFilteringServiceGroupsState, opts?: pulumi.CustomResourceOptions): FirewallFilteringServiceGroups {
        return new FirewallFilteringServiceGroups(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zia:Firewall/firewallFilteringServiceGroups:FirewallFilteringServiceGroups';

    /**
     * Returns true if the given object is an instance of FirewallFilteringServiceGroups.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallFilteringServiceGroups {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallFilteringServiceGroups.__pulumiType;
    }

    /**
     * Description of the network services group
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the network service group
     */
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly networkServiceGroupId!: pulumi.Output<number>;
    /**
     * Any number of network services ID to be added to the group
     */
    public readonly services!: pulumi.Output<outputs.Firewall.FirewallFilteringServiceGroupsService[]>;

    /**
     * Create a FirewallFilteringServiceGroups resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallFilteringServiceGroupsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallFilteringServiceGroupsArgs | FirewallFilteringServiceGroupsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallFilteringServiceGroupsState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkServiceGroupId"] = state ? state.networkServiceGroupId : undefined;
            resourceInputs["services"] = state ? state.services : undefined;
        } else {
            const args = argsOrState as FirewallFilteringServiceGroupsArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["services"] = args ? args.services : undefined;
            resourceInputs["networkServiceGroupId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallFilteringServiceGroups.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallFilteringServiceGroups resources.
 */
export interface FirewallFilteringServiceGroupsState {
    /**
     * Description of the network services group
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the network service group
     */
    name?: pulumi.Input<string>;
    networkServiceGroupId?: pulumi.Input<number>;
    /**
     * Any number of network services ID to be added to the group
     */
    services?: pulumi.Input<pulumi.Input<inputs.Firewall.FirewallFilteringServiceGroupsService>[]>;
}

/**
 * The set of arguments for constructing a FirewallFilteringServiceGroups resource.
 */
export interface FirewallFilteringServiceGroupsArgs {
    /**
     * Description of the network services group
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the network service group
     */
    name?: pulumi.Input<string>;
    /**
     * Any number of network services ID to be added to the group
     */
    services?: pulumi.Input<pulumi.Input<inputs.Firewall.FirewallFilteringServiceGroupsService>[]>;
}
