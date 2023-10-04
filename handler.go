package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

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

type ResponseWriterWithCode struct {
	rw   http.ResponseWriter
	code int
}

func (m_rw *ResponseWriterWithCode) Header() http.Header {
	return m_rw.rw.Header()
}

func (m_rw *ResponseWriterWithCode) Write(b []byte) (int, error) {
	return m_rw.rw.Write(b)
}

func (m_rw *ResponseWriterWithCode) WriteHeader(statusCode int) {
	m_rw.code = statusCode
	m_rw.rw.WriteHeader(statusCode)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			rw := &ResponseWriterWithCode{rw: w}
			start := time.Now()

			next.ServeHTTP(rw, r)

			log.Printf("%s | %s | %d | %s", r.Method, r.URL, rw.code, time.Since(start))
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
			response, err := GetAnime(c)
			if err != nil {
				w.WriteHeader(http.StatusBadGateway)
				enc.Encode(res{"message": err.Error()})
				return
			}
			w.WriteHeader(http.StatusOK)
			enc.Encode(response)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			enc.Encode(res{"message": "method not allowed"})
		}
	}
}
