package lib

import "github.com/charmbracelet/bubbles/key"

var QuitConfirmKeys = QuitConfirmKeyMap{
	Quit:    key.NewBinding(key.WithKeys("ctrl+c"), key.WithHelp("^C", "Quit")),
	Confirm: key.NewBinding(key.WithKeys("enter"), key.WithHelp("⏎", "Confirm")),
}

var QuitConfirmCancelKeys = QuitConfirmCancelKeyMap{
	Quit:    key.NewBinding(key.WithKeys("ctrl+c"), key.WithHelp("^C", "Quit")),
	Confirm: key.NewBinding(key.WithKeys("enter"), key.WithHelp("⏎", "Confirm")),
	Cancel:  key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "Cancel")),
}
