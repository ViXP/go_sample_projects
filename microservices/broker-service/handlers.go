package main

import (
	"encoding/json"
	"net/http"
)

func (app *App) Broker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	responsePayload := jsonResponse{
		Error:   false,
		Message: "Broker Service is triggered",
	}

	output, _ := json.MarshalIndent(responsePayload, "", "\t")

	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
