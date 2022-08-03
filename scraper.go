package main

import (
	"github.com/gocolly/colly"
	//"strconv"
)

func main() {
	c := colly.NewCollector()
	c.OnHTML("div[vk_bk TylWce SGNhVe]", func(h *colly.HTMLElement) {
		println(h.ChildText("span[wob_tm]"))

	})
	c.Visit("https://www.google.com/search?q=weather")

}
