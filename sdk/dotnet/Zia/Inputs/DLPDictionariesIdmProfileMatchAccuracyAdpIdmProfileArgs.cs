// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zia.Inputs
{

    public sealed class DLPDictionariesIdmProfileMatchAccuracyAdpIdmProfileArgs : global::Pulumi.ResourceArgs
    {
        [Input("extensions")]
        private InputMap<string>? _extensions;
        public InputMap<string> Extensions
        {
            get => _extensions ?? (_extensions = new InputMap<string>());
            set => _extensions = value;
        }

        [Input("id")]
        public Input<int>? Id { get; set; }

        public DLPDictionariesIdmProfileMatchAccuracyAdpIdmProfileArgs()
        {
        }
        public static new DLPDictionariesIdmProfileMatchAccuracyAdpIdmProfileArgs Empty => new DLPDictionariesIdmProfileMatchAccuracyAdpIdmProfileArgs();
    }
}
