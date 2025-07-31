/*
 * Copyright 2025 The Go-Spring Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-spring/gs-mock/gsmock"
)

// stdOut is the standard output writer, just for test.
var stdOut io.Writer = os.Stdout

// ToolVersion is the version of the mock generation tool.
const ToolVersion = "v0.0.3"

// flagVar holds the command-line flag values.
var flagVar struct {
	outputFile     string // Path to the output Go file for generated mocks.
	mockInterfaces string // Comma-separated list of interface names to mock.
}

func init() {
	flag.StringVar(&flagVar.outputFile, "o", "", "Path to the output file. If not specified, the generated mock code is printed to stdout.")
	flag.StringVar(&flagVar.outputFile, "output", "", "Same as -o. Path to the output file for generated mocks. Defaults to stdout if omitted.")
	flag.StringVar(&flagVar.mockInterfaces, "i", "", "Comma-separated list of interface names to mock (e.g., 'Reader,Writer'). Prefix with '!' to exclude (e.g., '!Logger'). If empty, all interfaces are mocked.")
	flag.StringVar(&flagVar.mockInterfaces, "interfaces", "", "Same as -i. Specifies which interfaces to mock. Use '!' prefix for exclusions. Defaults to all interfaces.")
}

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Println("A tool used to generate Go mock code.")
		fmt.Println(ToolVersion)
		return
	}
	flag.Parse()
	run(runParam{
		sourceDir:      ".",
		outputFile:     flagVar.outputFile,
		mockInterfaces: flagVar.mockInterfaces,
	})
}

// runParam holds configuration parameters for running the generator.
type runParam struct {
	sourceDir      string
	outputFile     string
	mockInterfaces string
}

// run runs the main logic for scanning interfaces and generating mocks.
func run(param runParam) {
	ctx := scanContext{
		outputFile:        param.outputFile,
		includeInterfaces: make(map[string]struct{}),
		excludeInterfaces: make(map[string]struct{}),
	}

	// Parse user-defined mock interface filters.
	if s := param.mockInterfaces; len(s) > 0 {
		if s[0] == '\'' || s[0] == '"' {
			param.mockInterfaces = s[1 : len(s)-1]
		}
		ctx.parse(param.mockInterfaces)
	}

	// Map from import path to package name.
	pkgs := make(map[string]string)
	interfaces := scanDir(param.sourceDir, ctx, pkgs)

	// Collect necessary imports.
	imports := make(map[string]string)
	imports["reflect"] = "reflect"
	imports["gsmock"] = "github.com/go-spring/gs-mock/gsmock"
	for _, m := range interfaces {
		for pkgName, pkgPath := range m.Imports {
			imports[pkgName] = pkgPath
		}
	}

	s := bytes.NewBuffer(nil)

	// Generate import section for the output file.
	h := bytes.NewBuffer(nil)
	for pkgName, pkgPath := range imports {
		ss := strings.Split(pkgPath, "/")
		if pkgName == ss[len(ss)-1] {
			h.WriteString(fmt.Sprintf("\t\"%s\"\n", pkgPath))
		} else {
			h.WriteString(fmt.Sprintf("\t%s \"%s\"\n", pkgName, pkgPath))
		}
	}

	var toolCommand string
	if len(param.outputFile) > 0 {
		toolCommand += "-o " + param.outputFile
	}
	if len(param.mockInterfaces) > 0 {
		toolCommand += " -i '" + param.mockInterfaces + "'"
	}

	packageName := interfaces[0].PackageName

	if err := tmplFileHeader.Execute(s, map[string]any{
		"ToolVersion": ToolVersion,
		"ToolCommand": toolCommand,
		"Package":     packageName,
		"Imports":     h.String(),
	}); err != nil {
		panic(err)
	}

	for _, i := range interfaces {
		if err := tmplInterface.Execute(s, i); err != nil {
			panic(err)
		}
		for _, m := range i.Methods {
			tmpl := getTmplMethod(m.ParamCount, m.ResultCount)
			if err := tmpl.Execute(s, map[string]any{
				"i": i,
				"m": m,
			}); err != nil {
				panic(err)
			}
		}
	}

	// Format the generated source code.
	b, err := format.Source(s.Bytes())
	if err != nil {
		panic(err)
	}

	// Output to file or stdout.
	switch {
	case param.outputFile == "":
		if _, err = stdOut.Write(b); err != nil {
			panic(err)
		}
	default:
		outputFile := filepath.Join(param.sourceDir, param.outputFile)
		if err = os.WriteFile(outputFile, b, os.ModePerm); err != nil {
			panic(err)
		}
	}
}

// scanContext holds configuration during interface scanning.
type scanContext struct {
	outputFile        string
	includeInterfaces map[string]struct{}
	excludeInterfaces map[string]struct{}
}

// parse parses interface filters into inclusion and exclusion sets.
func (ctx *scanContext) parse(mockInterfaces string) {
	if len(mockInterfaces) == 0 {
		return
	}
	ss := strings.Split(mockInterfaces, ",")
	for _, s := range ss {
		if len(s) == 0 {
			continue
		}
		if s[0] == '!' {
			ctx.excludeInterfaces[s[1:]] = struct{}{}
		} else {
			ctx.includeInterfaces[s] = struct{}{}
		}
	}
}

// mock determines if a given interface should be mocked.
func (ctx *scanContext) mock(name string) bool {
	if len(ctx.includeInterfaces) > 0 {
		_, ok := ctx.includeInterfaces[name]
		return ok
	}
	_, ok := ctx.excludeInterfaces[name]
	return !ok
}

// Interface describes a mockable interface.
type Interface struct {
	PackageName     string
	Name            string
	TypeParams      string
	TypeParamNames  string
	EmbedInterfaces string
	Methods         []Method
	File            string
	Imports         map[string]string
}

// Method describes a method inside an interface.
type Method struct {
	Name        string
	Params      string
	ParamNames  string
	ParamTypes  string
	ParamCount  int
	ResultTypes string
	ResultCount int
}

// scanDir scans a directory for interfaces to be mocked.
func scanDir(dir string, ctx scanContext, pkgs map[string]string) []Interface {
	entries, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	var ret []Interface
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".go" {
			continue
		}
		if strings.HasSuffix(entry.Name(), "_test.go") {
			continue
		}
		if entry.Name() == ctx.outputFile {
			continue
		}
		arr := scanFile(ctx, filepath.Join(dir, entry.Name()), pkgs)
		ret = append(ret, arr...)
	}
	return ret
}

// scanFile scans a Go file for interface definitions.
func scanFile(ctx scanContext, file string, pkgs map[string]string) []Interface {
	mode := parser.AllErrors
	node, err := parser.ParseFile(token.NewFileSet(), file, nil, mode)
	if err != nil {
		panic(err)
	}

	needImports := make(map[string]string)
	totalImports := make(map[string]string)

	for _, spec := range node.Imports {
		pkgPath := strings.Trim(spec.Path.Value, "\"")

		var pkgName string
		if spec.Name != nil {
			pkgName = spec.Name.Name
		} else {
			ss := strings.Split(pkgPath, "/")
			pkgName = ss[len(ss)-1]
		}

		if v, ok := pkgs[pkgPath]; ok && v != pkgName {
			panic(fmt.Sprintf("import package name conflict: %s, %s", v, pkgName))
		} else {
			pkgs[pkgPath] = pkgName
		}
		totalImports[pkgName] = pkgPath
	}

	putImport := func(pkgNames []string) {
		for _, s := range pkgNames {
			pkgName := s[:len(s)-1]
			if pkgPath, ok := totalImports[pkgName]; ok {
				needImports[pkgName] = pkgPath
			}
		}
	}

	var ret []Interface
	for _, decl := range node.Decls {
		d, ok := decl.(*ast.GenDecl)
		if !ok || d.Tok != token.TYPE {
			continue
		}

		for _, spec := range d.Specs {
			s := spec.(*ast.TypeSpec)
			t, ok := s.Type.(*ast.InterfaceType)
			if !ok || len(t.Methods.List) == 0 {
				continue
			}

			name := s.Name.String()
			if !ctx.mock(name) {
				continue
			}

			var (
				typeParams     []string
				typeParamNames []string
			)
			if s.TypeParams != nil {
				for _, f := range s.TypeParams.List {
					fName := f.Names[0].Name
					typeText, pkgNames := getTypeText(f.Type)
					typeParams = append(typeParams, fName+" "+typeText)
					typeParamNames = append(typeParamNames, fName)
					putImport(pkgNames)
				}
			}

			var embedInterfaces string
			for _, method := range t.Methods.List {
				if len(method.Names) == 0 {
					embedInterfaces += "\t"
					typeText, pkgNames := getTypeText(method.Type)
					embedInterfaces += typeText
					embedInterfaces += "\n"
					putImport(pkgNames)
				}
			}

			var methods []Method
			for _, method := range t.Methods.List {
				if len(method.Names) == 0 {
					continue
				}
				ft := method.Type.(*ast.FuncType)
				methodName := method.Names[0].Name

				paramCount := 0
				resultCount := 0

				var (
					params     []string
					paramNames []string
					paramTypes []string
				)
				if ft.Params != nil {
					for _, param := range ft.Params.List {

						var tempNames []string
						if len(param.Names) == 0 {
							tempNames = append(tempNames, "r"+strconv.Itoa(paramCount))
						} else {
							for _, r := range param.Names {
								tempNames = append(tempNames, r.Name)
							}
						}

						typeText, pkgNames := getTypeText(param.Type)
						for _, paramName := range tempNames {
							if strings.HasPrefix(typeText, "...") {
								paramTypes = append(paramTypes, "[]"+typeText[3:])
							} else {
								paramTypes = append(paramTypes, typeText)
							}
							paramNames = append(paramNames, paramName)
							params = append(params, paramName+" "+typeText)
						}
						putImport(pkgNames)
						paramCount += len(tempNames)
					}
				}

				if paramCount > gsmock.MaxParamCount {
					panic(fmt.Sprintf("have more than %d parameters", gsmock.MaxParamCount))
				}

				var (
					resultTypes []string
				)
				if ft.Results != nil {
					for _, result := range ft.Results.List {

						var tempNames []string
						if len(result.Names) == 0 {
							tempNames = append(tempNames, "r"+strconv.Itoa(resultCount))
						} else {
							for _, r := range result.Names {
								tempNames = append(tempNames, r.Name)
							}
						}

						typeText, pkgNames := getTypeText(result.Type)
						for range tempNames {
							resultTypes = append(resultTypes, typeText)
						}
						putImport(pkgNames)
						resultCount += len(tempNames)
					}
				}

				if resultCount > gsmock.MaxResultCount {
					panic(fmt.Sprintf("have more than %d results", gsmock.MaxResultCount))
				}

				methods = append(methods, Method{
					Name:        methodName,
					Params:      strings.Join(params, ", "),
					ParamNames:  strings.Join(paramNames, ", "),
					ParamTypes:  strings.Join(paramTypes, ", "),
					ResultTypes: strings.Join(resultTypes, ", "),
					ParamCount:  paramCount,
					ResultCount: resultCount,
				})
			}

			ret = append(ret, Interface{
				PackageName:     node.Name.String(),
				Name:            name,
				TypeParams:      strings.Join(typeParams, ", "),
				TypeParamNames:  strings.Join(typeParamNames, ", "),
				EmbedInterfaces: embedInterfaces,
				Methods:         methods,
				File:            file,
				Imports:         needImports,
			})
		}
	}
	return ret
}

var (
	typeTextBuffer  bytes.Buffer
	typeTextFileSet = token.NewFileSet()
	pkgNameSelector = regexp.MustCompile(`([a-zA-Z0-9_]+\.)`)
)

// getTypeText converts an AST type expression to its string form,
// and extracting used package names.
func getTypeText(t ast.Expr) (typeText string, pkgNames []string) {
	typeTextBuffer.Reset()
	_ = printer.Fprint(&typeTextBuffer, typeTextFileSet, t)
	typeText = typeTextBuffer.String()
	pkgNames = pkgNameSelector.FindAllString(typeText, -1)
	return
}
