package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"proj6/gomoon/database"

	"github.com/go-chi/chi/v5"
)

func getAllMessage(w http.ResponseWriter, r *http.Request) {
	enableCors := func(w *http.ResponseWriter) {
		(*w).Header().Set("Access-Control-Allow-Origin", "*")
	}
	enableCors(&w)
	allMessages := database.FindMessagesInDB()

	json.NewEncoder(w).Encode(&allMessages)
	fmt.Printf("running within getAllMessage middleware")
}

func getMessageFromTicker(w http.ResponseWriter, r *http.Request) {
	enableCors := func(w *http.ResponseWriter) {
		(*w).Header().Set("Access-Control-Allow-Origin", "*")
	}
	enableCors(&w)
	myTicker := chi.URLParam(r, "ticker")
	fmt.Println(myTicker)
	allMessages := database.FindMessagesByTicker(myTicker)

	json.NewEncoder(w).Encode(&allMessages)
	fmt.Printf("running within getMessageFromTicker middleware")
}

func MessageRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", getAllMessage)
	r.Get("/{ticker}", getMessageFromTicker)
	// r.Get("/{id}", getUser)
	return r
}
