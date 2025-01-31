package app

import (
	"fmt"

	"github.com/sebferrer/poc-sqlc/internal/database"
	"github.com/sebferrer/poc-sqlc/internal/models"

	"gopkg.in/gorp.v2"
)

func Run(dbMap *gorp.DbMap) {
	author := &models.Author{Email: "test@example.com", Bio: "An author"}
	database.InsertAuthor(dbMap, author)
	fmt.Println("Inserted Author:", author)

	retrievedAuthor, _ := database.GetAuthor(dbMap, author.ID)
	fmt.Println("Retrieved Author:", retrievedAuthor)

	retrievedAuthor.Bio = "Updated Bio"
	database.UpdateAuthor(dbMap, retrievedAuthor)
	fmt.Println("Updated Author:", retrievedAuthor)

	database.DeleteAuthor(dbMap, retrievedAuthor.ID)
	fmt.Println("Deleted Author")
}
