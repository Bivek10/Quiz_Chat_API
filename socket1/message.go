package socket1

import (
	"encoding/json"
	"fmt"
)

// Action: What action is the message requesting
const SendMessageAction = "send-message"
const LeaveRoomAction = "leave-room"

type Message struct {
	Action   string `json:"action"`
	Message  string `json:"message"`
	RoomId   int64  `json:"roomId"`
	SenderId int    `json:"senderId"`
}

func (message *Message) encode() []byte {
	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		fmt.Printf("Message struct values: %#v\n", message)
		return nil
	}
	return jsonData
}
