package database

import (
	"math/rand"
	"proj6/gomoon/types"
)

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func AUser(email string, username string) *types.User {
	p := &types.User{}
	p.Email = email
	p.Username = username
	return p
}

func GetOrCreate(email string) (*types.User, bool) {
	return &types.User{}, false
}
