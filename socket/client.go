package socket

import (
	"encoding/json"
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
	rooms    map[*Room]bool //

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
	//register clients to multiple room at a time
	//get room from database and do
	go client.writeMessage()
	go client.readMessage()
}

func newClient(conn *websocket.Conn, wsServer *WsServer, id int) *Client {
	return &Client{
		ID:       id,
		Conn:     conn,
		rooms:    make(map[*Room]bool),
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
	for room := range client.rooms {
		room.Unregister <- client
	}
}

func (client *Client) findRoomByID(ID int) *Room {
	var Room *Room
	for room := range client.rooms {
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

	case JoinRoomAction:
		client.handleJoinRoomMessage(message)

	case LeaveRoomAction:
		client.handleLeaveRoomMessage(message)
	}
	//  user online status message  to all the clients in the user room
}

func (client *Client) handleSendMessage(message Message) {
	room := client.findRoomByID(message.RoomId)
	if room == nil {
		println("The room you are trying to send message doesnot present")
		return
	}
	room.Broadcast <- &message
}


func (client *Client) handleJoinRoomMessage(message Message) {
	roomID := message.RoomId
	client.joinRoom(roomID, message.RoomName)
}

//there should be another create room function

func (client *Client) joinRoom(roomID int, roomName string) {

	room := client.wsServer.findRoomByID(roomID)
	if room == nil {
		room = client.wsServer.createRoom(roomID, roomName)
	}
	//check if client is in the room (database)before or not
	//if not add the room to this user
	client.rooms[room] = true
	room.Register <- client
}

func (client *Client) handleLeaveRoomMessage(message Message) {
	// delete  this room from client in  the database
	roomId := message.RoomId
	room := client.wsServer.findRoomByID(roomId)
	if _, ok := client.rooms[room]; ok {
		delete(client.rooms, room)
	}
	room.Unregister <- client
}
