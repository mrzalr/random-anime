package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("visiting %s\n", r.URL.String())
	})
	GetAnime(c)
}
