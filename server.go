package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"proj6/gomoon/routes"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"

	"github.com/markbates/goth/gothic"
)

func newAuthSessionStore() *sessions.CookieStore {

	key := "Secret-session-key" // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30        // 30 days
	isProd := false             // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	return store
}
func main() {

	gothic.Store = newAuthSessionStore()

	goth.UseProviders(google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), "http://localhost:3000/auth/google/callback"))

	// Routes

	r := chi.NewRouter()

	r.Mount("/", routes.StaticRouter())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, "Hi")
	})
	r.Mount("/dummy", routes.DummyRouter())

	r.Mount("/auth", routes.HTTPAuthRouter())

	// if err := http.ListenAndServe(":8080", nil); err != nil {
	//     log.Fatal(err)
	// }

	fmt.Println("Listening")
	if err := http.ListenAndServeTLS(":8080", "/Users/noellim/customkeystore/server.cert", "/Users/noellim/customkeystore/server.key", r); err != nil {
		log.Fatal(err)
	}

}
