package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"grpc_user_service/pkg/user"
	pb "grpc_user_service/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Update the endpoint to match your gRPC server's address
	grpcServerEndpoint := "localhost:50051"
	err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	go func() {
		if err := StartGRPCServer(); err != nil {
			log.Fatalf("failed to serve gRPC server: %v", err)
		}
	}()
	log.Printf("Starting gRPC-Gateway server on :8080, connecting to gRPC server at %s", grpcServerEndpoint)

	return http.ListenAndServe(":8080", mux)
}

func StartGRPCServer() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, user.NewUserServiceServer())

	log.Println("gRPC server started on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return nil
}
func main() {
	if err := run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
