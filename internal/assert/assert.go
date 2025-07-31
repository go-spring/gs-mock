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

package assert

import (
	"fmt"
	"reflect"
	"regexp"
)

// T is the minimum interface of *testing.T.
type T interface {
	Helper()
	Error(args ...interface{})
}

// isNil reports v is nil, but will not panic.
func isNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Interface,
		reflect.Map,
		reflect.Ptr,
		reflect.Slice,
		reflect.UnsafePointer:
		return v.IsNil()
	default:
		return !v.IsValid()
	}
}

// Nil assertion failed when got is not nil.
func Nil(t T, got interface{}) {
	t.Helper()
	// Why can't we use got==nil to judge？Because if
	// a := (*int)(nil)        // %T == *int
	// b := (interface{})(nil) // %T == <nil>
	// then a==b is false, because they are different types.
	if !isNil(reflect.ValueOf(got)) {
		t.Error(fmt.Sprintf("got (%T) %v but expect nil", got, got))
	}
}

// Equal assertion failed when got and expect are not `deeply equal`.
func Equal(t T, got interface{}, expect interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, expect) {
		t.Error(fmt.Sprintf("got (%T) %v but expect (%T) %v", got, got, expect, expect))
	}
}

func recovery(fn func()) (str string) {
	defer func() {
		if r := recover(); r != nil {
			str = fmt.Sprint(r)
		}
	}()
	fn()
	return "<<SUCCESS>>"
}

// Panic assertion failed when fn doesn't panic or not match expr expression.
func Panic(t T, fn func(), expr string) {
	t.Helper()
	str := recovery(fn)
	if str == "<<SUCCESS>>" {
		t.Error("did not panic")
	} else {
		matches(t, str, expr)
	}
}

func matches(t T, got string, expr string) {
	t.Helper()
	if ok, err := regexp.MatchString(expr, got); err != nil {
		t.Error("invalid pattern")
	} else if !ok {
		t.Error(fmt.Sprintf("got %q which does not match %q", got, expr))
	}
}
