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

func setFavourite(w http.ResponseWriter, r *http.Request) {

	tokenString := chi.URLParam(r, "token")
	ticker := chi.URLParam(r, "ticker")

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecret), nil
	})

	if err != nil {
		http.Error(w, "404 not found, from handler.", http.StatusNotFound)
		return
	}

	database.SetFavourite(&database.Favourite{
		Ticker: ticker, Email: claims["username"].(string),
	})

	w.WriteHeader(http.StatusOK)

}

func FavouritesRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/addtickertofavourite/{ticker}/{token}", setFavourite)
	r.Get("/getuserfavourite/{token}", getFavouritesOfUser)
	return r
}
