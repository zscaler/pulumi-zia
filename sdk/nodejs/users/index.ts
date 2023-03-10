// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetUserManagementArgs, GetUserManagementResult, GetUserManagementOutputArgs } from "./getUserManagement";
export const getUserManagement: typeof import("./getUserManagement").getUserManagement = null as any;
export const getUserManagementOutput: typeof import("./getUserManagement").getUserManagementOutput = null as any;
utilities.lazyLoad(exports, ["getUserManagement","getUserManagementOutput"], () => require("./getUserManagement"));

export { UserManagementArgs, UserManagementState } from "./userManagement";
export type UserManagement = import("./userManagement").UserManagement;
export const UserManagement: typeof import("./userManagement").UserManagement = null as any;
utilities.lazyLoad(exports, ["UserManagement"], () => require("./userManagement"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "zia:Users/userManagement:UserManagement":
                return new UserManagement(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("zia", "Users/userManagement", _module)
