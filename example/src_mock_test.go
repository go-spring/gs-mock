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

package example

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	exp "github.com/go-spring/gs-mock/example/inner"
	"github.com/go-spring/gs-mock/gsmock"
	"github.com/go-spring/gs-mock/internal/assert"
)

type ItemType int

func TestRepositoryMockImpl_FindByID(t *testing.T) {
	s := NewRepositoryMockImpl[ItemType](gsmock.NewManager())

	assert.Panic(t, func() {
		_, _ = s.FindByID("1")
	}, "no mock code matched")

	s.MockFindByID().Handle(func(s string) (ItemType, error) {
		return ItemType(666), nil
	})

	v, err := s.FindByID("1")
	assert.Nil(t, err)
	assert.Equal(t, v, ItemType(666))

	s.MockFindByID().Handle(func(s string) (ItemType, error) {
		return ItemType(666), nil
	})

	assert.Panic(t, func() {
		_, _ = s.FindByID("1")
	}, "found multiple Handle functions for .*FindByID")
}

func TestRepositoryMockImpl_Save(t *testing.T) {
	s := NewRepositoryMockImpl[ItemType](gsmock.NewManager())

	assert.Panic(t, func() {
		_ = s.Save(ItemType(666))
	}, "no mock code matched")

	s.MockSave().Handle(func(v ItemType) error {
		return nil
	})

	err := s.Save(ItemType(666))
	assert.Nil(t, err)
}

func TestGenericServiceMockImpl_M00(t *testing.T) {
	s := NewGenericServiceMockImpl[string, int](gsmock.NewManager())

	assert.Panic(t, func() {
		s.M00()
	}, "no mock code matched")

	s.MockM00().Ignore()

	s.M00()
}

func TestGenericServiceMockImpl_M01(t *testing.T) {
	s := NewGenericServiceMockImpl[string, int](gsmock.NewManager())

	assert.Panic(t, func() {
		s.M01()
	}, "no mock code matched")

	s.MockM01().Handle(func() int {
		return 5
	})

	resp := s.M01()
	assert.Equal(t, resp, 5)
}

func TestGenericServiceMockImpl_M10(t *testing.T) {
	s := NewGenericServiceMockImpl[string, int](gsmock.NewManager())

	assert.Panic(t, func() {
		s.M10("")
	}, "no mock code matched")

	s.MockM10().Ignore()

	s.M10("abc")
}

func TestGenericServiceMockImpl_M11(t *testing.T) {
	s := NewGenericServiceMockImpl[string, int](gsmock.NewManager())

	assert.Panic(t, func() {
		s.M11("")
	}, "no mock code matched")

	s.MockM11().Handle(func(s string) int {
		return 5
	})

	resp := s.M11("abc")
	assert.Equal(t, resp, 5)
}

func TestGenericServiceMockImpl_M02(t *testing.T) {
	s := NewGenericServiceMockImpl[string, int](gsmock.NewManager())

	assert.Panic(t, func() {
		s.M02()
	}, "no mock code matched")

	s.MockM02().Handle(func() (int, bool) {
		return 5, false
	})

	resp, ok := s.M02()
	assert.Equal(t, ok, false)
	assert.Equal(t, resp, 5)
}

func TestGenericServiceMockImpl_M12(t *testing.T) {
	s := NewGenericServiceMockImpl[string, int](gsmock.NewManager())

	assert.Panic(t, func() {
		s.M12("")
	}, "no mock code matched")

	s.MockM12().Handle(func(s string) (int, bool) {
		return 5, false
	})

	resp, ok := s.M12("abc")
	assert.Equal(t, ok, false)
	assert.Equal(t, resp, 5)
}

func TestGenericServiceMockImpl_M22(t *testing.T) {
	r := gsmock.NewManager()
	s := NewGenericServiceMockImpl[string, int](r)
	ctx := r.BindTo(t.Context())

	assert.Panic(t, func() {
		s.M22(ctx, map[string]string{})
	}, "no mock code matched")

	s.MockM22().Handle(func(ctx context.Context, m map[string]string) (*Response, bool) {
		return &Response{Value: 5}, false
	})

	resp, ok := s.M22(ctx, map[string]string{})
	assert.Equal(t, ok, false)
	assert.Equal(t, resp.Value, 5)
}

