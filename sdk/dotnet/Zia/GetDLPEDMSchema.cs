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
    public static class GetDLPEDMSchema
    {
        /// <summary>
        /// Use the **zia_dlp_edm_schema** data source to get information about a the list of DLP Exact Data Match (EDM) templates in the Zscaler Internet Access cloud or via the API.
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
        ///     var @this = Zia.GetDLPEDMSchema.Invoke(new()
        ///     {
        ///         ProjectName = "Example",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetDLPEDMSchemaResult> InvokeAsync(GetDLPEDMSchemaArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDLPEDMSchemaResult>("zia:index/getDLPEDMSchema:getDLPEDMSchema", args ?? new GetDLPEDMSchemaArgs(), options.WithDefaults());

        /// <summary>
        /// Use the **zia_dlp_edm_schema** data source to get information about a the list of DLP Exact Data Match (EDM) templates in the Zscaler Internet Access cloud or via the API.
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
        ///     var @this = Zia.GetDLPEDMSchema.Invoke(new()
        ///     {
        ///         ProjectName = "Example",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDLPEDMSchemaResult> Invoke(GetDLPEDMSchemaInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDLPEDMSchemaResult>("zia:index/getDLPEDMSchema:getDLPEDMSchema", args ?? new GetDLPEDMSchemaInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDLPEDMSchemaArgs : global::Pulumi.InvokeArgs
    {
        [Input("projectName")]
        public string? ProjectName { get; set; }

        public GetDLPEDMSchemaArgs()
        {
        }
        public static new GetDLPEDMSchemaArgs Empty => new GetDLPEDMSchemaArgs();
    }

    public sealed class GetDLPEDMSchemaInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        public GetDLPEDMSchemaInvokeArgs()
        {
        }
        public static new GetDLPEDMSchemaInvokeArgs Empty => new GetDLPEDMSchemaInvokeArgs();
    }


    [OutputType]
    public sealed class GetDLPEDMSchemaResult
    {
        public readonly int CellsUsed;
        public readonly ImmutableArray<Outputs.GetDLPEDMSchemaCreatedByResult> CreatedBies;
        public readonly ImmutableArray<Outputs.GetDLPEDMSchemaEdmClientResult> EdmClients;
        public readonly string FileName;
        public readonly string FileUploadStatus;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetDLPEDMSchemaLastModifiedByResult> LastModifiedBies;
        public readonly int LastModifiedTime;
        public readonly int OrigColCount;
        public readonly string OriginalFileName;
        public readonly string? ProjectName;
        public readonly int Revision;
        public readonly bool SchedulePresent;
        public readonly ImmutableArray<Outputs.GetDLPEDMSchemaScheduleResult> Schedules;
        public readonly bool SchemaActive;
        public readonly int SchemaId;
        public readonly ImmutableArray<Outputs.GetDLPEDMSchemaTokenListResult> TokenLists;

        [OutputConstructor]
        private GetDLPEDMSchemaResult(
            int cellsUsed,

            ImmutableArray<Outputs.GetDLPEDMSchemaCreatedByResult> createdBies,

            ImmutableArray<Outputs.GetDLPEDMSchemaEdmClientResult> edmClients,

            string fileName,

            string fileUploadStatus,

            string id,

            ImmutableArray<Outputs.GetDLPEDMSchemaLastModifiedByResult> lastModifiedBies,

            int lastModifiedTime,

            int origColCount,

            string originalFileName,

            string? projectName,

            int revision,

            bool schedulePresent,

            ImmutableArray<Outputs.GetDLPEDMSchemaScheduleResult> schedules,

            bool schemaActive,

            int schemaId,

            ImmutableArray<Outputs.GetDLPEDMSchemaTokenListResult> tokenLists)
        {
            CellsUsed = cellsUsed;
            CreatedBies = createdBies;
            EdmClients = edmClients;
            FileName = fileName;
            FileUploadStatus = fileUploadStatus;
            Id = id;
            LastModifiedBies = lastModifiedBies;
            LastModifiedTime = lastModifiedTime;
            OrigColCount = origColCount;
            OriginalFileName = originalFileName;
            ProjectName = projectName;
            Revision = revision;
            SchedulePresent = schedulePresent;
            Schedules = schedules;
            SchemaActive = schemaActive;
            SchemaId = schemaId;
            TokenLists = tokenLists;
        }
    }
}
