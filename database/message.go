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

func NewMessage(ticker string, message string, name string, time time.Time) *Message {
	msg := Message{}
	msg.RoomTicker = ticker
	msg.Message = message
	msg.Username = name
	msg.Time = time

	return &msg
}

func AddToMessageFromValues(ticker string, message string, name string, time time.Time) {
	newMessage := NewMessage(ticker, message, name, time)

	Db.Create(&newMessage)
}

func AddToMessage(msg *Message) {
	Db.Create(msg)
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

func FindMessagesInDB() *[]Message {

	var messages []Message

	Db.Limit(10).Find(&messages)

	// result := Db.Where(tickerMessage).Find(&message)

	// Db.Find(&message, Message{RoomTicker: ticker})

	return &messages

}
