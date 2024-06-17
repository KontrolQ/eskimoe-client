package shared

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	BaseStyle = lipgloss.NewStyle()

	ServerTextStyle = BaseStyle.
			Padding(0, 1).
			Foreground(lipgloss.AdaptiveColor{
			Light: DefaultPreferences.ServerStatusText.Light,
			Dark:  DefaultPreferences.ServerStatusText.Dark,
		}).
		Background(lipgloss.AdaptiveColor{
			Light: DefaultPreferences.ServerStatusBackground.Light,
			Dark:  DefaultPreferences.ServerStatusBackground.Dark,
		})

	RoomTextStyle = BaseStyle.
			Padding(0, 1).
			Foreground(lipgloss.AdaptiveColor{
			Light: DefaultPreferences.RoomStatusText.Light,
			Dark:  DefaultPreferences.RoomStatusText.Dark,
		}).
		Background(lipgloss.AdaptiveColor{
			Light: DefaultPreferences.RoomStatusBackground.Light,
			Dark:  DefaultPreferences.RoomStatusBackground.Dark,
		})

	UserInfoStyle = BaseStyle.
			Padding(0, 1).
			Foreground(lipgloss.AdaptiveColor{
			Light: DefaultPreferences.UserBarText.Light,
			Dark:  DefaultPreferences.UserBarText.Dark,
		}).
		Background(lipgloss.AdaptiveColor{
			Light: DefaultPreferences.UserBarBackground.Light,
			Dark:  DefaultPreferences.UserBarBackground.Dark,
		})

	NormalAreaStyle = BaseStyle.
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.AdaptiveColor{
			Light: DefaultPreferences.AreaNormal.Light,
			Dark:  DefaultPreferences.AreaNormal.Dark,
		})

	HighlightedAreaStyle = BaseStyle.
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.AdaptiveColor{
			Light: DefaultPreferences.AreaHighlight.Light,
			Dark:  DefaultPreferences.AreaHighlight.Dark,
		})

	CategoryListTextStyle = BaseStyle.
				Foreground(lipgloss.AdaptiveColor{
			Light: DefaultPreferences.CategoryListText.Light,
			Dark:  DefaultPreferences.CategoryListText.Dark,
		})

	HighlightedMessageBorderStyle = BaseStyle.
					BorderStyle(lipgloss.NormalBorder()).
					BorderBackground(lipgloss.AdaptiveColor{
			Light: DefaultPreferences.HighLightedMessageBorder.Light,
			Dark:  DefaultPreferences.HighLightedMessageBorder.Dark,
		}).
		BorderForeground(lipgloss.AdaptiveColor{
			Light: DefaultPreferences.HighLightedMessageBorder.Light,
			Dark:  DefaultPreferences.HighLightedMessageBorder.Dark,
		}).
		PaddingLeft(1).
		BorderLeft(true).
		BorderRight(false).
		BorderTop(false).
		BorderBottom(false).
		Foreground(lipgloss.AdaptiveColor{
			Light: DefaultPreferences.MessageText.Light,
			Dark:  DefaultPreferences.MessageText.Dark,
		})

	RoomListTextStyle = BaseStyle.
				Foreground(lipgloss.AdaptiveColor{
			Light: DefaultPreferences.RoomListText.Light,
			Dark:  DefaultPreferences.RoomListText.Dark,
		})

	HighlightedRoomTextStyle = BaseStyle.
					Foreground(lipgloss.AdaptiveColor{
			Light: DefaultPreferences.HighlightedRoomText.Light,
			Dark:  DefaultPreferences.HighlightedRoomText.Dark,
		}).
		Background(lipgloss.AdaptiveColor{
			Light: DefaultPreferences.HighlightedRoomBackground.Light,
			Dark:  DefaultPreferences.HighlightedRoomBackground.Dark,
		})

	MessageAuthorStyle = BaseStyle.
				Foreground(lipgloss.AdaptiveColor{
			Light: DefaultPreferences.MessageAuthor.Light,
			Dark:  DefaultPreferences.MessageAuthor.Dark,
		})

	MessageTextStyle = BaseStyle.
				Foreground(lipgloss.AdaptiveColor{
			Light: DefaultPreferences.MessageText.Light,
			Dark:  DefaultPreferences.MessageText.Dark,
		})
)
