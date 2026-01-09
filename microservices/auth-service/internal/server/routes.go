package server

import (
	"auth-service/internal/controllers"
	"auth-service/internal/data"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type App struct {
	DB     *sql.DB
	Models *data.Models
}

func (app *App) Routes() http.Handler {
	router := chi.NewRouter()
	router.Use(
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

	router.Use(middleware.Heartbeat("/api/v1/health"))

	router.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			c := controllers.UsersController{}
			r.Post("/authenticate", c.Authenticate)
			r.Post("/", c.Create)
		})
	})

	return router
}
