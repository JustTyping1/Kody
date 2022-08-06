package main

import (
	"github.com/gocolly/colly"
	//"strconv"
)

func main() {
	c := colly.NewCollector()
	c.OnHTML("div[aria-live=polite]", func(h *colly.HTMLElement) {
		println(h.ChildText("div[data-val]"))

	})
	c.Visit("https://www.bing.com/search?q=weather")

}
