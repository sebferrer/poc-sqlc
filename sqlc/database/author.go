package database

import (
	"context"
	"database/sql"
)

func (q *Queries) CreateAuthor(ctx context.Context, email string, bio string) (int32, error) {
	params := createAuthorParams{
		Email: email,
		Bio:   sql.NullString{String: bio, Valid: bio != ""},
	}
	return q.createAuthor(ctx, params)
}

func (q *Queries) GetAuthor(ctx context.Context, id int32) (*Author, error) {
	author, err := q.getAuthor(ctx, id)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (q *Queries) UpdateAuthor(ctx context.Context, id int32, email string, bio string) error {
	params := updateAuthorParams{
		ID:    id,
		Email: email,
		Bio:   sql.NullString{String: bio, Valid: bio != ""},
	}
	return q.updateAuthor(ctx, params)
}

func (q *Queries) DeleteAuthor(ctx context.Context, id int32) error {
	return q.deleteAuthor(ctx, id)
}
