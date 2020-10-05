package main

import (
	"fmt"
	"strings"
)

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatString(name ...string) {
	fmt.Println(name)
}

func main() {
	totalLength, _ := lenAndUpper("nico")
	fmt.Println(totalLength)

	repeatString("A", "BB", "CCC")
}
