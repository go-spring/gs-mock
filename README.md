# gs-mock

[English](README.md) | [中文](README_CN.md)

> The project has been officially released, welcome to use!

`gs-mock` is a modern, type-safe Go mock library with full support for generics.
It solves the shortcomings of traditional Go mocking tools in type safety and complexity,
and naturally supports concurrent testing through `context.Context`.
Whether it’s interfaces, regular functions, or struct methods,
everything can be mocked easily—making it ideal for unit tests in microservice projects.

## Features

* **Type-Safe & Generic Support**\
  Supports generic functions and interfaces, with IDE auto-completion for improved development experience
* **Multiple Mocking Modes**
    * `Handle` mode: handle all mock logic in a single callback
    * `When/Return` mode: execute different return logic based on conditions
* **Multiple Parameters & Return Values**\
  Supports up to 5 parameters and 4 return values, enough for most function signatures
* **Interface & Struct Method Mocking**
    * Interfaces: mock code is generated automatically
    * Regular functions & struct methods: mocked via `context.Context`, avoiding the need for interface abstraction
* **Concurrency Test Support**\
  Mock behavior is bound to the context chain, ensuring safe concurrent mocking
* **Simplified API**\
  API is similar to `gomock` but cleaner and easier to use

## Installation

Install directly:

```bash
go install github.com/go-spring/gs-mock@latest
```

Install via the `gs` toolkit:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/go-spring/gs/HEAD/install.sh)"
```

## Quick Start

### Interface Mocking

1. Define an interface

```go
type Service interface {
    Do(n int) int
}
```

2. Generate mock code

```go
//go:generate gs mock -o src_mock.go
```

> Add the above directive at the top of the interface file to generate mock code for all interfaces in the package.\
> If you only need to generate mocks for specific interfaces, use `-i`.
> Prefixing an interface name with `!` means exclusion.

```go
//go:generate gs mock -o src_mock.go -i '!RepositoryV2,Repository'
```

3. Use the generated mock

```go
r := gsmock.NewManager()
s := NewServiceMockImpl(r)

// Handle mode
s.MockDo().Handle(func (impl *ServiceMockImpl, n int) int {
    if n%2 == 0 {
        return n * 2
    }
    return n + 1
})

fmt.Println(s.Do(1)) // 2
fmt.Println(s.Do(2)) // 4
```

```go
r.Reset()

// When/Return mode
s.MockDo().When(func (impl *ServiceMockImpl, n int) bool {
    return n%2 == 0
}).ReturnValue(2)

s.MockDo().When(func (impl *ServiceMockImpl, n int) bool {
    return n%2 == 1
}).ReturnValue(1)

fmt.Println(s.Do(1)) // 1
fmt.Println(s.Do(2)) // 2
```

### Function Mocking

1. Define a regular function

```go
//go:noinline // prevent inline optimization
func Do(ctx context.Context, n int) int { return n }
```

2. Mock the function

```go
r := gsmock.NewManager()
ctx := r.WithManager(context.TODO())

gsmock.Mock21(Do, r).Handle(func (ctx context.Context, n int) int {
    return n * 2
})

fmt.Println(Do(ctx, 1)) // 2
```

### Method Mocking

1. Define a struct

```go
type Service struct{ m int }
func (s *Service) Do(ctx context.Context, n int) int { return n }
```

2. Mock the struct method

```go
r := gsmock.NewManager()
ctx := r.WithManager(context.TODO())

gsmock.Mock31((*Service).Do, r).Handle(func (s *Service, ctx context.Context, n int) int {
    return n + s.m
})

fmt.Println((&Service{m: 1}).Do(ctx, 1)) // 2
fmt.Println((&Service{m: 2}).Do(ctx, 1)) // 3
```

> ⚠️ When running `go test`, add `-gcflags="all=-N -l"` to prevent method inlining.

## License

This project is licensed under the Apache License Version 2.0.
