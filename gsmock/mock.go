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

package gsmock

import (
	"context"
	"fmt"
	"reflect"
	"runtime"
)

// Mode represents the mocking mode of an Invoker.
type Mode int

const (
	// ModeHandle indicates that the Invoker uses a custom Handle function.
	ModeHandle = Mode(iota)
	// ModeWhenReturn indicates that the Invoker uses a When + Return mechanism.
	ModeWhenReturn
)

// Invoker defines the interface that all mock implementations must satisfy.
type Invoker interface {
	// Mode returns the mocking mode of the Invoker.
	Mode() Mode
	// When determines if the current mock applies to the given parameters.
	When(params []any) bool
	// Return returns the mock values.
	Return() []any
	// Handle executes the custom handler function and returns its results.
	Handle(params []any) []any
}

// Manager manages a collection of mockers for top-level functions,
// using function identifiers as map keys.
type Manager struct {
	mockers map[string][]Invoker
}

// NewManager creates and returns a new Manager instance.
func NewManager() *Manager {
	return &Manager{
		mockers: make(map[string][]Invoker),
	}
}

// Reset clears all mockers in the Manager.
func (r *Manager) Reset() {
	r.mockers = make(map[string][]Invoker)
}

var managerKey int

// WithManager returns a new context with the Manager attached.
func (r *Manager) WithManager(ctx context.Context) context.Context {
	return context.WithValue(ctx, &managerKey, r)
}

// GetFuncID returns a unique identifier string for a function.
func GetFuncID(f any) string {
	v := reflect.ValueOf(f)
	e := runtime.FuncForPC(v.Pointer())
	return e.Name() + ":" + v.Type().String()
}

// addMocker registers a new Invoker for a specific function.
func (r *Manager) addMocker(f any, i Invoker) {
	k := GetFuncID(f)
	r.mockers[k] = append(r.mockers[k], i)
}

// Invoke searches for a matching Invoker for the given function and parameters,
// executes it based on its mocking mode, and returns the result slice along with
// a boolean indicating whether a mock was applied.
// The function `f` is usually a top-level function or a method with a receiver.
// It cannot be an instance method.
// If `f` is a method with a receiver, the first element of `params` must be the receiver.
func Invoke(r *Manager, f any, params ...any) ([]any, bool) {
	k := GetFuncID(f)
	var defaultHandler Invoker
	for _, m := range r.mockers[k] {
		switch m.Mode() {
		case ModeHandle:
			// Panic if multiple Handle mocks exist for the same function.
			if defaultHandler != nil {
				panic(fmt.Sprintf("found multiple Handle functions for %s", k))
			}
			defaultHandler = m
		case ModeWhenReturn:
			if m.When(params) {
				ret := m.Return()
				return ret, true
			}
		default: // for linter
		}
	}
	// Execute the Handle function if available.
	if defaultHandler != nil {
		return defaultHandler.Handle(params), true
	}
	return nil, false
}

// InvokeContext retrieves the Manager from the context and invokes the mock.
// The function `f` is usually a top-level function or a method with a receiver.
// It cannot be an instance method.
// If `f` is a method with a receiver, the first element of `params` must be the receiver.
func InvokeContext(ctx context.Context, f any, params ...any) ([]any, bool) {
	if r, ok := ctx.Value(&managerKey).(*Manager); ok {
		return Invoke(r, f, params...)
	}
	return nil, false
}

// Unbox1 extracts a single return value from a slice of interfaces.
func Unbox1[R1 any](ret []any) (r1 R1) {
	if len(ret) == 1 {
		r1, _ = ret[0].(R1)
	} else {
		panic(fmt.Sprintf("expected 1 return value, but got %d", len(ret)))
	}
	return
}

// Unbox2 extracts two return values from a slice of interfaces.
func Unbox2[R1, R2 any](ret []any) (r1 R1, r2 R2) {
	if len(ret) == 2 {
		r1, _ = ret[0].(R1)
		r2, _ = ret[1].(R2)
	} else {
		panic(fmt.Sprintf("expected 2 return values, but got %d", len(ret)))
	}
	return
}

// Unbox3 extracts three return values from a slice of interfaces.
func Unbox3[R1, R2, R3 any](ret []any) (r1 R1, r2 R2, r3 R3) {
	if len(ret) == 3 {
		r1, _ = ret[0].(R1)
		r2, _ = ret[1].(R2)
		r3, _ = ret[2].(R3)
	} else {
		panic(fmt.Sprintf("expected 3 return values, but got %d", len(ret)))
	}
	return
}

// Unbox4 extracts four return values from a slice of interfaces.
func Unbox4[R1, R2, R3, R4 any](ret []any) (r1 R1, r2 R2, r3 R3, r4 R4) {
	if len(ret) == 4 {
		r1, _ = ret[0].(R1)
		r2, _ = ret[1].(R2)
		r3, _ = ret[2].(R3)
		r4, _ = ret[3].(R4)
	} else {
		panic(fmt.Sprintf("expected 4 return values, but got %d", len(ret)))
	}
	return
}

// Unbox5 extracts five return values from a slice of interfaces.
func Unbox5[R1, R2, R3, R4, R5 any](ret []any) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
	if len(ret) == 5 {
		r1, _ = ret[0].(R1)
		r2, _ = ret[1].(R2)
		r3, _ = ret[2].(R3)
		r4, _ = ret[3].(R4)
		r5, _ = ret[4].(R5)
	} else {
		panic(fmt.Sprintf("expected 5 return values, but got %d", len(ret)))
	}
	return
}
