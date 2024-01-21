package storage

import (
	"gitlab.com/samandar_tukhtayev/trello/config"
	"gitlab.com/samandar_tukhtayev/trello/pkg/db"
	"gitlab.com/samandar_tukhtayev/trello/pkg/logger"
	"gitlab.com/samandar_tukhtayev/trello/storage/postgres"
)

type StorageI interface {
	Postgres() postgres.PostgresI
}

type StoragePg struct {
	postgres postgres.PostgresI
}

// NewStoragePg
func New(db *db.Postgres, log *logger.Logger, cfg config.Config) StorageI {
	return &StoragePg{
		postgres: postgres.New(db, log, cfg),
	}
}

func (s *StoragePg) Postgres() postgres.PostgresI {
	return s.postgres
}
