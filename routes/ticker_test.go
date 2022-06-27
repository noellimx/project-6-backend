package routes

import (
	"fmt"
	"testing"
)

func TestGetTickers(t *testing.T) {

	var val string = "GME"

	tickerResults, err := SearchTickers(val)

	if err != nil {

		t.Fatal("Error getting results")

	}

	wantCount := 2
	gotCount := len(*tickerResults)
	if gotCount != wantCount {
		t.Fatal("Error getting results. Third party should return " + fmt.Sprint(wantCount) + "values." + "Got " + fmt.Sprint(gotCount))
	}

}
