package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/amese1225/LearnGO/scrapper"
	"github.com/labstack/echo"

	queue "./src/queue"
	stack "./src/stack"
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

func main() {
	//var idMap map[int]string

	//idMap = make(map[int]string)

	// idMap := map[string]string{
	// 	"Tester": "task",
	// }

	// fmt.Println(idMap["Tester"])

	// var a []int

	// a = append(a, 1)
	// a = append(a, 2)

	// for i := 0; i < 2; i++ {
	// 	fmt.Println(a[i])
	// }

	a := queue.Queue{}

	a.Push(30)
	a.Push(40)

	fmt.Println(a.Size())

	fmt.Println(a.Peek())

	data, err := a.Pop()

	if err == nil {
		fmt.Println(data)
	}

	b := stack.Stack{}

	b.Push(10)
	b.Push(20)
	b.Push(300)

	stkData, stkerr := b.Top()

	if stkerr == nil {
		fmt.Println(stkData)
	} else {
		fmt.Println(stkerr)
	}

	for i := 0; i < 3; i++ {
		popData, popErr := b.Pop()

		if popErr == nil {
			fmt.Println(popData)
		}
	}

	fmt.Println(b.Size())
}
