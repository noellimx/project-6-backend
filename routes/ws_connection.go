package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/gorilla/websocket"

	"proj6/gomoon/wss"
)

var connections = []*websocket.Conn{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func StoreMessageInTickerRoom(result *map[string]interface{}) {

	token := (*result)["token"]
	_ = token
	// from token need to get username

	message := (*result)["message"]
	_ = message

	roomId := (*result)["roomId"]
	_ = roomId

	t := time.Now()
	_ = t

	// Store the result

	// And broadcase message to room
}
func ManageEvents(conn *websocket.Conn, p []byte) {

	var result map[string]interface{}

	json.Unmarshal(p, &result)

	event := result["event"].(string)

	fmt.Println(event + "event: ")

	if event == "send-to-ticker-room" {

		StoreMessageInTickerRoom(&result)
	}

}

// listen indefinitely for new messages coming
// through on our WebSocket connection
func listenToWsConnection(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// print out that message for clarity
		log.Println("[listenToWsConnection]")
		fmt.Println("text" + string(p))

		wss.Broadcast(connections, p)

		// TODO try catch

		ManageEvents(conn, p)

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {

	fmt.Println("wsEndpoint")
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

	connections = append(connections, ws)

	listenToWsConnection(ws)

}

func UpGradeToWsRouter() http.Handler {
	r := chi.NewRouter()
	r.HandleFunc("/", wsEndpoint)

	return r
}
