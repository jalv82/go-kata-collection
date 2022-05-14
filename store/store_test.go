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

func TestCalculateDiscount(t *testing.T) {
	tests := []struct {
		description        string
		record             Record
		discountPercentage float32
		want               float32
	}{
		{
			description:        "Should calculated the 10% of discount",
			record:             Record{Price: 25.0, IsDiscounted: true},
			discountPercentage: 10.0,
			want:               2.5,
		},
		{
			description:        "Should return an error when not is possible calculated some discount",
			record:             Record{Price: 25.0, IsDiscounted: false},
			discountPercentage: 10.0,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got, err := test.record.Discount(test.discountPercentage)
			if err != nil {
				assert.Error(t, err)
			}
			assert.Equal(t, test.want, got)
		})
	}
}
