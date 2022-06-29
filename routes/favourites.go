package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"proj6/gomoon/database"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"
)

func getFavouritesOfUser(w http.ResponseWriter, r *http.Request) {

	tokenString := chi.URLParam(r, "token")

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecret), nil
	})
	if err != nil {
		http.Error(w, "404 not found, from handler.", http.StatusNotFound)
		return
	}

	favourites := database.GetFavouritesOfEmail(claims["username"].(string))

	resBody := &struct {
		Favourites []database.Favourite `json:"favourites"`
	}{
		Favourites: *favourites,
	}

	fmt.Println("getFavouritesOfUser")
	fmt.Println(resBody)
	json.NewEncoder(w).Encode(resBody)

	w.WriteHeader(http.StatusOK)

}

type setFavBody struct {
	Ticker string `json:"ticker"`
	Token  string `json:"token"`
}

func setFavourite(w http.ResponseWriter, r *http.Request) {

	fmt.Println("setFavourite")
	fmt.Println(r.Body)

	decoder := json.NewDecoder(r.Body)
	var body setFavBody
	err2 := decoder.Decode(&body)

	if err2 != nil {

		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	fmt.Println(body)

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(body.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecret), nil
	})
	if err != nil {
		http.Error(w, "404 not found, from handler.", http.StatusUnauthorized)
		return
	}

	database.SetFavourite(&database.Favourite{
		Ticker: body.Ticker, Email: claims["username"].(string),
	})

	w.WriteHeader(http.StatusOK)

	return

}

func FavouritesRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/addtickertofavourite", setFavourite)
	r.Get("/getuserfavourite/{token}", getFavouritesOfUser)
	return r
}
