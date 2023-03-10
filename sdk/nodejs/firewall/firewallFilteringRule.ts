// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The **zia_firewall_filtering_rule** resource allows the creation and management of ZIA Cloud Firewall filtering rules in the Zscaler Internet Access.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@pulumi/zia";
 * import * as zia from "@zscaler/pulumi-zia";
 *
 * const zscalerProxyNwServices = zia.Firewall.getFirewallFilteringNetworkServices({
 *     name: "ZSCALER_PROXY_NW_SERVICES",
 * });
 * const engineering = zia.Departments.getDepartmentManagement({
 *     name: "Engineering",
 * });
 * const normalInternet = zia.Groups.getGroupManagement({
 *     name: "Normal_Internet",
 * });
 * const workHours = zia.TimeWindow.getTimeWindow({
 *     name: "Work hours",
 * });
 * const example = new zia.firewall.FirewallFilteringRule("example", {
 *     description: "Example",
 *     action: "ALLOW",
 *     state: "ENABLED",
 *     order: 1,
 *     enableFullLogging: true,
 *     nwServices: {
 *         ids: [zscalerProxyNwServices.then(zscalerProxyNwServices => zscalerProxyNwServices.id)],
 *     },
 *     departments: {
 *         ids: [engineering.then(engineering => engineering.id)],
 *     },
 *     groups: {
 *         ids: [normalInternet.then(normalInternet => normalInternet.id)],
 *     },
 *     timeWindows: {
 *         ids: [workHours.then(workHours => workHours.id)],
 *     },
 * });
 * ```
 */
export class FirewallFilteringRule extends pulumi.CustomResource {
    /**
     * Get an existing FirewallFilteringRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallFilteringRuleState, opts?: pulumi.CustomResourceOptions): FirewallFilteringRule {
        return new FirewallFilteringRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zia:Firewall/firewallFilteringRule:FirewallFilteringRule';

    /**
     * Returns true if the given object is an instance of FirewallFilteringRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallFilteringRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallFilteringRule.__pulumiType;
    }

    public readonly accessControl!: pulumi.Output<string>;
    /**
     * Choose the action of the service when packets match the rule. The following actions are accepted: `ALLOW`, `BLOCK_DROP`, `BLOCK_RESET`, `BLOCK_ICMP`, `EVAL_NWAPP`
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * Application service groups on which this rule is applied
     */
    public readonly appServiceGroups!: pulumi.Output<outputs.Firewall.FirewallFilteringRuleAppServiceGroups>;
    /**
     * Application services on which this rule is applied
     */
    public readonly appServices!: pulumi.Output<outputs.Firewall.FirewallFilteringRuleAppServices>;
    /**
     * If set to true, the default rule is applied
     */
    public readonly defaultRule!: pulumi.Output<boolean>;
    /**
     * Apply to any number of departments When not used it implies `Any` to apply the rule to all departments.
     */
    public readonly departments!: pulumi.Output<outputs.Firewall.FirewallFilteringRuleDepartments>;
    /**
     * Enter additional notes or information. The description cannot exceed 10,240 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * ** - (Optional) -  IP addresses and fully qualified domain names (FQDNs), if the domain has multiple destination IP addresses or if its IP addresses may change. For IP addresses, you can enter individual IP addresses, subnets, or address ranges. If adding multiple items, hit Enter after each entry.
     */
    public readonly destAddresses!: pulumi.Output<string[]>;
    /**
     * ** - (Optional) Identify destinations based on the location of a server, select Any to apply the rule to all countries or select the countries to which you want to control traffic.
     */
    public readonly destCountries!: pulumi.Output<string[]>;
    /**
     * ** - (Optional) identify destinations based on the URL category of the domain, select Any to apply the rule to all categories or select the specific categories you want to control.
     */
    public readonly destIpCategories!: pulumi.Output<string[] | undefined>;
    /**
     * ** - (Optional) Any number of destination IP address groups that you want to control with this rule.
     */
    public readonly destIpGroups!: pulumi.Output<outputs.Firewall.FirewallFilteringRuleDestIpGroups>;
    public readonly enableFullLogging!: pulumi.Output<boolean | undefined>;
    /**
     * You can manually select up to `8` groups. When not used it implies `Any` to apply the rule to all groups.
     */
    public readonly groups!: pulumi.Output<outputs.Firewall.FirewallFilteringRuleGroups>;
    /**
     * Labels that are applicable to the rule.
     */
    public readonly labels!: pulumi.Output<outputs.Firewall.FirewallFilteringRuleLabels>;
    public readonly lastModifiedBies!: pulumi.Output<outputs.Firewall.FirewallFilteringRuleLastModifiedBy[]>;
    public readonly lastModifiedTime!: pulumi.Output<number>;
    /**
     * You can manually select up to `32` location groups. When not used it implies `Any` to apply the rule to all location groups.
     */
    public readonly locationGroups!: pulumi.Output<outputs.Firewall.FirewallFilteringRuleLocationGroups>;
    /**
     * You can manually select up to `8` locations. When not used it implies `Any` to apply the rule to all groups.
     */
    public readonly locations!: pulumi.Output<outputs.Firewall.FirewallFilteringRuleLocations>;
    /**
     * Name of the network service group
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Any number of application groups that you want to control with this rule. The service provides predefined applications that you can group, but not modify
     */
    public readonly nwApplicationGroups!: pulumi.Output<outputs.Firewall.FirewallFilteringRuleNwApplicationGroups>;
    /**
     * When not used it applies the rule to all applications. The service provides predefined applications, which you can group, but not modify.
     */
    public readonly nwApplications!: pulumi.Output<string[]>;
    /**
     * Any number of predefined or custom network service groups to which the rule applies.
     */
    public readonly nwServiceGroups!: pulumi.Output<outputs.Firewall.FirewallFilteringRuleNwServiceGroups>;
    /**
     * When not used it applies the rule to all network services or you can select specific network services. The Zscaler firewall has predefined services and you can configure up to `1,024` additional custom services.
     */
    public readonly nwServices!: pulumi.Output<outputs.Firewall.FirewallFilteringRuleNwServices>;
    /**
     * Policy rules are evaluated in ascending numerical order (Rule 1 before Rule 2, and so on), and the Rule Order reflects this rule's place in the order.
     */
    public readonly order!: pulumi.Output<number>;
    /**
     * If set to true, a predefined rule is applied
     */
    public readonly predefined!: pulumi.Output<boolean>;
    /**
     * By default, the admin ranking is disabled. To use this feature, you must enable admin rank. The default value is `7`.
     */
    public readonly rank!: pulumi.Output<number>;
    public /*out*/ readonly ruleId!: pulumi.Output<number>;
    /**
     * Any number of source IP address groups that you want to control with this rule.
     */
    public readonly srcIpGroups!: pulumi.Output<outputs.Firewall.FirewallFilteringRuleSrcIpGroups>;
    /**
     * You can enter individual IP addresses, subnets, or address ranges.
     */
    public readonly srcIps!: pulumi.Output<string[] | undefined>;
    /**
     * An enabled rule is actively enforced. A disabled rule is not actively enforced but does not lose its place in the Rule Order. The service skips it and moves to the next rule.
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * You can manually select up to `2` time intervals. When not used it implies `always` to apply the rule to all time intervals.
     */
    public readonly timeWindows!: pulumi.Output<outputs.Firewall.FirewallFilteringRuleTimeWindows>;
    /**
     * You can manually select up to `4` general and/or special users. When not used it implies `Any` to apply the rule to all users.
     */
    public readonly users!: pulumi.Output<outputs.Firewall.FirewallFilteringRuleUsers>;

    /**
     * Create a FirewallFilteringRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallFilteringRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallFilteringRuleArgs | FirewallFilteringRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallFilteringRuleState | undefined;
            resourceInputs["accessControl"] = state ? state.accessControl : undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["appServiceGroups"] = state ? state.appServiceGroups : undefined;
            resourceInputs["appServices"] = state ? state.appServices : undefined;
            resourceInputs["defaultRule"] = state ? state.defaultRule : undefined;
            resourceInputs["departments"] = state ? state.departments : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["destAddresses"] = state ? state.destAddresses : undefined;
            resourceInputs["destCountries"] = state ? state.destCountries : undefined;
            resourceInputs["destIpCategories"] = state ? state.destIpCategories : undefined;
            resourceInputs["destIpGroups"] = state ? state.destIpGroups : undefined;
            resourceInputs["enableFullLogging"] = state ? state.enableFullLogging : undefined;
            resourceInputs["groups"] = state ? state.groups : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["lastModifiedBies"] = state ? state.lastModifiedBies : undefined;
            resourceInputs["lastModifiedTime"] = state ? state.lastModifiedTime : undefined;
            resourceInputs["locationGroups"] = state ? state.locationGroups : undefined;
            resourceInputs["locations"] = state ? state.locations : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nwApplicationGroups"] = state ? state.nwApplicationGroups : undefined;
            resourceInputs["nwApplications"] = state ? state.nwApplications : undefined;
            resourceInputs["nwServiceGroups"] = state ? state.nwServiceGroups : undefined;
            resourceInputs["nwServices"] = state ? state.nwServices : undefined;
            resourceInputs["order"] = state ? state.order : undefined;
            resourceInputs["predefined"] = state ? state.predefined : undefined;
            resourceInputs["rank"] = state ? state.rank : undefined;
            resourceInputs["ruleId"] = state ? state.ruleId : undefined;
            resourceInputs["srcIpGroups"] = state ? state.srcIpGroups : undefined;
            resourceInputs["srcIps"] = state ? state.srcIps : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["timeWindows"] = state ? state.timeWindows : undefined;
            resourceInputs["users"] = state ? state.users : undefined;
        } else {
            const args = argsOrState as FirewallFilteringRuleArgs | undefined;
            resourceInputs["accessControl"] = args ? args.accessControl : undefined;
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["appServiceGroups"] = args ? args.appServiceGroups : undefined;
            resourceInputs["appServices"] = args ? args.appServices : undefined;
            resourceInputs["defaultRule"] = args ? args.defaultRule : undefined;
            resourceInputs["departments"] = args ? args.departments : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destAddresses"] = args ? args.destAddresses : undefined;
            resourceInputs["destCountries"] = args ? args.destCountries : undefined;
            resourceInputs["destIpCategories"] = args ? args.destIpCategories : undefined;
            resourceInputs["destIpGroups"] = args ? args.destIpGroups : undefined;
            resourceInputs["enableFullLogging"] = args ? args.enableFullLogging : undefined;
            resourceInputs["groups"] = args ? args.groups : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["lastModifiedBies"] = args ? args.lastModifiedBies : undefined;
            resourceInputs["lastModifiedTime"] = args ? args.lastModifiedTime : undefined;
            resourceInputs["locationGroups"] = args ? args.locationGroups : undefined;
            resourceInputs["locations"] = args ? args.locations : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nwApplicationGroups"] = args ? args.nwApplicationGroups : undefined;
            resourceInputs["nwApplications"] = args ? args.nwApplications : undefined;
            resourceInputs["nwServiceGroups"] = args ? args.nwServiceGroups : undefined;
            resourceInputs["nwServices"] = args ? args.nwServices : undefined;
            resourceInputs["order"] = args ? args.order : undefined;
            resourceInputs["predefined"] = args ? args.predefined : undefined;
            resourceInputs["rank"] = args ? args.rank : undefined;
            resourceInputs["srcIpGroups"] = args ? args.srcIpGroups : undefined;
            resourceInputs["srcIps"] = args ? args.srcIps : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["timeWindows"] = args ? args.timeWindows : undefined;
            resourceInputs["users"] = args ? args.users : undefined;
            resourceInputs["ruleId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallFilteringRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallFilteringRule resources.
 */
export interface FirewallFilteringRuleState {
    accessControl?: pulumi.Input<string>;
    /**
     * Choose the action of the service when packets match the rule. The following actions are accepted: `ALLOW`, `BLOCK_DROP`, `BLOCK_RESET`, `BLOCK_ICMP`, `EVAL_NWAPP`
     */
    action?: pulumi.Input<string>;
    /**
     * Application service groups on which this rule is applied
     */
    appServiceGroups?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleAppServiceGroups>;
    /**
     * Application services on which this rule is applied
     */
    appServices?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleAppServices>;
    /**
     * If set to true, the default rule is applied
     */
    defaultRule?: pulumi.Input<boolean>;
    /**
     * Apply to any number of departments When not used it implies `Any` to apply the rule to all departments.
     */
    departments?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleDepartments>;
    /**
     * Enter additional notes or information. The description cannot exceed 10,240 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * ** - (Optional) -  IP addresses and fully qualified domain names (FQDNs), if the domain has multiple destination IP addresses or if its IP addresses may change. For IP addresses, you can enter individual IP addresses, subnets, or address ranges. If adding multiple items, hit Enter after each entry.
     */
    destAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ** - (Optional) Identify destinations based on the location of a server, select Any to apply the rule to all countries or select the countries to which you want to control traffic.
     */
    destCountries?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ** - (Optional) identify destinations based on the URL category of the domain, select Any to apply the rule to all categories or select the specific categories you want to control.
     */
    destIpCategories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ** - (Optional) Any number of destination IP address groups that you want to control with this rule.
     */
    destIpGroups?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleDestIpGroups>;
    enableFullLogging?: pulumi.Input<boolean>;
    /**
     * You can manually select up to `8` groups. When not used it implies `Any` to apply the rule to all groups.
     */
    groups?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleGroups>;
    /**
     * Labels that are applicable to the rule.
     */
    labels?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleLabels>;
    lastModifiedBies?: pulumi.Input<pulumi.Input<inputs.Firewall.FirewallFilteringRuleLastModifiedBy>[]>;
    lastModifiedTime?: pulumi.Input<number>;
    /**
     * You can manually select up to `32` location groups. When not used it implies `Any` to apply the rule to all location groups.
     */
    locationGroups?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleLocationGroups>;
    /**
     * You can manually select up to `8` locations. When not used it implies `Any` to apply the rule to all groups.
     */
    locations?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleLocations>;
    /**
     * Name of the network service group
     */
    name?: pulumi.Input<string>;
    /**
     * Any number of application groups that you want to control with this rule. The service provides predefined applications that you can group, but not modify
     */
    nwApplicationGroups?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleNwApplicationGroups>;
    /**
     * When not used it applies the rule to all applications. The service provides predefined applications, which you can group, but not modify.
     */
    nwApplications?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Any number of predefined or custom network service groups to which the rule applies.
     */
    nwServiceGroups?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleNwServiceGroups>;
    /**
     * When not used it applies the rule to all network services or you can select specific network services. The Zscaler firewall has predefined services and you can configure up to `1,024` additional custom services.
     */
    nwServices?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleNwServices>;
    /**
     * Policy rules are evaluated in ascending numerical order (Rule 1 before Rule 2, and so on), and the Rule Order reflects this rule's place in the order.
     */
    order?: pulumi.Input<number>;
    /**
     * If set to true, a predefined rule is applied
     */
    predefined?: pulumi.Input<boolean>;
    /**
     * By default, the admin ranking is disabled. To use this feature, you must enable admin rank. The default value is `7`.
     */
    rank?: pulumi.Input<number>;
    ruleId?: pulumi.Input<number>;
    /**
     * Any number of source IP address groups that you want to control with this rule.
     */
    srcIpGroups?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleSrcIpGroups>;
    /**
     * You can enter individual IP addresses, subnets, or address ranges.
     */
    srcIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An enabled rule is actively enforced. A disabled rule is not actively enforced but does not lose its place in the Rule Order. The service skips it and moves to the next rule.
     */
    state?: pulumi.Input<string>;
    /**
     * You can manually select up to `2` time intervals. When not used it implies `always` to apply the rule to all time intervals.
     */
    timeWindows?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleTimeWindows>;
    /**
     * You can manually select up to `4` general and/or special users. When not used it implies `Any` to apply the rule to all users.
     */
    users?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleUsers>;
}

/**
 * The set of arguments for constructing a FirewallFilteringRule resource.
 */
export interface FirewallFilteringRuleArgs {
    accessControl?: pulumi.Input<string>;
    /**
     * Choose the action of the service when packets match the rule. The following actions are accepted: `ALLOW`, `BLOCK_DROP`, `BLOCK_RESET`, `BLOCK_ICMP`, `EVAL_NWAPP`
     */
    action?: pulumi.Input<string>;
    /**
     * Application service groups on which this rule is applied
     */
    appServiceGroups?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleAppServiceGroups>;
    /**
     * Application services on which this rule is applied
     */
    appServices?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleAppServices>;
    /**
     * If set to true, the default rule is applied
     */
    defaultRule?: pulumi.Input<boolean>;
    /**
     * Apply to any number of departments When not used it implies `Any` to apply the rule to all departments.
     */
    departments?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleDepartments>;
    /**
     * Enter additional notes or information. The description cannot exceed 10,240 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * ** - (Optional) -  IP addresses and fully qualified domain names (FQDNs), if the domain has multiple destination IP addresses or if its IP addresses may change. For IP addresses, you can enter individual IP addresses, subnets, or address ranges. If adding multiple items, hit Enter after each entry.
     */
    destAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ** - (Optional) Identify destinations based on the location of a server, select Any to apply the rule to all countries or select the countries to which you want to control traffic.
     */
    destCountries?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ** - (Optional) identify destinations based on the URL category of the domain, select Any to apply the rule to all categories or select the specific categories you want to control.
     */
    destIpCategories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ** - (Optional) Any number of destination IP address groups that you want to control with this rule.
     */
    destIpGroups?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleDestIpGroups>;
    enableFullLogging?: pulumi.Input<boolean>;
    /**
     * You can manually select up to `8` groups. When not used it implies `Any` to apply the rule to all groups.
     */
    groups?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleGroups>;
    /**
     * Labels that are applicable to the rule.
     */
    labels?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleLabels>;
    lastModifiedBies?: pulumi.Input<pulumi.Input<inputs.Firewall.FirewallFilteringRuleLastModifiedBy>[]>;
    lastModifiedTime?: pulumi.Input<number>;
    /**
     * You can manually select up to `32` location groups. When not used it implies `Any` to apply the rule to all location groups.
     */
    locationGroups?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleLocationGroups>;
    /**
     * You can manually select up to `8` locations. When not used it implies `Any` to apply the rule to all groups.
     */
    locations?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleLocations>;
    /**
     * Name of the network service group
     */
    name?: pulumi.Input<string>;
    /**
     * Any number of application groups that you want to control with this rule. The service provides predefined applications that you can group, but not modify
     */
    nwApplicationGroups?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleNwApplicationGroups>;
    /**
     * When not used it applies the rule to all applications. The service provides predefined applications, which you can group, but not modify.
     */
    nwApplications?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Any number of predefined or custom network service groups to which the rule applies.
     */
    nwServiceGroups?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleNwServiceGroups>;
    /**
     * When not used it applies the rule to all network services or you can select specific network services. The Zscaler firewall has predefined services and you can configure up to `1,024` additional custom services.
     */
    nwServices?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleNwServices>;
    /**
     * Policy rules are evaluated in ascending numerical order (Rule 1 before Rule 2, and so on), and the Rule Order reflects this rule's place in the order.
     */
    order?: pulumi.Input<number>;
    /**
     * If set to true, a predefined rule is applied
     */
    predefined?: pulumi.Input<boolean>;
    /**
     * By default, the admin ranking is disabled. To use this feature, you must enable admin rank. The default value is `7`.
     */
    rank?: pulumi.Input<number>;
    /**
     * Any number of source IP address groups that you want to control with this rule.
     */
    srcIpGroups?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleSrcIpGroups>;
    /**
     * You can enter individual IP addresses, subnets, or address ranges.
     */
    srcIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An enabled rule is actively enforced. A disabled rule is not actively enforced but does not lose its place in the Rule Order. The service skips it and moves to the next rule.
     */
    state?: pulumi.Input<string>;
    /**
     * You can manually select up to `2` time intervals. When not used it implies `always` to apply the rule to all time intervals.
     */
    timeWindows?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleTimeWindows>;
    /**
     * You can manually select up to `4` general and/or special users. When not used it implies `Any` to apply the rule to all users.
     */
    users?: pulumi.Input<inputs.Firewall.FirewallFilteringRuleUsers>;
}
