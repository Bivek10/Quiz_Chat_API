package models

type ChatMessage struct {
	Base
	Message string `json:"message"`
	UserID  int    `json:"user_id"`
	RoomID  int    `json:"room_id"`
}

func (c ChatMessage) TableName() string {
	return "chatmessage"
}
