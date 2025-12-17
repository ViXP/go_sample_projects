package main

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *App) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	var maxBytes int64 = 1 * 1024 * 1024

	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(data)

	if err != nil {
		return err
	}

	return nil
}

func (app *App) writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	serialized, err := json.Marshal(data)

	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(serialized)

	if err != nil {
		return err
	}

	return nil
}

func (app *App) errorJSON(w http.ResponseWriter, err error, status ...int) error {
	var statusCode int = http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var responsePayload jsonResponse
	responsePayload.Error = true
	responsePayload.Message = err.Error()

	return app.writeJSON(w, statusCode, responsePayload)
}
