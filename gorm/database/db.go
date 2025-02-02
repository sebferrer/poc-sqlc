package database

import (
	"github.com/sebferrer/poc-sqlc/gorm/models"
	"github.com/sebferrer/poc-sqlc/server/configuration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb(cfg configuration.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Author{}, &models.Book{})

	return db
}
