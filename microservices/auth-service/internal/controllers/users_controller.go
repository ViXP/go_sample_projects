package controllers

import (
	"auth-service/internal/data"
	"auth-service/internal/utils"
	"errors"
	"fmt"
	"net/http"

	apiview "github.com/ViXP/go_sample_projects/microservices/api-view-helpers"
)

type UsersController struct {
	models *data.Models
}

func (controller *UsersController) Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	payload := requestPayload{}

	err := apiview.ReadJSON(w, r, &payload)

	if err != nil {
		apiview.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := controller.models.User.FindByEmail(payload.Email)

	if err != nil {
		apiview.ErrorJSON(w, errors.New("Credentials are incorrect"), http.StatusNotFound)
		return
	}

	if user.IsCorrectPassword(payload.Password) {
		utils.Log(fmt.Sprintf("User #%v is authenticated", user.ID))
		apiview.WriteJSON(w, http.StatusAccepted, "Authenticated")
	} else {
		apiview.ErrorJSON(w, errors.New("Wrong password"), http.StatusUnauthorized)
	}
}

func (controller *UsersController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func NewUsersController(store *data.Store) *UsersController {
	return &UsersController{
		models: store.Models,
	}
}
