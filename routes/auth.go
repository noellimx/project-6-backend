package routes

import (
	"errors"
	"fmt"
	"proj6/gomoon/config"
	"proj6/gomoon/database"
	"proj6/gomoon/utils"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"

	"net/http"

	jwt "github.com/golang-jwt/jwt/v4"
)

func CustomGetProviderNameFromRequestWithChiFramework(r *http.Request) (string, error) {
	fmt.Println("CustomGetProviderNameFromRequestWithChiFramework")
	str := chi.URLParam(r, "provider")
	fmt.Println(str)

	if str == "" {
		return "", errors.New("provider not found")

	}

	return str, nil
}

var routerName = "Router - HTTPAuth"

var jwtSecret = []byte(utils.RandomString(256))
var JwtSecret = jwtSecret

func HTTPAuthRouter(networkConfig config.Network) http.Handler {

	r := chi.NewRouter()

	r.HandleFunc("/login/{provider}", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(routerName + "/{provider}")
		fmt.Println(gothic.Store)

		s, err := r.Cookie("_gothic_session")

		if err == nil {
			fmt.Println("GothicSession" + s.Value)

		}

		fmt.Println(gothic.Store)
		provider, err := goth.GetProvider("google")
		if err != nil {
			fmt.Println("cannot GetFromSession")
		}

		value, err := gothic.GetFromSession("google", r)
		if err != nil {
			fmt.Println("cannot GetFromSession")
		}
		fmt.Println("can GetFromSession")

		fmt.Println(value)
		sess, err := provider.UnmarshalSession(value)

		sess.GetAuthURL()
		if err != nil {
			fmt.Println("cannot UnmarshalSession")
		} else {

			fmt.Println("can UnmarshalSession")

			fmt.Println(sess)

		}

		sss, err := provider.FetchUser(sess)
		_ = sss
		if err != nil {
			// user can be found with existing session data
			fmt.Println("cannot fetch")
		}
		fmt.Println("can fetch")

		fmt.Println(r.URL.Query())
		if _, err := gothic.CompleteUserAuth(w, r); err == nil {
			fmt.Println("existing login")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			gothic.BeginAuthHandler(w, r)
		}
		fmt.Println(gothic.Store)

	})

	r.HandleFunc("/{provider}/callback", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(routerName + "/{provider}/callback")

		user, err := gothic.CompleteUserAuth(w, r)

		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		dbUser, _ := database.GetByEmailOrCreateUser(user.Email)

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": dbUser.Email,
		})

		jwtStr, _ := token.SignedString(JwtSecret)

		cookie := &http.Cookie{
			Name:  "gm-token",
			Value: jwtStr,
			Path:  "/",
		}

		http.SetCookie(w, cookie)
		http.Redirect(w, r, "https://"+networkConfig.Domain+":"+networkConfig.Port+"/main.html", http.StatusMovedPermanently)
	})

	// option 1: login with google
	// User: {id, email, username }
	// when someone logins with google, check DB for the same email, if doesnt exist = new user/give random username
	// response will include the jwt credentials as cookies for browser

	r.HandleFunc("/logout/{provider}", func(w http.ResponseWriter, r *http.Request) {

		gothic.Logout(w, r)
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusTemporaryRedirect)
	})

	r.HandleFunc("/test/{provider}", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(routerName + "/test/{provider}")
		fmt.Println(gothic.Store)

		fmt.Println(r.URL.Query())

		user, err := gothic.CompleteUserAuth(w, r)

		defer fmt.Println(gothic.Store)
		if err != nil {

			fmt.Println("Error ")
			fmt.Println(r.URL.Query())
			fmt.Println(r.URL.Query())

			w.WriteHeader(http.StatusNonAuthoritativeInfo)
			fmt.Fprint(w, err)
			return
		}

		fmt.Fprintf(w, user.Email)

	})

	return r
}
