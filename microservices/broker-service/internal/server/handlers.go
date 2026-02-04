package server

import (
	"broker-service/internal/event"
	"broker-service/internal/grpc_server"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/rpc"
	"os"
	"time"

	apiview "github.com/ViXP/go_sample_projects/microservices/api-view-helpers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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
		app.AuthenticateGRPC(w, r, requestPayload.Auth)
	case "log":
		app.PublishLogEvent(w, r, requestPayload.Log)
	case "mail":
		app.SendMailRPC(w, requestPayload.Mail)
	default:
		apiview.ErrorJSON(w, errors.New("unknown action"), http.StatusBadRequest)
	}
}

// REST handlers
func (app *App) TriggerBroker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	responsePayload := apiview.JsonResponse{
		Error:   false,
		Message: "Broker Service is triggered",
	}

	_ = apiview.WriteJSON(w, http.StatusOK, responsePayload)
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

// RPC handlers
func (app *App) SendMailRPC(w http.ResponseWriter, payload MailerPayload) {
	rpcClient, err := rpc.Dial("tcp", os.Getenv("MAILER_RPC_URL"))

	if err != nil {
		apiview.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	var reply string

	err = rpcClient.Call("RPCProcedures.SendMail", payload, &reply)

	if err != nil {
		apiview.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	apiview.WriteJSON(w, http.StatusCreated, apiview.JsonResponse{
		Error:   false,
		Message: reply,
	})
}

// gRPC handlers
const grpcTimeout = 2 * time.Second

func (app *App) AuthenticateGRPC(w http.ResponseWriter, r *http.Request, payload AuthPayload) {
	conn, err := grpc.NewClient(os.Getenv("AUTH_GRPC_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		apiview.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	serviceClient := grpc_server.NewAuthServiceClient(conn)

	requestPayload := grpc_server.AuthRequestPayload{
		Email:    payload.Email,
		Password: payload.Password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), grpcTimeout)

	defer cancel()

	responsePayload, err := serviceClient.Authenticate(ctx, &requestPayload)

	if err != nil {
		apiview.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if responsePayload.Error {
		apiview.ErrorJSON(w, errors.New(responsePayload.Message), http.StatusUnauthorized)
		return
	}

	apiview.WriteJSON(w, http.StatusAccepted, responsePayload)
}

// RabbitMQ publishing handlers
const rabbitExchangeName = "microservices_topics"

func (app *App) PublishLogEvent(w http.ResponseWriter, r *http.Request, payload LogPayload) {
	serializedPayload, err := json.MarshalIndent(payload, "", "  ")

	if err != nil {
		apiview.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	eventPayload := EventPayload{
		Name: "log",
		Data: string(serializedPayload),
	}

	err = app.emitEvent("log.INFO", eventPayload)

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
