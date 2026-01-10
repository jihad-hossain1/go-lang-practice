package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Println("ami middleware: age print hobo")

		next.ServeHTTP(w, r)

		log.Println("ami middleware: pore print hbo")
		diff := time.Since(start)

		log.Println(r.Method, r.URL.Path, diff)
	})
}
