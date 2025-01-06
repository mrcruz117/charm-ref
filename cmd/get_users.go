// filepath: /C:/Users/mcruz/Documents/work_repos/study/charm-ref/cmd/adduser.go
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// addUserCmd represents the adduser command
var getUsersCmd = &cobra.Command{
	Use:   "getusers",
	Short: "Get all users",
	Long:  `This command shows all users in the database.`,
	// Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

		ctx := context.Background()
		users, err := Cfg.db.GetAllUsers(ctx)
		if err != nil {
			log.Fatalf("Failed to get users: %v", err)
		}
		for _, user := range users {
			fmt.Printf("User_ID=%s, FirstName=%s, LastName=%s, Email=%s\n", user.ID, user.FirstName, user.LastName, user.Email)
		}
	},
}

func init() {
	rootCmd.AddCommand(getUsersCmd)
}
