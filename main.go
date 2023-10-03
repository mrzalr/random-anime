package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

var (
	baseUrl = "https://myanimelist.net/topanime.php"
)

func main() {
	limit := GetRandomPage()
	animeNum := GetRandomNumber(1, 50)

	c := colly.NewCollector()

	counter := 0
	c.OnHTML("a[href].hoverinfo_trigger", func(h *colly.HTMLElement) {
		class := strings.Split(h.Attr("class"), " ")
		if len(class) > 1 || class[0] != "hoverinfo_trigger" {
			return
		}

		counter++
		if counter != animeNum {
			return
		}

		href := h.Attr("href")
		fmt.Printf("%d | %s | %s\n", animeNum, strings.TrimSpace(h.Text), href)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("visiting %s\n", r.URL.String())
	})

	if limit != 0 {
		baseUrl = fmt.Sprintf("%s?limit=%d", baseUrl, limit)
	}
	c.Visit(baseUrl)
}
