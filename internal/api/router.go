package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	Router *chi.Mux
	Server *http.Server
}

func NewHandler() *Handler {
	r := &Handler{}

	r.Router = chi.NewRouter()
	r.mapRoutes()

	r.Server = &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("APP_PORT")),
		Handler: r.Router,
	}

	return r
}

func (h *Handler) mapRoutes() {
	h.Router.Use(middleware.Logger)

	// v1 routes
	h.Router.Route("/v1", func(r chi.Router) {
		r.Group(buyerRoutes)
		r.Group(sellerRoutes)
	})

	h.Router.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("health check ok"))
	})

	h.Router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("route does not exist"))
	})
}

func (h *Handler) Serve() error {
	go func() {
		log.Println("starting web server on port:", os.Getenv("APP_PORT"))
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	h.Server.Shutdown(ctx)

	log.Println("shut down gracefully")

	return nil
}
