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
	SidebarWidth              int
	MessageInputHeight        int
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
