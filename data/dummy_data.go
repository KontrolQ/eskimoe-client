package data

import (
	"time"

	"github.com/charmbracelet/lipgloss"

	"eskimoe-client/shared"

	catppuccin "github.com/catppuccin/go"
)

var (
	GeneralRoomMessages = []shared.Message{
		{
			Author:  "John Doe",
			Content: "Hello, World!",
			SentAt:  time.Now().Add(-time.Minute * 10),
			Reactions: []shared.Reactions{
				{Reaction: "LIKE", Count: 1, Users: []string{"Jane Doe"}, Color: lipgloss.AdaptiveColor{Light: catppuccin.Mocha.Sapphire().Hex, Dark: catppuccin.Mocha.Sapphire().Hex}},
				{Reaction: "LOVE", Count: 2, Users: []string{"Conrad Reeves", "Jane Doe"}, Color: lipgloss.AdaptiveColor{Light: catppuccin.Latte.Maroon().Hex, Dark: catppuccin.Latte.Maroon().Hex}},
			},
		},
		{
			Author:    "Jane Doe",
			Content:   "Hi, John!",
			SentAt:    time.Now().Add(-time.Minute * 9),
			Reactions: []shared.Reactions{},
		},
		{
			Author:    "Conrad Reeves",
			Content:   "Hey, Guys!",
			SentAt:    time.Now().Add(-time.Minute * 8),
			Reactions: []shared.Reactions{},
		},
		{
			Author:    "John Doe",
			Content:   "How are you doing?",
			SentAt:    time.Now().Add(-time.Minute * 7),
			Reactions: []shared.Reactions{},
		},
		{
			Author:    "Jane Doe",
			Content:   "I'm good, thanks!",
			SentAt:    time.Now().Add(-time.Minute * 6),
			Reactions: []shared.Reactions{},
		},
		{
			Author:    "Conrad Reeves",
			Content:   "I'm doing great! Thanks for asking! What are you guys up to?",
			SentAt:    time.Now().Add(-time.Minute * 5),
			Reactions: []shared.Reactions{},
		},
		{
			Author:    "John Doe",
			Content:   "Just working on a project. It's going well! How about you?",
			SentAt:    time.Now().Add(-time.Minute * 4),
			Reactions: []shared.Reactions{},
		},
		{
			Author:  "Conrad Reeves",
			Content: "I'm working on a project too! It's going great! I am currently building a discord like chat client called Eskimoe! It is built using Golang and Bubbletea! This client will run locally and connect to multiple remote servers. Auth is device-based, so the client will generate a unique token for each device. This token will be used to join a server - servers will store the token and the user's ID. This server will be a chat server like Discord, but every server will be self-hosted separately. So the user can join multiple servers and chat with different people on each server. Servers will joined by their URLs, which will return JSON responses. Eskimoe Client will only connect to Eskimoe servers, which will return responses in a known format. I'm really excited about it!",
			SentAt:  time.Now().Add(-time.Minute * 3),
			Reactions: []shared.Reactions{
				{Reaction: "LIKE", Count: 2, Users: []string{"Jane Doe", "John Doe"}, Color: lipgloss.AdaptiveColor{Light: catppuccin.Mocha.Sapphire().Hex, Dark: catppuccin.Mocha.Sapphire().Hex}},
				{Reaction: "LOVE", Count: 1, Users: []string{"Jane Doe"}, Color: lipgloss.AdaptiveColor{Light: catppuccin.Latte.Maroon().Hex, Dark: catppuccin.Latte.Maroon().Hex}},
				{Reaction: "WOW", Count: 1, Users: []string{"John Doe"}, Color: lipgloss.AdaptiveColor{Light: catppuccin.Frappe.Yellow().Hex, Dark: catppuccin.Frappe.Yellow().Hex}},
				{Reaction: "YAY", Count: 1, Users: []string{"Jane Doe"}, Color: lipgloss.AdaptiveColor{Light: catppuccin.Frappe.Peach().Hex, Dark: catppuccin.Frappe.Peach().Hex}},
			},
		},
		{
			Author:    "Jane Doe",
			Content:   "That's awesome! I'd love to try it out when it's ready!",
			SentAt:    time.Now().Add(-time.Minute * 2),
			Reactions: []shared.Reactions{},
		},
		{
			Author:    "John Doe",
			Content:   "Me too!",
			SentAt:    time.Now().Add(-time.Minute * 2),
			Reactions: []shared.Reactions{},
		},
		{
			Author:    "Conrad Reeves",
			Content:   "Thanks, guys! I'm hoping to have a beta version ready by next month. Any suggestions or features you'd like to see?",
			SentAt:    time.Now().Add(-time.Minute * 1),
			Reactions: []shared.Reactions{},
		},
		{
			Author:  "Jane Doe",
			Content: "Maybe some cool themes and customizable avatars!",
			SentAt:  time.Now().Add(-time.Minute * 1),
			Reactions: []shared.Reactions{
				{Reaction: "LIKE", Count: 1, Users: []string{"John Doe"}, Color: lipgloss.AdaptiveColor{Light: catppuccin.Mocha.Sapphire().Hex, Dark: catppuccin.Mocha.Sapphire().Hex}},
			},
		},
		{
			Author:    "John Doe",
			Content:   "Yeah, themes would be great. Also, maybe some sort of channel categorization?",
			SentAt:    time.Now(),
			Reactions: []shared.Reactions{},
		},
		{
			Author:  "Conrad Reeves",
			Content: "Awesome ideas! I'll definitely add those to my to-do list. Thanks!",
			SentAt:  time.Now(),
			Reactions: []shared.Reactions{
				{Reaction: "LIKE", Count: 2, Users: []string{"Jane Doe", "John Doe"}, Color: lipgloss.AdaptiveColor{Light: catppuccin.Mocha.Sapphire().Hex, Dark: catppuccin.Mocha.Sapphire().Hex}},
			},
		},
		{
			Author:    "Jane Doe",
			Content:   "Conrad, do you need any help with testing? I could spare some time over the weekend.",
			SentAt:    time.Now().Add(time.Minute * 1),
			Reactions: []shared.Reactions{},
		},
		{
			Author:    "John Doe",
			Content:   "I can help too. I have some experience with Golang, so maybe I can assist with coding or debugging.",
			SentAt:    time.Now().Add(time.Minute * 2),
			Reactions: []shared.Reactions{},
		},
		{
			Author:  "Conrad Reeves",
			Content: "That would be fantastic! I'll set up a GitHub repo and share the link with you both. We can start with some basic tests and go from there.",
			SentAt:  time.Now().Add(time.Minute * 3),
			Reactions: []shared.Reactions{
				{Reaction: "LIKE", Count: 2, Users: []string{"Jane Doe", "John Doe"}, Color: lipgloss.AdaptiveColor{Light: catppuccin.Mocha.Sapphire().Hex, Dark: catppuccin.Mocha.Sapphire().Hex}},
			},
		},
		{
			Author:  "Jane Doe",
			Content: "Great! Looking forward to it.",
			SentAt:  time.Now().Add(time.Minute * 4),
			Reactions: []shared.Reactions{
				{Reaction: "LOVE", Count: 1, Users: []string{"John Doe"}, Color: lipgloss.AdaptiveColor{Light: catppuccin.Latte.Maroon().Hex, Dark: catppuccin.Latte.Maroon().Hex}},
			},
		},
		{
			Author:    "John Doe",
			Content:   "Same here. It's going to be fun!",
			SentAt:    time.Now().Add(time.Minute * 5),
			Reactions: []shared.Reactions{},
		},
		{
			Author:  "Conrad Reeves",
			Content: "I appreciate the support, guys. This project means a lot to me. Let's make Eskimoe something special!",
			SentAt:  time.Now().Add(time.Minute * 6),
			Reactions: []shared.Reactions{
				{Reaction: "LOVE", Count: 2, Users: []string{"Jane Doe", "John Doe"}, Color: lipgloss.AdaptiveColor{Light: catppuccin.Latte.Maroon().Hex, Dark: catppuccin.Latte.Maroon().Hex}},
			},
		},
		{
			Author:    "Jane Doe",
			Content:   "Absolutely! By the way, have you thought about incorporating video or voice chat in the future?",
			SentAt:    time.Now().Add(time.Minute * 7),
			Reactions: []shared.Reactions{},
		},
		{
			Author:    "Conrad Reeves",
			Content:   "Yes, that's definitely on the roadmap. It's a bit down the line, but I want to get the basics right first.",
			SentAt:    time.Now().Add(time.Minute * 8),
			Reactions: []shared.Reactions{},
		},
		{
			Author:  "John Doe",
			Content: "That makes sense. Starting with a solid foundation is key. Looking forward to seeing how Eskimoe evolves!",
			SentAt:  time.Now().Add(time.Minute * 9),
			Reactions: []shared.Reactions{
				{Reaction: "LIKE", Count: 1, Users: []string{"Jane Doe"}, Color: lipgloss.AdaptiveColor{Light: catppuccin.Mocha.Sapphire().Hex, Dark: catppuccin.Mocha.Sapphire().Hex}},
			},
		},
		{
			Author:    "Jane Doe",
			Content:   "Same here. Let us know if you need any more ideas or feedback.",
			SentAt:    time.Now().Add(time.Minute * 10),
			Reactions: []shared.Reactions{},
		},
		{
			Author:    "Conrad Reeves",
			Content:   "Will do! Thanks again for the support, you two. I'm excited for this journey!",
			SentAt:    time.Now().Add(time.Minute * 11),
			Reactions: []shared.Reactions{},
		},
		{
			Author:    "John Doe",
			Content:   "You're welcome, Conrad! Let's do this!",
			SentAt:    time.Now().Add(time.Minute * 12),
			Reactions: []shared.Reactions{},
		},
		{
			Author:  "Jane Doe",
			Content: "Go Team Eskimoe!",
			SentAt:  time.Now().Add(time.Minute * 13),
			Reactions: []shared.Reactions{
				{Reaction: "LOVE", Count: 1, Users: []string{"Conrad Reeves"}, Color: lipgloss.AdaptiveColor{Light: catppuccin.Latte.Maroon().Hex, Dark: catppuccin.Latte.Maroon().Hex}},
			},
		},
		{
			Author:  "Conrad Reeves",
			Content: "Team Eskimoe for the win!",
			SentAt:  time.Now().Add(time.Minute * 14),
			Reactions: []shared.Reactions{
				{Reaction: "LIKE", Count: 2, Users: []string{"Jane Doe", "John Doe"}, Color: lipgloss.AdaptiveColor{Light: catppuccin.Mocha.Sapphire().Hex, Dark: catppuccin.Mocha.Sapphire().Hex}},
			},
		},
	}

	AnnouncementsRoomMessages = []shared.Message{
		{
			Author:  "Admin",
			Content: "Welcome to Eskimoe!",
			SentAt:  time.Now().Add(-time.Minute * 5),
			Reactions: []shared.Reactions{
				{Reaction: "LIKE", Count: 1, Users: []string{"Jane Doe"}, Color: lipgloss.AdaptiveColor{Light: catppuccin.Mocha.Sapphire().Hex, Dark: catppuccin.Mocha.Sapphire().Hex}},
				{Reaction: "LOVE", Count: 2, Users: []string{"Conrad Reeves", "Jane Doe"}, Color: lipgloss.AdaptiveColor{Light: catppuccin.Latte.Maroon().Hex, Dark: catppuccin.Latte.Maroon().Hex}},
			},
		},
		{
			Author:    "Admin",
			Content:   "This is the announcements room!",
			SentAt:    time.Now().Add(-time.Minute * 4),
			Reactions: []shared.Reactions{},
		},
		{
			Author:    "Admin",
			Content:   "Please read the rules before posting!",
			SentAt:    time.Now().Add(-time.Minute * 3),
			Reactions: []shared.Reactions{},
		},
		{
			Author:    "Admin",
			Content:   "Have fun!",
			SentAt:    time.Now().Add(-time.Minute * 2),
			Reactions: []shared.Reactions{},
		},
		{
			Author:    "Admin",
			Content:   "Welcome to Eskimoe!",
			SentAt:    time.Now().Add(-time.Minute * 1),
			Reactions: []shared.Reactions{},
		},
	}

	IntroductionsRoomMessages = []shared.Message{
		{
			Author:  "John Doe",
			Content: "Hi, I'm John!",
			SentAt:  time.Now().Add(-time.Minute * 5),
			Reactions: []shared.Reactions{
				{Reaction: "LIKE", Count: 1, Users: []string{"Jane Doe"}, Color: lipgloss.AdaptiveColor{Light: catppuccin.Mocha.Sapphire().Hex, Dark: catppuccin.Mocha.Sapphire().Hex}},
			},
		},
		{
			Author:    "Jane Doe",
			Content:   "Hi, I'm Jane!",
			SentAt:    time.Now().Add(-time.Minute * 4),
			Reactions: []shared.Reactions{},
		},
		{
			Author:  "Admin",
			Content: "Welcome, John and Jane!",
			SentAt:  time.Now().Add(-time.Minute * 3),
			Reactions: []shared.Reactions{
				{Reaction: "LIKE", Count: 1, Users: []string{"John Doe"}, Color: lipgloss.AdaptiveColor{Light: catppuccin.Mocha.Sapphire().Hex, Dark: catppuccin.Mocha.Sapphire().Hex}},
				{Reaction: "LOVE", Count: 1, Users: []string{"Jane Doe"}, Color: lipgloss.AdaptiveColor{Light: catppuccin.Latte.Maroon().Hex, Dark: catppuccin.Latte.Maroon().Hex}},
			},
		},
	}

	PythonRoomMessages = []shared.Message{
		{
			Author:  "John Doe",
			Content: "I am a Python Developer and officially in love with Python!",
			SentAt:  time.Now().Add(-time.Minute * 5),
			Reactions: []shared.Reactions{
				{Reaction: "LIKE", Count: 1, Users: []string{"Jane Doe"}, Color: lipgloss.AdaptiveColor{Light: catppuccin.Mocha.Sapphire().Hex, Dark: catppuccin.Mocha.Sapphire().Hex}},
				{Reaction: "LOVE", Count: 2, Users: []string{"Conrad Reeves", "Jane Doe"}, Color: lipgloss.AdaptiveColor{Light: catppuccin.Latte.Maroon().Hex, Dark: catppuccin.Latte.Maroon().Hex}},
			},
		},
	}

	GoRoomMessages = []shared.Message{}

	JavascriptRoomMessages = []shared.Message{}

	UiUxRoomMessages = []shared.Message{}

	GraphicDesignRoomMessages = []shared.Message{}

	MemesRoomMessages = []shared.Message{}

	RandomRoomMessages = []shared.Message{}

	Categories = []shared.Category{
		{Name: "General", Rooms: []shared.Room{{Name: "General", RoomId: 1, Messages: GeneralRoomMessages}, {Name: "Announcements", RoomId: 2, Messages: AnnouncementsRoomMessages}, {Name: "Introductions", RoomId: 3, Messages: IntroductionsRoomMessages}}},
		{Name: "Development", Rooms: []shared.Room{{Name: "Python", RoomId: 4, Messages: PythonRoomMessages}, {Name: "Go", RoomId: 5, Messages: GoRoomMessages}, {Name: "JavaScript", RoomId: 6, Messages: JavascriptRoomMessages}}},
		{Name: "Design", Rooms: []shared.Room{{Name: "UI/UX", RoomId: 7, Messages: UiUxRoomMessages}, {Name: "Graphic Design", RoomId: 8, Messages: GraphicDesignRoomMessages}}},
		{Name: "Random", Rooms: []shared.Room{{Name: "Memes", RoomId: 9, Messages: MemesRoomMessages}, {Name: "Random", RoomId: 10, Messages: RandomRoomMessages}}},
	}
)

func GetRoomMessages(roomId int) []shared.Message {
	for _, category := range Categories {
		for _, room := range category.Rooms {
			if room.RoomId == roomId {
				return room.Messages
			}
		}
	}
	return []shared.Message{}
}

func GetRoomName(roomId int) string {
	for _, category := range Categories {
		for _, room := range category.Rooms {
			if room.RoomId == roomId {
				return room.Name
			}
		}
	}
	return ""
}
