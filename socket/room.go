package socket

import "github.com/bivek/fmt_backend/models"

type Room struct {
	ID         int64 `json:"id"`
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
	Isonline   chan bool
}

// NewRoom creates a new Room
func NewRoom(id int64) *Room {
	return &Room{
		ID:         id,
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message),
	}
}

// RunRoom runs our room, accepting various requests
func (room *Room) RunRoom() {
	for {
		select {
		case client := <-room.Register:
			room.registerClientInRoom(client)
		case client := <-room.Unregister:
			room.unregisterClientInRoom(client)
		case message := <-room.Broadcast:
			room.broadcastToClientsInRoom(message.encode(), message)
		}
	}
}

func (room *Room) GetId() int64 {
	println("room - room ID ", room.ID)
	return room.ID
}

// register and notify others users
func (room *Room) registerClientInRoom(client *Client) {

	room.Clients[client] = true
}

// unregister client from the room
func (room *Room) unregisterClientInRoom(client *Client) {
	// delete this client from the room database
	if _, ok := room.Clients[client]; ok {
		delete(room.Clients, client)
	}
}
func (room *Room) broadcastToClientsInRoom(message []byte, messages *Message) {
	// save this message to room
	// broadcast to all online users
	// get all client from the database and if some clients are not online than send them messages as notification when they are online
	for client := range room.Clients {
		// print("client list", client.ID)
		if client.ID != messages.SenderId {
			messageModel := models.ChatMessage{Message: messages.Message, UserID: messages.SenderId, RoomID: int(messages.RoomId)}
			dbMessage := client.wsServer.saveMessage(messageModel)
			if dbMessage != nil {
				client.wsServer.saveMessage(messageModel)
				client.Send <- message
				println("message detail", message)
			}

		}
	}
}
