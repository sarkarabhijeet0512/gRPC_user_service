package user

import (
	"context"

	pb "grpc_user_service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server represents the gRPC server
type Server struct {
	pb.UnimplementedUserServiceServer
	Users map[int32]*pb.User
}

func NewUserServiceServer() *Server {
	// Simulate a database with a list of users
	users := map[int32]*pb.User{
		1: {Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
		2: {Id: 2, Fname: "Bob", City: "CA", Phone: 1234567890, Height: 6.8, Married: false},
		// Add more users as needed
	}
	return &Server{Users: users}
}

// GetUser handles fetching a user by ID
func (s *Server) GetUser(ctx context.Context, req *pb.UserIdRequest) (*pb.UserResponse, error) {
	if req.Id <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user ID")
	}

	user, exists := s.Users[req.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	return &pb.UserResponse{User: user}, nil
}

// GetUsers handles fetching users by a list of IDs
func (s *Server) GetUsers(ctx context.Context, req *pb.UserIdsRequest) (*pb.UsersResponse, error) {
	if len(req.Ids) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "no user IDs provided")
	}

	var users []*pb.User
	for _, id := range req.Ids {
		if id <= 0 {
			return nil, status.Errorf(codes.InvalidArgument, "invalid user ID: %d", id)
		}

		if user, exists := s.Users[id]; exists {
			users = append(users, user)
		}
	}

	return &pb.UsersResponse{Users: users}, nil
}

// SearchUsers handles searching users by criteria
func (s *Server) SearchUsers(ctx context.Context, req *pb.SearchRequest) (*pb.UsersResponse, error) {
	if req.City == "" && req.Phone == 0 && !req.Married {
		return nil, status.Errorf(codes.InvalidArgument, "at least one search criterion must be provided")
	}

	var users []*pb.User
	for _, user := range s.Users {
		if req.City != "" && user.City != req.City {
			continue
		}
		if req.Phone != 0 && user.Phone != req.Phone {
			continue
		}
		if req.Married && !user.Married {
			continue
		}
		users = append(users, user)
	}

	return &pb.UsersResponse{Users: users}, nil
}
