package main

import (
	"bufio"

	"github.com/gocolly/colly"
	//"strconv"
)

func main() {
	scanner = bufio.NewScanner(os)

	c := colly.NewCollector()
	c.OnHTML("span[class=wr-value--temperature--c]", func(h *colly.HTMLElement) {
		println(h.ChildText("span[aria-hidden=true]"))
	})
	c.Visit("https://www.bbc.co.uk/weather/2649452")

}
