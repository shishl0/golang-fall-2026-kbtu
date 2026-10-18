package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	Owner   string
	balance float64
}

func (a *BankAccount) Deposit(amount float64) {
	a.balance += amount
}

func (a *BankAccount) Withdraw(amount float64) error {
	if amount > a.balance {
		return errors.New("not enough money")
	}
	a.balance -= amount
	return nil
}

func (a *BankAccount) Balance() float64 {
	return a.balance
}

func task4() {
	accounts := []BankAccount{{Owner: "Ali"}, {Owner: "Bota"}, {Owner: "Dias"}}
	accounts[0].Deposit(100)
	accounts[1].Deposit(200)
	accounts[2].Deposit(300)

	for i := range accounts {
		accounts[i].Deposit(accounts[i].Balance() * 0.05)
	}

	for i := range accounts {
		fmt.Println(accounts[i].Owner, accounts[i].Balance())
	}

	err := accounts[0].Withdraw(1000)
	if err != nil {
		fmt.Println(err)
	}
}
