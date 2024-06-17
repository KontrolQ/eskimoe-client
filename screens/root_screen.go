package screens

import (
	"eskimoe-client/data"
	"eskimoe-client/generators"
	"eskimoe-client/shared"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type RootScreen struct {
	messagesViewPort            viewport.Model
	currentlyHighlightedArea    string
	currentlySelectedRoom       int
	currentlyHighlightedMessage int
	messages                    []shared.Message
	roomChanged                 bool
}

func rootScreen() tea.Model {
	vw, vh := generators.GetViewportDimensions(globals.width, globals.height)
	messagesVP := viewport.New(vw, vh)

	return RootScreen{
		messagesViewPort:            messagesVP,
		currentlyHighlightedArea:    "messageInput",
		currentlySelectedRoom:       1,
		currentlyHighlightedMessage: len(data.GeneralRoomMessages) - 1,
		messages:                    data.GeneralRoomMessages,
		roomChanged:                 true,
	}
}

func (r RootScreen) Init() tea.Cmd {
	return tea.SetWindowTitle("Eskimoe")
}

func (r RootScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		viewPortCommand tea.Cmd
	)

	areas := []string{"sidebar", "mainArea", "messageInput"}

	if msg, ok := msg.(tea.KeyMsg); ok {
		// Area Switching: Tab and Shift+Tab to switch between areas
		if msg.String() == "tab" {
			for i, area := range areas {
				if r.currentlyHighlightedArea == area {
					if i == len(areas)-1 {
						r.currentlyHighlightedArea = areas[0]
					} else {
						r.currentlyHighlightedArea = areas[i+1]
					}
					break
				}
			}
		}

		if msg.String() == "shift+tab" {
			for i, area := range areas {
				if r.currentlyHighlightedArea == area {
					if i == 0 {
						r.currentlyHighlightedArea = areas[len(areas)-1]
					} else {
						r.currentlyHighlightedArea = areas[i-1]
					}
					break
				}
			}
		}

		// Sidebar Navigation: j/k or arrow keys to navigate
		if r.currentlyHighlightedArea == "sidebar" {
			if msg.String() == "j" || msg.String() == "down" {
				if r.currentlySelectedRoom == 10 {
					r.currentlySelectedRoom = 1
				} else {
					r.currentlySelectedRoom++
				}
			}

			if msg.String() == "k" || msg.String() == "up" {
				if r.currentlySelectedRoom == 1 {
					r.currentlySelectedRoom = 10
				} else {
					r.currentlySelectedRoom--
				}
			}

			// Load messages for the selected room
			r.messages = data.GetRoomMessages(r.currentlySelectedRoom)

			r.currentlyHighlightedMessage = len(r.messages) - 1
			r.roomChanged = true
		}

		// Message Navigation: j/k or arrow keys to navigate
		if r.currentlyHighlightedArea == "mainArea" {
			if msg.String() == "j" || msg.String() == "down" {
				if r.currentlyHighlightedMessage < len(r.messages)-1 {
					r.currentlyHighlightedMessage++
				}
			}

			if msg.String() == "k" || msg.String() == "up" {
				if r.currentlyHighlightedMessage > 0 {
					r.currentlyHighlightedMessage--
				}
			}
		}

		// Message Input: ctrl+c to quit the application
		if msg.String() == "ctrl+c" {
			return r, tea.Quit
		}
	}

	viewPortString := generators.MessageViewBuilder(r.messages, r.currentlyHighlightedMessage, r.messagesViewPort.Width, r.currentlyHighlightedArea == "mainArea")

	r.messagesViewPort.SetContent(viewPortString)

	if r.roomChanged {
		r.messagesViewPort.GotoBottom()
		r.roomChanged = false
	}

	r.messagesViewPort, viewPortCommand = r.messagesViewPort.Update(msg)

	return r, tea.Batch(viewPortCommand)
}

func (r RootScreen) View() string {
	doc := strings.Builder{}

	currentRoomName := data.GetRoomName(r.currentlySelectedRoom)

	topBar := generators.TopBarView("Eskimoe", currentRoomName, globals.currentUser.DisplayName, globals.width)

	sidebar := generators.SidebarView(data.Categories, r.currentlySelectedRoom, globals.height, topBar, r.currentlyHighlightedArea == "sidebar")

	messageInput := generators.MessageInputView(globals.width, r.currentlyHighlightedArea == "messageInput")

	mainArea := generators.MainAreaView(r.messagesViewPort.View(), globals.width, globals.height, topBar, r.currentlyHighlightedArea == "mainArea")

	chatArea := lipgloss.JoinVertical(lipgloss.Top, mainArea, messageInput)

	mainW := lipgloss.JoinHorizontal(lipgloss.Top, sidebar, chatArea)

	doc.WriteString(topBar + "\n" + mainW)

	return doc.String()
}
