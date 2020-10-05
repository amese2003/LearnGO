package main

import (
	"fmt"
)

func main() {
	names := map[string]string{"name": "asd", "age": "12"}
	//fmt.Println(names)

	for key, _ := range names {
		fmt.Println(key)
	}
}
