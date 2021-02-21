package main

import (
	"fmt"

	"github.com/Learngo/accounts"
)

func main() {
	account := accounts.NewAccount("dojun")
	account.Deposit(10)

	// err := account.Withdraw(20)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)
}
