package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/amese1225/LearnGO/scrapper"
	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
}

func StartScrapperServer() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":4000"))
}

func test(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

func main() {

	go test("Async")

	var wait sync.WaitGroup
	wait.Add(1)

	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("Test!")
		defer wait.Done()
	}()

	wait.Wait()
}
