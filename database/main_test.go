package database

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestMain(m *testing.M) {

	rand.Seed(time.Now().UnixNano())

	Init("dbmoontest")

	if Db == nil {
		log.Fatal("db is zero")
	}

	m.Run()
}
