package database

import (
	"proj6/gomoon/types"
	"proj6/gomoon/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Username string `gorm:"unique"`
}

func AUser(email string, username string) *types.User {
	p := &types.User{}
	p.Email = email
	p.Username = username
	return p
}

func GetByEmailOrCreateUser(email string) (newOrExistingUser *User, isNew bool) {

	u := FindUserByEmail(email)

	if u != nil {
		return u, false
	}

	username := utils.RandomString(12)

	CreateUser(email, username)

	createdUser := FindUserByEmail(email)

	return createdUser, true
}

func CreateUser(email string, username string) {

	user := User{}
	user.Email = email
	user.Username = username

	Db.Create(&user)
}

func FindUserByEmail(email string) *User {

	var users []User

	user := User{}
	user.Email = email

	Db.Where(user).First(&users)

	if len(users) == 0 {
		return nil
	}

	return &users[0]

}

// user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

// result := db.Create(&user) /

// // User not found, initialize it with give conditions and Assign attributes
// db.Where(User{Name: "non_existing"}).Assign(User{Age: 20}).FirstOrInit(&user)
// // user -> User{Name: "non_existing", Age: 20}

// // Found user with `name` = `jinzhu`, update it with Assign attributes
// db.Where(User{Name: "Jinzhu"}).Assign(User{Age: 20}).FirstOrInit(&user)
// result := db.Where(User{Name: "jinzhu"}).FirstOrCreate(&user)
