package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func buyerRoutes(r chi.Router) {
	r.Use(LogBuyer)
	r.Get("/buyer", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("buyer"))
	})
}

func LogBuyer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("buyer hit")
		next.ServeHTTP(w, r)
	})
}
