// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.KubernetesCoreDNS
{
    /// <summary>
    /// Enable fast and flexible in-cluster DNS.
    /// </summary>
    [KubernetesCoreDNSResourceType("kubernetes-coredns:index:CoreDNS")]
    public partial class CoreDNS : global::Pulumi.ComponentResource
    {
        /// <summary>
        /// Detailed information about the status of the underlying Helm deployment.
        /// </summary>
        [Output("status")]
        public Output<Outputs.ReleaseStatus> Status { get; private set; } = null!;


        /// <summary>
        /// Create a CoreDNS resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CoreDNS(string name, CoreDNSArgs? args = null, ComponentResourceOptions? options = null)
            : base("kubernetes-coredns:index:CoreDNS", name, args ?? new CoreDNSArgs(), MakeResourceOptions(options, ""), remote: true)
        {
        }

        private static ComponentResourceOptions MakeResourceOptions(ComponentResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new ComponentResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = ComponentResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class CoreDNSArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Affinity settings for pod assignment.
        /// </summary>
        [Input("affinity")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.AffinityArgs>? Affinity { get; set; }

        /// <summary>
        /// Configue a cluster-proportional-autoscaler for coredns. See https://github.com/kubernetes-incubator/cluster-proportional-autoscaler.
        /// </summary>
        [Input("autoscaler")]
        public Input<Inputs.CoreDNSAutoscalerArgs>? Autoscaler { get; set; }

        /// <summary>
        /// Create HorizontalPodAutoscaler object.
        /// </summary>
        [Input("autoscaling")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2.HorizontalPodAutoscalerSpecArgs>? Autoscaling { get; set; }

        /// <summary>
        /// Configure the CoreDNS Deployment.
        /// </summary>
        [Input("deployment")]
        public Input<Inputs.CoreDNSDeploymentArgs>? Deployment { get; set; }

        [Input("extraSecrets")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs>? _extraSecrets;

        /// <summary>
        /// Optional array of secrets to mount inside coredns container. Possible usecase: need for secure connection with etcd backend. Optional array of mount points for extraVolumes.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs> ExtraSecrets
        {
            get => _extraSecrets ?? (_extraSecrets = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs>());
            set => _extraSecrets = value;
        }

        [Input("extraVolumeMounts")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs>? _extraVolumeMounts;

        /// <summary>
        /// Optional array of mount points for extraVolumes.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs> ExtraVolumeMounts
        {
            get => _extraVolumeMounts ?? (_extraVolumeMounts = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.VolumeMountArgs>());
            set => _extraVolumeMounts = value;
        }

        [Input("extraVolumes")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.VolumeArgs>? _extraVolumes;

        /// <summary>
        /// Optional array of extra volumes to create.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.VolumeArgs> ExtraVolumes
        {
            get => _extraVolumes ?? (_extraVolumes = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.VolumeArgs>());
            set => _extraVolumes = value;
        }

        /// <summary>
        /// HelmOptions is an escape hatch that lets the end user control any aspect of the Helm deployment. This exposes the entirety of the underlying Helm Release component args.
        /// </summary>
        [Input("helmOptions")]
        public Input<Inputs.ReleaseArgs>? HelmOptions { get; set; }

        /// <summary>
        /// Alternative configuration for HPA deployment if wanted.
        /// </summary>
        [Input("hpa")]
        public Input<Inputs.CoreDNSHPAArgs>? Hpa { get; set; }

        /// <summary>
        /// The image to pull.
        /// </summary>
        [Input("image")]
        public Input<Inputs.CoreDNSImageArgs>? Image { get; set; }

        /// <summary>
        /// Specifies whether chart should be deployed as cluster-service or normal k8s app.
        /// </summary>
        [Input("isClusterService")]
        public Input<bool>? IsClusterService { get; set; }

        /// <summary>
        /// Configure the liveness probe. To use the livenessProbe, the health plugin needs to be enabled in CoreDNS' server config.
        /// </summary>
        [Input("livenessProbe")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.ProbeArgs>? LivenessProbe { get; set; }

        [Input("nodeSelector")]
        private InputMap<string>? _nodeSelector;

        /// <summary>
        /// Node labels for pod assignment.
        /// </summary>
        public InputMap<string> NodeSelector
        {
            get => _nodeSelector ?? (_nodeSelector = new InputMap<string>());
            set => _nodeSelector = value;
        }

        [Input("podAnnotations")]
        private InputMap<string>? _podAnnotations;

        /// <summary>
        /// Optional Pod only Annotations.
        /// </summary>
        public InputMap<string> PodAnnotations
        {
            get => _podAnnotations ?? (_podAnnotations = new InputMap<string>());
            set => _podAnnotations = value;
        }

        /// <summary>
        /// Optional PodDisruptionBudget.
        /// </summary>
        [Input("podDisruptionBudget")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Policy.V1.PodDisruptionBudgetSpecArgs>? PodDisruptionBudget { get; set; }

        /// <summary>
        /// Under heavy load it takes more that standard time to remove Pod endpoint from a cluster. This will delay termination of our pod by `preStopSleep`. To make sure kube-proxy has enough time to catch up.
        /// </summary>
        [Input("preStopSleep")]
        public Input<int>? PreStopSleep { get; set; }

        /// <summary>
        /// Optional priority class to be used for the coredns pods. Used for autoscaler if autoscaler.priorityClassName not set.
        /// </summary>
        [Input("priorityClassName")]
        public Input<string>? PriorityClassName { get; set; }

        /// <summary>
        /// Configure Prometheus installation.
        /// </summary>
        [Input("prometheus")]
        public Input<Inputs.CoreDNSPrometheusArgs>? Prometheus { get; set; }

        /// <summary>
        /// Configure CoreDNS RBAC resources.
        /// </summary>
        [Input("rbac")]
        public Input<Inputs.CoreDNSRBACArgs>? Rbac { get; set; }

        /// <summary>
        /// Configure the readiness probe. To use the readinessProbe, the health plugin needs to be enabled in CoreDNS' server config.
        /// </summary>
        [Input("readinessProbe")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.ProbeArgs>? ReadinessProbe { get; set; }

        /// <summary>
        /// Number of replicas.
        /// </summary>
        [Input("replicaCount")]
        public Input<int>? ReplicaCount { get; set; }

        /// <summary>
        /// Container resource limits.
        /// </summary>
        [Input("resources")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.ResourceRequirementsArgs>? Resources { get; set; }

        [Input("rollingUpdate")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Apps.V1.RollingUpdateDeploymentArgs>? RollingUpdate { get; set; }

        [Input("servers")]
        private InputList<Inputs.CoreDNSServerArgs>? _servers;

        /// <summary>
        /// Configuration for CoreDNS and plugins. Default zone is what Kubernetes recommends: https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/#coredns-configmap-options
        /// </summary>
        public InputList<Inputs.CoreDNSServerArgs> Servers
        {
            get => _servers ?? (_servers = new InputList<Inputs.CoreDNSServerArgs>());
            set => _servers = value;
        }

        /// <summary>
        /// Configure CoreDNS Service parameters.
        /// </summary>
        [Input("service")]
        public Input<Inputs.CoreDNSServiceArgs>? Service { get; set; }

        /// <summary>
        /// Configure CoreDNS Service Account.
        /// </summary>
        [Input("serviceAccount")]
        public Input<Inputs.CoreDNSServiceAccountArgs>? ServiceAccount { get; set; }

        /// <summary>
        /// Kubernetes Service type.
        /// </summary>
        [Input("serviceType")]
        public Input<string>? ServiceType { get; set; }

        /// <summary>
        /// Optional duration in seconds the pod needs to terminate gracefully.
        /// </summary>
        [Input("terminationGracePeriodSeconds")]
        public Input<int>? TerminationGracePeriodSeconds { get; set; }

        [Input("tolerations")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.TolerationArgs>? _tolerations;

        /// <summary>
        /// Tolerations for pod assignment.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.TolerationArgs> Tolerations
        {
            get => _tolerations ?? (_tolerations = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.TolerationArgs>());
            set => _tolerations = value;
        }

        [Input("zoneFiles")]
        private InputList<Inputs.CoreDNSZoneFileArgs>? _zoneFiles;

        /// <summary>
        /// Configure custom Zone files.
        /// </summary>
        public InputList<Inputs.CoreDNSZoneFileArgs> ZoneFiles
        {
            get => _zoneFiles ?? (_zoneFiles = new InputList<Inputs.CoreDNSZoneFileArgs>());
            set => _zoneFiles = value;
        }

        public CoreDNSArgs()
        {
        }
        public static new CoreDNSArgs Empty => new CoreDNSArgs();
    }
}
