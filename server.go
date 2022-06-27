package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"proj6/gomoon/config"
	"proj6/gomoon/database"
	"proj6/gomoon/routes"
	"proj6/gomoon/session"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"

	"github.com/markbates/goth/gothic"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func EnvSubPath(env config.Environment) string {
	var subpath string

	if env == config.Production {
		subpath = "production"
	} else if env == config.Test {
		subpath = "test"
	} else {
		log.Fatal("Environment not supported")
	}

	return subpath

}

func main() {

	var globalConfig = config.ReadConfig(config.Production)

	var staticDirectory string

	if globalConfig.StaticDirectory == "" {
		defaultStaticDirectory := "./static"
		fmt.Println("static directory not specified in config. default to : " + defaultStaticDirectory)
		staticDirectory = defaultStaticDirectory
	} else {
		staticDirectory = globalConfig.StaticDirectory
	}
	fmt.Println("static directory is : " + staticDirectory)

	var certFileParentVar = globalConfig.Https.Paths.CertFileParentVar
	var certFilePathFileParent = os.Getenv(certFileParentVar)

	var envSubPath = "/" + EnvSubPath(config.Production)
	var certificatePath = certFilePathFileParent + "/customkeystore" + envSubPath + globalConfig.Https.Paths.Certificate
	var keyPath = certFilePathFileParent + "/customkeystore" + envSubPath + globalConfig.Https.Paths.Key

	rand.Seed(time.Now().UnixNano())

	database.Init(&globalConfig.PSQL)
	defer database.Db.Close()

	gothic.Store = session.NewAuthSessionStore(globalConfig.Session.Key)

	gothic.GetProviderName = routes.CustomGetProviderNameFromRequestWithChiFramework

	googleAuthCredentials := globalConfig.OAuth.Google
	googleCallbackUrl := "https://" + globalConfig.Network.Domain + ":" + globalConfig.Network.Port + "/auth/google/callback"
	goth.UseProviders(google.New(googleAuthCredentials.ClientId, googleAuthCredentials.ClientSecret, googleCallbackUrl))

	routes.JwtSecret = []byte(globalConfig.JWT.Secret)
	// Routes
	r := chi.NewRouter()

	// Welcome Message
	r.Mount("/", routes.StaticRouter(staticDirectory))

	serverInitTime := time.Now()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		s, err := r.Cookie("_gothic_session")
		fmt.Println("accessing / route")
		if err == nil {
			fmt.Println(s.Value)

		}
		fmt.Fprint(w, "Hi: "+serverInitTime.String())
	})
	r.Mount("/dummy", routes.DummyRouter())

	r.Mount("/users", routes.UserRouter())

	r.Mount("/auth", routes.HTTPAuthRouter(globalConfig.Network))

	r.Mount("/ws", routes.UpGradeToWsRouter())

	r.Mount("/history", routes.MessageRouter())

	r.Mount("/ticker", routes.TickerRouter())

	func() {
		fqdn := ":" + globalConfig.Network.Port
		fmt.Println("Server listening on " + fqdn + "...")
		if err := http.ListenAndServeTLS(fqdn, certificatePath, keyPath, r); err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Println("Server gracefully ended.")
}
