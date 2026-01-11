package middleware

import (
	"log"
	"net/http"
)

func MdlwrTwo(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("res from mdlwr two")
		next.ServeHTTP(w, r)
	})
}
