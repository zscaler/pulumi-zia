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
    public static class GetWorkloadGroups
    {
        /// <summary>
        /// Use the **zia_workload_groups** data source to get information about Workload Groups in the Zscaler Internet Access cloud or via the API. This data source can then be used as a criterion in ZIA policies such as, Firewall Filtering, URL Filtering, and Data Loss Prevention (DLP) to apply security policies to the workload traffic.
        /// 
        /// ## Example Usage
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
        ///     var ios = Zia.GetWorkloadGroups.Invoke(new()
        ///     {
        ///         Name = "Example",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetWorkloadGroupsResult> InvokeAsync(GetWorkloadGroupsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWorkloadGroupsResult>("zia:index/getWorkloadGroups:getWorkloadGroups", args ?? new GetWorkloadGroupsArgs(), options.WithDefaults());

        /// <summary>
        /// Use the **zia_workload_groups** data source to get information about Workload Groups in the Zscaler Internet Access cloud or via the API. This data source can then be used as a criterion in ZIA policies such as, Firewall Filtering, URL Filtering, and Data Loss Prevention (DLP) to apply security policies to the workload traffic.
        /// 
        /// ## Example Usage
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
        ///     var ios = Zia.GetWorkloadGroups.Invoke(new()
        ///     {
        ///         Name = "Example",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetWorkloadGroupsResult> Invoke(GetWorkloadGroupsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWorkloadGroupsResult>("zia:index/getWorkloadGroups:getWorkloadGroups", args ?? new GetWorkloadGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWorkloadGroupsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the workload group to be exported.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetWorkloadGroupsArgs()
        {
        }
        public static new GetWorkloadGroupsArgs Empty => new GetWorkloadGroupsArgs();
    }

    public sealed class GetWorkloadGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the workload group to be exported.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetWorkloadGroupsInvokeArgs()
        {
        }
        public static new GetWorkloadGroupsInvokeArgs Empty => new GetWorkloadGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetWorkloadGroupsResult
    {
        /// <summary>
        /// (String) The description of the workload group.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (String) The workload group expression containing tag types, tags, and their relationships.
        /// </summary>
        public readonly string Expression;
        /// <summary>
        /// (List) The workload group expression containing tag types, tags, and their relationships represented in a JSON format.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetWorkloadGroupsExpressionJsonResult> ExpressionJsons;
        /// <summary>
        /// (Number) Identifier that uniquely identifies an entity
        /// </summary>
        public readonly int Id;
        public readonly ImmutableArray<Outputs.GetWorkloadGroupsLastModifiedByResult> LastModifiedBies;
        /// <summary>
        /// (Number) When the rule was last modified
        /// </summary>
        public readonly int LastModifiedTime;
        /// <summary>
        /// (String) The configured name of the entity
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private GetWorkloadGroupsResult(
            string description,

            string expression,

            ImmutableArray<Outputs.GetWorkloadGroupsExpressionJsonResult> expressionJsons,

            int id,

            ImmutableArray<Outputs.GetWorkloadGroupsLastModifiedByResult> lastModifiedBies,

            int lastModifiedTime,

            string? name)
        {
            Description = description;
            Expression = expression;
            ExpressionJsons = expressionJsons;
            Id = id;
            LastModifiedBies = lastModifiedBies;
            LastModifiedTime = lastModifiedTime;
            Name = name;
        }
    }
}
