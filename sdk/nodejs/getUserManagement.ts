// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use the **zia_user_management** data source to get information about a user account that may have been created in the Zscaler Internet Access portal or via API. This data source can then be associated with a ZIA cloud firewall filtering rule, and URL filtering rules.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@pulumi/zia";
 *
 * const adamAshcroft = zia.getUserManagement({
 *     name: "Adam Ashcroft",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getUserManagement(args?: GetUserManagementArgs, opts?: pulumi.InvokeOptions): Promise<GetUserManagementResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zia:index/getUserManagement:getUserManagement", {
        "authMethods": args.authMethods,
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getUserManagement.
 */
export interface GetUserManagementArgs {
    /**
     * (String) Type of authentication method to be enabled. Supported values are: ``BASIC`` and ``DIGEST``
     */
    authMethods?: string[];
    /**
     * The ID of the time window resource.
     */
    id?: number;
    /**
     * User name. This appears when choosing users for policies.
     */
    name?: string;
}

/**
 * A collection of values returned by getUserManagement.
 */
export interface GetUserManagementResult {
    /**
     * (String) True if this user is an Admin user. readOnly: `true` default: `false`
     */
    readonly adminUser: boolean;
    /**
     * (String) Type of authentication method to be enabled. Supported values are: ``BASIC`` and ``DIGEST``
     */
    readonly authMethods?: string[];
    /**
     * (String) Additional information about the group
     */
    readonly comments: string;
    /**
     * (String) Department a user belongs to
     */
    readonly departments: outputs.GetUserManagementDepartment[];
    /**
     * (Required) User email consists of a user name and domain name. It does not have to be a valid email address, but it must be unique and its domain must belong to the organization
     */
    readonly email: string;
    /**
     * (String) List of Groups a user belongs to. Groups are used in policies.
     */
    readonly groups: outputs.GetUserManagementGroup[];
    /**
     * (Number) Unique identfier for the group
     */
    readonly id?: number;
    readonly isAuditor: string;
    /**
     * (String) Group name
     */
    readonly name?: string;
    /**
     * (String) Temporary Authentication Email. If you enabled one-time tokens or links, enter the email address to which the Zscaler service sends the tokens or links. If this is empty, the service will send the email to the User email.
     */
    readonly tempAuthEmail: string;
    /**
     * (String) User type. Provided only if this user is not an end user. The supported types are:
     */
    readonly type: string;
}
/**
 * Use the **zia_user_management** data source to get information about a user account that may have been created in the Zscaler Internet Access portal or via API. This data source can then be associated with a ZIA cloud firewall filtering rule, and URL filtering rules.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@pulumi/zia";
 *
 * const adamAshcroft = zia.getUserManagement({
 *     name: "Adam Ashcroft",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getUserManagementOutput(args?: GetUserManagementOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserManagementResult> {
    return pulumi.output(args).apply((a: any) => getUserManagement(a, opts))
}

/**
 * A collection of arguments for invoking getUserManagement.
 */
export interface GetUserManagementOutputArgs {
    /**
     * (String) Type of authentication method to be enabled. Supported values are: ``BASIC`` and ``DIGEST``
     */
    authMethods?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the time window resource.
     */
    id?: pulumi.Input<number>;
    /**
     * User name. This appears when choosing users for policies.
     */
    name?: pulumi.Input<string>;
}
