package shared

import (
	"time"

	"github.com/charmbracelet/lipgloss"
)

type ColorPreference struct {
	Light string
	Dark  string
}

type Preferences struct {
	AreaNormal                ColorPreference
	AreaHighlight             ColorPreference
	CategoryListText          ColorPreference
	HighLightedMessageBorder  ColorPreference
	HighlightedRoomText       ColorPreference
	HighlightedRoomBackground ColorPreference
	MessageAuthor             ColorPreference
	MessageText               ColorPreference
	RoomListText              ColorPreference
	RoomStatusText            ColorPreference
	RoomStatusBackground      ColorPreference
	ServerStatusText          ColorPreference
	ServerStatusBackground    ColorPreference
	UserBarText               ColorPreference
	UserBarBackground         ColorPreference
	ErrorBackground           ColorPreference
	StatusBarBackground       ColorPreference
	SidebarWidth              int
	MessageInputHeight        int
}

type StyleSheet struct {
	BaseStyle                     lipgloss.Style
	ServerTextStyle               lipgloss.Style
	RoomTextStyle                 lipgloss.Style
	UserInfoStyle                 lipgloss.Style
	NormalAreaStyle               lipgloss.Style
	HighlightedAreaStyle          lipgloss.Style
	CategoryListTextStyle         lipgloss.Style
	HighlightedMessageBorderStyle lipgloss.Style
	RoomListTextStyle             lipgloss.Style
	HighlightedRoomTextStyle      lipgloss.Style
	MessageAuthorStyle            lipgloss.Style
	MessageTextStyle              lipgloss.Style
	ErrorTextStyle                lipgloss.Style
	StatusBarStyle                lipgloss.Style
}

type Reactions struct {
	Reaction string
	Count    int
	Users    []string
	Color    lipgloss.AdaptiveColor
}

type Message struct {
	Author    string
	Content   string
	SentAt    time.Time
	Reactions []Reactions
}

type Room struct {
	Name     string
	RoomId   int
	Messages []Message
}

type Category struct {
	Name  string
	Rooms []Room
}
