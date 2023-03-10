// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetTrafficForwardingGREInternalIPRangeArgs, GetTrafficForwardingGREInternalIPRangeResult, GetTrafficForwardingGREInternalIPRangeOutputArgs } from "./getTrafficForwardingGREInternalIPRange";
export const getTrafficForwardingGREInternalIPRange: typeof import("./getTrafficForwardingGREInternalIPRange").getTrafficForwardingGREInternalIPRange = null as any;
export const getTrafficForwardingGREInternalIPRangeOutput: typeof import("./getTrafficForwardingGREInternalIPRange").getTrafficForwardingGREInternalIPRangeOutput = null as any;
utilities.lazyLoad(exports, ["getTrafficForwardingGREInternalIPRange","getTrafficForwardingGREInternalIPRangeOutput"], () => require("./getTrafficForwardingGREInternalIPRange"));

export { GetTrafficForwardingGRETunnelArgs, GetTrafficForwardingGRETunnelResult, GetTrafficForwardingGRETunnelOutputArgs } from "./getTrafficForwardingGRETunnel";
export const getTrafficForwardingGRETunnel: typeof import("./getTrafficForwardingGRETunnel").getTrafficForwardingGRETunnel = null as any;
export const getTrafficForwardingGRETunnelOutput: typeof import("./getTrafficForwardingGRETunnel").getTrafficForwardingGRETunnelOutput = null as any;
utilities.lazyLoad(exports, ["getTrafficForwardingGRETunnel","getTrafficForwardingGRETunnelOutput"], () => require("./getTrafficForwardingGRETunnel"));

export { GetTrafficForwardingGRETunnelInfoArgs, GetTrafficForwardingGRETunnelInfoResult, GetTrafficForwardingGRETunnelInfoOutputArgs } from "./getTrafficForwardingGRETunnelInfo";
export const getTrafficForwardingGRETunnelInfo: typeof import("./getTrafficForwardingGRETunnelInfo").getTrafficForwardingGRETunnelInfo = null as any;
export const getTrafficForwardingGRETunnelInfoOutput: typeof import("./getTrafficForwardingGRETunnelInfo").getTrafficForwardingGRETunnelInfoOutput = null as any;
utilities.lazyLoad(exports, ["getTrafficForwardingGRETunnelInfo","getTrafficForwardingGRETunnelInfoOutput"], () => require("./getTrafficForwardingGRETunnelInfo"));

export { GetTrafficForwardingNodeVIPsArgs, GetTrafficForwardingNodeVIPsResult, GetTrafficForwardingNodeVIPsOutputArgs } from "./getTrafficForwardingNodeVIPs";
export const getTrafficForwardingNodeVIPs: typeof import("./getTrafficForwardingNodeVIPs").getTrafficForwardingNodeVIPs = null as any;
export const getTrafficForwardingNodeVIPsOutput: typeof import("./getTrafficForwardingNodeVIPs").getTrafficForwardingNodeVIPsOutput = null as any;
utilities.lazyLoad(exports, ["getTrafficForwardingNodeVIPs","getTrafficForwardingNodeVIPsOutput"], () => require("./getTrafficForwardingNodeVIPs"));

export { GetTrafficForwardingStaticIPArgs, GetTrafficForwardingStaticIPResult, GetTrafficForwardingStaticIPOutputArgs } from "./getTrafficForwardingStaticIP";
export const getTrafficForwardingStaticIP: typeof import("./getTrafficForwardingStaticIP").getTrafficForwardingStaticIP = null as any;
export const getTrafficForwardingStaticIPOutput: typeof import("./getTrafficForwardingStaticIP").getTrafficForwardingStaticIPOutput = null as any;
utilities.lazyLoad(exports, ["getTrafficForwardingStaticIP","getTrafficForwardingStaticIPOutput"], () => require("./getTrafficForwardingStaticIP"));

