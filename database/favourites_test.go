package database

import (
	"testing"
)

func TestAFavourite(t *testing.T) {

	email := "2@2.com"
	ticker := "NASDAQ:AAPL"

	gotFav := AFavourite(email, ticker)

	if gotFav.Email != email || gotFav.Ticker != ticker {

		t.Fatal("Type def error: Favourite")
	}

}
