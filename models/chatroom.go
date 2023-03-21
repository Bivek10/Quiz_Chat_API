package models

type ChatRoom struct {
	Base
	Name string `json:"name"`
}

func (c ChatRoom) TableName() string {
	return "chatroom"
}
