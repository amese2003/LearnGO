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

	// resultOne := <-c
	// resultTwo := <-c

	// fmt.Println("Waiting for Messages")
	// fmt.Println("Received this message:", resultOne)
	// fmt.Println("Received this message:", resultTwo)

	for i := 0; i < len(people); i++ {
		fmt.Println("waiting for", i)
		fmt.Println(<-c)
	}
}

func isCheck(person string, c chan string) {
	time.Sleep(time.Second * 2)
	c <- person + "check"
}
