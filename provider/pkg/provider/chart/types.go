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

package chart

import (
	"encoding/json"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

const releaseJSON string = `{
            "description": "A Release is an instance of a chart running in a Kubernetes cluster.\nA Chart is a Helm package. It contains all of the resource definitions necessary to run an application, tool, or service inside of a Kubernetes cluster.\nNote - Helm Release is currently in BETA and may change. Use in production environment is discouraged.",
            "properties": {
                "atomic": {
                    "type": "boolean",
                    "description": "If set, installation process purges chart on fail. ` + "`skipAwait`" + ` will be disabled automatically if atomic is used."
                },
                "chart": {
                    "type": "string",
                    "description": "Chart name to be installed. A path may be used."
                },
                "cleanupOnFail": {
                    "type": "boolean",
                    "description": "Allow deletion of new resources created in this upgrade when upgrade fails."
                },
                "createNamespace": {
                    "type": "boolean",
                    "description": "Create the namespace if it does not exist."
                },
                "dependencyUpdate": {
                    "type": "boolean",
                    "description": "Run helm dependency update before installing the chart."
                },
                "description": {
                    "type": "string",
                    "description": "Add a custom description"
                },
                "devel": {
                    "type": "boolean",
                    "description": "Use chart development versions, too. Equivalent to version '\u003e0.0.0-0'. If ` + "`version`" + ` is set, this is ignored."
                },
                "disableCRDHooks": {
                    "type": "boolean",
                    "description": "Prevent CRD hooks from, running, but run other hooks.  See helm install --no-crd-hook"
                },
                "disableOpenapiValidation": {
                    "type": "boolean",
                    "description": "If set, the installation process will not validate rendered templates against the Kubernetes OpenAPI Schema"
                },
                "disableWebhooks": {
                    "type": "boolean",
                    "description": "Prevent hooks from running."
                },
                "forceUpdate": {
                    "type": "boolean",
                    "description": "Force resource update through delete/recreate if needed."
                },
                "keyring": {
                    "type": "string",
                    "description": "Location of public keys used for verification. Used only if ` + "`verify`" + ` is true"
                },
                "lint": {
                    "type": "boolean",
                    "description": "Run helm lint when planning."
                },
                "manifest": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "pulumi.json#/Any"
                    },
                    "description": "The rendered manifests as JSON. Not yet supported."
                },
                "maxHistory": {
                    "type": "integer",
                    "description": "Limit the maximum number of revisions saved per release. Use 0 for no limit."
                },
                "name": {
                    "type": "string",
                    "description": "Release name."
                },
                "namespace": {
                    "type": "string",
                    "description": "Namespace to install the release into."
                },
                "postrender": {
                    "type": "string",
                    "description": "Postrender command to run."
                },
                "recreatePods": {
                    "type": "boolean",
                    "description": "Perform pods restart during upgrade/rollback."
                },
                "renderSubchartNotes": {
                    "type": "boolean",
                    "description": "If set, render subchart notes along with the parent."
                },
                "replace": {
                    "type": "boolean",
                    "description": "Re-use the given name, even if that name is already used. This is unsafe in production"
                },
                "repositoryOpts": {
                    "$ref": "#/types/kubernetes-coredns:index:RepositoryOpts",
                    "description": "Specification defining the Helm chart repository to use."
                },
                "resetValues": {
                    "type": "boolean",
                    "description": "When upgrading, reset the values to the ones built into the chart."
                },
                "resourceNames": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "array",
                        "items": {
                            "type": "string"
                        }
                    },
                    "description": "Names of resources created by the release grouped by \"kind/version\"."
                },
                "reuseValues": {
                    "type": "boolean",
                    "description": "When upgrading, reuse the last release's values and merge in any overrides. If 'resetValues' is specified, this is ignored"
                },
                "skipAwait": {
                    "type": "boolean",
                    "description": "By default, the provider waits until all resources are in a ready state before marking the release as successful. Setting this to true will skip such await logic."
                },
                "skipCrds": {
                    "type": "boolean",
                    "description": "If set, no CRDs will be installed. By default, CRDs are installed if not already present."
                },
                "timeout": {
                    "type": "integer",
                    "description": "Time in seconds to wait for any individual kubernetes operation."
                },
                "valueYamlFiles": {
                    "type": "array",
                    "items": {
                        "$ref": "pulumi.json#/Asset"
                    },
                    "description": "List of assets (raw yaml files). Content is read and merged with values. Not yet supported."
                },
                "values": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "pulumi.json#/Any"
                    },
                    "description": "Custom values set for the release."
                },
                "verify": {
                    "type": "boolean",
                    "description": "Verify the package before installing it."
                },
                "version": {
                    "type": "string",
                    "description": "Specify the exact chart version to install. If this is not specified, the latest version is installed."
                },
                "waitForJobs": {
                    "type": "boolean",
                    "description": "Will wait until all Jobs have been completed before marking the release as successful. This is ignored if ` + "`skipAwait`" + ` is enabled."
                }
            },
            "type": "object"
        }`

const repositoryOptsJSON string = `{
            "description": "Specification defining the Helm chart repository to use.",
            "properties": {
                "caFile": {
                    "type": "string",
                    "description": "The Repository's CA File"
                },
                "certFile": {
                    "type": "string",
                    "description": "The repository's cert file"
                },
                "keyFile": {
                    "type": "string",
                    "description": "The repository's cert key file"
                },
                "password": {
                    "type": "string",
                    "description": "Password for HTTP basic authentication",
                    "secret": true
                },
                "repo": {
                    "type": "string",
                    "description": "Repository where to locate the requested chart. If is a URL the chart is installed without installing the repository."
                },
                "username": {
                    "type": "string",
                    "description": "Username for HTTP basic authentication"
                }
            },
            "type": "object",
            "language": {
                "nodejs": {
                    "requiredOutputs": [
                        "repo",
                        "keyFile",
                        "certFile",
                        "caFile",
                        "username",
                        "password"
                    ]
                }
            }
        }`

const releaseStatusJSON = `{
            "properties": {
                "appVersion": {
                    "type": "string",
                    "description": "The version number of the application being deployed."
                },
                "chart": {
                    "type": "string",
                    "description": "The name of the chart."
                },
                "name": {
                    "type": "string",
                    "description": "Name is the name of the release."
                },
                "namespace": {
                    "type": "string",
                    "description": "Namespace is the kubernetes namespace of the release."
                },
                "revision": {
                    "type": "integer",
                    "description": "Version is an int32 which represents the version of the release."
                },
                "status": {
                    "type": "string",
                    "description": "Status of the release."
                },
                "version": {
                    "type": "string",
                    "description": "A SemVer 2 conformant version string of the chart."
                }
            },
            "type": "object",
            "required": [
                "name",
                "revision",
                "namespace",
                "chart",
                "version",
                "appVersion",
                "status"
            ]
        }`

func specOf(data string) pschema.ComplexTypeSpec {
	var out pschema.ComplexTypeSpec
	err := json.Unmarshal([]byte(data), &out)
	if err != nil {
		panic(err)
	}
	return out
}
