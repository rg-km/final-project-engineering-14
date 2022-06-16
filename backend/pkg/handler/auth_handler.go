package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rg-km/final-project-engineering-14/backend/helper"
	"github.com/rg-km/final-project-engineering-14/backend/model/web"
	"github.com/rg-km/final-project-engineering-14/backend/pkg/service"
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

func (handler *Handler) signIn(writer http.ResponseWriter, request *http.Request) {
	handler.AllowOrigin(writer, request)
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

	loginRequest.PrepareLogin()
	err = loginRequest.ValidateLogin("login")
	if err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}

	userLoginResponse, err := handler.services.AuthService.GenerateToken(
		request.Context(), loginRequest.Email, loginRequest.Password,
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
		Message: http.StatusText(http.StatusOK),
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

func (handler *Handler) signOut(writer http.ResponseWriter, request *http.Request) {
	handler.AllowOrigin(writer, request)
	token, err := request.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(web.WebResponse{
				Status:  http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
				Data:    "No token provided",
			})
			return
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}

	if token.Value == "" {
		writer.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusUnauthorized,
			Message: http.StatusText(http.StatusUnauthorized),
			Data:    "No token provided",
		})
		return
	}
	userId, _ := handler.services.AuthService.ParseToken(request.Context(), token.Value)

	_, err = handler.services.AuthService.Logout(request.Context(), uint32(userId))
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
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		MaxAge:   -1,
	}

	http.SetCookie(writer, &cookie)

	webResponse := web.WebResponse{
		Status:  http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    "Successfully logged out",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
