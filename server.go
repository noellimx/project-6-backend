package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
    "log"
	"proj6/gomoon/routes"
)



func main() {

	r := chi.NewRouter()

	r.Mount("/", routes.StaticRouter())

	r.Mount("/", routes.DummyRouter())

	// if err := http.ListenAndServe(":8080", nil); err != nil {
	//     log.Fatal(err)
	// }
	if err := http.ListenAndServeTLS(":8080", "/Users/noellim/customkeystore/server.cert", "/Users/noellim/customkeystore/server.key", r); err != nil {
		log.Fatal(err)
	}

}
