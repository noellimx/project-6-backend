package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Favourite struct {
	gorm.Model
	Email  string `gorm:"primaryKey"`
	Ticker string `gorm:"primaryKey"`
}

func AFavourite(email string, username string) *Favourite {
	p := &Favourite{}
	p.Email = email
	p.Ticker = username
	return p
}

func SetFavourite(favourite *Favourite) {

	var favs []Favourite

	Db.Where(favourite).First(&favs)

	if len(favs) == 0 {
		Db.Create(favourite)
	}
}

func RemoveFavourite(favourite *Favourite) {

	var favs []Favourite

	Db.Where(favourite).First(&favs)
	if len(favs) != 0 {

		fmt.Println("RemoveFavourite")
		fmt.Println(favs)
		Db.Delete(favs[0])
	}

}
func GetFavouritesOfEmail(email string) *[]Favourite {
	var favs []Favourite
	targetFav := Favourite{}
	targetFav.Email = email

	Db.Find(&favs, targetFav)

	return &favs

}

// func GetByEmailOrCreateUser(email string) (newOrExistingUser *User, isNew bool) {

// 	u := FindUserByEmail(email)

// 	if u != nil {
// 		return u, false
// 	}

// 	username := utils.RandomString(12)

// 	CreateUser(email, username)

// 	createdUser := FindUserByEmail(email)

// 	return createdUser, true
// }

// func CreateFavourite(email string, ticker string) {

// 	user := User{}
// 	user.Email = email
// 	user.Username = username

// 	Db.Create(&user)
// }

// func FindUserByEmail(email string) *User {

// 	var users []User

// 	user := User{}
// 	user.Email = email

// 	Db.Where(user).First(&users)

// 	if len(users) == 0 {
// 		return nil
// 	}

// 	return &users[0]

// }
