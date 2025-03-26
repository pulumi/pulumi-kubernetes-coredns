// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetescoredns;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.kubernetes.apps.v1.inputs.RollingUpdateDeploymentArgs;
import com.pulumi.kubernetes.autoscaling.v2beta2.inputs.HorizontalPodAutoscalerSpecArgs;
import com.pulumi.kubernetes.core.v1.inputs.AffinityArgs;
import com.pulumi.kubernetes.core.v1.inputs.ProbeArgs;
import com.pulumi.kubernetes.core.v1.inputs.ResourceRequirementsArgs;
import com.pulumi.kubernetes.core.v1.inputs.TolerationArgs;
import com.pulumi.kubernetes.core.v1.inputs.VolumeArgs;
import com.pulumi.kubernetes.core.v1.inputs.VolumeMountArgs;
import com.pulumi.kubernetes.policy.v1.inputs.PodDisruptionBudgetSpecArgs;
import com.pulumi.kubernetescoredns.inputs.CoreDNSAutoscalerArgs;
import com.pulumi.kubernetescoredns.inputs.CoreDNSDeploymentArgs;
import com.pulumi.kubernetescoredns.inputs.CoreDNSHPAArgs;
import com.pulumi.kubernetescoredns.inputs.CoreDNSImageArgs;
import com.pulumi.kubernetescoredns.inputs.CoreDNSPrometheusArgs;
import com.pulumi.kubernetescoredns.inputs.CoreDNSRBACArgs;
import com.pulumi.kubernetescoredns.inputs.CoreDNSServerArgs;
import com.pulumi.kubernetescoredns.inputs.CoreDNSServiceAccountArgs;
import com.pulumi.kubernetescoredns.inputs.CoreDNSServiceArgs;
import com.pulumi.kubernetescoredns.inputs.CoreDNSZoneFileArgs;
import com.pulumi.kubernetescoredns.inputs.ReleaseArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CoreDNSArgs extends com.pulumi.resources.ResourceArgs {

    public static final CoreDNSArgs Empty = new CoreDNSArgs();

    /**
     * Affinity settings for pod assignment.
     * 
     */
    @Import(name="affinity")
    private @Nullable Output<AffinityArgs> affinity;

    /**
     * @return Affinity settings for pod assignment.
     * 
     */
    public Optional<Output<AffinityArgs>> affinity() {
        return Optional.ofNullable(this.affinity);
    }

    /**
     * Configue a cluster-proportional-autoscaler for coredns. See https://github.com/kubernetes-incubator/cluster-proportional-autoscaler.
     * 
     */
    @Import(name="autoscaler")
    private @Nullable Output<CoreDNSAutoscalerArgs> autoscaler;

    /**
     * @return Configue a cluster-proportional-autoscaler for coredns. See https://github.com/kubernetes-incubator/cluster-proportional-autoscaler.
     * 
     */
    public Optional<Output<CoreDNSAutoscalerArgs>> autoscaler() {
        return Optional.ofNullable(this.autoscaler);
    }

    /**
     * Create HorizontalPodAutoscaler object.
     * 
     */
    @Import(name="autoscaling")
    private @Nullable Output<HorizontalPodAutoscalerSpecArgs> autoscaling;

    /**
     * @return Create HorizontalPodAutoscaler object.
     * 
     */
    public Optional<Output<HorizontalPodAutoscalerSpecArgs>> autoscaling() {
        return Optional.ofNullable(this.autoscaling);
    }

    /**
     * Configure the CoreDNS Deployment.
     * 
     */
    @Import(name="deployment")
    private @Nullable Output<CoreDNSDeploymentArgs> deployment;

    /**
     * @return Configure the CoreDNS Deployment.
     * 
     */
    public Optional<Output<CoreDNSDeploymentArgs>> deployment() {
        return Optional.ofNullable(this.deployment);
    }

    /**
     * Optional array of secrets to mount inside coredns container. Possible usecase: need for secure connection with etcd backend. Optional array of mount points for extraVolumes.
     * 
     */
    @Import(name="extraSecrets")
    private @Nullable Output<List<VolumeMountArgs>> extraSecrets;

    /**
     * @return Optional array of secrets to mount inside coredns container. Possible usecase: need for secure connection with etcd backend. Optional array of mount points for extraVolumes.
     * 
     */
    public Optional<Output<List<VolumeMountArgs>>> extraSecrets() {
        return Optional.ofNullable(this.extraSecrets);
    }

    /**
     * Optional array of mount points for extraVolumes.
     * 
     */
    @Import(name="extraVolumeMounts")
    private @Nullable Output<List<VolumeMountArgs>> extraVolumeMounts;

    /**
     * @return Optional array of mount points for extraVolumes.
     * 
     */
    public Optional<Output<List<VolumeMountArgs>>> extraVolumeMounts() {
        return Optional.ofNullable(this.extraVolumeMounts);
    }

    /**
     * Optional array of extra volumes to create.
     * 
     */
    @Import(name="extraVolumes")
    private @Nullable Output<List<VolumeArgs>> extraVolumes;

    /**
     * @return Optional array of extra volumes to create.
     * 
     */
    public Optional<Output<List<VolumeArgs>>> extraVolumes() {
        return Optional.ofNullable(this.extraVolumes);
    }

    /**
     * HelmOptions is an escape hatch that lets the end user control any aspect of the Helm deployment. This exposes the entirety of the underlying Helm Release component args.
     * 
     */
    @Import(name="helmOptions")
    private @Nullable ReleaseArgs helmOptions;

    /**
     * @return HelmOptions is an escape hatch that lets the end user control any aspect of the Helm deployment. This exposes the entirety of the underlying Helm Release component args.
     * 
     */
    public Optional<ReleaseArgs> helmOptions() {
        return Optional.ofNullable(this.helmOptions);
    }

    /**
     * Alternative configuration for HPA deployment if wanted.
     * 
     */
    @Import(name="hpa")
    private @Nullable Output<CoreDNSHPAArgs> hpa;

    /**
     * @return Alternative configuration for HPA deployment if wanted.
     * 
     */
    public Optional<Output<CoreDNSHPAArgs>> hpa() {
        return Optional.ofNullable(this.hpa);
    }

    /**
     * The image to pull.
     * 
     */
    @Import(name="image")
    private @Nullable Output<CoreDNSImageArgs> image;

    /**
     * @return The image to pull.
     * 
     */
    public Optional<Output<CoreDNSImageArgs>> image() {
        return Optional.ofNullable(this.image);
    }

    /**
     * Specifies whether chart should be deployed as cluster-service or normal k8s app.
     * 
     */
    @Import(name="isClusterService")
    private @Nullable Output<Boolean> isClusterService;

    /**
     * @return Specifies whether chart should be deployed as cluster-service or normal k8s app.
     * 
     */
    public Optional<Output<Boolean>> isClusterService() {
        return Optional.ofNullable(this.isClusterService);
    }

    /**
     * Configure the liveness probe. To use the livenessProbe, the health plugin needs to be enabled in CoreDNS&#39; server config.
     * 
     */
    @Import(name="livenessProbe")
    private @Nullable Output<ProbeArgs> livenessProbe;

    /**
     * @return Configure the liveness probe. To use the livenessProbe, the health plugin needs to be enabled in CoreDNS&#39; server config.
     * 
     */
    public Optional<Output<ProbeArgs>> livenessProbe() {
        return Optional.ofNullable(this.livenessProbe);
    }

    /**
     * Node labels for pod assignment.
     * 
     */
    @Import(name="nodeSelector")
    private @Nullable Output<Map<String,String>> nodeSelector;

    /**
     * @return Node labels for pod assignment.
     * 
     */
    public Optional<Output<Map<String,String>>> nodeSelector() {
        return Optional.ofNullable(this.nodeSelector);
    }

    /**
     * Optional Pod only Annotations.
     * 
     */
    @Import(name="podAnnotations")
    private @Nullable Output<Map<String,String>> podAnnotations;

    /**
     * @return Optional Pod only Annotations.
     * 
     */
    public Optional<Output<Map<String,String>>> podAnnotations() {
        return Optional.ofNullable(this.podAnnotations);
    }

    /**
     * Optional PodDisruptionBudget.
     * 
     */
    @Import(name="podDisruptionBudget")
    private @Nullable Output<PodDisruptionBudgetSpecArgs> podDisruptionBudget;

    /**
     * @return Optional PodDisruptionBudget.
     * 
     */
    public Optional<Output<PodDisruptionBudgetSpecArgs>> podDisruptionBudget() {
        return Optional.ofNullable(this.podDisruptionBudget);
    }

    /**
     * Under heavy load it takes more that standard time to remove Pod endpoint from a cluster. This will delay termination of our pod by `preStopSleep`. To make sure kube-proxy has enough time to catch up.
     * 
     */
    @Import(name="preStopSleep")
    private @Nullable Output<Integer> preStopSleep;

    /**
     * @return Under heavy load it takes more that standard time to remove Pod endpoint from a cluster. This will delay termination of our pod by `preStopSleep`. To make sure kube-proxy has enough time to catch up.
     * 
     */
    public Optional<Output<Integer>> preStopSleep() {
        return Optional.ofNullable(this.preStopSleep);
    }

    /**
     * Optional priority class to be used for the coredns pods. Used for autoscaler if autoscaler.priorityClassName not set.
     * 
     */
    @Import(name="priorityClassName")
    private @Nullable Output<String> priorityClassName;

    /**
     * @return Optional priority class to be used for the coredns pods. Used for autoscaler if autoscaler.priorityClassName not set.
     * 
     */
    public Optional<Output<String>> priorityClassName() {
        return Optional.ofNullable(this.priorityClassName);
    }

    /**
     * Configure Prometheus installation.
     * 
     */
    @Import(name="prometheus")
    private @Nullable Output<CoreDNSPrometheusArgs> prometheus;

    /**
     * @return Configure Prometheus installation.
     * 
     */
    public Optional<Output<CoreDNSPrometheusArgs>> prometheus() {
        return Optional.ofNullable(this.prometheus);
    }

    /**
     * Configure CoreDNS RBAC resources.
     * 
     */
    @Import(name="rbac")
    private @Nullable Output<CoreDNSRBACArgs> rbac;

    /**
     * @return Configure CoreDNS RBAC resources.
     * 
     */
    public Optional<Output<CoreDNSRBACArgs>> rbac() {
        return Optional.ofNullable(this.rbac);
    }

    /**
     * Configure the readiness probe. To use the readinessProbe, the health plugin needs to be enabled in CoreDNS&#39; server config.
     * 
     */
    @Import(name="readinessProbe")
    private @Nullable Output<ProbeArgs> readinessProbe;

    /**
     * @return Configure the readiness probe. To use the readinessProbe, the health plugin needs to be enabled in CoreDNS&#39; server config.
     * 
     */
    public Optional<Output<ProbeArgs>> readinessProbe() {
        return Optional.ofNullable(this.readinessProbe);
    }

    /**
     * Number of replicas.
     * 
     */
    @Import(name="replicaCount")
    private @Nullable Output<Integer> replicaCount;

    /**
     * @return Number of replicas.
     * 
     */
    public Optional<Output<Integer>> replicaCount() {
        return Optional.ofNullable(this.replicaCount);
    }

    /**
     * Container resource limits.
     * 
     */
    @Import(name="resources")
    private @Nullable Output<ResourceRequirementsArgs> resources;

    /**
     * @return Container resource limits.
     * 
     */
    public Optional<Output<ResourceRequirementsArgs>> resources() {
        return Optional.ofNullable(this.resources);
    }

    @Import(name="rollingUpdate")
    private @Nullable Output<RollingUpdateDeploymentArgs> rollingUpdate;

    public Optional<Output<RollingUpdateDeploymentArgs>> rollingUpdate() {
        return Optional.ofNullable(this.rollingUpdate);
    }

    /**
     * Configuration for CoreDNS and plugins. Default zone is what Kubernetes recommends: https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/#coredns-configmap-options
     * 
     */
    @Import(name="servers")
    private @Nullable Output<List<CoreDNSServerArgs>> servers;

    /**
     * @return Configuration for CoreDNS and plugins. Default zone is what Kubernetes recommends: https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/#coredns-configmap-options
     * 
     */
    public Optional<Output<List<CoreDNSServerArgs>>> servers() {
        return Optional.ofNullable(this.servers);
    }

    /**
     * Configure CoreDNS Service parameters.
     * 
     */
    @Import(name="service")
    private @Nullable Output<CoreDNSServiceArgs> service;

    /**
     * @return Configure CoreDNS Service parameters.
     * 
     */
    public Optional<Output<CoreDNSServiceArgs>> service() {
        return Optional.ofNullable(this.service);
    }

    /**
     * Configure CoreDNS Service Account.
     * 
     */
    @Import(name="serviceAccount")
    private @Nullable Output<CoreDNSServiceAccountArgs> serviceAccount;

    /**
     * @return Configure CoreDNS Service Account.
     * 
     */
    public Optional<Output<CoreDNSServiceAccountArgs>> serviceAccount() {
        return Optional.ofNullable(this.serviceAccount);
    }

    /**
     * Kubernetes Service type.
     * 
     */
    @Import(name="serviceType")
    private @Nullable Output<String> serviceType;

    /**
     * @return Kubernetes Service type.
     * 
     */
    public Optional<Output<String>> serviceType() {
        return Optional.ofNullable(this.serviceType);
    }

    /**
     * Optional duration in seconds the pod needs to terminate gracefully.
     * 
     */
    @Import(name="terminationGracePeriodSeconds")
    private @Nullable Output<Integer> terminationGracePeriodSeconds;

    /**
     * @return Optional duration in seconds the pod needs to terminate gracefully.
     * 
     */
    public Optional<Output<Integer>> terminationGracePeriodSeconds() {
        return Optional.ofNullable(this.terminationGracePeriodSeconds);
    }

    /**
     * Tolerations for pod assignment.
     * 
     */
    @Import(name="tolerations")
    private @Nullable Output<List<TolerationArgs>> tolerations;

    /**
     * @return Tolerations for pod assignment.
     * 
     */
    public Optional<Output<List<TolerationArgs>>> tolerations() {
        return Optional.ofNullable(this.tolerations);
    }

    /**
     * Configure custom Zone files.
     * 
     */
    @Import(name="zoneFiles")
    private @Nullable Output<List<CoreDNSZoneFileArgs>> zoneFiles;

    /**
     * @return Configure custom Zone files.
     * 
     */
    public Optional<Output<List<CoreDNSZoneFileArgs>>> zoneFiles() {
        return Optional.ofNullable(this.zoneFiles);
    }

    private CoreDNSArgs() {}

    private CoreDNSArgs(CoreDNSArgs $) {
        this.affinity = $.affinity;
        this.autoscaler = $.autoscaler;
        this.autoscaling = $.autoscaling;
        this.deployment = $.deployment;
        this.extraSecrets = $.extraSecrets;
        this.extraVolumeMounts = $.extraVolumeMounts;
        this.extraVolumes = $.extraVolumes;
        this.helmOptions = $.helmOptions;
        this.hpa = $.hpa;
        this.image = $.image;
        this.isClusterService = $.isClusterService;
        this.livenessProbe = $.livenessProbe;
        this.nodeSelector = $.nodeSelector;
        this.podAnnotations = $.podAnnotations;
        this.podDisruptionBudget = $.podDisruptionBudget;
        this.preStopSleep = $.preStopSleep;
        this.priorityClassName = $.priorityClassName;
        this.prometheus = $.prometheus;
        this.rbac = $.rbac;
        this.readinessProbe = $.readinessProbe;
        this.replicaCount = $.replicaCount;
        this.resources = $.resources;
        this.rollingUpdate = $.rollingUpdate;
        this.servers = $.servers;
        this.service = $.service;
        this.serviceAccount = $.serviceAccount;
        this.serviceType = $.serviceType;
        this.terminationGracePeriodSeconds = $.terminationGracePeriodSeconds;
        this.tolerations = $.tolerations;
        this.zoneFiles = $.zoneFiles;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CoreDNSArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CoreDNSArgs $;

        public Builder() {
            $ = new CoreDNSArgs();
        }

        public Builder(CoreDNSArgs defaults) {
            $ = new CoreDNSArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param affinity Affinity settings for pod assignment.
         * 
         * @return builder
         * 
         */
        public Builder affinity(@Nullable Output<AffinityArgs> affinity) {
            $.affinity = affinity;
            return this;
        }

        /**
         * @param affinity Affinity settings for pod assignment.
         * 
         * @return builder
         * 
         */
        public Builder affinity(AffinityArgs affinity) {
            return affinity(Output.of(affinity));
        }

        /**
         * @param autoscaler Configue a cluster-proportional-autoscaler for coredns. See https://github.com/kubernetes-incubator/cluster-proportional-autoscaler.
         * 
         * @return builder
         * 
         */
        public Builder autoscaler(@Nullable Output<CoreDNSAutoscalerArgs> autoscaler) {
            $.autoscaler = autoscaler;
            return this;
        }

        /**
         * @param autoscaler Configue a cluster-proportional-autoscaler for coredns. See https://github.com/kubernetes-incubator/cluster-proportional-autoscaler.
         * 
         * @return builder
         * 
         */
        public Builder autoscaler(CoreDNSAutoscalerArgs autoscaler) {
            return autoscaler(Output.of(autoscaler));
        }

        /**
         * @param autoscaling Create HorizontalPodAutoscaler object.
         * 
         * @return builder
         * 
         */
        public Builder autoscaling(@Nullable Output<HorizontalPodAutoscalerSpecArgs> autoscaling) {
            $.autoscaling = autoscaling;
            return this;
        }

        /**
         * @param autoscaling Create HorizontalPodAutoscaler object.
         * 
         * @return builder
         * 
         */
        public Builder autoscaling(HorizontalPodAutoscalerSpecArgs autoscaling) {
            return autoscaling(Output.of(autoscaling));
        }

        /**
         * @param deployment Configure the CoreDNS Deployment.
         * 
         * @return builder
         * 
         */
        public Builder deployment(@Nullable Output<CoreDNSDeploymentArgs> deployment) {
            $.deployment = deployment;
            return this;
        }

        /**
         * @param deployment Configure the CoreDNS Deployment.
         * 
         * @return builder
         * 
         */
        public Builder deployment(CoreDNSDeploymentArgs deployment) {
            return deployment(Output.of(deployment));
        }

        /**
         * @param extraSecrets Optional array of secrets to mount inside coredns container. Possible usecase: need for secure connection with etcd backend. Optional array of mount points for extraVolumes.
         * 
         * @return builder
         * 
         */
        public Builder extraSecrets(@Nullable Output<List<VolumeMountArgs>> extraSecrets) {
            $.extraSecrets = extraSecrets;
            return this;
        }

        /**
         * @param extraSecrets Optional array of secrets to mount inside coredns container. Possible usecase: need for secure connection with etcd backend. Optional array of mount points for extraVolumes.
         * 
         * @return builder
         * 
         */
        public Builder extraSecrets(List<VolumeMountArgs> extraSecrets) {
            return extraSecrets(Output.of(extraSecrets));
        }

        /**
         * @param extraSecrets Optional array of secrets to mount inside coredns container. Possible usecase: need for secure connection with etcd backend. Optional array of mount points for extraVolumes.
         * 
         * @return builder
         * 
         */
        public Builder extraSecrets(VolumeMountArgs... extraSecrets) {
            return extraSecrets(List.of(extraSecrets));
        }

        /**
         * @param extraVolumeMounts Optional array of mount points for extraVolumes.
         * 
         * @return builder
         * 
         */
        public Builder extraVolumeMounts(@Nullable Output<List<VolumeMountArgs>> extraVolumeMounts) {
            $.extraVolumeMounts = extraVolumeMounts;
            return this;
        }

        /**
         * @param extraVolumeMounts Optional array of mount points for extraVolumes.
         * 
         * @return builder
         * 
         */
        public Builder extraVolumeMounts(List<VolumeMountArgs> extraVolumeMounts) {
            return extraVolumeMounts(Output.of(extraVolumeMounts));
        }

        /**
         * @param extraVolumeMounts Optional array of mount points for extraVolumes.
         * 
         * @return builder
         * 
         */
        public Builder extraVolumeMounts(VolumeMountArgs... extraVolumeMounts) {
            return extraVolumeMounts(List.of(extraVolumeMounts));
        }

        /**
         * @param extraVolumes Optional array of extra volumes to create.
         * 
         * @return builder
         * 
         */
        public Builder extraVolumes(@Nullable Output<List<VolumeArgs>> extraVolumes) {
            $.extraVolumes = extraVolumes;
            return this;
        }

        /**
         * @param extraVolumes Optional array of extra volumes to create.
         * 
         * @return builder
         * 
         */
        public Builder extraVolumes(List<VolumeArgs> extraVolumes) {
            return extraVolumes(Output.of(extraVolumes));
        }

        /**
         * @param extraVolumes Optional array of extra volumes to create.
         * 
         * @return builder
         * 
         */
        public Builder extraVolumes(VolumeArgs... extraVolumes) {
            return extraVolumes(List.of(extraVolumes));
        }

        /**
         * @param helmOptions HelmOptions is an escape hatch that lets the end user control any aspect of the Helm deployment. This exposes the entirety of the underlying Helm Release component args.
         * 
         * @return builder
         * 
         */
        public Builder helmOptions(@Nullable ReleaseArgs helmOptions) {
            $.helmOptions = helmOptions;
            return this;
        }

        /**
         * @param hpa Alternative configuration for HPA deployment if wanted.
         * 
         * @return builder
         * 
         */
        public Builder hpa(@Nullable Output<CoreDNSHPAArgs> hpa) {
            $.hpa = hpa;
            return this;
        }

        /**
         * @param hpa Alternative configuration for HPA deployment if wanted.
         * 
         * @return builder
         * 
         */
        public Builder hpa(CoreDNSHPAArgs hpa) {
            return hpa(Output.of(hpa));
        }

        /**
         * @param image The image to pull.
         * 
         * @return builder
         * 
         */
        public Builder image(@Nullable Output<CoreDNSImageArgs> image) {
            $.image = image;
            return this;
        }

        /**
         * @param image The image to pull.
         * 
         * @return builder
         * 
         */
        public Builder image(CoreDNSImageArgs image) {
            return image(Output.of(image));
        }

        /**
         * @param isClusterService Specifies whether chart should be deployed as cluster-service or normal k8s app.
         * 
         * @return builder
         * 
         */
        public Builder isClusterService(@Nullable Output<Boolean> isClusterService) {
            $.isClusterService = isClusterService;
            return this;
        }

        /**
         * @param isClusterService Specifies whether chart should be deployed as cluster-service or normal k8s app.
         * 
         * @return builder
         * 
         */
        public Builder isClusterService(Boolean isClusterService) {
            return isClusterService(Output.of(isClusterService));
        }

        /**
         * @param livenessProbe Configure the liveness probe. To use the livenessProbe, the health plugin needs to be enabled in CoreDNS&#39; server config.
         * 
         * @return builder
         * 
         */
        public Builder livenessProbe(@Nullable Output<ProbeArgs> livenessProbe) {
            $.livenessProbe = livenessProbe;
            return this;
        }

        /**
         * @param livenessProbe Configure the liveness probe. To use the livenessProbe, the health plugin needs to be enabled in CoreDNS&#39; server config.
         * 
         * @return builder
         * 
         */
        public Builder livenessProbe(ProbeArgs livenessProbe) {
            return livenessProbe(Output.of(livenessProbe));
        }

        /**
         * @param nodeSelector Node labels for pod assignment.
         * 
         * @return builder
         * 
         */
        public Builder nodeSelector(@Nullable Output<Map<String,String>> nodeSelector) {
            $.nodeSelector = nodeSelector;
            return this;
        }

        /**
         * @param nodeSelector Node labels for pod assignment.
         * 
         * @return builder
         * 
         */
        public Builder nodeSelector(Map<String,String> nodeSelector) {
            return nodeSelector(Output.of(nodeSelector));
        }

        /**
         * @param podAnnotations Optional Pod only Annotations.
         * 
         * @return builder
         * 
         */
        public Builder podAnnotations(@Nullable Output<Map<String,String>> podAnnotations) {
            $.podAnnotations = podAnnotations;
            return this;
        }

        /**
         * @param podAnnotations Optional Pod only Annotations.
         * 
         * @return builder
         * 
         */
        public Builder podAnnotations(Map<String,String> podAnnotations) {
            return podAnnotations(Output.of(podAnnotations));
        }

        /**
         * @param podDisruptionBudget Optional PodDisruptionBudget.
         * 
         * @return builder
         * 
         */
        public Builder podDisruptionBudget(@Nullable Output<PodDisruptionBudgetSpecArgs> podDisruptionBudget) {
            $.podDisruptionBudget = podDisruptionBudget;
            return this;
        }

        /**
         * @param podDisruptionBudget Optional PodDisruptionBudget.
         * 
         * @return builder
         * 
         */
        public Builder podDisruptionBudget(PodDisruptionBudgetSpecArgs podDisruptionBudget) {
            return podDisruptionBudget(Output.of(podDisruptionBudget));
        }

        /**
         * @param preStopSleep Under heavy load it takes more that standard time to remove Pod endpoint from a cluster. This will delay termination of our pod by `preStopSleep`. To make sure kube-proxy has enough time to catch up.
         * 
         * @return builder
         * 
         */
        public Builder preStopSleep(@Nullable Output<Integer> preStopSleep) {
            $.preStopSleep = preStopSleep;
            return this;
        }

        /**
         * @param preStopSleep Under heavy load it takes more that standard time to remove Pod endpoint from a cluster. This will delay termination of our pod by `preStopSleep`. To make sure kube-proxy has enough time to catch up.
         * 
         * @return builder
         * 
         */
        public Builder preStopSleep(Integer preStopSleep) {
            return preStopSleep(Output.of(preStopSleep));
        }

        /**
         * @param priorityClassName Optional priority class to be used for the coredns pods. Used for autoscaler if autoscaler.priorityClassName not set.
         * 
         * @return builder
         * 
         */
        public Builder priorityClassName(@Nullable Output<String> priorityClassName) {
            $.priorityClassName = priorityClassName;
            return this;
        }

        /**
         * @param priorityClassName Optional priority class to be used for the coredns pods. Used for autoscaler if autoscaler.priorityClassName not set.
         * 
         * @return builder
         * 
         */
        public Builder priorityClassName(String priorityClassName) {
            return priorityClassName(Output.of(priorityClassName));
        }

        /**
         * @param prometheus Configure Prometheus installation.
         * 
         * @return builder
         * 
         */
        public Builder prometheus(@Nullable Output<CoreDNSPrometheusArgs> prometheus) {
            $.prometheus = prometheus;
            return this;
        }

        /**
         * @param prometheus Configure Prometheus installation.
         * 
         * @return builder
         * 
         */
        public Builder prometheus(CoreDNSPrometheusArgs prometheus) {
            return prometheus(Output.of(prometheus));
        }

        /**
         * @param rbac Configure CoreDNS RBAC resources.
         * 
         * @return builder
         * 
         */
        public Builder rbac(@Nullable Output<CoreDNSRBACArgs> rbac) {
            $.rbac = rbac;
            return this;
        }

        /**
         * @param rbac Configure CoreDNS RBAC resources.
         * 
         * @return builder
         * 
         */
        public Builder rbac(CoreDNSRBACArgs rbac) {
            return rbac(Output.of(rbac));
        }

        /**
         * @param readinessProbe Configure the readiness probe. To use the readinessProbe, the health plugin needs to be enabled in CoreDNS&#39; server config.
         * 
         * @return builder
         * 
         */
        public Builder readinessProbe(@Nullable Output<ProbeArgs> readinessProbe) {
            $.readinessProbe = readinessProbe;
            return this;
        }

        /**
         * @param readinessProbe Configure the readiness probe. To use the readinessProbe, the health plugin needs to be enabled in CoreDNS&#39; server config.
         * 
         * @return builder
         * 
         */
        public Builder readinessProbe(ProbeArgs readinessProbe) {
            return readinessProbe(Output.of(readinessProbe));
        }

        /**
         * @param replicaCount Number of replicas.
         * 
         * @return builder
         * 
         */
        public Builder replicaCount(@Nullable Output<Integer> replicaCount) {
            $.replicaCount = replicaCount;
            return this;
        }

        /**
         * @param replicaCount Number of replicas.
         * 
         * @return builder
         * 
         */
        public Builder replicaCount(Integer replicaCount) {
            return replicaCount(Output.of(replicaCount));
        }

        /**
         * @param resources Container resource limits.
         * 
         * @return builder
         * 
         */
        public Builder resources(@Nullable Output<ResourceRequirementsArgs> resources) {
            $.resources = resources;
            return this;
        }

        /**
         * @param resources Container resource limits.
         * 
         * @return builder
         * 
         */
        public Builder resources(ResourceRequirementsArgs resources) {
            return resources(Output.of(resources));
        }

        public Builder rollingUpdate(@Nullable Output<RollingUpdateDeploymentArgs> rollingUpdate) {
            $.rollingUpdate = rollingUpdate;
            return this;
        }

        public Builder rollingUpdate(RollingUpdateDeploymentArgs rollingUpdate) {
            return rollingUpdate(Output.of(rollingUpdate));
        }

        /**
         * @param servers Configuration for CoreDNS and plugins. Default zone is what Kubernetes recommends: https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/#coredns-configmap-options
         * 
         * @return builder
         * 
         */
        public Builder servers(@Nullable Output<List<CoreDNSServerArgs>> servers) {
            $.servers = servers;
            return this;
        }

        /**
         * @param servers Configuration for CoreDNS and plugins. Default zone is what Kubernetes recommends: https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/#coredns-configmap-options
         * 
         * @return builder
         * 
         */
        public Builder servers(List<CoreDNSServerArgs> servers) {
            return servers(Output.of(servers));
        }

        /**
         * @param servers Configuration for CoreDNS and plugins. Default zone is what Kubernetes recommends: https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/#coredns-configmap-options
         * 
         * @return builder
         * 
         */
        public Builder servers(CoreDNSServerArgs... servers) {
            return servers(List.of(servers));
        }

        /**
         * @param service Configure CoreDNS Service parameters.
         * 
         * @return builder
         * 
         */
        public Builder service(@Nullable Output<CoreDNSServiceArgs> service) {
            $.service = service;
            return this;
        }

        /**
         * @param service Configure CoreDNS Service parameters.
         * 
         * @return builder
         * 
         */
        public Builder service(CoreDNSServiceArgs service) {
            return service(Output.of(service));
        }

        /**
         * @param serviceAccount Configure CoreDNS Service Account.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccount(@Nullable Output<CoreDNSServiceAccountArgs> serviceAccount) {
            $.serviceAccount = serviceAccount;
            return this;
        }

        /**
         * @param serviceAccount Configure CoreDNS Service Account.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccount(CoreDNSServiceAccountArgs serviceAccount) {
            return serviceAccount(Output.of(serviceAccount));
        }

        /**
         * @param serviceType Kubernetes Service type.
         * 
         * @return builder
         * 
         */
        public Builder serviceType(@Nullable Output<String> serviceType) {
            $.serviceType = serviceType;
            return this;
        }

        /**
         * @param serviceType Kubernetes Service type.
         * 
         * @return builder
         * 
         */
        public Builder serviceType(String serviceType) {
            return serviceType(Output.of(serviceType));
        }

        /**
         * @param terminationGracePeriodSeconds Optional duration in seconds the pod needs to terminate gracefully.
         * 
         * @return builder
         * 
         */
        public Builder terminationGracePeriodSeconds(@Nullable Output<Integer> terminationGracePeriodSeconds) {
            $.terminationGracePeriodSeconds = terminationGracePeriodSeconds;
            return this;
        }

        /**
         * @param terminationGracePeriodSeconds Optional duration in seconds the pod needs to terminate gracefully.
         * 
         * @return builder
         * 
         */
        public Builder terminationGracePeriodSeconds(Integer terminationGracePeriodSeconds) {
            return terminationGracePeriodSeconds(Output.of(terminationGracePeriodSeconds));
        }

        /**
         * @param tolerations Tolerations for pod assignment.
         * 
         * @return builder
         * 
         */
        public Builder tolerations(@Nullable Output<List<TolerationArgs>> tolerations) {
            $.tolerations = tolerations;
            return this;
        }

        /**
         * @param tolerations Tolerations for pod assignment.
         * 
         * @return builder
         * 
         */
        public Builder tolerations(List<TolerationArgs> tolerations) {
            return tolerations(Output.of(tolerations));
        }

        /**
         * @param tolerations Tolerations for pod assignment.
         * 
         * @return builder
         * 
         */
        public Builder tolerations(TolerationArgs... tolerations) {
            return tolerations(List.of(tolerations));
        }

        /**
         * @param zoneFiles Configure custom Zone files.
         * 
         * @return builder
         * 
         */
        public Builder zoneFiles(@Nullable Output<List<CoreDNSZoneFileArgs>> zoneFiles) {
            $.zoneFiles = zoneFiles;
            return this;
        }

        /**
         * @param zoneFiles Configure custom Zone files.
         * 
         * @return builder
         * 
         */
        public Builder zoneFiles(List<CoreDNSZoneFileArgs> zoneFiles) {
            return zoneFiles(Output.of(zoneFiles));
        }

        /**
         * @param zoneFiles Configure custom Zone files.
         * 
         * @return builder
         * 
         */
        public Builder zoneFiles(CoreDNSZoneFileArgs... zoneFiles) {
            return zoneFiles(List.of(zoneFiles));
        }

        public CoreDNSArgs build() {
            return $;
        }
    }

}
