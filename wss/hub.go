package wss

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func Broadcast(connections []*websocket.Conn, p []byte) {
	fmt.Println("broadcasting to everyone in channel")
	for _, connection := range connections {

		connection.WriteMessage(websocket.BinaryMessage, p)

	}
}
