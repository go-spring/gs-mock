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

package assert_test

import (
	"errors"
	"testing"

	"github.com/go-spring/mock"
	"github.com/go-spring/mock/internal/assert"
)

func TestNil(t *testing.T) {
	r, _ := mock.Init(t.Context())
	s := assert.NewTMockImpl(r)

	helperCount := 0
	s.MockHelper().Handle(func() bool {
		helperCount++
		return true
	})

	assert.Nil(s, nil)
	assert.Equal(t, helperCount, 1)
	assert.Nil(s, (*int)(nil))
	assert.Equal(t, helperCount, 2)

	errorCount := 0
	s.MockError().Handle(func(i []interface{}) bool {
		errorCount++
		assert.Equal(t, len(i), 1)
		assert.Equal(t, i[0], "got (*errors.errorString) error but expect nil")
		return true
	})

	assert.Nil(s, errors.New("error"))
	assert.Equal(t, errorCount, 1)
}

func TestEqual(t *testing.T) {
	r, _ := mock.Init(t.Context())
	s := assert.NewTMockImpl(r)

	helperCount := 0
	s.MockHelper().Handle(func() bool {
		helperCount++
		return true
	})

	errorCount := 0
	s.MockError().Handle(func(i []interface{}) bool {
		errorCount++
		assert.Equal(t, len(i), 1)
		assert.Equal(t, i[0], "got (int) 1 but expect (int) 2")
		return true
	})

	assert.Equal(s, 1, 2)
	assert.Equal(t, helperCount, 1)
	assert.Equal(t, errorCount, 1)
}

func TestPanic(t *testing.T) {

	t.Run("did not panic", func(t *testing.T) {
		r, _ := mock.Init(t.Context())
		s := assert.NewTMockImpl(r)

		helperCount := 0
		s.MockHelper().Handle(func() bool {
			helperCount++
			return true
		})

		errorCount := 0
		s.MockError().Handle(func(i []interface{}) bool {
			errorCount++
			assert.Equal(t, len(i), 1)
			assert.Equal(t, i[0], "did not panic")
			return true
		})

		assert.Panic(s, func() {
			// not panic
		}, "")

		assert.Equal(t, helperCount, 1)
		assert.Equal(t, errorCount, 1)
	})

	t.Run("invalid pattern", func(t *testing.T) {
		r, _ := mock.Init(t.Context())
		s := assert.NewTMockImpl(r)

		helperCount := 0
		s.MockHelper().Handle(func() bool {
			helperCount++
			return true
		})

		errorCount := 0
		s.MockError().Handle(func(i []interface{}) bool {
			errorCount++
			assert.Equal(t, len(i), 1)
			assert.Equal(t, i[0], "invalid pattern")
			return true
		})

		assert.Panic(s, func() {
			panic("error")
		}, "\\9")

		assert.Equal(t, helperCount, 2)
		assert.Equal(t, errorCount, 1)
	})

	t.Run("success", func(t *testing.T) {
		r, _ := mock.Init(t.Context())
		s := assert.NewTMockImpl(r)

		helperCount := 0
		s.MockHelper().Handle(func() bool {
			helperCount++
			return true
		})

		errorCount := 0
		s.MockError().Handle(func(i []interface{}) bool {
			errorCount++
			assert.Equal(t, len(i), 1)
			assert.Equal(t, i[0], `got "error" which does not match "panic"`)
			return true
		})

		assert.Panic(s, func() {
			panic("error")
		}, "panic")

		assert.Equal(t, helperCount, 2)
		assert.Equal(t, errorCount, 1)
	})
}
