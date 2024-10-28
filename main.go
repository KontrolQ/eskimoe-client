package main

/*
This client will run locally and connect to multiple remote servers.
Auth is device-based, so the client will generate a unique token for each device.
This token will be used to join a server - servers will store the token and the user's ID.
This server will be a chat server like Discord, but every server will be self-hosted separately.
So the user can join multiple servers and chat with different people on each server.
Servers will joined by their URLs, which will return JSON responses. Eskimoe Client will only
Connect to Eskimoe servers, which will return responses in a known format.
*/

import (
	"eskimoe-client/database"
	"eskimoe-client/screens"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	databasePath := "eskimoe.db"

	// Read database poth from args
	if len(os.Args) > 1 {
		fmt.Println("Using database path from command line:", os.Args[1])
		databasePath = os.Args[1]
	}

	database.Initialize(databasePath)

	p := tea.NewProgram(screens.Initialize(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting program:", err)
		os.Exit(1)
	}
}
