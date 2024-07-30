// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zia
{
    public static class GetDeviceGroups
    {
        /// <summary>
        /// Use the **zia_device_groups** data source to get information about a device group in the Zscaler Internet Access cloud or via the API. This data source can then be associated with resources such as: URL Filtering Rules
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zia = Pulumi.Zia;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ios = Zia.GetDeviceGroups.Invoke(new()
        ///     {
        ///         Name = "IOS",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zia = Pulumi.Zia;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var android = Zia.GetDeviceGroups.Invoke(new()
        ///     {
        ///         Name = "Android",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetDeviceGroupsResult> InvokeAsync(GetDeviceGroupsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDeviceGroupsResult>("zia:index/getDeviceGroups:getDeviceGroups", args ?? new GetDeviceGroupsArgs(), options.WithDefaults());

        /// <summary>
        /// Use the **zia_device_groups** data source to get information about a device group in the Zscaler Internet Access cloud or via the API. This data source can then be associated with resources such as: URL Filtering Rules
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zia = Pulumi.Zia;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ios = Zia.GetDeviceGroups.Invoke(new()
        ///     {
        ///         Name = "IOS",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zia = Pulumi.Zia;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var android = Zia.GetDeviceGroups.Invoke(new()
        ///     {
        ///         Name = "Android",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDeviceGroupsResult> Invoke(GetDeviceGroupsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDeviceGroupsResult>("zia:index/getDeviceGroups:getDeviceGroups", args ?? new GetDeviceGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDeviceGroupsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the device group to be exported.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetDeviceGroupsArgs()
        {
        }
        public static new GetDeviceGroupsArgs Empty => new GetDeviceGroupsArgs();
    }

    public sealed class GetDeviceGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the device group to be exported.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetDeviceGroupsInvokeArgs()
        {
        }
        public static new GetDeviceGroupsInvokeArgs Empty => new GetDeviceGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetDeviceGroupsResult
    {
        /// <summary>
        /// (String) The device group's description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (int) The number of devices within the group.
        /// </summary>
        public readonly int DeviceCount;
        /// <summary>
        /// (String) The names of devices that belong to the device group. The device names are comma-separated.
        /// </summary>
        public readonly string DeviceNames;
        /// <summary>
        /// (String) The device group type. i.e ``ZCC_OS``, ``NON_ZCC``, ``CBI``
        /// </summary>
        public readonly string GroupType;
        /// <summary>
        /// (String) The unique identifer for the device group.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// (String) The device group name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// (String) The operating system (OS).
        /// </summary>
        public readonly string OsType;
        /// <summary>
        /// (Boolean) Indicates whether this is a predefined device group. If this value is set to true, the group is predefined.
        /// </summary>
        public readonly bool Predefined;

        [OutputConstructor]
        private GetDeviceGroupsResult(
            string description,

            int deviceCount,

            string deviceNames,

            string groupType,

            int id,

            string? name,

            string osType,

            bool predefined)
        {
            Description = description;
            DeviceCount = deviceCount;
            DeviceNames = deviceNames;
            GroupType = groupType;
            Id = id;
            Name = name;
            OsType = osType;
            Predefined = predefined;
        }
    }
}
