/*******************************************************************************
* Copyright 2019 Samsung Electronics All Rights Reserved.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*
*******************************************************************************/

package injectionchecker

import (
	"errors"
	"strings"
	"unicode/utf8"
)

var injectionOperators = []string{
	";",
	"&",
	"&&",
	"||",
	"'",
	"(",
	")",
	"#",
}

// HasInjectionOperator checks the presence of injection operators in the string
func HasInjectionOperator(str string) bool {
	for _, injection := range injectionOperators {
		if strings.Contains(str, injection) {
			return true
		}
	}
	return false
}

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}
