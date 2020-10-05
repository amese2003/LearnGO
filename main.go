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

func superAdd(numbers ...int) int {
	var sum int = 0

	// for index, number := range numbers {
	// 	fmt.Println(index, number)
	// 	sum += number
	// }

	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	return sum
}

func main() {
	total := superAdd(1, 2, 3, 4, 5, 6, 7)

	var test int = superAdd(total)

	fmt.Println(test)
}
