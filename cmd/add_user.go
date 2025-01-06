// filepath: /C:/Users/mcruz/Documents/work_repos/study/charm-ref/cmd/adduser.go
package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mrcruz117/charm-ref/internal/database"
	"github.com/spf13/cobra"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// addUserCmd represents the adduser command
var addUserCmd = &cobra.Command{
	Use:   "adduser [first_name] [last_name] [email]",
	Short: "Add a new user",
	Long:  `This command adds a new user to the database.`,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		firstName := args[0]
		lastName := args[1]
		email := args[2]

		// Generate a new UUID for the user ID
		id := uuid.New().String()

		// Create a new user

		ctx := context.Background()
		user, err := Cfg.db.CreateUser(ctx, database.CreateUserParams{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
		})
		if err != nil {
			log.Fatalf("Failed to create user: %v", err)
		}
		fmt.Printf("User created: ID=%s, FirstName=%s, LastName=%s, Email=%s\n", user.ID, user.FirstName, user.LastName, user.Email)
	},
}

func init() {
	rootCmd.AddCommand(addUserCmd)
}
