package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func sellerRoutes(r chi.Router) {
	r.Use(LogSeller)
	r.Get("/seller", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("seller"))
	})
}

func LogSeller(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("seller hit")
		next.ServeHTTP(w, r)
	})
}
