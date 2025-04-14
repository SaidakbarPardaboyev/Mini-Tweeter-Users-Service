package graph

import (
	"context"
	"fmt"
	"testing"

	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/config"
	pb "github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/genproto/users_service"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/db"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/logger"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/storage/repo"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func Init() (repo.FollowUnFollowRepo, neo4j.Driver, neo4j.Session, error) {
	cfg := config.Load()

	log := logger.New(cfg.Environment, "users_service_grpc")
	driverNeo4j, sessionNeo4j, err := db.NewDgraphClient(cfg) // <-- consider renaming db.NewDgraphClient
	if err != nil {
		return nil, nil, nil, err
	}

	userRepo := NewFollowUnFollowRepo(sessionNeo4j, log, cfg)

	return userRepo, driverNeo4j, sessionNeo4j, nil
}

func TestCreateUser(t *testing.T) {
	followUnFollowRepo, driver, session, err := Init()
	if err != nil {
		t.Fatalf("error while connecting to database: %v", err)
	}
	defer driver.Close()
	defer session.Close()

	_, err = followUnFollowRepo.CreateUser(context.Background(), &pb.User{
		Id:           "4ea5c837-6843-43a9-84b8-08809a413a3d",
		Name:         "Mudorjon",
		Surname:      "Muronjonov",
		Bio:          "Bio",
		ProfileImage: "",
		Username:     "MudorjonMuronjonov",
		PasswordHash: "11062006",
	})
	if err != nil {
		t.Fatalf("error creating user: %s", err.Error())
	}
}

func TestGetUser(t *testing.T) {
	followUnFollowRepo, driver, session, err := Init()
	if err != nil {
		t.Fatalf("error while connecting to database: %v", err)
	}
	defer driver.Close()
	defer session.Close()

	resp, err := followUnFollowRepo.GetUser(context.Background(), &pb.GetUserRequest{
		Id: "4ea5c837-6843-43a9-84b8-08809a413a3a",
	})
	if err != nil {
		t.Error(fmt.Errorf("error creating user: %v", err))
	}

	fmt.Println("User info: ", resp)
}

func TestUpdateUser(t *testing.T) {
	followUnFollowRepo, driver, session, err := Init()
	if err != nil {
		t.Fatalf("error while connecting to database: %v", err)
	}
	defer driver.Close()
	defer session.Close()

	err = followUnFollowRepo.UpdateUser(context.Background(), &pb.User{
		Id:           "4ea5c837-6843-43a9-84b8-08809a413a3c",
		Name:         "Mudorjon",
		Surname:      "Muronjonov",
		Bio:          "Bio",
		ProfileImage: "",
		Username:     "MudorjonMuronjonov",
		PasswordHash: "11062006",
	})
	if err != nil {
		t.Error(fmt.Errorf("error creating user: %v", err))
	}
}

func TestDeleteUser(t *testing.T) {
	followUnFollowRepo, driver, session, err := Init()
	if err != nil {
		t.Fatalf("error while connecting to database: %v", err)
	}
	defer driver.Close()
	defer session.Close()

	err = followUnFollowRepo.DeleteUser(context.Background(), "4ea5c837-6843-43a9-84b8-08809a413a3d")
	if err != nil {
		t.Error(fmt.Errorf("error creating user: %v", err))
	}
}

func TestFollow(t *testing.T) {
	followUnFollowRepo, driver, session, err := Init()
	if err != nil {
		t.Fatalf("error while connecting to database: %v", err)
	}
	defer driver.Close()
	defer session.Close()

	err = followUnFollowRepo.Follow(context.Background(), &pb.FollowRequest{
		FollowerId:  "4ea5c837-6843-43a9-84b8-08809a413a3a",
		FollowingId: "4ea5c837-6843-43a9-84b8-08809a413a3b",
	})
	if err != nil {
		t.Fatalf("error creating user: %s", err.Error())
	}
}

func TestUnFollow(t *testing.T) {
	followUnFollowRepo, driver, session, err := Init()
	if err != nil {
		t.Fatalf("error while connecting to database: %v", err)
	}
	defer driver.Close()
	defer session.Close()

	err = followUnFollowRepo.UnFollow(context.Background(), &pb.UnFollowRequest{
		FollowerId:  "4ea5c837-6843-43a9-84b8-08809a413a3a",
		FollowingId: "4ea5c837-6843-43a9-84b8-08809a413a3b",
	})
	if err != nil {
		t.Fatalf("error creating user: %s", err.Error())
	}
}

func TestGetListFollowings(t *testing.T) {
	followUnFollowRepo, driver, session, err := Init()
	if err != nil {
		t.Fatalf("error while connecting to database: %v", err)
	}
	defer driver.Close()
	defer session.Close()

	var search = "Bio"
	resp, err := followUnFollowRepo.GetListFollowings(context.Background(), &pb.GetListFollowingRequest{
		Id:     "4ea5c837-6843-43a9-84b8-08809a413a3c",
		Page:   1,
		Limit:  10,
		Search: &search,
	})
	if err != nil {
		t.Fatalf("error creating user: %s", err.Error())
	}

	for _, user := range resp.Followings {
		fmt.Println("User info: ", user)
	}
}

func TestGetListFollowers(t *testing.T) {
	followUnFollowRepo, driver, session, err := Init()
	if err != nil {
		t.Fatalf("error while connecting to database: %v", err)
	}
	defer driver.Close()
	defer session.Close()

	var search = "Bio"
	resp, err := followUnFollowRepo.GetListFollowers(context.Background(), &pb.GetListFollowerRequest{
		Id:     "4ea5c837-6843-43a9-84b8-08809a413a3c",
		Page:   1,
		Limit:  10,
		Search: &search,
	})
	if err != nil {
		t.Fatalf("error creating user: %s", err.Error())
	}

	for _, user := range resp.Followers {
		fmt.Println("User info: ", user)
	}
}
