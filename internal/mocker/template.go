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

// mocker00Tmpl for no request parameters, returns nothing.
var mocker00Tmpl = template.Must(template.New("mocker00").Parse(`
/******************************** Mocker00 ***********************************/

type Mocker00 struct {
	fnHandle func()
	fnWhen   func() bool
	fnReturn func()
}

// Handle sets a custom function to handle requests.
func (m *Mocker00) Handle(fn func()) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker00) When(fn func() bool) *Mocker00 {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker00) Return(fn func()) {
	if m.fnWhen == nil {
		m.fnWhen = func() bool { return true }
	}
	m.fnReturn = fn
}

// ReturnValue sets a return function with predefined values.
func (m *Mocker00) ReturnValue() {
	if m.fnWhen == nil {
		m.fnWhen = func() bool { return true }
	}
	m.fnReturn = func() {}
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker00) ReturnDefault() {
	if m.fnWhen == nil {
		m.fnWhen = func() bool { return true }
	}
	m.Return(func() {})
}

// Invoker00 is an Invoker implementation for Mocker00.
type Invoker00 struct {
	*Mocker00
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker00) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker00) Handle(params []any) []any {
	m.fnHandle({{.cvtParams}})
	return []any{}
}

// When checks if the condition function evaluates to true.
func (m *Invoker00) When(params []any) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen({{.cvtParams}})
}

// Return provides predefined response and error values.
func (m *Invoker00) Return(params []any) []any {
	m.fnReturn()
	return []any{}
}

// NewMocker00 creates a new Mocker00 instance.
func NewMocker00(r *Manager, typ reflect.Type, method string) *Mocker00 {
	m := &Mocker00{}
	i := &Invoker00{ Mocker00: m}
	r.addMocker(typ, method, i)
	return m
}
`))

// mocker0NTmpl for no request parameters, returns N values.
var mocker0NTmpl = template.Must(template.New("mocker0N").Parse(`
/******************************** {{.mockerName}} ***********************************/

type {{.mockerName}}[{{.resp}} any] struct {
	fnHandle func() ({{.resp}})
	fnWhen   func() bool
	fnReturn func() ({{.resp}})
}

// Handle sets a custom function to handle requests.
func (m *{{.mockerName}}[{{.resp}}]) Handle(fn func() ({{.resp}})) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *{{.mockerName}}[{{.resp}}]) When(fn func() bool) *{{.mockerName}}[{{.resp}}] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *{{.mockerName}}[{{.resp}}]) Return(fn func() ({{.resp}})) {
	if m.fnWhen == nil {
		m.fnWhen = func() bool { return true }
	}
	m.fnReturn = fn
}

// ReturnValue sets a return function with predefined values.
func (m *{{.mockerName}}[{{.resp}}]) ReturnValue({{.respWithArg}}) {
	if m.fnWhen == nil {
		m.fnWhen = func() bool { return true }
	}
	m.fnReturn = func() ({{.resp}}) { return {{.respOnlyArg}} }
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *{{.mockerName}}[{{.resp}}]) ReturnDefault() {
	if m.fnWhen == nil {
		m.fnWhen = func() bool { return true }
	}
	m.Return(func() ({{.respWithArg}}) { return {{.respOnlyArg}} })
}

// {{.invokerName}} is an Invoker implementation for {{.mockerName}}.
type {{.invokerName}}[{{.resp}} any] struct {
	*{{.mockerName}}[{{.resp}}]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *{{.invokerName}}[{{.resp}}]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *{{.invokerName}}[{{.resp}}]) Handle(params []any) []any {
	{{.respOnlyArg}} := m.fnHandle({{.cvtParams}})
	return []any{ {{.respOnlyArg}}}
}

// When checks if the condition function evaluates to true.
func (m *{{.invokerName}}[{{.resp}}]) When(params []any) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen({{.cvtParams}})
}

// Return provides predefined response and error values.
func (m *{{.invokerName}}[{{.resp}}]) Return(params []any) []any {
	{{.respOnlyArg}} := m.fnReturn()
	return []any{ {{.respOnlyArg}}}
}

// New{{.mockerName}} creates a new {{.mockerName}} instance.
func New{{.mockerName}}[{{.resp}} any](r *Manager, typ reflect.Type, method string) *{{.mockerName}}[{{.resp}}] {
	m := &{{.mockerName}}[{{.resp}}]{}
	i := &{{.invokerName}}[{{.resp}}]{ {{.mockerName}}: m}
	r.addMocker(typ, method, i)
	return m
}
`))

