package routes

import "net/http"

var EnableCors = func(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
