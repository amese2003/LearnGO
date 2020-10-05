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

func canDrink(age int) bool {
	// if koreanAge := age + 2; koreanAge < 18 {
	// 	return false
	// }

	switch age {
	case 10:
		return false
	case 18:
		return true
	}

	return true
}

func main() {
	a := 2
	b := &a

	*b = 5

	fmt.Println(a, *b)

}
