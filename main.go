package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gocolly/colly"
)

type res map[string]string

func main() {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		log.Printf("visiting %s\n", r.URL.String())
	})

	mux := http.NewServeMux()
	mux.Handle("/", CORS(Logger(JSON(HandlerHome()))))
	mux.Handle("/anime", CORS(Logger(JSON(HandlerAnime(c)))))

	port := os.Getenv("PORT")
	log.Printf("server is running on port :%s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}
