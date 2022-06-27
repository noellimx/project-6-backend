package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetAllTickers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

type TickerSearchResponse struct {
	Symbol       string   `json:"symbol"`
	Description  string   `json:"description"`
	Type         string   `json:"type"`
	Exchange     string   `json:"exchange"`
	CurrencyCode string   `json:"currency_code"`
	Logoid       string   `json:"logoid,omitempty"`
	ProviderID   string   `json:"provider_id"`
	Country      string   `json:"country,omitempty"`
	Typespecs    []string `json:"typespecs"`
	Prefix       string   `json:"prefix,omitempty"`
}
type TickerSearchResponses = []TickerSearchResponse

type DadJokeReponse struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func SearchTickers(val string) (*TickerSearchResponses, error) {
	fmt.Println("Calling API...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://symbol-search.tradingview.com/symbol_search/?text="+val+"&hl=1&exchange=&lang=en&type=stock&domain=production", nil)
	if err != nil {
		return &TickerSearchResponses{}, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return &TickerSearchResponses{}, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &TickerSearchResponses{}, err
	}
	var tickerResults TickerSearchResponses
	json.Unmarshal(bodyBytes, &tickerResults)
	fmt.Printf("API Response as struct %+v\n", tickerResults)

	return &tickerResults, nil
}

func GetSearchValue(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)
	param := chi.URLParam(r, "val")

	// TODO
	result, _ := SearchTickers(param)

	json.NewEncoder(w).Encode(result)

}
func TickerRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/ticker/getallticker/{val}", GetSearchValue)

	return r
}
