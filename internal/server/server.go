package server

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewRouter creates and configures the chi router with all routes
func NewRouter(deps *Dependencies) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.StripSlashes)

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	router.Route("/numeric", func(r chi.Router) {
		r.Get("/{shortURL}", deps.NumericHandler.GetLongURL)
		r.Post("/", deps.NumericHandler.CreateShortURL)
	})

	router.Route("/random", func(r chi.Router) {
		r.Get("/test", deps.URLRandomHandler.Test)
	})

	return router
}

// Run starts the HTTP server
func Run(deps *Dependencies) {
	router := NewRouter(deps)

	addr := os.Getenv("PORT")
	if addr == "" {
		addr = ":8080"
	}
	log.Printf("Server starting on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
