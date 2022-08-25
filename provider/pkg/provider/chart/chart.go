// Copyright 2022, Pulumi Corporation.
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

// A middleware for creating typed helm components
package chart

import (
	"encoding/json"
	"fmt"
	"sync"

	helmbase "github.com/pulumi/pulumi-go-helmbase"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	pp "github.com/pulumi/pulumi/sdk/v3/go/pulumi/provider"
)

func New[C infer.ComponentResource[I, O], I helmbase.ChartArgs, O helmbase.Chart]() (p.Provider, *schema.Provider) {
	var cachedSchema string
	var schemaMutex sync.Mutex
	inferredComponent := infer.Component[C, I, O]()
	schema := schema.Wrap(nil).
		WithResources(inferredComponent)
	return &middleware.Scaffold{
		ConstructFn: func(pctx p.Context, typ string, name string,
			ctx *pulumi.Context, inputs pp.ConstructInputs, opts pulumi.ResourceOption) (pulumi.ComponentResource, error) {
			var i I
			var o O
			state := o
			_, err := helmbase.Construct(ctx, state, typ, name, i, inputs, opts)
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
				pkgName := ctx.RuntimeInformation().PackageName
				tk, err := inferredComponent.GetToken()
				token := fmt.Sprintf("%s:index:%s", pkgName, tk.Name().String())
				if err != nil {
					return p.GetSchemaResponse{}, err
				}
				var spec pschema.PackageSpec
				err = json.Unmarshal([]byte(s.Schema), &spec)
				if err != nil {
					return p.GetSchemaResponse{}, err
				}

				// Fix the type ref
				r := spec.Resources[token]
				r.InputProperties["helmOptions"] = pschema.PropertySpec{
					Description: "HelmOptions is an escape hatch that lets the end user control any aspect of the Helm deployment. This exposes the entirety of the underlying Helm Release component args.",
					TypeSpec: pschema.TypeSpec{
						Ref: fmt.Sprintf("#/types/%s:index:Release", pkgName),
					},
				}
				spec.Resources[token] = r
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
	}, schema
}
