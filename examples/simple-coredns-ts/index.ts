import * as k8s from "@pulumi/kubernetes";
import * as coredns from "@pulumi/kubernetes-coredns";

// Create a sandbox namespace.
const ns = new k8s.core.v1.Namespace("sandbox-ns");

const dns = new coredns.CoreDNS("dns", {
    servers: [{
        zones: [
            {
                zone: "hello.world.",
                scheme: "tls://",
            },
            {
                zone: "foo.bar.",
                scheme: "dns://",
                use_tcp: true,
            },
        ],
        port: 12345,
        plugins: [
            {
                name: "kubernetes",
                parameters: "foo bar",
                configBlock: "hello world\nfoo bar",
            },
        ],
    }],
    helmOptions: {
        namespace: ns.metadata.name,
    },
});

export const namespace = ns.metadata.name;
export const corednsStatus = dns.status;
