package generators

import (
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wrap"

	"eskimoe-client/shared"
)

func GetViewportDimensions(width, height int) (int, int) {
	w := width - shared.DefaultPreferences.SidebarWidth - 4
	h := height - shared.DefaultPreferences.MessageInputHeight - 6

	return w, h
}

func SingleMessageView(message shared.Message, width int) string {
	spacer := shared.BaseStyle.Render(" ")

	author := shared.MessageAuthorStyle.Render(message.Author)

	var timestring string

	if message.SentAt.Day() == time.Now().Day() {
		timestring = message.SentAt.Format("15:04")
	} else {
		timestring = message.SentAt.Format("2006-01-02 15:04")
	}

	time := shared.MessageAuthorStyle.Padding(0, 1).Render(timestring)

	var reactionsString string

	for _, reaction := range message.Reactions {
		reactionsString += shared.BaseStyle.Render("[ ")
		reactionsString += shared.BaseStyle.
			Foreground(reaction.Color).
			Render(reaction.Reaction)
		reactionsString += spacer
		reactionsString += shared.BaseStyle.
			Foreground(reaction.Color).
			Render(strconv.Itoa(reaction.Count))
		reactionsString += shared.BaseStyle.Render(" ] ")
	}

	reactions := shared.BaseStyle.Padding(0, 1).Render(reactionsString)

	content := wrap.String(message.Content, width-4)

	content = shared.MessageTextStyle.Render(content)

	m := lipgloss.JoinHorizontal(
		lipgloss.Top,
		author,
		time,
		reactions,
	)

	m = lipgloss.JoinVertical(
		lipgloss.Top,
		m,
		content,
	)

	return m
}

func MessageViewBuilder(messages []shared.Message, currentlyHighlightedMessage int, width int, mainAreaHighlighted bool) string {
	var messagesArray []string

	if len(messages) == 0 {
		centralizedMessage := shared.BaseStyle.Render("No messages in this room!")
		return shared.BaseStyle.Padding(0, 0, 0, (width-lipgloss.Width(centralizedMessage))/2).Render(centralizedMessage)
	}

	for i, message := range messages {
		messageView := SingleMessageView(message, width)

		// Add border to currently highlighted message

		if i == currentlyHighlightedMessage && mainAreaHighlighted {
			messageView = shared.HighlightedMessageBorderStyle.Render(messageView)
		}

		messagesArray = append(messagesArray, messageView)
	}

	return strings.Join(messagesArray, "\n\n")
}

func TopBarView(serverName, currentRoomName, currentUser string, width int) string {
	w := lipgloss.Width

	spacer := shared.BaseStyle.Render(" ")

	serverText := shared.ServerTextStyle.Bold(true).Render("Server: ")
	serverNameText := shared.ServerTextStyle.Render(serverName)

	roomText := shared.RoomTextStyle.Bold(true).Render("Room: ")
	roomNameText := shared.RoomTextStyle.Render(currentRoomName)

	userAreaText := shared.UserInfoStyle.Bold(true).Render("User: ")
	userText := shared.UserInfoStyle.Render(currentUser)

	separator := shared.BaseStyle.Width(width - w(serverText) - w(serverNameText) - w(roomText) - w(roomNameText) - w(userAreaText) - w(userText) - w(spacer)).Render("")

	topBar := lipgloss.JoinHorizontal(
		lipgloss.Top,
		serverText,
		serverNameText,
		spacer,
		roomText,
		roomNameText,
		separator,
		userAreaText,
		userText,
	)

	return topBar
}

func SidebarView(categories []shared.Category, currentlySelectedRoom int, height int, topBar string, currentlyHighlighted bool) string {
	h := lipgloss.Height

	var sidebarDoc strings.Builder

	for i, category := range categories {
		var categoryText string

		if i == 0 {
			categoryText = shared.CategoryListTextStyle.Render(category.Name)
		} else {
			categoryText = shared.CategoryListTextStyle.Margin(1, 0, 0, 0).Render(category.Name)
		}

		sidebarDoc.WriteString(categoryText + "\n")

		for _, room := range category.Rooms {
			roomText := shared.RoomListTextStyle.MarginLeft(2).Render(room.Name)
			if room.RoomId == currentlySelectedRoom {
				roomText = shared.HighlightedRoomTextStyle.MarginLeft(2).Render(room.Name)
			}
			sidebarDoc.WriteString(roomText + "\n")
		}

	}

	var sidebar string

	if currentlyHighlighted {
		sidebar = shared.HighlightedAreaStyle.Padding(0, 1).
			Width(shared.DefaultPreferences.SidebarWidth).
			Height(height - h(topBar) - 2).
			Render(sidebarDoc.String())
	} else {
		sidebar = shared.NormalAreaStyle.Padding(0, 1).
			Width(shared.DefaultPreferences.SidebarWidth).
			Height(height - h(topBar) - 2).
			Render(sidebarDoc.String())
	}

	return sidebar
}

func MainAreaView(viewport string, width, height int, topBar string, currentlyHighlighted bool) string {
	h := lipgloss.Height

	var mainArea string

	if currentlyHighlighted {
		mainArea = shared.HighlightedAreaStyle.Padding(0, 1).
			Width(width - shared.DefaultPreferences.SidebarWidth - 4).
			Height(height - h(topBar) - shared.DefaultPreferences.MessageInputHeight - 4).
			Render(viewport)
	} else {
		mainArea = shared.NormalAreaStyle.Padding(0, 1).
			Width(width - shared.DefaultPreferences.SidebarWidth - 4).
			Height(height - h(topBar) - shared.DefaultPreferences.MessageInputHeight - 4).
			Render(viewport)
	}

	return mainArea
}

func MessageInputView(width int, currentlyHighlighted bool) string {
	var messageInputDoc strings.Builder

	messageInputDoc.WriteString(shared.BaseStyle.Render("Message: "))

	var messageInput string

	if currentlyHighlighted {
		messageInput = shared.HighlightedAreaStyle.Padding(0, 1).
			Width(width - shared.DefaultPreferences.SidebarWidth - 4).
			Height(shared.DefaultPreferences.MessageInputHeight).
			Render(messageInputDoc.String())
	} else {
		messageInput = shared.NormalAreaStyle.Padding(0, 1).
			Width(width - shared.DefaultPreferences.SidebarWidth - 4).
			Height(shared.DefaultPreferences.MessageInputHeight).
			Render(messageInputDoc.String())
	}

	return messageInput
}
