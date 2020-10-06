package main

import (
	"fmt"
	"log"

	"github.com/amese1225/LearnGO/accounts"
)

func main() {
	account := accounts.NewAccount("tester")
	account.Deposit(10)
	fmt.Println(account)

	err := account.Withdraw(5)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(account)
}
