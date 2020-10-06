package main

import (
	accounts "Bank_dictionary/Accounts"
	"fmt"
)

func main() {
	account := accounts.NewAccount("tester")
	account.Deposit(5)
	fmt.Println(account.Balance())

}
