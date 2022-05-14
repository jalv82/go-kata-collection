package generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		description string
		numbers     []int
		want        int
	}{
		{description: "Should sum two numbers of type integer", numbers: []int{1, 2}, want: 3},
		{description: "Should sum four numbers of type integer", numbers: []int{1, 2, 3, 5}, want: 11},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got := Sum(test.numbers[:]...)

			assert.Equal(t, test.want, got, test.description)
		})
	}
}
