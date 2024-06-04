package lib

import "github.com/charmbracelet/bubbles/key"

type QuitConfirmKeyMap struct {
	Quit    key.Binding
	Confirm key.Binding
}

// FullHelp implements help.KeyMap.
func (k QuitConfirmKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Quit, k.Confirm},
	}
}

// ShortHelp implements help.KeyMap.
func (k QuitConfirmKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Confirm}
}

type QuitConfirmCancelKeyMap struct {
	Quit    key.Binding
	Confirm key.Binding
	Cancel  key.Binding
}

// FullHelp implements help.KeyMap.
func (k QuitConfirmCancelKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Quit, k.Confirm, k.Cancel},
	}
}

// ShortHelp implements help.KeyMap.
func (k QuitConfirmCancelKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Confirm, k.Cancel}
}
