package main

import (
	"github.com/gocolly/colly"
)

type reading struct {
	Temp string `json:"temp"`
}

func main() {
	c := colly.NewCollector()
	c.OnHTML("div[class = vk_bk TylWce SGNhVe]", func(h *colly.HTMLElement) {
		println(h.Text)
	})
	c.Visit("https://www.google.com/search?q=weather+in+patchway&oq=weather+in+patchway&aqs=chrome..69i64j0i22i30l5j0i15i22i30j0i390l3.9585j1j7&sourceid=chrome&ie=UTF-8")

}
