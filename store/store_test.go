package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuy(t *testing.T) {
	tests := []struct {
		description string
		record      Record
		copiesToBuy uint
		want        uint
	}{
		{
			description: "Should be nine copies left when exists ten copies and ones is bought",
			record:      Record{Title: "The Jar Dance", Artist: "Chiquito de la calzada", Copies: 10},
			copiesToBuy: 1,
			want:        9,
		},
		{
			description: "Should return an error when exists zero copies and ones is bought",
			record:      Record{Title: "The Jar Dance", Artist: "Chiquito de la calzada", Copies: 0},
			copiesToBuy: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			err := test.record.Buy(test.copiesToBuy)
			if err != nil {
				assert.Error(t, err)
			}

			got := test.record.Copies

			assert.Equal(t, test.want, got)
		})
	}
}
