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
    public static class GetDLPIDMProfiles
    {
        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zia/data-loss-prevention#/idmprofile-get)
        /// * [API documentation](https://help.zscaler.com/zia/about-indexed-document-match)
        /// 
        /// Use the **zia_dlp_idm_profile** data source to get information about a ZIA DLP IDM Profile in the Zscaler Internet Access cloud or via the API.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl
        /// # Retrieve a DLP IDM Profile by name
        /// data "zia_dlp_idm_profile" "example"{
        ///     name = "Example"
        /// }
        /// ```
        /// 
        /// ```hcl
        /// # Retrieve a DLP IDM Profile by ID
        /// data "zia_dlp_idm_profile" "example"{
        ///     name = "Example"
        /// }
        /// ```
        /// </summary>
        public static Task<GetDLPIDMProfilesResult> InvokeAsync(GetDLPIDMProfilesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDLPIDMProfilesResult>("zia:index/getDLPIDMProfiles:getDLPIDMProfiles", args ?? new GetDLPIDMProfilesArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zia/data-loss-prevention#/idmprofile-get)
        /// * [API documentation](https://help.zscaler.com/zia/about-indexed-document-match)
        /// 
        /// Use the **zia_dlp_idm_profile** data source to get information about a ZIA DLP IDM Profile in the Zscaler Internet Access cloud or via the API.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl
        /// # Retrieve a DLP IDM Profile by name
        /// data "zia_dlp_idm_profile" "example"{
        ///     name = "Example"
        /// }
        /// ```
        /// 
        /// ```hcl
        /// # Retrieve a DLP IDM Profile by ID
        /// data "zia_dlp_idm_profile" "example"{
        ///     name = "Example"
        /// }
        /// ```
        /// </summary>
        public static Output<GetDLPIDMProfilesResult> Invoke(GetDLPIDMProfilesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDLPIDMProfilesResult>("zia:index/getDLPIDMProfiles:getDLPIDMProfiles", args ?? new GetDLPIDMProfilesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://help.zscaler.com/zia/data-loss-prevention#/idmprofile-get)
        /// * [API documentation](https://help.zscaler.com/zia/about-indexed-document-match)
        /// 
        /// Use the **zia_dlp_idm_profile** data source to get information about a ZIA DLP IDM Profile in the Zscaler Internet Access cloud or via the API.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl
        /// # Retrieve a DLP IDM Profile by name
        /// data "zia_dlp_idm_profile" "example"{
        ///     name = "Example"
        /// }
        /// ```
        /// 
        /// ```hcl
        /// # Retrieve a DLP IDM Profile by ID
        /// data "zia_dlp_idm_profile" "example"{
        ///     name = "Example"
        /// }
        /// ```
        /// </summary>
        public static Output<GetDLPIDMProfilesResult> Invoke(GetDLPIDMProfilesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDLPIDMProfilesResult>("zia:index/getDLPIDMProfiles:getDLPIDMProfiles", args ?? new GetDLPIDMProfilesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDLPIDMProfilesArgs : global::Pulumi.InvokeArgs
    {
        [Input("profileName")]
        public string? ProfileName { get; set; }

        public GetDLPIDMProfilesArgs()
        {
        }
        public static new GetDLPIDMProfilesArgs Empty => new GetDLPIDMProfilesArgs();
    }

    public sealed class GetDLPIDMProfilesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("profileName")]
        public Input<string>? ProfileName { get; set; }

        public GetDLPIDMProfilesInvokeArgs()
        {
        }
        public static new GetDLPIDMProfilesInvokeArgs Empty => new GetDLPIDMProfilesInvokeArgs();
    }


    [OutputType]
    public sealed class GetDLPIDMProfilesResult
    {
        public readonly string Host;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetDLPIDMProfilesIdmClientResult> IdmClients;
        public readonly ImmutableArray<Outputs.GetDLPIDMProfilesLastModifiedByResult> LastModifiedBies;
        public readonly int LastModifiedTime;
        public readonly int NumDocuments;
        public readonly int Port;
        public readonly string ProfileDesc;
        public readonly string ProfileDirPath;
        public readonly int ProfileId;
        public readonly string? ProfileName;
        public readonly string ProfileType;
        public readonly int ScheduleDay;
        public readonly ImmutableArray<string> ScheduleDayOfMonths;
        public readonly ImmutableArray<string> ScheduleDayOfWeeks;
        public readonly bool ScheduleDisabled;
        public readonly int ScheduleTime;
        public readonly string ScheduleType;
        public readonly string UploadStatus;
        public readonly string Username;
        public readonly int Version;
        public readonly int VolumeOfDocuments;

        [OutputConstructor]
        private GetDLPIDMProfilesResult(
            string host,

            string id,

            ImmutableArray<Outputs.GetDLPIDMProfilesIdmClientResult> idmClients,

            ImmutableArray<Outputs.GetDLPIDMProfilesLastModifiedByResult> lastModifiedBies,

            int lastModifiedTime,

            int numDocuments,

            int port,

            string profileDesc,

            string profileDirPath,

            int profileId,

            string? profileName,

            string profileType,

            int scheduleDay,

            ImmutableArray<string> scheduleDayOfMonths,

            ImmutableArray<string> scheduleDayOfWeeks,

            bool scheduleDisabled,

            int scheduleTime,

            string scheduleType,

            string uploadStatus,

            string username,

            int version,

            int volumeOfDocuments)
        {
            Host = host;
            Id = id;
            IdmClients = idmClients;
            LastModifiedBies = lastModifiedBies;
            LastModifiedTime = lastModifiedTime;
            NumDocuments = numDocuments;
            Port = port;
            ProfileDesc = profileDesc;
            ProfileDirPath = profileDirPath;
            ProfileId = profileId;
            ProfileName = profileName;
            ProfileType = profileType;
            ScheduleDay = scheduleDay;
            ScheduleDayOfMonths = scheduleDayOfMonths;
            ScheduleDayOfWeeks = scheduleDayOfWeeks;
            ScheduleDisabled = scheduleDisabled;
            ScheduleTime = scheduleTime;
            ScheduleType = scheduleType;
            UploadStatus = uploadStatus;
            Username = username;
            Version = version;
            VolumeOfDocuments = volumeOfDocuments;
        }
    }
}
