package routes

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"log"
	"fmt"

	"github.com/gorilla/websocket"

    "proj6/gomoon/wss"

)


var users = []*websocket.Conn{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// print out that message for clarity
		log.Println("sending from here")
		fmt.Println("text" + string(p))

		log.Println(users)

		wss.Boardcast(users, p)

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	users = append(users, ws)

	reader(ws)

}

func UpGradeToWsRouter() http.Handler {
    r := chi.NewRouter()
    r.HandleFunc("/ws", wsEndpoint)

    return r
}