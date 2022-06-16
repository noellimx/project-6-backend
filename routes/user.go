package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"proj6/gomoon/database"

	"github.com/go-chi/chi/v5"
)

func getAllUsers(w http.ResponseWriter, r *http.Request) {

	var users []database.User

	allUser := database.Db.Find(&users)

	json.NewEncoder(w).Encode(&allUser)
	fmt.Printf("running within getAllUsers middleware")
	fmt.Println(allUser)
}

func newUser(w http.ResponseWriter, r *http.Request) {

	// email := chi.URLParam(r, "email")
	// username := chi.URLParam(r, "username")

	// var users []User
	// user := db.Create(&user)

	w.WriteHeader(http.StatusNotImplemented)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("running within getUser middleware")

	param := chi.URLParam(r, "id")
	var users []database.User

	data := database.Db.First(&users, param)
	if len(users) == 0 {
		fmt.Println("no user found")
		return
	}

	json.NewEncoder(w).Encode(&data)
}

func UserRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", getAllUsers)
	r.Get("/{id}", getUser)
	r.Post("/newuser/{email}/{username}", newUser)
	return r
}
