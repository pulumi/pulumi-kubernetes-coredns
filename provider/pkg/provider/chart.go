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
	infer.ComponentResource[Args, *State]
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
	CoreDNSImage *CoreDNSImage `pulumi:"image"`
	// Number of replicas.
	ReplicaCount *int `pulumi:"replicaCount"`
	// Container resource limits.
	Resources *corev1.ResourceRequirements `pulumi:"resources"
											pschema:"ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:ResourceRequirements"
											provider:"type=kubernetes@3.8.1:core/v1:ResourceRequirements"`
	// Create HorizontalPodAutoscaler object.
	Autoscaling *autoscaling.HorizontalPodAutoscalerSpec `pulumi:"autoscaling"
															pschema:"ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:autoscaling/v2beta2:HorizontalPodAutoscalerSpec"
															provider:"type=kubernetes@3.8.1:autoscaling/v2beta2:HorizontalPodAutoscalerSpec"`
	RollingUpdate *appsv1.RollingUpdateDeployment `pulumi:"rollingUpdate"
												   pschema:"ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:apps/v1:RollingUpdateDeployment"
												   provider:"type=kubernetes@3.8.1:apps/v1:RollingUpdateDeployment"`
	// Under heavy load it takes more that standard time to remove Pod endpoint from a cluster.
	// This will delay termination of our pod by `preStopSleep`. To make sure kube-proxy has
	// enough time to catch up.
	PreStopSleep *int `pulumi:"preStopSleep"`
	// Optional duration in seconds the pod needs to terminate gracefully.
	TerminationGracePeriodSeconds *int `pulumi:"terminationGracePeriodSeconds"`
	// Optional Pod only Annotations.
	PodAnnotations *map[string]string `pulumi:"podAnnotations"`
	// Kubernetes Service type.
	ServiceType *string `pulumi:"serviceType"`
	// Configure Prometheus installation.
	Prometheus *CoreDNSPrometheus `pulumi:"prometheus"`
	// Configure CoreDNS Service parameters.
	Service *CoreDNSService `pulumi:"service"`
	// Configure CoreDNS Service Account.
	ServiceAccount *CoreDNSServiceAccount `pulumi:"serviceAccount"`
	// Configure CoreDNS RBAC resources.
	RBAC *CoreDNSRBAC `pulumi:"rbac"`
	// Specifies whether chart should be deployed as cluster-service or normal k8s app.
	IsClusterService *bool `pulumi:"isClusterService"`
	// Optional priority class to be used for the coredns pods. Used for autoscaler if autoscaler.priorityClassName not set.
	PriorityClassName *string `pulumi:"priorityClassName"`
	// Configuration for CoreDNS and plugins. Default zone is what Kubernetes recommends:
	// https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/#coredns-configmap-options
	CoreDNSServers *[]CoreDNSServer `pulumi:"servers"`
	// Configure the liveness probe. To use the livenessProbe, the health plugin needs to be enabled in CoreDNS' server config.
	LivenessProbe *corev1.Probe `pulumi:"livenessProbe"
								 pschema:"ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:Probe"
								 provider:"typekubernetes@3.8.1:core/v1:Probe"`
	// Configure the readiness probe. To use the readinessProbe, the health plugin needs to be enabled in CoreDNS' server config.
	ReadinessProbe *corev1.Probe `pulumi:"readinessProbe"
								  pschema:"ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:Probe"
								  provider:"typekubernetes@3.8.1:core/v1:Probe"`
	// Affinity settings for pod assignment	.
	Affinity *corev1.Affinity `pulumi:"affinity"
							   pschema:"ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:Affinity"
							   provider:"type=kubernetes@v3.8.1:core/v1:Affinity"`
	// Node labels for pod assignment.
	NodeSelector *map[string]string `pulumi:"nodeSelector"`
	// Tolerations for pod assignment.
	Tolerations *[]corev1.Toleration `pulumi:"tolerations"
									  pschema:"ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:Toleration"
									  provider:"type=kubernetes@3.8.1:core/v1:Toleration"`
	// Optional PodDisruptionBudget.
	PodDisruptionBudget *policyv1.PodDisruptionBudgetSpec `pulumi:"podDisruptionBudget"
														   pschema:"ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:policy/v1:PodDisruptionBudgetSpec"
														   provider:"type=kubernetes@3.8.1:policy/v1:PodDisruptionBudgetSpec"`
	// Configure custom Zone files.
	ZoneFiles *[]CoreDNSZoneFile `pulumi:"zoneFiles"`
	// Optional array of extra volumes to create.
	ExtraVolumes *[]corev1.Volume `pulumi:"extraVolumes"
								   pschema:"ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:Volume"
								   provider:"type=kubernetes@3.8.1:core/v1:Volume"`
	// Optional array of mount points for extraVolumes.
	ExtraVolumeMounts *[]corev1.VolumeMount `pulumi:"extraVolumeMounts"
											 pschema:"ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:VolumeMount"
											 provider:"type=kubernetes@3.8.1:core/v1:VolumeMount"`
	// Optional array of secrets to mount inside coredns container.
	// Possible usecase: need for secure connection with etcd backend.
	// Optional array of mount points for extraVolumes.
	ExtraSecrets *[]corev1.VolumeMount `pulumi:"extraSecrets"
										pschema:"ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:core/v1:VolumeMount"
										provider:"type=kubernetes@3.8.1:core/v1:VolumeMount"`
	// Custom labels to apply to Deployment, Pod, Configmap, Service, ServiceMonitor. Including autoscaler if enabled.
	CustomLabels *map[string]string
	// Alternative configuration for HPA deployment if wanted.
	HPA *CoreDNSHPA `pulumi:"hpa"`
	// Configue a cluster-proportional-autoscaler for coredns.
	// See https://github.com/kubernetes-incubator/cluster-proportional-autoscaler.
	Autoscaler *CoreDNSAutoscaler `pulumi:"autoscaler"`
	// Configure the CoreDNS Deployment.
	Deployment *CoreDNSDeployment `pulumi:"deployment"`

	// HelmOptions is an escape hatch that lets the end user control any aspect of the
	// Helm deployment. This exposes the entirety of the underlying Helm Release component args.
	HelmOptions *helmbase.ReleaseType `pulumi:"helmOptions"
									   pschema:"ref=#/types/chart-coredns:index:Release" json:"-"
                                       provider:"type=index:Release"`
}

