package database

import (
	"proj6/gomoon/utils"
	"testing"
)

func TestCreateUser(t *testing.T) {
	testEmail := "a" + utils.RandomString(12) + "@b.com"
	username := "test" + utils.RandomString(12)

	CreateUser(testEmail, username)

	newUser := FindUserByEmail(testEmail)
	if newUser == nil {
		t.Fatal("user not created")
	}

	if newUser.Email != testEmail {
		t.Fatal("created user, but different username")
	}
}
func TestUserGetOrCreate(t *testing.T) {

	testEmail := "a" + utils.RandomString(12) + "@b.com"

	user, isNew := GetByEmailOrCreateUser(testEmail)
	if isNew != true {
		t.Fatal("user should be new")
	}

	if user.Email != testEmail {
		t.Fatal("user email supplied not the same as retrieved")
	}
	if user.Username == "" {
		t.Fatal("username should not be empty.")
	}
}
