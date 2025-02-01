package app

import (
	"context"
	"fmt"

	"github.com/sebferrer/poc-sqlc/sqlc/database"
)

func RunWithSQLC(ctx context.Context, queries *database.Queries) {

	authorID, err := queries.CreateAuthor(ctx, "test@example.com", "An author")
	if err != nil {
		fmt.Println("Error inserting author:", err)
		return
	}
	fmt.Println("Inserted Author with ID:", authorID)

	retrievedAuthor, err := queries.GetAuthor(ctx, authorID)
	if err != nil {
		fmt.Println("Error retrieving author:", err)
		return
	}
	fmt.Println("Retrieved Author:", retrievedAuthor)

	err = queries.UpdateAuthor(ctx, retrievedAuthor.ID, retrievedAuthor.Email, "Updated Bio")
	if err != nil {
		fmt.Println("Error updating author:", err)
		return
	}
	fmt.Println("Updated Author:", retrievedAuthor)

	err = queries.DeleteAuthor(ctx, retrievedAuthor.ID)
	if err != nil {
		fmt.Println("Error deleting author:", err)
		return
	}
	fmt.Println("Deleted Author")
}
