package database

// Ideally only one user will be there on a device, but the user can switch accounts on the same client.
// So the client will store multiple user tokens and the user can switch between them.
type User struct {
	ID              uint     `gorm:"primaryKey;autoIncrement:true" json:"-"`
	UniqueID        string   `gorm:"unique;not null" json:"unique_id"`
	UniqueToken     string   `gorm:"unique;not null" json:"unique_token"`
	Servers         []Server `gorm:"many2many:user_servers" json:"-"`
	DisplayName     string   `json:"display_name"`
	Current         bool     `gorm:"default:false" json:"-"`
	CurrentServerID uint     `json:"-"`
	CurrentServer   Server   `gorm:"foreignKey:CurrentServerID" json:"-"`
}

// Servers are just a collection of URLs. All user based settings are stored on the server's database.
// We will only store the URL and the user's ID here to know what servers the user is a part of.
type Server struct {
	ID    uint   `gorm:"primaryKey;autoIncrement:true"`
	URL   string `gorm:"unique;not null"`
	Users []User `gorm:"many2many:user_servers;"`
}
