package database

import (
	"context"
	"time"
)

func (q *Queries) CreateBook(ctx context.Context, title string, publicationDate time.Time, authorID int32) (int32, error) {
	params := createBookParams{
		Title:           title,
		PublicationDate: publicationDate,
		AuthorID:        authorID,
	}
	return q.createBook(ctx, params)
}

func (q *Queries) GetBook(ctx context.Context, id int32) (*Book, error) {
	book, err := q.getBook(ctx, id)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (q *Queries) UpdateBook(ctx context.Context, id int32, title string, publicationDate time.Time, authorID int32) error {
	params := updateBookParams{
		ID:              id,
		Title:           title,
		PublicationDate: publicationDate,
		AuthorID:        authorID,
	}
	return q.updateBook(ctx, params)
}

func (q *Queries) DeleteBook(ctx context.Context, id int32) error {
	return q.deleteBook(ctx, id)
}
