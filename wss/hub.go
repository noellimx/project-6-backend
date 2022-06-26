package wss

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func Broadcast(connections []*websocket.Conn, p []byte) {

	count := fmt.Sprint(len(connections))
	fmt.Println("broadcasting to " + count + " channel")
	fmt.Println(p)
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
