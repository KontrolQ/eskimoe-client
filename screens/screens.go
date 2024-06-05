package screens

import (
	"eskimoe-client/database"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type ScreenSwitcher struct {
	currentScreen tea.Model
}

type Globals struct {
	currentUser database.User
	servers     []database.Server
	width       int
	height      int
}

var globals Globals
var theme *huh.Theme = huh.ThemeCatppuccin()

func (s ScreenSwitcher) Init() tea.Cmd {
	return s.currentScreen.Init()
}

func (s ScreenSwitcher) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var model tea.Model

	switch m := msg.(type) {
	case tea.WindowSizeMsg:
		globals.width, globals.height = m.Width, m.Height
	}

	model, cmd = s.currentScreen.Update(msg)

	return ScreenSwitcher{currentScreen: model}, cmd
}

func (s ScreenSwitcher) View() string {
	return s.currentScreen.View()
}

func (s ScreenSwitcher) Switch(screen tea.Model) (tea.Model, tea.Cmd) {
	s.currentScreen = screen
	return s.currentScreen, s.currentScreen.Init()
}

func screen() ScreenSwitcher {
	var screen tea.Model

	if globals.currentUser.DisplayName == "" {
		screen = setupScreen()
	} else if len(globals.servers) == 0 {
		screen = joinServerScreen()
	} else {
		if globals.currentUser.CurrentServer.ID == 0 {
			globals.currentUser.CurrentServer = globals.servers[0]
		}
		screen = rootScreen()
	}

	return ScreenSwitcher{
		currentScreen: screen,
	}
}

// Initialize just calls Screen()
func Initialize() tea.Model {
	currentUser := database.GetCurrentUser()
	servers := database.GetServers(currentUser)

	globals.currentUser = currentUser
	globals.servers = servers

	return screen()
}
