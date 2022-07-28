package main

import (
	"bufio"
	"os"

	"github.com/gocolly/colly"
	//"strconv"
)

func main() {
	println("ENTER CMD >")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	c := colly.NewCollector()
	c.OnHTML("span[class=wr-value--temperature--c]", func(h *colly.HTMLElement) {
		println(h.ChildText("span[aria-hidden=true]"))
	})
	c.Visit("")

}
