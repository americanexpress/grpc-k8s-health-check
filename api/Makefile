# Copyright 2019 American Express Travel Related Services Company, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
# in compliance with the License. You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software distributed under the License
# is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
# or implied. See the License for the specific language governing permissions and limitations under
# the License.

# Downlaod dependencies. 
dep:
	go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

# Genrate the api.pb.go
generate-proto:
	protoc -I . api.proto --go_out=plugins=grpc:.

## Use this command to generate the api.pb.go inside a docker image and then return the file to the mac and remove the docker. 
## This command is useful for when VPN does not allow you to run protoc -I
generate-proto-in-docker:
	docker run --rm -v $(CURDIR):$(CURDIR) -w $(CURDIR) znly/protoc -I . api.proto --go_out=plugins=grpc:.

