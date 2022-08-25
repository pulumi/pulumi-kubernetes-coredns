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
	"encoding/json"
	"fmt"
	"sync"

	"github.com/blang/semver"
	helmbase "github.com/pulumi/pulumi-go-helmbase"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi-go-provider/middleware"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/cmdutil"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	pp "github.com/pulumi/pulumi/sdk/v3/go/pulumi/provider"
)

const (
	ProviderName  = "kubernetes-coredns"
	ComponentName = ProviderName + ":index:CoreDNS"
)

func Provider() p.Provider {
	schema := schema.Wrap(nil).
		WithResources(infer.Component[*CoreDNS, Args, *State]()).
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
	var cachedSchema string
	var schemaMutex sync.Mutex
	return &middleware.Scaffold{
		ConstructFn: func(pctx p.Context, typ string, name string,
			ctx *pulumi.Context, inputs pp.ConstructInputs, opts pulumi.ResourceOption) (pulumi.ComponentResource, error) {
			state := &State{}
			_, err := helmbase.Construct(ctx, state, typ, name, &Args{}, inputs, opts)
			if err != nil {
				return nil, err
			}
			return state, nil
		},
		GetSchemaFn: func(ctx p.Context, gsr p.GetSchemaRequest) (p.GetSchemaResponse, error) {
			schemaMutex.Lock()
			defer schemaMutex.Unlock()
			if cachedSchema == "" {
				s, err := schema.GetSchema(ctx, gsr)
				if err != nil {
					return p.GetSchemaResponse{}, err
				}
				var spec pschema.PackageSpec
				err = json.Unmarshal([]byte(s.Schema), &spec)
				if err != nil {
					return p.GetSchemaResponse{}, err
				}
				pkgName := ctx.RuntimeInformation().PackageName

				// Fix the type ref
				r := spec.Resources[fmt.Sprintf("%s:index:CoreDNS", pkgName)]
				r.InputProperties["helmOptions"] = pschema.PropertySpec{
					Description: "HelmOptions is an escape hatch that lets the end user control any aspect of the Helm deployment. This exposes the entirety of the underlying Helm Release component args.",
					TypeSpec: pschema.TypeSpec{
						Ref: fmt.Sprintf("#/types/%s:index:Release", pkgName),
					},
				}
				spec.Resources[fmt.Sprintf("%s:index:CoreDNS", pkgName)] = r
				for k, v := range r.InputProperties {
					var makePlain func(p *pschema.TypeSpec)
					makePlain = func(p *pschema.TypeSpec) {
						if p == nil {
							return
						}
						p.Plain = false
						makePlain(p.AdditionalProperties)
						makePlain(p.Items)
						for i, v := range p.OneOf {
							makePlain(&v)
							p.OneOf[i] = v
						}
					}
					makePlain(&v.TypeSpec)
					r.InputProperties[k] = v
				}

				// Add the release type manually
				spec.Types[pkgName+":index:Release"] = specOf(releaseJSON)
				spec.Types[pkgName+":index:ReleaseStatus"] = specOf(releaseStatusJSON)
				spec.Types[pkgName+":index:RepositoryOpts"] = specOf(repositoryOptsJSON)
				delete(spec.Types, pkgName+":v3:ReleaseStatus")

				bytes, err := json.Marshal(spec)
				if err != nil {
					return p.GetSchemaResponse{}, err
				}
				cachedSchema = string(bytes)
			}
			return p.GetSchemaResponse{
				Schema: cachedSchema,
			}, nil
		},
	}
}

func Schema(version semver.Version) (string, error) {
	s, err := integration.NewServer(ProviderName, version, Provider()).
		GetSchema(p.GetSchemaRequest{})
	return s.Schema, err
}

// Serve launches the gRPC server for the resource provider.
func Serve(version string, schema []byte) {
	if err := provider.ComponentMain(ProviderName, version, schema, Construct); err != nil {
		cmdutil.ExitError(err.Error())
	}
}

// Construct is the RPC call that initiates the creation of a new component resource. It
// creates, registers, and returns the resulting object.
func Construct(ctx *pulumi.Context, typ, name string, inputs pp.ConstructInputs,
	opts pulumi.ResourceOption) (*pp.ConstructResult, error) {
	return helmbase.Construct(ctx, &State{}, typ, name, &Args{}, inputs, opts)
}
