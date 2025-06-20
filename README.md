# mock

[English](README.md) | [中文](README_CN.md)

**mock** is a modern, type-safe mocking library for the Go programming language, fully supporting generic programming.
It provides a simple and easy-to-use interface that helps developers easily create and manage mock objects, thereby
improving the quality and efficiency of unit testing. The library is designed to address the lack of type safety and the
complexity of traditional mocking tools in Go.

## Key Features

* **Type Safety**: Utilizes Go 1.18+ generics to ensure compile-time safety and avoid runtime type errors
* **Multiple Mocking Modes**:
    * `Handle` Mode: Directly handle function calls
    * `When/Return` Mode: Conditional mock returns
* **Flexible Method Matching**: Supports different numbers and types of parameters and return values (up to 5 parameters
  and 5 return values)
* **Context Support**: Provides integration with the `context` package, making it easier to test in distributed systems
* **Auto Reset Functionality**: The `Manager` provides a `Reset` method to easily reset all mockers to their initial
  state
* **Detailed Error Messages**: Offers clear error prompts when no matching mock code is found or when multiple matches
  exist

## Installation Tool

**gsmock** is a tool used to generate Go mock code. You can install it with the following command:

```bash
go install github.com/go-spring/mock/gsmock@latest
```

### Basic Usage

1. **Define an Interface**

First, define the interface you want to mock in your project. For example, create a file named `service.go` and add the
following code:

```go
package main

type Service interface {
	Save(r1, r2, r3, r4, r5, r6 int)
}
```

2. **Generate Mock Code**

Then, add a `go:generate` directive to the `service.go` file to generate the mock code:

```go
//go:generate gsmock
```

You need to specify an output filename, such as `service_mock.go`, otherwise the output will be printed to the console.

```go
//go:generate gsmock -o src_mock.go
```

You can also specify which interfaces to generate mocks for and which to exclude (prefix the interface name with `!` to
exclude it).

```go
//go:generate gsmock -o src_mock.go -i '!RepositoryV2,Repository'
```

## Usage Example

Below is a simple usage example:

```go
package mock_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-spring/mock"
	"github.com/go-spring/mock/internal/assert"
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
	if ret, ok := mock.InvokeContext(ctx, clientType, "Get", ctx, req, trace); ok {
		return mock.Unbox2[*Response, error](ret)
	}
	return &Response{Message: "9:xxx"}, nil
}

// MockGet registers a mock implementation for the Get method.
func MockGet(r *mock.Manager) *mock.Mocker32[context.Context, *Request, *Trace, *Response, error] {
	return mock.NewMocker32[context.Context, *Request, *Trace, *Response, error](r, clientType, "Get")
}

func TestMockWithContext(t *testing.T) {
	var c Client

	// Test case: Unmocked
	{
		resp, err := c.Get(t.Context(), &Request{}, &Trace{})
		assert.Nil(t, err)
		assert.Equal(t, resp.Message, "9:xxx")
	}

	r := mock.NewManager()
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
	r *mock.Manager
}

var mockClientType = reflect.TypeFor[MockClient]()

// NewMockClient creates a new instance of MockClient.
func NewMockClient(r *mock.Manager) *MockClient {
	return &MockClient{r}
}

// Query mocks the Query method by invoking a registered mock implementation.
func (c *MockClient) Query(req *Request, trace *Trace) (*Response, error) {
	if ret, ok := mock.Invoke(c.r, mockClientType, "Query", req, trace); ok {
		return mock.Unbox2[*Response, error](ret)
	}
	panic("mock error")
}

// MockQuery registers a mock implementation for the Query method.
func (c *MockClient) MockQuery() *mock.Mocker22[*Request, *Trace, *Response, error] {
	return mock.NewMocker22[*Request, *Trace, *Response, error](c.r, mockClientType, "Query")
}

func TestMockNoContext(t *testing.T) {
	r := mock.NewManager()

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
```

## License

This project is licensed under the Apache License Version 2.0.
