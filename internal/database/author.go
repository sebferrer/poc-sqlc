package database

import (
	"github.com/sebferrer/poc-sqlc/internal/models"

	"gopkg.in/gorp.v2"
)

func InsertAuthor(dbMap *gorp.DbMap, author *models.Author) error {
	return dbMap.Insert(author)
}

func GetAuthor(dbMap *gorp.DbMap, id int64) (*models.Author, error) {
	var author models.Author
	err := dbMap.SelectOne(&author, `SELECT * FROM author WHERE id=$1`, id)
	return &author, err
}

func UpdateAuthor(dbMap *gorp.DbMap, author *models.Author) error {
	_, err := dbMap.Exec(`UPDATE author SET email=$1, bio=$2 WHERE id=$3`, author.Email, author.Bio, author.ID)
	return err
}

func DeleteAuthor(dbMap *gorp.DbMap, id int64) error {
	_, err := dbMap.Exec(`DELETE FROM author WHERE id=$1`, id)
	return err
}