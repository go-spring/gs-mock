# gs-mock

[English](README.md) | [中文](README_CN.md)

> 该项目已经正式发布，欢迎使用！

`gs-mock` 是一个现代、类型安全的 Go mock 库，全面支持泛型。
它解决了传统 Go mock 工具在类型安全和使用复杂性上的不足，
同时通过 `context.Context` 天然支持并发测试。
无论接口、普通函数还是结构体方法，都可以轻松 mock，
非常适合微服务项目的单元测试。

## 特性

* **类型安全 & 泛型支持**\
  支持泛型函数和接口，IDE 提供自动类型补全，提高开发效率
* **多种 Mock 模式**
    * `Handle` 模式：在单个回调中处理所有 mock 逻辑
    * `When/Return` 模式：根据条件执行对应的返回逻辑
* **多参数与多返回值**\
  支持最多 5 个参数和 4 个返回值，满足绝大多数函数签名需求
* **接口与结构体方法 Mock**
    * 接口：通过代码生成自动生成 mock
    * 普通函数 & 结构体方法：通过 `context.Context` 传递 mock，避免使用接口
* **并发测试支持**\
  通过 `context.Context` 链路确保并发场景下的 mock 安全
* **简化 API 体验**\
  与官方 `gomock` 用法类似，同时 API 更加简洁易用

## 安装

单独安装：

```bash
go install github.com/go-spring/gs-mock@latest
```

通过 `gs` 工具集安装：

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/go-spring/gs/HEAD/install.sh)"
```

## 快速开始

### 接口 Mock

1. 定义接口

```go
type Service interface {
    Do(n int, s string) (int, error)
    Format(s string, args ...any) string
}
```

2. 生成 Mock 代码

```go
//go:generate gs mock -o src_mock.go
```

> 在接口文件开头添加上述指令，即可为当前目录（包）下的所有接口生成 mock 代码。\
> 如果只需要为特定接口生成，可以使用 `-i` 参数，接口名前加 `!` 表示排除。

```go
//go:generate gs mock -o src_mock.go -i '!RepositoryV2,Repository'
```

3. 使用 Mock 代码

```go
r := gsmock.NewManager()
s := NewServiceMockImpl(r)

// Handle 模式
s.MockDo().Handle(func(impl *ServiceMockImpl, n int, s string) (int, error) {
    if n%2 == 0 {
        return n * 2, nil
    }
    return n + 1, errors.New("error")
})

fmt.Println(s.Do(1, "abc")) // 2 error
fmt.Println(s.Do(2, "abc")) // 4 <nil>
```

```go
r := gsmock.NewManager()
s := NewServiceMockImpl(r)

// When/Return 模式
s.MockFormat().When(func(impl *ServiceMockImpl, s string, args []any) bool {
    return args[0] == "abc"
}).ReturnValue("abc")

// When/Return 模式
s.MockFormat().When(func(impl *ServiceMockImpl, s string, args []any) bool {
    return args[0] == "123"
}).ReturnValue("123")

fmt.Println(s.Format("", "abc", "123")) // abc
fmt.Println(s.Format("", "123", "abc")) // 123
```

### 函数 Mock

1. 定义普通函数

```go
//go:noinline // 防止函数被内联
func Do(ctx context.Context, n int) int { return n }
```

2. mock 普通函数

```go
r := gsmock.NewManager()
ctx := r.WithManager(context.TODO())

gsmock.Mock21(Do, r).Handle(func(ctx context.Context, n int) int {
    return n * 2
})

fmt.Println(Do(ctx, 1)) // 2
```

### 方法 Mock

1. 创建一个结构体

```go
type Service struct{ m int }
func (s *Service) Do(ctx context.Context, n int) int { return n }
```

2. mock 结构体方法

```go
r := gsmock.NewManager()
ctx := r.WithManager(context.TODO())

gsmock.Mock31((*Service).Do, r).Handle(func(s *Service, ctx context.Context, n int) int {
    return n + s.m
})

fmt.Println((&Service{m: 1}).Do(ctx, 1)) // 2
fmt.Println((&Service{m: 2}).Do(ctx, 1)) // 3
```

> ⚠️ 执行 go test 时添加 `-gcflags="all=-N -l"`，防止方法被内联优化。

## 许可证

本项目采用 Apache License Version 2.0 许可证。