func (args *Args) R() **helmbase.ReleaseType { return &args.HelmOptions }

type CoreDNSImage struct {
	// The image repository to pull from.
	Repository *string `pulumi:"repository"`
	// The image tag to pull from.
	Tag *string `pulumi:"tag"`
	// Image pull policy.
	PullPolicy *string `pulumi:"pullPolicy"`
	// Specify container image pull secrets.
	PullSecrets *[]string `pulumi:"pullSecrets"`
}

type CoreDNSPrometheus struct {
	Service *CoreDNSPrometheusService `pulumi:"service"`
	Monitor *CoreDNSPrometheusMonitor `pulumi:"monitor"`
}

type CoreDNSPrometheusService struct {
	// Set this to true to create Service for Prometheus metrics.
	Enabled *bool `pulumi:"enabled"`
	// Annotations to add to the metrics Service.
	Annotations *map[string]string `pulumi:"annotations"`
}

type CoreDNSPrometheusMonitor struct {
	// Set this to true to create ServiceMonitor for Prometheus operator.
	Enabled *bool `pulumi:"enabled"`
	// Additional labels that can be used so ServiceMonitor will be discovered by Prometheus.
	AdditionalLabels *map[string]string `pulumi:"additionalLabels"`
	// Selector to select which namespaces the Endpoints objects are discovered from.
	Namespace *string `pulumi:"namespace"`
}

type CoreDNSService struct {
	// IP address to assign to service.
	ClusterIP *string `pulumi:"clusterIP"`
	// IP address to assign to load balancer (if supported).
	LoadBalancerIP *string `pulumi:"loadBalancerIP"`
	// External IP addresses.
	ExternalIPs *[]string `pulumi:"externalIPs"`
	// Enable client source IP preservation.
	ExternalTrafficPolicy *string `pulumi:"externalTrafficPolicy"`
	// The name of the Service. If not set, a name is generated using the fullname template.
	Name *string `pulumi:"name"`
	// Annotations to add to service.
	Annotations *map[string]string `pulumi:"annotations"`
}

