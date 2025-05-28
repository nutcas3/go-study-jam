package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	Owner  string
	Balance float64
	mu     sync.Mutex
}

func (a *BankAccount) Deposit(amount float64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance += amount
}

func (a *BankAccount) Withdraw(amount float64) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.Balance >= amount {
		a.Balance -= amount
		return true
	}
	return false
}

func main() {
	acc := &BankAccount{Owner: "Alice", Balance: 1000}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		acc.Deposit(500)
		fmt.Println("Deposited 500")
	}()
	go func() {
		defer wg.Done()
		ok := acc.Withdraw(200)
		if ok {
			fmt.Println("Withdrew 200")
		} else {
			fmt.Println("Withdraw failed")
		}
	}()
	wg.Wait()
	fmt.Printf("Final balance for %s: %.2f\n", acc.Owner, acc.Balance)
}
