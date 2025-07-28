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

package gsmock_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-spring/gs-mock/gsmock"
	"github.com/go-spring/gs-mock/internal/assert"
)

type Trace struct {
	TraceId string
}

type Request struct {
	Token string
}

type Response struct {
	Message string
}

type Client struct{}

var clientType = reflect.TypeFor[Client]()

func (c *Client) Get(ctx context.Context, req *Request, trace *Trace) (*Response, error) {
	if ret, ok := gsmock.InvokeContext(ctx, clientType, "Get", ctx, req, trace); ok {
		return gsmock.Unbox2[*Response, error](ret)
	}
	return &Response{Message: "9:xxx"}, nil
}

// MockGet registers a mock implementation for the Get method.
func MockGet(r *gsmock.Manager) *gsmock.Mocker32[context.Context, *Request, *Trace, *Response, error] {
	return gsmock.NewMocker32[context.Context, *Request, *Trace, *Response, error](r, clientType, "Get")
}

func TestMockWithContext(t *testing.T) {
	var c Client

	// Test case: Unmocked
	{
		resp, err := c.Get(t.Context(), &Request{}, &Trace{})
		assert.Nil(t, err)
		assert.Equal(t, resp.Message, "9:xxx")
	}

	r := gsmock.NewManager()
	ctx := r.BindTo(t.Context())

	// Test case: When && Return
	{
		r.Reset()
		MockGet(r).
			When(func(ctx context.Context, req *Request, trace *Trace) bool {
				return req.Token == "1:abc"
			}).
			Return(func() (resp *Response, err error) {
				return &Response{Message: "1:abc"}, nil
			})

		resp, err := c.Get(ctx, &Request{Token: "1:abc"}, &Trace{})
		assert.Nil(t, err)
		assert.Equal(t, resp.Message, "1:abc")
	}

	// Test case: Handle
	{
		r.Reset()
		MockGet(r).
			Handle(func(ctx context.Context, req *Request, trace *Trace) (resp *Response, err error) {
				return &Response{Message: "4:xyz"}, nil
			})

		resp, err := c.Get(ctx, &Request{Token: "4:xyz"}, &Trace{})
		assert.Nil(t, err)
		assert.Equal(t, resp.Message, "4:xyz")
	}

	// Test case: Invalid Handle
	{
		r.Reset()
		MockGet(r).Handle(nil)

		resp, err := c.Get(ctx, &Request{}, &Trace{})
		assert.Nil(t, err)
		assert.Equal(t, resp.Message, "9:xxx")
	}
}

type ClientInterface interface {
	Query(req *Request, trace *Trace) (*Response, error)
}

// MockClient is a mock implementation of ClientInterface.
type MockClient struct {
	r *gsmock.Manager
}

var mockClientType = reflect.TypeFor[MockClient]()

// NewMockClient creates a new instance of MockClient.
func NewMockClient(r *gsmock.Manager) *MockClient {
	return &MockClient{r}
}

// Query mocks the Query method by invoking a registered mock implementation.
func (c *MockClient) Query(req *Request, trace *Trace) (*Response, error) {
	if ret, ok := gsmock.Invoke(c.r, mockClientType, "Query", req, trace); ok {
		return gsmock.Unbox2[*Response, error](ret)
	}
	panic("mock error")
}

// MockQuery registers a mock implementation for the Query method.
func (c *MockClient) MockQuery() *gsmock.Mocker22[*Request, *Trace, *Response, error] {
	return gsmock.NewMocker22[*Request, *Trace, *Response, error](c.r, mockClientType, "Query")
}

func TestMockNoContext(t *testing.T) {
	r := gsmock.NewManager()

	var c ClientInterface
	mc := NewMockClient(r)
	c = mc

	// Test case: When && Return
	{
		r.Reset()
		mc.MockQuery().
			When(func(req *Request, trace *Trace) bool {
				return req.Token == "1:abc"
			}).
			Return(func() (resp *Response, err error) {
				return &Response{Message: "1:abc"}, nil
			})

		resp, err := c.Query(&Request{Token: "1:abc"}, &Trace{})
		assert.Nil(t, err)
		assert.Equal(t, resp.Message, "1:abc")
	}

	// Test case: Handle
	{
		r.Reset()
		mc.MockQuery().
			Handle(func(req *Request, trace *Trace) (resp *Response, err error) {
				return &Response{Message: "4:xyz"}, nil
			})

		resp, err := c.Query(&Request{Token: "4:xyz"}, &Trace{})
		assert.Nil(t, err)
		assert.Equal(t, resp.Message, "4:xyz")
	}

	// Test case: Invalid Handle
	{
		r.Reset()
		mc.MockQuery().Handle(nil)

		assert.Panic(t, func() {
			_, _ = c.Query(&Request{}, &Trace{})
		}, "mock error")
	}
}
