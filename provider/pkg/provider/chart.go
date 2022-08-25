// Copyright 2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	helmbase "github.com/pulumi/pulumi-go-helmbase"
	"github.com/pulumi/pulumi-go-provider/infer"
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/apps/v1"
	autoscaling "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/autoscaling/v2beta2"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
	policyv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/policy/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CoreDNS struct {
	infer.ComponentResource[*Args, *State]
}

func (c *CoreDNS) Annotate(a infer.Annotator) {
	a.Describe(&c, "Enable fast and flexible in-cluster DNS.")
}

// State installs a fully-configured State stack in Kubernetes.
type State struct {
	pulumi.ResourceState
	Status helmv3.ReleaseStatusOutput `pulumi:"status" pschema:"out" provider:"type=index:ReleaseStatus"`
}

func (s *State) Annotate(a infer.Annotator) {
	a.Describe(&s.Status, "Detailed information about the status of the underlying Helm deployment.")
}

func (c *State) SetOutputs(out helmv3.ReleaseStatusOutput) { c.Status = out }
func (c *State) Type() string                              { return ComponentName }
func (c *State) DefaultChartName() string                  { return "coredns" }
func (c *State) DefaultRepoURL() string                    { return "https://coredns.github.io/helm" }

// Args contains the set of arguments for creating a CoreDNS component resource.
type Args struct {
	// The image to pull.
	CoreDNSImage *CoreDNSImage `pulumi:"image,optional"`
	// Number of replicas.
	ReplicaCount *int `pulumi:"replicaCount,optional"`
	// Container resource limits.
	Resources *corev1.ResourceRequirements `pulumi:"resources,optional" pschema:"ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:ResourceRequirements" provider:"type=kubernetes@3.8.1:core/v1:ResourceRequirements"`
	// Create HorizontalPodAutoscaler object.
	Autoscaling   *autoscaling.HorizontalPodAutoscalerSpec `pulumi:"autoscaling,optional" pschema:"ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:autoscaling/v2beta2:HorizontalPodAutoscalerSpec" provider:"type=kubernetes@3.8.1:autoscaling/v2beta2:HorizontalPodAutoscalerSpec"`
	RollingUpdate *appsv1.RollingUpdateDeployment          `pulumi:"rollingUpdate,optional" pschema:"ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:apps/v1:RollingUpdateDeployment" provider:"type=kubernetes@3.8.1:apps/v1:RollingUpdateDeployment"`
	// Under heavy load it takes more that standard time to remove Pod endpoint from a cluster.
	// This will delay termination of our pod by `preStopSleep`. To make sure kube-proxy has
	// enough time to catch up.
	PreStopSleep *int `pulumi:"preStopSleep,optional"`
	// Optional duration in seconds the pod needs to terminate gracefully.
	TerminationGracePeriodSeconds *int `pulumi:"terminationGracePeriodSeconds,optional"`
	// Optional Pod only Annotations.
	PodAnnotations *map[string]string `pulumi:"podAnnotations,optional"`
	// Kubernetes Service type.
	ServiceType *string `pulumi:"serviceType,optional"`
	// Configure Prometheus installation.
	Prometheus *CoreDNSPrometheus `pulumi:"prometheus,optional"`
	// Configure CoreDNS Service parameters.
	Service *CoreDNSService `pulumi:"service,optional"`
	// Configure CoreDNS Service Account.
	ServiceAccount *CoreDNSServiceAccount `pulumi:"serviceAccount,optional"`
	// Configure CoreDNS RBAC resources.
	RBAC *CoreDNSRBAC `pulumi:"rbac,optional"`
	// Specifies whether chart should be deployed as cluster-service or normal k8s app.
	IsClusterService *bool `pulumi:"isClusterService,optional"`
	// Optional priority class to be used for the coredns pods. Used for autoscaler if autoscaler.priorityClassName not set.
	PriorityClassName *string `pulumi:"priorityClassName,optional"`
	// Configuration for CoreDNS and plugins. Default zone is what Kubernetes recommends:
	// https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/#coredns-configmap-options
	CoreDNSServers *[]CoreDNSServer `pulumi:"servers,optional"`
	// Configure the liveness probe. To use the livenessProbe, the health plugin needs to be enabled in CoreDNS' server config.
	LivenessProbe *corev1.Probe `pulumi:"livenessProbe,optional" pschema:"ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:Probe" provider:"type=kubernetes@3.8.1:core/v1:Probe"`
	// Configure the readiness probe. To use the readinessProbe, the health plugin needs to be enabled in CoreDNS' server config.
	ReadinessProbe *corev1.Probe `pulumi:"readinessProbe,optional" pschema:"ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:Probe" provider:"type=kubernetes@3.8.1:core/v1:Probe"`
	// Affinity settings for pod assignment	.
	Affinity *corev1.Affinity `pulumi:"affinity,optional" pschema:"ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:Affinity" provider:"type=kubernetes@v3.8.1:core/v1:Affinity"`
	// Node labels for pod assignment.
	NodeSelector *map[string]string `pulumi:"nodeSelector,optional"`
	// Tolerations for pod assignment.
	Tolerations *[]corev1.Toleration `pulumi:"tolerations,optional" pschema:"ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:Toleration" provider:"type=kubernetes@3.8.1:core/v1:Toleration"`
	// Optional PodDisruptionBudget.
	PodDisruptionBudget *policyv1.PodDisruptionBudgetSpec `pulumi:"podDisruptionBudget,optional" pschema:"ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:policy/v1:PodDisruptionBudgetSpec" provider:"type=kubernetes@3.8.1:policy/v1:PodDisruptionBudgetSpec"`
	// Configure custom Zone files.
	ZoneFiles *[]CoreDNSZoneFile `pulumi:"zoneFiles,optional"`
	// Optional array of extra volumes to create.
	ExtraVolumes *[]corev1.Volume `pulumi:"extraVolumes,optional" pschema:"ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:Volume" provider:"type=kubernetes@3.8.1:core/v1:Volume"`
	// Optional array of mount points for extraVolumes.
	ExtraVolumeMounts *[]corev1.VolumeMount `pulumi:"extraVolumeMounts,optional" pschema:"ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:VolumeMount" provider:"type=kubernetes@3.8.1:core/v1:VolumeMount"`
	// Optional array of secrets to mount inside coredns container.
	// Possible usecase: need for secure connection with etcd backend.
	// Optional array of mount points for extraVolumes.
	ExtraSecrets *[]corev1.VolumeMount `pulumi:"extraSecrets,optional" pschema:"ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:VolumeMount" provider:"type=kubernetes@3.8.1:core/v1:VolumeMount"`
	// Custom labels to apply to Deployment, Pod, Configmap, Service, ServiceMonitor. Including autoscaler if enabled.
	CustomLabels *map[string]string
	// Alternative configuration for HPA deployment if wanted.
	HPA *CoreDNSHPA `pulumi:"hpa,optional"`
	// Configue a cluster-proportional-autoscaler for coredns.
	// See https://github.com/kubernetes-incubator/cluster-proportional-autoscaler.
	Autoscaler *CoreDNSAutoscaler `pulumi:"autoscaler,optional"`
	// Configure the CoreDNS Deployment.
	Deployment *CoreDNSDeployment `pulumi:"deployment,optional"`

	// HelmOptions is an escape hatch that lets the end user control any aspect of the
	// Helm deployment. This exposes the entirety of the underlying Helm Release component args.
	HelmOptions *helmbase.ReleaseType `pulumi:"helmOptions,optional" pschema:"ref=#/types/chart-coredns:index:Release" json:"-" provider:"type=external@0.0.0:index:Release"`
}

