package database

import (
	"gopkg.in/gorp.v2"

	"github.com/sebferrer/poc-sqlc/gorp/models"
)

func CreateBook(dbMap *gorp.DbMap, book *models.Book) error {
	return dbMap.Insert(book)
}

func GetBook(dbMap *gorp.DbMap, id int64) (*models.Book, error) {
	var book models.Book
	err := dbMap.SelectOne(&book, "SELECT * FROM book WHERE id=$1", id)
	return &book, err
}

func UpdateBook(dbMap *gorp.DbMap, book *models.Book) error {
	_, err := dbMap.Update(book)
	return err
}

func DeleteBook(dbMap *gorp.DbMap, id int64) error {
	_, err := dbMap.Exec("DELETE FROM book WHERE id=$1", id)
	return err
}
