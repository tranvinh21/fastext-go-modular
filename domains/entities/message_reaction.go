package entity

import "time"

type MessageReaction struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	MessageID uint      `json:"message_id"`
	UserID    uint      `json:"user_id"`
	Reaction  string    `json:"reaction"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Message Message `gorm:"foreignKey:MessageID,references:ID,onDelete:CASCADE"`
	User    User    `gorm:"foreignKey:UserID,references:ID,onDelete:CASCADE"`
}

func (MessageReaction) TableName() string {
	return "message_reactions"
}
