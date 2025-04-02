package repo

import (
	"context"

	pb "github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/genproto/users_service"
)

type UsersRepo interface {
	CreateUser(ctx context.Context, in *pb.User) (*pb.CreateUserResponse, error)
	GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error)
	GetUserWithLogin(ctx context.Context, in *pb.GetUserWithUsernameRequest) (*pb.User, error)
	GetListUser(ctx context.Context, in *pb.GetListUserRequest) (*pb.UserList, error)
	UpdateUser(ctx context.Context, in *pb.User) error
	DeleteUser(ctx context.Context, id string) error
}
