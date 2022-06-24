package database

import (
	"proj6/gomoon/utils"
	"testing"
	"time"
)

func TestCreateAndGetOneMessage(t *testing.T) {

	testMessage := "message" + utils.RandomString(12)
	username := "user" + utils.RandomString(12)
	timee := time.Now()
	ticker := "ticker" + utils.RandomString(12)

	AddToMessageFromValues(ticker, testMessage, username, timee)

	newMessage := FindMessageByUsernameTime(username, timee)

	if newMessage == nil {
		t.Fatal("message not created")
	}

	if newMessage.Username != username {
		t.Fatal("created message, but different username")
	}

	gotButRoundedTime := timee.Round(time.Millisecond)
	actualButRoundedTime := newMessage.Time.Round(time.Millisecond)

	if actualButRoundedTime.GoString() == gotButRoundedTime.GoString() {
		t.Log(actualButRoundedTime.Unix())
		t.Log(gotButRoundedTime.Unix())
		t.Fatal("created message, but different time as compared to offset from GoString")
	}

}

func TestGetMessagesByTickers(t *testing.T) {

	username := "user" + utils.RandomString(12)
	ticker1 := "ticker" + utils.RandomString(12)
	ticker2 := "ticker" + utils.RandomString(12)

	tickers := []string{
		ticker1, ticker2,
	}

	messageCount := 10

	for i := 0; i < messageCount; i++ {
		testMessage := "message" + utils.RandomString(12)

		ticker := tickers[i%2]
		timee := time.Now()

		AddToMessageFromValues(ticker, testMessage, username, timee)

	}

	AddToMessageFromValues(ticker2, "feeder message", username, time.Now())

	allTicker1Messages := FindMessagesByTicker(ticker1)
	allTicker2Messages := FindMessagesByTicker(ticker2)

	if allTicker1Messages == nil {
		t.Fatal("nessage array is empty")
	}

	if len(*allTicker1Messages) != 5 {
		t.Fatal("message array length not tally")
	}

	if allTicker2Messages == nil {
		t.Fatal("nessage array is empty")
	}

	if len(*allTicker2Messages) != 6 {
		t.Fatal("message array length not tally")
	}

}
