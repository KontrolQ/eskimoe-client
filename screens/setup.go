package screens

import (
	"eskimoe-client/database"
	"eskimoe-client/lib"
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"

	catppuccin "github.com/catppuccin/go"
)

type SetupScreen struct {
	form *huh.Form
	help help.Model
	keys lib.QuitConfirmKeyMap
}

func setupScreen() tea.Model {
	var form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Display Name").Placeholder("Enter your display name").Validate(huh.ValidateLength(1, 64)).CharLimit(64).Key("display_name").WithWidth(64),
		),
	).WithTheme(theme).WithWidth(64).WithShowHelp(false)

	return SetupScreen{
		form: form,
		help: help.New(),
		keys: lib.QuitConfirmKeys,
	}
}

func (s SetupScreen) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, tea.SetWindowTitle("Eskimoe — Account Setup"))
	cmds = append(cmds, s.form.Init())
	return tea.Batch(cmds...)
}

func (s SetupScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, s.keys.Quit):
			return s, tea.Quit
		}
	}

	// Process the form
	form, cmd := s.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		s.form = f
		cmds = append(cmds, cmd)
	}

	if s.form.State == huh.StateCompleted {
		// Save the display name
		displayname := s.form.GetString("display_name")

		globals.currentUser.DisplayName = displayname
		database.Database.Save(&globals.currentUser)

		if len(globals.servers) == 0 {
			// Switch to the join server screen
			joinServerScreen := joinServerScreen()
			return screen().Switch(joinServerScreen)
		} else {
			// Switch to the root screen
			rootScreen := rootScreen()
			return screen().Switch(rootScreen)
		}
	}

	return s, tea.Batch(cmds...)
}

func (s SetupScreen) View() string {
	box := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.AdaptiveColor{Light: catppuccin.Latte.Rosewater().Hex, Dark: catppuccin.Mocha.Rosewater().Hex}).
		Foreground(lipgloss.AdaptiveColor{Light: catppuccin.Latte.Text().Hex, Dark: catppuccin.Mocha.Text().Hex}).
		Width(120).
		Padding(2)

	welcomeText := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.AdaptiveColor{Light: catppuccin.Latte.Mauve().Hex, Dark: catppuccin.Mocha.Mauve().Hex}).Render("Welcome to Eskimoe Client!")

	uniqueIDText := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.AdaptiveColor{Light: catppuccin.Latte.Text().Hex, Dark: catppuccin.Mocha.Text().Hex}).Render("Your Unique ID:")

	uniqueIDValue := lipgloss.NewStyle().Bold(false).Foreground(lipgloss.AdaptiveColor{Light: catppuccin.Latte.Overlay0().Hex, Dark: catppuccin.Mocha.Overlay0().Hex}).Render(globals.currentUser.UniqueID)

	uniqueTokenText := lipgloss.NewStyle().Bold(true).Render("Your Unique Token:")

	uniqueTokenValue := lipgloss.NewStyle().Bold(false).Foreground(lipgloss.AdaptiveColor{Light: catppuccin.Latte.Overlay0().Hex, Dark: catppuccin.Mocha.Overlay0().Hex}).Render(globals.currentUser.UniqueToken)

	formView := lipgloss.NewStyle().Margin(1, 0).Render(s.form.View())

	helpView := lipgloss.NewStyle().Margin(1, 0).Render(s.help.View(s.keys))

	boxContent := box.Render(fmt.Sprintf("%s\n\n%s %s\n\n%s %s\n\n%s\n\n%s", welcomeText, uniqueIDText, uniqueIDValue, uniqueTokenText, uniqueTokenValue, formView, helpView))

	return lipgloss.Place(
		globals.width,
		globals.height,
		lipgloss.Center,
		lipgloss.Center,
		boxContent,
	)
}
