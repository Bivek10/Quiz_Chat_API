package socket

import (
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/services"
)

type WsServer struct {
	Client          map[*Client]bool
	Broadcast       chan []byte
	Register        chan *Client
	Unregister      chan *Client
	Rooms           map[*Room]bool
	ChatRoomService services.ChatRoomService
	Logger          infrastructure.Logger
}

func NewWebsocketServer(chatRoomServices services.ChatRoomService, logger infrastructure.Logger) *WsServer {
	chatServer := &WsServer{
		Broadcast:       make(chan []byte),
		Register:        make(chan *Client),
		Unregister:      make(chan *Client),
		Client:          make(map[*Client]bool),
		Rooms:           make(map[*Room]bool),
		ChatRoomService: chatRoomServices,
		Logger:          logger,
	}
	go chatServer.Run()
	return chatServer
}
func (server *WsServer) Run() {
	for {
		select {

		case client := <-server.Register:
			server.registerClient(client)

		case client := <-server.Unregister:
			server.unregisterClient(client)

		}
	}
}

// register client to the server
func (server *WsServer) registerClient(client *Client) {
	server.Client[client] = true
}

// Delete the clent from the server after it's is lost its connection
func (server *WsServer) unregisterClient(client *Client) {
	// broadcast its connectiion to all clients in room associated with this clients
	// make user status offline
	if _, ok := server.Client[client]; ok {
		delete(server.Client, client)
	}
}

// create room inside server
// it should be changed to create by id method

// create a room in database
func (server *WsServer) createRoom(id int, roomName string) *Room {
	roomModel := models.ChatRoom{Name: roomName}

	chatroom, err := server.ChatRoomService.CreateChatRoom(roomModel)
	if err != nil {
		server.Logger.Zap.Error("failed to create chat room")
		return nil
	}
	room := NewRoom(int(chatroom.ID))
	go room.RunRoom()
	server.Rooms[room] = true
	return room
}

// all the created room are saved in server . This features
// find room by name we may not need this later inted use search by id
// To find room by id .To add clients there . leave clients and send message to the room clients.

func (server *WsServer) findRoomByID(ID int) *Room {

	var foundRoom *Room
	
	chatRoom, err := server.ChatRoomService.GetOneChatRoom(int64(ID))

	if err != nil {
		server.Logger.Zap.Error("Unable to get chat room")
		return nil
	}

	for room := range server.Rooms {
		if room.GetId() == int(chatRoom.ID) {
			foundRoom = room
			break
		}
	}
	return foundRoom
}

func (server *WsServer) findClientByID(ID int) *Client {
	var foundClient *Client
	for client := range server.Client {
		if client.ID == ID {
			foundClient = client
			break
		}
	}
	return foundClient
}
