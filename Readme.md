# gRPC User Service
## Overview

This project is a Golang gRPC service for managing user details and includes a search capability.

## Project Structure
```
grpc-user-service/
├── cmd/
│   └── server/
│       └── main.go
├── pkg/
│   └── user/
│       ├── user.go
│       └── user_test.go
├── proto/
│   └── user.proto
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```
## Prerequisites

1. **Install Protocol Buffers Compiler (`protoc`):**

   - **macOS:** Using Homebrew
     ```sh
     brew install protobuf
     ```

   - **Ubuntu/Debian:** Using apt
     ```sh
     sudo apt update
     sudo apt install -y protobuf-compiler
     ```

   - **Windows:** Download from the [official releases page](https://github.com/protocolbuffers/protobuf/releases) and add the `bin` directory to your PATH.

2. **Install Go Plugins for Protobuf and gRPC:**

   ```sh
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

3. **Install dependencies:**

    ```sh
    go mod tidy

4. **Postman Rest to gRPC brdige:**
    ```sh
    go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
    go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

5. **Install dependencies:**
    ```sh
    git clone https://github.com/googleapis/googleapis.git

## Generate proto files
    ```sh
    protoc --go_out=. --go-grpc_out=. proto/user.proto
    protoc -I . -I ./googleapis --grpc-gateway_out ./proto --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative proto/user.proto
    ```
## Build the application
    go build -o grpc_user_service ./cmd/server

## Run the application:
    ./grpc_user_service

## Using Docker

   - **Build the Docker image::** 
     ```sh
     docker build -t grpc_user_service .
     ```

   - **Run the Docker container:** 
     ```sh
     docker run -p 50051:50051 grpc_user_service
     ```
