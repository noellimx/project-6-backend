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

	"proj6/gomoon/config"
)

var configFileParent = os.Getenv("HOME")
var configFilePath = configFileParent + "/customkeystore/config.json"

var globalConfig = config.ReadConfig(configFilePath)

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

func main() {

	gothic.Store = newAuthSessionStore()

	gothic.GetProviderName = routes.CustomGetProviderNameFromRequestWithChiFramework

	googleAuthCredentials := globalConfig.OAuth.Google
	googleCallbackUrl := "https://" + globalConfig.Network.Domain + ":" + globalConfig.Network.Port + "/auth/google/callback"
	goth.UseProviders(google.New(googleAuthCredentials.ClientId, googleAuthCredentials.ClientSecret, googleCallbackUrl))

	// Routes

	r := chi.NewRouter()

	r.Mount("/", routes.StaticRouter())

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
