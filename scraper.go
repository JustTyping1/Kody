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
			})
			a.Visit("https://www.cnbc.com/quotes/.DJI?qsearchterm=Dow%20Jones%20Industrial%20Average")

			b := colly.NewCollector()
			b.OnHTML("div.QuoteStrip-lastPriceStripContainer", func(h *colly.HTMLElement) {
				println("S&P500 :")
				println(h.ChildText("span.QuoteStrip-lastPrice"))
			})
			b.Visit("https://www.cnbc.com/quotes/.SPX?qsearchterm=S&P")

			d := colly.NewCollector()
			d.OnHTML("div.QuoteStrip-lastPriceStripContainer", func(h *colly.HTMLElement) {
				println("NASDAQ :")
				println(h.ChildText("span.QuoteStrip-lastPrice"))
			})
			d.Visit("https://www.cnbc.com/quotes/.IXIC?qsearchterm=nasdaq")
		}
		if command == "n" {
			e := colly.NewCollector()
			e.OnHTML("div.story-content", func(h *colly.HTMLElement) {
				println(h.ChildText("a"))
			})
			e.Visit("https://www.reuters.com/news/archive/worldNews?date=today")
		}
		if command == "e" {
			running = false
		}
	}
}
