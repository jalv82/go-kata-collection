package store

import "testing"

func TestBuy(t *testing.T) {

	record := Record{
		Title:  "The Jar Dance",
		Artist: "Chiquito de la calzada",
		Copies: 10,
	}

	record.Buy(1)

	got := record.Copies
	var want uint = 9

	if got != want {
		t.Errorf("Got %v != Want %v", got, want)
	}
}
