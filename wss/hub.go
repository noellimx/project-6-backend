package wss

import "github.com/gorilla/websocket"

func Boardcast(users []*websocket.Conn, p []byte) {
	for _, user := range users {
		user.WriteMessage(websocket.BinaryMessage, p)
	}
}
