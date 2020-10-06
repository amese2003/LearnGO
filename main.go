package main

import (
	accounts "Bank_dictionary/Accounts"
	"fmt"
	"log"
)

func main() {
	account := accounts.NewAccount("tester")
	account.Deposit(5)
	fmt.Println(account.Balance())
	err := account.WithDraw(6)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(account.Balance())
}
