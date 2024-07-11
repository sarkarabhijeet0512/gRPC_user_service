# Use the official Golang image from Docker Hub
FROM golang:1.18-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Clean and download dependencies
RUN go clean -modcache
RUN go mod download

# Copy the entire project directory into the container
COPY . .

# Build the Go application
RUN go build -o /grpc_user_service ./cmd/server

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /grpc_user_service .

# Expose ports for gRPC and gRPC-Gateway
EXPOSE 8080
EXPOSE 50051

# Command to run the executable
CMD ["./grpc_user_service"]
