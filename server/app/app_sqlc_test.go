package app_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sebferrer/poc-sqlc/server/app"
	"github.com/sebferrer/poc-sqlc/sqlc/database"
	"gotest.tools/v3/assert"
)

func TestRunWithSQLC(t *testing.T) {
	db, mock, _ := sqlmock.New()
	queries := database.New(db)
	ctx := context.Background()

	mock.ExpectQuery(`-- name: createAuthor :one
					  INSERT INTO author \(email, bio\)
					  VALUES \(\$1, \$2\)
					  RETURNING id`).
		WithArgs("test@example.com", "An author").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectQuery(`-- name: getAuthor :one
					  SELECT id, email, bio
					  FROM author
					  WHERE id = \$1`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "bio"}).
			AddRow(1, "test@example.com", "An author"))

	mock.ExpectExec(`-- name: updateAuthor :exec
					 UPDATE author
					 SET email = \$1, bio = \$2
					 WHERE id = \$3`).
		WithArgs("test@example.com", "Updated Bio", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(`-- name: deleteAuthor :exec
					 DELETE FROM author
					 WHERE id = \$1`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	app.RunWithSQLC(ctx, queries)

	assert.NilError(t, mock.ExpectationsWereMet())
}
