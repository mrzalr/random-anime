package main

import (
	"log"
	"net/http"

	"github.com/gocolly/colly"
)

type res map[string]string

func main() {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		log.Printf("visiting %s\n", r.URL.String())
	})

	mux := http.NewServeMux()
	mux.Handle("/", JSON(HandlerHome()))
	mux.Handle("/anime", JSON(HandlerAnime(c)))

	log.Println("server is running on port :8080")
	http.ListenAndServe(":8080", mux)
}
