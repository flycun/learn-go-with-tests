package main

import "fmt"

// Bitcoin represents a number of Bitcoins.
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet stores the number of Bitcoin someone owns.
type Wallet struct {
	balance int
}

// Deposit will add some Bitcoin to a wallet.
func (w Wallet) Deposit(amount int) {
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance+=amount
}

// Balance returns the number of Bitcoin a wallet has.
func (w Wallet) Balance() int {
	return w.balance
}