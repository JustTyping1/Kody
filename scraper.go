package main

import (
	"github.com/gocolly/colly"
)

type reading struct {
	Temp string `json:"temp"`
}

func main() {
	c := colly.NewCollector()
	c.OnHTML("span[class=wr-value--temperature--c]", func(h *colly.HTMLElement) {
		println(h.ChildText("span[aria-hidden=true]"))
	})
	c.Visit("https://www.bbc.co.uk/weather/2649452")

}
