// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.KubernetesCoreDNS.Inputs
{

    public sealed class CoreDNSZoneFileArgs : global::Pulumi.ResourceArgs
    {
        [Input("contents")]
        public Input<string>? Contents { get; set; }

        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("string")]
        public Input<string>? String { get; set; }

        public CoreDNSZoneFileArgs()
        {
        }
        public static new CoreDNSZoneFileArgs Empty => new CoreDNSZoneFileArgs();
    }
}
