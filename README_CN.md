# mock

[English](README.md) | [中文](README_CN.md)

mock 是一个现代化的、类型安全的 Go 语言 mocking 库，完全支持泛型编程。它提供了简单易用的接口，
可以帮助开发者轻松创建和管理模拟对象，从而提高单元测试的质量和效率。该库旨在解决 Go 语言中传统
mocking 工具存在的类型安全性不足和使用复杂性问题。

## 主要特性

- **类型安全**：利用 Go 1.18+ 的泛型特性，确保编译时的安全性，避免运行时类型错误
- **多种 Mock 模式**：
    - `Handle` 模式：直接处理函数调用
    - `When/Return` 模式：基于条件的模拟返回
- **灵活的方法匹配**：支持不同数量和类型的参数及返回值（最多支持5个参数和5个返回值）
- **上下文支持**：提供与 context 包的集成，方便在分布式系统中进行测试
- **自动重置功能**：Manager 提供 Reset 方法，可轻松重置所有模拟器到初始状态
- **详细的错误信息**：当没有匹配的 mock 代码或存在多个匹配时，提供清晰的错误提示

## 安装工具

**gs-mock** 是一个用于生成 Go mock 代码的工具，你可以通过以下方式安装它：

```bash
go install github.com/go-spring/gs-mock@latest
```

### 基本用法

1. 定义接口

首先，在你的项目中定义需要 mock 的接口。例如，创建一个名为 service.go 的文件，并添加如下代码：

```go
package main

type Service interface {
	Save(r1, r2, r3, r4, r5, r6 int)
}
```

2. 生成 Mock 代码

然后在 service.go 文件中加入 go:generate 指令，即可生成 mock 代码：

```go
//go:generate gs mock
```

你需要指定一个输出文件名，例如 service_mock.go，否则会输出到控制台上。

```go
//go:generate gs mock -o src_mock.go
```

你还可以指定哪些接口生成 mock，哪些接口不生成 mock (在接口名前面加!即可)。

```go
//go:generate gs mock -o src_mock.go -i '!RepositoryV2,Repository'
```

## 使用示例

以下是一个简单的使用示例：

```go
package mock_test

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
```

## 许可证

本项目采用 Apache License Version 2.0 许可证。
