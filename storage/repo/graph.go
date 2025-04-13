package repo

import (
	"context"

	pb "github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/genproto/users_service"
)

type FollowUnFollowRepo interface {
	CreateUser(ctx context.Context, in *pb.User) (string, error)
	GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error)
	UpdateUser(ctx context.Context, in *pb.User) error
	DeleteUser(ctx context.Context, id string) error

	Follow(ctx context.Context, in *pb.FollowRequest) error
	UnFollow(ctx context.Context, in *pb.UnFollowRequest) error
	GetListFollowings(ctx context.Context, in *pb.GetListFollowingRequest) (*pb.FollowingList, error)
	GetListFollowers(ctx context.Context, in *pb.GetListFollowerRequest) (*pb.FollowerList, error)
}
