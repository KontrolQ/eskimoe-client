package screens

import (
	"encoding/json"
	"eskimoe-client/api"
	"eskimoe-client/generators"
	"eskimoe-client/shared"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gorilla/websocket"
)

type RootScreen struct {
	messagesViewPort            viewport.Model
	currentServerInfo           api.ServerInfoAuthorized
	totalRoomCount              int
	currentlyHighlightedArea    string
	currentlySelectedRoom       int
	currentlyHighlightedMessage int
	messages                    []api.Message
	roomChanged                 bool
	textarea                    textarea.Model
	wsConn                      *websocket.Conn
	messageChannel              chan api.Message
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

	currentServerInfo, err := api.GetAuthorizedServerInfo(globals.currentUser.CurrentServer.URL, globals.currentUser.AuthToken)
	if err != nil {
		log.Fatal(err)
	}

	categories := currentServerInfo.Categories
	currentRoom := categories[0].Rooms[0]

	currentRoomsMessages, err := api.GetMessagesInRoom(globals.currentUser.CurrentServer.URL, currentRoom.ID, globals.currentUser.AuthToken)
	if err != nil {
		log.Fatal(err)
	}

	wsURL := strings.Replace(globals.currentUser.CurrentServer.URL, "http", "ws", 1) + "/ws/connect"
	wsConn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		log.Fatal("Failed to connect to WebSocket:", err)
	}

	messageChannel := make(chan api.Message)

	go func() {
		defer wsConn.Close()
		for {
			_, msg, err := wsConn.ReadMessage()
			if err != nil {
				log.Println("Error reading WebSocket message:", err)
				return
			}
			var message api.Message
			err = json.Unmarshal(msg, &message)
			if err != nil {
				log.Println("Error unmarshalling WebSocket message:", err)
				continue
			}
			messageChannel <- message
		}
	}()

	return RootScreen{
		messagesViewPort:            messagesVP,
		currentServerInfo:           currentServerInfo,
		totalRoomCount:              shared.GetTotalRoomCount(categories),
		currentlyHighlightedArea:    "messageInput",
		currentlySelectedRoom:       currentRoom.ID,
		currentlyHighlightedMessage: 0,
		messages:                    currentRoomsMessages,
		roomChanged:                 true,
		textarea:                    ti,
		wsConn:                      wsConn,
		messageChannel:              messageChannel,
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
	// Handle WebSocket messages
	select {
	case message := <-r.messageChannel:
		// Append message to messages slice
		r.messages = append(r.messages, message)
		messageLen := len(r.messages)
		if r.currentlyHighlightedMessage == messageLen-2 {
			r.currentlyHighlightedMessage++
		}

		if r.currentlyHighlightedMessage == messageLen-1 {
			r.roomChanged = true
		}
	default:
	}

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
				r.currentlySelectedRoom = shared.GetNextRoomId(r.currentServerInfo.Categories, r.currentlySelectedRoom)
				r.messages, _ = api.GetMessagesInRoom(globals.currentUser.CurrentServer.URL, r.currentlySelectedRoom, globals.currentUser.AuthToken)
				r.currentlyHighlightedMessage = len(r.messages) - 1
				r.roomChanged = true
			}

			if msg.String() == "k" || msg.String() == "up" {
				r.currentlySelectedRoom = shared.GetPreviousRoomId(r.currentServerInfo.Categories, r.currentlySelectedRoom)
				r.messages, _ = api.GetMessagesInRoom(globals.currentUser.CurrentServer.URL, r.currentlySelectedRoom, globals.currentUser.AuthToken)
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
				if strings.TrimSpace(r.textarea.Value()) != "" {
					newMessage := api.SendRoomMessage{
						Content: r.textarea.Value(),
					}
					err := api.SendMessageToRoom(globals.currentUser.CurrentServer.URL, r.currentlySelectedRoom, globals.currentUser.AuthToken, newMessage)

					if err != nil {
						log.Fatal(err)
					}

					// r.messages, _ = api.GetMessagesInRoom(globals.currentUser.CurrentServer.URL, r.currentlySelectedRoom, globals.currentUser.AuthToken)
					r.currentlyHighlightedMessage = len(r.messages)
					// r.roomChanged = true

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

	// Generate the viewport content based on the updated messages
	viewPortString := generators.MessageViewBuilder(r.messages, r.currentlyHighlightedMessage, r.messagesViewPort.Width, r.currentlyHighlightedArea == "mainArea")

	// Set the content of the viewport
	r.messagesViewPort.SetContent(viewPortString)

	if r.roomChanged {
		r.messagesViewPort.GotoBottom()
		r.roomChanged = false
	}

	// Update viewport and textarea models
	r.messagesViewPort, viewPortCommand = r.messagesViewPort.Update(msg)
	r.textarea, textAreaCommand = r.textarea.Update(msg)

	return r, tea.Batch(viewPortCommand, textAreaCommand)
}

func (r RootScreen) View() string {
	currentRoomName := shared.GetRoomNameFromId(r.currentServerInfo.Categories, r.currentlySelectedRoom)

	topBar := generators.TopBarView("Eskimoe", currentRoomName, globals.currentUser.DisplayName, globals.width)

	sidebar := generators.SidebarView(r.currentServerInfo.Categories, r.currentlySelectedRoom, globals.height, topBar, r.currentlyHighlightedArea == "sidebar")

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