func (args *Args) Annotate(a infer.Annotator) {
	a.Describe(&args.CoreDNSImage, "The image to pull.")
	a.Describe(&args.ReplicaCount, "Number of replicas.")
	a.Describe(&args.Resources, "Container resource limits.")
	a.Describe(&args.Autoscaling, "Create HorizontalPodAutoscaler object.")
	a.Describe(&args.RollingUpdate, "")
	a.Describe(&args.PreStopSleep, "Under heavy load it takes more that standard time to remove Pod endpoint from a cluster. This will delay termination of our pod by `preStopSleep`. To make sure kube-proxy has enough time to catch up.")
	a.Describe(&args.TerminationGracePeriodSeconds, "Optional duration in seconds the pod needs to terminate gracefully.")
	a.Describe(&args.PodAnnotations, "Optional Pod only Annotations.")
	a.Describe(&args.ServiceType, "Kubernetes Service type.")
	a.Describe(&args.Prometheus, "Configure Prometheus installation.")
	a.Describe(&args.Service, "Configure CoreDNS Service parameters.")
	a.Describe(&args.ServiceAccount, "Configure CoreDNS Service Account.")
	a.Describe(&args.RBAC, "Configure CoreDNS RBAC resources.")
	a.Describe(&args.IsClusterService, "Specifies whether chart should be deployed as cluster-service or normal k8s app.")
	a.Describe(&args.PriorityClassName, "Optional priority class to be used for the coredns pods. Used for autoscaler if autoscaler.priorityClassName not set.")
	a.Describe(&args.CoreDNSServers, "Configuration for CoreDNS and plugins. Default zone is what Kubernetes recommends: https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/#coredns-configmap-options")
	a.Describe(&args.LivenessProbe, "Configure the liveness probe. To use the livenessProbe, the health plugin needs to be enabled in CoreDNS' server config.")
	a.Describe(&args.ReadinessProbe, "Configure the readiness probe. To use the readinessProbe, the health plugin needs to be enabled in CoreDNS' server config.")
	a.Describe(&args.Affinity, "Affinity settings for pod assignment\t.")
	a.Describe(&args.NodeSelector, "Node labels for pod assignment.")
	a.Describe(&args.Tolerations, "Tolerations for pod assignment.")
	a.Describe(&args.PodDisruptionBudget, "Optional PodDisruptionBudget.")
	a.Describe(&args.ZoneFiles, "Configure custom Zone files.")
	a.Describe(&args.ExtraVolumes, "Optional array of extra volumes to create.")
	a.Describe(&args.ExtraVolumeMounts, "Optional array of mount points for extraVolumes.")
	a.Describe(&args.ExtraSecrets, "Optional array of secrets to mount inside coredns container. Possible usecase: need for secure connection with etcd backend. Optional array of mount points for extraVolumes.")
	a.Describe(&args.HPA, "Alternative configuration for HPA deployment if wanted.")
	a.Describe(&args.Autoscaler, "Configue a cluster-proportional-autoscaler for coredns. See https://github.com/kubernetes-incubator/cluster-proportional-autoscaler.")
	a.Describe(&args.Deployment, "Configure the CoreDNS Deployment.")
	a.Describe(&args.HelmOptions, "HelmOptions is an escape hatch that lets the end user control any aspect of the Helm deployment. This exposes the entirety of the underlying Helm Release component args.")
}