func TestGenericServiceMockImpl_Print(t *testing.T) {
	s := NewGenericServiceMockImpl[string, int](gsmock.NewManager())

	assert.Panic(t, func() {
		s.Printf("%s\n", "123")
	}, "no mock code matched")

	var buf bytes.Buffer
	s.MockPrintf().Handle(func(format string, args []any) {
		buf.WriteString(fmt.Sprintf(format, args...))
	})

	s.Printf("%s\n", "123")
	assert.Equal(t, buf.String(), "123\n")
}

func TestServiceMockImpl_M00(t *testing.T) {
	s := NewServiceMockImpl(gsmock.NewManager())

	assert.Panic(t, func() {
		s.M00()
	}, "no mock code matched")

	s.MockM00().Ignore()

	s.M00()
}

func TestServiceMockImpl_M01(t *testing.T) {
	s := NewServiceMockImpl(gsmock.NewManager())

	assert.Panic(t, func() {
		s.M01()
	}, "no mock code matched")

	s.MockM01().Handle(func() *Response {
		return &Response{Value: 5}
	})

	resp := s.M01()
	assert.Equal(t, resp.Value, 5)
}

func TestServiceMockImpl_M10(t *testing.T) {
	s := NewServiceMockImpl(gsmock.NewManager())

	assert.Panic(t, func() {
		s.M10(&exp.Request{})
	}, "no mock code matched")

	s.MockM10().Ignore()

	s.M10(&exp.Request{})
}

func TestServiceMockImpl_M11(t *testing.T) {
	s := NewServiceMockImpl(gsmock.NewManager())

	assert.Panic(t, func() {
		s.M11(&exp.Request{})
	}, "no mock code matched")

	s.MockM11().Handle(func(req *exp.Request) *Response {
		return &Response{Value: 5}
	})

	resp := s.M11(&exp.Request{})
	assert.Equal(t, resp.Value, 5)
}

func TestServiceMockImpl_M02(t *testing.T) {
	s := NewServiceMockImpl(gsmock.NewManager())

	assert.Panic(t, func() {
		s.M02()
	}, "no mock code matched")

	s.MockM02().Handle(func() (*Response, bool) {
		return &Response{Value: 5}, false
	})

	resp, ok := s.M02()
	assert.Equal(t, ok, false)
	assert.Equal(t, resp.Value, 5)
}

func TestServiceMockImpl_M12(t *testing.T) {
	s := NewServiceMockImpl(gsmock.NewManager())

	assert.Panic(t, func() {
		s.M12(&exp.Request{})
	}, "no mock code matched")

	s.MockM12().Handle(func(req *exp.Request) (*Response, bool) {
		return &Response{Value: 5}, false
	})

	resp, ok := s.M12(&exp.Request{})
	assert.Equal(t, ok, false)
	assert.Equal(t, resp.Value, 5)
}

func TestServiceMockImpl_M22(t *testing.T) {
	r := gsmock.NewManager()
	s := NewServiceMockImpl(r)
	ctx := r.BindTo(t.Context())

	assert.Panic(t, func() {
		s.M22(ctx, map[string]*exp.Request{})
	}, "no mock code matched")

	s.MockM22().Handle(func(ctx context.Context, m map[string]*exp.Request) (*Response, bool) {
		return &Response{Value: 5}, false
	})

	resp, ok := s.M22(ctx, map[string]*exp.Request{})
	assert.Equal(t, ok, false)
	assert.Equal(t, resp.Value, 5)
}

func TestServiceMockImpl_Print(t *testing.T) {
	s := NewServiceMockImpl(gsmock.NewManager())

	assert.Panic(t, func() {
		s.Printf("%s\n", "123")
	}, "no mock code matched")

	var buf bytes.Buffer
	s.MockPrintf().Handle(func(format string, args []any) {
		buf.WriteString(fmt.Sprintf(format, args...))
	})

	s.Printf("%s\n", "123")
	assert.Equal(t, buf.String(), "123\n")
}

func TestServiceMockImpl_Writer(t *testing.T) {
	s := NewServiceMockImpl(gsmock.NewManager())

	assert.Panic(t, func() {
		_, _ = s.Write([]byte("123"))
	}, "runtime error: invalid memory address or nil pointer dereference")

	buf := bytes.NewBuffer(nil)
	s.Writer = buf

	buf.Reset()
	_, _ = s.Write([]byte("abc"))
	assert.Equal(t, buf.String(), "abc")

	buf.Reset()
	_, _ = s.Write([]byte("123"))
	assert.Equal(t, buf.String(), "123")
}
