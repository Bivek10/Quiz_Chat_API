package models

type ChatMember struct {
	Base
	UserID int `json:"user_id"`
	RoomID int `json:"room_id"`
}

func (c ChatMember) TableName() string {
	return "chatmember"
}
