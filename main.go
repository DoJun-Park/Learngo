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

	// Search
	// dictionary := mydict.Dictionary{"first","first word"}
	// definition, err := dictionary.Search("first")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	// Add
	// dictionary := mydict.Dictionary{}
	// word := "hello"
	// definition := "greeting"
	// err := dictionary.Add(word, definition)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// hello, _ := dictionary.Search(word)
	// fmt.Println(hello)
	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }
	// //------------------------------------------------------

	// Update
	// dictionary := mydict.Dictionary{}
	// baseWord := "hello"
	// dictionary.Add(baseWord, "Firsts")
	// err := dictionary.Update(baseWord, "Second")

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// word, _ := dictionary.Search(baseWord)
	// fmt.Println(word)

	// Delete
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
