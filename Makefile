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

# Ensures a dependency is safely vendored in the project
deps:
	dep ensure

# Deploy to kubernetes
deploy: 
	kubectl apply -f kubernetes/deploy.yaml

# Delete deployment and images
clean-minikube:
	kubectl delete deploy grpc-deploy || true
	docker rmi -f serverimg || true
	docker rmi -f clientimg || true

# Remove the server binary and container
clean-server:
	rm ./server-grpc/server-grpc || true
	docker rm myserver 

# Remove the client binary and container
clean-client:
	rm ./client-grpc/client-grpc || true
	docker rm myclient 

# Remove binaries for server and client and stop and remove all the running containers	
clean-all: clean-server clean-client
	docker stop $(docker ps -a -q) || true
	docker rm $(docker ps -a -q)

# docker build -f, --file string Name of the Dockerfile (Default is 'PATH/Dockerfile')
# We use the -f flag with docker build to point to a Dockerfile anywhere in the file system
# Build a docker image for server
build-server-img:
	docker build -t serverimg -f server.Dockerfile .	

# Build a docker image for client
build-client-img:
	docker build -t clientimg -f client.Dockerfile .

# Builds both client and server images
build-imgs: build-server-img build-client-img

# Run client container	
# Example usage: make n=123 run-client
run-client:
	docker run --name myclient${n} -it --network="host" clientimg

# Run server container	
run-server:
	docker run --name myserver -p3000:3000 serverimg

# ---------------------------------------------- #
# additional useful commands					 #
# ---------------------------------------------- #
## to delete a docker image use this command:
## docker rmi -f NAMEOFIMAGE
## to see list of images:  
## docker images 