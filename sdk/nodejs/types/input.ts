// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export namespace AdminUsers {
    export interface AdminUsersAdminScope {
        /**
         * Based on the admin scope type, the entities can be the ID/name pair of departments, locations, or location groups.
         */
        scopeEntities?: pulumi.Input<inputs.AdminUsers.AdminUsersAdminScopeScopeEntities>;
        /**
         * Only applicable for the LOCATION_GROUP admin scope type, in which case this attribute gives the list of ID/name pairs of locations within the location group.
         */
        scopeGroupMemberEntities?: pulumi.Input<inputs.AdminUsers.AdminUsersAdminScopeScopeGroupMemberEntities>;
        /**
         * The admin scope type. The attribute name is subject to change.
         */
        type?: pulumi.Input<string>;
    }

    export interface AdminUsersAdminScopeScopeEntities {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface AdminUsersAdminScopeScopeGroupMemberEntities {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface AdminUsersRole {
        extensions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        /**
         * Identifier that uniquely identifies an entity
         */
        id?: pulumi.Input<number>;
        isNameL10nTag?: pulumi.Input<boolean>;
        /**
         * The configured name of the entity
         */
        name?: pulumi.Input<string>;
    }

}

export namespace DLP {
    export interface DLPDictionariesExactDataMatchDetail {
        /**
         * The unique identifier for the EDM mapping.
         */
        dictionaryEdmMappingId?: pulumi.Input<number>;
        /**
         * The EDM template's primary field.
         */
        primaryField?: pulumi.Input<number>;
        /**
         * The unique identifier for the EDM template (or schema).
         */
        schemaId?: pulumi.Input<number>;
        /**
         * The EDM secondary field to match on.
         * - `"MATCHON_NONE"`
         * - `"MATCHON_ANY_1"`
         * - `"MATCHON_ANY_2"`
         * - `"MATCHON_ANY_3"`
         * - `"MATCHON_ANY_4"`
         * - `"MATCHON_ANY_5"`
         * - `"MATCHON_ANY_6"`
         * - `"MATCHON_ANY_7"`
         * - `"MATCHON_ANY_8"`
         * - `"MATCHON_ANY_9"`
         * - `"MATCHON_ANY_10"`
         * - `"MATCHON_ANY_11"`
         * - `"MATCHON_ANY_12"`
         * - `"MATCHON_ANY_13"`
         * - `"MATCHON_ANY_14"`
         * - `"MATCHON_ANY_15"`
         * - `"MATCHON_ALL"`
         */
        secondaryFieldMatchOn?: pulumi.Input<string>;
        /**
         * The EDM template's secondary fields.
         */
        secondaryFields?: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface DLPDictionariesIdmProfileMatchAccuracy {
        /**
         * The IDM template reference.
         */
        adpIdmProfile?: pulumi.Input<inputs.DLP.DLPDictionariesIdmProfileMatchAccuracyAdpIdmProfile>;
        /**
         * The IDM template match accuracy.
         * - `"LOW"`
         * - `"MEDIUM"`
         * - `"HEAVY"`
         */
        matchAccuracy?: pulumi.Input<string>;
    }

    export interface DLPDictionariesIdmProfileMatchAccuracyAdpIdmProfile {
        extensions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        id?: pulumi.Input<number>;
    }

    export interface DLPDictionariesPattern {
        /**
         * The action applied to a DLP dictionary using patterns. The following values are supported:
         */
        action?: pulumi.Input<string>;
        /**
         * DLP dictionary pattern
         */
        pattern?: pulumi.Input<string>;
    }

    export interface DLPDictionariesPhrase {
        /**
         * The action applied to a DLP dictionary using patterns. The following values are supported:
         */
        action?: pulumi.Input<string>;
        /**
         * DLP dictionary phrase
         */
        phrase?: pulumi.Input<string>;
    }

    export interface DLPWebRulesAuditor {
        /**
         * Identifier that uniquely identifies an entity
         */
        id: pulumi.Input<number>;
    }

    export interface DLPWebRulesDepartments {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface DLPWebRulesDlpEngines {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface DLPWebRulesExcludedDepartments {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface DLPWebRulesExcludedGroups {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface DLPWebRulesExcludedUsers {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface DLPWebRulesGroups {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface DLPWebRulesIcapServer {
        /**
         * Identifier that uniquely identifies an entity
         */
        id: pulumi.Input<number>;
    }

    export interface DLPWebRulesLabels {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface DLPWebRulesLocationGroups {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface DLPWebRulesLocations {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface DLPWebRulesNotificationTemplate {
        /**
         * Identifier that uniquely identifies an entity
         */
        id: pulumi.Input<number>;
    }

    export interface DLPWebRulesTimeWindows {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface DLPWebRulesUrlCategories {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface DLPWebRulesUsers {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

}

export namespace Firewall {
    export interface FirewallFilteringNetworkServicesDestTcpPort {
        end?: pulumi.Input<number>;
        start?: pulumi.Input<number>;
    }

    export interface FirewallFilteringNetworkServicesDestUdpPort {
        end?: pulumi.Input<number>;
        start?: pulumi.Input<number>;
    }

    export interface FirewallFilteringNetworkServicesSrcTcpPort {
        end?: pulumi.Input<number>;
        start?: pulumi.Input<number>;
    }

    export interface FirewallFilteringNetworkServicesSrcUdpPort {
        end?: pulumi.Input<number>;
        start?: pulumi.Input<number>;
    }

    export interface FirewallFilteringRuleAppServiceGroups {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface FirewallFilteringRuleAppServices {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface FirewallFilteringRuleDepartments {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface FirewallFilteringRuleDestIpGroups {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface FirewallFilteringRuleGroups {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface FirewallFilteringRuleLabels {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface FirewallFilteringRuleLastModifiedBy {
        extensions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        /**
         * Identifier that uniquely identifies an entity
         */
        id?: pulumi.Input<number>;
    }

    export interface FirewallFilteringRuleLocationGroups {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface FirewallFilteringRuleLocations {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface FirewallFilteringRuleNwApplicationGroups {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface FirewallFilteringRuleNwServiceGroups {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface FirewallFilteringRuleNwServices {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface FirewallFilteringRuleSrcIpGroups {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface FirewallFilteringRuleTimeWindows {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface FirewallFilteringRuleUsers {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface FirewallFilteringServiceGroupsService {
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

}

export namespace LocationGroups {
    export interface GetLocationGroupsDynamicLocationGroupCriteria {
        /**
         * (Block List)
         */
        cities?: inputs.LocationGroups.GetLocationGroupsDynamicLocationGroupCriteriaCity[];
        /**
         * (List of String) One or more countries from a predefined set
         */
        countries?: string[];
        /**
         * (Boolean) Enable Bandwidth Control. When set to true, Bandwidth Control is enabled for the location.
         */
        enableBandwidthControl?: boolean;
        /**
         * (Boolean) Enable Caution. When set to true, a caution notifcation is enabled for the location.
         */
        enableCaution?: boolean;
        /**
         * (Boolean) Enable `XFF` Forwarding. When set to true, traffic is passed to Zscaler Cloud via the X-Forwarded-For (XFF) header.
         */
        enableXffForwarding?: boolean;
        /**
         * (Boolean) Enable AUP. When set to true, AUP is enabled for the location.
         */
        enforceAup?: boolean;
        /**
         * (Boolean) Enforce Authentication. Required when ports are enabled, IP Surrogate is enabled, or Kerberos Authentication is enabled.
         */
        enforceAuthentication?: boolean;
        /**
         * (Boolean) Enable Firewall. When set to true, Firewall is enabled for the location.
         */
        enforceFirewallControl?: boolean;
        /**
         * (Block List)
         */
        managedBies?: inputs.LocationGroups.GetLocationGroupsDynamicLocationGroupCriteriaManagedBy[];
        /**
         * Location group name
         */
        names?: inputs.LocationGroups.GetLocationGroupsDynamicLocationGroupCriteriaName[];
        /**
         * (List of String) One or more location profiles from a predefined set
         */
        profiles?: string[];
    }

    export interface GetLocationGroupsDynamicLocationGroupCriteriaArgs {
        /**
         * (Block List)
         */
        cities?: pulumi.Input<pulumi.Input<inputs.LocationGroups.GetLocationGroupsDynamicLocationGroupCriteriaCityArgs>[]>;
        /**
         * (List of String) One or more countries from a predefined set
         */
        countries?: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * (Boolean) Enable Bandwidth Control. When set to true, Bandwidth Control is enabled for the location.
         */
        enableBandwidthControl?: pulumi.Input<boolean>;
        /**
         * (Boolean) Enable Caution. When set to true, a caution notifcation is enabled for the location.
         */
        enableCaution?: pulumi.Input<boolean>;
        /**
         * (Boolean) Enable `XFF` Forwarding. When set to true, traffic is passed to Zscaler Cloud via the X-Forwarded-For (XFF) header.
         */
        enableXffForwarding?: pulumi.Input<boolean>;
        /**
         * (Boolean) Enable AUP. When set to true, AUP is enabled for the location.
         */
        enforceAup?: pulumi.Input<boolean>;
        /**
         * (Boolean) Enforce Authentication. Required when ports are enabled, IP Surrogate is enabled, or Kerberos Authentication is enabled.
         */
        enforceAuthentication?: pulumi.Input<boolean>;
        /**
         * (Boolean) Enable Firewall. When set to true, Firewall is enabled for the location.
         */
        enforceFirewallControl?: pulumi.Input<boolean>;
        /**
         * (Block List)
         */
        managedBies?: pulumi.Input<pulumi.Input<inputs.LocationGroups.GetLocationGroupsDynamicLocationGroupCriteriaManagedByArgs>[]>;
        /**
         * Location group name
         */
        names?: pulumi.Input<pulumi.Input<inputs.LocationGroups.GetLocationGroupsDynamicLocationGroupCriteriaNameArgs>[]>;
        /**
         * (List of String) One or more location profiles from a predefined set
         */
        profiles?: pulumi.Input<pulumi.Input<string>[]>;
    }

    export interface GetLocationGroupsDynamicLocationGroupCriteriaCity {
        /**
         * (String) String value to be matched or partially matched
         */
        matchString?: string;
        /**
         * (String) Operator that performs match action
         */
        matchType?: string;
    }

    export interface GetLocationGroupsDynamicLocationGroupCriteriaCityArgs {
        /**
         * (String) String value to be matched or partially matched
         */
        matchString?: pulumi.Input<string>;
        /**
         * (String) Operator that performs match action
         */
        matchType?: pulumi.Input<string>;
    }

    export interface GetLocationGroupsDynamicLocationGroupCriteriaManagedBy {
        /**
         * (Map of String)
         */
        extensions?: {[key: string]: string};
        /**
         * Unique identifier for the location group
         */
        id?: number;
        /**
         * Location group name
         */
        name?: string;
    }

    export interface GetLocationGroupsDynamicLocationGroupCriteriaManagedByArgs {
        /**
         * (Map of String)
         */
        extensions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        /**
         * Unique identifier for the location group
         */
        id?: pulumi.Input<number>;
        /**
         * Location group name
         */
        name?: pulumi.Input<string>;
    }

    export interface GetLocationGroupsDynamicLocationGroupCriteriaName {
        /**
         * (String) String value to be matched or partially matched
         */
        matchString?: string;
        /**
         * (String) Operator that performs match action
         */
        matchType?: string;
    }

    export interface GetLocationGroupsDynamicLocationGroupCriteriaNameArgs {
        /**
         * (String) String value to be matched or partially matched
         */
        matchString?: pulumi.Input<string>;
        /**
         * (String) Operator that performs match action
         */
        matchType?: pulumi.Input<string>;
    }

}

export namespace LocationManagement {
    export interface LocationManagementVpnCredential {
        comments?: pulumi.Input<string>;
        fqdn?: pulumi.Input<string>;
        /**
         * VPN credential resource id. The value is required if `ipAddresses` are not defined.
         */
        id?: pulumi.Input<number>;
        ipAddress?: pulumi.Input<string>;
        preSharedKey?: pulumi.Input<string>;
        type?: pulumi.Input<string>;
    }
}

export namespace RuleLabels {
    export interface RuleLabelsCreatedBy {
        extensions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        id?: pulumi.Input<number>;
        /**
         * The name of the devices to be created.
         */
        name?: pulumi.Input<string>;
    }

    export interface RuleLabelsLastModifiedBy {
        extensions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        id?: pulumi.Input<number>;
        /**
         * The name of the devices to be created.
         */
        name?: pulumi.Input<string>;
    }
}

export namespace TrafficForwarding {
    export interface TrafficForwardingGRETunnelLastModifiedBy {
        extensions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        /**
         * Unique identifer of the GRE virtual IP address (VIP)
         */
        id?: pulumi.Input<number>;
        name?: pulumi.Input<string>;
    }

    export interface TrafficForwardingGRETunnelPrimaryDestVip {
        datacenter?: pulumi.Input<string>;
        /**
         * Unique identifer of the GRE virtual IP address (VIP)
         */
        id?: pulumi.Input<number>;
        privateServiceEdge?: pulumi.Input<boolean>;
        /**
         * GRE cluster virtual IP address (VIP)
         */
        virtualIp?: pulumi.Input<string>;
    }

    export interface TrafficForwardingGRETunnelSecondaryDestVip {
        datacenter?: pulumi.Input<string>;
        /**
         * Unique identifer of the GRE virtual IP address (VIP)
         */
        id?: pulumi.Input<number>;
        privateServiceEdge?: pulumi.Input<boolean>;
        /**
         * GRE cluster virtual IP address (VIP)
         */
        virtualIp?: pulumi.Input<string>;
    }

    export interface TrafficForwardingStaticIPLastModifiedBy {
        extensions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        id?: pulumi.Input<number>;
        name?: pulumi.Input<string>;
    }

    export interface TrafficForwardingStaticIPManagedBy {
        extensions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        id?: pulumi.Input<number>;
        name?: pulumi.Input<string>;
    }
}

export namespace URLCategory {
    export interface URLCategoriesScope {
        scopeEntities?: pulumi.Input<inputs.URLCategory.URLCategoriesScopeScopeEntities>;
        /**
         * Only applicable for the LOCATION_GROUP admin scope type, in which case this attribute gives the list of ID/name pairs of locations within the location group. The attribute name is subject to change.
         */
        scopeGroupMemberEntities?: pulumi.Input<inputs.URLCategory.URLCategoriesScopeScopeGroupMemberEntities>;
        /**
         * Type of the custom categories. `URL_CATEGORY`, `TLD_CATEGORY`, `ALL`
         */
        type?: pulumi.Input<string>;
    }

    export interface URLCategoriesScopeScopeEntities {
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface URLCategoriesScopeScopeGroupMemberEntities {
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface URLCategoriesUrlKeywordCounts {
        /**
         * Count of total keywords with retain parent category.
         */
        retainParentKeywordCount?: pulumi.Input<number>;
        /**
         * Count of URLs with retain parent category.
         */
        retainParentUrlCount?: pulumi.Input<number>;
        /**
         * Total keyword count for the category.
         */
        totalKeywordCount?: pulumi.Input<number>;
        /**
         * Custom URL count for the category.
         */
        totalUrlCount?: pulumi.Input<number>;
    }
}

export namespace URLFiltering {
    export interface URLFilteringRulesDepartments {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface URLFilteringRulesDeviceGroups {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface URLFilteringRulesDevices {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface URLFilteringRulesGroups {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface URLFilteringRulesLabels {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface URLFilteringRulesLastModifiedBy {
        extensions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        /**
         * Identifier that uniquely identifies an entity
         */
        id?: pulumi.Input<number>;
        /**
         * Name of the Firewall Filtering policy rule
         */
        name?: pulumi.Input<string>;
    }

    export interface URLFilteringRulesLocationGroups {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface URLFilteringRulesLocations {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface URLFilteringRulesOverrideGroups {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface URLFilteringRulesOverrideUsers {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface URLFilteringRulesTimeWindows {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }

    export interface URLFilteringRulesUsers {
        /**
         * Identifier that uniquely identifies an entity
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }
}

export namespace Users {
    export interface UserManagementDepartment {
        comments?: pulumi.Input<string>;
        deleted?: pulumi.Input<boolean>;
        /**
         * Department ID
         */
        id?: pulumi.Input<number>;
        idpId?: pulumi.Input<number>;
        /**
         * User name. This appears when choosing users for policies.
         */
        name?: pulumi.Input<string>;
    }

    export interface UserManagementGroups {
        /**
         * Department ID
         */
        ids: pulumi.Input<pulumi.Input<number>[]>;
    }
}
