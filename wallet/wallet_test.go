package wallet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeposit(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(BTC(12.5))

	want := BTC(12.5)
	got := wallet.Balance()

	assert.Equal(t, want, got)
}
