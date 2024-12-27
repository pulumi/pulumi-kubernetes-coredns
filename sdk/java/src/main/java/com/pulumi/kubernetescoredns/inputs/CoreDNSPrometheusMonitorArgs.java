// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetescoredns.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CoreDNSPrometheusMonitorArgs extends com.pulumi.resources.ResourceArgs {

    public static final CoreDNSPrometheusMonitorArgs Empty = new CoreDNSPrometheusMonitorArgs();

    /**
     * Additional labels that can be used so ServiceMonitor will be discovered by Prometheus.
     * 
     */
    @Import(name="additionalLabels")
    private @Nullable Output<Map<String,String>> additionalLabels;

    /**
     * @return Additional labels that can be used so ServiceMonitor will be discovered by Prometheus.
     * 
     */
    public Optional<Output<Map<String,String>>> additionalLabels() {
        return Optional.ofNullable(this.additionalLabels);
    }

    /**
     * Set this to true to create ServiceMonitor for Prometheus operator.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Set this to true to create ServiceMonitor for Prometheus operator.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * Selector to select which namespaces the Endpoints objects are discovered from.
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return Selector to select which namespaces the Endpoints objects are discovered from.
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    private CoreDNSPrometheusMonitorArgs() {}

    private CoreDNSPrometheusMonitorArgs(CoreDNSPrometheusMonitorArgs $) {
        this.additionalLabels = $.additionalLabels;
        this.enabled = $.enabled;
        this.namespace = $.namespace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CoreDNSPrometheusMonitorArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CoreDNSPrometheusMonitorArgs $;

        public Builder() {
            $ = new CoreDNSPrometheusMonitorArgs();
        }

        public Builder(CoreDNSPrometheusMonitorArgs defaults) {
            $ = new CoreDNSPrometheusMonitorArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param additionalLabels Additional labels that can be used so ServiceMonitor will be discovered by Prometheus.
         * 
         * @return builder
         * 
         */
        public Builder additionalLabels(@Nullable Output<Map<String,String>> additionalLabels) {
            $.additionalLabels = additionalLabels;
            return this;
        }

        /**
         * @param additionalLabels Additional labels that can be used so ServiceMonitor will be discovered by Prometheus.
         * 
         * @return builder
         * 
         */
        public Builder additionalLabels(Map<String,String> additionalLabels) {
            return additionalLabels(Output.of(additionalLabels));
        }

        /**
         * @param enabled Set this to true to create ServiceMonitor for Prometheus operator.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Set this to true to create ServiceMonitor for Prometheus operator.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param namespace Selector to select which namespaces the Endpoints objects are discovered from.
         * 
         * @return builder
         * 
         */
        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param namespace Selector to select which namespaces the Endpoints objects are discovered from.
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        public CoreDNSPrometheusMonitorArgs build() {
            return $;
        }
    }

}