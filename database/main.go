package database

import (
	"fmt"
	"net/url"

	"proj6/gomoon/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Db *gorm.DB

var err error

func autoMigrate() {

	if Db == nil {
		panic("This method should only be executed after initializing the global db instance")
	}
	fmt.Println("Auto-migrating...")
	Db.AutoMigrate(&User{}, &Message{}, &Favourite{})
}

func Init(dbConfig *config.PSQL) {

	fmt.Println("Initializing database")

	host := dbConfig.Host
	port := dbConfig.Port
	username := dbConfig.Username
	dbName := dbConfig.DatabaseName
	password := dbConfig.Password

	dsn := url.URL{
		User:     url.UserPassword(username, password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%s", host, port),
		Path:     dbName,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	dsnString := dsn.String()

	if password == "" {
		fmt.Println("Info: No password is supplied for the database.")
	} else {
		fmt.Println("Info: Password is supplied to the database")
	}

	connString := "host=" + host + " port=" + port + " user=" + username + " password=" + password + " dbname=" + dbName + " sslmode=disable"
	_ = connString

	fmt.Println("Connection String to DB : " + dsnString)
	Db, err = gorm.Open("postgres", dsnString)

	if err != nil {

		defer fmt.Println(err)
		panic("failed to connect database with configuration: " + "host=" + host + " port=" + port + " user=" + username + " dbname=" + dbName + " sslmode=disable")
	} else {
		fmt.Println("dbName: " + dbName)
	}

	autoMigrate()

}
