package entity

import (
	"time"
)

type MessageStatus struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	MessageID uint      `json:"message_id"`
	UserID    uint      `json:"user_id"`
	Status    string    `json:"status"`
	SeenAt    time.Time `json:"seen_at"`
	CreatedAt time.Time `json:"created_at"`

	Message Message `gorm:"foreignKey:MessageID,references:ID,onDelete:CASCADE"`
	User    User    `gorm:"foreignKey:UserID,references:ID,onDelete:CASCADE"`
}

func (MessageStatus) TableName() string {
	return "message_statuses"
}
