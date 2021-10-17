// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

import * as pulumiKubernetes from "@pulumi/kubernetes";

/**
 * Enable fast and flexible in-cluster DNS.
 */
export class CoreDNS extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'kubernetes-coredns:index:CoreDNS';

    /**
     * Returns true if the given object is an instance of CoreDNS.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CoreDNS {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CoreDNS.__pulumiType;
    }

    /**
     * Detailed information about the status of the underlying Helm deployment.
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.ReleaseStatus>;

    /**
     * Create a CoreDNS resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: CoreDNSArgs, opts?: pulumi.ComponentResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["affinity"] = args ? args.affinity : undefined;
            inputs["autoscaler"] = args ? args.autoscaler : undefined;
            inputs["autoscaling"] = args ? args.autoscaling : undefined;
            inputs["deployment"] = args ? args.deployment : undefined;
            inputs["extraSecrets"] = args ? args.extraSecrets : undefined;
            inputs["extraVolumeMounts"] = args ? args.extraVolumeMounts : undefined;
            inputs["extraVolumes"] = args ? args.extraVolumes : undefined;
            inputs["helmOptions"] = args ? args.helmOptions : undefined;
            inputs["hpa"] = args ? args.hpa : undefined;
            inputs["image"] = args ? args.image : undefined;
            inputs["isClusterService"] = args ? args.isClusterService : undefined;
            inputs["livenessProbe"] = args ? args.livenessProbe : undefined;
            inputs["nodeSelector"] = args ? args.nodeSelector : undefined;
            inputs["podAnnotations"] = args ? args.podAnnotations : undefined;
            inputs["podDisruptionBudget"] = args ? args.podDisruptionBudget : undefined;
            inputs["preStopSleep"] = args ? args.preStopSleep : undefined;
            inputs["priorityClassName"] = args ? args.priorityClassName : undefined;
            inputs["prometheus"] = args ? args.prometheus : undefined;
            inputs["rbac"] = args ? args.rbac : undefined;
            inputs["readinessProbe"] = args ? args.readinessProbe : undefined;
            inputs["replicaCount"] = args ? args.replicaCount : undefined;
            inputs["resources"] = args ? args.resources : undefined;
            inputs["rollingUpdate"] = args ? args.rollingUpdate : undefined;
            inputs["servers"] = args ? args.servers : undefined;
            inputs["service"] = args ? args.service : undefined;
            inputs["serviceAccount"] = args ? args.serviceAccount : undefined;
            inputs["serviceType"] = args ? args.serviceType : undefined;
            inputs["terminationGracePeriodSeconds"] = args ? args.terminationGracePeriodSeconds : undefined;
            inputs["tolerations"] = args ? args.tolerations : undefined;
            inputs["zoneFiles"] = args ? args.zoneFiles : undefined;
            inputs["status"] = undefined /*out*/;
        } else {
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(CoreDNS.__pulumiType, name, inputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a CoreDNS resource.
 */
export interface CoreDNSArgs {
    /**
     * Affinity settings for pod assignment	.
     */
    affinity?: pulumi.Input<pulumiKubernetes.types.input.core.v1.Affinity>;
    /**
     * Configue a cluster-proportional-autoscaler for coredns. See https://github.com/kubernetes-incubator/cluster-proportional-autoscaler.
     */
    autoscaler?: pulumi.Input<inputs.CoreDNSAutoscalerArgs>;
    /**
     * Create HorizontalPodAutoscaler object.
     */
    autoscaling?: pulumi.Input<pulumiKubernetes.types.input.autoscaling.v2beta2.HorizontalPodAutoscalerSpec>;
    /**
     * Configure the CoreDNS Deployment.
     */
    deployment?: pulumi.Input<inputs.CoreDNSDeploymentArgs>;
    /**
     * Optional array of secrets to mount inside coredns container. Possible usecase: need for secure connection with etcd backend. Optional array of mount points for extraVolumes.
     */
    extraSecrets?: pulumi.Input<pulumi.Input<pulumiKubernetes.types.input.core.v1.VolumeMount>[]>;
    /**
     * Optional array of mount points for extraVolumes.
     */
    extraVolumeMounts?: pulumi.Input<pulumi.Input<pulumiKubernetes.types.input.core.v1.VolumeMount>[]>;
    /**
     * Optional array of extra volumes to create.
     */
    extraVolumes?: pulumi.Input<pulumi.Input<pulumiKubernetes.types.input.core.v1.Volume>[]>;
    /**
     * HelmOptions is an escape hatch that lets the end user control any aspect of the Helm deployment. This exposes the entirety of the underlying Helm Release component args.
     */
    helmOptions?: pulumi.Input<inputs.ReleaseArgs>;
    /**
     * Alternative configuration for HPA deployment if wanted.
     */
    hpa?: pulumi.Input<inputs.CoreDNSHPAArgs>;
    /**
     * The image to pull.
     */
    image?: pulumi.Input<inputs.CoreDNSImageArgs>;
    /**
     * Specifies whether chart should be deployed as cluster-service or normal k8s app.
     */
    isClusterService?: pulumi.Input<boolean>;
    /**
     * Configure the liveness probe. To use the livenessProbe, the health plugin needs to be enabled in CoreDNS' server config.
     */
    livenessProbe?: pulumi.Input<pulumiKubernetes.types.input.core.v1.Probe>;
    /**
     * Node labels for pod assignment.
     */
    nodeSelector?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Optional Pod only Annotations.
     */
    podAnnotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Optional PodDisruptionBudget.
     */
    podDisruptionBudget?: pulumi.Input<pulumiKubernetes.types.input.policy.v1.PodDisruptionBudgetSpec>;
    /**
     * Under heavy load it takes more that standard time to remove Pod endpoint from a cluster. This will delay termination of our pod by `preStopSleep`. To make sure kube-proxy has enough time to catch up.
     */
    preStopSleep?: pulumi.Input<number>;
    /**
     * Optional priority class to be used for the coredns pods. Used for autoscaler if autoscaler.priorityClassName not set.
     */
    priorityClassName?: pulumi.Input<string>;
    /**
     * Configure Prometheus installation.
     */
    prometheus?: pulumi.Input<inputs.CoreDNSPrometheusArgs>;
    /**
     * Configure CoreDNS RBAC resources.
     */
    rbac?: pulumi.Input<inputs.CoreDNSRBACArgs>;
    /**
     * Configure the readiness probe. To use the readinessProbe, the health plugin needs to be enabled in CoreDNS' server config.
     */
    readinessProbe?: pulumi.Input<pulumiKubernetes.types.input.core.v1.Probe>;
    /**
     * Number of replicas.
     */
    replicaCount?: pulumi.Input<number>;
    /**
     * Container resource limits.
     */
    resources?: pulumi.Input<pulumiKubernetes.types.input.core.v1.ResourceRequirements>;
    rollingUpdate?: pulumi.Input<pulumiKubernetes.types.input.apps.v1.RollingUpdateDeployment>;
    /**
     * Configuration for CoreDNS and plugins. Default zone is what Kubernetes recommends: https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/#coredns-configmap-options
     */
    servers?: pulumi.Input<pulumi.Input<inputs.CoreDNSServerArgs>[]>;
    /**
     * Configure CoreDNS Service parameters.
     */
    service?: pulumi.Input<inputs.CoreDNSServiceArgs>;
    /**
     * Configure CoreDNS Service Account.
     */
    serviceAccount?: pulumi.Input<inputs.CoreDNSServiceAccountArgs>;
    /**
     * Kubernetes Service type.
     */
    serviceType?: pulumi.Input<string>;
    /**
     * Optional duration in seconds the pod needs to terminate gracefully.
     */
    terminationGracePeriodSeconds?: pulumi.Input<number>;
    /**
     * Tolerations for pod assignment.
     */
    tolerations?: pulumi.Input<pulumi.Input<pulumiKubernetes.types.input.core.v1.Toleration>[]>;
    /**
     * Configure custom Zone files.
     */
    zoneFiles?: pulumi.Input<pulumi.Input<inputs.CoreDNSZoneFileArgs>[]>;
}