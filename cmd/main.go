package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	db1 "github.com/sebferrer/poc-sqlc/gorp/database"
	"github.com/sebferrer/poc-sqlc/server/app"
	"github.com/sebferrer/poc-sqlc/server/configuration"
	db2 "github.com/sebferrer/poc-sqlc/sqlc/database"
)

func main() {
	cfg := configuration.LoadConfig()

	// Run with gorp
	dbMap := db1.InitDb(cfg)
	defer dbMap.Db.Close()
	app.RunWithGorp(dbMap)

	// Run with SQLC
	connString := cfg.DatabaseURL
	db, err := sql.Open("pgx", connString)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Database is unreachable: %v", err)
	}

	ctx := context.Background()
	queries := db2.New(db)

	app.RunWithSQLC(ctx, queries)
}
