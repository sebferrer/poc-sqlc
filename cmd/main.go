package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	database1 "github.com/sebferrer/poc-sqlc/gorm/database"
	"github.com/sebferrer/poc-sqlc/server/app"
	"github.com/sebferrer/poc-sqlc/server/configuration"
	database2 "github.com/sebferrer/poc-sqlc/sqlc/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Run with gorm
	cfg := configuration.LoadConfig()
	db1, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database1.InitDb(cfg)
	app.RunWithGorm(db1)

	// Run with SQLC
	connString := cfg.DatabaseURL
	db2, err := sql.Open("pgx", connString)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db2.Close()

	if err := db2.Ping(); err != nil {
		log.Fatalf("Database is unreachable: %v", err)
	}

	ctx := context.Background()
	queries := database2.New(db2)

	app.RunWithSQLC(ctx, queries)
}
