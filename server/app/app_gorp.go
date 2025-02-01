package app

import (
	"fmt"

	"github.com/sebferrer/poc-sqlc/gorp/database"
	"github.com/sebferrer/poc-sqlc/gorp/models"
	"gopkg.in/gorp.v2"
)

func RunWithGorp(dbMap *gorp.DbMap) {
	author := &models.Author{Email: "test@example.com", Bio: "An author"}
	if err := database.CreateAuthor(dbMap, author); err != nil {
		fmt.Println("Error inserting author:", err)
		return
	}
	fmt.Println("Inserted Author:", author)

	retrievedAuthor, err := database.GetAuthor(dbMap, author.ID)
	if err != nil {
		fmt.Println("Error retrieving author:", err)
		return
	}
	fmt.Println("Retrieved Author:", retrievedAuthor)

	retrievedAuthor.Bio = "Updated Bio"
	if err := database.UpdateAuthor(dbMap, retrievedAuthor); err != nil {
		fmt.Println("Error updating author:", err)
		return
	}
	fmt.Println("Updated Author:", retrievedAuthor)

	if err := database.DeleteAuthor(dbMap, retrievedAuthor.ID); err != nil {
		fmt.Println("Error deleting author:", err)
		return
	}
	fmt.Println("Deleted Author")
}
