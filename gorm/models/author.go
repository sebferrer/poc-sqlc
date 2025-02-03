package models

type Author struct {
	ID    int64  `gorm:"primaryKey"`
	Email string `gorm:"unique;not null"`
	Bio   string
	Books []Book `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE"`
}

func (Author) TableName() string {
	return "author"
}
