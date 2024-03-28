// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zia.Inputs
{

    public sealed class GetSandboxReportStealthArgs : global::Pulumi.InvokeArgs
    {
        [Input("risk", required: true)]
        public string Risk { get; set; } = null!;

        [Input("signature", required: true)]
        public string Signature { get; set; } = null!;

        [Input("signatureSources", required: true)]
        private List<string>? _signatureSources;
        public List<string> SignatureSources
        {
            get => _signatureSources ?? (_signatureSources = new List<string>());
            set => _signatureSources = value;
        }

        public GetSandboxReportStealthArgs()
        {
        }
        public static new GetSandboxReportStealthArgs Empty => new GetSandboxReportStealthArgs();
    }
}
