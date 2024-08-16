package generators

import (
	"eskimoe-client/database"
	"eskimoe-client/shared"

	"github.com/charmbracelet/lipgloss"
)

func GenerateStyles(pref database.UserPreferences) shared.StyleSheet {
	BaseStyle := lipgloss.NewStyle()

	return shared.StyleSheet{
		BaseStyle: BaseStyle,

		ServerTextStyle: BaseStyle.
			Padding(0, 1).
			Foreground(lipgloss.AdaptiveColor{
				Light: pref.ServerStatusTextLight,
				Dark:  pref.ServerStatusTextDark,
			}).
			Background(lipgloss.AdaptiveColor{
				Light: pref.ServerStatusBackgroundLight,
				Dark:  pref.ServerStatusBackgroundDark,
			}),

		RoomTextStyle: BaseStyle.
			Padding(0, 1).
			Foreground(lipgloss.AdaptiveColor{
				Light: pref.RoomStatusTextLight,
				Dark:  pref.RoomStatusTextDark,
			}).
			Background(lipgloss.AdaptiveColor{
				Light: pref.RoomStatusBackgroundLight,
				Dark:  pref.RoomStatusBackgroundDark,
			}),
		UserInfoStyle: BaseStyle.
			Padding(0, 1).
			Foreground(lipgloss.AdaptiveColor{
				Light: pref.UserBarTextLight,
				Dark:  pref.UserBarTextDark,
			}).
			Background(lipgloss.AdaptiveColor{
				Light: pref.UserBarBackgroundLight,
				Dark:  pref.UserBarBackgroundDark,
			}),
		ErrorTextStyle: BaseStyle.
			Padding(0, 1).
			Foreground(lipgloss.AdaptiveColor{
				Light: pref.MessageTextLight,
				Dark:  pref.MessageTextDark,
			}).
			Background(lipgloss.AdaptiveColor{
				Light: pref.ErrorBackgroundLight,
				Dark:  pref.ErrorBackgroundDark,
			}),
		StatusBarStyle: BaseStyle.
			Foreground(lipgloss.AdaptiveColor{
				Light: pref.MessageTextDark,
				Dark:  pref.MessageTextLight,
			}).
			Background(lipgloss.AdaptiveColor{
				Light: pref.StatusBarBackgroundLight,
				Dark:  pref.StatusBarBackgroundDark,
			}),
		NormalAreaStyle: BaseStyle.
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.AdaptiveColor{
				Light: pref.AreaNormalLight,
				Dark:  pref.AreaNormalDark,
			}),
		HighlightedAreaStyle: BaseStyle.
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.AdaptiveColor{
				Light: pref.AreaHighlightLight,
				Dark:  pref.AreaHighlightDark,
			}),
		CategoryListTextStyle: BaseStyle.
			Foreground(lipgloss.AdaptiveColor{
				Light: pref.CategoryListTextLight,
				Dark:  pref.CategoryListTextDark,
			}),
		HighlightedMessageBorderStyle: BaseStyle.
			BorderStyle(lipgloss.NormalBorder()).
			BorderBackground(lipgloss.AdaptiveColor{
				Light: pref.HighLightedMessageBorderLight,
				Dark:  pref.HighLightedMessageBorderDark,
			}).
			BorderForeground(lipgloss.AdaptiveColor{
				Light: pref.HighLightedMessageBorderLight,
				Dark:  pref.HighLightedMessageBorderDark,
			}).
			PaddingLeft(1).
			BorderLeft(true).
			BorderRight(false).
			BorderTop(false).
			BorderBottom(false).
			Foreground(lipgloss.AdaptiveColor{
				Light: pref.MessageTextLight,
				Dark:  pref.MessageTextDark,
			}),
		RoomListTextStyle: BaseStyle.
			Foreground(lipgloss.AdaptiveColor{
				Light: pref.RoomListTextLight,
				Dark:  pref.RoomListTextDark,
			}),
		HighlightedRoomTextStyle: BaseStyle.
			Foreground(lipgloss.AdaptiveColor{
				Light: pref.HighlightedRoomTextLight,
				Dark:  pref.HighlightedRoomTextDark,
			}).
			Background(lipgloss.AdaptiveColor{
				Light: pref.HighlightedRoomBackgroundLight,
				Dark:  pref.HighlightedRoomBackgroundDark,
			}),
		MessageAuthorStyle: BaseStyle.
			Foreground(lipgloss.AdaptiveColor{
				Light: pref.MessageAuthorLight,
				Dark:  pref.MessageAuthorDark,
			}),
		MessageTextStyle: BaseStyle.
			Foreground(lipgloss.AdaptiveColor{
				Light: pref.MessageTextLight,
				Dark:  pref.MessageTextDark,
			}),
	}
}
