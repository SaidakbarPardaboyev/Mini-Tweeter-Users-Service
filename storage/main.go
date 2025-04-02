package storage

import (
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/config"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/db"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/logger"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/storage/postgres"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/storage/repo"
)

type StorageI interface {
	Users() repo.UsersRepo
}

type storagePg struct {
	users repo.UsersRepo
}

func New(db *db.Postgres, log logger.Logger, cfg *config.Config) StorageI {
	return &storagePg{
		users: postgres.NewUsersRepo(db, log, cfg),
	}
}

func (s *storagePg) Users() repo.UsersRepo {
	return s.users
}
