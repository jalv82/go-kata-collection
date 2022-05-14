package generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	got := Sum(2, 3, 5)

	want := 10

	assert.Equal(t, want, got, "Should sum any quantity of numbers of type integer")
}
