package models

type ChatMember struct {
	Base
	UserID int   `json:"user_id"`
	RoomID int64 `json:"room_id"`
}

func (c ChatMember) TableName() string {
	return "chatmember"
}
