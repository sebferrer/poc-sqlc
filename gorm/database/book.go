package database

import (
	"github.com/sebferrer/poc-sqlc/gorm/models"
	"gorm.io/gorm"
)

func CreateBook(db *gorm.DB, book *models.Book) error {
	return db.Create(book).Error
}

func GetBook(db *gorm.DB, id int64) (*models.Book, error) {
	var book models.Book
	if err := db.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func UpdateBook(db *gorm.DB, book *models.Book) error {
	return db.Save(book).Error
}

func DeleteBook(db *gorm.DB, id int64) error {
	return db.Delete(&models.Book{}, id).Error
}
