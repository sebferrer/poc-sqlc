package app_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sebferrer/poc-sqlc/internal/database"
	"github.com/sebferrer/poc-sqlc/internal/models"
	"gopkg.in/gorp.v2"
	"gotest.tools/v3/assert"
)

func TestRun(t *testing.T) {
	db, mock, _ := sqlmock.New()
	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbMap.AddTableWithName(models.Author{}, "author").SetKeys(true, "ID")
	dbMap.AddTableWithName(models.Book{}, "book").SetKeys(true, "ID")

	insertSQL := regexp.QuoteMeta(`insert into "author" ("id","email","bio") values (default,$1,$2) returning "id"`)
	mock.ExpectQuery(insertSQL).
		WithArgs("test@example.com", "An author").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	selectSQL := regexp.QuoteMeta(`SELECT * FROM author WHERE id=$1`)
	mock.ExpectQuery(selectSQL).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "bio"}).
			AddRow(1, "test@example.com", "An author"))

	updateSQL := regexp.QuoteMeta(`UPDATE author SET email=$1, bio=$2 WHERE id=$3`)
	mock.ExpectExec(updateSQL).
		WithArgs("test@example.com", "Updated Bio", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	deleteSQL := regexp.QuoteMeta(`DELETE FROM author WHERE id=$1`)
	mock.ExpectExec(deleteSQL).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	author := &models.Author{Email: "test@example.com", Bio: "An author"}
	err := database.InsertAuthor(dbMap, author)
	assert.NilError(t, err)

	retrievedAuthor, err := database.GetAuthor(dbMap, author.ID)
	assert.NilError(t, err)
	assert.Equal(t, retrievedAuthor.Email, "test@example.com")
	assert.Equal(t, retrievedAuthor.Bio, "An author")

	author.Bio = "Updated Bio"
	err = database.UpdateAuthor(dbMap, author)
	assert.NilError(t, err)

	err = database.DeleteAuthor(dbMap, author.ID)
	assert.NilError(t, err)

	assert.NilError(t, mock.ExpectationsWereMet())
}
