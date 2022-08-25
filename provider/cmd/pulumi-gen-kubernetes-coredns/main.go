// Copyright 2016-2021, Pulumi Corporation.
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

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/blang/semver"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-kubernetes-coredns/pkg/provider"
	"github.com/pulumi/pulumi-kubernetes-coredns/pkg/version"
	dotnetgen "github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	nodejsgen "github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	pygen "github.com/pulumi/pulumi/pkg/v3/codegen/python"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <language> <out-dir> \n", os.Args[0])
		os.Exit(1)
	}

	language, outdir := os.Args[1], os.Args[2]

	err := emitSDK(language, outdir)
	if err != nil {
		fmt.Printf("Failed: %s", err.Error())
	}
}

func emitSDK(language, outdir string) error {
	pkg, err := readSchema(version.Version)
	if err != nil {
		return err
	}

	tool := "Pulumi SDK Generator"
	extraFiles := map[string][]byte{}

	var generator func() (map[string][]byte, error)
	switch language {
	case "dotnet":
		generator = func() (map[string][]byte, error) { return dotnetgen.GeneratePackage(tool, pkg, extraFiles) }
	case "go":
		generator = func() (map[string][]byte, error) { return gogen.GeneratePackage(tool, pkg) }
	case "nodejs":
		generator = func() (map[string][]byte, error) { return nodejsgen.GeneratePackage(tool, pkg, extraFiles) }
	case "python":
		generator = func() (map[string][]byte, error) { return pygen.GeneratePackage(tool, pkg, extraFiles) }
	case "schema":
		generator = func() (map[string][]byte, error) {
			schemaBytes, err := provider.Schema(semver.MustParse(version.Version))
			if err != nil {
				return nil, err
			}
			indented := new(bytes.Buffer)
			err = json.Indent(indented, []byte(schemaBytes), "", "    ")
			if err != nil {
				return nil, err
			}
			return map[string][]byte{"schema.json": indented.Bytes()}, nil
		}

	default:
		return errors.Errorf("Unrecognized language %q", language)
	}

	files, err := generator()
	if err != nil {
		return errors.Wrapf(err, "generating %s package", language)
	}

	for f, contents := range files {
		if err := emitFile(outdir, f, contents); err != nil {
			return errors.Wrapf(err, "emitting file %v", f)
		}
	}

	return nil
}

func readSchema(version string) (*schema.Package, error) {
	schemaBytes, err := provider.Schema(semver.MustParse(version))
	if err != nil {
		return nil, errors.Wrap(err, "reading schema")
	}

	var spec schema.PackageSpec
	if err = json.Unmarshal([]byte(schemaBytes), &spec); err != nil {
		return nil, errors.Wrap(err, "unmarshalling schema")
	}

	pkg, err := schema.ImportSpec(spec, nil)
	if err != nil {
		return nil, errors.Wrap(err, "importing schema")
	}
	return pkg, nil
}

func emitFile(rootDir, filename string, contents []byte) error {
	outPath := filepath.Join(rootDir, filename)
	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		return err
	}
	if err := ioutil.WriteFile(outPath, contents, 0600); err != nil {
		return err
	}
	return nil
}
