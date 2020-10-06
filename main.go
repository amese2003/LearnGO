package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	people := [2]string{"nico", "flynn"}

	for _, person := range people {
		go isCheck(person, c)
	}

	fmt.Println(<-c)
	fmt.Println(<-c)
}

func isCheck(person string, c chan bool) {
	time.Sleep(time.Second * 2)
	fmt.Println(person)
	c <- true
}
