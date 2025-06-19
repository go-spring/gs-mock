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

package mock

import (
	"reflect"
)

const (
	MaxParamCount  = 5
	MaxResultCount = 5
)

/******************************** Mocker00 ***********************************/

type Mocker00 struct {
	fnHandle func() bool
	fnWhen   func() bool
	fnReturn func()
}

// Handle sets a custom function to handle requests.
func (m *Mocker00) Handle(fn func() bool) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker00) When(fn func() bool) *Mocker00 {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker00) Return(fn func()) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker00) Always() *Mocker00 {
	return m.When(func() bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker00) ReturnDefault() {
	m.Return(func() {})
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker00) Ignore() {
	m.Always().ReturnDefault()
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
func (m *Invoker00) Handle(params []interface{}) ([]interface{}, bool) {
	ok := m.fnHandle()
	return []interface{}{}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker00) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen()
}

// Return provides predefined response and error values.
func (m *Invoker00) Return(params []interface{}) []interface{} {
	m.fnReturn()
	return []interface{}{}
}

// NewMocker00 creates a new Mocker00 instance.
func NewMocker00(r *Manager, typ reflect.Type, method string) *Mocker00 {
	m := &Mocker00{}
	i := &Invoker00{Mocker00: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker01 ***********************************/

type Mocker01[R1 any] struct {
	fnHandle func() (R1, bool)
	fnWhen   func() bool
	fnReturn func() R1
}

// Handle sets a custom function to handle requests.
func (m *Mocker01[R1]) Handle(fn func() (R1, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker01[R1]) When(fn func() bool) *Mocker01[R1] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker01[R1]) Return(fn func() R1) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker01[R1]) Always() *Mocker01[R1] {
	return m.When(func() bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker01[R1]) ReturnDefault() {
	m.Return(func() (r1 R1) { return r1 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker01[R1]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker01 is an Invoker implementation for Mocker01.
type Invoker01[R1 any] struct {
	*Mocker01[R1]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker01[R1]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker01[R1]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, ok := m.fnHandle()
	return []interface{}{r1}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker01[R1]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen()
}

// Return provides predefined response and error values.
func (m *Invoker01[R1]) Return(params []interface{}) []interface{} {
	r1 := m.fnReturn()
	return []interface{}{r1}
}

// NewMocker01 creates a new Mocker01 instance.
func NewMocker01[R1 any](r *Manager, typ reflect.Type, method string) *Mocker01[R1] {
	m := &Mocker01[R1]{}
	i := &Invoker01[R1]{Mocker01: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker02 ***********************************/

type Mocker02[R1, R2 any] struct {
	fnHandle func() (R1, R2, bool)
	fnWhen   func() bool
	fnReturn func() (R1, R2)
}

// Handle sets a custom function to handle requests.
func (m *Mocker02[R1, R2]) Handle(fn func() (R1, R2, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker02[R1, R2]) When(fn func() bool) *Mocker02[R1, R2] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker02[R1, R2]) Return(fn func() (R1, R2)) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker02[R1, R2]) Always() *Mocker02[R1, R2] {
	return m.When(func() bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker02[R1, R2]) ReturnDefault() {
	m.Return(func() (r1 R1, r2 R2) { return r1, r2 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker02[R1, R2]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker02 is an Invoker implementation for Mocker02.
type Invoker02[R1, R2 any] struct {
	*Mocker02[R1, R2]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker02[R1, R2]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker02[R1, R2]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, r2, ok := m.fnHandle()
	return []interface{}{r1, r2}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker02[R1, R2]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen()
}

// Return provides predefined response and error values.
func (m *Invoker02[R1, R2]) Return(params []interface{}) []interface{} {
	r1, r2 := m.fnReturn()
	return []interface{}{r1, r2}
}

// NewMocker02 creates a new Mocker02 instance.
func NewMocker02[R1, R2 any](r *Manager, typ reflect.Type, method string) *Mocker02[R1, R2] {
	m := &Mocker02[R1, R2]{}
	i := &Invoker02[R1, R2]{Mocker02: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker03 ***********************************/

type Mocker03[R1, R2, R3 any] struct {
	fnHandle func() (R1, R2, R3, bool)
	fnWhen   func() bool
	fnReturn func() (R1, R2, R3)
}

// Handle sets a custom function to handle requests.
func (m *Mocker03[R1, R2, R3]) Handle(fn func() (R1, R2, R3, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker03[R1, R2, R3]) When(fn func() bool) *Mocker03[R1, R2, R3] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker03[R1, R2, R3]) Return(fn func() (R1, R2, R3)) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker03[R1, R2, R3]) Always() *Mocker03[R1, R2, R3] {
	return m.When(func() bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker03[R1, R2, R3]) ReturnDefault() {
	m.Return(func() (r1 R1, r2 R2, r3 R3) { return r1, r2, r3 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker03[R1, R2, R3]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker03 is an Invoker implementation for Mocker03.
type Invoker03[R1, R2, R3 any] struct {
	*Mocker03[R1, R2, R3]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker03[R1, R2, R3]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker03[R1, R2, R3]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, r2, r3, ok := m.fnHandle()
	return []interface{}{r1, r2, r3}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker03[R1, R2, R3]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen()
}

// Return provides predefined response and error values.
func (m *Invoker03[R1, R2, R3]) Return(params []interface{}) []interface{} {
	r1, r2, r3 := m.fnReturn()
	return []interface{}{r1, r2, r3}
}

// NewMocker03 creates a new Mocker03 instance.
func NewMocker03[R1, R2, R3 any](r *Manager, typ reflect.Type, method string) *Mocker03[R1, R2, R3] {
	m := &Mocker03[R1, R2, R3]{}
	i := &Invoker03[R1, R2, R3]{Mocker03: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker04 ***********************************/

type Mocker04[R1, R2, R3, R4 any] struct {
	fnHandle func() (R1, R2, R3, R4, bool)
	fnWhen   func() bool
	fnReturn func() (R1, R2, R3, R4)
}

// Handle sets a custom function to handle requests.
func (m *Mocker04[R1, R2, R3, R4]) Handle(fn func() (R1, R2, R3, R4, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker04[R1, R2, R3, R4]) When(fn func() bool) *Mocker04[R1, R2, R3, R4] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker04[R1, R2, R3, R4]) Return(fn func() (R1, R2, R3, R4)) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker04[R1, R2, R3, R4]) Always() *Mocker04[R1, R2, R3, R4] {
	return m.When(func() bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker04[R1, R2, R3, R4]) ReturnDefault() {
	m.Return(func() (r1 R1, r2 R2, r3 R3, r4 R4) { return r1, r2, r3, r4 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker04[R1, R2, R3, R4]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker04 is an Invoker implementation for Mocker04.
type Invoker04[R1, R2, R3, R4 any] struct {
	*Mocker04[R1, R2, R3, R4]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker04[R1, R2, R3, R4]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker04[R1, R2, R3, R4]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, r2, r3, r4, ok := m.fnHandle()
	return []interface{}{r1, r2, r3, r4}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker04[R1, R2, R3, R4]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen()
}

// Return provides predefined response and error values.
func (m *Invoker04[R1, R2, R3, R4]) Return(params []interface{}) []interface{} {
	r1, r2, r3, r4 := m.fnReturn()
	return []interface{}{r1, r2, r3, r4}
}

// NewMocker04 creates a new Mocker04 instance.
func NewMocker04[R1, R2, R3, R4 any](r *Manager, typ reflect.Type, method string) *Mocker04[R1, R2, R3, R4] {
	m := &Mocker04[R1, R2, R3, R4]{}
	i := &Invoker04[R1, R2, R3, R4]{Mocker04: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker05 ***********************************/

type Mocker05[R1, R2, R3, R4, R5 any] struct {
	fnHandle func() (R1, R2, R3, R4, R5, bool)
	fnWhen   func() bool
	fnReturn func() (R1, R2, R3, R4, R5)
}

// Handle sets a custom function to handle requests.
func (m *Mocker05[R1, R2, R3, R4, R5]) Handle(fn func() (R1, R2, R3, R4, R5, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker05[R1, R2, R3, R4, R5]) When(fn func() bool) *Mocker05[R1, R2, R3, R4, R5] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker05[R1, R2, R3, R4, R5]) Return(fn func() (R1, R2, R3, R4, R5)) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker05[R1, R2, R3, R4, R5]) Always() *Mocker05[R1, R2, R3, R4, R5] {
	return m.When(func() bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker05[R1, R2, R3, R4, R5]) ReturnDefault() {
	m.Return(func() (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) { return r1, r2, r3, r4, r5 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker05[R1, R2, R3, R4, R5]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker05 is an Invoker implementation for Mocker05.
type Invoker05[R1, R2, R3, R4, R5 any] struct {
	*Mocker05[R1, R2, R3, R4, R5]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker05[R1, R2, R3, R4, R5]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker05[R1, R2, R3, R4, R5]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, r2, r3, r4, r5, ok := m.fnHandle()
	return []interface{}{r1, r2, r3, r4, r5}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker05[R1, R2, R3, R4, R5]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen()
}

// Return provides predefined response and error values.
func (m *Invoker05[R1, R2, R3, R4, R5]) Return(params []interface{}) []interface{} {
	r1, r2, r3, r4, r5 := m.fnReturn()
	return []interface{}{r1, r2, r3, r4, r5}
}

// NewMocker05 creates a new Mocker05 instance.
func NewMocker05[R1, R2, R3, R4, R5 any](r *Manager, typ reflect.Type, method string) *Mocker05[R1, R2, R3, R4, R5] {
	m := &Mocker05[R1, R2, R3, R4, R5]{}
	i := &Invoker05[R1, R2, R3, R4, R5]{Mocker05: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker10 ***********************************/

type Mocker10[T1 any] struct {
	fnHandle func(T1) bool
	fnWhen   func(T1) bool
	fnReturn func()
}

// Handle sets a custom function to handle requests.
func (m *Mocker10[T1]) Handle(fn func(T1) bool) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker10[T1]) When(fn func(T1) bool) *Mocker10[T1] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker10[T1]) Return(fn func()) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker10[T1]) Always() *Mocker10[T1] {
	return m.When(func(T1) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker10[T1]) ReturnDefault() {
	m.Return(func() {})
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker10[T1]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker10 is an Invoker implementation for Mocker10.
type Invoker10[T1 any] struct {
	*Mocker10[T1]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker10[T1]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker10[T1]) Handle(params []interface{}) ([]interface{}, bool) {
	ok := m.fnHandle(params[0].(T1))
	return []interface{}{}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker10[T1]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1))
}

// Return provides predefined response and error values.
func (m *Invoker10[T1]) Return(params []interface{}) []interface{} {
	m.fnReturn()
	return []interface{}{}
}

// NewMocker10 creates a new Mocker10 instance.
func NewMocker10[T1 any](r *Manager, typ reflect.Type, method string) *Mocker10[T1] {
	m := &Mocker10[T1]{}
	i := &Invoker10[T1]{Mocker10: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker11 ***********************************/

type Mocker11[T1 any, R1 any] struct {
	fnHandle func(T1) (R1, bool)
	fnWhen   func(T1) bool
	fnReturn func() R1
}

// Handle sets a custom function to handle requests.
func (m *Mocker11[T1, R1]) Handle(fn func(T1) (R1, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker11[T1, R1]) When(fn func(T1) bool) *Mocker11[T1, R1] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker11[T1, R1]) Return(fn func() R1) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker11[T1, R1]) Always() *Mocker11[T1, R1] {
	return m.When(func(T1) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker11[T1, R1]) ReturnDefault() {
	m.Return(func() (r1 R1) { return r1 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker11[T1, R1]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker11 is an Invoker implementation for Mocker11.
type Invoker11[T1 any, R1 any] struct {
	*Mocker11[T1, R1]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker11[T1, R1]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker11[T1, R1]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, ok := m.fnHandle(params[0].(T1))
	return []interface{}{r1}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker11[T1, R1]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1))
}

// Return provides predefined response and error values.
func (m *Invoker11[T1, R1]) Return(params []interface{}) []interface{} {
	r1 := m.fnReturn()
	return []interface{}{r1}
}

// NewMocker11 creates a new Mocker11 instance.
func NewMocker11[T1 any, R1 any](r *Manager, typ reflect.Type, method string) *Mocker11[T1, R1] {
	m := &Mocker11[T1, R1]{}
	i := &Invoker11[T1, R1]{Mocker11: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker12 ***********************************/

type Mocker12[T1 any, R1, R2 any] struct {
	fnHandle func(T1) (R1, R2, bool)
	fnWhen   func(T1) bool
	fnReturn func() (R1, R2)
}

// Handle sets a custom function to handle requests.
func (m *Mocker12[T1, R1, R2]) Handle(fn func(T1) (R1, R2, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker12[T1, R1, R2]) When(fn func(T1) bool) *Mocker12[T1, R1, R2] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker12[T1, R1, R2]) Return(fn func() (R1, R2)) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker12[T1, R1, R2]) Always() *Mocker12[T1, R1, R2] {
	return m.When(func(T1) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker12[T1, R1, R2]) ReturnDefault() {
	m.Return(func() (r1 R1, r2 R2) { return r1, r2 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker12[T1, R1, R2]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker12 is an Invoker implementation for Mocker12.
type Invoker12[T1 any, R1, R2 any] struct {
	*Mocker12[T1, R1, R2]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker12[T1, R1, R2]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker12[T1, R1, R2]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, r2, ok := m.fnHandle(params[0].(T1))
	return []interface{}{r1, r2}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker12[T1, R1, R2]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1))
}

// Return provides predefined response and error values.
func (m *Invoker12[T1, R1, R2]) Return(params []interface{}) []interface{} {
	r1, r2 := m.fnReturn()
	return []interface{}{r1, r2}
}

// NewMocker12 creates a new Mocker12 instance.
func NewMocker12[T1 any, R1, R2 any](r *Manager, typ reflect.Type, method string) *Mocker12[T1, R1, R2] {
	m := &Mocker12[T1, R1, R2]{}
	i := &Invoker12[T1, R1, R2]{Mocker12: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker13 ***********************************/

type Mocker13[T1 any, R1, R2, R3 any] struct {
	fnHandle func(T1) (R1, R2, R3, bool)
	fnWhen   func(T1) bool
	fnReturn func() (R1, R2, R3)
}

// Handle sets a custom function to handle requests.
func (m *Mocker13[T1, R1, R2, R3]) Handle(fn func(T1) (R1, R2, R3, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker13[T1, R1, R2, R3]) When(fn func(T1) bool) *Mocker13[T1, R1, R2, R3] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker13[T1, R1, R2, R3]) Return(fn func() (R1, R2, R3)) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker13[T1, R1, R2, R3]) Always() *Mocker13[T1, R1, R2, R3] {
	return m.When(func(T1) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker13[T1, R1, R2, R3]) ReturnDefault() {
	m.Return(func() (r1 R1, r2 R2, r3 R3) { return r1, r2, r3 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker13[T1, R1, R2, R3]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker13 is an Invoker implementation for Mocker13.
type Invoker13[T1 any, R1, R2, R3 any] struct {
	*Mocker13[T1, R1, R2, R3]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker13[T1, R1, R2, R3]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker13[T1, R1, R2, R3]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, r2, r3, ok := m.fnHandle(params[0].(T1))
	return []interface{}{r1, r2, r3}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker13[T1, R1, R2, R3]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1))
}

// Return provides predefined response and error values.
func (m *Invoker13[T1, R1, R2, R3]) Return(params []interface{}) []interface{} {
	r1, r2, r3 := m.fnReturn()
	return []interface{}{r1, r2, r3}
}

// NewMocker13 creates a new Mocker13 instance.
func NewMocker13[T1 any, R1, R2, R3 any](r *Manager, typ reflect.Type, method string) *Mocker13[T1, R1, R2, R3] {
	m := &Mocker13[T1, R1, R2, R3]{}
	i := &Invoker13[T1, R1, R2, R3]{Mocker13: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker14 ***********************************/

type Mocker14[T1 any, R1, R2, R3, R4 any] struct {
	fnHandle func(T1) (R1, R2, R3, R4, bool)
	fnWhen   func(T1) bool
	fnReturn func() (R1, R2, R3, R4)
}

// Handle sets a custom function to handle requests.
func (m *Mocker14[T1, R1, R2, R3, R4]) Handle(fn func(T1) (R1, R2, R3, R4, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker14[T1, R1, R2, R3, R4]) When(fn func(T1) bool) *Mocker14[T1, R1, R2, R3, R4] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker14[T1, R1, R2, R3, R4]) Return(fn func() (R1, R2, R3, R4)) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker14[T1, R1, R2, R3, R4]) Always() *Mocker14[T1, R1, R2, R3, R4] {
	return m.When(func(T1) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker14[T1, R1, R2, R3, R4]) ReturnDefault() {
	m.Return(func() (r1 R1, r2 R2, r3 R3, r4 R4) { return r1, r2, r3, r4 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker14[T1, R1, R2, R3, R4]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker14 is an Invoker implementation for Mocker14.
type Invoker14[T1 any, R1, R2, R3, R4 any] struct {
	*Mocker14[T1, R1, R2, R3, R4]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker14[T1, R1, R2, R3, R4]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker14[T1, R1, R2, R3, R4]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, r2, r3, r4, ok := m.fnHandle(params[0].(T1))
	return []interface{}{r1, r2, r3, r4}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker14[T1, R1, R2, R3, R4]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1))
}

// Return provides predefined response and error values.
func (m *Invoker14[T1, R1, R2, R3, R4]) Return(params []interface{}) []interface{} {
	r1, r2, r3, r4 := m.fnReturn()
	return []interface{}{r1, r2, r3, r4}
}

// NewMocker14 creates a new Mocker14 instance.
func NewMocker14[T1 any, R1, R2, R3, R4 any](r *Manager, typ reflect.Type, method string) *Mocker14[T1, R1, R2, R3, R4] {
	m := &Mocker14[T1, R1, R2, R3, R4]{}
	i := &Invoker14[T1, R1, R2, R3, R4]{Mocker14: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker15 ***********************************/

type Mocker15[T1 any, R1, R2, R3, R4, R5 any] struct {
	fnHandle func(T1) (R1, R2, R3, R4, R5, bool)
	fnWhen   func(T1) bool
	fnReturn func() (R1, R2, R3, R4, R5)
}

// Handle sets a custom function to handle requests.
func (m *Mocker15[T1, R1, R2, R3, R4, R5]) Handle(fn func(T1) (R1, R2, R3, R4, R5, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker15[T1, R1, R2, R3, R4, R5]) When(fn func(T1) bool) *Mocker15[T1, R1, R2, R3, R4, R5] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker15[T1, R1, R2, R3, R4, R5]) Return(fn func() (R1, R2, R3, R4, R5)) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker15[T1, R1, R2, R3, R4, R5]) Always() *Mocker15[T1, R1, R2, R3, R4, R5] {
	return m.When(func(T1) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker15[T1, R1, R2, R3, R4, R5]) ReturnDefault() {
	m.Return(func() (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) { return r1, r2, r3, r4, r5 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker15[T1, R1, R2, R3, R4, R5]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker15 is an Invoker implementation for Mocker15.
type Invoker15[T1 any, R1, R2, R3, R4, R5 any] struct {
	*Mocker15[T1, R1, R2, R3, R4, R5]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker15[T1, R1, R2, R3, R4, R5]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker15[T1, R1, R2, R3, R4, R5]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, r2, r3, r4, r5, ok := m.fnHandle(params[0].(T1))
	return []interface{}{r1, r2, r3, r4, r5}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker15[T1, R1, R2, R3, R4, R5]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1))
}

// Return provides predefined response and error values.
func (m *Invoker15[T1, R1, R2, R3, R4, R5]) Return(params []interface{}) []interface{} {
	r1, r2, r3, r4, r5 := m.fnReturn()
	return []interface{}{r1, r2, r3, r4, r5}
}

// NewMocker15 creates a new Mocker15 instance.
func NewMocker15[T1 any, R1, R2, R3, R4, R5 any](r *Manager, typ reflect.Type, method string) *Mocker15[T1, R1, R2, R3, R4, R5] {
	m := &Mocker15[T1, R1, R2, R3, R4, R5]{}
	i := &Invoker15[T1, R1, R2, R3, R4, R5]{Mocker15: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker20 ***********************************/

type Mocker20[T1, T2 any] struct {
	fnHandle func(T1, T2) bool
	fnWhen   func(T1, T2) bool
	fnReturn func()
}

// Handle sets a custom function to handle requests.
func (m *Mocker20[T1, T2]) Handle(fn func(T1, T2) bool) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker20[T1, T2]) When(fn func(T1, T2) bool) *Mocker20[T1, T2] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker20[T1, T2]) Return(fn func()) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker20[T1, T2]) Always() *Mocker20[T1, T2] {
	return m.When(func(T1, T2) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker20[T1, T2]) ReturnDefault() {
	m.Return(func() {})
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker20[T1, T2]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker20 is an Invoker implementation for Mocker20.
type Invoker20[T1, T2 any] struct {
	*Mocker20[T1, T2]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker20[T1, T2]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker20[T1, T2]) Handle(params []interface{}) ([]interface{}, bool) {
	ok := m.fnHandle(params[0].(T1), params[1].(T2))
	return []interface{}{}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker20[T1, T2]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1), params[1].(T2))
}

// Return provides predefined response and error values.
func (m *Invoker20[T1, T2]) Return(params []interface{}) []interface{} {
	m.fnReturn()
	return []interface{}{}
}

// NewMocker20 creates a new Mocker20 instance.
func NewMocker20[T1, T2 any](r *Manager, typ reflect.Type, method string) *Mocker20[T1, T2] {
	m := &Mocker20[T1, T2]{}
	i := &Invoker20[T1, T2]{Mocker20: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker21 ***********************************/

type Mocker21[T1, T2 any, R1 any] struct {
	fnHandle func(T1, T2) (R1, bool)
	fnWhen   func(T1, T2) bool
	fnReturn func() R1
}

// Handle sets a custom function to handle requests.
func (m *Mocker21[T1, T2, R1]) Handle(fn func(T1, T2) (R1, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker21[T1, T2, R1]) When(fn func(T1, T2) bool) *Mocker21[T1, T2, R1] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker21[T1, T2, R1]) Return(fn func() R1) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker21[T1, T2, R1]) Always() *Mocker21[T1, T2, R1] {
	return m.When(func(T1, T2) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker21[T1, T2, R1]) ReturnDefault() {
	m.Return(func() (r1 R1) { return r1 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker21[T1, T2, R1]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker21 is an Invoker implementation for Mocker21.
type Invoker21[T1, T2 any, R1 any] struct {
	*Mocker21[T1, T2, R1]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker21[T1, T2, R1]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker21[T1, T2, R1]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, ok := m.fnHandle(params[0].(T1), params[1].(T2))
	return []interface{}{r1}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker21[T1, T2, R1]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1), params[1].(T2))
}

// Return provides predefined response and error values.
func (m *Invoker21[T1, T2, R1]) Return(params []interface{}) []interface{} {
	r1 := m.fnReturn()
	return []interface{}{r1}
}

// NewMocker21 creates a new Mocker21 instance.
func NewMocker21[T1, T2 any, R1 any](r *Manager, typ reflect.Type, method string) *Mocker21[T1, T2, R1] {
	m := &Mocker21[T1, T2, R1]{}
	i := &Invoker21[T1, T2, R1]{Mocker21: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker22 ***********************************/

type Mocker22[T1, T2 any, R1, R2 any] struct {
	fnHandle func(T1, T2) (R1, R2, bool)
	fnWhen   func(T1, T2) bool
	fnReturn func() (R1, R2)
}

// Handle sets a custom function to handle requests.
func (m *Mocker22[T1, T2, R1, R2]) Handle(fn func(T1, T2) (R1, R2, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker22[T1, T2, R1, R2]) When(fn func(T1, T2) bool) *Mocker22[T1, T2, R1, R2] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker22[T1, T2, R1, R2]) Return(fn func() (R1, R2)) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker22[T1, T2, R1, R2]) Always() *Mocker22[T1, T2, R1, R2] {
	return m.When(func(T1, T2) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker22[T1, T2, R1, R2]) ReturnDefault() {
	m.Return(func() (r1 R1, r2 R2) { return r1, r2 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker22[T1, T2, R1, R2]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker22 is an Invoker implementation for Mocker22.
type Invoker22[T1, T2 any, R1, R2 any] struct {
	*Mocker22[T1, T2, R1, R2]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker22[T1, T2, R1, R2]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker22[T1, T2, R1, R2]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, r2, ok := m.fnHandle(params[0].(T1), params[1].(T2))
	return []interface{}{r1, r2}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker22[T1, T2, R1, R2]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1), params[1].(T2))
}

// Return provides predefined response and error values.
func (m *Invoker22[T1, T2, R1, R2]) Return(params []interface{}) []interface{} {
	r1, r2 := m.fnReturn()
	return []interface{}{r1, r2}
}

// NewMocker22 creates a new Mocker22 instance.
func NewMocker22[T1, T2 any, R1, R2 any](r *Manager, typ reflect.Type, method string) *Mocker22[T1, T2, R1, R2] {
	m := &Mocker22[T1, T2, R1, R2]{}
	i := &Invoker22[T1, T2, R1, R2]{Mocker22: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker23 ***********************************/

type Mocker23[T1, T2 any, R1, R2, R3 any] struct {
	fnHandle func(T1, T2) (R1, R2, R3, bool)
	fnWhen   func(T1, T2) bool
	fnReturn func() (R1, R2, R3)
}

// Handle sets a custom function to handle requests.
func (m *Mocker23[T1, T2, R1, R2, R3]) Handle(fn func(T1, T2) (R1, R2, R3, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker23[T1, T2, R1, R2, R3]) When(fn func(T1, T2) bool) *Mocker23[T1, T2, R1, R2, R3] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker23[T1, T2, R1, R2, R3]) Return(fn func() (R1, R2, R3)) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker23[T1, T2, R1, R2, R3]) Always() *Mocker23[T1, T2, R1, R2, R3] {
	return m.When(func(T1, T2) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker23[T1, T2, R1, R2, R3]) ReturnDefault() {
	m.Return(func() (r1 R1, r2 R2, r3 R3) { return r1, r2, r3 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker23[T1, T2, R1, R2, R3]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker23 is an Invoker implementation for Mocker23.
type Invoker23[T1, T2 any, R1, R2, R3 any] struct {
	*Mocker23[T1, T2, R1, R2, R3]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker23[T1, T2, R1, R2, R3]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker23[T1, T2, R1, R2, R3]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, r2, r3, ok := m.fnHandle(params[0].(T1), params[1].(T2))
	return []interface{}{r1, r2, r3}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker23[T1, T2, R1, R2, R3]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1), params[1].(T2))
}

// Return provides predefined response and error values.
func (m *Invoker23[T1, T2, R1, R2, R3]) Return(params []interface{}) []interface{} {
	r1, r2, r3 := m.fnReturn()
	return []interface{}{r1, r2, r3}
}

// NewMocker23 creates a new Mocker23 instance.
func NewMocker23[T1, T2 any, R1, R2, R3 any](r *Manager, typ reflect.Type, method string) *Mocker23[T1, T2, R1, R2, R3] {
	m := &Mocker23[T1, T2, R1, R2, R3]{}
	i := &Invoker23[T1, T2, R1, R2, R3]{Mocker23: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker24 ***********************************/

type Mocker24[T1, T2 any, R1, R2, R3, R4 any] struct {
	fnHandle func(T1, T2) (R1, R2, R3, R4, bool)
	fnWhen   func(T1, T2) bool
	fnReturn func() (R1, R2, R3, R4)
}

// Handle sets a custom function to handle requests.
func (m *Mocker24[T1, T2, R1, R2, R3, R4]) Handle(fn func(T1, T2) (R1, R2, R3, R4, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker24[T1, T2, R1, R2, R3, R4]) When(fn func(T1, T2) bool) *Mocker24[T1, T2, R1, R2, R3, R4] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker24[T1, T2, R1, R2, R3, R4]) Return(fn func() (R1, R2, R3, R4)) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker24[T1, T2, R1, R2, R3, R4]) Always() *Mocker24[T1, T2, R1, R2, R3, R4] {
	return m.When(func(T1, T2) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker24[T1, T2, R1, R2, R3, R4]) ReturnDefault() {
	m.Return(func() (r1 R1, r2 R2, r3 R3, r4 R4) { return r1, r2, r3, r4 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker24[T1, T2, R1, R2, R3, R4]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker24 is an Invoker implementation for Mocker24.
type Invoker24[T1, T2 any, R1, R2, R3, R4 any] struct {
	*Mocker24[T1, T2, R1, R2, R3, R4]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker24[T1, T2, R1, R2, R3, R4]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker24[T1, T2, R1, R2, R3, R4]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, r2, r3, r4, ok := m.fnHandle(params[0].(T1), params[1].(T2))
	return []interface{}{r1, r2, r3, r4}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker24[T1, T2, R1, R2, R3, R4]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1), params[1].(T2))
}

// Return provides predefined response and error values.
func (m *Invoker24[T1, T2, R1, R2, R3, R4]) Return(params []interface{}) []interface{} {
	r1, r2, r3, r4 := m.fnReturn()
	return []interface{}{r1, r2, r3, r4}
}

// NewMocker24 creates a new Mocker24 instance.
func NewMocker24[T1, T2 any, R1, R2, R3, R4 any](r *Manager, typ reflect.Type, method string) *Mocker24[T1, T2, R1, R2, R3, R4] {
	m := &Mocker24[T1, T2, R1, R2, R3, R4]{}
	i := &Invoker24[T1, T2, R1, R2, R3, R4]{Mocker24: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker25 ***********************************/

type Mocker25[T1, T2 any, R1, R2, R3, R4, R5 any] struct {
	fnHandle func(T1, T2) (R1, R2, R3, R4, R5, bool)
	fnWhen   func(T1, T2) bool
	fnReturn func() (R1, R2, R3, R4, R5)
}

// Handle sets a custom function to handle requests.
func (m *Mocker25[T1, T2, R1, R2, R3, R4, R5]) Handle(fn func(T1, T2) (R1, R2, R3, R4, R5, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker25[T1, T2, R1, R2, R3, R4, R5]) When(fn func(T1, T2) bool) *Mocker25[T1, T2, R1, R2, R3, R4, R5] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker25[T1, T2, R1, R2, R3, R4, R5]) Return(fn func() (R1, R2, R3, R4, R5)) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker25[T1, T2, R1, R2, R3, R4, R5]) Always() *Mocker25[T1, T2, R1, R2, R3, R4, R5] {
	return m.When(func(T1, T2) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker25[T1, T2, R1, R2, R3, R4, R5]) ReturnDefault() {
	m.Return(func() (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) { return r1, r2, r3, r4, r5 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker25[T1, T2, R1, R2, R3, R4, R5]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker25 is an Invoker implementation for Mocker25.
type Invoker25[T1, T2 any, R1, R2, R3, R4, R5 any] struct {
	*Mocker25[T1, T2, R1, R2, R3, R4, R5]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker25[T1, T2, R1, R2, R3, R4, R5]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker25[T1, T2, R1, R2, R3, R4, R5]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, r2, r3, r4, r5, ok := m.fnHandle(params[0].(T1), params[1].(T2))
	return []interface{}{r1, r2, r3, r4, r5}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker25[T1, T2, R1, R2, R3, R4, R5]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1), params[1].(T2))
}

// Return provides predefined response and error values.
func (m *Invoker25[T1, T2, R1, R2, R3, R4, R5]) Return(params []interface{}) []interface{} {
	r1, r2, r3, r4, r5 := m.fnReturn()
	return []interface{}{r1, r2, r3, r4, r5}
}

// NewMocker25 creates a new Mocker25 instance.
func NewMocker25[T1, T2 any, R1, R2, R3, R4, R5 any](r *Manager, typ reflect.Type, method string) *Mocker25[T1, T2, R1, R2, R3, R4, R5] {
	m := &Mocker25[T1, T2, R1, R2, R3, R4, R5]{}
	i := &Invoker25[T1, T2, R1, R2, R3, R4, R5]{Mocker25: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker30 ***********************************/

type Mocker30[T1, T2, T3 any] struct {
	fnHandle func(T1, T2, T3) bool
	fnWhen   func(T1, T2, T3) bool
	fnReturn func()
}

// Handle sets a custom function to handle requests.
func (m *Mocker30[T1, T2, T3]) Handle(fn func(T1, T2, T3) bool) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker30[T1, T2, T3]) When(fn func(T1, T2, T3) bool) *Mocker30[T1, T2, T3] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker30[T1, T2, T3]) Return(fn func()) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker30[T1, T2, T3]) Always() *Mocker30[T1, T2, T3] {
	return m.When(func(T1, T2, T3) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker30[T1, T2, T3]) ReturnDefault() {
	m.Return(func() {})
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker30[T1, T2, T3]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker30 is an Invoker implementation for Mocker30.
type Invoker30[T1, T2, T3 any] struct {
	*Mocker30[T1, T2, T3]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker30[T1, T2, T3]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker30[T1, T2, T3]) Handle(params []interface{}) ([]interface{}, bool) {
	ok := m.fnHandle(params[0].(T1), params[1].(T2), params[2].(T3))
	return []interface{}{}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker30[T1, T2, T3]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1), params[1].(T2), params[2].(T3))
}

// Return provides predefined response and error values.
func (m *Invoker30[T1, T2, T3]) Return(params []interface{}) []interface{} {
	m.fnReturn()
	return []interface{}{}
}

// NewMocker30 creates a new Mocker30 instance.
func NewMocker30[T1, T2, T3 any](r *Manager, typ reflect.Type, method string) *Mocker30[T1, T2, T3] {
	m := &Mocker30[T1, T2, T3]{}
	i := &Invoker30[T1, T2, T3]{Mocker30: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker31 ***********************************/

type Mocker31[T1, T2, T3 any, R1 any] struct {
	fnHandle func(T1, T2, T3) (R1, bool)
	fnWhen   func(T1, T2, T3) bool
	fnReturn func() R1
}

// Handle sets a custom function to handle requests.
func (m *Mocker31[T1, T2, T3, R1]) Handle(fn func(T1, T2, T3) (R1, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker31[T1, T2, T3, R1]) When(fn func(T1, T2, T3) bool) *Mocker31[T1, T2, T3, R1] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker31[T1, T2, T3, R1]) Return(fn func() R1) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker31[T1, T2, T3, R1]) Always() *Mocker31[T1, T2, T3, R1] {
	return m.When(func(T1, T2, T3) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker31[T1, T2, T3, R1]) ReturnDefault() {
	m.Return(func() (r1 R1) { return r1 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker31[T1, T2, T3, R1]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker31 is an Invoker implementation for Mocker31.
type Invoker31[T1, T2, T3 any, R1 any] struct {
	*Mocker31[T1, T2, T3, R1]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker31[T1, T2, T3, R1]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker31[T1, T2, T3, R1]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, ok := m.fnHandle(params[0].(T1), params[1].(T2), params[2].(T3))
	return []interface{}{r1}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker31[T1, T2, T3, R1]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1), params[1].(T2), params[2].(T3))
}

// Return provides predefined response and error values.
func (m *Invoker31[T1, T2, T3, R1]) Return(params []interface{}) []interface{} {
	r1 := m.fnReturn()
	return []interface{}{r1}
}

// NewMocker31 creates a new Mocker31 instance.
func NewMocker31[T1, T2, T3 any, R1 any](r *Manager, typ reflect.Type, method string) *Mocker31[T1, T2, T3, R1] {
	m := &Mocker31[T1, T2, T3, R1]{}
	i := &Invoker31[T1, T2, T3, R1]{Mocker31: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker32 ***********************************/

type Mocker32[T1, T2, T3 any, R1, R2 any] struct {
	fnHandle func(T1, T2, T3) (R1, R2, bool)
	fnWhen   func(T1, T2, T3) bool
	fnReturn func() (R1, R2)
}

// Handle sets a custom function to handle requests.
func (m *Mocker32[T1, T2, T3, R1, R2]) Handle(fn func(T1, T2, T3) (R1, R2, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker32[T1, T2, T3, R1, R2]) When(fn func(T1, T2, T3) bool) *Mocker32[T1, T2, T3, R1, R2] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker32[T1, T2, T3, R1, R2]) Return(fn func() (R1, R2)) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker32[T1, T2, T3, R1, R2]) Always() *Mocker32[T1, T2, T3, R1, R2] {
	return m.When(func(T1, T2, T3) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker32[T1, T2, T3, R1, R2]) ReturnDefault() {
	m.Return(func() (r1 R1, r2 R2) { return r1, r2 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker32[T1, T2, T3, R1, R2]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker32 is an Invoker implementation for Mocker32.
type Invoker32[T1, T2, T3 any, R1, R2 any] struct {
	*Mocker32[T1, T2, T3, R1, R2]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker32[T1, T2, T3, R1, R2]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker32[T1, T2, T3, R1, R2]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, r2, ok := m.fnHandle(params[0].(T1), params[1].(T2), params[2].(T3))
	return []interface{}{r1, r2}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker32[T1, T2, T3, R1, R2]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1), params[1].(T2), params[2].(T3))
}

// Return provides predefined response and error values.
func (m *Invoker32[T1, T2, T3, R1, R2]) Return(params []interface{}) []interface{} {
	r1, r2 := m.fnReturn()
	return []interface{}{r1, r2}
}

// NewMocker32 creates a new Mocker32 instance.
func NewMocker32[T1, T2, T3 any, R1, R2 any](r *Manager, typ reflect.Type, method string) *Mocker32[T1, T2, T3, R1, R2] {
	m := &Mocker32[T1, T2, T3, R1, R2]{}
	i := &Invoker32[T1, T2, T3, R1, R2]{Mocker32: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker33 ***********************************/

type Mocker33[T1, T2, T3 any, R1, R2, R3 any] struct {
	fnHandle func(T1, T2, T3) (R1, R2, R3, bool)
	fnWhen   func(T1, T2, T3) bool
	fnReturn func() (R1, R2, R3)
}

// Handle sets a custom function to handle requests.
func (m *Mocker33[T1, T2, T3, R1, R2, R3]) Handle(fn func(T1, T2, T3) (R1, R2, R3, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker33[T1, T2, T3, R1, R2, R3]) When(fn func(T1, T2, T3) bool) *Mocker33[T1, T2, T3, R1, R2, R3] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker33[T1, T2, T3, R1, R2, R3]) Return(fn func() (R1, R2, R3)) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker33[T1, T2, T3, R1, R2, R3]) Always() *Mocker33[T1, T2, T3, R1, R2, R3] {
	return m.When(func(T1, T2, T3) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker33[T1, T2, T3, R1, R2, R3]) ReturnDefault() {
	m.Return(func() (r1 R1, r2 R2, r3 R3) { return r1, r2, r3 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker33[T1, T2, T3, R1, R2, R3]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker33 is an Invoker implementation for Mocker33.
type Invoker33[T1, T2, T3 any, R1, R2, R3 any] struct {
	*Mocker33[T1, T2, T3, R1, R2, R3]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker33[T1, T2, T3, R1, R2, R3]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker33[T1, T2, T3, R1, R2, R3]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, r2, r3, ok := m.fnHandle(params[0].(T1), params[1].(T2), params[2].(T3))
	return []interface{}{r1, r2, r3}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker33[T1, T2, T3, R1, R2, R3]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1), params[1].(T2), params[2].(T3))
}

// Return provides predefined response and error values.
func (m *Invoker33[T1, T2, T3, R1, R2, R3]) Return(params []interface{}) []interface{} {
	r1, r2, r3 := m.fnReturn()
	return []interface{}{r1, r2, r3}
}

// NewMocker33 creates a new Mocker33 instance.
func NewMocker33[T1, T2, T3 any, R1, R2, R3 any](r *Manager, typ reflect.Type, method string) *Mocker33[T1, T2, T3, R1, R2, R3] {
	m := &Mocker33[T1, T2, T3, R1, R2, R3]{}
	i := &Invoker33[T1, T2, T3, R1, R2, R3]{Mocker33: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker34 ***********************************/

type Mocker34[T1, T2, T3 any, R1, R2, R3, R4 any] struct {
	fnHandle func(T1, T2, T3) (R1, R2, R3, R4, bool)
	fnWhen   func(T1, T2, T3) bool
	fnReturn func() (R1, R2, R3, R4)
}

// Handle sets a custom function to handle requests.
func (m *Mocker34[T1, T2, T3, R1, R2, R3, R4]) Handle(fn func(T1, T2, T3) (R1, R2, R3, R4, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker34[T1, T2, T3, R1, R2, R3, R4]) When(fn func(T1, T2, T3) bool) *Mocker34[T1, T2, T3, R1, R2, R3, R4] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker34[T1, T2, T3, R1, R2, R3, R4]) Return(fn func() (R1, R2, R3, R4)) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker34[T1, T2, T3, R1, R2, R3, R4]) Always() *Mocker34[T1, T2, T3, R1, R2, R3, R4] {
	return m.When(func(T1, T2, T3) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker34[T1, T2, T3, R1, R2, R3, R4]) ReturnDefault() {
	m.Return(func() (r1 R1, r2 R2, r3 R3, r4 R4) { return r1, r2, r3, r4 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker34[T1, T2, T3, R1, R2, R3, R4]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker34 is an Invoker implementation for Mocker34.
type Invoker34[T1, T2, T3 any, R1, R2, R3, R4 any] struct {
	*Mocker34[T1, T2, T3, R1, R2, R3, R4]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker34[T1, T2, T3, R1, R2, R3, R4]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker34[T1, T2, T3, R1, R2, R3, R4]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, r2, r3, r4, ok := m.fnHandle(params[0].(T1), params[1].(T2), params[2].(T3))
	return []interface{}{r1, r2, r3, r4}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker34[T1, T2, T3, R1, R2, R3, R4]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1), params[1].(T2), params[2].(T3))
}

// Return provides predefined response and error values.
func (m *Invoker34[T1, T2, T3, R1, R2, R3, R4]) Return(params []interface{}) []interface{} {
	r1, r2, r3, r4 := m.fnReturn()
	return []interface{}{r1, r2, r3, r4}
}

// NewMocker34 creates a new Mocker34 instance.
func NewMocker34[T1, T2, T3 any, R1, R2, R3, R4 any](r *Manager, typ reflect.Type, method string) *Mocker34[T1, T2, T3, R1, R2, R3, R4] {
	m := &Mocker34[T1, T2, T3, R1, R2, R3, R4]{}
	i := &Invoker34[T1, T2, T3, R1, R2, R3, R4]{Mocker34: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker35 ***********************************/

type Mocker35[T1, T2, T3 any, R1, R2, R3, R4, R5 any] struct {
	fnHandle func(T1, T2, T3) (R1, R2, R3, R4, R5, bool)
	fnWhen   func(T1, T2, T3) bool
	fnReturn func() (R1, R2, R3, R4, R5)
}

// Handle sets a custom function to handle requests.
func (m *Mocker35[T1, T2, T3, R1, R2, R3, R4, R5]) Handle(fn func(T1, T2, T3) (R1, R2, R3, R4, R5, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker35[T1, T2, T3, R1, R2, R3, R4, R5]) When(fn func(T1, T2, T3) bool) *Mocker35[T1, T2, T3, R1, R2, R3, R4, R5] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker35[T1, T2, T3, R1, R2, R3, R4, R5]) Return(fn func() (R1, R2, R3, R4, R5)) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker35[T1, T2, T3, R1, R2, R3, R4, R5]) Always() *Mocker35[T1, T2, T3, R1, R2, R3, R4, R5] {
	return m.When(func(T1, T2, T3) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker35[T1, T2, T3, R1, R2, R3, R4, R5]) ReturnDefault() {
	m.Return(func() (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) { return r1, r2, r3, r4, r5 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker35[T1, T2, T3, R1, R2, R3, R4, R5]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker35 is an Invoker implementation for Mocker35.
type Invoker35[T1, T2, T3 any, R1, R2, R3, R4, R5 any] struct {
	*Mocker35[T1, T2, T3, R1, R2, R3, R4, R5]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker35[T1, T2, T3, R1, R2, R3, R4, R5]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker35[T1, T2, T3, R1, R2, R3, R4, R5]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, r2, r3, r4, r5, ok := m.fnHandle(params[0].(T1), params[1].(T2), params[2].(T3))
	return []interface{}{r1, r2, r3, r4, r5}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker35[T1, T2, T3, R1, R2, R3, R4, R5]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1), params[1].(T2), params[2].(T3))
}

// Return provides predefined response and error values.
func (m *Invoker35[T1, T2, T3, R1, R2, R3, R4, R5]) Return(params []interface{}) []interface{} {
	r1, r2, r3, r4, r5 := m.fnReturn()
	return []interface{}{r1, r2, r3, r4, r5}
}

// NewMocker35 creates a new Mocker35 instance.
func NewMocker35[T1, T2, T3 any, R1, R2, R3, R4, R5 any](r *Manager, typ reflect.Type, method string) *Mocker35[T1, T2, T3, R1, R2, R3, R4, R5] {
	m := &Mocker35[T1, T2, T3, R1, R2, R3, R4, R5]{}
	i := &Invoker35[T1, T2, T3, R1, R2, R3, R4, R5]{Mocker35: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker40 ***********************************/

type Mocker40[T1, T2, T3, T4 any] struct {
	fnHandle func(T1, T2, T3, T4) bool
	fnWhen   func(T1, T2, T3, T4) bool
	fnReturn func()
}

// Handle sets a custom function to handle requests.
func (m *Mocker40[T1, T2, T3, T4]) Handle(fn func(T1, T2, T3, T4) bool) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker40[T1, T2, T3, T4]) When(fn func(T1, T2, T3, T4) bool) *Mocker40[T1, T2, T3, T4] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker40[T1, T2, T3, T4]) Return(fn func()) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker40[T1, T2, T3, T4]) Always() *Mocker40[T1, T2, T3, T4] {
	return m.When(func(T1, T2, T3, T4) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker40[T1, T2, T3, T4]) ReturnDefault() {
	m.Return(func() {})
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker40[T1, T2, T3, T4]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker40 is an Invoker implementation for Mocker40.
type Invoker40[T1, T2, T3, T4 any] struct {
	*Mocker40[T1, T2, T3, T4]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker40[T1, T2, T3, T4]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker40[T1, T2, T3, T4]) Handle(params []interface{}) ([]interface{}, bool) {
	ok := m.fnHandle(params[0].(T1), params[1].(T2), params[2].(T3), params[3].(T4))
	return []interface{}{}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker40[T1, T2, T3, T4]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1), params[1].(T2), params[2].(T3), params[3].(T4))
}

// Return provides predefined response and error values.
func (m *Invoker40[T1, T2, T3, T4]) Return(params []interface{}) []interface{} {
	m.fnReturn()
	return []interface{}{}
}

// NewMocker40 creates a new Mocker40 instance.
func NewMocker40[T1, T2, T3, T4 any](r *Manager, typ reflect.Type, method string) *Mocker40[T1, T2, T3, T4] {
	m := &Mocker40[T1, T2, T3, T4]{}
	i := &Invoker40[T1, T2, T3, T4]{Mocker40: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker41 ***********************************/

type Mocker41[T1, T2, T3, T4 any, R1 any] struct {
	fnHandle func(T1, T2, T3, T4) (R1, bool)
	fnWhen   func(T1, T2, T3, T4) bool
	fnReturn func() R1
}

// Handle sets a custom function to handle requests.
func (m *Mocker41[T1, T2, T3, T4, R1]) Handle(fn func(T1, T2, T3, T4) (R1, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker41[T1, T2, T3, T4, R1]) When(fn func(T1, T2, T3, T4) bool) *Mocker41[T1, T2, T3, T4, R1] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker41[T1, T2, T3, T4, R1]) Return(fn func() R1) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker41[T1, T2, T3, T4, R1]) Always() *Mocker41[T1, T2, T3, T4, R1] {
	return m.When(func(T1, T2, T3, T4) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker41[T1, T2, T3, T4, R1]) ReturnDefault() {
	m.Return(func() (r1 R1) { return r1 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker41[T1, T2, T3, T4, R1]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker41 is an Invoker implementation for Mocker41.
type Invoker41[T1, T2, T3, T4 any, R1 any] struct {
	*Mocker41[T1, T2, T3, T4, R1]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker41[T1, T2, T3, T4, R1]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker41[T1, T2, T3, T4, R1]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, ok := m.fnHandle(params[0].(T1), params[1].(T2), params[2].(T3), params[3].(T4))
	return []interface{}{r1}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker41[T1, T2, T3, T4, R1]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1), params[1].(T2), params[2].(T3), params[3].(T4))
}

// Return provides predefined response and error values.
func (m *Invoker41[T1, T2, T3, T4, R1]) Return(params []interface{}) []interface{} {
	r1 := m.fnReturn()
	return []interface{}{r1}
}

// NewMocker41 creates a new Mocker41 instance.
func NewMocker41[T1, T2, T3, T4 any, R1 any](r *Manager, typ reflect.Type, method string) *Mocker41[T1, T2, T3, T4, R1] {
	m := &Mocker41[T1, T2, T3, T4, R1]{}
	i := &Invoker41[T1, T2, T3, T4, R1]{Mocker41: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker42 ***********************************/

type Mocker42[T1, T2, T3, T4 any, R1, R2 any] struct {
	fnHandle func(T1, T2, T3, T4) (R1, R2, bool)
	fnWhen   func(T1, T2, T3, T4) bool
	fnReturn func() (R1, R2)
}

// Handle sets a custom function to handle requests.
func (m *Mocker42[T1, T2, T3, T4, R1, R2]) Handle(fn func(T1, T2, T3, T4) (R1, R2, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker42[T1, T2, T3, T4, R1, R2]) When(fn func(T1, T2, T3, T4) bool) *Mocker42[T1, T2, T3, T4, R1, R2] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker42[T1, T2, T3, T4, R1, R2]) Return(fn func() (R1, R2)) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker42[T1, T2, T3, T4, R1, R2]) Always() *Mocker42[T1, T2, T3, T4, R1, R2] {
	return m.When(func(T1, T2, T3, T4) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker42[T1, T2, T3, T4, R1, R2]) ReturnDefault() {
	m.Return(func() (r1 R1, r2 R2) { return r1, r2 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker42[T1, T2, T3, T4, R1, R2]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker42 is an Invoker implementation for Mocker42.
type Invoker42[T1, T2, T3, T4 any, R1, R2 any] struct {
	*Mocker42[T1, T2, T3, T4, R1, R2]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker42[T1, T2, T3, T4, R1, R2]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker42[T1, T2, T3, T4, R1, R2]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, r2, ok := m.fnHandle(params[0].(T1), params[1].(T2), params[2].(T3), params[3].(T4))
	return []interface{}{r1, r2}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker42[T1, T2, T3, T4, R1, R2]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1), params[1].(T2), params[2].(T3), params[3].(T4))
}

// Return provides predefined response and error values.
func (m *Invoker42[T1, T2, T3, T4, R1, R2]) Return(params []interface{}) []interface{} {
	r1, r2 := m.fnReturn()
	return []interface{}{r1, r2}
}

// NewMocker42 creates a new Mocker42 instance.
func NewMocker42[T1, T2, T3, T4 any, R1, R2 any](r *Manager, typ reflect.Type, method string) *Mocker42[T1, T2, T3, T4, R1, R2] {
	m := &Mocker42[T1, T2, T3, T4, R1, R2]{}
	i := &Invoker42[T1, T2, T3, T4, R1, R2]{Mocker42: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker43 ***********************************/

type Mocker43[T1, T2, T3, T4 any, R1, R2, R3 any] struct {
	fnHandle func(T1, T2, T3, T4) (R1, R2, R3, bool)
	fnWhen   func(T1, T2, T3, T4) bool
	fnReturn func() (R1, R2, R3)
}

// Handle sets a custom function to handle requests.
func (m *Mocker43[T1, T2, T3, T4, R1, R2, R3]) Handle(fn func(T1, T2, T3, T4) (R1, R2, R3, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker43[T1, T2, T3, T4, R1, R2, R3]) When(fn func(T1, T2, T3, T4) bool) *Mocker43[T1, T2, T3, T4, R1, R2, R3] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker43[T1, T2, T3, T4, R1, R2, R3]) Return(fn func() (R1, R2, R3)) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker43[T1, T2, T3, T4, R1, R2, R3]) Always() *Mocker43[T1, T2, T3, T4, R1, R2, R3] {
	return m.When(func(T1, T2, T3, T4) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker43[T1, T2, T3, T4, R1, R2, R3]) ReturnDefault() {
	m.Return(func() (r1 R1, r2 R2, r3 R3) { return r1, r2, r3 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker43[T1, T2, T3, T4, R1, R2, R3]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker43 is an Invoker implementation for Mocker43.
type Invoker43[T1, T2, T3, T4 any, R1, R2, R3 any] struct {
	*Mocker43[T1, T2, T3, T4, R1, R2, R3]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker43[T1, T2, T3, T4, R1, R2, R3]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker43[T1, T2, T3, T4, R1, R2, R3]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, r2, r3, ok := m.fnHandle(params[0].(T1), params[1].(T2), params[2].(T3), params[3].(T4))
	return []interface{}{r1, r2, r3}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker43[T1, T2, T3, T4, R1, R2, R3]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1), params[1].(T2), params[2].(T3), params[3].(T4))
}

// Return provides predefined response and error values.
func (m *Invoker43[T1, T2, T3, T4, R1, R2, R3]) Return(params []interface{}) []interface{} {
	r1, r2, r3 := m.fnReturn()
	return []interface{}{r1, r2, r3}
}

// NewMocker43 creates a new Mocker43 instance.
func NewMocker43[T1, T2, T3, T4 any, R1, R2, R3 any](r *Manager, typ reflect.Type, method string) *Mocker43[T1, T2, T3, T4, R1, R2, R3] {
	m := &Mocker43[T1, T2, T3, T4, R1, R2, R3]{}
	i := &Invoker43[T1, T2, T3, T4, R1, R2, R3]{Mocker43: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker44 ***********************************/

type Mocker44[T1, T2, T3, T4 any, R1, R2, R3, R4 any] struct {
	fnHandle func(T1, T2, T3, T4) (R1, R2, R3, R4, bool)
	fnWhen   func(T1, T2, T3, T4) bool
	fnReturn func() (R1, R2, R3, R4)
}

// Handle sets a custom function to handle requests.
func (m *Mocker44[T1, T2, T3, T4, R1, R2, R3, R4]) Handle(fn func(T1, T2, T3, T4) (R1, R2, R3, R4, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker44[T1, T2, T3, T4, R1, R2, R3, R4]) When(fn func(T1, T2, T3, T4) bool) *Mocker44[T1, T2, T3, T4, R1, R2, R3, R4] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker44[T1, T2, T3, T4, R1, R2, R3, R4]) Return(fn func() (R1, R2, R3, R4)) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker44[T1, T2, T3, T4, R1, R2, R3, R4]) Always() *Mocker44[T1, T2, T3, T4, R1, R2, R3, R4] {
	return m.When(func(T1, T2, T3, T4) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker44[T1, T2, T3, T4, R1, R2, R3, R4]) ReturnDefault() {
	m.Return(func() (r1 R1, r2 R2, r3 R3, r4 R4) { return r1, r2, r3, r4 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker44[T1, T2, T3, T4, R1, R2, R3, R4]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker44 is an Invoker implementation for Mocker44.
type Invoker44[T1, T2, T3, T4 any, R1, R2, R3, R4 any] struct {
	*Mocker44[T1, T2, T3, T4, R1, R2, R3, R4]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker44[T1, T2, T3, T4, R1, R2, R3, R4]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker44[T1, T2, T3, T4, R1, R2, R3, R4]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, r2, r3, r4, ok := m.fnHandle(params[0].(T1), params[1].(T2), params[2].(T3), params[3].(T4))
	return []interface{}{r1, r2, r3, r4}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker44[T1, T2, T3, T4, R1, R2, R3, R4]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1), params[1].(T2), params[2].(T3), params[3].(T4))
}

// Return provides predefined response and error values.
func (m *Invoker44[T1, T2, T3, T4, R1, R2, R3, R4]) Return(params []interface{}) []interface{} {
	r1, r2, r3, r4 := m.fnReturn()
	return []interface{}{r1, r2, r3, r4}
}

// NewMocker44 creates a new Mocker44 instance.
func NewMocker44[T1, T2, T3, T4 any, R1, R2, R3, R4 any](r *Manager, typ reflect.Type, method string) *Mocker44[T1, T2, T3, T4, R1, R2, R3, R4] {
	m := &Mocker44[T1, T2, T3, T4, R1, R2, R3, R4]{}
	i := &Invoker44[T1, T2, T3, T4, R1, R2, R3, R4]{Mocker44: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker45 ***********************************/

type Mocker45[T1, T2, T3, T4 any, R1, R2, R3, R4, R5 any] struct {
	fnHandle func(T1, T2, T3, T4) (R1, R2, R3, R4, R5, bool)
	fnWhen   func(T1, T2, T3, T4) bool
	fnReturn func() (R1, R2, R3, R4, R5)
}

// Handle sets a custom function to handle requests.
func (m *Mocker45[T1, T2, T3, T4, R1, R2, R3, R4, R5]) Handle(fn func(T1, T2, T3, T4) (R1, R2, R3, R4, R5, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker45[T1, T2, T3, T4, R1, R2, R3, R4, R5]) When(fn func(T1, T2, T3, T4) bool) *Mocker45[T1, T2, T3, T4, R1, R2, R3, R4, R5] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker45[T1, T2, T3, T4, R1, R2, R3, R4, R5]) Return(fn func() (R1, R2, R3, R4, R5)) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker45[T1, T2, T3, T4, R1, R2, R3, R4, R5]) Always() *Mocker45[T1, T2, T3, T4, R1, R2, R3, R4, R5] {
	return m.When(func(T1, T2, T3, T4) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker45[T1, T2, T3, T4, R1, R2, R3, R4, R5]) ReturnDefault() {
	m.Return(func() (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) { return r1, r2, r3, r4, r5 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker45[T1, T2, T3, T4, R1, R2, R3, R4, R5]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker45 is an Invoker implementation for Mocker45.
type Invoker45[T1, T2, T3, T4 any, R1, R2, R3, R4, R5 any] struct {
	*Mocker45[T1, T2, T3, T4, R1, R2, R3, R4, R5]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker45[T1, T2, T3, T4, R1, R2, R3, R4, R5]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker45[T1, T2, T3, T4, R1, R2, R3, R4, R5]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, r2, r3, r4, r5, ok := m.fnHandle(params[0].(T1), params[1].(T2), params[2].(T3), params[3].(T4))
	return []interface{}{r1, r2, r3, r4, r5}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker45[T1, T2, T3, T4, R1, R2, R3, R4, R5]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1), params[1].(T2), params[2].(T3), params[3].(T4))
}

// Return provides predefined response and error values.
func (m *Invoker45[T1, T2, T3, T4, R1, R2, R3, R4, R5]) Return(params []interface{}) []interface{} {
	r1, r2, r3, r4, r5 := m.fnReturn()
	return []interface{}{r1, r2, r3, r4, r5}
}

// NewMocker45 creates a new Mocker45 instance.
func NewMocker45[T1, T2, T3, T4 any, R1, R2, R3, R4, R5 any](r *Manager, typ reflect.Type, method string) *Mocker45[T1, T2, T3, T4, R1, R2, R3, R4, R5] {
	m := &Mocker45[T1, T2, T3, T4, R1, R2, R3, R4, R5]{}
	i := &Invoker45[T1, T2, T3, T4, R1, R2, R3, R4, R5]{Mocker45: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker50 ***********************************/

type Mocker50[T1, T2, T3, T4, T5 any] struct {
	fnHandle func(T1, T2, T3, T4, T5) bool
	fnWhen   func(T1, T2, T3, T4, T5) bool
	fnReturn func()
}

// Handle sets a custom function to handle requests.
func (m *Mocker50[T1, T2, T3, T4, T5]) Handle(fn func(T1, T2, T3, T4, T5) bool) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker50[T1, T2, T3, T4, T5]) When(fn func(T1, T2, T3, T4, T5) bool) *Mocker50[T1, T2, T3, T4, T5] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker50[T1, T2, T3, T4, T5]) Return(fn func()) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker50[T1, T2, T3, T4, T5]) Always() *Mocker50[T1, T2, T3, T4, T5] {
	return m.When(func(T1, T2, T3, T4, T5) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker50[T1, T2, T3, T4, T5]) ReturnDefault() {
	m.Return(func() {})
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker50[T1, T2, T3, T4, T5]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker50 is an Invoker implementation for Mocker50.
type Invoker50[T1, T2, T3, T4, T5 any] struct {
	*Mocker50[T1, T2, T3, T4, T5]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker50[T1, T2, T3, T4, T5]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker50[T1, T2, T3, T4, T5]) Handle(params []interface{}) ([]interface{}, bool) {
	ok := m.fnHandle(params[0].(T1), params[1].(T2), params[2].(T3), params[3].(T4), params[4].(T5))
	return []interface{}{}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker50[T1, T2, T3, T4, T5]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1), params[1].(T2), params[2].(T3), params[3].(T4), params[4].(T5))
}

// Return provides predefined response and error values.
func (m *Invoker50[T1, T2, T3, T4, T5]) Return(params []interface{}) []interface{} {
	m.fnReturn()
	return []interface{}{}
}

// NewMocker50 creates a new Mocker50 instance.
func NewMocker50[T1, T2, T3, T4, T5 any](r *Manager, typ reflect.Type, method string) *Mocker50[T1, T2, T3, T4, T5] {
	m := &Mocker50[T1, T2, T3, T4, T5]{}
	i := &Invoker50[T1, T2, T3, T4, T5]{Mocker50: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker51 ***********************************/

type Mocker51[T1, T2, T3, T4, T5 any, R1 any] struct {
	fnHandle func(T1, T2, T3, T4, T5) (R1, bool)
	fnWhen   func(T1, T2, T3, T4, T5) bool
	fnReturn func() R1
}

// Handle sets a custom function to handle requests.
func (m *Mocker51[T1, T2, T3, T4, T5, R1]) Handle(fn func(T1, T2, T3, T4, T5) (R1, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker51[T1, T2, T3, T4, T5, R1]) When(fn func(T1, T2, T3, T4, T5) bool) *Mocker51[T1, T2, T3, T4, T5, R1] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker51[T1, T2, T3, T4, T5, R1]) Return(fn func() R1) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker51[T1, T2, T3, T4, T5, R1]) Always() *Mocker51[T1, T2, T3, T4, T5, R1] {
	return m.When(func(T1, T2, T3, T4, T5) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker51[T1, T2, T3, T4, T5, R1]) ReturnDefault() {
	m.Return(func() (r1 R1) { return r1 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker51[T1, T2, T3, T4, T5, R1]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker51 is an Invoker implementation for Mocker51.
type Invoker51[T1, T2, T3, T4, T5 any, R1 any] struct {
	*Mocker51[T1, T2, T3, T4, T5, R1]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker51[T1, T2, T3, T4, T5, R1]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker51[T1, T2, T3, T4, T5, R1]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, ok := m.fnHandle(params[0].(T1), params[1].(T2), params[2].(T3), params[3].(T4), params[4].(T5))
	return []interface{}{r1}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker51[T1, T2, T3, T4, T5, R1]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1), params[1].(T2), params[2].(T3), params[3].(T4), params[4].(T5))
}

// Return provides predefined response and error values.
func (m *Invoker51[T1, T2, T3, T4, T5, R1]) Return(params []interface{}) []interface{} {
	r1 := m.fnReturn()
	return []interface{}{r1}
}

// NewMocker51 creates a new Mocker51 instance.
func NewMocker51[T1, T2, T3, T4, T5 any, R1 any](r *Manager, typ reflect.Type, method string) *Mocker51[T1, T2, T3, T4, T5, R1] {
	m := &Mocker51[T1, T2, T3, T4, T5, R1]{}
	i := &Invoker51[T1, T2, T3, T4, T5, R1]{Mocker51: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker52 ***********************************/

type Mocker52[T1, T2, T3, T4, T5 any, R1, R2 any] struct {
	fnHandle func(T1, T2, T3, T4, T5) (R1, R2, bool)
	fnWhen   func(T1, T2, T3, T4, T5) bool
	fnReturn func() (R1, R2)
}

// Handle sets a custom function to handle requests.
func (m *Mocker52[T1, T2, T3, T4, T5, R1, R2]) Handle(fn func(T1, T2, T3, T4, T5) (R1, R2, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker52[T1, T2, T3, T4, T5, R1, R2]) When(fn func(T1, T2, T3, T4, T5) bool) *Mocker52[T1, T2, T3, T4, T5, R1, R2] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker52[T1, T2, T3, T4, T5, R1, R2]) Return(fn func() (R1, R2)) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker52[T1, T2, T3, T4, T5, R1, R2]) Always() *Mocker52[T1, T2, T3, T4, T5, R1, R2] {
	return m.When(func(T1, T2, T3, T4, T5) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker52[T1, T2, T3, T4, T5, R1, R2]) ReturnDefault() {
	m.Return(func() (r1 R1, r2 R2) { return r1, r2 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker52[T1, T2, T3, T4, T5, R1, R2]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker52 is an Invoker implementation for Mocker52.
type Invoker52[T1, T2, T3, T4, T5 any, R1, R2 any] struct {
	*Mocker52[T1, T2, T3, T4, T5, R1, R2]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker52[T1, T2, T3, T4, T5, R1, R2]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker52[T1, T2, T3, T4, T5, R1, R2]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, r2, ok := m.fnHandle(params[0].(T1), params[1].(T2), params[2].(T3), params[3].(T4), params[4].(T5))
	return []interface{}{r1, r2}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker52[T1, T2, T3, T4, T5, R1, R2]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1), params[1].(T2), params[2].(T3), params[3].(T4), params[4].(T5))
}

// Return provides predefined response and error values.
func (m *Invoker52[T1, T2, T3, T4, T5, R1, R2]) Return(params []interface{}) []interface{} {
	r1, r2 := m.fnReturn()
	return []interface{}{r1, r2}
}

// NewMocker52 creates a new Mocker52 instance.
func NewMocker52[T1, T2, T3, T4, T5 any, R1, R2 any](r *Manager, typ reflect.Type, method string) *Mocker52[T1, T2, T3, T4, T5, R1, R2] {
	m := &Mocker52[T1, T2, T3, T4, T5, R1, R2]{}
	i := &Invoker52[T1, T2, T3, T4, T5, R1, R2]{Mocker52: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker53 ***********************************/

type Mocker53[T1, T2, T3, T4, T5 any, R1, R2, R3 any] struct {
	fnHandle func(T1, T2, T3, T4, T5) (R1, R2, R3, bool)
	fnWhen   func(T1, T2, T3, T4, T5) bool
	fnReturn func() (R1, R2, R3)
}

// Handle sets a custom function to handle requests.
func (m *Mocker53[T1, T2, T3, T4, T5, R1, R2, R3]) Handle(fn func(T1, T2, T3, T4, T5) (R1, R2, R3, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker53[T1, T2, T3, T4, T5, R1, R2, R3]) When(fn func(T1, T2, T3, T4, T5) bool) *Mocker53[T1, T2, T3, T4, T5, R1, R2, R3] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker53[T1, T2, T3, T4, T5, R1, R2, R3]) Return(fn func() (R1, R2, R3)) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker53[T1, T2, T3, T4, T5, R1, R2, R3]) Always() *Mocker53[T1, T2, T3, T4, T5, R1, R2, R3] {
	return m.When(func(T1, T2, T3, T4, T5) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker53[T1, T2, T3, T4, T5, R1, R2, R3]) ReturnDefault() {
	m.Return(func() (r1 R1, r2 R2, r3 R3) { return r1, r2, r3 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker53[T1, T2, T3, T4, T5, R1, R2, R3]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker53 is an Invoker implementation for Mocker53.
type Invoker53[T1, T2, T3, T4, T5 any, R1, R2, R3 any] struct {
	*Mocker53[T1, T2, T3, T4, T5, R1, R2, R3]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker53[T1, T2, T3, T4, T5, R1, R2, R3]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker53[T1, T2, T3, T4, T5, R1, R2, R3]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, r2, r3, ok := m.fnHandle(params[0].(T1), params[1].(T2), params[2].(T3), params[3].(T4), params[4].(T5))
	return []interface{}{r1, r2, r3}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker53[T1, T2, T3, T4, T5, R1, R2, R3]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1), params[1].(T2), params[2].(T3), params[3].(T4), params[4].(T5))
}

// Return provides predefined response and error values.
func (m *Invoker53[T1, T2, T3, T4, T5, R1, R2, R3]) Return(params []interface{}) []interface{} {
	r1, r2, r3 := m.fnReturn()
	return []interface{}{r1, r2, r3}
}

// NewMocker53 creates a new Mocker53 instance.
func NewMocker53[T1, T2, T3, T4, T5 any, R1, R2, R3 any](r *Manager, typ reflect.Type, method string) *Mocker53[T1, T2, T3, T4, T5, R1, R2, R3] {
	m := &Mocker53[T1, T2, T3, T4, T5, R1, R2, R3]{}
	i := &Invoker53[T1, T2, T3, T4, T5, R1, R2, R3]{Mocker53: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker54 ***********************************/

type Mocker54[T1, T2, T3, T4, T5 any, R1, R2, R3, R4 any] struct {
	fnHandle func(T1, T2, T3, T4, T5) (R1, R2, R3, R4, bool)
	fnWhen   func(T1, T2, T3, T4, T5) bool
	fnReturn func() (R1, R2, R3, R4)
}

// Handle sets a custom function to handle requests.
func (m *Mocker54[T1, T2, T3, T4, T5, R1, R2, R3, R4]) Handle(fn func(T1, T2, T3, T4, T5) (R1, R2, R3, R4, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker54[T1, T2, T3, T4, T5, R1, R2, R3, R4]) When(fn func(T1, T2, T3, T4, T5) bool) *Mocker54[T1, T2, T3, T4, T5, R1, R2, R3, R4] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker54[T1, T2, T3, T4, T5, R1, R2, R3, R4]) Return(fn func() (R1, R2, R3, R4)) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker54[T1, T2, T3, T4, T5, R1, R2, R3, R4]) Always() *Mocker54[T1, T2, T3, T4, T5, R1, R2, R3, R4] {
	return m.When(func(T1, T2, T3, T4, T5) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker54[T1, T2, T3, T4, T5, R1, R2, R3, R4]) ReturnDefault() {
	m.Return(func() (r1 R1, r2 R2, r3 R3, r4 R4) { return r1, r2, r3, r4 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker54[T1, T2, T3, T4, T5, R1, R2, R3, R4]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker54 is an Invoker implementation for Mocker54.
type Invoker54[T1, T2, T3, T4, T5 any, R1, R2, R3, R4 any] struct {
	*Mocker54[T1, T2, T3, T4, T5, R1, R2, R3, R4]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker54[T1, T2, T3, T4, T5, R1, R2, R3, R4]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker54[T1, T2, T3, T4, T5, R1, R2, R3, R4]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, r2, r3, r4, ok := m.fnHandle(params[0].(T1), params[1].(T2), params[2].(T3), params[3].(T4), params[4].(T5))
	return []interface{}{r1, r2, r3, r4}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker54[T1, T2, T3, T4, T5, R1, R2, R3, R4]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1), params[1].(T2), params[2].(T3), params[3].(T4), params[4].(T5))
}

// Return provides predefined response and error values.
func (m *Invoker54[T1, T2, T3, T4, T5, R1, R2, R3, R4]) Return(params []interface{}) []interface{} {
	r1, r2, r3, r4 := m.fnReturn()
	return []interface{}{r1, r2, r3, r4}
}

// NewMocker54 creates a new Mocker54 instance.
func NewMocker54[T1, T2, T3, T4, T5 any, R1, R2, R3, R4 any](r *Manager, typ reflect.Type, method string) *Mocker54[T1, T2, T3, T4, T5, R1, R2, R3, R4] {
	m := &Mocker54[T1, T2, T3, T4, T5, R1, R2, R3, R4]{}
	i := &Invoker54[T1, T2, T3, T4, T5, R1, R2, R3, R4]{Mocker54: m}
	r.AddMocker(typ, method, i)
	return m
}

/******************************** Mocker55 ***********************************/

type Mocker55[T1, T2, T3, T4, T5 any, R1, R2, R3, R4, R5 any] struct {
	fnHandle func(T1, T2, T3, T4, T5) (R1, R2, R3, R4, R5, bool)
	fnWhen   func(T1, T2, T3, T4, T5) bool
	fnReturn func() (R1, R2, R3, R4, R5)
}

// Handle sets a custom function to handle requests.
func (m *Mocker55[T1, T2, T3, T4, T5, R1, R2, R3, R4, R5]) Handle(fn func(T1, T2, T3, T4, T5) (R1, R2, R3, R4, R5, bool)) {
	m.fnHandle = fn
}

// When sets a condition function that determines if the mock should apply.
func (m *Mocker55[T1, T2, T3, T4, T5, R1, R2, R3, R4, R5]) When(fn func(T1, T2, T3, T4, T5) bool) *Mocker55[T1, T2, T3, T4, T5, R1, R2, R3, R4, R5] {
	m.fnWhen = fn
	return m
}

// Return sets a function that returns predefined values.
func (m *Mocker55[T1, T2, T3, T4, T5, R1, R2, R3, R4, R5]) Return(fn func() (R1, R2, R3, R4, R5)) {
	m.fnReturn = fn
}

// Always sets the condition to always return true, meaning the mock will be applied for any input.
func (m *Mocker55[T1, T2, T3, T4, T5, R1, R2, R3, R4, R5]) Always() *Mocker55[T1, T2, T3, T4, T5, R1, R2, R3, R4, R5] {
	return m.When(func(T1, T2, T3, T4, T5) bool { return true })
}

// ReturnDefault sets a return function that returns zero values for all return types.
func (m *Mocker55[T1, T2, T3, T4, T5, R1, R2, R3, R4, R5]) ReturnDefault() {
	m.Return(func() (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) { return r1, r2, r3, r4, r5 })
}

// Ignore sets the mock to always apply and return default zero values.
func (m *Mocker55[T1, T2, T3, T4, T5, R1, R2, R3, R4, R5]) Ignore() {
	m.Always().ReturnDefault()
}

// Invoker55 is an Invoker implementation for Mocker55.
type Invoker55[T1, T2, T3, T4, T5 any, R1, R2, R3, R4, R5 any] struct {
	*Mocker55[T1, T2, T3, T4, T5, R1, R2, R3, R4, R5]
}

// Mode determines whether the mock operates in Handle mode or WhenReturn mode.
func (m *Invoker55[T1, T2, T3, T4, T5, R1, R2, R3, R4, R5]) Mode() Mode {
	if m.fnHandle != nil {
		return ModeHandle
	}
	return ModeWhenReturn
}

// Handle executes the custom function if set.
func (m *Invoker55[T1, T2, T3, T4, T5, R1, R2, R3, R4, R5]) Handle(params []interface{}) ([]interface{}, bool) {
	r1, r2, r3, r4, r5, ok := m.fnHandle(params[0].(T1), params[1].(T2), params[2].(T3), params[3].(T4), params[4].(T5))
	return []interface{}{r1, r2, r3, r4, r5}, ok
}

// When checks if the condition function evaluates to true.
func (m *Invoker55[T1, T2, T3, T4, T5, R1, R2, R3, R4, R5]) When(params []interface{}) bool {
	if m.fnWhen == nil {
		return false
	}
	return m.fnWhen(params[0].(T1), params[1].(T2), params[2].(T3), params[3].(T4), params[4].(T5))
}

// Return provides predefined response and error values.
func (m *Invoker55[T1, T2, T3, T4, T5, R1, R2, R3, R4, R5]) Return(params []interface{}) []interface{} {
	r1, r2, r3, r4, r5 := m.fnReturn()
	return []interface{}{r1, r2, r3, r4, r5}
}

// NewMocker55 creates a new Mocker55 instance.
func NewMocker55[T1, T2, T3, T4, T5 any, R1, R2, R3, R4, R5 any](r *Manager, typ reflect.Type, method string) *Mocker55[T1, T2, T3, T4, T5, R1, R2, R3, R4, R5] {
	m := &Mocker55[T1, T2, T3, T4, T5, R1, R2, R3, R4, R5]{}
	i := &Invoker55[T1, T2, T3, T4, T5, R1, R2, R3, R4, R5]{Mocker55: m}
	r.AddMocker(typ, method, i)
	return m
}
