package main

import (
	"fmt"

	"github.com/Learngo/mydict"
)

func main() {
	// //------------------------Bank------------------------
	// account := accounts.NewAccount("dojun")
	// account.Deposit(10)

	// // err := account.Withdraw(20)
	// // if err != nil {
	// // 	fmt.Println(err)
	// // }

	// // fmt.Println(account.Balance(), account.Owner())
	// fmt.Println(account)
	// //------------------------------------------------------

	// //------------------------Dictionary------------------------

	dictionary := mydict.Dictionary{"first": "first word"}
	definition, err := dictionary.Search("first")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
	// //------------------------------------------------------

}
