package main

import (
	accounts "Bank_dictionary/Accounts"
	"fmt"
)

func main() {
	account := accounts.NewAccount("tester")
	fmt.Println(account)
}
