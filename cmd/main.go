package main

import (
	"github.com/sebferrer/poc-sqlc/internal/app"
	"github.com/sebferrer/poc-sqlc/internal/configuration"
	"github.com/sebferrer/poc-sqlc/internal/database"
)

func main() {
	cfg := configuration.LoadConfig()
	dbMap := database.InitDb(cfg)
	defer dbMap.Db.Close()

	app.Run(dbMap)
}
