package socket

import (
	"encoding/json"
	"fmt"
)

// Action: What action is the message requesting
const SendMessageAction = "send"
const LeaveRoomAction = "leave"

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
