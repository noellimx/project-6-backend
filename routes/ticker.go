package routes

import (
	"encoding/json"
	"net/http"
)

func GetAllTickers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func GetSearchValue(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	json.NewEncoder(w).Encode(&struct{}{})

}
