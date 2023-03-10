package middlewares

import (
	"log"
	"net/http"
)

func PrintPath(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL
		log.Printf("You access %s %s", url.Host, url.Path)
		next.ServeHTTP(w, r)
	})
}
