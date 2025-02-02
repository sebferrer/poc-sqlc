package database

import (
	"github.com/sebferrer/poc-sqlc/gorm/models"
	"gorm.io/gorm"
)

func CreateAuthor(db *gorm.DB, author *models.Author) error {
	return db.Create(author).Error
}

func GetAuthor(db *gorm.DB, id int64) (*models.Author, error) {
	var author models.Author
	if err := db.First(&author, id).Error; err != nil {
		return nil, err
	}
	return &author, nil
}

func UpdateAuthor(db *gorm.DB, author *models.Author) error {
	return db.Save(author).Error
}

func DeleteAuthor(db *gorm.DB, id int64) error {
	return db.Delete(&models.Author{}, id).Error
}
