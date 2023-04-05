package socket

import (
	"fmt"

	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/services"
)

type WsServer struct {
	Client           map[*Client]bool
	Broadcast        chan []byte
	Register         chan *Client
	Unregister       chan *Client
	Rooms            map[int64]*Room
	ChatRoomService  services.ChatRoomService
	ChatMembeService services.ChatMemberService
	Logger           infrastructure.Logger
	MessageService   services.ChatMessageService
}


func NewWebsocketServer(chatRoomServices services.ChatRoomService, chatmemeberService services.ChatMemberService, logger infrastructure.Logger, messageService services.ChatMessageService) *WsServer {
	chatServer := &WsServer{
		Broadcast:        make(chan []byte),
		Register:         make(chan *Client),
		Unregister:       make(chan *Client),
		Client:           make(map[*Client]bool),
		Rooms:            make(map[int64]*Room),
		ChatRoomService:  chatRoomServices,
		Logger:           logger,
		ChatMembeService: chatmemeberService,
		MessageService:   messageService,
	}
	//go chatServer.Run()
	return chatServer
}

func (server *WsServer) Run() {
	for {
		select {
		case client := <-server.Register:
			fmt.Println("rooms :::::::",client.rooms)
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

func (server *WsServer) createRoom(roomName string, roomID int64) *models.ChatRoom {
	roomModel := models.ChatRoom{Name: roomName, Base: models.Base{ID: roomID}}
	chatroom, err := server.ChatRoomService.CreateChatRoom(roomModel)
	if err != nil {
		server.Logger.Zap.Error("failed to create chat room")
		return nil
	}
	return &chatroom
}

func (server *WsServer) addMemberInRoom(roomID int64, userID int64) *models.ChatMember {
	chatMember, err := server.ChatMembeService.CreateChatMember(models.ChatMember{UserID: userID, RoomID: roomID})
	if err != nil {
		server.Logger.Zap.Error("failed to add member in  room")
		return nil
	}
	return &chatMember
}

func (server *WsServer) findMemberInRoom(roomID int64, userID int64) *models.ChatMember {
	chatMember, err := server.ChatMembeService.GetOneChatMember(int64(userID), roomID)
	if err != nil {
		server.Logger.Zap.Error("Unable to get member in room ID", roomID)
		return nil
	}
	return &chatMember
}

func (server *WsServer) findRoomByID(ID int64) *models.ChatRoom {
	chatRoom, err := server.ChatRoomService.GetOneChatRoom(int64(ID))

	if err != nil {
		server.Logger.Zap.Error("Unable to get chat room")
		return nil
	}

	return &chatRoom
}

func (server *WsServer) saveMessage(messageModel models.ChatMessage) *models.ChatMessage {
	
	dbMessage, err := server.MessageService.CreateChatMessage(messageModel)
	if err != nil {
		server.Logger.Zap.Error("Failed to save data in database")
		return nil
	}
	return &dbMessage

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
