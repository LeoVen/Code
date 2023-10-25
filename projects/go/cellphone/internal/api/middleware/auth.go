package middleware

import (
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// auth := r.Header.Get("Authorization")

		// TODO
		log.Default().Println("Authentication passed")

		next.ServeHTTP(w, r)
	})
}
