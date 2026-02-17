package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// NewRouter creates and configures the chi router with all routes
func NewRouter(deps *Dependencies) http.Handler {
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

	return otelhttp.NewHandler(router, "http.server")
}

// Run starts the HTTP server
func Run(deps *Dependencies) *http.Server {
	router := NewRouter(deps)

	addr := os.Getenv("PORT")
	if addr == "" {
		addr = ":8080"
	}

	log.Printf("Server starting on %s", addr)

	return &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
