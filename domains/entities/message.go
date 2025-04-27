package entity

import "time"

type Message struct {
	ID               uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	ConversationID   uint      `json:"conversation_id"`
	UserID           uint      `json:"user_id"`
	ReplyToMessageID uint      `json:"reply_to_message_id"`
	Parts            string    `json:"parts"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	DeletedAt        time.Time `json:"deleted_at"`

	Conversation Conversation `gorm:"foreignKey:ConversationID,references:ID,onDelete:CASCADE"`
	User         User         `gorm:"foreignKey:UserID,references:ID,onDelete:CASCADE"`
}

func (Message) TableName() string {
	return "messages"
}
