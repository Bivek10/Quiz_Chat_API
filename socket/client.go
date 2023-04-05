package socket

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/bivek/fmt_backend/helpers"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	// Max wait time when writing message to peer
	writeWait = 10 * time.Second
	// Max time till next pong from peer
	pongWait = 60 * time.Second
	// Send ping interval, must be less then pong wait time
	pingPeriod = (pongWait * 9) / 10
	// Maximum message size allowed from peer.
	maxMessageSize = 10000
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// client must send their own id whiler joing the room	q
type Client struct {
	ID       int `json:"id"`
	wsServer *WsServer
	Conn     *websocket.Conn
	Send     chan []byte
	rooms    map[int64]*Room //
}

func ServeWs(wsServer *WsServer, c *gin.Context) {
	conn, err := helpers.Upgrade(c.Writer, c.Request)
	if err != nil {
		println("the errror is", err)
	}
	id := c.Query("id")
	userid, errs := strconv.Atoi(id)

	if errs != nil {
		log.Println("Url Param 'Id' is invalid", errs)
		return
	}
	client := *newClient(conn, wsServer, userid)

	// set the user status to online
	// broadcast to all the users of the users room online message
	wsServer.Register <- &client

	client.joinUserInAllRoom(int64(client.ID))
	go client.writeMessage()
	go client.readMessage()
}

func newClient(conn *websocket.Conn, wsServer *WsServer, id int) *Client {
	return &Client{
		ID:       id,
		Conn:     conn,
		rooms:    make(map[int64]*Room),
		wsServer: wsServer,
		Send:     make(chan []byte),
	}
}

func (client *Client) readMessage() {
	defer func() {
		client.disconnect()
	}()
	client.Conn.SetReadLimit(maxMessageSize)
	client.Conn.SetReadDeadline(time.Now().Add(pongWait))
	client.Conn.SetPongHandler(func(string) error { client.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, jsonMessage, err := client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected close error: %v", err)
			}
			break
		}
		println("json handle mesesae", jsonMessage)
		client.handleNewMessages(jsonMessage)
	}
}

func (client *Client) writeMessage() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		client.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-client.Send:
			client.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The WsServer closed the channel.
				client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := client.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			// Attach queued chat messages to the current websocket message.
			n := len(client.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-client.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			client.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (client *Client) GetId() int {
	return client.ID
}

// disconnect client from the websocket server and all the rooms he/she was present in
func (client *Client) disconnect() {
	client.wsServer.Unregister <- client
	for _, room := range client.rooms {
		room.Unregister <- client
	}
}

func (client *Client) findRoomByID(ID int64) *Room {
	var Room *Room
	for _, room := range client.rooms {
		fmt.Println("room id", room.ID)
		fmt.Println("length", len(client.rooms))
		fmt.Println("get id room", room.GetId())
		fmt.Println("get id room ID", ID)
		if room.GetId() == ID {
			Room = room
			break
		}
	}
	return Room
}

func (client *Client) handleNewMessages(jsonMessage []byte) {
	var message Message
	if err := json.Unmarshal(jsonMessage, &message); err != nil {
		log.Printf("Error on unmarshal JSON message %s", err)
		return
	}
	//Attach the client id as the sender of the messsage.
	message.SenderId = client.ID

	switch message.Action {
	case SendMessageAction:
		//save the message in database over here
		client.handleSendMessage(message)

	case LeaveRoomAction:
		client.handleLeaveRoomMessage(message)
	}

}

func (client *Client) handleSendMessage(message Message) {
	room := client.findRoomByID(message.RoomId)
	fmt.Println("room =======id", room.ID)
	if room == nil {
		println("The room you are trying to send message doesnot present")
		return
	}

	room.Broadcast <- &message

}

// there should be another create room function

func (client *Client) handleLeaveRoomMessage(message Message) {
	// delete  this room from client in  the database
	roomId := message.RoomId
	dbRoom := client.wsServer.findRoomByID(roomId)
	if dbRoom != nil {
		if room, ok := client.rooms[message.RoomId]; ok {
			fmt.Println("clients::::::", client.ID)
			err := client.wsServer.ChatMembeService.DeleteOneChatMember(int64(client.ID))
			if err != nil {
				client.wsServer.Logger.Zap.Error("Failed to delete member", err)
			}
			room.Unregister <- client
			delete(client.rooms, room.ID)
		}
	}
}

// join the member in room once connected to web server.
func (client *Client) joinUserInAllRoom(userID int64) {
	var chatRoom *Room
	// db room finding and creation
	dbAllRoom, _, err := client.wsServer.ChatRoomService.GetAllRoomByUserID(userID)

	if err != nil {
		fmt.Println("error:get all room by user", err)
	}
	for i := range dbAllRoom {
		dbRoom := dbAllRoom[i]
		chatRoom = client.wsServer.Rooms[dbRoom.RoomID]
		if chatRoom == nil {
			room := NewRoom(dbRoom.RoomID)
			client.wsServer.Rooms[dbRoom.RoomID] = room
			chatRoom = client.wsServer.Rooms[dbRoom.RoomID]
			go room.RunRoom()
		}
		client.rooms[chatRoom.ID] = chatRoom
		chatRoom.Register <- client
		
	}
	fmt.Println("chat room length", len(client.rooms))
}
