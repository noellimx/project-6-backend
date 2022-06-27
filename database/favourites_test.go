package database

import (
	"fmt"
	"proj6/gomoon/utils"
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

func TestAddAFavourite(t *testing.T) {
	email := "s" + utils.RandomString(6) + "@a.com"
	ticker1 := "NASDAQ:AAPL"
	ticker4 := "NASDAQ:AAPL"
	ticker2 := "NYSE:GME"
	ticker3 := "NYSE:DDD"

	for _, ticker := range []string{ticker4, ticker3, ticker1, ticker2} {
		f := AFavourite(email, ticker)
		SetFavourite(f)
	}

	favs := GetFavouritesOfEmail(email)
	gotCount := len(*favs)
	wantCount := 3
	if gotCount != wantCount {
		t.Fatal("Favs count wrong. got " + fmt.Sprint(gotCount) + " want " + fmt.Sprint(wantCount))

	}
	t.Log("got fav " + fmt.Sprint(gotCount))

	RemoveFavourite(&Favourite{
		Email: email, Ticker: ticker1,
	})

	favsAfterRemove := GetFavouritesOfEmail(email)
	gotCount = len(*favsAfterRemove)
	wantCount = 2

	if gotCount != wantCount {
		t.Fatal("Favs count wrong after remove. got " + fmt.Sprint(gotCount) + " want " + fmt.Sprint(wantCount))

	}
	t.Log("got fav " + fmt.Sprint(gotCount))

}
