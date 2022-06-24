package wss

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func Broadcast(connections []*websocket.Conn, p []byte) {
	fmt.Println("broadcasting to everyone all channel")
	for _, connection := range connections {
		connection.WriteMessage(websocket.BinaryMessage, p)
	}
}

func BroadcastToTickerRoom(roomId string, message string) {
	fmt.Println("[BroadcastToTickerRoom]")
	// data := string(p)
	// fmt.Println(data.roomId)

	var conns []*websocket.Conn

	// get the collection of conns we want to send

	var payload []byte

	// prepare the payload we want to provide
	Broadcast(conns, payload)
}
