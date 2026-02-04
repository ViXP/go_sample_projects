package controllers

import (
	"auth-service/internal/data"
	"auth-service/internal/grpc_server"
	"context"
	"log"
)

type UsersGRPCController struct {
	grpc_server.UnimplementedAuthServiceServer
	models *data.Models
}

func (service *UsersGRPCController) Authenticate(ctx context.Context, request *grpc_server.AuthRequestPayload) (*grpc_server.AuthResponsePayload, error) {
	response := grpc_server.AuthResponsePayload{}

	user, err := service.models.User.FindByEmail(request.GetEmail())

	if err != nil {
		response.Error = true
		response.Message = "User not found"

		return &response, err
	}

	if user.IsCorrectPassword(request.GetPassword()) {
		log.Printf("User #%v is authenticated", user.ID)
		response.Message = "Authenticated"
	} else {
		response.Error = true
		response.Message = "Wrong password"
	}

	return &response, nil
}

func NewUsersGRPCController(store *data.Store) *UsersGRPCController {
	return &UsersGRPCController{
		models: store.Models,
	}
}
