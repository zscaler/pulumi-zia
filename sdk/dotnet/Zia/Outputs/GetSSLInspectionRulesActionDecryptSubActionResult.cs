// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zia.Outputs
{

    [OutputType]
    public sealed class GetSSLInspectionRulesActionDecryptSubActionResult
    {
        /// <summary>
        /// Whether to block SSL traffic when SNI is not present.
        /// </summary>
        public readonly bool BlockSslTrafficWithNoSniEnabled;
        public readonly bool BlockUndecrypt;
        public readonly bool Http2Enabled;
        public readonly string MinClientTlsVersion;
        public readonly string MinServerTlsVersion;
        /// <summary>
        /// Whether to enable OCSP check.
        /// </summary>
        public readonly bool OcspCheck;
        /// <summary>
        /// Action to take on server certificates. Valid values might include `ALLOW`, `BLOCK`, or `PASS_THRU`.
        /// </summary>
        public readonly string ServerCertificates;

        [OutputConstructor]
        private GetSSLInspectionRulesActionDecryptSubActionResult(
            bool blockSslTrafficWithNoSniEnabled,

            bool blockUndecrypt,

            bool http2Enabled,

            string minClientTlsVersion,

            string minServerTlsVersion,

            bool ocspCheck,

            string serverCertificates)
        {
            BlockSslTrafficWithNoSniEnabled = blockSslTrafficWithNoSniEnabled;
            BlockUndecrypt = blockUndecrypt;
            Http2Enabled = http2Enabled;
            MinClientTlsVersion = minClientTlsVersion;
            MinServerTlsVersion = minServerTlsVersion;
            OcspCheck = ocspCheck;
            ServerCertificates = serverCertificates;
        }
    }
}
