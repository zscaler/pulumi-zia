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
    public static class GetLocationGroups
    {
        /// <summary>
        /// Use the **zia_location_groups** data source to get information about a location group option available in the Zscaler Internet Access.
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zia = Pulumi.Zia;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zia.GetLocationGroups.Invoke(new()
        ///     {
        ///         Name = "Corporate User Traffic Group",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zia = Pulumi.Zia;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zia.GetLocationGroups.Invoke(new()
        ///     {
        ///         Name = "Guest Wifi Group",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zia = Pulumi.Zia;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zia.GetLocationGroups.Invoke(new()
        ///     {
        ///         Name = "IoT Traffic Group",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zia = Pulumi.Zia;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zia.GetLocationGroups.Invoke(new()
        ///     {
        ///         Name = "Server Traffic Group",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zia = Pulumi.Zia;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zia.GetLocationGroups.Invoke(new()
        ///     {
        ///         Name = "Server Traffic Group",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetLocationGroupsResult> InvokeAsync(GetLocationGroupsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLocationGroupsResult>("zia:index/getLocationGroups:getLocationGroups", args ?? new GetLocationGroupsArgs(), options.WithDefaults());

        /// <summary>
        /// Use the **zia_location_groups** data source to get information about a location group option available in the Zscaler Internet Access.
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zia = Pulumi.Zia;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zia.GetLocationGroups.Invoke(new()
        ///     {
        ///         Name = "Corporate User Traffic Group",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zia = Pulumi.Zia;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zia.GetLocationGroups.Invoke(new()
        ///     {
        ///         Name = "Guest Wifi Group",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zia = Pulumi.Zia;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zia.GetLocationGroups.Invoke(new()
        ///     {
        ///         Name = "IoT Traffic Group",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zia = Pulumi.Zia;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zia.GetLocationGroups.Invoke(new()
        ///     {
        ///         Name = "Server Traffic Group",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zia = Pulumi.Zia;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zia.GetLocationGroups.Invoke(new()
        ///     {
        ///         Name = "Server Traffic Group",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetLocationGroupsResult> Invoke(GetLocationGroupsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLocationGroupsResult>("zia:index/getLocationGroups:getLocationGroups", args ?? new GetLocationGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLocationGroupsArgs : global::Pulumi.InvokeArgs
    {
        [Input("dynamicLocationGroupCriterias")]
        private List<Inputs.GetLocationGroupsDynamicLocationGroupCriteriaArgs>? _dynamicLocationGroupCriterias;

        /// <summary>
        /// (Block Set) Dynamic location group information.
        /// </summary>
        public List<Inputs.GetLocationGroupsDynamicLocationGroupCriteriaArgs> DynamicLocationGroupCriterias
        {
            get => _dynamicLocationGroupCriterias ?? (_dynamicLocationGroupCriterias = new List<Inputs.GetLocationGroupsDynamicLocationGroupCriteriaArgs>());
            set => _dynamicLocationGroupCriterias = value;
        }

        /// <summary>
        /// Location group name
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetLocationGroupsArgs()
        {
        }
        public static new GetLocationGroupsArgs Empty => new GetLocationGroupsArgs();
    }

    public sealed class GetLocationGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("dynamicLocationGroupCriterias")]
        private InputList<Inputs.GetLocationGroupsDynamicLocationGroupCriteriaInputArgs>? _dynamicLocationGroupCriterias;

        /// <summary>
        /// (Block Set) Dynamic location group information.
        /// </summary>
        public InputList<Inputs.GetLocationGroupsDynamicLocationGroupCriteriaInputArgs> DynamicLocationGroupCriterias
        {
            get => _dynamicLocationGroupCriterias ?? (_dynamicLocationGroupCriterias = new InputList<Inputs.GetLocationGroupsDynamicLocationGroupCriteriaInputArgs>());
            set => _dynamicLocationGroupCriterias = value;
        }

        /// <summary>
        /// Location group name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetLocationGroupsInvokeArgs()
        {
        }
        public static new GetLocationGroupsInvokeArgs Empty => new GetLocationGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetLocationGroupsResult
    {
        /// <summary>
        /// (List of Object)
        /// </summary>
        public readonly string Comments;
        /// <summary>
        /// (Boolean) Indicates the location group was deleted
        /// </summary>
        public readonly bool Deleted;
        /// <summary>
        /// (Block Set) Dynamic location group information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLocationGroupsDynamicLocationGroupCriteriaResult> DynamicLocationGroupCriterias;
        /// <summary>
        /// (String) The location group's type (i.e., Static or Dynamic)
        /// </summary>
        public readonly string GroupType;
        /// <summary>
        /// (Number) Identifier that uniquely identifies an entity
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// (List of Object) Automatically populated with the current time, after a successful POST or PUT request.
        /// </summary>
        public readonly int LastModTime;
        /// <summary>
        /// (List of Object)
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLocationGroupsLastModUserResult> LastModUsers;
        /// <summary>
        /// (List of Object) The Name-ID pairs of the locations that are assigned to the static location group. This is ignored if the groupType is Dynamic.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLocationGroupsLocationResult> Locations;
        /// <summary>
        /// (String) The configured name of the entity
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// (Boolean)
        /// </summary>
        public readonly bool Predefined;

        [OutputConstructor]
        private GetLocationGroupsResult(
            string comments,

            bool deleted,

            ImmutableArray<Outputs.GetLocationGroupsDynamicLocationGroupCriteriaResult> dynamicLocationGroupCriterias,

            string groupType,

            int id,

            int lastModTime,

            ImmutableArray<Outputs.GetLocationGroupsLastModUserResult> lastModUsers,

            ImmutableArray<Outputs.GetLocationGroupsLocationResult> locations,

            string? name,

            bool predefined)
        {
            Comments = comments;
            Deleted = deleted;
            DynamicLocationGroupCriterias = dynamicLocationGroupCriterias;
            GroupType = groupType;
            Id = id;
            LastModTime = lastModTime;
            LastModUsers = lastModUsers;
            Locations = locations;
            Name = name;
            Predefined = predefined;
        }
    }
}
