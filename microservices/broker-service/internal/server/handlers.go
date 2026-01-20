package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	apiview "github.com/ViXP/go_sample_projects/microservices/api-view-helpers"
)

func (app *App) TriggerBroker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	responsePayload := apiview.JsonResponse{
		Error:   false,
		Message: "Broker Service is triggered",
	}

	_ = apiview.WriteJSON(w, http.StatusOK, responsePayload)
}

func (app *App) HandleProxyRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var requestPayload ProxyRequestPayload

	err := apiview.ReadJSON(w, r, &requestPayload)

	if err != nil {
		apiview.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	switch requestPayload.Action {
	case "auth":
		app.Authenticate(w, r, requestPayload.Auth)
	default:
		apiview.ErrorJSON(w, errors.New("unknown action"), http.StatusBadRequest)
	}
}

func (app *App) Authenticate(w http.ResponseWriter, r *http.Request, a AuthPayload) {
	authData, _ := json.MarshalIndent(a, "", "  ")
	body := bytes.NewBuffer(authData)

	response, err := http.Post(os.Getenv("AUTH_URL")+"/api/v1/users/authenticate", "application/json", body)

	if err != nil {
		apiview.ErrorJSON(w, err)
		return
	}

	defer response.Body.Close()

	var responsePayload apiview.JsonResponse

	err = json.NewDecoder(response.Body).Decode(&responsePayload)

	if err != nil {
		apiview.ErrorJSON(w, err)
		return
	}

	if responsePayload.Error {
		apiview.ErrorJSON(w, errors.New(responsePayload.Message), response.StatusCode)
		return
	}

	apiview.WriteJSON(w, response.StatusCode, responsePayload)
}
