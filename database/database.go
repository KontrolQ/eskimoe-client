package database

import (
	"eskimoe-client/api"
	"eskimoe-client/lib"
	"eskimoe-client/shared"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *gorm.DB

func makePrefs() UserPreferences {
	return UserPreferences{
		AreaNormalLight:                shared.DefaultPreferences.AreaNormal.Light,
		AreaNormalDark:                 shared.DefaultPreferences.AreaNormal.Dark,
		AreaHighlightLight:             shared.DefaultPreferences.AreaHighlight.Light,
		AreaHighlightDark:              shared.DefaultPreferences.AreaHighlight.Dark,
		CategoryListTextLight:          shared.DefaultPreferences.CategoryListText.Light,
		CategoryListTextDark:           shared.DefaultPreferences.CategoryListText.Dark,
		HighLightedMessageBorderLight:  shared.DefaultPreferences.HighLightedMessageBorder.Light,
		HighLightedMessageBorderDark:   shared.DefaultPreferences.HighLightedMessageBorder.Dark,
		HighlightedRoomTextLight:       shared.DefaultPreferences.HighlightedRoomText.Light,
		HighlightedRoomTextDark:        shared.DefaultPreferences.HighlightedRoomText.Dark,
		HighlightedRoomBackgroundLight: shared.DefaultPreferences.HighlightedRoomBackground.Light,
		HighlightedRoomBackgroundDark:  shared.DefaultPreferences.HighlightedRoomBackground.Dark,
		MessageAuthorLight:             shared.DefaultPreferences.MessageAuthor.Light,
		MessageAuthorDark:              shared.DefaultPreferences.MessageAuthor.Dark,
		MessageTextLight:               shared.DefaultPreferences.MessageText.Light,
		MessageTextDark:                shared.DefaultPreferences.MessageText.Dark,
		RoomListTextLight:              shared.DefaultPreferences.RoomListText.Light,
		RoomListTextDark:               shared.DefaultPreferences.RoomListText.Dark,
		RoomStatusTextLight:            shared.DefaultPreferences.RoomStatusText.Light,
		RoomStatusTextDark:             shared.DefaultPreferences.RoomStatusText.Dark,
		RoomStatusBackgroundLight:      shared.DefaultPreferences.RoomStatusBackground.Light,
		RoomStatusBackgroundDark:       shared.DefaultPreferences.RoomStatusBackground.Dark,
		ServerStatusTextLight:          shared.DefaultPreferences.ServerStatusText.Light,
		ServerStatusTextDark:           shared.DefaultPreferences.ServerStatusText.Dark,
		ServerStatusBackgroundLight:    shared.DefaultPreferences.ServerStatusBackground.Light,
		ServerStatusBackgroundDark:     shared.DefaultPreferences.ServerStatusBackground.Dark,
		UserBarTextLight:               shared.DefaultPreferences.UserBarText.Light,
		UserBarTextDark:                shared.DefaultPreferences.UserBarText.Dark,
		UserBarBackgroundLight:         shared.DefaultPreferences.UserBarBackground.Light,
		UserBarBackgroundDark:          shared.DefaultPreferences.UserBarBackground.Dark,
		SidebarWidth:                   shared.DefaultPreferences.SidebarWidth,
		MessageInputHeight:             shared.DefaultPreferences.MessageInputHeight,
	}
}

/*
User UniqueID is a UUID generated by the client.
User UniqueToken is a new random 256-bit token generated by the client and hashed with device's MAC address.
*/
func newUser() *User {
	uniqueID := lib.GenerateUUID()
	uniqueToken := lib.GenerateToken()

	return &User{
		UniqueID:    uniqueID,
		UniqueToken: uniqueToken,
		Current:     true,
		Preferences: makePrefs(),
	}
}

func GetCurrentUser() User {
	var user User

	Database.Where("current = ?", true).First(&user)

	return user
}

func GetServers(user User) []Server {
	var servers []Server

	Database.Model(&user).Association("Servers").Find(&servers)

	return servers
}

func JoinServer(user User, server Server) (User, error) {
	serverURL := server.URL
	member := api.JoinMemberRequest{
		UniqueID:    user.UniqueID,
		UniqueToken: user.UniqueToken,
		DisplayName: user.DisplayName,
	}

	// Attempt to join the server via an API call
	resp, err := api.JoinServerAsMember(serverURL, member)
	if err != nil {
		return user, err
	}

	// Check if the server is already in the database
	var existingServer Server
	if err := Database.Where("url = ?", server.URL).First(&existingServer).Error; err != nil {
		// If not, save it
		if err := Database.Create(&server).Error; err != nil {
			return user, err
		}
	} else {
		// Use the existing server
		server = existingServer
	}

	// Update the user's auth token with the one from the server
	user.AuthToken = resp.Member.AuthToken

	// Set the current server to the newly joined server
	user.CurrentServer = server
	user.CurrentServerID = server.ID

	// Mark the user as the current user
	user.Current = true

	// Save the updated user record in the database
	if err := Database.Save(&user).Error; err != nil {
		return user, err
	}

	// Add the server to the user's list of servers
	if err := Database.Model(&user).Association("Servers").Append(&server); err != nil {
		return user, err
	}

	return user, nil
}

func Initialize() {
	var err error
	Database, err = gorm.Open(sqlite.Open("eskimoe.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	Database.AutoMigrate(&User{}, &Server{}, &UserPreferences{})

	// Create a user if there isn't one
	currentUser := GetCurrentUser()

	if currentUser.ID == 0 {
		user := *newUser()
		Database.Create(&user)
	}
}
