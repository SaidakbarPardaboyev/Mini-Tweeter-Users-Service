package postgres

import (
	"context"
	"fmt"
	"testing"

	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/config"
	pb "github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/genproto/users_service"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/db"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/logger"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/storage/repo"
)

func Init() (repo.UsersRepo, error) {
	cfg := config.Load()

	fmt.Println(cfg.PostgresPassword)

	log := logger.New(cfg.Environment, "users_service_grpc")
	psql, err := db.NewPostgresDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("error while connecting to database: %v", err)
	}

	userRepo := NewUsersRepo(psql, log, cfg)

	return userRepo, nil
}

func TestCreateUser(t *testing.T) {
	userRepo, err := Init()
	if err != nil {
		t.Error(fmt.Errorf("error while connecting to database: %v", err))
	}

	resp, err := userRepo.CreateUser(context.Background(), &pb.User{
		Name:         "Salom",
		Surname:      "salom",
		Bio:          "Something",
		ProfileImage: "",
		Username:     "salom",
		PasswordHash: "11062006",
	})
	if err != nil {
		t.Error(fmt.Errorf("error creating user: %v", err))
	}

	fmt.Println("ID: ", resp.Id)
}

func TestGetUser(t *testing.T) {
	userRepo, err := Init()
	if err != nil {
		t.Error(fmt.Errorf("error while connecting to database: %v", err))
	}

	resp, err := userRepo.GetUser(context.Background(), &pb.GetUserRequest{
		Id: "4ea5c837-6843-43a9-84b8-08809a413a3a",
	})
	if err != nil {
		t.Error(fmt.Errorf("error creating user: %v", err))
	}

	fmt.Println("User info: ", resp)
}

func TestGetUserWithLogin(t *testing.T) {
	userRepo, err := Init()
	if err != nil {
		t.Error(fmt.Errorf("error while connecting to database: %v", err))
	}

	resp, err := userRepo.GetUserWithLogin(context.Background(), &pb.GetUserWithUsernameRequest{
		Username: "string",
	})
	if err != nil {
		t.Error(fmt.Errorf("error creating user: %v", err))
	}

	fmt.Println("User info: ", resp)
}

func TestGetListUser(t *testing.T) {
	userRepo, err := Init()
	if err != nil {
		t.Error(fmt.Errorf("error while connecting to database: %v", err))
	}

	var search = "parda"
	resp, err := userRepo.GetListUser(context.Background(), &pb.GetListUserRequest{
		Search: &search,
		Page:   1,
		Limit:  10,
	})
	if err != nil {
		t.Error(fmt.Errorf("error creating user: %v", err))
	}

	for _, user := range resp.Users {
		fmt.Println("User info: ", user)
	}
}

func TestUpdateUser(t *testing.T) {
	userRepo, err := Init()
	if err != nil {
		t.Error(fmt.Errorf("error while connecting to database: %v", err))
	}

	err = userRepo.UpdateUser(context.Background(), &pb.User{
		Id:           "4ea5c837-6843-43a9-84b8-08809a413a3a",
		Name:         "Salom01",
		Surname:      "salom01",
		Bio:          "Something01",
		ProfileImage: "",
		Username:     "salom01",
		PasswordHash: "1106200601",
	})
	if err != nil {
		t.Error(fmt.Errorf("error creating user: %v", err))
	}
}

func TestDeleteUser(t *testing.T) {
	userRepo, err := Init()
	if err != nil {
		t.Error(fmt.Errorf("error while connecting to database: %v", err))
	}

	err = userRepo.DeleteUser(context.Background(), "4ea5c837-6843-43a9-84b8-08809a413a3a")
	if err != nil {
		t.Error(fmt.Errorf("error creating user: %v", err))
	}
}
