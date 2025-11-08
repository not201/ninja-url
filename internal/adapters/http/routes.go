package http

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func SetupRoutes(handler *handler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(httprate.LimitByIP(20, time.Minute))

	r.Route("/api", func(r chi.Router) {
		r.Get("/health", handler.Health)
		r.Post("/shorten", handler.Shorten)
	})

	r.Get("/{shortCode:[A-Za-z0-9_-]{6}}", handler.Resolver)
	r.Get("/*", handler.Static)

	r.NotFound(handler.NotFound)
	r.MethodNotAllowed(handler.MethodNotAllowed)

	return r
}
