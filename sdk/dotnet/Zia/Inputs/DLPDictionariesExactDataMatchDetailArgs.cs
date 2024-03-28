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

    public sealed class DLPDictionariesExactDataMatchDetailArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique identifier for the EDM mapping.
        /// </summary>
        [Input("dictionaryEdmMappingId")]
        public Input<int>? DictionaryEdmMappingId { get; set; }

        /// <summary>
        /// The EDM template's primary field.
        /// </summary>
        [Input("primaryField")]
        public Input<int>? PrimaryField { get; set; }

        /// <summary>
        /// The unique identifier for the EDM template (or schema).
        /// </summary>
        [Input("schemaId")]
        public Input<int>? SchemaId { get; set; }

        /// <summary>
        /// The EDM secondary field to match on.
        /// - `"MATCHON_NONE"`
        /// - `"MATCHON_ANY_1"`
        /// - `"MATCHON_ANY_2"`
        /// - `"MATCHON_ANY_3"`
        /// - `"MATCHON_ANY_4"`
        /// - `"MATCHON_ANY_5"`
        /// - `"MATCHON_ANY_6"`
        /// - `"MATCHON_ANY_7"`
        /// - `"MATCHON_ANY_8"`
        /// - `"MATCHON_ANY_9"`
        /// - `"MATCHON_ANY_10"`
        /// - `"MATCHON_ANY_11"`
        /// - `"MATCHON_ANY_12"`
        /// - `"MATCHON_ANY_13"`
        /// - `"MATCHON_ANY_14"`
        /// - `"MATCHON_ANY_15"`
        /// - `"MATCHON_ALL"`
        /// </summary>
        [Input("secondaryFieldMatchOn")]
        public Input<string>? SecondaryFieldMatchOn { get; set; }

        [Input("secondaryFields")]
        private InputList<int>? _secondaryFields;

        /// <summary>
        /// The EDM template's secondary fields.
        /// </summary>
        public InputList<int> SecondaryFields
        {
            get => _secondaryFields ?? (_secondaryFields = new InputList<int>());
            set => _secondaryFields = value;
        }

        public DLPDictionariesExactDataMatchDetailArgs()
        {
        }
        public static new DLPDictionariesExactDataMatchDetailArgs Empty => new DLPDictionariesExactDataMatchDetailArgs();
    }
}
