package socket

const welcomeMessage = "%s joined the room"

type Room struct {
	ID         int `json:"id"`
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

// NewRoom creates a new Room
func NewRoom(id int) *Room {
	return &Room{
		ID:        id,	
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
			room.broadcastToClientsInRoom(message.encode(), message.SenderId)
		}
	}
}
func (room *Room) GetId() int {
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

func (room *Room) broadcastToClientsInRoom(message []byte, clientId int) {
	// save this message to room
	// broadcast to all online users
	// get all client from the database and if some clients are not online than send them messages as notification when they are online
	for client := range room.Clients {
		if client.ID != clientId {
			client.Send <- message
		}
	}
}

