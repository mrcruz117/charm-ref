package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	count int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		default:
			m.count++
			return m, tea.Tick(time.Second, func(t time.Time) tea.Msg {
				return nil
			})
		}
	default:
		return m, nil
	}
}

func (m model) View() string {
	return fmt.Sprintf("Count: %d\nPress any key to increment.", m.count)
}

func main() {
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error starting app: %v\n", err)
	}
}
