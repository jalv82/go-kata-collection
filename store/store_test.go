package store

import "testing"

func TestBuy(t *testing.T) {
	tests := []struct {
		description string
		record      Record
		copiesToBuy uint
		want        uint
	}{
		{
			description: "Should be nine copies left  when exists ten copies and ones is bought",
			record:      Record{Title: "The Jar Dance", Artist: "Chiquito de la calzada", Copies: 10},
			copiesToBuy: 1,
			want:        9,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			test.record.Buy(test.copiesToBuy)

			got := test.record.Copies

			if got != test.want {
				t.Errorf("Got %v != Want %v", got, test.want)
			}
		})
	}
}
