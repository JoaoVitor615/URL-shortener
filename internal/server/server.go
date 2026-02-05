package server

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Run(deps *Dependencies) {

	// Setup routes with chi
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	numeric := router.Group(func(r chi.Router) {
		r.Get("/test", deps.URLNumericHandler.Test)
	})

	random := router.Group(func(r chi.Router) {
		r.Get("/test", deps.URLRandomHandler.Test)
	})

	router.Mount("/numeric", numeric)
	router.Mount("/random", random)

	// Start server
	addr := os.Getenv("PORT")
	if addr == "" {
		addr = ":8080"
	}
	log.Printf("Server starting on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
