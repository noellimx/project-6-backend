package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func StaticRouter(staticDirectory string) http.Handler {
	r := chi.NewRouter()

	fileServer := http.FileServer(http.Dir(staticDirectory)) // desktop/static, but i want <repo>/static
	r.Handle("/*", fileServer)

	return r
}
