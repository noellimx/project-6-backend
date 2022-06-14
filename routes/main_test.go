package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTickerListHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	GetAllTickers(w, req)
	res := w.Result()
	defer res.Body.Close()

	wantResponseType := "application/json"

	gotResponseType := res.Header.Get("Content-Type")

	if wantResponseType != gotResponseType {
		t.Error("Wrong Response Type")
	}
	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	//     t.Errorf("expected error to be nil got %v", err)
	// }
	// if string(data) != "sdgadgasdgasdgasdg" {
	//     t.Errorf("expected ABC got %v", string(data))
	// }
}
