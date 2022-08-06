package main

import (
	"fmt"

	"github.com/gocolly/colly"
	//"strconv"
)

func main() {
	var running bool = true
	for running {
		println("ENTER CMD >")
		var command string
		fmt.Scanln(&command)
		if command == "w" {
			c := colly.NewCollector()
			c.OnHTML("div.main-temp", func(h *colly.HTMLElement) {
				println(h.ChildText("span.currTemp") + " Degrees Celcius")

			})
			c.Visit("https://uk.search.yahoo.com/search?p=weather")
		}
		if command == "s" {
			a := colly.NewCollector()
			a.OnHTML("div.QuoteStrip-lastPriceStripContainer", func(h *colly.HTMLElement) {
				println("DJI :")
				println(h.ChildText("span.QuoteStrip-lastPrice"))
				println(h.ChildText("span.QuoteStrip-changeUp"))
			})
			a.Visit("https://www.cnbc.com/quotes/.DJI?qsearchterm=Dow%20Jones%20Industrial%20Average")

			b := colly.NewCollector()
			b.OnHTML("div.QuoteStrip-lastPriceStripContainer", func(h *colly.HTMLElement) {
				println("DJI :")
				println(h.ChildText("span.QuoteStrip-lastPrice"))
				println(h.ChildText("span.QuoteStrip-changeUp"))
			})
			b.Visit("https://www.cnbc.com/quotes/.DJI?qsearchterm=Dow%20Jones%20Industrial%20Average")
		}
	}
}
