package database

import (
	"log"
	"math/rand"
	"proj6/gomoon/config"
	"testing"
	"time"
)

func TestMain(m *testing.M) {

	rand.Seed(time.Now().UnixNano())

	globalConfig := config.ReadConfig(config.Test)

	Init(&globalConfig.PSQL)

	if Db == nil {
		log.Fatal("db is zero")
	}

	m.Run()
}
