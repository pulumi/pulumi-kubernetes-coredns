import pulumi
from pulumi_kubernetes_coredns import CoreDNS, CoreDNSServerArgs, CoreDNSServerZoneArgs, CoreDNSServerPluginArgs

dns = CoreDNS('dns',
    servers=[
        CoreDNSServerArgs(
            zones=[
                CoreDNSServerZoneArgs(
                    zone='hello.world.',
                    scheme='tls://',
                ),
                CoreDNSServerZoneArgs(
                    zone='foo.bar.',
                    scheme='dns://',
                    use_tcp=True,
                ),
            ],
            port=12345,
            plugins=[
                CoreDNSServerPluginArgs(
                    name='kubernetes',
                    parameters='foo bar',
                    config_block='hello world\nfoo bar',
                ),
            ],
        ),
    ],
)

pulumi.export('coredns_status', dns.status)
