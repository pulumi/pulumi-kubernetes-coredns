import * as k8s from "@pulumi/kubernetes";
import * as coredns from "@pulumi/kubernetes-coredns";

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
});

export const corednsStatus = dns.status;