func (args *Args) R() **helmbase.ReleaseType { return &args.HelmOptions }

type CoreDNSImage struct {
	// The image repository to pull from.
	Repository *string `pulumi:"repository,optional"`
	// The image tag to pull from.
	Tag *string `pulumi:"tag,optional"`
	// Image pull policy.
	PullPolicy *string `pulumi:"pullPolicy,optional"`
	// Specify container image pull secrets.
	PullSecrets *[]string `pulumi:"pullSecrets,optional"`
}

func (c *CoreDNSImage) Annotate(a infer.Annotator) {
	a.Describe(&c.Repository, "The image repository to pull from.")
	a.Describe(&c.Tag, "The image tag to pull from.")
	a.Describe(&c.PullPolicy, "Image pull policy.")
	a.Describe(&c.PullSecrets, "Specify container image pull secrets.")
}

type CoreDNSPrometheus struct {
	Service *CoreDNSPrometheusService `pulumi:"service,optional"`
	Monitor *CoreDNSPrometheusMonitor `pulumi:"monitor,optional"`
}

type CoreDNSPrometheusService struct {
	// Set this to true to create Service for Prometheus metrics.
	Enabled *bool `pulumi:"enabled,optional"`
	// Annotations to add to the metrics Service.
	Annotations *map[string]string `pulumi:"annotations,optional"`
}

func (c *CoreDNSPrometheusService) Annotate(a infer.Annotator) {
	a.Describe(&c.Enabled, "Set this to true to create Service for Prometheus metrics.")
	a.Describe(&c.Annotations, "Annotations to add to the metrics Service.")
}

type CoreDNSPrometheusMonitor struct {
	// Set this to true to create ServiceMonitor for Prometheus operator.
	Enabled *bool `pulumi:"enabled,optional"`
	// Additional labels that can be used so ServiceMonitor will be discovered by Prometheus.
	AdditionalLabels *map[string]string `pulumi:"additionalLabels,optional"`
	// Selector to select which namespaces the Endpoints objects are discovered from.
	Namespace *string `pulumi:"namespace,optional"`
}

