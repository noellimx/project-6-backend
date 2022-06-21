package wss

import "github.com/gorilla/websocket"

func Broadcast(users []*websocket.Conn, p []byte) {
	for _, user := range users {
		user.WriteMessage(websocket.BinaryMessage, p)
	}
}
