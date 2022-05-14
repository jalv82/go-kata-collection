package generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testsTable[T MyNumber] struct {
	description string
	numbers     []T
	want        T
}

func TestSum(t *testing.T) {

	tests := []testsTable[int]{
		{description: "Should sum two numbers of type integer", numbers: []int{1, 2}, want: 3},
		{description: "Should sum four numbers of type integer", numbers: []int{1, 2, 3, 5}, want: 11},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got := Sum[int](test.numbers[:]...)

			assert.Equal(t, test.want, got, test.description)
		})
	}
}