type CoreDNSServiceAccount struct {
	// If true, create & use serviceAccount.
	Create *bool `pulumi:"create"`
	// The name of the ServiceAccount to use.
	// If not set and create is true, a name is generated using the fullname template
	Name        *string            `pulumi:"name"`
	Annotations *map[string]string `pulumi:"annotations"`
}

type CoreDNSRBAC struct {
	// If true, create & use RBAC resources
	Create *bool `pulumi:"create"`
	// If true, create and use PodSecurityPolicy
	PspEnable *bool `pulumi:"pspEnable"`
	// The name of the ServiceAccount to use.
	// If not set and create is true, a name is generated using the fullname template.
	Name *string `pulumi:"name"`
}

type CoreDNSServer struct {
	// the `zones` block can be left out entirely, defaults to "."
	Zones *[]CoreDNSServerZone `pulumi:"zones"`
	// optional, defaults to "" (which equals 53 in CoreDNS).
	Port *int `pulumi:"port"`
	// the plugins to use for this server block.
	Plugins *[]CoreDNSServerPlugin `pulumi:"plugins"`
}

type CoreDNSServerZone struct {
	// optional, defaults to "."
	Zone *string `pulumi:"zone"`
	// optional, defaults to "" (which equals "dns://" in CoreDNS)
	Scheme *string `pulumi:"scheme"`
	// set this parameter to optionally expose the port on tcp as well as udp for the DNS protocol.
	// Note that this will not work if you are also exposing tls or grpc on the same server.
	UseTcp *bool `pulumi:"use_tcp"`
}

type CoreDNSServerPlugin struct {
	// name of plugin, if used multiple times ensure that the plugin supports it!
	Name *string `pulumi:"name"`
	// list of parameters after the plugin
	Parameters *string `pulumi:"parameters"`
	// if the plugin supports extra block style config, supply it here
	ConfigBlock *string `pulumi:"configBlock"`
}

type CoreDNSZoneFile struct {
	Filename *string `pulumi:"string"`
	Domain   *string `pulumi:"domain"`
	Contents *string `pulumi:"contents"`
}

type CoreDNSHPA struct {
	Enabled     *bool                   `pulumi:"enabled"`
	MinReplicas *int                    `pulumi:"minReplicas"`
	MaxReplicas *int                    `pulumi:"maxReplicas"`
	Metrics     *autoscaling.MetricSpec `pulumi:"metrics"
										 pschema:"ref=/kubernetes/v3.8.1/schema.json#/types/kubernetes:autoscaling/v2beta2:MetricSpec"
										 provider:"type=kubernetes@3.8.1:autoscaling/v2beta2:MetricSpec"`
}

type CoreDNSAutoscaler struct {
	// Enabled the cluster-proportional-autoscaler.
	Enabled *bool `pulumi:"enabled"`
	// Number of cores in the cluster per coredns replica.
	CoresPerReplica *int `pulumi:"coresPerReplica"`
	// Number of nodes in the cluster per coredns replica.
	NodesPerReplica *int `pulumi:"nodesPerReplica"`
	// Min size of replicaCount
	Min *int `pulumi:"min"`
	// Max size of replicaCount
	Max *int `pulumi:"max"`
	// Whether to include unschedulable nodes in the nodes/cores calculations - this requires version 1.8.0+ of the autoscaler.
	IncludeUnschedulableNodes *bool `pulumi:"includeUnschedulableNodes"`
	// If true does not allow single points of failure to form.
	PreventSinglePointFailure *bool `pulumi:"preventSinglePointFailure"`
	// The image to pull from for the autoscaler.
	Image *CoreDNSImage `pulumi:"image"`
}

type CoreDNSDeployment struct {
	// Optionally disable the main deployment and its respective resources.
	Enabled *bool `pulumi:"enabled"`
	// Name of the deployment if deployment.enabled is true. Otherwise the name
	// of an existing deployment for the autoscaler or HPA to target.
	Name *string `pulumi:"name"`
}
