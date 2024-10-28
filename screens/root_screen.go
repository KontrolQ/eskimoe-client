package screens

import (
	"encoding/json"
	"eskimoe-client/api"
	"eskimoe-client/database"
	"eskimoe-client/generators"
	"eskimoe-client/shared"
	"strings"
	"time"

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
	socketBroadcast             chan api.SocketBroadcast
	errorMessage                string
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
	ti.KeyMap.InsertNewline.SetEnabled(false)

	var errorMessage string
	currentServerInfo, err := api.GetAuthorizedServerInfo(globals.currentUser.CurrentServer.URL, globals.currentUser.AuthToken)
	if err != nil {
		errorMessage = err.Error()
	}

	categories := currentServerInfo.Categories
	if len(categories) == 0 {
		categories = append(categories, api.Category{
			Name: "Failed...",
			Rooms: []api.Room{
				{
					ID:   0,
					Name: "Failed to load server info",
				},
			},
		})
	}
	currentRoom := categories[0].Rooms[0]

	currentRoomsMessages, err := api.GetMessagesInRoom(globals.currentUser.CurrentServer.URL, currentRoom.ID, globals.currentUser.AuthToken)
	if err != nil {
		errorMessage = err.Error()
	}

	rootScreen := RootScreen{
		messagesViewPort:            messagesVP,
		currentServerInfo:           currentServerInfo,
		totalRoomCount:              shared.GetTotalRoomCount(categories),
		currentlyHighlightedArea:    "messageInput",
		currentlySelectedRoom:       currentRoom.ID,
		currentlyHighlightedMessage: len(currentRoomsMessages) - 1,
		messages:                    currentRoomsMessages,
		roomChanged:                 true,
		textarea:                    ti,
		errorMessage:                errorMessage,
	}

	rootScreen.ConnectToSocket()

	return rootScreen
}

// tea.Cmd to clear the error message after 3 seconds
func clearErrorMessage() tea.Cmd {
	return tea.Tick(time.Second*3, func(time.Time) tea.Msg {
		return clearError{}
	})
}

type clearError struct{}

func (r RootScreen) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle("Eskimoe"),
		textarea.Blink,
		clearErrorMessage(),
	)
}

func (r *RootScreen) ConnectToSocket() {
	headers := make(map[string][]string)
	headers["Authorization"] = []string{globals.currentUser.AuthToken}
	wsURL := strings.Replace(globals.currentUser.CurrentServer.URL, "http", "ws", 1) + "/ws/listen"
	wsConn, _, err := websocket.DefaultDialer.Dial(wsURL, headers)
	if err != nil {
		r.errorMessage = "failed to connect to the server's socket endpoint"
	}

	socketBroadcast := make(chan api.SocketBroadcast)

	if wsConn != nil {
		go func() {
			defer wsConn.Close()
			for {
				var broadcast api.SocketBroadcast
				err := wsConn.ReadJSON(&broadcast)
				if err != nil {
					r.errorMessage = "failed to read a recently received message"
					break
				}
				socketBroadcast <- broadcast
			}
		}()
	}

	r.wsConn = wsConn
	r.socketBroadcast = socketBroadcast
}

func (r *RootScreen) AttemptReconnect() {
	r.ConnectToSocket()
	currentServerInfo, err := api.GetAuthorizedServerInfo(globals.currentUser.CurrentServer.URL, globals.currentUser.AuthToken)
	if err == nil {
		r.currentServerInfo = currentServerInfo
		r.totalRoomCount = shared.GetTotalRoomCount(currentServerInfo.Categories)
		r.currentlySelectedRoom = currentServerInfo.Categories[0].Rooms[0].ID
		r.messages, _ = api.GetMessagesInRoom(globals.currentUser.CurrentServer.URL, r.currentlySelectedRoom, globals.currentUser.AuthToken)
		r.currentlyHighlightedMessage = len(r.messages) - 1
		r.roomChanged = true
		r.errorMessage = ""
	}
}

