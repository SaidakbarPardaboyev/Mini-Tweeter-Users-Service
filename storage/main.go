package storage

import (
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/config"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/db"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/logger"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/storage/graph"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/storage/postgres"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/storage/repo"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type StorageI interface {
	Users() repo.UsersRepo
	FollowUnFollow() repo.FollowUnFollowRepo
}

type storagePg struct {
	users          repo.UsersRepo
	followUnFollow repo.FollowUnFollowRepo
}

func New(db *db.Postgres, sessionNeo4j neo4j.Session, log logger.Logger, cfg *config.Config) StorageI {
	return &storagePg{
		users:          postgres.NewUsersRepo(db, log, cfg),
		followUnFollow: graph.NewFollowUnFollowRepo(sessionNeo4j, log, cfg),
	}
}

func (s *storagePg) Users() repo.UsersRepo {
	return s.users
}

func (s *storagePg) FollowUnFollow() repo.FollowUnFollowRepo {
	return s.followUnFollow
}
