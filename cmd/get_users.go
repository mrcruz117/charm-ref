// filepath: /C:/Users/mcruz/Documents/work_repos/study/charm-ref/cmd/adduser.go
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

type model struct {
	table table.Model
}

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

		columns := []table.Column{
			{Title: "ID", Width: 15},
			{Title: "First Name", Width: 12},
			{Title: "Last Name", Width: 12},
			{Title: "Email", Width: 20},
			{Title: "Created At", Width: 19},
			{Title: "Updated At", Width: 19},
		}
		rows := make([]table.Row, 0)
		location := time.Now().Location()

		for _, user := range users {
			createdAt := user.CreatedAt.In(location).Format("1/02/2006 3:04PM")
			updatedAt := user.UpdatedAt.In(location).Format("1/02/2006 3:04PM")
			rows = append(rows, table.Row{user.ID, user.FirstName, user.LastName, user.Email, createdAt, updatedAt})
		}

		t := table.New(
			table.WithColumns(columns),
			table.WithRows(rows),
			table.WithFocused(true),
			table.WithHeight(7),
		)
		s := table.DefaultStyles()
		s.Header = s.Header.
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")).
			BorderBottom(true).
			Bold(false)
		s.Selected = s.Selected.
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("57")).
			Bold(false)
		t.SetStyles(s)
		m := model{t}
		if _, err := tea.NewProgram(m).Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
		// for _, user := range users {
		// 	fmt.Printf("User_ID=%s, FirstName=%s, LastName=%s, Email=%s\n", user.ID, user.FirstName, user.LastName, user.Email)
		// }
	},
}

func init() {
	rootCmd.AddCommand(getUsersCmd)
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("%s ID: %s", m.table.SelectedRow()[3], m.table.SelectedRow()[0]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n  " + m.table.HelpView() + "\n"
}
