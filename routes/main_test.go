package routes

import (
	"testing"
)

func TestTickerListHandler(t *testing.T) {
	t.Skip("Skipping test ticker handler")

	// type tickerListResponseBody struct {
	// 	Tickers []string
	// }

	// req := httptest.NewRequest(http.MethodGet, "/", nil)
	// w := httptest.NewRecorder()
	// GetAllTickers(w, req)
	// res := w.Result()
	// defer res.Body.Close()

	// wantResponseType := "application/json"

	// gotResponseType := res.Header.Get("Content-Type")

	// if wantResponseType != gotResponseType {
	// 	t.Fatal("Wrong Response Type")
	// }

	// var bodyBytes []byte
	// tlRB := &tickerListResponseBody{}

	// res.Body.Read(bodyBytes)

	// json.Unmarshal(bodyBytes, tlRB)

	// notWantedTicketCount := 0
	// gotTickerCount := len(tlRB.Tickers)

	// if notWantedTicketCount == gotTickerCount {
	// 	t.Fatal("ticker Count Empty")
	// }
	//////////
	// {
	//     "tickers" : ["AAPL", "MOON"]
	// }

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	//     t.Fatal("expected error to be nil got %v", err)
	// }
	// if string(data) != "sdgadgasdgasdgasdg" {
	//     t.Fatal("expected ABC got %v", string(data))
	// }
}
