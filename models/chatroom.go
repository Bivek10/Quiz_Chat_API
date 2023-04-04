package models

type ChatRoom struct {
	Base
	Name       string       `json:"name"`
	ChatMember []ChatMember `json:"chatmember" gorm:"foreignKey:room_id"`
}

func (c ChatRoom) TableName() string {
	return "chatroom"
}

type ChatRoomMember struct {
	ChatRoom
	ChatMember []ChatMember `json:"chatmember" gorm:"foreignKey:room_id"`
}

type RoomMember struct {
	RoomID   int64  `json:"room_id"`
	RoomName string `json:"room_name"`
	Member   Member `json:"member"`
}

type Member struct {
	SendeID    int64 `json:"sender_id"`
	ReceiverID int64 `json:"receiver_id"`
}
