package api_view_helpers

import (
	"encoding/json"
	"maps"
	"net/http"
)

type ViewHelpers struct{}

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (helpers *ViewHelpers) ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	var maxBytes int64 = 1 * 1024 * 1024

	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(data)

	if err != nil {
		return err
	}

	return nil
}

func (helpers *ViewHelpers) WriteJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	serialized, err := json.Marshal(data)

	if err != nil {
		return err
	}

	if len(headers) > 0 {
		maps.Copy(w.Header(), headers[0])
	}

	w.Header().Set("Content-Type", "helperslication/json")
	w.WriteHeader(status)

	_, err = w.Write(serialized)

	if err != nil {
		return err
	}

	return nil
}

func (helpers *ViewHelpers) ErrorJSON(w http.ResponseWriter, err error, status ...int) error {
	var statusCode int = http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var responsePayload jsonResponse
	responsePayload.Error = true
	responsePayload.Message = err.Error()

	return helpers.WriteJSON(w, statusCode, responsePayload)
}