export { GetTrafficForwardingVIPRecommendedListArgs, GetTrafficForwardingVIPRecommendedListResult, GetTrafficForwardingVIPRecommendedListOutputArgs } from "./getTrafficForwardingVIPRecommendedList";
export const getTrafficForwardingVIPRecommendedList: typeof import("./getTrafficForwardingVIPRecommendedList").getTrafficForwardingVIPRecommendedList = null as any;
export const getTrafficForwardingVIPRecommendedListOutput: typeof import("./getTrafficForwardingVIPRecommendedList").getTrafficForwardingVIPRecommendedListOutput = null as any;
utilities.lazyLoad(exports, ["getTrafficForwardingVIPRecommendedList","getTrafficForwardingVIPRecommendedListOutput"], () => require("./getTrafficForwardingVIPRecommendedList"));

export { GetTrafficForwardingVPNCredentialsArgs, GetTrafficForwardingVPNCredentialsResult, GetTrafficForwardingVPNCredentialsOutputArgs } from "./getTrafficForwardingVPNCredentials";
export const getTrafficForwardingVPNCredentials: typeof import("./getTrafficForwardingVPNCredentials").getTrafficForwardingVPNCredentials = null as any;
export const getTrafficForwardingVPNCredentialsOutput: typeof import("./getTrafficForwardingVPNCredentials").getTrafficForwardingVPNCredentialsOutput = null as any;
utilities.lazyLoad(exports, ["getTrafficForwardingVPNCredentials","getTrafficForwardingVPNCredentialsOutput"], () => require("./getTrafficForwardingVPNCredentials"));

export { TrafficForwardingGRETunnelArgs, TrafficForwardingGRETunnelState } from "./trafficForwardingGRETunnel";
export type TrafficForwardingGRETunnel = import("./trafficForwardingGRETunnel").TrafficForwardingGRETunnel;
export const TrafficForwardingGRETunnel: typeof import("./trafficForwardingGRETunnel").TrafficForwardingGRETunnel = null as any;
utilities.lazyLoad(exports, ["TrafficForwardingGRETunnel"], () => require("./trafficForwardingGRETunnel"));

export { TrafficForwardingStaticIPArgs, TrafficForwardingStaticIPState } from "./trafficForwardingStaticIP";
export type TrafficForwardingStaticIP = import("./trafficForwardingStaticIP").TrafficForwardingStaticIP;
export const TrafficForwardingStaticIP: typeof import("./trafficForwardingStaticIP").TrafficForwardingStaticIP = null as any;
utilities.lazyLoad(exports, ["TrafficForwardingStaticIP"], () => require("./trafficForwardingStaticIP"));

export { TrafficForwardingVPNCredentialsArgs, TrafficForwardingVPNCredentialsState } from "./trafficForwardingVPNCredentials";
export type TrafficForwardingVPNCredentials = import("./trafficForwardingVPNCredentials").TrafficForwardingVPNCredentials;
export const TrafficForwardingVPNCredentials: typeof import("./trafficForwardingVPNCredentials").TrafficForwardingVPNCredentials = null as any;
utilities.lazyLoad(exports, ["TrafficForwardingVPNCredentials"], () => require("./trafficForwardingVPNCredentials"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "zia:TrafficForwarding/trafficForwardingGRETunnel:TrafficForwardingGRETunnel":
                return new TrafficForwardingGRETunnel(name, <any>undefined, { urn })
            case "zia:TrafficForwarding/trafficForwardingStaticIP:TrafficForwardingStaticIP":
                return new TrafficForwardingStaticIP(name, <any>undefined, { urn })
            case "zia:TrafficForwarding/trafficForwardingVPNCredentials:TrafficForwardingVPNCredentials":
                return new TrafficForwardingVPNCredentials(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("zia", "TrafficForwarding/trafficForwardingGRETunnel", _module)
pulumi.runtime.registerResourceModule("zia", "TrafficForwarding/trafficForwardingStaticIP", _module)
pulumi.runtime.registerResourceModule("zia", "TrafficForwarding/trafficForwardingVPNCredentials", _module)
