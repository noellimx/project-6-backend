package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"

	"github.com/gorilla/websocket"

	"proj6/gomoon/database"
	"proj6/gomoon/wss"
)

type BroadcastPayload struct {
	Event   string
	Message *database.Message
}

var connections = []*websocket.Conn{}

var exists = struct{}{}

type hub struct {
	connectionsInRoom map[string]map[*websocket.Conn]struct{}
	roomOfConnection  map[*websocket.Conn]string
}

func NewHub() *hub {
	s := &hub{}
	s.connectionsInRoom = make(map[string]map[*websocket.Conn]struct{})
	s.roomOfConnection = make(map[*websocket.Conn]string)
	return s
}

// 1. Subscribe connection to room. If connection is subscriber, it drops the subscription
// 2. if the target room has nil collection, construct new empty collection
// 3. Subscribe.
func (s *hub) AddConnectionToRoom(roomIdToSubscribe string, conn *websocket.Conn) {
	fmt.Println("[(s *hub) AddConnectionToRoom]")
	// 1.
	subscribingToRoomId, isSubscribed := s.roomOfConnection[conn]

	if isSubscribed {
		fmt.Println("[(s *hub) AddConnectionToRoom] Connection was a subscriber. Dropping previous subscription.")

		delete(s.connectionsInRoom[subscribingToRoomId], conn)
	}
	// 2.
	if s.connectionsInRoom[roomIdToSubscribe] == nil {
		s.connectionsInRoom[roomIdToSubscribe] = make(map[*websocket.Conn]struct{})
	}

	// 3.
	s.connectionsInRoom[roomIdToSubscribe][conn] = exists

}

func (s *hub) GetConnectionsInRoom(roomId string) []*websocket.Conn {

	connarray := []*websocket.Conn{}

	for conn, _ := range s.connectionsInRoom[roomId] {
		connarray = append(connarray, conn)
	}

	return connarray
}

var myhub = NewHub()

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func newEventMessageFromFE(p []byte) (string, *database.Message, string) {

	var result map[string]interface{}

	json.Unmarshal(p, &result)
	roomId := result["roomId"].(string)
	username := result["username"].(string)
	message := result["message"].(string)
	tString := result["time"].(string)

	fmt.Println("from newEventMessage printing Time", tString)

	// t := result["time"].(time.Time)
	t, _ := time.Parse(time.RFC3339, tString)
	//2022-06-24T15:10:44.238Z
	tokenString := result["token"].(string)

	return "", database.NewMessage(roomId, message, username, t), tokenString

}

func newEventSubscribeFromFE(p []byte) (string, string) {
	fmt.Println("newEventSubscribeFromFE")
	var result map[string]interface{}
	json.Unmarshal(p, &result)
	roomId := result["roomId"].(string)
	tokenString := result["token"].(string)
	fmt.Println(tokenString + roomId)
	return roomId, tokenString
}

func getEventNameFromFE(p []byte) string {
	var result map[string]interface{}
	json.Unmarshal(p, &result)
	return result["event"].(string)
}

func StoreMessageInTickerRoom(message *database.Message, tokenString string) {

	// Authentication before storing data

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecret), nil
	})

	if err != nil {
		fmt.Println("Error authenticating for database activity")
		return
	}

	// storing data

	if err != nil {
		fmt.Println(err)
		return
	}

	// // do something with decoded claims
	// for key, val := range claims {
	// }
	// // from token need to get username

	// Store the result
	fmt.Println("storing message to db")
	fmt.Println(message)

	database.AddToMessage(message)
	// And broadcase message to room
}

func ChatHistory(result *map[string]interface{}) {
	thisResult := *result
	roomId := (thisResult)["roomId"]
	ticker := roomId
	fmt.Println("chatHistory func, ticker", ticker)

}

// listen indefinitely for new messages coming
// through on our WebSocket connection
func listenToWsConnection(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, rawPayload, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		event := getEventNameFromFE(rawPayload)
		log.Println("[listenToWsConnection] Event: " + event)

		if event == "send-to-ticker-room" {
			_, message, tokenString := newEventMessageFromFE(rawPayload)
			//1. store message sent from frontend to backend
			StoreMessageInTickerRoom(message, tokenString)
			// 2. we need to get roomId and username received from frontend, roomId to make sure that socket only sends msg to users that are in the same roomId

			messageThatIsStoredInDatabase := database.FindMessageByUsernameTime(message.Username, message.Time)

			if messageThatIsStoredInDatabase == nil {
				return
			}

			var messagePayload BroadcastPayload
			messagePayload.Event = event
			messagePayload.Message = message
			fmt.Println("messagePayload ", messagePayload)
			messageMarshal, _ := json.Marshal(messagePayload)
			fmt.Println(string(messageMarshal))

			roomId := messageThatIsStoredInDatabase.RoomTicker

			connArray := myhub.GetConnectionsInRoom(roomId)

			wss.Broadcast(connArray, []byte(messageMarshal))

			//4. only send msg to specific roomId
			fmt.Println("event for sendToTicker fire")

		} else if event == "subsribe-to-ticker-room" {
			roomId, _ := newEventSubscribeFromFE(rawPayload)
			myhub.AddConnectionToRoom(roomId, conn)
		}

		//on connection, username will be added to

		//3. add username to the room that they are in
		//   remove username when user disconnects

		// TODO try catch

		// ManageEvents(conn, p)

		if err := conn.WriteMessage(messageType, rawPayload); err != nil {
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

	// wss.Broadcast(connections, p)

	connections = append(connections, ws)

	// tickerRoom := result["roomId"].(string)
	// fmt.Println("ticker room in listentoWsConnect func", tickerRoom)

	listenToWsConnection(ws)

}

func UpGradeToWsRouter() http.Handler {
	r := chi.NewRouter()
	r.HandleFunc("/", wsEndpoint)

	return r
}

func GetHistoryData(w http.ResponseWriter, r *http.Request) *[]database.Message {
	data := database.FindMessagesInDB()
	fmt.Println(data)
	dataType := reflect.TypeOf(data)
	fmt.Println(dataType)

	return data
}
