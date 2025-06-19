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
	"context"
	"log"
	"reflect"
	"testing"
)

// Mode represents the mocking mode.
type Mode int

const (
	ModeHandle = Mode(iota)
	ModeWhenReturn
)

// Invoker defines the interface that all mock implementations must satisfy.
type Invoker interface {
	// Mode returns the mocking mode
	Mode() Mode
	// When determines if the current mock applies to the given parameters
	When(params []interface{}) bool
	// Return returns mock values
	Return(params []interface{}) []interface{}
	// Handle handles the function call and indicates if it was handled
	Handle(params []interface{}) ([]interface{}, bool)
}

// mockerKey is used as a key in the map to identify mockers by type and method.
type mockerKey struct {
	typ    reflect.Type
	method string
}

// Manager manages a collection of mockers for different types and methods.
type Manager struct {
	mockers map[mockerKey][]Invoker
}

// NewManager creates and returns a new Manager instance.
func NewManager() *Manager {
	return &Manager{
		mockers: make(map[mockerKey][]Invoker),
	}
}

var managerKey int

// getManager retrieves the Manager instance from the context.
func getManager(ctx context.Context) *Manager {
	if r, ok := ctx.Value(&managerKey).(*Manager); ok {
		return r
	}
	return nil
}

// BindTo returns a new context with the Manager attached to it.
func (r *Manager) BindTo(ctx context.Context) context.Context {
	return context.WithValue(ctx, &managerKey, r)
}

// getMockers retrieves all mockers for a given type and method.
func (r *Manager) getMockers(typ reflect.Type, method string) []Invoker {
	return r.mockers[mockerKey{typ, method}]
}

// addMocker adds a new mocker for a specific type and method.
func (r *Manager) addMocker(typ reflect.Type, method string, i Invoker) {
	k := mockerKey{typ, method}
	r.mockers[k] = append(r.mockers[k], i)
}

// Invoke finds a matching Invoker and calls it based on the mocking mode.
func Invoke(r *Manager, typ reflect.Type, method string, params ...interface{}) ([]interface{}, bool) {
	if r == nil || !testing.Testing() {
		return nil, false
	}
	mockers := r.getMockers(typ, method)
	for _, f := range mockers {
		switch f.Mode() {
		case ModeHandle:
			ret, ok := f.Handle(params)
			if ok {
				return ret, true
			}
		case ModeWhenReturn:
			if f.When(params) {
				ret := f.Return(params)
				return ret, true
			}
		default: // for linter
		}
	}
	return nil, false
}

// InvokeContext is a convenience function that invokes a mock using context to retrieve the Manager.
func InvokeContext(ctx context.Context, typ reflect.Type, method string, params ...interface{}) ([]interface{}, bool) {
	if !testing.Testing() {
		return nil, false
	}
	return Invoke(getManager(ctx), typ, method, params...)
}

// Unbox1 extracts a single return value from a slice of interfaces.
func Unbox1[R1 any](ret []interface{}) (r1 R1) {
	if len(ret) == 1 {
		r1, _ = ret[0].(R1)
	} else {
		log.Printf("Warning: expected 1 return value, but got %d", len(ret))
	}
	return
}

// Unbox2 extracts two return values from a slice of interfaces.
func Unbox2[R1, R2 any](ret []interface{}) (r1 R1, r2 R2) {
	if len(ret) == 2 {
		r1, _ = ret[0].(R1)
		r2, _ = ret[1].(R2)
	} else {
		log.Printf("Warning: expected 2 return values, but got %d", len(ret))
	}
	return
}

// Unbox3 extracts three return values from a slice of interfaces.
func Unbox3[R1, R2, R3 any](ret []interface{}) (r1 R1, r2 R2, r3 R3) {
	if len(ret) == 3 {
		r1, _ = ret[0].(R1)
		r2, _ = ret[1].(R2)
		r3, _ = ret[2].(R3)
	} else {
		log.Printf("Warning: expected 3 return values, but got %d", len(ret))
	}
	return
}

// Unbox4 extracts four return values from a slice of interfaces.
func Unbox4[R1, R2, R3, R4 any](ret []interface{}) (r1 R1, r2 R2, r3 R3, r4 R4) {
	if len(ret) == 4 {
		r1, _ = ret[0].(R1)
		r2, _ = ret[1].(R2)
		r3, _ = ret[2].(R3)
		r4, _ = ret[3].(R4)
	} else {
		log.Printf("Warning: expected 4 return values, but got %d", len(ret))
	}
	return
}

// Unbox5 extracts five return values from a slice of interfaces.
func Unbox5[R1, R2, R3, R4, R5 any](ret []interface{}) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
	if len(ret) == 5 {
		r1, _ = ret[0].(R1)
		r2, _ = ret[1].(R2)
		r3, _ = ret[2].(R3)
		r4, _ = ret[3].(R4)
		r5, _ = ret[4].(R5)
	} else {
		log.Printf("Warning: expected 5 return values, but got %d", len(ret))
	}
	return
}
