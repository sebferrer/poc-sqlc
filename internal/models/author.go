package models

type Author struct {
	ID    int64  `db:"id"`
	Email string `db:"email"`
	Bio   string `db:"bio"`
}