func (r RootScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		viewPortCommand tea.Cmd
		textAreaCommand tea.Cmd
	)

	totalMessages := len(r.messages)
	areas := []string{"sidebar", "mainArea", "messageInput"}

	// Handle WebSocket messages
	select {
	case message := <-r.socketBroadcast:
		databytes, _ := json.Marshal(message.Data)
		switch message.BroadcastType {
		case api.MessageCreated:
			var newMessage api.Message
			_ = json.Unmarshal(databytes, &newMessage)
			if newMessage.RoomID == r.currentlySelectedRoom {
				r.messages = append(r.messages, newMessage)
				totalMessages++
				if r.currentlyHighlightedMessage == totalMessages-2 {
					r.currentlyHighlightedMessage++
				}
			}

		case api.MessageDeleted:
			var deletedMessage api.DeletedMessage
			_ = json.Unmarshal(databytes, &deletedMessage)
			for i, m := range r.messages {
				if m.ID == deletedMessage.MessageID && m.RoomID == r.currentlySelectedRoom {
					r.messages = append(r.messages[:i], r.messages[i+1:]...)
					totalMessages--
					if r.currentlyHighlightedMessage == totalMessages-1 {
						r.currentlyHighlightedMessage--
					}
					break
				}
			}
		}
	default:
	}

	// Handle clearing the error message
	if _, ok := msg.(clearError); ok {
		r.errorMessage = ""
	}

	// Check if connection is still alive
	if r.wsConn != nil {
		err := r.wsConn.WriteMessage(websocket.PingMessage, nil)
		if err != nil {
			r.AttemptReconnect()
		}
	} else {
		r.AttemptReconnect()
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

	if r.currentlyHighlightedMessage >= totalMessages {
		r.currentlyHighlightedMessage = totalMessages - 1
	}
	if r.currentlyHighlightedMessage < 0 {
		r.currentlyHighlightedMessage = 0
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

			// d on a message to delete it
			if msg.String() == "d" {
				isMessageAuthor := r.messages[r.currentlyHighlightedMessage].Author.UID == globals.currentUser.UniqueID
				if !isMessageAuthor {
					r.errorMessage = "You can only delete your own messages"
				} else {
					if r.currentlyHighlightedMessage < len(r.messages) && len(r.messages) > 0 {
						r.errorMessage = "Deleting message..."
						messageId := r.messages[r.currentlyHighlightedMessage].ID
						err := api.DeleteMessageInRoom(globals.currentUser.CurrentServer.URL, r.currentlySelectedRoom, messageId, globals.currentUser.AuthToken)
						if err != nil {
							r.wsConn = nil
							r.errorMessage = "Failed to delete message"
						} else {
							r.errorMessage = ""
							r.messages = append(r.messages[:r.currentlyHighlightedMessage], r.messages[r.currentlyHighlightedMessage+1:]...)
						}
					}
				}
			}
		}

		// Message Input:
		if r.currentlyHighlightedArea != "messageInput" {
			if r.textarea.Focused() {
				r.textarea.Blur()
			}
		} else {
			if msg.String() == "alt+enter" {
				// Shift+Enter to insert a newline
				r.textarea.SetValue(r.textarea.Value() + "\n")
			}

			if msg.String() == "enter" {
				r.errorMessage = "Sending message..."
				if strings.TrimSpace(r.textarea.Value()) != "" {
					newMessage := api.SendRoomMessage{
						Content: r.textarea.Value(),
					}
					err := api.SendMessageToRoom(globals.currentUser.CurrentServer.URL, r.currentlySelectedRoom, globals.currentUser.AuthToken, newMessage)

					if err != nil {
						r.wsConn = nil
					} else {
						r.errorMessage = ""
						r.currentlyHighlightedMessage = len(r.messages)
						r.textarea.SetValue("")
					}
				}
				r.errorMessage = ""
			}

			if !r.textarea.Focused() {
				r.textarea.Focus()
			}
		}

		// ctrl+c to quit the application
		if msg.String() == "ctrl+c" {
			return r, tea.Quit
		}

		// ctrl+l to leave the current server
		if msg.String() == "ctrl+l" {
			r.errorMessage = "Leaving server..."
			user, err := database.LeaveServer(globals.currentUser, globals.currentUser.CurrentServer)
			if err != nil {
				r.errorMessage = "Failed to leave server: " + err.Error()
			} else {
				globals.currentUser = user
				if user.CurrentServer.ID == 0 {
					globals.servers = database.GetServers(globals.currentUser)
					return screen().Switch(joinServerScreen())
				} else {
					return screen().Switch(rootScreen())
				}
			}
		}
	}

	// Generate the viewport content based on the updated messages
	viewPortString := generators.MessageViewBuilder(r.messages, r.currentlyHighlightedMessage, r.messagesViewPort.Width, r.currentlyHighlightedArea == "mainArea")

	// Set the content of the viewport
	r.messagesViewPort.SetContent(viewPortString)

	if r.roomChanged || r.currentlyHighlightedMessage == totalMessages-1 {
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

	statusBar := generators.StatusBarView(r.errorMessage, r.currentlyHighlightedArea, globals.width)

	sidebar := generators.SidebarView(r.currentServerInfo.Categories, r.currentlySelectedRoom, globals.height, topBar, statusBar, r.currentlyHighlightedArea == "sidebar")

	memberList := generators.MemberListView(r.currentServerInfo.Members, globals.height, topBar, statusBar, r.currentlyHighlightedArea == "memberList")

	messageInput := generators.MessageInputView(r.textarea.View(), globals.width, r.currentlyHighlightedArea == "messageInput")

	mainArea := generators.MainAreaView(r.messagesViewPort.View(), globals.width, globals.height, topBar, statusBar, r.currentlyHighlightedArea == "mainArea")

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
			memberList,
		),
		statusBar,
	)

	return rootView
}
