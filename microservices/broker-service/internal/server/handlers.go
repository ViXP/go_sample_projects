package server

import (
	"net/http"

	apiview "github.com/ViXP/go_sample_projects/microservices/api-view-helpers"
)

func (app *App) Broker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	responsePayload := apiview.JsonResponse{
		Error:   false,
		Message: "Broker Service is triggered",
	}

	_ = apiview.WriteJSON(w, http.StatusOK, responsePayload)
}
