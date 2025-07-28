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

//func TestNil(t *testing.T) {
//	s := assert.NewTMockImpl(gsmock.NewManager())
//
//	s.MockHelper().Ignore()
//
//	assert.Nil(s, nil)
//	assert.Nil(s, (*int)(nil))
//
//	var buf bytes.Buffer
//	s.MockError().Handle(func(args []interface{}) {
//		buf.WriteString(args[0].(string))
//	})
//
//	assert.Nil(s, errors.New("error"))
//	assert.Equal(t, buf.String(), "got (*errors.errorString) error but expect nil")
//}
//
//func TestEqual(t *testing.T) {
//	s := assert.NewTMockImpl(gsmock.NewManager())
//
//	s.MockHelper().Ignore()
//
//	var buf bytes.Buffer
//	s.MockError().Handle(func(args []interface{}) {
//		buf.WriteString(args[0].(string))
//	})
//
//	assert.Equal(s, 1, 2)
//	assert.Equal(t, buf.String(), "got (int) 1 but expect (int) 2")
//}
//
//func TestPanic(t *testing.T) {
//	s := assert.NewTMockImpl(gsmock.NewManager())
//
//	s.MockHelper().Ignore()
//
//	var buf bytes.Buffer
//	s.MockError().Handle(func(args []interface{}) {
//		buf.WriteString(args[0].(string))
//	})
//
//	t.Run("did not panic", func(t *testing.T) {
//		buf.Reset()
//		assert.Panic(s, func() {}, "")
//		assert.Equal(t, buf.String(), "did not panic")
//	})
//
//	t.Run("invalid pattern", func(t *testing.T) {
//		buf.Reset()
//		assert.Panic(s, func() { panic("error") }, "\\9")
//		assert.Equal(t, buf.String(), "invalid pattern")
//	})
//
//	t.Run("success", func(t *testing.T) {
//		buf.Reset()
//		assert.Panic(s, func() { panic("error") }, "panic")
//		assert.Equal(t, buf.String(), `got "error" which does not match "panic"`)
//	})
//}
