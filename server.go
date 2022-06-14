package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"proj6/gomoon/database"
	"proj6/gomoon/routes"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"

	"github.com/markbates/goth/gothic"

	"proj6/gomoon/config"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Environment int

const (
	Production Environment = iota
	Test
)

func ReadConfigV2(env Environment) *config.GlobalConfig {

	configFileParent := os.Getenv("HOME")

	var subpath string

	if env == Production {
		subpath = "production"
	} else if env == Test {
		subpath = "test"
	} else {
		log.Fatal("Environment not supported")
	}

	configFilePath := configFileParent + "/customkeystore/" + subpath + "/config.json"

	return config.ReadConfig(configFilePath)
}

var globalConfig = ReadConfigV2(Production)

var certFileParentVar = globalConfig.Https.Paths.CertFileParentVar
var certFilePathFileParent = os.Getenv(certFileParentVar)

var certificatePath = certFilePathFileParent + globalConfig.Https.Paths.Certificate
var keyPath = certFilePathFileParent + globalConfig.Https.Paths.Key

func newAuthSessionStore() *sessions.CookieStore {

	key := globalConfig.Session.Key
	maxAge := 60 * 60
	isProd := false

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = false // HttpOnly should always be enabled
	store.Options.Secure = isProd
	fmt.Println(store.Options.MaxAge)
	return store
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	var users []database.User

	allUser := database.Db.Find(&users)

	json.NewEncoder(w).Encode(&allUser)
	fmt.Printf("running within GetAllUsers function")
	fmt.Println(allUser)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("running within GetUser function")

	params := chi.URLParam(r, "id")
	fmt.Println(params)
	var users []database.User

	data := database.Db.First(&users, params)
	if len(users) == 0 {
		fmt.Println("no user found")
		return
	}
	fmt.Println(users[0].Email)
	fmt.Println(users[0].Username)

	json.NewEncoder(w).Encode(&data)
}

// func NewUser(w http.ResponseWriter, r *http.Request){

// 	email := chi.URLParam(r, "email")
// 	username := chi.URLParam(r, "username")

// 	var users []User
// 	user := db.Create(&user)

// }

func main() {

	rand.Seed(time.Now().UnixNano())

	database.Init("gomoon", &globalConfig.PSQL)
	defer database.Db.Close()

	gothic.Store = newAuthSessionStore()

	gothic.GetProviderName = routes.CustomGetProviderNameFromRequestWithChiFramework

	googleAuthCredentials := globalConfig.OAuth.Google
	googleCallbackUrl := "https://" + globalConfig.Network.Domain + ":" + globalConfig.Network.Port + "/auth/google/callback"
	goth.UseProviders(google.New(googleAuthCredentials.ClientId, googleAuthCredentials.ClientSecret, googleCallbackUrl))

	// Routes

	r := chi.NewRouter()

	r.Mount("/", routes.StaticRouter())
	r.Get("/users", GetAllUsers)
	r.Get("/users/{id}", GetUser)
	// r.Post("/newuser/{email}/{username}", NewUser)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		s, err := r.Cookie("_gothic_session")

		if err == nil {
			fmt.Println(s.Value)

		}
		fmt.Fprint(w, "Hi")
	})
	r.Mount("/dummy", routes.DummyRouter())

	r.Mount("/auth", routes.HTTPAuthRouter())

	func() {
		fqdn := globalConfig.Network.Domain + ":" + globalConfig.Network.Port
		fmt.Println("Server listening on " + fqdn + "...")
		if err := http.ListenAndServeTLS(fqdn, certificatePath, keyPath, r); err != nil {
			log.Fatal(err)
		}

	}()
	fmt.Println("Server gracefully ended.")
}
