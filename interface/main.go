package main

import (
	"fmt"
	"learning-golang/interface/accounts"
)

//Interfaces are a great way to enforce a blueprint for what your app can do,
// similar to object oriented programming. For example,
// there are many different types of bank accounts,
// but you can use a single interface to interact with all of them:
// GetBalance(), Deposit() and Withdraw().
// This makes your high-level code easier to work with because the implementation is abstracted away.


type BankAccount interface {
	GetBalance() int 
	Deposit(amount int)
	Withdraw(amount int) error
}

func main() {
	// they implement the interface then their are type of BanckAccount too.
	myAccounts := []BankAccount{
		accounts.NewBtcAccount(),
		accounts.NewWellsFargo(),
	}

	for _, account := range myAccounts{
		account.Deposit(500)
		if err := account.Withdraw(100); err != nil {
			fmt.Printf("%d\n", err)
		}

		balance := account.GetBalance()
		fmt.Printf("balance = %d\n", balance)
	}

}