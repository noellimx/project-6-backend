package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/gorilla/websocket"

	"proj6/gomoon/database"
	"proj6/gomoon/wss"
)

var connections = []*websocket.Conn{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func StoreMessageInTickerRoom(result *map[string]interface{}) {
	thisResult := *result

	//removed JWT for now, reason being that cant find username

	// tokenString := (thisResult)["token"]
	// fmt.Println(tokenString)

	// ddd := tokenString.(string)

	// claims := jwt.MapClaims{}
	// _, err := jwt.ParseWithClaims(ddd, claims, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte(JwtSecret), nil
	// })

	// username := claims["username"]
	// ... error handling
	message := (thisResult)["message"]
	_ = message
	username := (thisResult)["token"]
	roomId := (thisResult)["roomId"]
	ticker := roomId
	// if err != nil {
	// 	fmt.Println(username)
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("running storemessage func3")

	// // do something with decoded claims
	// for key, val := range claims {
	// 	fmt.Printf("Key: %v, value: %v\n", key, val)
	// }
	// // from token need to get username

	t := time.Now()
	_ = t

	// Store the result
	fmt.Println("storing message to db")
	database.AddToMessage(ticker.(string), message.(string), username.(string), t)
	// And broadcase message to room
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
		var result map[string]interface{}

		json.Unmarshal(p, &result)
		fmt.Println("text" + string(p))
		event := result["event"].(string)

		fmt.Println(event + "line 104")
		if event == "send-to-ticker-room" {
			fmt.Println("event for sendToTicker fire")
			StoreMessageInTickerRoom(&result)
			// get the result
			// broadcast to people who are online to this ticker room
		}
		fmt.Println("value of P", p)
		s := string(p)
		fmt.Println(s)
		wss.Broadcast(connections, p)

		// TODO try catch

		// ManageEvents(conn, p)

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

func GetHistoryData(w http.ResponseWriter, r *http.Request) {

}
