package server

import (
	"broker-service/internal/event"
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
	case "log":
		app.PublishLogEvent(w, r, requestPayload.Log)
	case "mail":
		app.SendMail(w, r, requestPayload.Mail)
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

func (app *App) LogRequest(w http.ResponseWriter, r *http.Request, payload LogPayload) {
	logData, _ := json.MarshalIndent(payload, "", "  ")
	body := bytes.NewBuffer(logData)

	response, err := http.Post(os.Getenv("LOGGER_URL")+"/api/v1/logs", "application/json", body)

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

func (app *App) SendMail(w http.ResponseWriter, r *http.Request, payload MailerPayload) {
	mailData, _ := json.MarshalIndent(payload, "", "  ")
	body := bytes.NewBuffer(mailData)

	response, err := http.Post(os.Getenv("MAILER_URL")+"/api/v1/messages", "application/json", body)
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

func (app *App) PublishLogEvent(w http.ResponseWriter, r *http.Request, payload LogPayload) {
	eventPayload := EventPayload{
		Name: payload.Name,
		Data: payload.Data,
	}

	err := app.emitEvent("log.INFO", eventPayload)

	if err != nil {
		apiview.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	apiview.WriteJSON(w, http.StatusCreated, apiview.JsonResponse{
		Error:   false,
		Message: "Event emitted successfully.",
	})
}

func (app *App) PublishAuthEvent(w http.ResponseWriter, r *http.Request, payload AuthPayload) {
	serializedPayload, err := json.MarshalIndent(payload, "", "  ")

	if err != nil {
		apiview.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	eventPayload := EventPayload{
		Name: "auth",
		Data: string(serializedPayload),
	}

	err = app.emitEvent("auth.INFO", eventPayload)

	if err != nil {
		apiview.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	apiview.WriteJSON(w, http.StatusCreated, apiview.JsonResponse{
		Error:   false,
		Message: "Event emitted successfully.",
	})
}

const rabbitExchangeName = "microservices_topics"

func (app *App) emitEvent(routingKey string, payload EventPayload) error {
	emitter, err := event.NewEmitter(app.RabbitConn, rabbitExchangeName)

	if err != nil {
		return err
	}

	body, err := json.MarshalIndent(&payload, "", "  ")

	if err != nil {
		return err
	}

	return emitter.Emit(string(body), routingKey)
}
