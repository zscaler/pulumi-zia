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

    public sealed class SSLInspectionRulesActionDecryptSubActionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Boolean) - Whether to block SSL traffic when SNI is not present.
        /// </summary>
        [Input("blockSslTrafficWithNoSniEnabled")]
        public Input<bool>? BlockSslTrafficWithNoSniEnabled { get; set; }

        /// <summary>
        /// (Boolean) - Enable to block traffic from servers that use non-standard encryption methods or require mutual TLS authentication.
        /// </summary>
        [Input("blockUndecrypt")]
        public Input<bool>? BlockUndecrypt { get; set; }

        /// <summary>
        /// (Boolean)
        /// </summary>
        [Input("http2Enabled")]
        public Input<bool>? Http2Enabled { get; set; }

        /// <summary>
        /// (String) - The minimum TLS version allowed on the client side: Supported Values are: `CLIENT_TLS_1_0`, `CLIENT_TLS_1_1`, `CLIENT_TLS_1_2`,  `CLIENT_TLS_1_3`.
        /// </summary>
        [Input("minClientTlsVersion")]
        public Input<string>? MinClientTlsVersion { get; set; }

        /// <summary>
        /// (String) - The minimum TLS version allowed on the server side: Supported Values are: `SERVER_TLS_1_0`, `SERVER_TLS_1_1`, `SERVER_TLS_1_2`,  `SERVER_TLS_1_3`.
        /// </summary>
        [Input("minServerTlsVersion")]
        public Input<string>? MinServerTlsVersion { get; set; }

        /// <summary>
        /// (Boolean) - Whether to enable OCSP check.
        /// </summary>
        [Input("ocspCheck")]
        public Input<bool>? OcspCheck { get; set; }

        /// <summary>
        /// (String) - Action to take on server certificates. Valid values might include `ALLOW`, `BLOCK`, or `PASS_THRU`.
        /// </summary>
        [Input("serverCertificates")]
        public Input<string>? ServerCertificates { get; set; }

        public SSLInspectionRulesActionDecryptSubActionGetArgs()
        {
        }
        public static new SSLInspectionRulesActionDecryptSubActionGetArgs Empty => new SSLInspectionRulesActionDecryptSubActionGetArgs();
    }
}
