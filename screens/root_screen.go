package screens

import (
	"eskimoe-client/data"
	"eskimoe-client/generators"
	"eskimoe-client/shared"
	"time"

	"github.com/charmbracelet/bubbles/textarea"
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
	textarea                    textarea.Model
}

func rootScreen() tea.Model {
	generators.InitializeStyles(globals.currentUser)

	vw, vh := generators.GetViewportDimensions(globals.width, globals.height)
	messagesVP := viewport.New(vw, vh)
	messagesVP.KeyMap = viewport.KeyMap{}

	ti := textarea.New()
	ti.Placeholder = "Write Something Epic... (Ctrl + S to send)"
	ti.Focus()
	ti.SetWidth(vw - 4)
	ti.SetHeight(shared.DefaultPreferences.MessageInputHeight)
	ti.CharLimit = 4096
	ti.ShowLineNumbers = false

	return RootScreen{
		messagesViewPort:            messagesVP,
		currentlyHighlightedArea:    "messageInput",
		currentlySelectedRoom:       1,
		currentlyHighlightedMessage: len(data.GeneralRoomMessages) - 1,
		messages:                    data.GeneralRoomMessages,
		roomChanged:                 true,
		textarea:                    ti,
	}
}

func (r RootScreen) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle("Eskimoe"),
		textarea.Blink,
	)
}

func (r RootScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		viewPortCommand tea.Cmd
		textAreaCommand tea.Cmd
	)

	areas := []string{"sidebar", "mainArea", "messageInput"}

	// OnWindowResize
	if msg, ok := msg.(tea.WindowSizeMsg); ok {
		globals.width = msg.Width
		globals.height = msg.Height

		vw, vh := generators.GetViewportDimensions(globals.width, globals.height)
		r.messagesViewPort.Width = vw
		r.messagesViewPort.Height = vh

		r.textarea.SetWidth(vw - 4)
		r.textarea.SetHeight(shared.DefaultPreferences.MessageInputHeight)
	}

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

				r.messages = data.GetRoomMessages(r.currentlySelectedRoom)
				r.currentlyHighlightedMessage = len(r.messages) - 1
				r.roomChanged = true
			}

			if msg.String() == "k" || msg.String() == "up" {
				if r.currentlySelectedRoom == 1 {
					r.currentlySelectedRoom = 10
				} else {
					r.currentlySelectedRoom--
				}

				r.messages = data.GetRoomMessages(r.currentlySelectedRoom)
				r.currentlyHighlightedMessage = len(r.messages) - 1
				r.roomChanged = true
			}
		}

		// Message Navigation: j/k or arrow keys to navigate
		if r.currentlyHighlightedArea == "mainArea" {
			if msg.String() == "j" || msg.String() == "down" {
				if r.currentlyHighlightedMessage < len(r.messages)-1 {
					r.currentlyHighlightedMessage++
					r.messagesViewPort.LineDown(3)
				}
			}

			if msg.String() == "k" || msg.String() == "up" {
				if r.currentlyHighlightedMessage > 0 {
					r.currentlyHighlightedMessage--
					r.messagesViewPort.LineUp(3)
				}
			}
		}

		// Message Input:
		if r.currentlyHighlightedArea != "messageInput" {
			if r.textarea.Focused() {
				r.textarea.Blur()
			}
		} else {
			if msg.String() == "ctrl+s" {
				if r.textarea.Value() != "" {
					newMessage := shared.Message{
						Author:    globals.currentUser.DisplayName,
						Content:   r.textarea.Value(),
						Reactions: []shared.Reactions{},
						SentAt:    time.Now(),
					}

					r.messages = append(r.messages, newMessage)
					r.currentlyHighlightedMessage = len(r.messages) - 1
					r.roomChanged = true
					r.textarea.SetValue("")
				}
			}

			if !r.textarea.Focused() {
				r.textarea.Focus()
			}
		}

		// ctrl+c to quit the application
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
	r.textarea, textAreaCommand = r.textarea.Update(msg)

	return r, tea.Batch(viewPortCommand, textAreaCommand)
}

func (r RootScreen) View() string {
	currentRoomName := data.GetRoomName(r.currentlySelectedRoom)

	topBar := generators.TopBarView("Eskimoe", currentRoomName, globals.currentUser.DisplayName, globals.width)

	sidebar := generators.SidebarView(data.Categories, r.currentlySelectedRoom, globals.height, topBar, r.currentlyHighlightedArea == "sidebar")

	messageInput := generators.MessageInputView(r.textarea.View(), globals.width, r.currentlyHighlightedArea == "messageInput")

	mainArea := generators.MainAreaView(r.messagesViewPort.View(), globals.width, globals.height, topBar, r.currentlyHighlightedArea == "mainArea")

	rootView := lipgloss.JoinVertical(
		lipgloss.Top,
		topBar,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			sidebar,
			lipgloss.JoinVertical(
				lipgloss.Top,
				mainArea,
				messageInput,
			),
		),
	)

	return rootView
}
