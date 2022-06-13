package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func StaticRouter() http.Handler {
	r := chi.NewRouter()
	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/", fileServer)

	return r
}
