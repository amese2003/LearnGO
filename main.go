package main

import (
	"fmt"

	"github.com/amese1225/LearnGO/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}

	baseword := "hello"

	dictionary.Add(baseword, "First")
	err := dictionary.Update(baseword, "Second")

	if err != nil {
		fmt.Println(err)
	}

	word, _ := dictionary.Search(baseword)
	fmt.Println(word)

	err2 := dictionary.Delete(baseword)

	if err2 != nil {
		fmt.Println(err2)
	}

}
