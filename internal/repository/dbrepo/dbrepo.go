package dbrepo

import (
	"database/sql"

	"github.com/jofosuware/mindease/internal/config"
	"github.com/jofosuware/mindease/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

func NewDB(conn *sql.DB) repository.DatabaseRepo {
	return &postgresDBRepo{
		DB: conn,
	}
}
