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
	"net"

	"client-server-grpc/api"
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// server struct
type server struct {
	Name string
}

// happyUpper takes a string and converts it to upper case and adds a smiley face emoji at the end of the string.
func happyUpper(s string) string {
	return strings.ToUpper(s) + "😊"
}

// This function takes in the context named c, and an api (api is the name of our proto file) request named req and outputs a pointer to the OutputResponse.
func (s server) Upper(c context.Context, req *api.InputRequest) (*api.OutputResponse, error) {

	x := happyUpper(req.GetText())

	log.Printf("➡️ Received message from client %v: %v ", req.GetClientName(), req.GetText())

	// Return a pointer to the api OutputResponse struct where server name equals to serverName and Text equals to x, and error is nil.
	return &api.OutputResponse{ServerName: serverName, Text: x}, nil
}

// Health struct
type Health struct{}

// Check does the health check and changes the status of the server based on wether the db is ready or not.
func (h *Health) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	log.Println("🏥 K8s is health checking")
	if isDatabaseReady == true {
		log.Printf("✅ Server's status is %s", grpc_health_v1.HealthCheckResponse_SERVING)
		return &grpc_health_v1.HealthCheckResponse{
			Status: grpc_health_v1.HealthCheckResponse_SERVING,
		}, nil
	} else if isDatabaseReady == false {
		log.Printf("🚫 Server's status is %s", grpc_health_v1.HealthCheckResponse_NOT_SERVING)
		return &grpc_health_v1.HealthCheckResponse{
			Status: grpc_health_v1.HealthCheckResponse_NOT_SERVING,
		}, nil
	} else {
		log.Printf("🚫 Server's status is %s", grpc_health_v1.HealthCheckResponse_UNKNOWN)
		return &grpc_health_v1.HealthCheckResponse{
			Status: grpc_health_v1.HealthCheckResponse_UNKNOWN,
		}, nil
	}

}

// Watch is used by clients to receive updates when the service status changes.
// Watch only dummy implemented just to satisfy the interface.
func (h *Health) Watch(req *grpc_health_v1.HealthCheckRequest, w grpc_health_v1.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "Watching is not supported")
}

func startGrpcServer() {
	// Opening a tcp listener. Listen on all interfaces.
	ln, err := net.Listen("tcp", port)

	if err != nil {
		// Break if can't listen.
		log.Fatalf("❌ Error listening to interfaces. %v ", err)
	}

	log.Println("👂 Server is listening on port", port)
	log.Printf("💡 Server's name is %s", serverName)

	// Using NewServer() function from "google.golang.org/grpc" library to create a new grpc server object.
	grpcServer := grpc.NewServer()
	srv := &server{Name: serverName}

	// Register our service on the server. Use a function from the api (proto) called RegisterProcessTextServer, pass the server and a reference from server struct type.
	api.RegisterProcessTextServer(grpcServer, srv)
	// Register the health service.
	grpc_health_v1.RegisterHealthServer(grpcServer, &Health{})

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	// Pass our listener to the Server. This sets up the grpc server.
	err = grpcServer.Serve(ln)
	if err != nil {
		// Break if can't set up server.
		log.Fatalf("❌ Error setting up the server. %v ", err)
	}

}
