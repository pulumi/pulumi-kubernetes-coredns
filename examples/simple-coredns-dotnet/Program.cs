using Pulumi;
using Pulumi.Kubernetes.Types.Inputs.Core.V1;
using Pulumi.Kubernetes.Types.Inputs.Apps.V1;
using Pulumi.Kubernetes.Types.Inputs.Meta.V1;
using Pulumi.Kubernetes.Core.V1;
using Pulumi.KubernetesCoreDNS;
using Pulumi.KubernetesCoreDNS.Inputs;
using System.Collections.Generic;

return await Deployment.RunAsync(() =>
{
  // Create a sandbox namespace
  var ns = new Namespace("sandbox-ns");

  var dns = new CoreDNS("dns", new CoreDNSArgs
  {
    Servers = new[]
      {
            new CoreDNSServerArgs
            {
                Zones = new[]
                {
                    new CoreDNSServerZoneArgs
                    {
                        Zone = "hello.world.",
                        Scheme = "tls://"
                    },
                    new CoreDNSServerZoneArgs
                    {
                        Zone = "foo.bar.",
                        Scheme = "dns://",
                        Use_tcp = true
                    }
                },
                Port = 12345,
                Plugins = new[]
                {
                    new CoreDNSServerPluginArgs
                    {
                        Name = "kubernetes",
                        Parameters = "foo bar",
                        ConfigBlock = "hello world\nfoo bar"
                    }
                }
            }
        },
    HelmOptions = new ReleaseArgs
    {
      Namespace = ns.Metadata.Apply(m => m.Name)
    }
  });

  return new Dictionary<string, object?>
    {
        { "namespace", ns.Metadata.Apply(m => m.Name) },
        { "corednsStatus", dns.Status.Apply(s => s.Status.ToString()) }
    };
});


