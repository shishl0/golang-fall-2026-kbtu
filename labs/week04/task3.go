package main

import "fmt"

type Counter struct {
	n int
}

func (c Counter) Get() int {
	return c.n
}

func (c *Counter) Inc() {
	c.n++
}

func NewCounter() Counter {
	return Counter{}
}

type Account struct {
	Owner   string
	Balance float64
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

func task3() {
	var c Counter
	c.Inc()
	(&c).Inc()
	Counter{}.Get()

	m := map[string]Counter{"a": {}}
	m["a"].Get()

	s := []Counter{{}}
	s[0].Inc()

	pm := map[string]*Counter{"a": {}}
	pm["a"].Inc()

	fmt.Println(c.Get(), s[0].Get(), pm["a"].Get())

	accounts := []Account{{"Ali", 100}, {"Bota", 200}}
	for _, acc := range accounts {
		acc.Deposit(50)
	}
	fmt.Println(accounts)

	for i := range accounts {
		accounts[i].Deposit(50)
	}
	fmt.Println(accounts)
}
