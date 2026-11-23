// Copyright (c) 2023 Zscaler Technology Alliances, <devrel@zscaler.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package tests

import (
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/property"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// createOnlyResources have no Update method (any change is a replacement).
var createOnlyResources = map[string]bool{
	"zia:index:Activation": true,
}

// TestDryRunGuardIsFirstStatement statically enforces that every resource
// Create and Update method short-circuits on req.DryRun before doing anything
// else. `pulumi preview` calls Create/Update with DryRun=true, and any ZIA API
// call made before the guard (POST/PUT, reorder, activation) mutates the live
// tenant while Pulumi state stays unchanged. See
// https://github.com/zscaler/pulumi-zia/issues/83.
func TestDryRunGuardIsFirstStatement(t *testing.T) {
	t.Parallel()

	files, err := filepath.Glob(filepath.Join("..", "provider", "*.go"))
	require.NoError(t, err)
	require.NotEmpty(t, files)

	fset := token.NewFileSet()
	checked := 0
	for _, f := range files {
		if strings.HasSuffix(f, "_test.go") {
			continue
		}
		af, err := parser.ParseFile(fset, f, nil, 0)
		require.NoError(t, err)
		for _, d := range af.Decls {
			fd, ok := d.(*ast.FuncDecl)
			if !ok || fd.Recv == nil || fd.Body == nil || !isCreateOrUpdate(fd) {
				continue
			}
			checked++
			pos := fset.Position(fd.Pos())
			where := filepath.Base(pos.Filename) + ":" + strconv.Itoa(pos.Line) + " " + fd.Name.Name
			if !assert.NotEmpty(t, fd.Body.List, "%s has an empty body", where) {
				continue
			}
			assert.True(t, isDryRunGuard(fd.Body.List[0]),
				"%s must start with `if req.DryRun { return ... }` so `pulumi preview` never calls the ZIA API", where)
		}
	}
	// 70 Creates + 69 Updates at the time of writing; guard against the scan
	// silently matching nothing if signatures change.
	assert.GreaterOrEqual(t, checked, 139, "expected to find every resource Create/Update method")
}

// TestDryRunMakesNoAPICalls exercises Create and Update with DryRun=true for
// every resource on a provider that was never configured. Any code path that
// reaches the ZIA client fails with "ZIA provider not configured", so success
// proves preview returns before touching the API.
func TestDryRunMakesNoAPICalls(t *testing.T) {
	t.Parallel()

	prov := providerServer(t)
	schemaResp, err := prov.GetSchema(p.GetSchemaRequest{})
	require.NoError(t, err)

	var schema schemaSpec
	require.NoError(t, json.Unmarshal([]byte(schemaResp.Schema), &schema))
	require.NotEmpty(t, schema.Resources)

	tokensList := make([]string, 0, len(schema.Resources))
	for tok := range schema.Resources {
		tokensList = append(tokensList, tok)
	}
	sort.Strings(tokensList)

	for _, tok := range tokensList {
		tok := tok
		res := schema.Resources[tok]
		t.Run(tok, func(t *testing.T) {
			u := resource.NewURN("stack", "proj", "", tokens.Type(tok), "name")
			inputs := schema.minimalObject(res.InputProperties, res.RequiredInputs)

			_, err := prov.Create(p.CreateRequest{Urn: u, Properties: inputs, DryRun: true})
			assert.NoError(t, err, "Create preview must not call the ZIA API")

			if createOnlyResources[tok] {
				return
			}
			state := schema.minimalObject(res.Properties, res.Required)
			_, err = prov.Update(p.UpdateRequest{
				ID: "1", Urn: u, State: state, Inputs: inputs, OldInputs: inputs, DryRun: true,
			})
			assert.NoError(t, err, "Update preview must not call the ZIA API")
		})
	}
}

// schemaSpec is the subset of the Pulumi package schema needed to build
// minimal, decodable property maps for each resource.
type schemaSpec struct {
	Resources map[string]struct {
		InputProperties map[string]propSpec `json:"inputProperties"`
		RequiredInputs  []string            `json:"requiredInputs"`
		Properties      map[string]propSpec `json:"properties"`
		Required        []string            `json:"required"`
	} `json:"resources"`
	Types map[string]struct {
		Properties map[string]propSpec `json:"properties"`
		Required   []string            `json:"required"`
	} `json:"types"`
}

type propSpec struct {
	Type  string    `json:"type"`
	Ref   string    `json:"$ref"`
	Items *propSpec `json:"items"`
}

// minimalObject returns a property map populated with placeholder values for
// the required properties only.
func (s schemaSpec) minimalObject(props map[string]propSpec, required []string) property.Map {
	m := map[string]property.Value{}
	for _, name := range required {
		if spec, ok := props[name]; ok {
			m[name] = s.placeholder(spec)
		}
	}
	return property.NewMap(m)
}

func (s schemaSpec) placeholder(spec propSpec) property.Value {
	if strings.HasPrefix(spec.Ref, "#/types/") {
		typ := s.Types[strings.TrimPrefix(spec.Ref, "#/types/")]
		return property.New(s.minimalObject(typ.Properties, typ.Required))
	}
	switch spec.Type {
	case "string":
		return property.New("placeholder")
	case "integer", "number":
		return property.New(1.0)
	case "boolean":
		return property.New(false)
	case "array":
		return property.New([]property.Value{})
	default: // object / map / any
		return property.New(property.NewMap(map[string]property.Value{}))
	}
}

func isCreateOrUpdate(fd *ast.FuncDecl) bool {
	if fd.Name.Name != "Create" && fd.Name.Name != "Update" {
		return false
	}
	params := fd.Type.Params.List
	if len(params) != 2 {
		return false
	}
	sel := requestTypeSelector(params[1].Type)
	return sel == "CreateRequest" || sel == "UpdateRequest"
}

// requestTypeSelector returns the Sel of infer.CreateRequest[...] / infer.UpdateRequest[...].
func requestTypeSelector(e ast.Expr) string {
	switch t := e.(type) {
	case *ast.IndexExpr:
		e = t.X
	case *ast.IndexListExpr:
		e = t.X
	}
	if s, ok := e.(*ast.SelectorExpr); ok {
		if x, ok := s.X.(*ast.Ident); ok && x.Name == "infer" {
			return s.Sel.Name
		}
	}
	return ""
}

// isDryRunGuard reports whether stmt is `if req.DryRun { ...; return ... }`.
func isDryRunGuard(stmt ast.Stmt) bool {
	ifs, ok := stmt.(*ast.IfStmt)
	if !ok || ifs.Init != nil {
		return false
	}
	sel, ok := ifs.Cond.(*ast.SelectorExpr)
	if !ok || sel.Sel.Name != "DryRun" {
		return false
	}
	if x, ok := sel.X.(*ast.Ident); !ok || x.Name != "req" {
		return false
	}
	body := ifs.Body.List
	if len(body) == 0 {
		return false
	}
	_, ok = body[len(body)-1].(*ast.ReturnStmt)
	return ok
}
