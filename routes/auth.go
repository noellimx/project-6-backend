package routes

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"

	"net/http"
)

var routerName = "Router - HTTPAuth"

func HTTPAuthRouter() http.Handler {

	r := chi.NewRouter()

	r.HandleFunc("/{provider}", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(routerName + "/{provider}")

		fmt.Println(r.URL.Query())
		fmt.Println(r.URL.Query().Get("provider"))

		newQ := r.URL.Query()
		newQ.Add("provider", chi.URLParam(r, "provider"))
		r.URL.RawQuery = newQ.Encode()

		fmt.Println(r.URL.Query())


		if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {

			fmt.Println("authenticated")
			fmt.Println("id" + gothUser.UserID)
			fmt.Println("email" + gothUser.Email)

		} else {
			gothic.BeginAuthHandler(w, r)
		}
	})

	return r
}
