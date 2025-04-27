package entity

import "time"

type Friend struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID      uint      `json:"user_id"`
	FriendID    uint      `json:"friend_id"`
	InitiatorID uint      `json:"initiator_id"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	User      User `gorm:"foreignKey:UserID,references:ID,onDelete:CASCADE"`
	Friend    User `gorm:"foreignKey:FriendID,references:ID,onDelete:CASCADE"`
	Initiator User `gorm:"foreignKey:InitiatorID,references:ID,onDelete:CASCADE"`
}

func (Friend) TableName() string {
	return "friends"
}
