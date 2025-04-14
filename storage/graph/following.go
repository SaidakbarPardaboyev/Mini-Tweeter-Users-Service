package graph

import (
	"context"
	"fmt"
	"strings"

	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/config"
	pb "github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/genproto/users_service"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/logger"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/storage/repo"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type followUnFollow struct {
	db  neo4j.Session // Correct type (not *neo4j.Session)
	cff *config.Config
	log logger.Logger
}

func NewFollowUnFollowRepo(db neo4j.Session, log logger.Logger, cff *config.Config) repo.FollowUnFollowRepo {
	return &followUnFollow{db: db, log: log, cff: cff}
}

func (f *followUnFollow) CreateUser(ctx context.Context, in *pb.User) (string, error) {
	query := `
		CREATE (n:User {
			user_id: $user_id,
			username: $username,
			name: $name,
			surname: $surname,
			bio: $bio,
			profile_image: $profile_image
		})
	`

	params := map[string]interface{}{
		"user_id":       in.Id,
		"username":      in.Username,
		"name":          in.Name,
		"surname":       in.Surname,
		"bio":           in.Bio,
		"profile_image": in.ProfileImage,
	}

	_, err := f.db.Run(query, params)
	if err != nil {
		return "", fmt.Errorf("user creation failed error: " + err.Error())
	}

	return "", nil
}

func (f *followUnFollow) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	query := `
		MATCH (n:User {user_id: $id})
		RETURN 
			n.user_id, 
			n.username, 
			n.name, 
			n.surname, 
			n.bio, 
			n.profile_image
	`

	params := map[string]interface{}{
		"id": in.Id,
	}

	res, err := f.db.Run(query, params)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}
	if res.Next() {
		record := res.Record()
		user := &pb.User{
			Id:           record.GetByIndex(0).(string),
			Username:     record.GetByIndex(1).(string),
			Name:         record.GetByIndex(2).(string),
			Surname:      record.GetByIndex(3).(string),
			Bio:          record.GetByIndex(4).(string),
			ProfileImage: record.GetByIndex(5).(string),
		}
		return user, nil
	}

	// If no result found, return nil
	if res.Err() != nil {
		return nil, res.Err()
	}

	return nil, nil
}

func (f *followUnFollow) UpdateUser(ctx context.Context, in *pb.User) error {
	query := `
		MATCH (n:User {user_id: $id})
		SET n.username = $username,
		    n.name = $name,
		    n.surname = $surname,
		    n.bio = $bio,
		    n.profile_image = $profile_image
	`
	params := map[string]interface{}{
		"id":            in.Id,
		"username":      in.Username,
		"name":          in.Name,
		"surname":       in.Surname,
		"bio":           in.Bio,
		"profile_image": in.ProfileImage,
	}

	_, err := f.db.Run(query, params)
	return err
}

func (f *followUnFollow) DeleteUser(ctx context.Context, id string) error {
	query := `
		MATCH (n:User {user_id: $id})
		DETACH DELETE n
	`
	params := map[string]interface{}{
		"id": id,
	}

	_, err := f.db.Run(query, params)
	return err
}

func (f *followUnFollow) Follow(ctx context.Context, in *pb.FollowRequest) error {
	query := `
		MATCH (a:User {user_id: $followerId}), (b:User {user_id: $followingId})
		MERGE (a)-[:FOLLOWS]->(b)
	`
	params := map[string]interface{}{
		"followerId":  in.FollowerId,
		"followingId": in.FollowingId,
	}

	_, err := f.db.Run(query, params)
	return err
}

func (f *followUnFollow) UnFollow(ctx context.Context, in *pb.UnFollowRequest) error {
	query := `
		MATCH (a:User {user_id: $followerId})-[r:FOLLOWS]->(b:User {user_id: $followingId})
		DELETE r
	`
	params := map[string]interface{}{
		"followerId":  in.FollowerId,
		"followingId": in.FollowingId,
	}

	_, err := f.db.Run(query, params)
	return err
}

func (f *followUnFollow) GetListFollowings(ctx context.Context, in *pb.GetListFollowingRequest) (*pb.FollowingList, error) {
	query := `
		MATCH (n:User {user_id: $id})-[:FOLLOWS]->(f:User)
	`
	params := map[string]interface{}{
		"id": in.Id,
	}

	if in.Search != nil && strings.TrimSpace(*in.Search) != "" {
		query += `
		WHERE f.username CONTAINS $search
		   OR f.name CONTAINS $search
		   OR f.surname CONTAINS $search
		   OR f.bio CONTAINS $search
		`
		params["search"] = *in.Search
	}

	query += `
		RETURN f.user_id, f.username, f.name
		SKIP $skip
		LIMIT $limit
	`

	params["skip"] = (in.Page - 1) * in.Limit // Page 1 -> skip 0, Page 2 -> skip limit, etc
	params["limit"] = in.Limit

	res, err := f.db.Run(query, params)
	if err != nil {
		return nil, err
	}

	var followings []*pb.User
	for res.Next() {
		record := res.Record()
		followings = append(followings, &pb.User{
			Id:       record.GetByIndex(0).(string),
			Username: record.GetByIndex(1).(string),
			Name:     record.GetByIndex(2).(string),
		})
	}

	return &pb.FollowingList{
		Followings: followings,
	}, nil
}

func (f *followUnFollow) GetListFollowers(ctx context.Context, in *pb.GetListFollowerRequest) (*pb.FollowerList, error) {
	query := `
		MATCH (f:User)-[:FOLLOWS]->(u:User {user_id: $id})
	`
	params := map[string]interface{}{
		"id": in.Id,
	}

	if in.Search != nil && strings.TrimSpace(*in.Search) != "" {
		query += `
		WHERE f.username CONTAINS $search
		   OR f.name CONTAINS $search
		   OR f.surname CONTAINS $search
		   OR f.bio CONTAINS $search
		`
		params["search"] = *in.Search
	}

	query += `
		RETURN f.user_id, f.username, f.name
		SKIP $skip
		LIMIT $limit
	`

	params["skip"] = (in.Page - 1) * in.Limit
	params["limit"] = in.Limit

	res, err := f.db.Run(query, params)
	if err != nil {
		return nil, err
	}

	var followers []*pb.User
	for res.Next() {
		record := res.Record()
		followers = append(followers, &pb.User{
			Id:       record.GetByIndex(0).(string),
			Username: record.GetByIndex(1).(string),
			Name:     record.GetByIndex(2).(string),
		})
	}

	return &pb.FollowerList{
		Followers: followers,
	}, nil
}
