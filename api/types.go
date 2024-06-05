package api

// Request Types
type JoinMemberRequest struct {
	UniqueID    string `json:"unique_id"`
	UniqueToken string `json:"unique_token"`
	DisplayName string `json:"display_name"`
}

// Response Types
type JoinMemberSuccessResponse struct {
	Message string `json:"message"`
	Member  Member `json:"member"`
}

type JoinMemberErrorResponse struct {
	ErrorCode int    `json:"error_code"`
	Error     string `json:"error"`
}

// Server Types
type Permission string

const (
	ViewRoom           Permission = "view_room"
	SendMessage        Permission = "send_message"
	AddLink            Permission = "add_link"
	AddFile            Permission = "add_file"
	AddReaction        Permission = "add_reaction"
	CreatePoll         Permission = "create_poll"
	DeleteMessage      Permission = "delete_message"
	ManageRoles        Permission = "manage_roles"
	ChangeName         Permission = "change_name"
	MuteMembers        Permission = "mute_members"
	KickMembers        Permission = "kick_members"
	BanMembers         Permission = "ban_members"
	ManageRooms        Permission = "manage_rooms"
	RunCommands        Permission = "run_commands"
	ViewLogs           Permission = "view_logs"
	ViewMessageHistory Permission = "view_message_history"
	CreateEvents       Permission = "create_events"
	ManageEvents       Permission = "manage_events"
	Administrator      Permission = "administrator"
)

// Room Types: Announcement, Text, Commands, Archive
type RoomType string

const (
	Announcement RoomType = "announcement"
	Text         RoomType = "text"
	Commands     RoomType = "commands"
	Archive      RoomType = "archive"
)

type Member struct {
	UniqueID    string `gorm:"unique" json:"unique_id"`
	DisplayName string `json:"display_name"`
	About       string `json:"about"`
	JoinedAt    string `json:"joined_at"`
	Pronouns    string `json:"pronouns"`
	Roles       []Role `json:"roles"`
}

type Role struct {
	Name        string             `gorm:"unique" json:"name"`
	Permissions []ServerPermission `json:"permissions"`
}

type ServerPermission struct {
	Role       Role       `json:"role"`
	Permission Permission `json:"permission"`
}
