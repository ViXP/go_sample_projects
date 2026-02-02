package event

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Handler interface {
	Handle(payload *Payload)
}

type LogHandler struct{}

func (h *LogHandler) Handle(payload *Payload) {
	logPayload := struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}{
		Name: payload.Name,
		Data: payload.Data,
	}
	logData, _ := json.MarshalIndent(logPayload, "", "  ")
	body := bytes.NewBuffer(logData)

	resp, err := http.Post(os.Getenv("LOGGER_URL")+"/api/v1/logs", "application/json", body)

	if err != nil {
		log.Printf("Error logging message: %v\n", err)
		return
	}

	defer resp.Body.Close()
}

type AuthHandler struct{}

func (h *AuthHandler) Handle(payload *Payload) {
	var authPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.Unmarshal([]byte(payload.Data), &authPayload)

	if err != nil {
		log.Printf("Error unmarshaling auth payload: %v\n", err)
		return
	}

	authData, _ := json.MarshalIndent(authPayload, "", "  ")
	body := bytes.NewBuffer(authData)

	resp, err := http.Post(os.Getenv("AUTH_URL")+"/api/v1/users/authenticate", "application/json", body)

	if err != nil {
		log.Printf("Error authenticating user: %v\n", err)
		return
	}

	defer resp.Body.Close()
}

type DefaultHandler struct{}

func (h *DefaultHandler) Handle(payload *Payload) {
	log.Printf("Received message %s with payload: %s\n", payload.Name, payload.Data)
}

var (
	_ Handler = &LogHandler{}
	_ Handler = &AuthHandler{}
	_ Handler = &DefaultHandler{}
)
