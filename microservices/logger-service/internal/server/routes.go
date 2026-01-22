package server

import (
	"logger-service/internal/controllers"
	"logger-service/internal/data"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Routes(store *data.Store) http.Handler {
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

			r.Route("/logs", func(r chi.Router) {
				logEntriesController := controllers.NewLogEntriesController(store)

				r.Post("/", logEntriesController.Create)
			})
		})
	})

	return mux
}
