// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
    public sealed class GetAdminUsersAdminScopeScopeEntityResult
    {
        public readonly ImmutableDictionary<string, string> Extensions;
        /// <summary>
        /// The ID of the admin user to be exported.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// (String)
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetAdminUsersAdminScopeScopeEntityResult(
            ImmutableDictionary<string, string> extensions,

            int id,

            string name)
        {
            Extensions = extensions;
            Id = id;
            Name = name;
        }
    }
}
