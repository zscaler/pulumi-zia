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
    public static class GetDLPWebRules
    {
        /// <summary>
        /// Use the **zia_dlp_web_rules** data source to get information about a ZIA DLP Web Rules in the Zscaler Internet Access cloud or via the API.
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
        ///     var example = Zia.GetDLPWebRules.Invoke(new()
        ///     {
        ///         Name = "Example",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetDLPWebRulesResult> InvokeAsync(GetDLPWebRulesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDLPWebRulesResult>("zia:index/getDLPWebRules:getDLPWebRules", args ?? new GetDLPWebRulesArgs(), options.WithDefaults());

        /// <summary>
        /// Use the **zia_dlp_web_rules** data source to get information about a ZIA DLP Web Rules in the Zscaler Internet Access cloud or via the API.
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
        ///     var example = Zia.GetDLPWebRules.Invoke(new()
        ///     {
        ///         Name = "Example",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetDLPWebRulesResult> Invoke(GetDLPWebRulesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDLPWebRulesResult>("zia:index/getDLPWebRules:getDLPWebRules", args ?? new GetDLPWebRulesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDLPWebRulesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier assigned to the workload group
        /// </summary>
        [Input("id")]
        public int? Id { get; set; }

        /// <summary>
        /// The name of the workload group
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetDLPWebRulesArgs()
        {
        }
        public static new GetDLPWebRulesArgs Empty => new GetDLPWebRulesArgs();
    }

    public sealed class GetDLPWebRulesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique identifier assigned to the workload group
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// The name of the workload group
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetDLPWebRulesInvokeArgs()
        {
        }
        public static new GetDLPWebRulesInvokeArgs Empty => new GetDLPWebRulesInvokeArgs();
    }


    [OutputType]
    public sealed class GetDLPWebRulesResult
    {
        public readonly string AccessControl;
        public readonly string Action;
        public readonly ImmutableArray<Outputs.GetDLPWebRulesAuditorResult> Auditors;
        public readonly ImmutableArray<string> CloudApplications;
        public readonly ImmutableArray<Outputs.GetDLPWebRulesDepartmentResult> Departments;
        public readonly string Description;
        public readonly bool DlpDownloadScanEnabled;
        public readonly ImmutableArray<Outputs.GetDLPWebRulesDlpEngineResult> DlpEngines;
        public readonly ImmutableArray<Outputs.GetDLPWebRulesExcludedDepartmentResult> ExcludedDepartments;
        public readonly ImmutableArray<Outputs.GetDLPWebRulesExcludedGroupResult> ExcludedGroups;
        public readonly ImmutableArray<Outputs.GetDLPWebRulesExcludedUserResult> ExcludedUsers;
        public readonly string ExternalAuditorEmail;
        public readonly ImmutableArray<string> FileTypes;
        public readonly ImmutableArray<Outputs.GetDLPWebRulesGroupResult> Groups;
        public readonly ImmutableArray<Outputs.GetDLPWebRulesIcapServerResult> IcapServers;
        public readonly int? Id;
        public readonly ImmutableArray<Outputs.GetDLPWebRulesLabelResult> Labels;
        public readonly ImmutableArray<Outputs.GetDLPWebRulesLastModifiedByResult> LastModifiedBies;
        public readonly int LastModifiedTime;
        public readonly ImmutableArray<Outputs.GetDLPWebRulesLocationGroupResult> LocationGroups;
        public readonly ImmutableArray<Outputs.GetDLPWebRulesLocationResult> Locations;
        public readonly bool MatchOnly;
        public readonly int MinSize;
        public readonly string? Name;
        public readonly ImmutableArray<Outputs.GetDLPWebRulesNotificationTemplateResult> NotificationTemplates;
        public readonly bool OcrEnabled;
        public readonly int Order;
        public readonly int ParentRule;
        public readonly ImmutableArray<string> Protocols;
        public readonly int Rank;
        public readonly string Severity;
        public readonly string State;
        public readonly ImmutableArray<string> SubRules;
        public readonly ImmutableArray<Outputs.GetDLPWebRulesTimeWindowResult> TimeWindows;
        public readonly ImmutableArray<Outputs.GetDLPWebRulesUrlCategoryResult> UrlCategories;
        public readonly ImmutableArray<Outputs.GetDLPWebRulesUserResult> Users;
        public readonly bool WithoutContentInspection;
        public readonly ImmutableArray<Outputs.GetDLPWebRulesWorkloadGroupResult> WorkloadGroups;
        public readonly bool ZccNotificationsEnabled;
        public readonly bool ZscalerIncidentReceiver;

        [OutputConstructor]
        private GetDLPWebRulesResult(
            string accessControl,

            string action,

            ImmutableArray<Outputs.GetDLPWebRulesAuditorResult> auditors,

            ImmutableArray<string> cloudApplications,

            ImmutableArray<Outputs.GetDLPWebRulesDepartmentResult> departments,

            string description,

            bool dlpDownloadScanEnabled,

            ImmutableArray<Outputs.GetDLPWebRulesDlpEngineResult> dlpEngines,

            ImmutableArray<Outputs.GetDLPWebRulesExcludedDepartmentResult> excludedDepartments,

            ImmutableArray<Outputs.GetDLPWebRulesExcludedGroupResult> excludedGroups,

            ImmutableArray<Outputs.GetDLPWebRulesExcludedUserResult> excludedUsers,

            string externalAuditorEmail,

            ImmutableArray<string> fileTypes,

            ImmutableArray<Outputs.GetDLPWebRulesGroupResult> groups,

            ImmutableArray<Outputs.GetDLPWebRulesIcapServerResult> icapServers,

            int? id,

            ImmutableArray<Outputs.GetDLPWebRulesLabelResult> labels,

            ImmutableArray<Outputs.GetDLPWebRulesLastModifiedByResult> lastModifiedBies,

            int lastModifiedTime,

            ImmutableArray<Outputs.GetDLPWebRulesLocationGroupResult> locationGroups,

            ImmutableArray<Outputs.GetDLPWebRulesLocationResult> locations,

            bool matchOnly,

            int minSize,

            string? name,

            ImmutableArray<Outputs.GetDLPWebRulesNotificationTemplateResult> notificationTemplates,

            bool ocrEnabled,

            int order,

            int parentRule,

            ImmutableArray<string> protocols,

            int rank,

            string severity,

            string state,

            ImmutableArray<string> subRules,

            ImmutableArray<Outputs.GetDLPWebRulesTimeWindowResult> timeWindows,

            ImmutableArray<Outputs.GetDLPWebRulesUrlCategoryResult> urlCategories,

            ImmutableArray<Outputs.GetDLPWebRulesUserResult> users,

            bool withoutContentInspection,

            ImmutableArray<Outputs.GetDLPWebRulesWorkloadGroupResult> workloadGroups,

            bool zccNotificationsEnabled,

            bool zscalerIncidentReceiver)
        {
            AccessControl = accessControl;
            Action = action;
            Auditors = auditors;
            CloudApplications = cloudApplications;
            Departments = departments;
            Description = description;
            DlpDownloadScanEnabled = dlpDownloadScanEnabled;
            DlpEngines = dlpEngines;
            ExcludedDepartments = excludedDepartments;
            ExcludedGroups = excludedGroups;
            ExcludedUsers = excludedUsers;
            ExternalAuditorEmail = externalAuditorEmail;
            FileTypes = fileTypes;
            Groups = groups;
            IcapServers = icapServers;
            Id = id;
            Labels = labels;
            LastModifiedBies = lastModifiedBies;
            LastModifiedTime = lastModifiedTime;
            LocationGroups = locationGroups;
            Locations = locations;
            MatchOnly = matchOnly;
            MinSize = minSize;
            Name = name;
            NotificationTemplates = notificationTemplates;
            OcrEnabled = ocrEnabled;
            Order = order;
            ParentRule = parentRule;
            Protocols = protocols;
            Rank = rank;
            Severity = severity;
            State = state;
            SubRules = subRules;
            TimeWindows = timeWindows;
            UrlCategories = urlCategories;
            Users = users;
            WithoutContentInspection = withoutContentInspection;
            WorkloadGroups = workloadGroups;
            ZccNotificationsEnabled = zccNotificationsEnabled;
            ZscalerIncidentReceiver = zscalerIncidentReceiver;
        }
    }
}
