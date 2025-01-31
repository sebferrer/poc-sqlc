package models

type Book struct {
	ID              int64  `db:"id"`
	Title           string `db:"title"`
	PublicationDate string `db:"publication_date"`
	AuthorID        int64  `db:"author_id"`
}
