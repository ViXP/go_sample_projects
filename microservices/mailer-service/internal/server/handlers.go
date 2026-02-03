package server

import (
	"fmt"
	"mailer-service/internal/mailing"
	"net/http"

	apiview "github.com/ViXP/go_sample_projects/microservices/api-view-helpers"
)

type EmailRequestPayload struct {
	To       string `json:"to"`
	From     string `json:"from,omitempty"`
	FromName string `json:"from_name,omitempty"`
	Subject  string `json:"subject"`
	Body     string `json:"body"`
}

func handleSendMessage(w http.ResponseWriter, r *http.Request) {
	requestPayload := EmailRequestPayload{}

	w.Header().Set("Content-Type", "application/json")

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

	apiview.WriteJSON(w, http.StatusCreated, fmt.Sprintf("Email is sent successfully to %s", requestPayload.To))
}
