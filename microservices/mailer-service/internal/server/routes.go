package server

import (
	"fmt"
	"mailer-service/internal/mailing"
	"net/http"

	apiview "github.com/ViXP/go_sample_projects/microservices/api-view-helpers"
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

func handleSendMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var requestPayload struct {
		To       string `json:"to"`
		From     string `json:"from,omitempty"`
		FromName string `json:"from_name,omitempty"`
		Subject  string `json:"subject"`
		Body     string `json:"body"`
	}

	err := apiview.ReadJSON(w, r, &requestPayload)

	if err != nil {
		apiview.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	message := mailing.NewMessage(requestPayload.To, requestPayload.Subject, requestPayload.Body)
	smtpClient := mailing.NewSMTPClient(requestPayload.From, requestPayload.FromName)

	err = smtpClient.SendMessage(message)

	if err != nil {
		apiview.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	apiview.WriteJSON(w, http.StatusCreated, fmt.Sprintf("Email is sent successfully to %s\n", requestPayload.To))
}
