package routes

import "net/http"

func GetAllTickers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
