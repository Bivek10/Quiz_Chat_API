package models

import (
	"time"

	"gorm.io/gorm"
)

type FriendRequest struct {
	Sender    int64          `json:"sender"`
	Receiver  int64          `json:"receiver"`
	Status    string         `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"` //add soft delete in gorm
}

func (f FriendRequest) TableName() string {
	return "friendrequest"
}

func (f FriendRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{}
}
