# Pulumi Kubernetes CoreDNS Component

This repo contains the Pulumi CoreDNS component for Kubernetes. CoreDNS is a fast and flexible
DNS server, providing DNS services to your cluster.

This component wraps [the official CoreDNS Helm Chart](https://github.com/coredns/helm),
and offers a Pulumi-friendly and strongly-typed way to manage CoreDNS installations.

For examples of usage, see [the official documentation](https://coredns.io/),
or refer to [the examples](/examples) in this repo.

## To Use

To use this component, first install the Pulumi Package:

Afterwards, import the library and instantiate it within your Pulumi program:

## Configuration

This component supports all of the configuration options of the [official Helm chart](
https://github.com/coredns/helm#configuration), except that these
are strongly typed so you will get IDE support and static error checking.

For complete details, refer to the Pulumi Package details within the Pulumi Registry.
