package database

import "testing"

func TestUserGetOrCreate(t *testing.T) {
	t.Fail()

	testEmail := "a" + RandomString(12) + "@b.com"
	user, isNew := GetOrCreate(testEmail)

	if user.Email != testEmail {
		t.Fatal("user email supplied not the same as retrieved")
	}
	if user.Username == "" {
		t.Fatal("username should not be empty.")
	}
	if isNew != true {
		t.Fatal("user should be new")
		t.Fail()
	}
}
