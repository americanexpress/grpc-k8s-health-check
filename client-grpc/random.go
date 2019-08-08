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

import (
	"math/rand"
	"time"
)

// Seed for the randomString function.
var seededRand = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

// randomString creates random strings of a certain length from a to z.
func randomString(length int) string {
	x := make([]byte, length)
	for i := 0; i < length; i++ {
		x[i] = byte(97 + seededRand.Intn(25))
	}
	return string(x)
}
