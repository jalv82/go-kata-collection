package wallet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDepositAmount(t *testing.T) {
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
		{description: "Should return an error when try deposit -22.0 BTC on the wallet", wallet: Wallet{}, amount: -22.0},
		{description: "Should return an error when try deposit 0.0 BTC on the wallet", wallet: Wallet{}, amount: 0.0},
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

func TestWithdrawAmount(t *testing.T) {
	tests := []struct {
		description   string
		wallet        Wallet
		amount        BTC
		balanceBefore BTC
		balanceAfter  BTC
	}{
		{description: "Should withdraw 22.0 BTC from the wallet than have 100.0 BTC of balance", wallet: Wallet{}, amount: 22.0, balanceBefore: 100.0, balanceAfter: 78.0},
		{description: "Should withdraw 1.11 BTC from the wallet than have 100.0 BTC of balance", wallet: Wallet{}, amount: 1.11, balanceBefore: 100.0, balanceAfter: 98.89},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			_ = test.wallet.Deposit(test.balanceBefore)

			want := test.balanceAfter
			err := test.wallet.Withdraw(test.amount)
			got := test.wallet.Balance()

			assert.Nil(t, err)
			assert.Equal(t, want, got, test.description)
		})
	}
}

func TestForbiddenWithdrawAmountGreaterThanBalance(t *testing.T) {
	tests := []struct {
		description   string
		wallet        Wallet
		amount        BTC
		balanceBefore BTC
		balanceAfter  BTC
	}{
		{description: "Should return an error when try withdraw 22.0 BTC from the wallet than have 10.0 BTC of balance", wallet: Wallet{}, amount: 22.0, balanceBefore: 10.0, balanceAfter: 10.0},
		{description: "Should return an error when try withdraw 1.11 BTC from the wallet than have 0.0 BTC of balance", wallet: Wallet{}, amount: 1.11, balanceBefore: 0.0, balanceAfter: 0.0},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			_ = test.wallet.Deposit(test.balanceBefore)

			want := test.balanceAfter
			err := test.wallet.Withdraw(test.amount)
			got := test.wallet.Balance()

			assert.NotNil(t, err)
			assert.Equal(t, want, got, test.description)
		})
	}
}
