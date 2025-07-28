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
	"context"
	"fmt"
	"io"

	exp "github.com/go-spring/gs-mock/example/inner"
)

//go:generate gs-mock -o src_mock.go -i '!RepositoryV2,,GenericService,Service,,Repository'

var _ = fmt.Println

type Response struct {
	Value int
}

type GenericService[T any, R any] interface {
	io.Writer
	M00()
	M01() R
	M10(T)
	M11(T) R
	M02() (R, bool)
	M12(T) (R, bool)
	M22(ctx context.Context, req map[string]T) (*Response, bool)
	Printf(format string, args ...any)
}

type Service interface {
	io.Writer
	M00()
	M01() *Response
	M10(*exp.Request)
	M11(*exp.Request) *Response
	M02() (*Response, bool)
	M12(*exp.Request) (*Response, bool)
	M22(ctx context.Context, req map[string]*exp.Request) (*Response, bool)
	Printf(format string, args ...any)
}
