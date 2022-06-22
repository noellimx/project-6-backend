package database

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"proj6/gomoon/config"
)

var Db *gorm.DB

var err error

func autoMigrate() {

	if Db == nil {

		panic("This method should only be executed after initializing the global db instance")

	}
	Db.AutoMigrate(&User{})
	Db.AutoMigrate(&Message{})

}

func Init(dbConfig *config.PSQL) {

	fmt.Println("Initializing database")

	host := dbConfig.Host
	port := dbConfig.Port
	username := dbConfig.Username
	dbName := dbConfig.DatabaseName
	password := dbConfig.Password

	if password == "" {
		fmt.Println("Info: No password is supplied for the database.")
	} else {
		fmt.Println("Info: Password is supplied to the database")

	}

	Db, err = gorm.Open("postgres", "host="+host+" port="+port+" user="+username+" password="+password+" dbname="+dbName+" sslmode=disable")

	if err != nil {

		defer fmt.Println(err)
		panic("failed to connect database with configuration: " + "host=" + host + " port=" + port + " user=" + username + " dbname=" + dbName + " sslmode=disable")
	} else {
		fmt.Println("dbName: " + dbName)
	}
	autoMigrate()

}
