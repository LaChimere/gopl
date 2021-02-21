// Package bank provides a concurrency-safe bank with one account.
package bank

import "sync"

var (
	mutex   sync.Mutex // guards balance
	rwmutex sync.RWMutex
	balance int
)

func Deposit(amount int) {
	mutex.Lock()
	defer mutex.Unlock()
	deposit(amount)
}

// deposit requires that the lock be held.
func deposit(amount int) {
	balance += amount
}

func Balance() int {
	rwmutex.RLock() // readers lock
	defer rwmutex.RUnlock()
	return balance
}

func Withdraw(amount int) bool {
	mutex.Lock()
	defer mutex.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false // insufficient funds
	}
	return true
}
