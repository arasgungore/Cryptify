package wallet

import (
	"fmt"
	"sync"
)

// Wallet represents a user's cryptocurrency wallet
type Wallet struct {
	balances     map[string]float64
	balancesLock sync.Mutex
}

// NewWallet creates a new instance of Wallet
func NewWallet() *Wallet {
	return &Wallet{
		balances: make(map[string]float64),
	}
}

// GetBalance returns the balance for a specific cryptocurrency
func (w *Wallet) GetBalance(currency string) float64 {
	w.balancesLock.Lock()
	defer w.balancesLock.Unlock()

	return w.balances[currency]
}

// AddToBalance adds funds to the user's wallet
func (w *Wallet) AddToBalance(currency string, amount float64) {
	w.balancesLock.Lock()
	defer w.balancesLock.Unlock()

	w.balances[currency] += amount
	fmt.Printf("Added %f %s to the wallet\n", amount, currency)
}

// SubtractFromBalance subtracts funds from the user's wallet
func (w *Wallet) SubtractFromBalance(currency string, amount float64) {
	w.balancesLock.Lock()
	defer w.balancesLock.Unlock()

	if w.balances[currency] >= amount {
		w.balances[currency] -= amount
		fmt.Printf("Subtracted %f %s from the wallet\n", amount, currency)
	} else {
		fmt.Printf("Insufficient funds in the wallet\n")
	}
}
