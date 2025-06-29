// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zia/ips-control-policy#/firewallIpsRules-get)
 * * [API documentation](https://help.zscaler.com/zia/configuring-ips-control-policy)
 *
 * Use the **zia_firewall_ips_rule** data source to get information about a cloud firewall IPS rule available in the Zscaler Internet Access.
 *
 * ## Example Usage
 */
export function getIPSFirewallRule(args?: GetIPSFirewallRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetIPSFirewallRuleResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zia:index/getIPSFirewallRule:getIPSFirewallRule", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getIPSFirewallRule.
 */
export interface GetIPSFirewallRuleArgs {
    /**
     * Unique identifier for the Firewall Filtering policy rule
     */
    id?: number;
    /**
     * Name of the Firewall Filtering policy rule
     */
    name?: string;
}

/**
 * A collection of values returned by getIPSFirewallRule.
 */
export interface GetIPSFirewallRuleResult {
    /**
     * (String) The action configured for the rule that must take place if the traffic matches the rule criteria, such as allowing or blocking the traffic or bypassing the rule. The following actions are accepted: `ALLOW`, `BLOCK_DROP`, `BLOCK_RESET`, `BYPASS_IPS`
     */
    readonly action: string;
    /**
     * (Boolean) Value that indicates whether packet capture (PCAP) is enabled or not
     */
    readonly capturePcap: boolean;
    /**
     * (Boolean) Value that indicates whether the rule is the Default Cloud IPS Rule or not
     */
    readonly defaultRule: boolean;
    /**
     * (List of Objects) Apply to any number of departments When not used it implies `Any` to apply the rule to all departments.
     */
    readonly departments: outputs.GetIPSFirewallRuleDepartment[];
    /**
     * (String) Enter additional notes or information. The description cannot exceed 10,240 characters.
     */
    readonly description: string;
    /**
     * (Set of String) Destination IP addresses or FQDNs to which the rule applies. If not set, the rule is not restricted to a specific destination IP address. Each IP entry can be a single IP address, CIDR (e.g., 10.10.33.0/24), or an IP range (e.g., 10.10.33.1-10.10.33.10).
     */
    readonly destAddresses: string[];
    /**
     * (Set of String) Identify destinations based on the location of a server, select Any to apply the rule to all countries or select the countries to which you want to control traffic.
     * **NOTE**: Provide a 2 letter [ISO3166 Alpha2 Country code](https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes). i.e ``"US"``, ``"CA"``
     */
    readonly destCountries: string[];
    /**
     * (Set of String)  identify destinations based on the URL category of the domain, select Any to apply the rule to all categories or select the specific categories you want to control.
     */
    readonly destIpCategories: string[];
    /**
     * ** - (List of Objects) Any number of destination IP address groups that you want to control with this rule.
     */
    readonly destIpGroups: outputs.GetIPSFirewallRuleDestIpGroup[];
    readonly destIpv6Groups: outputs.GetIPSFirewallRuleDestIpv6Group[];
    /**
     * (List of Objects) Device groups to which the rule applies. This field is applicable for devices that are managed using Zscaler Client Connector. If no value is set, this field is ignored during the policy evaluation.
     */
    readonly deviceGroups: outputs.GetIPSFirewallRuleDeviceGroup[];
    /**
     * (List of Objects) Devices to which the rule applies. This field is applicable for devices that are managed using Zscaler Client Connector. If no value is set, this field is ignored during the policy evaluation.
     */
    readonly devices: outputs.GetIPSFirewallRuleDevice[];
    /**
     * (Integer) A Boolean value that indicates whether full logging is enabled. A true value indicates that full logging is enabled, whereas a false value indicates that aggregate logging is enabled.
     */
    readonly enableFullLogging: boolean;
    /**
     * (List of Objects) You can manually select up to `8` groups. When not used it implies `Any` to apply the rule to all groups.
     */
    readonly groups: outputs.GetIPSFirewallRuleGroup[];
    /**
     * (Integer) Identifier that uniquely identifies an entity
     */
    readonly id: number;
    /**
     * (List of Objects) Labels that are applicable to the rule.
     */
    readonly labels: outputs.GetIPSFirewallRuleLabel[];
    readonly lastModifiedBies: outputs.GetIPSFirewallRuleLastModifiedBy[];
    readonly lastModifiedTime: number;
    /**
     * (List of Objects)You can manually select up to `32` location groups. When not used it implies `Any` to apply the rule to all location groups.
     */
    readonly locationGroups: outputs.GetIPSFirewallRuleLocationGroup[];
    /**
     * (List of Objects) You can manually select up to `8` locations. When not used it implies `Any` to apply the rule to all groups.
     */
    readonly locations: outputs.GetIPSFirewallRuleLocation[];
    /**
     * (String) The configured name of the entity
     */
    readonly name: string;
    /**
     * (List of Objects) Any number of predefined or custom network service groups to which the rule applies.
     */
    readonly nwServiceGroups: outputs.GetIPSFirewallRuleNwServiceGroup[];
    /**
     * (List of Objects) When not used it applies the rule to all network services or you can select specific network services. The Zscaler firewall has predefined services and you can configure up to `1,024` additional custom services.
     */
    readonly nwServices: outputs.GetIPSFirewallRuleNwService[];
    /**
     * (Integer) Policy rules are evaluated in ascending numerical order (Rule 1 before Rule 2, and so on), and the Rule Order reflects this rule's place in the order.
     */
    readonly order: number;
    /**
     * (Boolean) A Boolean field that indicates that the rule is predefined by using a true value
     */
    readonly predefined: boolean;
    /**
     * (Integer) By default, the admin ranking is disabled. To use this feature, you must enable admin rank. The default value is `7`.
     */
    readonly rank: number;
    /**
     * (Set of String) URL categories associated with resolved IP addresses to which the rule applies. If not set, the rule is not restricted to a specific URL category.
     */
    readonly resCategories: string[];
    /**
     * (Set of String) The countries of origin of traffic for which the rule is applicable. If not set, the rule is not restricted to specific source countries.
     * **NOTE**: Provide a 2 letter [ISO3166 Alpha2 Country code](https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes). i.e ``"US"``, ``"CA"``
     */
    readonly sourceCountries: string[];
    /**
     * (List of Objects)Source IP address groups for which the rule is applicable. If not set, the rule is not restricted to a specific source IP address group.
     */
    readonly srcIpGroups: outputs.GetIPSFirewallRuleSrcIpGroup[];
    /**
     * (Set of String) Source IP addresses or FQDNs to which the rule applies. If not set, the rule is not restricted to a specific source IP address. Each IP entry can be a single IP address, CIDR (e.g., 10.10.33.0/24), or an IP range (e.g., 10.10.33.1-10.10.33.10).
     */
    readonly srcIps: string[];
    /**
     * (List of Objects) Source IPv6 address groups for which the rule is applicable. If not set, the rule is not restricted to a specific source IPv6 address group.
     */
    readonly srcIpv6Groups: outputs.GetIPSFirewallRuleSrcIpv6Group[];
    /**
     * (String) An enabled rule is actively enforced. A disabled rule is not actively enforced but does not lose its place in the Rule Order. The service skips it and moves to the next rule.
     */
    readonly state: string;
    /**
     * (List of Objects) Advanced threat categories to which the rule applies
     */
    readonly threatCategories: outputs.GetIPSFirewallRuleThreatCategory[];
    /**
     * (List of Objects) You can manually select up to `1` time intervals. When not used it implies `always` to apply the rule to all time intervals.
     */
    readonly timeWindows: outputs.GetIPSFirewallRuleTimeWindow[];
    /**
     * (List of Objects) You can manually select up to `4` general and/or special users. When not used it implies `Any` to apply the rule to all users.
     */
    readonly users: outputs.GetIPSFirewallRuleUser[];
    /**
     * (List of Objects) The ZPA application segments to which the rule applies
     */
    readonly zpaAppSegments: outputs.GetIPSFirewallRuleZpaAppSegment[];
}
/**
 * * [Official documentation](https://help.zscaler.com/zia/ips-control-policy#/firewallIpsRules-get)
 * * [API documentation](https://help.zscaler.com/zia/configuring-ips-control-policy)
 *
 * Use the **zia_firewall_ips_rule** data source to get information about a cloud firewall IPS rule available in the Zscaler Internet Access.
 *
 * ## Example Usage
 */
export function getIPSFirewallRuleOutput(args?: GetIPSFirewallRuleOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetIPSFirewallRuleResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("zia:index/getIPSFirewallRule:getIPSFirewallRule", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getIPSFirewallRule.
 */
export interface GetIPSFirewallRuleOutputArgs {
    /**
     * Unique identifier for the Firewall Filtering policy rule
     */
    id?: pulumi.Input<number>;
    /**
     * Name of the Firewall Filtering policy rule
     */
    name?: pulumi.Input<string>;
}
