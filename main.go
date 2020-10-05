package main

import (
	"fmt"
	"strings"
)

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")

	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatString(name ...string) {
	fmt.Println(name)
}

func main() {
	totalLength, _ := lenAndUpper("nico")
	fmt.Println(totalLength)

	repeatString("A", "BB", "CCC")
}
