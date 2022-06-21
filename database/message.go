package database

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Message struct {
	gorm.Model
	RoomTicker string
	Message    string
	Username   string
	Time       time.Time
}

func AddToMessage(ticker string, message string, name string, time time.Time) {

	newMessage := Message{}
	newMessage.RoomTicker = ticker
	newMessage.Message = message
	newMessage.Username = name
	newMessage.Time = time

	Db.Create(&newMessage)
}

func FindMessageByUsernameTime(username string, time time.Time) *Message {

	var messages []Message

	message := Message{}
	message.Username = username
	message.Time = time

	Db.Where(message).First(&messages)

	if len(messages) == 0 {
		return nil
	}
	return &messages[0]
}

func FindMessagesByTicker(ticker string) *[]Message {

	var messages []Message

	Db.Find(&messages, Message{RoomTicker: ticker})

	// result := Db.Where(tickerMessage).Find(&message)

	// Db.Find(&message, Message{RoomTicker: ticker})

	return &messages

}
