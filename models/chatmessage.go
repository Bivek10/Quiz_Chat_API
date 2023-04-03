package models

type ChatMessage struct {
	Base
	Message string `json:"message"`
	UserID  string `json:"user_id"`
	RoomID  string `json:"room_id"`
}

func (c ChatMessage) TableName() string {
	return "chatmessage"
}
