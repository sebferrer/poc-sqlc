package database

import (
	"database/sql"

	"github.com/sebferrer/poc-sqlc/gorp/models"
	"github.com/sebferrer/poc-sqlc/server/configuration"

	_ "github.com/lib/pq"
	"gopkg.in/gorp.v2"
)

func InitDb(cfg configuration.Config) *gorp.DbMap {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		panic(err)
	}

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbMap.AddTableWithName(models.Author{}, "author").SetKeys(true, "ID")
	dbMap.AddTableWithName(models.Book{}, "book").SetKeys(true, "ID")

	return dbMap
}