func (c *CoreDNSPrometheusMonitor) Annotate(a infer.Annotator) {
	a.Describe(&c.Enabled, "Set this to true to create ServiceMonitor for Prometheus operator.")
	a.Describe(&c.AdditionalLabels, "Additional labels that can be used so ServiceMonitor will be discovered by Prometheus.")
	a.Describe(&c.Namespace, "Selector to select which namespaces the Endpoints objects are discovered from.")
}

type CoreDNSService struct {
	// IP address to assign to service.
	ClusterIP *string `pulumi:"clusterIP,optional"`
	// IP address to assign to load balancer (if supported).
	LoadBalancerIP *string `pulumi:"loadBalancerIP,optional"`
	// External IP addresses.
	ExternalIPs *[]string `pulumi:"externalIPs,optional"`
	// Enable client source IP preservation.
	ExternalTrafficPolicy *string `pulumi:"externalTrafficPolicy,optional"`
	// The name of the Service. If not set, a name is generated using the fullname template.
	Name *string `pulumi:"name,optional"`
	// Annotations to add to service.
	Annotations *map[string]string `pulumi:"annotations,optional"`
}

func (c *CoreDNSService) Annotate(a infer.Annotator) {
	a.Describe(&c.ClusterIP, "IP address to assign to service.")
	a.Describe(&c.LoadBalancerIP, "IP address to assign to load balancer (if supported).")
	a.Describe(&c.ExternalIPs, "External IP addresses.")
	a.Describe(&c.ExternalTrafficPolicy, "Enable client source IP preservation.")
	a.Describe(&c.Name, "The name of the Service. If not set, a name is generated using the fullname template.")
	a.Describe(&c.Annotations, "Annotations to add to service.")
}

type CoreDNSServiceAccount struct {
	// If true, create & use serviceAccount.
	Create *bool `pulumi:"create,optional"`
	// The name of the ServiceAccount to use.
	// If not set and create is true, a name is generated using the fullname template
	Name        *string            `pulumi:"name,optional"`
	Annotations *map[string]string `pulumi:"annotations,optional"`
}

func (c *CoreDNSServiceAccount) Annotate(a infer.Annotator) {
	a.Describe(&c.Create, "If true, create & use serviceAccount.")
	a.Describe(&c.Name, "The name of the ServiceAccount to use. If not set and create is true, a name is generated using the fullname template")
}

type CoreDNSRBAC struct {
	// If true, create & use RBAC resources
	Create *bool `pulumi:"create,optional"`
	// If true, create and use PodSecurityPolicy
	PspEnable *bool `pulumi:"pspEnable,optional"`
	// The name of the ServiceAccount to use.
	// If not set and create is true, a name is generated using the fullname template.
	Name *string `pulumi:"name,optional"`
}

func (c *CoreDNSRBAC) Annotate(a infer.Annotator) {
	a.Describe(&c.Create, "If true, create & use RBAC resources")
	a.Describe(&c.PspEnable, "If true, create and use PodSecurityPolicy")
	a.Describe(&c.Name, "The name of the ServiceAccount to use. If not set and create is true, a name is generated using the fullname template.")
}

type CoreDNSServer struct {
	// the `zones` block can be left out entirely, defaults to "."
	Zones *[]CoreDNSServerZone `pulumi:"zones,optional"`
	// optional, defaults to "" (which equals 53 in CoreDNS).
	Port *int `pulumi:"port,optional"`
	// the plugins to use for this server block.
	Plugins *[]CoreDNSServerPlugin `pulumi:"plugins,optional"`
}

func (c *CoreDNSServer) Annotate(a infer.Annotator) {
	a.Describe(&c.Zones, "the `zones` block can be left out entirely, defaults to \".\"")
	a.Describe(&c.Port, "optional, defaults to \"\" (which equals 53 in CoreDNS).")
	a.Describe(&c.Plugins, "the plugins to use for this server block.")
}

type CoreDNSServerZone struct {
	// optional, defaults to "."
	Zone *string `pulumi:"zone,optional"`
	// optional, defaults to "" (which equals "dns://" in CoreDNS)
	Scheme *string `pulumi:"scheme,optional"`
	// set this parameter to optionally expose the port on tcp as well as udp for the DNS protocol.
	// Note that this will not work if you are also exposing tls or grpc on the same server.
	UseTcp *bool `pulumi:"use_tcp,optional"`
}