// mockerN0Tmpl for N request parameters, returns nothing.
var mockerN0Tmpl = template.Must(template.New("mockerN0").Parse(`
/******************************** {{.mockerName}} ***********************************/

type {{.mockerName}}[{{.req}} any] struct {
	fnHandle func({{.req}})
	fnWhen   func({{.req}}) bool
	fnReturn func()
}

// Handle sets a custom function to handle requests.
func (m *{{.mockerName}}[{{.req}}]) Handle(fn func({{.req}})) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *{{.mockerName}}[{{.req}}]) When(fn func({{.req}}) bool) *{{.mockerName}}[{{.req}}] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *{{.mockerName}}[{{.req}}]) Return(fn func()) {
	if m.fnWhen == nil {
		m.fnWhen = func({{.req}}) bool { return true }
	}
	m.fnReturn = fn
}

// ReturnValue sets a return function with predefined values.
func (m *{{.mockerName}}[{{.req}}]) ReturnValue() {
	if m.fnWhen == nil {
		m.fnWhen = func({{.req}}) bool { return true }
	}
	m.fnReturn = func() {}
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *{{.mockerName}}[{{.req}}]) ReturnDefault() {
	if m.fnWhen == nil {
		m.fnWhen = func({{.req}}) bool { return true }
	}
	m.Return(func() {})
}

// {{.invokerName}} is an Invoker implementation for {{.mockerName}}.
type {{.invokerName}}[{{.req}} any] struct {
	*{{.mockerName}}[{{.req}}]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *{{.invokerName}}[{{.req}}]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *{{.invokerName}}[{{.req}}]) Handle(params []any) []any {
	m.fnHandle({{.cvtParams}})
	return []any{}
}

// When checks if the condition function evaluates to true.
func (m *{{.invokerName}}[{{.req}}]) When(params []any) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen({{.cvtParams}})
}

// Return provides predefined response and error values.
func (m *{{.invokerName}}[{{.req}}]) Return(params []any) []any {
	m.fnReturn()
	return []any{}
}

// New{{.mockerName}} creates a new {{.mockerName}} instance.
func New{{.mockerName}}[{{.req}} any](r *Manager, typ reflect.Type, method string) *{{.mockerName}}[{{.req}}] {
	m := &{{.mockerName}}[{{.req}}]{}
	i := &{{.invokerName}}[{{.req}}]{ {{.mockerName}}: m}
	r.addMocker(typ, method, i)
	return m
}
`))

// mockerNNTmpl for N request parameters, returns N values.
var mockerNNTmpl = template.Must(template.New("mockerNN").Parse(`
/******************************** {{.mockerName}} ***********************************/

type {{.mockerName}}[{{.req}} any, {{.resp}} any] struct {
	fnHandle func({{.req}}) ({{.resp}})
	fnWhen   func({{.req}}) bool
	fnReturn func() ({{.resp}})
}

// Handle sets a custom function to handle requests.
func (m *{{.mockerName}}[{{.req}}, {{.resp}}]) Handle(fn func({{.req}}) ({{.resp}})) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *{{.mockerName}}[{{.req}}, {{.resp}}]) When(fn func({{.req}}) bool) *{{.mockerName}}[{{.req}}, {{.resp}}] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *{{.mockerName}}[{{.req}}, {{.resp}}]) Return(fn func() ({{.resp}})) {
	if m.fnWhen == nil {
		m.fnWhen = func({{.req}}) bool { return true }
	}
	m.fnReturn = fn
}

// ReturnValue sets a return function with predefined values.
func (m *{{.mockerName}}[{{.req}}, {{.resp}}]) ReturnValue({{.respWithArg}}) {
	if m.fnWhen == nil {
		m.fnWhen = func({{.req}}) bool { return true }
	}
	m.fnReturn = func() ({{.resp}}) { return {{.respOnlyArg}} }
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *{{.mockerName}}[{{.req}}, {{.resp}}]) ReturnDefault() {
	if m.fnWhen == nil {
		m.fnWhen = func({{.req}}) bool { return true }
	}
	m.Return(func() ({{.respWithArg}}) { return {{.respOnlyArg}} })
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
func (m *{{.invokerName}}[{{.req}}, {{.resp}}]) Handle(params []any) []any {
	{{.respOnlyArg}} := m.fnHandle({{.cvtParams}})
	return []any{ {{.respOnlyArg}}}
}

// When checks if the condition function evaluates to true.
func (m *{{.invokerName}}[{{.req}}, {{.resp}}]) When(params []any) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen({{.cvtParams}})
}

// Return provides predefined response and error values.
func (m *{{.invokerName}}[{{.req}}, {{.resp}}]) Return(params []any) []any {
	{{.respOnlyArg}} := m.fnReturn()
	return []any{ {{.respOnlyArg}}}
}

// New{{.mockerName}} creates a new {{.mockerName}} instance.
func New{{.mockerName}}[{{.req}} any, {{.resp}} any](r *Manager, typ reflect.Type, method string) *{{.mockerName}}[{{.req}}, {{.resp}}] {
	m := &{{.mockerName}}[{{.req}}, {{.resp}}]{}
	i := &{{.invokerName}}[{{.req}}, {{.resp}}]{ {{.mockerName}}: m}
	r.addMocker(typ, method, i)
	return m
}
`))
