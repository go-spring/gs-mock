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
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

var mockerTmpl = template.Must(template.New("mocker").Parse(`
/******************************** {{.mockerName}} ***********************************/

type {{.mockerName}}[{{.req}} any, {{.resp}} any] struct {
	fnHandle func({{.req}}) ({{.resp}}, bool)
	fnWhen   func({{.req}}) bool
	fnReturn func() ({{.resp}})
}

// Handle sets a custom function to handle requests.
func (m *{{.mockerName}}[{{.req}}, {{.resp}}]) Handle(fn func({{.req}}) ({{.resp}}, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *{{.mockerName}}[{{.req}}, {{.resp}}]) When(fn func({{.req}}) bool) *{{.mockerName}}[{{.req}}, {{.resp}}] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *{{.mockerName}}[{{.req}}, {{.resp}}]) Return(fn func() ({{.resp}})) {
	m.fnReturn = fn
}

// {{.invokerName}} is an Invoker implementation for {{.mockerName}}.
type {{.invokerName}}[{{.req}} any, {{.resp}} any] struct {
	*{{.mockerName}}[{{.req}}, {{.resp}}]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *{{.invokerName}}[{{.req}}, {{.resp}}]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *{{.invokerName}}[{{.req}}, {{.resp}}]) Handle(params []interface{}) ([]interface{}, bool) {
	{{.respOnlyArg}}, ok := m.fnHandle({{.cvtParams}})
	return []interface{}{ {{.respOnlyArg}}}, ok
}

// When checks if the condition function evaluates to true.
func (m *{{.invokerName}}[{{.req}}, {{.resp}}]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen({{.cvtParams}})
}

// Return provides predefined response and error values.
func (m *{{.invokerName}}[{{.req}}, {{.resp}}]) Return(params []interface{}) []interface{} {
	{{.respOnlyArg}} := m.fnReturn()
	return []interface{}{ {{.respOnlyArg}}}
}

// New{{.mockerName}} creates a new {{.mockerName}} instance.
func New{{.mockerName}}[{{.req}} any, {{.resp}} any](r *Manager, typ reflect.Type, method string) *{{.mockerName}}[{{.req}}, {{.resp}}] {
	m := &{{.mockerName}}[{{.req}}, {{.resp}}]{}
	i := &{{.invokerName}}[{{.req}}, {{.resp}}]{ {{.mockerName}}: m}
	r.AddMocker(typ, method, i)
	return m
}
`))

// init sets the working directory of the application to the directory
// where this source file resides.
func init() {
	var execDir string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		execDir = filepath.Dir(filename)
	}
	err := os.Chdir(execDir)
	if err != nil {
		panic(err)
	}
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(workDir)
}

func main() {
	var s bytes.Buffer
	s.WriteString(`/*
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

package mock

import (
	"reflect"
)
`)

	const (
		MaxParamCount  = 5
		MaxResultCount = 5
	)

	s.WriteString(fmt.Sprintf(`
	const (
		MaxParamCount  = %d
		MaxResultCount = %d
	)
`, MaxParamCount, MaxResultCount))

	for i := 1; i <= MaxParamCount; i++ {
		for j := 1; j <= MaxResultCount; j++ {
			mockerName := fmt.Sprintf("Mocker%d%d", i, j)
			invokerName := fmt.Sprintf("Invoker%d%d", i, j)
			req := make([]string, i)
			for k := 0; k < i; k++ {
				req[k] = "T" + fmt.Sprint(k+1)
			}
			resp := make([]string, j)
			for k := 0; k < j; k++ {
				resp[k] = "R" + fmt.Sprint(k+1)
			}
			respOnlyArg := make([]string, j)
			for k := 0; k < j; k++ {
				respOnlyArg[k] = "r" + fmt.Sprint(k+1)
			}
			cvtParams := make([]string, i)
			for k := 0; k < i; k++ {
				cvtParams[k] = "params[" + fmt.Sprint(k) + "].(T" + fmt.Sprint(k+1) + ")"
			}
			data := map[string]interface{}{
				"mockerName":  mockerName,
				"invokerName": invokerName,
				"req":         strings.Join(req, ", "),
				"resp":        strings.Join(resp, ", "),
				"respOnlyArg": strings.Join(respOnlyArg, ", "),
				"cvtParams":   strings.Join(cvtParams, ", "),
			}
			err := mockerTmpl.Execute(&s, data)
			if err != nil {
				panic(err)
			}
		}
	}
	b, err := format.Source(s.Bytes())
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("../../mocker.go", b, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
