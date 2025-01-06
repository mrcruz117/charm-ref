// filepath: /C:/Users/mcruz/Documents/work_repos/study/charm-ref/cmd/adduser.go
package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/google/uuid"
	"github.com/mrcruz117/charm-ref/internal/database"

	// "github.com/mrcruz117/charm-ref/spinner"
	"github.com/spf13/cobra"
)

var base *huh.Theme = huh.ThemeBase()
var theme *huh.Theme = base

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
	Use:   "adduser",
	Short: "Add a new user",
	Long:  `This command adds a new user to the database.`,
	// Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		// type UserArgs struct {
		// 	FirstName string
		// 	LastName  string
		// 	Email     string
		// }

		userForm := User{}
		var confirm bool

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Value(&userForm.FirstName).
					Title("What's your first name?").
					Placeholder("Homer").
					Validate(func(s string) error {
						if s == "" {
							return errors.New("please enter your first name")
						}
						return nil
					}),
				// Description("Just first name, please."),
				huh.NewInput().
					Value(&userForm.LastName).
					Title("What's your last name?").
					Placeholder("Simpson").
					Validate(func(s string) error {
						if s == "" {
							return errors.New("please enter your last name")
						}
						return nil
					}),
				// Description("Just last name, please."),
				huh.NewInput().
					Value(&userForm.Email).
					Title("What's your email address?").
					Placeholder("fake@email.com").
					Validate(func(s string) error {
						if s == "" {
							return errors.New("please enter your email address")
						}
						if err := ValidateEmail(s); err != nil {
							return err
						}
						return nil
					}).
					Description("We'll never share your email with anyone."),
				huh.NewConfirm().Title("Confirm info is correct?").Value(&confirm).Description("White box is the choice you're selecting"),
			),
		)
		err := form.WithTheme(theme).Run()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		if !confirm {
			fmt.Println("Aborted")
			os.Exit(0)
		}

		ctx := context.Background()
		s := spinner.New()
		s.Title("Creating user...")

		var user database.User
		startTime := time.Now()
		action := func() {
			// Generate a new UUID for the user ID
			id := uuid.New().String()

			// Create the user
			user, err = Cfg.db.CreateUser(ctx, database.CreateUserParams{
				ID:        id,
				FirstName: userForm.FirstName,
				LastName:  userForm.LastName,
				Email:     userForm.Email,
			})
			if err != nil {
				log.Fatalf("Failed to create user: %v", err)
			}
		}

		// Run the spinner with the action
		err = s.Action(action).Run()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		endTime := time.Now()
		fmt.Printf("\nUser created: ID=%s, FirstName=%s, LastName=%s, Email=%s\n", user.ID, user.FirstName, user.LastName, user.Email)
		fmt.Printf("\nTime taken: %v\n", endTime.Sub(startTime))
	},
}

func init() {
	rootCmd.AddCommand(addUserCmd)
}

func ValidateEmail(email string) error {
	// Define a regular expression for validating an email address
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return errors.New("invalid email address")
	}
	return nil
}
