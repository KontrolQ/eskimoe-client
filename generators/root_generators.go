package generators

import (
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wrap"

	"eskimoe-client/database"
	"eskimoe-client/shared"
)

var Style shared.StyleSheet
var Preferences database.UserPreferences

func InitializeStyles(User database.User) {
	Preferences = User.Preferences
	Style = GenerateStyles(Preferences)
}

func GetViewportDimensions(width, height int) (int, int) {
	w := width - Preferences.SidebarWidth - 4
	h := height - Preferences.MessageInputHeight - 6

	return w, h
}

func SingleMessageView(message shared.Message, width int) string {
	spacer := Style.BaseStyle.Render(" ")

	author := Style.MessageAuthorStyle.Render(message.Author)

	var timestring string

	if message.SentAt.Day() == time.Now().Day() {
		timestring = message.SentAt.Format("15:04")
	} else {
		timestring = message.SentAt.Format("2006-01-02 15:04")
	}

	time := Style.MessageAuthorStyle.Padding(0, 1).Render(timestring)

	var reactionsString string

	for _, reaction := range message.Reactions {
		reactionsString += Style.BaseStyle.Render("[ ")
		reactionsString += Style.BaseStyle.
			Foreground(reaction.Color).
			Render(reaction.Reaction)
		reactionsString += spacer
		reactionsString += Style.BaseStyle.
			Foreground(reaction.Color).
			Render(strconv.Itoa(reaction.Count))
		reactionsString += Style.BaseStyle.Render(" ] ")
	}

	reactions := Style.BaseStyle.Padding(0, 1).Render(reactionsString)

	content := wrap.String(message.Content, width-4)

	content = Style.MessageTextStyle.Render(content)

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
		centralizedMessage := Style.BaseStyle.Render("No messages in this room!")
		return Style.BaseStyle.Padding(0, 0, 0, (width-lipgloss.Width(centralizedMessage))/2).Render(centralizedMessage)
	}

	for i, message := range messages {
		messageView := SingleMessageView(message, width)

		// Add border to currently highlighted message

		if i == currentlyHighlightedMessage && mainAreaHighlighted {
			messageView = Style.HighlightedMessageBorderStyle.Render(messageView)
		}

		messagesArray = append(messagesArray, messageView)
	}

	return strings.Join(messagesArray, "\n\n")
}

func TopBarView(serverName, currentRoomName, currentUser string, width int) string {
	w := lipgloss.Width

	spacer := Style.BaseStyle.Render(" ")

	serverText := Style.ServerTextStyle.Bold(true).Render("Server:")
	serverNameText := Style.ServerTextStyle.Render(serverName)

	roomText := Style.RoomTextStyle.Bold(true).Render("Room:")
	roomNameText := Style.RoomTextStyle.Render(currentRoomName)

	userAreaText := Style.UserInfoStyle.Bold(true).Render("User:")
	userText := Style.UserInfoStyle.Render(currentUser)

	separator := Style.BaseStyle.Width(width - w(serverText) - w(serverNameText) - w(roomText) - w(roomNameText) - w(userAreaText) - w(userText) - w(spacer)).Render("")

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
			categoryText = Style.CategoryListTextStyle.Render(category.Name)
		} else {
			categoryText = Style.CategoryListTextStyle.Margin(1, 0, 0, 0).Render(category.Name)
		}

		sidebarDoc.WriteString(categoryText + "\n")

		for _, room := range category.Rooms {
			roomText := Style.RoomListTextStyle.MarginLeft(2).Render(room.Name)
			if room.RoomId == currentlySelectedRoom {
				roomText = Style.HighlightedRoomTextStyle.MarginLeft(2).Render(room.Name)
			}
			sidebarDoc.WriteString(roomText + "\n")
		}

	}

	var sidebar string

	if currentlyHighlighted {
		sidebar = Style.HighlightedAreaStyle.Padding(0, 1).
			Width(Preferences.SidebarWidth).
			Height(height - h(topBar) - 2).
			Render(sidebarDoc.String())
	} else {
		sidebar = Style.NormalAreaStyle.Padding(0, 1).
			Width(Preferences.SidebarWidth).
			Height(height - h(topBar) - 2).
			Render(sidebarDoc.String())
	}

	return sidebar
}

func MainAreaView(viewport string, width, height int, topBar string, currentlyHighlighted bool) string {
	h := lipgloss.Height

	var mainArea string

	if currentlyHighlighted {
		mainArea = Style.HighlightedAreaStyle.Padding(0, 1).
			Width(width - Preferences.SidebarWidth - 4).
			Height(height - h(topBar) - Preferences.MessageInputHeight - 4).
			Render(viewport)
	} else {
		mainArea = Style.NormalAreaStyle.Padding(0, 1).
			Width(width - Preferences.SidebarWidth - 4).
			Height(height - h(topBar) - Preferences.MessageInputHeight - 4).
			Render(viewport)
	}

	return mainArea
}

func MessageInputView(textareaView string, width int, currentlyHighlighted bool) string {
	var messageInputDoc strings.Builder

	messageInputDoc.WriteString(Style.BaseStyle.Render(textareaView))

	var messageInput string

	if currentlyHighlighted {
		messageInput = Style.HighlightedAreaStyle.Padding(0, 1).
			Width(width - Preferences.SidebarWidth - 4).
			Height(Preferences.MessageInputHeight).
			Render(messageInputDoc.String())
	} else {
		messageInput = Style.NormalAreaStyle.Padding(0, 1).
			Width(width - Preferences.SidebarWidth - 4).
			Height(Preferences.MessageInputHeight).
			Render(messageInputDoc.String())
	}

	return messageInput
}
