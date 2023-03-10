// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));


// Export sub-modules:
import * as activation from "./activation";
import * as adminroles from "./adminroles";
import * as adminusers from "./adminusers";
import * as authsettingsurls from "./authsettingsurls";
import * as config from "./config";
import * as departments from "./departments";
import * as devicegroups from "./devicegroups";
import * as devices from "./devices";
import * as dlp from "./dlp";
import * as firewall from "./firewall";
import * as groups from "./groups";
import * as locationgroups from "./locationgroups";
import * as locationmanagement from "./locationmanagement";
import * as rulelabels from "./rulelabels";
import * as securitysettings from "./securitysettings";
import * as timewindow from "./timewindow";
import * as trafficforwarding from "./trafficforwarding";
import * as types from "./types";
import * as urlcategory from "./urlcategory";
import * as urlfiltering from "./urlfiltering";
import * as users from "./users";

export {
    activation,
    adminroles,
    adminusers,
    authsettingsurls,
    config,
    departments,
    devicegroups,
    devices,
    dlp,
    firewall,
    groups,
    locationgroups,
    locationmanagement,
    rulelabels,
    securitysettings,
    timewindow,
    trafficforwarding,
    types,
    urlcategory,
    urlfiltering,
    users,
};
pulumi.runtime.registerResourcePackage("zia", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:zia") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
