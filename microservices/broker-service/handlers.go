package main

import (
	"net/http"
)

func (app *App) Broker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	responsePayload := jsonResponse{
		Error:   false,
		Message: "Broker Service is triggered",
	}

	_ = app.writeJSON(w, http.StatusOK, responsePayload)
}
