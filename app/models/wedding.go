package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// WEDDING_TABLE follows Rails' naming convention. Nested documents that used
// to be embedded in the Mongo document are stored as jsonb columns, using
// GORM's built-in "serializer:json" tag.
const WEDDING_TABLE = "wedding_data"

const WEDDING_COLLECTION = WEDDING_TABLE

type SocialMedia struct {
	Username string `json:"username"`
	Link     string `json:"link"`
	Platform string `json:"platform"`
}

type Parent struct {
	Father struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	} `json:"father"`
	Mother struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	} `json:"mother"`
}

type AkadResepsi struct {
	Date    string `json:"date"`
	Start   string `json:"start"`
	End     string `json:"end"`
	Address string `json:"address"`
	MapURL  string `json:"mapUrl"`
}

type LoveStory struct {
	Date        string `json:"date"`
	Description string `json:"description"`
	ImageURL    string `json:"imageUrl"`
}

type Envelop struct {
	BackgroundImage string `json:"backgroundImage"`
	BackgroundColor string `json:"backgroundColor"`
}

type Guest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Streaming struct {
	Platform   string `json:"platform"`
	StreamLink string `json:"streamLink"`
}

type Gift struct {
	Name     string `json:"name"`
	Account  string `json:"account"`
	BankName string `json:"bankname"`
	Link     string `json:"link"`
}

type Wedding struct {
	Date    string      `json:"date"`
	Akad    AkadResepsi `json:"akad"`
	Resepsi AkadResepsi `json:"resepsi"`
}

type BrideGroom struct {
	Firstname   string        `json:"firstname"`
	Lastname    string        `json:"lastname"`
	ProfileURL  string        `json:"profileUrl"`
	SocialMedia []SocialMedia `json:"socialMedia"`
	Parent      Parent        `json:"parent"`
}

// WeddingData is the PostgreSQL version of the old Mongo document. Columns are
// snake_case like Rails migrations, embedded payloads are jsonb.
//
// UserID is a pointer so a missing/blank user_id is stored as NULL instead of
// an empty string, which PostgreSQL would reject for a uuid column.
type WeddingData struct {
	ID        string       `gorm:"column:id;type:uuid;primaryKey" json:"id,omitempty"`
	UserID    *string      `gorm:"column:user_id;type:uuid;index" json:"user_id,omitempty"`
	Bride     BrideGroom   `gorm:"column:bride;type:jsonb;serializer:json" json:"bride"`
	Groom     BrideGroom   `gorm:"column:groom;type:jsonb;serializer:json" json:"groom"`
	Wedding   Wedding      `gorm:"column:wedding;type:jsonb;serializer:json" json:"wedding"`
	LoveStory []LoveStory  `gorm:"column:love_story;type:jsonb;serializer:json" json:"loveStory"`
	Envelop   Envelop      `gorm:"column:envelop;type:jsonb;serializer:json" json:"envelop"`
	Gallery   []string     `gorm:"column:gallery;type:jsonb;serializer:json" json:"gallery"`
	Guest     Guest        `gorm:"column:guest;type:jsonb;serializer:json" json:"guest"`
	Streaming []Streaming  `gorm:"column:streaming;type:jsonb;serializer:json" json:"streaming"`
	Gifts     []Gift       `gorm:"column:gifts;type:jsonb;serializer:json" json:"gifts"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at,omitempty"`
}

// TableName pins the table name instead of relying on GORM's pluralizer.
func (WeddingData) TableName() string { return WEDDING_TABLE }

// BeforeCreate assigns a UUID primary key when the caller does not provide one
// and drops a user_id that is not a valid UUID.
func (w *WeddingData) BeforeCreate(tx *gorm.DB) error {
	if w.ID == "" {
		w.ID = uuid.NewString()
	}
	if w.UserID != nil {
		if _, err := uuid.Parse(*w.UserID); err != nil {
			w.UserID = nil
		}
	}
	return nil
}
