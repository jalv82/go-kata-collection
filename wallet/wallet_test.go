package wallet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeposit(t *testing.T) {
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
			err := test.wallet.Deposit(test.amount)
			got := test.wallet.Balance()

			assert.Nil(t, err)
			assert.Equal(t, test.amount, got, test.description)
		})
	}
}

func TestForbiddenNegativeOrZeroAmountDeposit(t *testing.T) {
	tests := []struct {
		description string
		wallet      Wallet
		amount      BTC
	}{
		{description: "Should return an error when when try deposit -22.0 BTC on the wallet", wallet: Wallet{}, amount: -22.0},
		{description: "Should return an error when when try deposit 0.0 BTC on the wallet", wallet: Wallet{}, amount: 0.0},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			want := test.wallet.Balance()
			err := test.wallet.Deposit(test.amount)
			got := test.wallet.Balance()

			assert.NotNil(t, err)
			assert.Equal(t, want, got, test.description)
		})
	}
}