func (c *CoreDNSServerZone) Annotate(a infer.Annotator) {
	a.Describe(&c.Zone, "optional, defaults to \".\"")
	a.Describe(&c.UseTcp, "set this parameter to optionally expose the port on tcp as well as udp for the DNS protocol. Note that this will not work if you are also exposing tls or grpc on the same server.")
	a.Describe(&c.Scheme, "optional, defaults to \"\" (which equals \"dns://\" in CoreDNS)")
}

type CoreDNSServerPlugin struct {
	// name of plugin, if used multiple times ensure that the plugin supports it!
	Name *string `pulumi:"name,optional"`
	// list of parameters after the plugin
	Parameters *string `pulumi:"parameters,optional"`
	// if the plugin supports extra block style config, supply it here
	ConfigBlock *string `pulumi:"configBlock,optional"`
}

func (c *CoreDNSServerPlugin) Annotate(a infer.Annotator) {
	a.Describe(&c.Name, "name of plugin, if used multiple times ensure that the plugin supports it!")
	a.Describe(&c.Parameters, "list of parameters after the plugin")
	a.Describe(&c.ConfigBlock, "if the plugin supports extra block style config, supply it here")
}

type CoreDNSZoneFile struct {
	Filename *string `pulumi:"string,optional"`
	Domain   *string `pulumi:"domain,optional"`
	Contents *string `pulumi:"contents,optional"`
}

type CoreDNSHPA struct {
	Enabled     *bool                   `pulumi:"enabled,optional"`
	MinReplicas *int                    `pulumi:"minReplicas,optional"`
	MaxReplicas *int                    `pulumi:"maxReplicas,optional"`
	Metrics     *autoscaling.MetricSpec `pulumi:"metrics,optional" pschema:"ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:autoscaling/v2beta2:MetricSpec" provider:"type=kubernetes@3.8.1:autoscaling/v2beta2:MetricSpec"`
}

type CoreDNSAutoscaler struct {
	// Enabled the cluster-proportional-autoscaler.
	Enabled *bool `pulumi:"enabled,optional"`
	// Number of cores in the cluster per coredns replica.
	CoresPerReplica *int `pulumi:"coresPerReplica,optional"`
	// Number of nodes in the cluster per coredns replica.
	NodesPerReplica *int `pulumi:"nodesPerReplica,optional"`
	// Min size of replicaCount
	Min *int `pulumi:"min,optional"`
	// Max size of replicaCount
	Max *int `pulumi:"max,optional"`
	// Whether to include unschedulable nodes in the nodes/cores calculations - this requires version 1.8.0+ of the autoscaler.
	IncludeUnschedulableNodes *bool `pulumi:"includeUnschedulableNodes,optional"`
	// If true does not allow single points of failure to form.
	PreventSinglePointFailure *bool `pulumi:"preventSinglePointFailure,optional"`
	// The image to pull from for the autoscaler.
	Image *CoreDNSImage `pulumi:"image,optional"`
}

func (c *CoreDNSAutoscaler) Annotate(a infer.Annotator) {
	a.Describe(&c.Enabled, "Enabled the cluster-proportional-autoscaler.")
	a.Describe(&c.CoresPerReplica, "Number of cores in the cluster per coredns replica.")
	a.Describe(&c.NodesPerReplica, "Number of nodes in the cluster per coredns replica.")
	a.Describe(&c.Min, "Min size of replicaCount")
	a.Describe(&c.Max, "Max size of replicaCount")
	a.Describe(&c.IncludeUnschedulableNodes, "Whether to include unschedulable nodes in the nodes/cores calculations - this requires version 1.8.0+ of the autoscaler.")
	a.Describe(&c.PreventSinglePointFailure, "If true does not allow single points of failure to form.")
	a.Describe(&c.Image, "The image to pull from for the autoscaler.")
}

type CoreDNSDeployment struct {
	// Optionally disable the main deployment and its respective resources.
	Enabled *bool `pulumi:"enabled,optional"`
	// Name of the deployment if deployment.enabled is true. Otherwise the name
	// of an existing deployment for the autoscaler or HPA to target.
	Name *string `pulumi:"name,optional"`
}

func (c *CoreDNSDeployment) Annotate(a infer.Annotator) {
	a.Describe(&c.Enabled, "Optionally disable the main deployment and its respective resources.")
	a.Describe(&c.Name, "Name of the deployment if deployment.enabled is true. Otherwise the name of an existing deployment for the autoscaler or HPA to target.")
}
