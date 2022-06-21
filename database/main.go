package database

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"proj6/gomoon/config"
)

var Db *gorm.DB

var err error

func Init(dbConfig *config.PSQL) {

	fmt.Println("Initializing database")

	host := dbConfig.Host
	port := dbConfig.Port
	username := dbConfig.Username
	dbName := dbConfig.DatabaseName

	Db, err = gorm.Open("postgres", "host="+host+" port="+port+" user="+username+" dbname="+dbName+" sslmode=disable")

	if err != nil {
		panic("failed to connect database with configuration: " + "host=" + host + " port=" + port + " user=" + username + " dbname=" + dbName + " sslmode=disable")
	}

	Db.AutoMigrate(&User{})
}
