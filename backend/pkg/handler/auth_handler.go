package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/rg-km/final-project-engineering-14/backend/helper"
	"github.com/rg-km/final-project-engineering-14/backend/model/web"
	"github.com/rg-km/final-project-engineering-14/backend/pkg/service"
)

func (handler *Handler) signUp(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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

	err = registerRequest.ValidateRegister()
	if err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}

	registerResponse, err := handler.service.AuthService.Create(
		registerRequest, registerRequest.Email,
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
		Message: "CREATE OK",
		Data:    registerResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *Handler) signIn(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var loginRequest web.LoginRequest
	err := json.NewDecoder(request.Body).Decode(&loginRequest)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}

	err = loginRequest.ValidateLogin()
	if err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}

	userLoginResponse, err := handler.service.AuthService.GenerateToken(
		loginRequest.Email, loginRequest.Password,
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
		Status:  http.StatusOK,
		Message: "LOGIN OK",
		Data:    userLoginResponse,
	}

	http.SetCookie(writer, &http.Cookie{
		Name:    "token",
		Value:   userLoginResponse.Token,
		Expires: time.Now().Add(service.TokenExpires),
		Path:    "/",
	})

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *Handler) signOut(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	token, _ := request.Cookie("token")

	userId, _, _ := handler.service.AuthService.ParseToken(
		request.Context(), token.Value,
	)
	_, err := handler.service.AuthService.Logout(int32(userId))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    err.Error(),
		})
		return
	}

	cookie := http.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
	}

	http.SetCookie(writer, &cookie)

	webResponse := web.WebResponse{
		Status:  http.StatusOK,
		Message: "LOGOUT OK",
		Data:    "Successfully logged out",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
