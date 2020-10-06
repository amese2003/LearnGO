package main

import (
	"fmt"

	"github.com/amese1225/LearnGO/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	def := "Greething"

	err := dictionary.Add(word, def)

	if err != nil {
		fmt.Println(err)
	}

	hello, _ := dictionary.Search(word)
	fmt.Println(hello)

	err2 := dictionary.Add(word, def)

	if err2 != nil {
		fmt.Println(err2)
	}

}
