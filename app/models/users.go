package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// USER_TABLE follows Rails' default table naming ("users", plural, snake_case).
// USER_COLLECTION is kept as an alias so older code/seeds still compile while
// the project moves from Mongo to PostgreSQL.
const USER_TABLE = "users"

const USER_COLLECTION = USER_TABLE

// User mirrors the users table. JSON tags stay the same as the old Mongo model
// so API responses are backwards compatible; bson tags are kept so the model
// can still be reused by the legacy Mongo package if needed.
type User struct {
	ID        string    `gorm:"column:id;type:uuid;primaryKey" json:"id,omitempty" bson:"_id,omitempty"`
	Username  string    `gorm:"column:username;size:255" json:"username,omitempty" bson:"username,omitempty"`
	Email     string    `gorm:"column:email;size:255;uniqueIndex;not null" json:"email,omitempty" bson:"email,omitempty"`
	Firstname string    `gorm:"column:firstname;size:255" json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string    `gorm:"column:lastname;size:255" json:"lastname,omitempty" bson:"lastname,omitempty"`
	Authtoken string    `gorm:"column:auth_token;size:255" json:"auth_token,omitempty" bson:"authtoken,omitempty"`
	Password  string    `gorm:"column:password;size:255" json:"password,omitempty" bson:"password,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// TableName pins the table name instead of relying on GORM's pluralizer.
func (User) TableName() string { return USER_TABLE }

// BeforeCreate assigns a UUID primary key when the caller does not provide one
// (the equivalent of a Postgres gen_random_uuid() column default, without
// requiring the pgcrypto extension).
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		u.ID = uuid.NewString()
	}
	return nil
}
