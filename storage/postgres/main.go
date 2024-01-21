package postgres

import (
	"time"

	"gitlab.com/samandar_tukhtayev/trello/config"
	"gitlab.com/samandar_tukhtayev/trello/pkg/db"
	"gitlab.com/samandar_tukhtayev/trello/pkg/logger"
)

var (
	CreatedAt time.Time
	UpdatedAt time.Time
)

type postgresRepo struct {
	Db  *db.Postgres
	Log *logger.Logger
	Cfg config.Config
}

func New(db *db.Postgres, log *logger.Logger, cfg config.Config) PostgresI {
	return &postgresRepo{
		Db:  db,
		Log: log,
		Cfg: cfg,
	}
}
