package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		}),
		middleware.Logger,
	)

	mux.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Use(middleware.Heartbeat("/health"))
			r.Post("/messages", handleSendMessage)
		})
	})

	return mux
}
