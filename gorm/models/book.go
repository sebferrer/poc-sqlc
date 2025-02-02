package models

type Book struct {
	ID              int64  `gorm:"primaryKey"`
	Title           string `gorm:"not null"`
	PublicationDate string `gorm:"not null"`
	AuthorID        int64  `gorm:"index"`
}
