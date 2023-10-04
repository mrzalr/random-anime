package main

import (
	"encoding/json"
	"net/http"

	"github.com/gocolly/colly"
)

func JSON(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		},
	)
}

func HandlerHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(true)

		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
			enc.Encode(res{
				"message": "Welcome to random anime suggestion API !",
			})
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			enc.Encode(res{"message": "method not allowed"})
		}
	}
}

func HandlerAnime(c *colly.Collector) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(true)

		switch r.Method {
		case http.MethodGet:
			response := GetAnime(c)
			w.WriteHeader(http.StatusOK)
			enc.Encode(response)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			enc.Encode(res{"message": "method not allowed"})
		}
	}
}
