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

FROM golang:1.12.5-alpine3.9 AS BUILD_STAGE
# Add vendor to GOPATH
ADD vendor /go/src/ 
# Add server-grpc to the docker image
ADD server-grpc /go/src/client-server-grpc/server-grpc
# Add api to the docker image
ADD api /go/src/client-server-grpc/api
# Set working dir in the container 
WORKDIR /go/src/client-server-grpc/server-grpc
# Build the program and output it to app
RUN go build -o /app
# Make the dockerfile more optimized by using multistage dockerbuild which we copy the binary from the BUILD_STAGE container to the final container
# The second FROM instruction starts a new build stage with the alpine image as its base
FROM alpine:3.9
# Get dependencies
RUN apk add wget
RUN GRPC_HEALTH_PROBE_VERSION=v0.3.0 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 
# Make it executable 
RUN chmod +x /bin/grpc_health_probe  
COPY --from=BUILD_STAGE /app /app
# Expose port to the outside once the container has launched
EXPOSE 3000
# The ENTRYPOINT of an image specifies what executable to run when the container starts
ENTRYPOINT [ "/app" ]