package database

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"proj6/gomoon/config"
)

var Db *gorm.DB

var err error

func Init(name string, dbConfig *config.PSQL) {

	fmt.Println("Initializing database")
	Db, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=kaichungyeo dbname="+name+" sslmode=disable")

	if err != nil {
		panic("failed to connect database")
	}

	Db.AutoMigrate(&User{})
}
