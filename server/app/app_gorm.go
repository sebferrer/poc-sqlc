package app

import (
	"fmt"

	"github.com/sebferrer/poc-sqlc/gorm/database"
	"github.com/sebferrer/poc-sqlc/gorm/models"
	"gorm.io/gorm"
)

func RunWithGorm(db *gorm.DB) {
	author := &models.Author{Email: "test@example.com", Bio: "An author"}
	if err := database.CreateAuthor(db, author); err != nil {
		fmt.Println("Error inserting author:", err)
		return
	}
	fmt.Println("Inserted Author:", author)

	retrievedAuthor, err := database.GetAuthor(db, author.ID)
	if err != nil {
		fmt.Println("Error retrieving author:", err)
		return
	}
	fmt.Println("Retrieved Author:", retrievedAuthor)

	retrievedAuthor.Bio = "Updated Bio"
	if err := database.UpdateAuthor(db, retrievedAuthor); err != nil {
		fmt.Println("Error updating author:", err)
		return
	}
	fmt.Println("Updated Author:", retrievedAuthor)

	if err := database.DeleteAuthor(db, retrievedAuthor.ID); err != nil {
		fmt.Println("Error deleting author:", err)
		return
	}
	fmt.Println("Deleted Author")
}
