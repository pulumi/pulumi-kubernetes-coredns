// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.KubernetesCoreDNS.Inputs
{

    public sealed class CoreDNSRBACArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true, create &amp; use RBAC resources
        /// </summary>
        [Input("create")]
        public Input<bool>? Create { get; set; }

        /// <summary>
        /// The name of the ServiceAccount to use. If not set and create is true, a name is generated using the fullname template.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// If true, create and use PodSecurityPolicy
        /// </summary>
        [Input("pspEnable")]
        public Input<bool>? PspEnable { get; set; }

        public CoreDNSRBACArgs()
        {
        }
    }
}
