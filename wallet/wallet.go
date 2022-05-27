package wallet

import "errors"

type BTC float64

type Wallet struct {
	balance BTC
}

var ErrorForbiddenDeposit = errors.New("forbidden deposit negative or zero amount")

func (w *Wallet) Deposit(amount BTC) error {
	if amount <= 0 {
		return ErrorForbiddenDeposit
	}

	w.balance += amount

	return nil
}

func (w Wallet) Balance() BTC {
	return w.balance
}
