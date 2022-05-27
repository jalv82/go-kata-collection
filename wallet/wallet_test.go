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

func Test(t *testing.T) {
	tests := []struct {
		description string
		wallet      Wallet
		amount      BTC
	}{
		{description: "Should deposit 22.0 BTC on the wallet", wallet: Wallet{}, amount: 22.0},
		{description: "Should deposit 1.11 BTC on the wallet", wallet: Wallet{}, amount: 1.11},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			test.wallet.Deposit(test.amount)
			got := test.wallet.Balance()

			assert.Equal(t, test.amount, got, test.description)
		})
	}
}
