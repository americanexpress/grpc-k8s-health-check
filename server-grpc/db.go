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
	"log"
	"time"
)

// connectDB mimics a dummy database that waits some time and then changes the isDatabaseReady flag to true.
// This service is used to later check the readiness of the server.
func connectDB() error {
	sleepTime := 30
	log.Println("⏳ Connecting to the dummy database. This might take up to", sleepTime, "seconds")
	time.Sleep(time.Duration(sleepTime) * time.Second)
	log.Println("📣 Database is ready now!")
	isDatabaseReady = true
	return nil
}
