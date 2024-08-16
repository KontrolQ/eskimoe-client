package database

// Ideally only one user will be there on a device, but the user can switch accounts on the same client.
// So the client will store multiple user tokens and the user can switch between them.
type User struct {
	ID              uint            `gorm:"primaryKey;autoIncrement:true" json:"-"`
	UniqueID        string          `gorm:"unique;not null" json:"unique_id"`
	UniqueToken     string          `gorm:"unique;not null" json:"unique_token"`
	AuthToken       string          `json:"auth_token"`
	Servers         []Server        `gorm:"many2many:user_servers" json:"-"`
	DisplayName     string          `json:"display_name"`
	Current         bool            `gorm:"default:false" json:"-"`
	CurrentServer   Server          `gorm:"foreignKey:CurrentServerID" json:"-"`
	CurrentServerID uint            `json:"-"`
	Preferences     UserPreferences `gorm:"embedded" json:"preferences"`
}

type Server struct {
	ID    uint   `gorm:"primaryKey;autoIncrement:true"`
	URL   string `gorm:"unique;not null"`
	Users []User `gorm:"many2many:user_servers;"`
}

type UserPreferences struct {
	ID                             uint   `gorm:"primaryKey;autoIncrement:true"`
	AreaNormalLight                string `json:"area_normal_light"`
	AreaNormalDark                 string `json:"area_normal_dark"`
	AreaHighlightLight             string `json:"area_highlight_light"`
	AreaHighlightDark              string `json:"area_highlight_dark"`
	CategoryListTextLight          string `json:"category_list_text_light"`
	CategoryListTextDark           string `json:"category_list_text_dark"`
	HighLightedMessageBorderLight  string `json:"highlighted_message_border_light"`
	HighLightedMessageBorderDark   string `json:"highlighted_message_border_dark"`
	HighlightedRoomTextLight       string `json:"highlighted_room_text_light"`
	HighlightedRoomTextDark        string `json:"highlighted_room_text_dark"`
	HighlightedRoomBackgroundLight string `json:"highlighted_room_background_light"`
	HighlightedRoomBackgroundDark  string `json:"highlighted_room_background_dark"`
	MessageAuthorLight             string `json:"message_author_light"`
	MessageAuthorDark              string `json:"message_author_dark"`
	MessageTextLight               string `json:"message_text_light"`
	MessageTextDark                string `json:"message_text_dark"`
	RoomListTextLight              string `json:"room_list_text_light"`
	RoomListTextDark               string `json:"room_list_text_dark"`
	RoomStatusTextLight            string `json:"room_status_text_light"`
	RoomStatusTextDark             string `json:"room_status_text_dark"`
	RoomStatusBackgroundLight      string `json:"room_status_background_light"`
	RoomStatusBackgroundDark       string `json:"room_status_background_dark"`
	ServerStatusTextLight          string `json:"server_status_text_light"`
	ServerStatusTextDark           string `json:"server_status_text_dark"`
	ServerStatusBackgroundLight    string `json:"server_status_background_light"`
	ServerStatusBackgroundDark     string `json:"server_status_background_dark"`
	UserBarTextLight               string `json:"user_bar_text_light"`
	UserBarTextDark                string `json:"user_bar_text_dark"`
	UserBarBackgroundLight         string `json:"user_bar_background_light"`
	UserBarBackgroundDark          string `json:"user_bar_background_dark"`
	ErrorBackgroundLight           string `json:"error_background_light"`
	ErrorBackgroundDark            string `json:"error_background_dark"`
	StatusBarBackgroundLight       string `json:"status_bar_background_light"`
	StatusBarBackgroundDark        string `json:"status_bar_background_dark"`
	SidebarWidth                   int    `json:"sidebar_width"`
	MessageInputHeight             int    `json:"message_input_height"`
}
