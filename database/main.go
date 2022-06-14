package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

var err error

func Init(name string) {

	fmt.Println("Initializing database")
	Db, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=kaichungyeo dbname="+name+" sslmode=disable")

	if err != nil {
		panic("failed to connect database")
	}

	Db.AutoMigrate(&User{})
}
