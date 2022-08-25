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
	"github.com/blang/semver"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	"github.com/pulumi/pulumi-kubernetes-coredns/pkg/provider/chart"
)

const (
	ProviderName  = "kubernetes-coredns"
	ComponentName = ProviderName + ":index:CoreDNS"
)

func Provider() p.Provider {
	chart, schema := chart.New[*CoreDNS, *Args, *State]()
	schema.
		WithModuleMap(map[tokens.ModuleName]tokens.ModuleName{
			"provider": "index",
		}).
		WithDisplayName("Kubernetes CoreDNS").
		WithKeywords([]string{
			"pulumi",
			"kubernetes",
			"coredns",
			"kind/component",
			"category/infrastructure",
		}).
		WithPublisher("Pulumi").
		WithLanguageMap(map[string]any{
			"csharp": map[string]map[string]string{
				"namespaces": {
					"kubernetes-coredns": "KubernetesCoreDNS",
				},
				"packageReferences": {
					"Pulumi":            "3.*",
					"Pulumi.Kubernetes": "3.*",
				},
			},
			"go": map[string]any{
				"generateResourceContainerTypes": true,
				"importBasePath":                 "github.com/pulumi/pulumi-kubernetes-coredns/sdk/go/kubernetes-coredns",
			},
			"nodejs": map[string]map[string]string{
				"dependencies": {
					"@pulumi/kubernetes": "^3.7.1",
				},
				"devDependencies": {
					"typescript": "^3.7.0",
				},
			},
			"python": map[string]map[string]string{
				"requires": {
					"pulumi":            ">=3.0.0,<4.0.0",
					"pulumi-kubernetes": ">=3.7.1,<4.0.0",
				},
			},
		})
	return chart
}

func Schema(version semver.Version) (string, error) {
	s, err := integration.NewServer(ProviderName, version, Provider()).
		GetSchema(p.GetSchemaRequest{})
	return s.Schema, err
}
