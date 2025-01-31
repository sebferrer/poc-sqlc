package database

import (
	"database/sql"
	"fmt"

	"github.com/sebferrer/poc-sqlc/internal/configuration"
	"github.com/sebferrer/poc-sqlc/internal/models"

	_ "github.com/lib/pq"
	"gopkg.in/gorp.v2"
)

func InitDb(cfg configuration.Config) *gorp.DbMap {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbHost, cfg.DbPort)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbMap.AddTableWithName(models.Author{}, "author").SetKeys(true, "ID")
	dbMap.AddTableWithName(models.Book{}, "book").SetKeys(true, "ID")

	return dbMap
}
