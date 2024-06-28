package shared

import (
	catppuccin "github.com/catppuccin/go"
)

var DefaultPreferences = Preferences{
	AreaNormal: ColorPreference{
		Light: catppuccin.Mocha.Base().Hex,
		Dark:  catppuccin.Latte.Base().Hex,
	},
	AreaHighlight: ColorPreference{
		Light: catppuccin.Latte.Teal().Hex,
		Dark:  catppuccin.Mocha.Teal().Hex,
	},
	CategoryListText: ColorPreference{
		Light: catppuccin.Mocha.Base().Hex,
		Dark:  catppuccin.Mocha.Rosewater().Hex,
	},
	HighLightedMessageBorder: ColorPreference{
		Light: catppuccin.Latte.Mauve().Hex,
		Dark:  catppuccin.Mocha.Mauve().Hex,
	},
	HighlightedRoomText: ColorPreference{
		Light: catppuccin.Mocha.Base().Hex,
		Dark:  catppuccin.Latte.Base().Hex,
	},
	HighlightedRoomBackground: ColorPreference{
		Light: catppuccin.Mocha.Lavender().Hex,
		Dark:  catppuccin.Latte.Lavender().Hex,
	},
	MessageAuthor: ColorPreference{
		Light: catppuccin.Mocha.Overlay2().Hex,
		Dark:  catppuccin.Latte.Overlay2().Hex,
	},
	MessageText: ColorPreference{
		Light: catppuccin.Mocha.Crust().Hex,
		Dark:  catppuccin.Latte.Crust().Hex,
	},
	RoomListText: ColorPreference{
		Light: catppuccin.Latte.Peach().Hex,
		Dark:  catppuccin.Mocha.Yellow().Hex,
	},
	RoomStatusText: ColorPreference{
		Light: catppuccin.Latte.Base().Hex,
		Dark:  catppuccin.Mocha.Base().Hex,
	},
	RoomStatusBackground: ColorPreference{
		Light: catppuccin.Latte.Teal().Hex,
		Dark:  catppuccin.Mocha.Teal().Hex,
	},
	ServerStatusText: ColorPreference{
		Light: catppuccin.Latte.Base().Hex,
		Dark:  catppuccin.Mocha.Base().Hex,
	},
	ServerStatusBackground: ColorPreference{
		Light: catppuccin.Latte.Green().Hex,
		Dark:  catppuccin.Mocha.Green().Hex,
	},
	UserBarText: ColorPreference{
		Light: catppuccin.Latte.Base().Hex,
		Dark:  catppuccin.Mocha.Base().Hex,
	},
	UserBarBackground: ColorPreference{
		Light: catppuccin.Latte.Yellow().Hex,
		Dark:  catppuccin.Mocha.Yellow().Hex,
	},
	SidebarWidth:       30,
	MessageInputHeight: 10,
}
