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

package testdata

import (
	"context"
	"testing"

	"github.com/go-spring/mock"
	"github.com/go-spring/mock/internal/assert"
	"github.com/go-spring/mock/mockgen/testdata/inner"
)

func runService(s Service) {
	_, _ = s.Get(context.Background(), nil, nil)
}

func TestMock(t *testing.T) {
	r, _ := mock.Init(t.Context())
	impl := NewServiceMockImpl(r)
	count := 0
	impl.MockGet().Handle(func(ctx context.Context, req *inner.Request, m map[string]string) (*Response, error, bool) {
		count++
		return nil, nil, count > 1
	})
	assert.Panic(t, func() {
		runService(impl)
	}, "no mock code matched")
	runService(impl)
	assert.Equal(t, count, 2)
}

func TestGenericMock(t *testing.T) {
	r, _ := mock.Init(t.Context())
	impl := NewRepositoryMockImpl[int](r)
	count := 0
	impl.MockSave().Handle(func(item int) (error, bool) {
		count++
		return nil, count > 1
	})
	assert.Panic(t, func() {
		_ = impl.Save(1)
	}, "no mock code matched")
	_ = impl.Save(1)
	assert.Equal(t, count, 2)
}
