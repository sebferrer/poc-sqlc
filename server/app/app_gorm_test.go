package app_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sebferrer/poc-sqlc/gorm/models"
	"github.com/sebferrer/poc-sqlc/server/app"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gotest.tools/v3/assert"
)

func TestRunWithGorm(t *testing.T) {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})

	testAuthor := models.Author{ID: 1, Email: "test@example.com", Bio: "An author"}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "authors" ("email","bio") VALUES ($1,$2) RETURNING "id"`)).
		WithArgs(testAuthor.Email, testAuthor.Bio).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAuthor.ID))
	mock.ExpectCommit()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "authors" WHERE "authors"."id" = $1 ORDER BY "authors"."id" LIMIT $2`)).
		WithArgs(testAuthor.ID, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "bio"}).
			AddRow(testAuthor.ID, testAuthor.Email, testAuthor.Bio))

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "authors" SET "email"=\$1,\s*"bio"=\$2 WHERE "id" = \$3`).
		WithArgs(testAuthor.Email, "Updated Bio", testAuthor.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "authors" WHERE "authors"."id" = $1`)).
		WithArgs(testAuthor.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	app.RunWithGorm(gormDB)

	assert.NilError(t, mock.ExpectationsWereMet())
}
