package main

import (
	"fmt"
	"net/http"
	"time"
)

type result struct {
	url    string
	status string
}

func main() {
	//results := make(map[string]string)
	c := make(chan result)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}
}

func hitURL(url string, c chan<- result) {
	fmt.Println("checking:", url)

	resp, err := http.Get(url)
	status := "OK"

	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}

	c <- result{url: url, status: status}
}

func isCheck(person string, c chan string) {
	time.Sleep(time.Second * 2)
	c <- person + "check"
}
