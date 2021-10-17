// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ChartCertManager.Inputs
{

    public sealed class CoreDNSHPAArgs : Pulumi.ResourceArgs
    {
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("maxReplicas")]
        public Input<int>? MaxReplicas { get; set; }

        [Input("metrics")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2.MetricSpecArgs>? Metrics { get; set; }

        [Input("minReplicas")]
        public Input<int>? MinReplicas { get; set; }

        public CoreDNSHPAArgs()
        {
        }
    }
}