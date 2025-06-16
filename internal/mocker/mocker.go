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

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

// init sets the working directory of the application to the directory
// where this source file resides.
func init() {
	var execDir string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		execDir = filepath.Dir(filename)
	}
	err := os.Chdir(execDir)
	if err != nil {
		panic(err)
	}
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(workDir)
}

func main() {
	s := bytes.NewBuffer(nil)

	s.WriteString(`
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
	
	package mock
	
	import (
		"reflect"
	)`)

	const (
		MaxParamCount  = 5
		MaxResultCount = 5
	)

	s.WriteString(fmt.Sprintf(`
	const (
		MaxParamCount  = %d
		MaxResultCount = %d
	)`, MaxParamCount, MaxResultCount))

	for i := 0; i <= MaxParamCount; i++ {
		for j := 0; j <= MaxResultCount; j++ {
			mockerName := fmt.Sprintf("Mocker%d%d", i, j)
			invokerName := fmt.Sprintf("Invoker%d%d", i, j)
			req := make([]string, i)
			for k := 0; k < i; k++ {
				req[k] = "T" + fmt.Sprint(k+1)
			}
			resp := make([]string, j)
			for k := 0; k < j; k++ {
				resp[k] = "R" + fmt.Sprint(k+1)
			}
			respOnlyArg := make([]string, j)
			for k := 0; k < j; k++ {
				respOnlyArg[k] = "r" + fmt.Sprint(k+1)
			}
			cvtParams := make([]string, i)
			for k := 0; k < i; k++ {
				cvtParams[k] = "params[" + fmt.Sprint(k) + "].(T" + fmt.Sprint(k+1) + ")"
			}
			data := map[string]interface{}{
				"mockerName":  mockerName,
				"invokerName": invokerName,
				"req":         strings.Join(req, ", "),
				"resp":        strings.Join(resp, ", "),
				"respOnlyArg": strings.Join(respOnlyArg, ", "),
				"cvtParams":   strings.Join(cvtParams, ", "),
			}
			if err := getTemplate(i, j).Execute(s, data); err != nil {
				panic(err)
			}
		}
	}
	b, err := format.Source(s.Bytes())
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("../../mocker.go", b, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

// getTemplate returns a template based on the number of input (i) and output (j) parameters.
func getTemplate(i, j int) *template.Template {
	if i == 0 {
		if j == 0 {
			return mocker00Tmpl
		} else {
			return mocker0NTmpl
		}
	} else {
		if j == 0 {
			return mockerN0Tmpl
		} else {
			return mockerNNTmpl
		}
	}
}
