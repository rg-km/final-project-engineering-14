package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rg-km/final-project-engineering-14/backend/helper"
	"github.com/rg-km/final-project-engineering-14/backend/model/web"
)

func (handler *Handler) signUp(writer http.ResponseWriter, request *http.Request) {
	handler.AllowOrigin(writer, request)
	var registerRequest web.RegisterRequest
	err := json.NewDecoder(request.Body).Decode(&registerRequest)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}

	registerRequest.PrepareRegister()
	err = registerRequest.ValidateRegister("create")
	if err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}

	registerResponse, err := handler.services.AuthService.Create(
		request.Context(), registerRequest,
	)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    err.Error(),
		})
		return
	}

	webResponse := web.WebResponse{
		Status:  http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
		Data:    registerResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
