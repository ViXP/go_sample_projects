package controllers

import (
	"logger-service/internal/data"
	"net/http"

	apiview "github.com/ViXP/go_sample_projects/microservices/api-view-helpers"
)

type LogEntriesController struct {
	models *data.Models
}

func (lec *LogEntriesController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type requestPayload struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}

	payload := requestPayload{}

	err := apiview.ReadJSON(w, r, &payload)

	if err != nil {
		apiview.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = lec.models.LogEntry.Create(&data.LogEntry{
		Name: payload.Name,
		Data: payload.Data,
	})

	if err != nil {
		apiview.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	apiview.WriteJSON(w, http.StatusCreated, "Log entry created.")
}

func NewLogEntriesController(store *data.Store) *LogEntriesController {
	return &LogEntriesController{store.Models}
}
