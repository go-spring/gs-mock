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
	"text/template"
)

// getMockTemplate returns the appropriate template
// based on the number of return values.
func getMockTemplate(j int) *template.Template {
	if j == 0 {
		return tmplMockN0
	}
	return tmplMockNN
}

// tmplMockN0 for N request parameters, returns nothing.
var tmplMockN0 = template.Must(template.New("MockN0").Parse(`
/******************************** {{.mockerName}} ***********************************/

type {{.mockerName}}[{{.tmplReq}} any] struct {
	fnHandle func({{.req}})
	fnWhen   func({{.req}}) bool
	fnReturn func()
}

// Handle sets a custom function to handle requests.
func (m *{{.mockerName}}[{{.tmplReq}}]) Handle(fn func({{.req}})) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *{{.mockerName}}[{{.tmplReq}}]) When(fn func({{.req}}) bool) *{{.mockerName}}[{{.tmplReq}}] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *{{.mockerName}}[{{.tmplReq}}]) Return(fn func()) {
	if m.fnWhen == nil {
		m.fnWhen = func({{.req}}) bool { return true }
	}
	m.fnReturn = fn
}

// ReturnValue sets a return function with predefined values.
func (m *{{.mockerName}}[{{.tmplReq}}]) ReturnValue() {
	m.Return(func() {})
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *{{.mockerName}}[{{.tmplReq}}]) ReturnDefault() {
	m.Return(func() {})
}

// {{.invokerName}} is an Invoker implementation for {{.mockerName}}.
type {{.invokerName}}[{{.tmplReq}} any] struct {
	*{{.mockerName}}[{{.tmplReq}}]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *{{.invokerName}}[{{.tmplReq}}]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *{{.invokerName}}[{{.tmplReq}}]) Handle(params []any) []any {
	m.fnHandle({{.cvtParams}})
	return []any{}
}

// When checks if the condition function evaluates to true.
func (m *{{.invokerName}}[{{.tmplReq}}]) When(params []any) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen({{.cvtParams}})
}

// Return provides predefined response and error values.
func (m *{{.invokerName}}[{{.tmplReq}}]) Return() []any {
	m.fnReturn()
	return []any{}
}

// {{.factoryName}} creates a new {{.mockerName}} instance.
func {{.factoryName}}[{{.tmplReq}} any](f func({{.funcReq}}), r *Manager) *{{.mockerName}}[{{.tmplReq}}] {
	PatchOnce(f)
	m := &{{.mockerName}}[{{.tmplReq}}]{}
	i := &{{.invokerName}}[{{.tmplReq}}]{ {{.mockerName}}: m}
	r.addMocker(f, i)
	return m
}
`))

// tmplMockNN for N request parameters, returns N values.
var tmplMockNN = template.Must(template.New("MockNN").Parse(`
/******************************** {{.mockerName}} ***********************************/

type {{.mockerName}}[{{.tmplReq}} any, {{.resp}} any] struct {
	fnHandle func({{.req}}) ({{.resp}})
	fnWhen   func({{.req}}) bool
	fnReturn func() ({{.resp}})
}

// Handle sets a custom function to handle requests.
func (m *{{.mockerName}}[{{.tmplReq}}, {{.resp}}]) Handle(fn func({{.req}}) ({{.resp}})) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *{{.mockerName}}[{{.tmplReq}}, {{.resp}}]) When(fn func({{.req}}) bool) *{{.mockerName}}[{{.tmplReq}}, {{.resp}}] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *{{.mockerName}}[{{.tmplReq}}, {{.resp}}]) Return(fn func() ({{.resp}})) {
	if m.fnWhen == nil {
		m.fnWhen = func({{.req}}) bool { return true }
	}
	m.fnReturn = fn
}

// ReturnValue sets a return function with predefined values.
func (m *{{.mockerName}}[{{.tmplReq}}, {{.resp}}]) ReturnValue({{.respWithArg}}) {
	m.Return(func() ({{.resp}}) { return {{.respOnlyArg}} })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *{{.mockerName}}[{{.tmplReq}}, {{.resp}}]) ReturnDefault() {
	m.Return(func() ({{.respWithArg}}) { return {{.respOnlyArg}} })
}

// {{.invokerName}} is an Invoker implementation for {{.mockerName}}.
type {{.invokerName}}[{{.tmplReq}} any, {{.resp}} any] struct {
	*{{.mockerName}}[{{.tmplReq}}, {{.resp}}]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *{{.invokerName}}[{{.tmplReq}}, {{.resp}}]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *{{.invokerName}}[{{.tmplReq}}, {{.resp}}]) Handle(params []any) []any {
	{{.respOnlyArg}} := m.fnHandle({{.cvtParams}})
	return []any{ {{.respOnlyArg}}}
}

// When checks if the condition function evaluates to true.
func (m *{{.invokerName}}[{{.tmplReq}}, {{.resp}}]) When(params []any) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen({{.cvtParams}})
}

// Return provides predefined response and error values.
func (m *{{.invokerName}}[{{.tmplReq}}, {{.resp}}]) Return() []any {
	{{.respOnlyArg}} := m.fnReturn()
	return []any{ {{.respOnlyArg}}}
}

// {{.factoryName}} creates a new {{.mockerName}} instance.
func {{.factoryName}}[{{.tmplReq}} any, {{.resp}} any](f func({{.funcReq}})({{.resp}}), r *Manager) *{{.mockerName}}[{{.tmplReq}}, {{.resp}}] {
	PatchOnce(f)
	m := &{{.mockerName}}[{{.tmplReq}}, {{.resp}}]{}
	i := &{{.invokerName}}[{{.tmplReq}}, {{.resp}}]{ {{.mockerName}}: m}
	r.addMocker(f, i)
	return m
}
`))
