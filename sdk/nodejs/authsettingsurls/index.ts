// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AuthSettingsURLsArgs, AuthSettingsURLsState } from "./authSettingsURLs";
export type AuthSettingsURLs = import("./authSettingsURLs").AuthSettingsURLs;
export const AuthSettingsURLs: typeof import("./authSettingsURLs").AuthSettingsURLs = null as any;
utilities.lazyLoad(exports, ["AuthSettingsURLs"], () => require("./authSettingsURLs"));

export { GetAuthSettingsURLsResult } from "./getAuthSettingsURLs";
export const getAuthSettingsURLs: typeof import("./getAuthSettingsURLs").getAuthSettingsURLs = null as any;
utilities.lazyLoad(exports, ["getAuthSettingsURLs"], () => require("./getAuthSettingsURLs"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "zia:AuthSettingsUrls/authSettingsURLs:AuthSettingsURLs":
                return new AuthSettingsURLs(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("zia", "AuthSettingsUrls/authSettingsURLs", _module)
