// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zia
{
    public static class GetURLFilteringCloudAppSettings
    {
        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zia/url-cloud-app-control-policy-settings#/advancedUrlFilterAndCloudAppSettings-get)
        /// * [API documentation](https://help.zscaler.com/zia/url-cloud-app-control-policy-settings#/advancedUrlFilterAndCloudAppSettings-get)
        /// 
        /// Use the **zia_url_filtering_and_cloud_app_settings** data source to get information about URL and Cloud App Control advanced policy settings.
        /// 
        /// ```hcl
        /// ```
        /// </summary>
        public static Task<GetURLFilteringCloudAppSettingsResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetURLFilteringCloudAppSettingsResult>("zia:index/getURLFilteringCloudAppSettings:getURLFilteringCloudAppSettings", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zia/url-cloud-app-control-policy-settings#/advancedUrlFilterAndCloudAppSettings-get)
        /// * [API documentation](https://help.zscaler.com/zia/url-cloud-app-control-policy-settings#/advancedUrlFilterAndCloudAppSettings-get)
        /// 
        /// Use the **zia_url_filtering_and_cloud_app_settings** data source to get information about URL and Cloud App Control advanced policy settings.
        /// 
        /// ```hcl
        /// ```
        /// </summary>
        public static Output<GetURLFilteringCloudAppSettingsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetURLFilteringCloudAppSettingsResult>("zia:index/getURLFilteringCloudAppSettings:getURLFilteringCloudAppSettings", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zia/url-cloud-app-control-policy-settings#/advancedUrlFilterAndCloudAppSettings-get)
        /// * [API documentation](https://help.zscaler.com/zia/url-cloud-app-control-policy-settings#/advancedUrlFilterAndCloudAppSettings-get)
        /// 
        /// Use the **zia_url_filtering_and_cloud_app_settings** data source to get information about URL and Cloud App Control advanced policy settings.
        /// 
        /// ```hcl
        /// ```
        /// </summary>
        public static Output<GetURLFilteringCloudAppSettingsResult> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetURLFilteringCloudAppSettingsResult>("zia:index/getURLFilteringCloudAppSettings:getURLFilteringCloudAppSettings", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetURLFilteringCloudAppSettingsResult
    {
        public readonly bool BlockSkype;
        public readonly bool ConsiderEmbeddedSites;
        public readonly bool EnableBlockOverrideForNonAuthUser;
        public readonly bool EnableChatgptPrompt;
        public readonly bool EnableCipaCompliance;
        public readonly bool EnableDynamicContentCat;
        public readonly bool EnableGeminiPrompt;
        public readonly bool EnableMetaPrompt;
        public readonly bool EnableMicrosoftCopilotPrompt;
        public readonly bool EnableMsftO365;
        public readonly bool EnableNewlyRegisteredDomains;
        public readonly bool EnableOffice365;
        public readonly bool EnablePerPlexityPrompt;
        public readonly bool EnablePoepPrompt;
        public readonly bool EnableUcaasLogmein;
        public readonly bool EnableUcaasRingCentral;
        public readonly bool EnableUcaasTalkdesk;
        public readonly bool EnableUcaasWebex;
        public readonly bool EnableUcaasZoom;
        public readonly bool EnforceSafeSearch;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetURLFilteringCloudAppSettingsResult(
            bool blockSkype,

            bool considerEmbeddedSites,

            bool enableBlockOverrideForNonAuthUser,

            bool enableChatgptPrompt,

            bool enableCipaCompliance,

            bool enableDynamicContentCat,

            bool enableGeminiPrompt,

            bool enableMetaPrompt,

            bool enableMicrosoftCopilotPrompt,

            bool enableMsftO365,

            bool enableNewlyRegisteredDomains,

            bool enableOffice365,

            bool enablePerPlexityPrompt,

            bool enablePoepPrompt,

            bool enableUcaasLogmein,

            bool enableUcaasRingCentral,

            bool enableUcaasTalkdesk,

            bool enableUcaasWebex,

            bool enableUcaasZoom,

            bool enforceSafeSearch,

            string id)
        {
            BlockSkype = blockSkype;
            ConsiderEmbeddedSites = considerEmbeddedSites;
            EnableBlockOverrideForNonAuthUser = enableBlockOverrideForNonAuthUser;
            EnableChatgptPrompt = enableChatgptPrompt;
            EnableCipaCompliance = enableCipaCompliance;
            EnableDynamicContentCat = enableDynamicContentCat;
            EnableGeminiPrompt = enableGeminiPrompt;
            EnableMetaPrompt = enableMetaPrompt;
            EnableMicrosoftCopilotPrompt = enableMicrosoftCopilotPrompt;
            EnableMsftO365 = enableMsftO365;
            EnableNewlyRegisteredDomains = enableNewlyRegisteredDomains;
            EnableOffice365 = enableOffice365;
            EnablePerPlexityPrompt = enablePerPlexityPrompt;
            EnablePoepPrompt = enablePoepPrompt;
            EnableUcaasLogmein = enableUcaasLogmein;
            EnableUcaasRingCentral = enableUcaasRingCentral;
            EnableUcaasTalkdesk = enableUcaasTalkdesk;
            EnableUcaasWebex = enableUcaasWebex;
            EnableUcaasZoom = enableUcaasZoom;
            EnforceSafeSearch = enforceSafeSearch;
            Id = id;
        }
    }
}
