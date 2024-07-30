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
    public static class GetDLPNotificationTemplates
    {
        /// <summary>
        /// Use the **zia_dlp_notification_templates** data source to get information about a ZIA DLP Notification Templates in the Zscaler Internet Access cloud or via the API.
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
        ///     var example = Zia.GetDLPNotificationTemplates.Invoke(new()
        ///     {
        ///         Name = "DLP Auditor Template Test",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetDLPNotificationTemplatesResult> InvokeAsync(GetDLPNotificationTemplatesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDLPNotificationTemplatesResult>("zia:index/getDLPNotificationTemplates:getDLPNotificationTemplates", args ?? new GetDLPNotificationTemplatesArgs(), options.WithDefaults());

        /// <summary>
        /// Use the **zia_dlp_notification_templates** data source to get information about a ZIA DLP Notification Templates in the Zscaler Internet Access cloud or via the API.
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
        ///     var example = Zia.GetDLPNotificationTemplates.Invoke(new()
        ///     {
        ///         Name = "DLP Auditor Template Test",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDLPNotificationTemplatesResult> Invoke(GetDLPNotificationTemplatesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDLPNotificationTemplatesResult>("zia:index/getDLPNotificationTemplates:getDLPNotificationTemplates", args ?? new GetDLPNotificationTemplatesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDLPNotificationTemplatesArgs : global::Pulumi.InvokeArgs
    {
        [Input("id")]
        public int? Id { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        public GetDLPNotificationTemplatesArgs()
        {
        }
        public static new GetDLPNotificationTemplatesArgs Empty => new GetDLPNotificationTemplatesArgs();
    }

    public sealed class GetDLPNotificationTemplatesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id")]
        public Input<int>? Id { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetDLPNotificationTemplatesInvokeArgs()
        {
        }
        public static new GetDLPNotificationTemplatesInvokeArgs Empty => new GetDLPNotificationTemplatesInvokeArgs();
    }


    [OutputType]
    public sealed class GetDLPNotificationTemplatesResult
    {
        public readonly bool AttachContent;
        public readonly string HtmlMessage;
        public readonly int Id;
        public readonly string Name;
        public readonly string PlainTextMessage;
        public readonly string Subject;
        public readonly bool TlsEnabled;

        [OutputConstructor]
        private GetDLPNotificationTemplatesResult(
            bool attachContent,

            string htmlMessage,

            int id,

            string name,

            string plainTextMessage,

            string subject,

            bool tlsEnabled)
        {
            AttachContent = attachContent;
            HtmlMessage = htmlMessage;
            Id = id;
            Name = name;
            PlainTextMessage = plainTextMessage;
            Subject = subject;
            TlsEnabled = tlsEnabled;
        }
    }
}
