// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.KubernetesCoreDNS.Inputs
{

    public sealed class CoreDNSServerArgs : global::Pulumi.ResourceArgs
    {
        [Input("plugins")]
        private InputList<Inputs.CoreDNSServerPluginArgs>? _plugins;

        /// <summary>
        /// the plugins to use for this server block.
        /// </summary>
        public InputList<Inputs.CoreDNSServerPluginArgs> Plugins
        {
            get => _plugins ?? (_plugins = new InputList<Inputs.CoreDNSServerPluginArgs>());
            set => _plugins = value;
        }

        /// <summary>
        /// optional, defaults to "" (which equals 53 in CoreDNS).
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("zones")]
        private InputList<Inputs.CoreDNSServerZoneArgs>? _zones;

        /// <summary>
        /// the `zones` block can be left out entirely, defaults to "."
        /// </summary>
        public InputList<Inputs.CoreDNSServerZoneArgs> Zones
        {
            get => _zones ?? (_zones = new InputList<Inputs.CoreDNSServerZoneArgs>());
            set => _zones = value;
        }

        public CoreDNSServerArgs()
        {
        }
        public static new CoreDNSServerArgs Empty => new CoreDNSServerArgs();
    }
}
