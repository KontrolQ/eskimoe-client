package screens

import (
	"eskimoe-client/api"
	"eskimoe-client/database"
	"eskimoe-client/lib"
	"fmt"
	"log"

	catppuccin "github.com/catppuccin/go"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type JoinServerScreen struct {
	form          *huh.Form
	help          help.Model
	keys          lib.QuitConfirmKeyMap
	serverEntered bool
	server        database.Server
	serverInfo    api.ServerInfo
}

func joinServerScreen() tea.Model {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Server URL").Placeholder("Enter the server URL").Validate(huh.ValidateLength(1, 256)).Validate(lib.ServerReachableValidator()).CharLimit(256).Key("server_url").WithWidth(64),
		),
	).WithTheme(theme).WithWidth(64).WithShowHelp(false)

	return JoinServerScreen{
		form:          form,
		help:          help.New(),
		keys:          lib.QuitConfirmKeys,
		serverEntered: false,
	}
}

func (j JoinServerScreen) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, tea.SetWindowTitle("Eskimoe — Join A Server"))
	cmds = append(cmds, j.form.Init())
	return tea.Batch(cmds...)
}

func (j JoinServerScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, j.keys.Quit):
			return j, tea.Quit
		}

		switch msg.Type {
		case tea.KeyEnter:
			if j.serverEntered {
				err := database.JoinServer(globals.currentUser, j.server)

				if err != nil {
					log.Fatal("Error joining server:", err)
					return j, tea.Quit
				}

				globals.servers = database.GetServers(globals.currentUser)
				globals.currentUser.CurrentServer = j.server

				return screen().Switch(rootScreen())
			}
		case tea.KeyEscape:
			return screen().Switch(joinServerScreen())
		}
	}

	form, cmd := j.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		j.form = f
		cmds = append(cmds, cmd)
	}

	if j.form.State == huh.StateCompleted {
		serverURL := j.form.GetString("server_url")
		j.server = database.Server{
			URL: serverURL,
		}
		j.serverEntered = true
		j.serverInfo, _ = api.GetServerInfo(serverURL)
	}

	return j, tea.Batch(cmds...)
}

func (j JoinServerScreen) View() string {
	box := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.AdaptiveColor{Light: catppuccin.Latte.Rosewater().Hex, Dark: catppuccin.Mocha.Rosewater().Hex}).
		Foreground(lipgloss.AdaptiveColor{Light: catppuccin.Latte.Text().Hex, Dark: catppuccin.Mocha.Text().Hex}).
		Width(120).
		Padding(2)

	introductoryText := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.AdaptiveColor{Light: catppuccin.Latte.Mauve().Hex, Dark: catppuccin.Mocha.Mauve().Hex}).Render("Join a Server")

	var description string
	if len(globals.servers) == 0 {
		if !j.serverEntered {
			description = "You are not on any servers yet. Join a server by entering the server URL below."
		} else {
			description = "You are about to join a server. Press verify the server details and press enter to join."
		}
	} else {
		description = "You are on " + fmt.Sprintf("%d", len(globals.servers)) + " servers. Join another server by entering the server URL below."
	}
	descriptionText := lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: catppuccin.Latte.Rosewater().Hex, Dark: catppuccin.Mocha.Rosewater().Hex}).Render(description)

	var boxContent string
	if j.serverEntered {
		boxContent = fmt.Sprintf("%s\n\n%s", introductoryText, descriptionText)
	} else {
		formView := lipgloss.NewStyle().Margin(1, 0).Render(j.form.View())
		helpView := lipgloss.NewStyle().Margin(1, 0).Render(j.help.View(j.keys))
		boxContent = fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s", introductoryText, descriptionText, formView, helpView)
	}

	if j.serverEntered {
		helpView := lipgloss.NewStyle().Margin(1, 0).Render(j.help.View(lib.QuitConfirmCancelKeys))
		boxContent = box.Render(boxContent + "\n\n" +
			lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: catppuccin.Latte.Mauve().Hex, Dark: catppuccin.Mocha.Mauve().Hex}).Render("Server Name: ") +
			lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: catppuccin.Latte.Text().Hex, Dark: catppuccin.Mocha.Text().Hex}).Render(j.serverInfo.Name) + "\n" +
			lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: catppuccin.Latte.Mauve().Hex, Dark: catppuccin.Mocha.Mauve().Hex}).Render("Server Message: ") +
			lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: catppuccin.Latte.Text().Hex, Dark: catppuccin.Mocha.Text().Hex}).Render(j.serverInfo.Message) + "\n" +
			lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: catppuccin.Latte.Mauve().Hex, Dark: catppuccin.Mocha.Mauve().Hex}).Render("Server Version: ") +
			lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: catppuccin.Latte.Text().Hex, Dark: catppuccin.Mocha.Text().Hex}).Render(j.serverInfo.Version) + "\n" +
			lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: catppuccin.Latte.Mauve().Hex, Dark: catppuccin.Mocha.Mauve().Hex}).Render("Server URL: ") +
			lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: catppuccin.Latte.Text().Hex, Dark: catppuccin.Mocha.Text().Hex}).Render(j.server.URL) + "\n\n" +
			lipgloss.NewStyle().Render(helpView),
		)
	} else {
		boxContent = box.Render(boxContent)
	}

	return lipgloss.Place(
		globals.width,
		globals.height,
		lipgloss.Center,
		lipgloss.Center,
		boxContent,
	)
}
