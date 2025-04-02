package services

import (
	"context"

	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/config"
	pb "github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/genproto/users_service"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/logger"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/server/grpc/client"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/storage"
)

type UsersService struct {
	storage  storage.StorageI
	logger   logger.Logger
	cfg      *config.Config
	services client.ServiceManager
	pb.UnimplementedUserServiceServer
}

func NewUsersService(options *ServiceOptions) pb.UserServiceServer {
	return &UsersService{
		storage:  options.Storage,
		logger:   options.Logger,
		services: options.ServiceManager,
		cfg:      options.Config,
	}
}

func (s *UsersService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	valid, err := validatingUser(in.User)
	if !valid || err != nil {
		return nil, err
	}

	return s.storage.Users().CreateUser(ctx, in.User)
}

func (s *UsersService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	return s.storage.Users().GetUser(ctx, in)
}

func (s *UsersService) GetUserWithLogin(ctx context.Context, in *pb.GetUserWithUsernameRequest) (*pb.User, error) {
	return s.storage.Users().GetUserWithLogin(ctx, in)
}

func (s *UsersService) GetListUser(ctx context.Context, in *pb.GetListUserRequest) (*pb.UserList, error) {
	return s.storage.Users().GetListUser(ctx, in)
}

func (s *UsersService) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	valid, err := validatingUser(in.User)
	if !valid || err != nil {
		return nil, err
	}

	err = s.storage.Users().UpdateUser(ctx, in.User)
	if err != nil {
		return &pb.UpdateUserResponse{Success: false}, err
	}

	return &pb.UpdateUserResponse{Success: true}, nil
}

func (s *UsersService) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := s.storage.Users().DeleteUser(ctx, in.Id)
	if err != nil {
		return &pb.DeleteUserResponse{Success: false}, err
	}

	return &pb.DeleteUserResponse{Success: true}, nil
}
