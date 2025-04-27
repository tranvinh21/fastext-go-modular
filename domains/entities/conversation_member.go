package entity

import "time"

type ConversationMember struct {
	ID             uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	ConversationID uint      `json:"conversation_id"`
	UserID         uint      `json:"user_id"`
	Nickname       string    `json:"nickname"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	Conversation Conversation `gorm:"foreignKey:ConversationID,references:ID,onDelete:CASCADE"`
	User         User         `gorm:"foreignKey:UserID,references:ID,onDelete:CASCADE"`
}

func (ConversationMember) TableName() string {
	return "conversation_members"
}
