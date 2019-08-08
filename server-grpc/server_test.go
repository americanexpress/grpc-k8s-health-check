/*
 * Copyright 2019 American Express Travel Related Services Company, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 */

package main

import "testing"

// TestHappyUpper tests the function happyUpper.
func TestHappyUpper(t *testing.T) {
	var testCases = []struct {
		input string
		want  string
	}{
		{"low", "LOW😊"},
		{"l", "L😊"},
		{"1234", "1234😊"},
		{"", "😊"},
	}

	for _, tc := range testCases {
		have := happyUpper(tc.input)
		if have != tc.want {
			t.Errorf("Error. Want %s have %s", tc.want, have)
		}
	}
}
