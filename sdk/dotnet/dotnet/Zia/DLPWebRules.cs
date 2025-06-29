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
    /// <summary>
    /// * [Official documentation](https://help.zscaler.com/zia/configuring-dlp-policy-rules-content-inspection#Rules)
    /// * [API documentation](https://help.zscaler.com/zia/data-loss-prevention#/webDlpRules-get)
    /// 
    /// The **zia_dlp_web_rules** resource allows the creation and management of ZIA DLP Web Rules in the Zscaler Internet Access cloud or via the API.
    /// 
    /// ⚠️ **WARNING:** Zscaler Internet Access DLP supports a maximum of 127 Web DLP Rules to be created via API.
    /// 
    /// ## Example Usage
    /// 
    /// ### "FTCATEGORY_ALL_OUTBOUND" File Type"
    /// 
    /// ### "Specify Incident Receiver Setting"
    /// 
    /// ### "Creating Parent Rules And SubRules"
    /// 
    /// ⚠️ **WARNING:** Destroying a parent rule will also destroy all subrules
    /// 
    ///  **NOTE** Exception rules can be configured only when the inline DLP rule evaluation type is set
    ///  to evaluate all DLP rules in the DLP Advanced Settings.
    ///  To learn more, see [Configuring DLP Advanced Settings](https://help.zscaler.com/%22/zia/configuring-dlp-advanced-settings/%22)
    /// 
    /// ## Import
    /// 
    /// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZIA configurations into Terraform-compliant HashiCorp Configuration Language.
    /// 
    /// Visit
    /// 
    /// **zia_dlp_web_rules** can be imported by using `&lt;RULE ID&gt;` or `&lt;RULE NAME&gt;` as the import ID.
    /// 
    /// For example:
    /// 
    /// ```sh
    /// $ pulumi import zia:index/dLPWebRules:DLPWebRules example &lt;rule_id&gt;
    /// ```
    /// 
    /// or
    /// 
    /// ```sh
    /// $ pulumi import zia:index/dLPWebRules:DLPWebRules example &lt;rule_name&gt;
    /// ```
    /// </summary>
    [ZiaResourceType("zia:index/dLPWebRules:DLPWebRules")]
    public partial class DLPWebRules : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The action taken when traffic matches the DLP policy rule criteria.
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// The auditor to which the DLP policy rule must be applied.
        /// </summary>
        [Output("auditors")]
        public Output<ImmutableArray<Outputs.DLPWebRulesAuditor>> Auditors { get; private set; } = null!;

        /// <summary>
        /// The list of cloud applications to which the DLP policy rule must be applied.
        /// </summary>
        [Output("cloudApplications")]
        public Output<ImmutableArray<string>> CloudApplications { get; private set; } = null!;

        /// <summary>
        /// The Name-ID pairs of departments to which the DLP policy rule must be applied.
        /// </summary>
        [Output("departments")]
        public Output<Outputs.DLPWebRulesDepartments?> Departments { get; private set; } = null!;

        /// <summary>
        /// The description of the DLP policy rule.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Indicates a DLP policy rule without content inspection, when the value is set to true.
        /// </summary>
        [Output("dlpDownloadScanEnabled")]
        public Output<bool> DlpDownloadScanEnabled { get; private set; } = null!;

        /// <summary>
        /// The list of DLP engines to which the DLP policy rule must be applied.
        /// </summary>
        [Output("dlpEngines")]
        public Output<Outputs.DLPWebRulesDlpEngines?> DlpEngines { get; private set; } = null!;

        /// <summary>
        /// The Name-ID pairs of users to which the DLP policy rule must be applied.
        /// </summary>
        [Output("excludedDepartments")]
        public Output<Outputs.DLPWebRulesExcludedDepartments?> ExcludedDepartments { get; private set; } = null!;

        /// <summary>
        /// The Name-ID pairs of users to which the DLP policy rule must be applied.
        /// </summary>
        [Output("excludedDomainProfiles")]
        public Output<Outputs.DLPWebRulesExcludedDomainProfiles?> ExcludedDomainProfiles { get; private set; } = null!;

        /// <summary>
        /// The Name-ID pairs of users to which the DLP policy rule must be applied.
        /// </summary>
        [Output("excludedGroups")]
        public Output<Outputs.DLPWebRulesExcludedGroups?> ExcludedGroups { get; private set; } = null!;

        /// <summary>
        /// The Name-ID pairs of users to which the DLP policy rule must be applied.
        /// </summary>
        [Output("excludedUsers")]
        public Output<Outputs.DLPWebRulesExcludedUsers?> ExcludedUsers { get; private set; } = null!;

        /// <summary>
        /// The email address of an external auditor to whom DLP email notifications are sent
        /// </summary>
        [Output("externalAuditorEmail")]
        public Output<string?> ExternalAuditorEmail { get; private set; } = null!;

        /// <summary>
        /// The list of file types for which the DLP policy rule must be applied.
        /// </summary>
        [Output("fileTypes")]
        public Output<ImmutableArray<string>> FileTypes { get; private set; } = null!;

        /// <summary>
        /// The Name-ID pairs of groups to which the DLP policy rule must be applied.
        /// </summary>
        [Output("groups")]
        public Output<Outputs.DLPWebRulesGroups?> Groups { get; private set; } = null!;

        /// <summary>
        /// The DLP server, using ICAP, to which the transaction content is forwarded.
        /// </summary>
        [Output("icapServers")]
        public Output<ImmutableArray<Outputs.DLPWebRulesIcapServer>> IcapServers { get; private set; } = null!;

        /// <summary>
        /// The Name-ID pairs of users to which the DLP policy rule must be applied.
        /// </summary>
        [Output("includedDomainProfiles")]
        public Output<Outputs.DLPWebRulesIncludedDomainProfiles?> IncludedDomainProfiles { get; private set; } = null!;

        [Output("inspectHttpGetEnabled")]
        public Output<bool?> InspectHttpGetEnabled { get; private set; } = null!;

        /// <summary>
        /// list of Labels that are applicable to the rule.
        /// </summary>
        [Output("labels")]
        public Output<Outputs.DLPWebRulesLabels?> Labels { get; private set; } = null!;

        /// <summary>
        /// The Name-ID pairs of locations groups to which the DLP policy rule must be applied.
        /// </summary>
        [Output("locationGroups")]
        public Output<Outputs.DLPWebRulesLocationGroups?> LocationGroups { get; private set; } = null!;

        /// <summary>
        /// The Name-ID pairs of locations to which the DLP policy rule must be applied.
        /// </summary>
        [Output("locations")]
        public Output<Outputs.DLPWebRulesLocations?> Locations { get; private set; } = null!;

        /// <summary>
        /// The match only criteria for DLP engines.
        /// </summary>
        [Output("matchOnly")]
        public Output<bool> MatchOnly { get; private set; } = null!;

        /// <summary>
        /// The minimum file size (in KB) used for evaluation of the DLP policy rule.
        /// </summary>
        [Output("minSize")]
        public Output<int> MinSize { get; private set; } = null!;

        /// <summary>
        /// The DLP policy rule name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The template used for DLP notification emails.
        /// </summary>
        [Output("notificationTemplates")]
        public Output<ImmutableArray<Outputs.DLPWebRulesNotificationTemplate>> NotificationTemplates { get; private set; } = null!;

        /// <summary>
        /// The rule order of execution for the DLP policy rule with respect to other rules.
        /// </summary>
        [Output("order")]
        public Output<int> Order { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the parent rule under which an exception rule is added
        /// </summary>
        [Output("parentRule")]
        public Output<int> ParentRule { get; private set; } = null!;

        /// <summary>
        /// The protocol criteria specified for the DLP policy rule.
        /// </summary>
        [Output("protocols")]
        public Output<ImmutableArray<string>> Protocols { get; private set; } = null!;

        /// <summary>
        /// Admin rank of the admin who creates this rule
        /// </summary>
        [Output("rank")]
        public Output<int> Rank { get; private set; } = null!;

        [Output("ruleId")]
        public Output<int> RuleId { get; private set; } = null!;

        /// <summary>
        /// Indicates the severity selected for the DLP rule violation
        /// </summary>
        [Output("severity")]
        public Output<string> Severity { get; private set; } = null!;

        /// <summary>
        /// list of source ip groups
        /// </summary>
        [Output("sourceIpGroups")]
        public Output<Outputs.DLPWebRulesSourceIpGroups?> SourceIpGroups { get; private set; } = null!;

        /// <summary>
        /// Enables or disables the DLP policy rule.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The list of exception rules added to a parent rule
        /// </summary>
        [Output("subRules")]
        public Output<ImmutableArray<string>> SubRules { get; private set; } = null!;

        /// <summary>
        /// list of time interval during which rule must be enforced.
        /// </summary>
        [Output("timeWindows")]
        public Output<Outputs.DLPWebRulesTimeWindows?> TimeWindows { get; private set; } = null!;

        /// <summary>
        /// The list of URL categories to which the DLP policy rule must be applied.
        /// </summary>
        [Output("urlCategories")]
        public Output<Outputs.DLPWebRulesUrlCategories?> UrlCategories { get; private set; } = null!;

        [Output("userRiskScoreLevels")]
        public Output<ImmutableArray<string>> UserRiskScoreLevels { get; private set; } = null!;

        /// <summary>
        /// The Name-ID pairs of users to which the DLP policy rule must be applied.
        /// </summary>
        [Output("users")]
        public Output<Outputs.DLPWebRulesUsers?> Users { get; private set; } = null!;

        /// <summary>
        /// Indicates a DLP policy rule without content inspection, when the value is set to true.
        /// </summary>
        [Output("withoutContentInspection")]
        public Output<bool> WithoutContentInspection { get; private set; } = null!;

        /// <summary>
        /// The list of preconfigured workload groups to which the policy must be applied
        /// </summary>
        [Output("workloadGroups")]
        public Output<ImmutableArray<Outputs.DLPWebRulesWorkloadGroup>> WorkloadGroups { get; private set; } = null!;

        /// <summary>
        /// Indicates a DLP policy rule without content inspection, when the value is set to true.
        /// </summary>
        [Output("zccNotificationsEnabled")]
        public Output<bool> ZccNotificationsEnabled { get; private set; } = null!;

        /// <summary>
        /// Indicates whether a Zscaler Incident Receiver is associated to the DLP policy rule.
        /// </summary>
        [Output("zscalerIncidentReceiver")]
        public Output<bool?> ZscalerIncidentReceiver { get; private set; } = null!;


        /// <summary>
        /// Create a DLPWebRules resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DLPWebRules(string name, DLPWebRulesArgs args, CustomResourceOptions? options = null)
            : base("zia:index/dLPWebRules:DLPWebRules", name, args ?? new DLPWebRulesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DLPWebRules(string name, Input<string> id, DLPWebRulesState? state = null, CustomResourceOptions? options = null)
            : base("zia:index/dLPWebRules:DLPWebRules", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/zscaler",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DLPWebRules resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DLPWebRules Get(string name, Input<string> id, DLPWebRulesState? state = null, CustomResourceOptions? options = null)
        {
            return new DLPWebRules(name, id, state, options);
        }
    }

    public sealed class DLPWebRulesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action taken when traffic matches the DLP policy rule criteria.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("auditors")]
        private InputList<Inputs.DLPWebRulesAuditorArgs>? _auditors;

        /// <summary>
        /// The auditor to which the DLP policy rule must be applied.
        /// </summary>
        public InputList<Inputs.DLPWebRulesAuditorArgs> Auditors
        {
            get => _auditors ?? (_auditors = new InputList<Inputs.DLPWebRulesAuditorArgs>());
            set => _auditors = value;
        }

        [Input("cloudApplications")]
        private InputList<string>? _cloudApplications;

        /// <summary>
        /// The list of cloud applications to which the DLP policy rule must be applied.
        /// </summary>
        public InputList<string> CloudApplications
        {
            get => _cloudApplications ?? (_cloudApplications = new InputList<string>());
            set => _cloudApplications = value;
        }

        /// <summary>
        /// The Name-ID pairs of departments to which the DLP policy rule must be applied.
        /// </summary>
        [Input("departments")]
        public Input<Inputs.DLPWebRulesDepartmentsArgs>? Departments { get; set; }

        /// <summary>
        /// The description of the DLP policy rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates a DLP policy rule without content inspection, when the value is set to true.
        /// </summary>
        [Input("dlpDownloadScanEnabled")]
        public Input<bool>? DlpDownloadScanEnabled { get; set; }

        /// <summary>
        /// The list of DLP engines to which the DLP policy rule must be applied.
        /// </summary>
        [Input("dlpEngines")]
        public Input<Inputs.DLPWebRulesDlpEnginesArgs>? DlpEngines { get; set; }

        /// <summary>
        /// The Name-ID pairs of users to which the DLP policy rule must be applied.
        /// </summary>
        [Input("excludedDepartments")]
        public Input<Inputs.DLPWebRulesExcludedDepartmentsArgs>? ExcludedDepartments { get; set; }

        /// <summary>
        /// The Name-ID pairs of users to which the DLP policy rule must be applied.
        /// </summary>
        [Input("excludedDomainProfiles")]
        public Input<Inputs.DLPWebRulesExcludedDomainProfilesArgs>? ExcludedDomainProfiles { get; set; }

        /// <summary>
        /// The Name-ID pairs of users to which the DLP policy rule must be applied.
        /// </summary>
        [Input("excludedGroups")]
        public Input<Inputs.DLPWebRulesExcludedGroupsArgs>? ExcludedGroups { get; set; }

        /// <summary>
        /// The Name-ID pairs of users to which the DLP policy rule must be applied.
        /// </summary>
        [Input("excludedUsers")]
        public Input<Inputs.DLPWebRulesExcludedUsersArgs>? ExcludedUsers { get; set; }

        /// <summary>
        /// The email address of an external auditor to whom DLP email notifications are sent
        /// </summary>
        [Input("externalAuditorEmail")]
        public Input<string>? ExternalAuditorEmail { get; set; }

        [Input("fileTypes")]
        private InputList<string>? _fileTypes;

        /// <summary>
        /// The list of file types for which the DLP policy rule must be applied.
        /// </summary>
        public InputList<string> FileTypes
        {
            get => _fileTypes ?? (_fileTypes = new InputList<string>());
            set => _fileTypes = value;
        }

        /// <summary>
        /// The Name-ID pairs of groups to which the DLP policy rule must be applied.
        /// </summary>
        [Input("groups")]
        public Input<Inputs.DLPWebRulesGroupsArgs>? Groups { get; set; }

        [Input("icapServers")]
        private InputList<Inputs.DLPWebRulesIcapServerArgs>? _icapServers;

        /// <summary>
        /// The DLP server, using ICAP, to which the transaction content is forwarded.
        /// </summary>
        public InputList<Inputs.DLPWebRulesIcapServerArgs> IcapServers
        {
            get => _icapServers ?? (_icapServers = new InputList<Inputs.DLPWebRulesIcapServerArgs>());
            set => _icapServers = value;
        }

        /// <summary>
        /// The Name-ID pairs of users to which the DLP policy rule must be applied.
        /// </summary>
        [Input("includedDomainProfiles")]
        public Input<Inputs.DLPWebRulesIncludedDomainProfilesArgs>? IncludedDomainProfiles { get; set; }

        [Input("inspectHttpGetEnabled")]
        public Input<bool>? InspectHttpGetEnabled { get; set; }

        /// <summary>
        /// list of Labels that are applicable to the rule.
        /// </summary>
        [Input("labels")]
        public Input<Inputs.DLPWebRulesLabelsArgs>? Labels { get; set; }

        /// <summary>
        /// The Name-ID pairs of locations groups to which the DLP policy rule must be applied.
        /// </summary>
        [Input("locationGroups")]
        public Input<Inputs.DLPWebRulesLocationGroupsArgs>? LocationGroups { get; set; }

        /// <summary>
        /// The Name-ID pairs of locations to which the DLP policy rule must be applied.
        /// </summary>
        [Input("locations")]
        public Input<Inputs.DLPWebRulesLocationsArgs>? Locations { get; set; }

        /// <summary>
        /// The match only criteria for DLP engines.
        /// </summary>
        [Input("matchOnly")]
        public Input<bool>? MatchOnly { get; set; }

        /// <summary>
        /// The minimum file size (in KB) used for evaluation of the DLP policy rule.
        /// </summary>
        [Input("minSize")]
        public Input<int>? MinSize { get; set; }

        /// <summary>
        /// The DLP policy rule name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationTemplates")]
        private InputList<Inputs.DLPWebRulesNotificationTemplateArgs>? _notificationTemplates;

        /// <summary>
        /// The template used for DLP notification emails.
        /// </summary>
        public InputList<Inputs.DLPWebRulesNotificationTemplateArgs> NotificationTemplates
        {
            get => _notificationTemplates ?? (_notificationTemplates = new InputList<Inputs.DLPWebRulesNotificationTemplateArgs>());
            set => _notificationTemplates = value;
        }

        /// <summary>
        /// The rule order of execution for the DLP policy rule with respect to other rules.
        /// </summary>
        [Input("order", required: true)]
        public Input<int> Order { get; set; } = null!;

        /// <summary>
        /// The unique identifier of the parent rule under which an exception rule is added
        /// </summary>
        [Input("parentRule")]
        public Input<int>? ParentRule { get; set; }

        [Input("protocols")]
        private InputList<string>? _protocols;

        /// <summary>
        /// The protocol criteria specified for the DLP policy rule.
        /// </summary>
        public InputList<string> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<string>());
            set => _protocols = value;
        }

        /// <summary>
        /// Admin rank of the admin who creates this rule
        /// </summary>
        [Input("rank")]
        public Input<int>? Rank { get; set; }

        /// <summary>
        /// Indicates the severity selected for the DLP rule violation
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        /// <summary>
        /// list of source ip groups
        /// </summary>
        [Input("sourceIpGroups")]
        public Input<Inputs.DLPWebRulesSourceIpGroupsArgs>? SourceIpGroups { get; set; }

        /// <summary>
        /// Enables or disables the DLP policy rule.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("subRules")]
        private InputList<string>? _subRules;

        /// <summary>
        /// The list of exception rules added to a parent rule
        /// </summary>
        public InputList<string> SubRules
        {
            get => _subRules ?? (_subRules = new InputList<string>());
            set => _subRules = value;
        }

        /// <summary>
        /// list of time interval during which rule must be enforced.
        /// </summary>
        [Input("timeWindows")]
        public Input<Inputs.DLPWebRulesTimeWindowsArgs>? TimeWindows { get; set; }

        /// <summary>
        /// The list of URL categories to which the DLP policy rule must be applied.
        /// </summary>
        [Input("urlCategories")]
        public Input<Inputs.DLPWebRulesUrlCategoriesArgs>? UrlCategories { get; set; }

        [Input("userRiskScoreLevels")]
        private InputList<string>? _userRiskScoreLevels;
        public InputList<string> UserRiskScoreLevels
        {
            get => _userRiskScoreLevels ?? (_userRiskScoreLevels = new InputList<string>());
            set => _userRiskScoreLevels = value;
        }

        /// <summary>
        /// The Name-ID pairs of users to which the DLP policy rule must be applied.
        /// </summary>
        [Input("users")]
        public Input<Inputs.DLPWebRulesUsersArgs>? Users { get; set; }

        /// <summary>
        /// Indicates a DLP policy rule without content inspection, when the value is set to true.
        /// </summary>
        [Input("withoutContentInspection")]
        public Input<bool>? WithoutContentInspection { get; set; }

        [Input("workloadGroups")]
        private InputList<Inputs.DLPWebRulesWorkloadGroupArgs>? _workloadGroups;

        /// <summary>
        /// The list of preconfigured workload groups to which the policy must be applied
        /// </summary>
        public InputList<Inputs.DLPWebRulesWorkloadGroupArgs> WorkloadGroups
        {
            get => _workloadGroups ?? (_workloadGroups = new InputList<Inputs.DLPWebRulesWorkloadGroupArgs>());
            set => _workloadGroups = value;
        }

        /// <summary>
        /// Indicates a DLP policy rule without content inspection, when the value is set to true.
        /// </summary>
        [Input("zccNotificationsEnabled")]
        public Input<bool>? ZccNotificationsEnabled { get; set; }

        /// <summary>
        /// Indicates whether a Zscaler Incident Receiver is associated to the DLP policy rule.
        /// </summary>
        [Input("zscalerIncidentReceiver")]
        public Input<bool>? ZscalerIncidentReceiver { get; set; }

        public DLPWebRulesArgs()
        {
        }
        public static new DLPWebRulesArgs Empty => new DLPWebRulesArgs();
    }

    public sealed class DLPWebRulesState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action taken when traffic matches the DLP policy rule criteria.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("auditors")]
        private InputList<Inputs.DLPWebRulesAuditorGetArgs>? _auditors;

        /// <summary>
        /// The auditor to which the DLP policy rule must be applied.
        /// </summary>
        public InputList<Inputs.DLPWebRulesAuditorGetArgs> Auditors
        {
            get => _auditors ?? (_auditors = new InputList<Inputs.DLPWebRulesAuditorGetArgs>());
            set => _auditors = value;
        }

        [Input("cloudApplications")]
        private InputList<string>? _cloudApplications;

        /// <summary>
        /// The list of cloud applications to which the DLP policy rule must be applied.
        /// </summary>
        public InputList<string> CloudApplications
        {
            get => _cloudApplications ?? (_cloudApplications = new InputList<string>());
            set => _cloudApplications = value;
        }

        /// <summary>
        /// The Name-ID pairs of departments to which the DLP policy rule must be applied.
        /// </summary>
        [Input("departments")]
        public Input<Inputs.DLPWebRulesDepartmentsGetArgs>? Departments { get; set; }

        /// <summary>
        /// The description of the DLP policy rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates a DLP policy rule without content inspection, when the value is set to true.
        /// </summary>
        [Input("dlpDownloadScanEnabled")]
        public Input<bool>? DlpDownloadScanEnabled { get; set; }

        /// <summary>
        /// The list of DLP engines to which the DLP policy rule must be applied.
        /// </summary>
        [Input("dlpEngines")]
        public Input<Inputs.DLPWebRulesDlpEnginesGetArgs>? DlpEngines { get; set; }

        /// <summary>
        /// The Name-ID pairs of users to which the DLP policy rule must be applied.
        /// </summary>
        [Input("excludedDepartments")]
        public Input<Inputs.DLPWebRulesExcludedDepartmentsGetArgs>? ExcludedDepartments { get; set; }

        /// <summary>
        /// The Name-ID pairs of users to which the DLP policy rule must be applied.
        /// </summary>
        [Input("excludedDomainProfiles")]
        public Input<Inputs.DLPWebRulesExcludedDomainProfilesGetArgs>? ExcludedDomainProfiles { get; set; }

        /// <summary>
        /// The Name-ID pairs of users to which the DLP policy rule must be applied.
        /// </summary>
        [Input("excludedGroups")]
        public Input<Inputs.DLPWebRulesExcludedGroupsGetArgs>? ExcludedGroups { get; set; }

        /// <summary>
        /// The Name-ID pairs of users to which the DLP policy rule must be applied.
        /// </summary>
        [Input("excludedUsers")]
        public Input<Inputs.DLPWebRulesExcludedUsersGetArgs>? ExcludedUsers { get; set; }

        /// <summary>
        /// The email address of an external auditor to whom DLP email notifications are sent
        /// </summary>
        [Input("externalAuditorEmail")]
        public Input<string>? ExternalAuditorEmail { get; set; }

        [Input("fileTypes")]
        private InputList<string>? _fileTypes;

        /// <summary>
        /// The list of file types for which the DLP policy rule must be applied.
        /// </summary>
        public InputList<string> FileTypes
        {
            get => _fileTypes ?? (_fileTypes = new InputList<string>());
            set => _fileTypes = value;
        }

        /// <summary>
        /// The Name-ID pairs of groups to which the DLP policy rule must be applied.
        /// </summary>
        [Input("groups")]
        public Input<Inputs.DLPWebRulesGroupsGetArgs>? Groups { get; set; }

        [Input("icapServers")]
        private InputList<Inputs.DLPWebRulesIcapServerGetArgs>? _icapServers;

        /// <summary>
        /// The DLP server, using ICAP, to which the transaction content is forwarded.
        /// </summary>
        public InputList<Inputs.DLPWebRulesIcapServerGetArgs> IcapServers
        {
            get => _icapServers ?? (_icapServers = new InputList<Inputs.DLPWebRulesIcapServerGetArgs>());
            set => _icapServers = value;
        }

        /// <summary>
        /// The Name-ID pairs of users to which the DLP policy rule must be applied.
        /// </summary>
        [Input("includedDomainProfiles")]
        public Input<Inputs.DLPWebRulesIncludedDomainProfilesGetArgs>? IncludedDomainProfiles { get; set; }

        [Input("inspectHttpGetEnabled")]
        public Input<bool>? InspectHttpGetEnabled { get; set; }

        /// <summary>
        /// list of Labels that are applicable to the rule.
        /// </summary>
        [Input("labels")]
        public Input<Inputs.DLPWebRulesLabelsGetArgs>? Labels { get; set; }

        /// <summary>
        /// The Name-ID pairs of locations groups to which the DLP policy rule must be applied.
        /// </summary>
        [Input("locationGroups")]
        public Input<Inputs.DLPWebRulesLocationGroupsGetArgs>? LocationGroups { get; set; }

        /// <summary>
        /// The Name-ID pairs of locations to which the DLP policy rule must be applied.
        /// </summary>
        [Input("locations")]
        public Input<Inputs.DLPWebRulesLocationsGetArgs>? Locations { get; set; }

        /// <summary>
        /// The match only criteria for DLP engines.
        /// </summary>
        [Input("matchOnly")]
        public Input<bool>? MatchOnly { get; set; }

        /// <summary>
        /// The minimum file size (in KB) used for evaluation of the DLP policy rule.
        /// </summary>
        [Input("minSize")]
        public Input<int>? MinSize { get; set; }

        /// <summary>
        /// The DLP policy rule name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationTemplates")]
        private InputList<Inputs.DLPWebRulesNotificationTemplateGetArgs>? _notificationTemplates;

        /// <summary>
        /// The template used for DLP notification emails.
        /// </summary>
        public InputList<Inputs.DLPWebRulesNotificationTemplateGetArgs> NotificationTemplates
        {
            get => _notificationTemplates ?? (_notificationTemplates = new InputList<Inputs.DLPWebRulesNotificationTemplateGetArgs>());
            set => _notificationTemplates = value;
        }

        /// <summary>
        /// The rule order of execution for the DLP policy rule with respect to other rules.
        /// </summary>
        [Input("order")]
        public Input<int>? Order { get; set; }

        /// <summary>
        /// The unique identifier of the parent rule under which an exception rule is added
        /// </summary>
        [Input("parentRule")]
        public Input<int>? ParentRule { get; set; }

        [Input("protocols")]
        private InputList<string>? _protocols;

        /// <summary>
        /// The protocol criteria specified for the DLP policy rule.
        /// </summary>
        public InputList<string> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<string>());
            set => _protocols = value;
        }

        /// <summary>
        /// Admin rank of the admin who creates this rule
        /// </summary>
        [Input("rank")]
        public Input<int>? Rank { get; set; }

        [Input("ruleId")]
        public Input<int>? RuleId { get; set; }

        /// <summary>
        /// Indicates the severity selected for the DLP rule violation
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        /// <summary>
        /// list of source ip groups
        /// </summary>
        [Input("sourceIpGroups")]
        public Input<Inputs.DLPWebRulesSourceIpGroupsGetArgs>? SourceIpGroups { get; set; }

        /// <summary>
        /// Enables or disables the DLP policy rule.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("subRules")]
        private InputList<string>? _subRules;

        /// <summary>
        /// The list of exception rules added to a parent rule
        /// </summary>
        public InputList<string> SubRules
        {
            get => _subRules ?? (_subRules = new InputList<string>());
            set => _subRules = value;
        }

        /// <summary>
        /// list of time interval during which rule must be enforced.
        /// </summary>
        [Input("timeWindows")]
        public Input<Inputs.DLPWebRulesTimeWindowsGetArgs>? TimeWindows { get; set; }

        /// <summary>
        /// The list of URL categories to which the DLP policy rule must be applied.
        /// </summary>
        [Input("urlCategories")]
        public Input<Inputs.DLPWebRulesUrlCategoriesGetArgs>? UrlCategories { get; set; }

        [Input("userRiskScoreLevels")]
        private InputList<string>? _userRiskScoreLevels;
        public InputList<string> UserRiskScoreLevels
        {
            get => _userRiskScoreLevels ?? (_userRiskScoreLevels = new InputList<string>());
            set => _userRiskScoreLevels = value;
        }

        /// <summary>
        /// The Name-ID pairs of users to which the DLP policy rule must be applied.
        /// </summary>
        [Input("users")]
        public Input<Inputs.DLPWebRulesUsersGetArgs>? Users { get; set; }

        /// <summary>
        /// Indicates a DLP policy rule without content inspection, when the value is set to true.
        /// </summary>
        [Input("withoutContentInspection")]
        public Input<bool>? WithoutContentInspection { get; set; }

        [Input("workloadGroups")]
        private InputList<Inputs.DLPWebRulesWorkloadGroupGetArgs>? _workloadGroups;

        /// <summary>
        /// The list of preconfigured workload groups to which the policy must be applied
        /// </summary>
        public InputList<Inputs.DLPWebRulesWorkloadGroupGetArgs> WorkloadGroups
        {
            get => _workloadGroups ?? (_workloadGroups = new InputList<Inputs.DLPWebRulesWorkloadGroupGetArgs>());
            set => _workloadGroups = value;
        }

        /// <summary>
        /// Indicates a DLP policy rule without content inspection, when the value is set to true.
        /// </summary>
        [Input("zccNotificationsEnabled")]
        public Input<bool>? ZccNotificationsEnabled { get; set; }

        /// <summary>
        /// Indicates whether a Zscaler Incident Receiver is associated to the DLP policy rule.
        /// </summary>
        [Input("zscalerIncidentReceiver")]
        public Input<bool>? ZscalerIncidentReceiver { get; set; }

        public DLPWebRulesState()
        {
        }
        public static new DLPWebRulesState Empty => new DLPWebRulesState();
    }
}
