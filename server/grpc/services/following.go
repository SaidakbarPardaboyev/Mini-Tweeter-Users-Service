package services

import (
	"context"

	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/config"
	pb "github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/genproto/users_service"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/logger"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/server/grpc/client"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/storage"
)

type FollowingService struct {
	storage  storage.StorageI
	logger   logger.Logger
	cfg      *config.Config
	services client.ServiceManager
	pb.UnimplementedFollowingServiceServer
}

func NewFollowUnFollowService(options *ServiceOptions) pb.FollowingServiceServer {
	return &FollowingService{
		storage:  options.Storage,
		logger:   options.Logger,
		services: options.ServiceManager,
		cfg:      options.Config,
	}
}

func (s *FollowingService) Follow(ctx context.Context, in *pb.FollowRequest) (*pb.FollowResponse, error) {
	err := s.storage.FollowUnFollow().Follow(ctx, in)
	if err != nil {
		return &pb.FollowResponse{Success: false}, err
	}

	return &pb.FollowResponse{Success: true}, nil
}

func (s *FollowingService) UnFollow(ctx context.Context, in *pb.UnFollowRequest) (*pb.UnFollowResponse, error) {
	err := s.storage.FollowUnFollow().UnFollow(ctx, in)
	if err != nil {
		return &pb.UnFollowResponse{Success: false}, err
	}

	return &pb.UnFollowResponse{Success: true}, nil
}

func (s *FollowingService) GetListFollowings(ctx context.Context, in *pb.GetListFollowingRequest) (*pb.FollowingList, error) {
	if in.Page == 0 {
		in.Page = 1
	}

	if in.Limit == 0 {
		in.Limit = 10
	}

	followings, err := s.storage.FollowUnFollow().GetListFollowings(ctx, in)
	if err != nil {
		return nil, err
	}

	return followings, nil
}

func (s *FollowingService) GetListFollowers(ctx context.Context, in *pb.GetListFollowerRequest) (*pb.FollowerList, error) {
	if in.Page == 0 {
		in.Page = 1
	}

	if in.Limit == 0 {
		in.Limit = 10
	}

	followers, err := s.storage.FollowUnFollow().GetListFollowers(ctx, in)
	if err != nil {
		return nil, err
	}

	return followers, nil
}
