package entity

import "time"

type MessageMedia struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	MessageID uint      `json:"message_id"`
	MediaURL  string    `json:"media_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Message Message `gorm:"foreignKey:MessageID,references:ID,onDelete:CASCADE"`
}

func (MessageMedia) TableName() string {
	return "message_media"
}
