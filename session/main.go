package session

import (
	"fmt"

	"github.com/gorilla/sessions"
)

func NewAuthSessionStore(key string) *sessions.CookieStore {
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
