package generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	got := Sum(2, 3)

	want := 5

	assert.Equal(t, want, got, "Should sum two number of type integer")
}
