// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The **zia_admin_users** resource allows the creation and management of ZIA admin user account created in the Zscaler Internet Access cloud or via the API.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zia from "@pulumi/zia";
 * import * as zia from "@zscaler/pulumi-zia";
 *
 * const superAdmin = zia.AdminRoles.getAdminRoles({
 *     name: "Super Admin",
 * });
 * const engineering = zia.Departments.getDepartmentManagement({
 *     name: "Engineering",
 * });
 * const johnSmith = new zia.adminusers.AdminUsers("johnSmith", {
 *     loginName: "john.smith@acme.com",
 *     userName: "John Smith",
 *     email: "john.smith@acme.com",
 *     isPasswordLoginAllowed: true,
 *     password: `AeQ9E5w8B$`,
 *     isSecurityReportCommEnabled: true,
 *     isServiceUpdateCommEnabled: true,
 *     isProductUpdateCommEnabled: true,
 *     comments: "Administrator User",
 *     roles: [{
 *         id: superAdmin.then(superAdmin => superAdmin.id),
 *     }],
 *     adminScopes: [{
 *         type: "DEPARTMENT",
 *         scopeEntities: {
 *             ids: [engineering.then(engineering => engineering.id)],
 *         },
 *     }],
 * });
 * ```
 */
export class AdminUsers extends pulumi.CustomResource {
    /**
     * Get an existing AdminUsers resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AdminUsersState, opts?: pulumi.CustomResourceOptions): AdminUsers {
        return new AdminUsers(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zia:AdminUsers/adminUsers:AdminUsers';

    /**
     * Returns true if the given object is an instance of AdminUsers.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AdminUsers {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AdminUsers.__pulumiType;
    }

    public /*out*/ readonly adminId!: pulumi.Output<number>;
    /**
     * The admin's scope. A scope is required for admins, but not applicable to auditors. This attribute is subject to change.
     */
    public readonly adminScopes!: pulumi.Output<outputs.AdminUsers.AdminUsersAdminScope[]>;
    /**
     * Additional information about the admin or auditor.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether or not the admin account is disabled.
     */
    public readonly disabled!: pulumi.Output<boolean>;
    /**
     * Admin or auditor's email address.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * Indicates whether the user is an auditor. This attribute is subject to change.
     */
    public readonly isAuditor!: pulumi.Output<boolean>;
    /**
     * Indicates whether or not Executive Insights App access is enabled for the admin.
     */
    public readonly isExecMobileAppEnabled!: pulumi.Output<boolean>;
    /**
     * Indicates whether or not the admin can be edited or deleted.
     */
    public readonly isNonEditable!: pulumi.Output<boolean>;
    /**
     * Indicates whether or not an admin's password has expired.
     */
    public readonly isPasswordExpired!: pulumi.Output<boolean>;
    /**
     * The default is true when SAML Authentication is disabled. When SAML Authentication is enabled, this can be set to false in order to force the admin to login via SSO only.
     */
    public readonly isPasswordLoginAllowed!: pulumi.Output<boolean>;
    /**
     * Communication setting for Product Update.
     */
    public readonly isProductUpdateCommEnabled!: pulumi.Output<boolean>;
    /**
     * Communication for Security Report is enabled.
     */
    public readonly isSecurityReportCommEnabled!: pulumi.Output<boolean>;
    /**
     * Communication setting for Service Update.
     */
    public readonly isServiceUpdateCommEnabled!: pulumi.Output<boolean>;
    /**
     * The email address of the admin user to be exported.
     */
    public readonly loginName!: pulumi.Output<string>;
    /**
     * The username of the admin user to be exported.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Role of the admin. This is not required for an auditor.
     */
    public readonly roles!: pulumi.Output<outputs.AdminUsers.AdminUsersRole[]>;
    /**
     * The username of the admin user to be exported.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a AdminUsers resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AdminUsersArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AdminUsersArgs | AdminUsersState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AdminUsersState | undefined;
            resourceInputs["adminId"] = state ? state.adminId : undefined;
            resourceInputs["adminScopes"] = state ? state.adminScopes : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["isAuditor"] = state ? state.isAuditor : undefined;
            resourceInputs["isExecMobileAppEnabled"] = state ? state.isExecMobileAppEnabled : undefined;
            resourceInputs["isNonEditable"] = state ? state.isNonEditable : undefined;
            resourceInputs["isPasswordExpired"] = state ? state.isPasswordExpired : undefined;
            resourceInputs["isPasswordLoginAllowed"] = state ? state.isPasswordLoginAllowed : undefined;
            resourceInputs["isProductUpdateCommEnabled"] = state ? state.isProductUpdateCommEnabled : undefined;
            resourceInputs["isSecurityReportCommEnabled"] = state ? state.isSecurityReportCommEnabled : undefined;
            resourceInputs["isServiceUpdateCommEnabled"] = state ? state.isServiceUpdateCommEnabled : undefined;
            resourceInputs["loginName"] = state ? state.loginName : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["roles"] = state ? state.roles : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as AdminUsersArgs | undefined;
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.loginName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'loginName'");
            }
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["adminScopes"] = args ? args.adminScopes : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["isAuditor"] = args ? args.isAuditor : undefined;
            resourceInputs["isExecMobileAppEnabled"] = args ? args.isExecMobileAppEnabled : undefined;
            resourceInputs["isNonEditable"] = args ? args.isNonEditable : undefined;
            resourceInputs["isPasswordExpired"] = args ? args.isPasswordExpired : undefined;
            resourceInputs["isPasswordLoginAllowed"] = args ? args.isPasswordLoginAllowed : undefined;
            resourceInputs["isProductUpdateCommEnabled"] = args ? args.isProductUpdateCommEnabled : undefined;
            resourceInputs["isSecurityReportCommEnabled"] = args ? args.isSecurityReportCommEnabled : undefined;
            resourceInputs["isServiceUpdateCommEnabled"] = args ? args.isServiceUpdateCommEnabled : undefined;
            resourceInputs["loginName"] = args ? args.loginName : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["roles"] = args ? args.roles : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["adminId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(AdminUsers.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AdminUsers resources.
 */
export interface AdminUsersState {
    adminId?: pulumi.Input<number>;
    /**
     * The admin's scope. A scope is required for admins, but not applicable to auditors. This attribute is subject to change.
     */
    adminScopes?: pulumi.Input<pulumi.Input<inputs.AdminUsers.AdminUsersAdminScope>[]>;
    /**
     * Additional information about the admin or auditor.
     */
    comments?: pulumi.Input<string>;
    /**
     * Indicates whether or not the admin account is disabled.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * Admin or auditor's email address.
     */
    email?: pulumi.Input<string>;
    /**
     * Indicates whether the user is an auditor. This attribute is subject to change.
     */
    isAuditor?: pulumi.Input<boolean>;
    /**
     * Indicates whether or not Executive Insights App access is enabled for the admin.
     */
    isExecMobileAppEnabled?: pulumi.Input<boolean>;
    /**
     * Indicates whether or not the admin can be edited or deleted.
     */
    isNonEditable?: pulumi.Input<boolean>;
    /**
     * Indicates whether or not an admin's password has expired.
     */
    isPasswordExpired?: pulumi.Input<boolean>;
    /**
     * The default is true when SAML Authentication is disabled. When SAML Authentication is enabled, this can be set to false in order to force the admin to login via SSO only.
     */
    isPasswordLoginAllowed?: pulumi.Input<boolean>;
    /**
     * Communication setting for Product Update.
     */
    isProductUpdateCommEnabled?: pulumi.Input<boolean>;
    /**
     * Communication for Security Report is enabled.
     */
    isSecurityReportCommEnabled?: pulumi.Input<boolean>;
    /**
     * Communication setting for Service Update.
     */
    isServiceUpdateCommEnabled?: pulumi.Input<boolean>;
    /**
     * The email address of the admin user to be exported.
     */
    loginName?: pulumi.Input<string>;
    /**
     * The username of the admin user to be exported.
     */
    password?: pulumi.Input<string>;
    /**
     * Role of the admin. This is not required for an auditor.
     */
    roles?: pulumi.Input<pulumi.Input<inputs.AdminUsers.AdminUsersRole>[]>;
    /**
     * The username of the admin user to be exported.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AdminUsers resource.
 */
export interface AdminUsersArgs {
    /**
     * The admin's scope. A scope is required for admins, but not applicable to auditors. This attribute is subject to change.
     */
    adminScopes?: pulumi.Input<pulumi.Input<inputs.AdminUsers.AdminUsersAdminScope>[]>;
    /**
     * Additional information about the admin or auditor.
     */
    comments?: pulumi.Input<string>;
    /**
     * Indicates whether or not the admin account is disabled.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * Admin or auditor's email address.
     */
    email: pulumi.Input<string>;
    /**
     * Indicates whether the user is an auditor. This attribute is subject to change.
     */
    isAuditor?: pulumi.Input<boolean>;
    /**
     * Indicates whether or not Executive Insights App access is enabled for the admin.
     */
    isExecMobileAppEnabled?: pulumi.Input<boolean>;
    /**
     * Indicates whether or not the admin can be edited or deleted.
     */
    isNonEditable?: pulumi.Input<boolean>;
    /**
     * Indicates whether or not an admin's password has expired.
     */
    isPasswordExpired?: pulumi.Input<boolean>;
    /**
     * The default is true when SAML Authentication is disabled. When SAML Authentication is enabled, this can be set to false in order to force the admin to login via SSO only.
     */
    isPasswordLoginAllowed?: pulumi.Input<boolean>;
    /**
     * Communication setting for Product Update.
     */
    isProductUpdateCommEnabled?: pulumi.Input<boolean>;
    /**
     * Communication for Security Report is enabled.
     */
    isSecurityReportCommEnabled?: pulumi.Input<boolean>;
    /**
     * Communication setting for Service Update.
     */
    isServiceUpdateCommEnabled?: pulumi.Input<boolean>;
    /**
     * The email address of the admin user to be exported.
     */
    loginName: pulumi.Input<string>;
    /**
     * The username of the admin user to be exported.
     */
    password?: pulumi.Input<string>;
    /**
     * Role of the admin. This is not required for an auditor.
     */
    roles?: pulumi.Input<pulumi.Input<inputs.AdminUsers.AdminUsersRole>[]>;
    /**
     * The username of the admin user to be exported.
     */
    username: pulumi.Input<string>;
}
