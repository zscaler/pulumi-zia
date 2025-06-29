// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://help.zscaler.com/zia/configuring-browser-control-policy)
 * * [API documentation](https://help.zscaler.com/zia/browser-control-policy#/browserControlSettings-get)
 *
 * Use the **zia_browser_control_policy** data source to retrieves information about the security exceptions configured for the Malware Protection policy. To learn more see [Configuring the Browser Control Policy](https://help.zscaler.com/zia/configuring-browser-control-policy)
 *
 * ## Example Usage
 */
export function getBrowserControlSettings(opts?: pulumi.InvokeOptions): Promise<GetBrowserControlSettingsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zia:index/getBrowserControlSettings:getBrowserControlSettings", {
    }, opts);
}

/**
 * A collection of values returned by getBrowserControlSettings.
 */
export interface GetBrowserControlSettingsResult {
    /**
     * (Boolean) A Boolean value that specifies whether or not to allow all the browsers and their respective versions access to the internet
     */
    readonly allowAllBrowsers: boolean;
    /**
     * (List) Versions of Google Chrome browser that need to be blocked. If not set, all Google Chrome versions are allowed. See all [Supported values](https://help.zscaler.com/zia/browser-control-policy#/browserControlSettings-get)
     */
    readonly blockedChromeVersions: string[];
    /**
     * (List) Versions of Mozilla Firefox browser that need to be blocked. If not set, all Mozilla Firefox versions are allowed. See all [Supported values](https://help.zscaler.com/zia/browser-control-policy#/browserControlSettings-get)
     */
    readonly blockedFirefoxVersions: string[];
    /**
     * (List) Versions of Microsoft browser that need to be blocked. If not set, all Microsoft browser versions are allowed. See all [Supported values](https://help.zscaler.com/zia/browser-control-policy#/browserControlSettings-get)
     */
    readonly blockedInternetExplorerVersions: string[];
    /**
     * (List) Versions of Opera browser that need to be blocked. If not set, all Opera versions are allowed. See all [Supported values](https://help.zscaler.com/zia/browser-control-policy#/browserControlSettings-get)
     */
    readonly blockedOperaVersions: string[];
    /**
     * (List) Versions of Apple Safari browser that need to be blocked. If not set, all Apple Safari versions are allowed. See all [Supported values](https://help.zscaler.com/zia/browser-control-policy#/browserControlSettings-get)
     */
    readonly blockedSafariVersions: string[];
    /**
     * (Boolean) If set to true, all the browsers are bypassed for warnings
     */
    readonly bypassAllBrowsers: boolean;
    /**
     * (List) List of applications that need to be bypassed for warnings. This attribute has effect only if the 'enableWarnings' attribute is set to true. If not set, all vulnerable applications are warned. Supported Values:
     * * `ANY`
     * * `NONE`
     * * `OUTLOOKEXP`
     * * `MSOFFICE`
     */
    readonly bypassApplications: string[];
    /**
     * (List) List of plugins that need to be bypassed for warnings. This attribute has effect only if the 'enableWarnings' attribute is set to true. If not set, all vulnerable plugins are warned.Supported Values:
     * * `ANY`
     * * `NONE`
     * * `ACROBAT`
     * * `FLASH`
     * * `SHOCKWAVE`
     * * `QUICKTIME`
     * * `DIVX`
     * * `GOOGLEGEARS`
     * * `DOTNET`
     * * `SILVERLIGHT`
     * * `REALPLAYER`
     * * `JAVA`
     * * `TOTEM`
     * * `WMP`
     */
    readonly bypassPlugins: string[];
    /**
     * (Boolean) A Boolean value that specifies if Smart Browser Isolation is enabled
     */
    readonly enableSmartBrowserIsolation: boolean;
    /**
     * (Boolean) A Boolean value that specifies if the warnings are enabled
     */
    readonly enableWarnings: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * (String) Specifies how frequently the service checks browsers and relevant applications to warn users regarding outdated or vulnerable browsers, plugins, and applications. If not set, the warnings are disabled. Supported Values:
     * * `DAILY`
     * * `WEEKLY`
     * * `MONTHLY`,
     * * `EVERY_2_HOURS`
     * * `EVERY_4_HOURS`
     * * `EVERY_6_HOURS`
     * * `EVERY_8_HOURS`
     * * `EVERY_12_HOURS`
     */
    readonly pluginCheckFrequency: string;
    readonly smartIsolationProfileId: number;
    /**
     * (Block, Max: 1) The isolation profile ID used for DLP email alerts sent to the auditor.
     */
    readonly smartIsolationProfiles: outputs.GetBrowserControlSettingsSmartIsolationProfile[];
}
/**
 * * [Official documentation](https://help.zscaler.com/zia/configuring-browser-control-policy)
 * * [API documentation](https://help.zscaler.com/zia/browser-control-policy#/browserControlSettings-get)
 *
 * Use the **zia_browser_control_policy** data source to retrieves information about the security exceptions configured for the Malware Protection policy. To learn more see [Configuring the Browser Control Policy](https://help.zscaler.com/zia/configuring-browser-control-policy)
 *
 * ## Example Usage
 */
export function getBrowserControlSettingsOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetBrowserControlSettingsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("zia:index/getBrowserControlSettings:getBrowserControlSettings", {
    }, opts);
}
