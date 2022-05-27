package main

import (
	"log"
	"wallet"
)

func main() {
	log.Println("# Bitcoin wallet")

	btcWallet := wallet.Wallet{}
	log.Printf("Wallet's balance: %.2f", btcWallet.Balance())

	err := btcWallet.Withdraw(112.3)
	if err != nil {
		log.Println(err.Error())
	}

	err = btcWallet.Deposit(114.0)
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("Wallet's balance: %.2f", btcWallet.Balance())

	err = btcWallet.Withdraw(112.3)
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("Wallet's balance: %.2f", btcWallet.Balance())
}
