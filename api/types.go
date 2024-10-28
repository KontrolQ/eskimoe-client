package api

// Socket Types
type BroadcastType int

const (
	MessageCreated BroadcastType = iota
	MessageDeleted
	MessageEdited
	MessageBulkDeleted
	MessageReactionCreated
	MessageReactionDeleted
	MessageReactionUpdated
	RoomCreated
	RoomDeleted
	RoomUpdated
	CategoryCreated
	CategoryDeleted
	CategoryUpdated
	CategoryOrderUpdated
	MemberJoined
	MemberLeft
	MemberBanned
	MemberKicked
	MemberUnbanned
	MemberUpdated
	RoleCreated
	RoleDeleted
	RoleUpdated
)

type SocketBroadcast struct {
	BroadcastType BroadcastType `json:"broadcast_type"`
	Data          interface{}   `json:"data"`
}

// Server Types
type Permission string

const (
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
	GenerateInvites    Permission = "generate_invites"
	Administrator      Permission = "administrator"
)

type RoomType string

const (
	Announcement RoomType = "announcement"
	Text         RoomType = "text"
	Commands     RoomType = "commands"
	Archive      RoomType = "archive"
)

type Member struct {
	UID         string `json:"uid"`
	DisplayName string `json:"display_name"`
	About       string `json:"about"`
	Pronouns    string `json:"pronouns"`
	Roles       []Role `json:"roles"`
	JoinedAt    string `json:"joined_at"`
}

type Role struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
	SystemRole  bool         `json:"system_role"`
	CreatedAt   string       `json:"created_at"`
}

type Category struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Rooms     []Room `json:"rooms"`
	RoomOrder []int  `json:"room_order"`
}

type Room struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Type        RoomType `json:"type"`
	CreatedAt   string   `json:"created_at"`
}

type Message struct {
	ID          int               `json:"id"`
	Content     string            `json:"content"`
	Author      Member            `json:"author"`
	Reactions   []MessageReaction `json:"reactions"`
	Attachments []Attachment      `json:"attachments"`
	RoomID      int               `json:"room_id"`
	Edited      bool              `json:"edited"`
	CreatedAt   string            `json:"created_at"`
}

type SendRoomMessage struct {
	Content string `json:"content"`
}

type Attachment struct {
	Type    string  `json:"type"`
	URL     string  `json:"url"`
	Message Message `json:"message"`
}

type MessageReaction struct {
	Reaction ServerReaction `json:"reaction"`
	Members  []Member       `json:"members"`
	Count    int            `json:"count"`
}

type ServerReaction struct {
	Reaction  string `json:"reaction"`
	Color     string `json:"color"`
	CreatedAt string `json:"created_at"`
}

type Event struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	StartTime   string   `json:"start_time"`
	EndTime     string   `json:"end_time"`
	CreatedBy   Member   `json:"created_by"`
	Interested  []Member `json:"interested"`
	CreatedAt   string   `json:"created_at"`
}

type ServerInfoUnauthorized struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Version string `json:"version"`
}

type ServerInfoAuthorized struct {
	Name            string           `json:"name"`
	Message         string           `json:"message"`
	PublicURL       string           `json:"public_url"`
	Mode            string           `json:"mode"`
	Categories      []Category       `json:"categories"`
	CategoryOrder   []int            `json:"category_order"`
	ServerReactions []ServerReaction `json:"server_reactions"`
	Roles           []Role           `json:"roles"`
	RoleOrder       []int            `json:"role_order"`
	Events          []Event          `json:"events"`
	Members         []Member         `json:"members"`
	CreatedAt       string           `json:"created_at"`
}

type JoinMemberRequest struct {
	UniqueID    string `json:"unique_id"`
	UniqueToken string `json:"unique_token"`
	DisplayName string `json:"display_name"`
}

type JoinMemberSuccessResponse struct {
	AuthToken string `json:"auth_token"`
	Message   string `json:"message"`
	Member    Member `json:"member"`
}

type JoinMemberErrorResponse struct {
	ErrorCode int    `json:"error_code"`
	Error     string `json:"error"`
}

type DeletedMessage struct {
	MessageID int  `json:"message_id"`
	RoomID    int  `json:"room_id"`
	Deleted   bool `json:"deleted"`
}
