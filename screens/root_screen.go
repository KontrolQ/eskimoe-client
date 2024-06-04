package screens

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type RootScreen struct{}

func rootScreen() tea.Model {
	return RootScreen{}
}

func (r RootScreen) Init() tea.Cmd {
	return tea.SetWindowTitle("Eskimoe")
}

func (r RootScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Exit on Ctrl+C
	if msg, ok := msg.(tea.KeyMsg); ok {
		if msg.String() == "ctrl+c" {
			return r, tea.Quit
		}
	}

	return r, nil
}

func (r RootScreen) View() string {
  text := lipgloss.NewStyle().Render(fmt.Sprintf("Welcome to Eskimoe, %s! You are on %d servers. Your current server URL is: %s.", globals.currentUser.DisplayName, len(globals.servers), globals.currentUser.CurrentServer.URL))
	return text
}
