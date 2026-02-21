// Package middleware serve request middleware
package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	middleware := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Println(r.Method, r.URL.Path, time.Since(start))
	}

	return http.HandlerFunc(middleware)
}

func Emoji(next http.Handler) http.Handler {
	middleware := func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		log.Println("...................")
	}

	return http.HandlerFunc(middleware)
}

func After(next http.Handler) http.Handler {
	middleware := func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		log.Println("-----did a lot's of works---------")
	}

	return http.HandlerFunc(middleware)
}
