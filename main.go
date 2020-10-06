package main

import (
	"fmt"

	"github.com/amese1225/LearnGO/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}

	def, err := dictionary.Search("second")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(def)
	}
}
