package entity

import "time"

type Conversation struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	ChatKey   string    `json:"chat_key"`
	IsGroup   bool      `json:"is_group"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (Conversation) TableName() string {
	return "conversations"
}
