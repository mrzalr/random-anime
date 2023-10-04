package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("visiting %s\n", r.URL.String())
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := map[string]string{
			"message": "Welcome to random anime suggestion API ! &&",
		}

		w.WriteHeader(http.StatusOK)

		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(true)
		enc.Encode(response)
	})

	http.HandleFunc("/anime", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		anime := GetAnime(c)

		w.WriteHeader(http.StatusOK)

		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(true)
		enc.Encode(anime)

	})

	http.ListenAndServe(":8080", nil)
}
