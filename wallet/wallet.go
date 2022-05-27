package wallet

type BTC float64

type Wallet struct {
	balance BTC
}

func (w *Wallet) Deposit(amount BTC) {
	w.balance += amount
}

func (w Wallet) Balance() BTC {
	return w.balance
}
