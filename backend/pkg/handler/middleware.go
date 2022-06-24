package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/rg-km/final-project-engineering-14/backend/model/web"
)

const (
	authorizationHeader = "Authorization"
	userCtxKey          = "user_id"
	roleCtxKey          = "role"
)

func (handler *Handler) AllowOrigin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		}

		next.ServeHTTP(w, r)
	})
}

func (handler *Handler) AuthMiddleWare(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		header := r.Header.Get(authorizationHeader)
		if header == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(web.WebResponse{
				Status:  http.StatusUnauthorized,
				Message: "UNAUTHORIZED",
				Data:    "Need Authorization header",
			})
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(web.WebResponse{
				Status:  http.StatusUnauthorized,
				Message: "UNAUTHORIZED",
				Data:    "Invalid Authorization header",
			})
			return
		}

		userId, userRole, err := handler.service.AuthService.ParseToken(r.Context(), headerParts[1])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(web.WebResponse{
				Status:  http.StatusUnauthorized,
				Message: "UNAUTHORIZED",
				Data:    err.Error(),
			})
			return
		}

		ctx := context.WithValue(r.Context(), userCtxKey, userId)
		ctx = context.WithValue(ctx, roleCtxKey, userRole)

		r = r.WithContext(ctx)

		next(w, r, ps)
	}
}

func (handler *Handler) AdminMiddleWare(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		encoder := json.NewEncoder(w)
		role := r.Context().Value(roleCtxKey)
		if role != "admin" {
			w.WriteHeader(http.StatusForbidden)
			encoder.Encode(web.WebResponse{
				Status:  http.StatusForbidden,
				Message: http.StatusText(http.StatusForbidden),
				Data:    "Forbidden access",
			})
			return
		}

		next(w, r, ps)
	}
}

func getUserId(request *http.Request) (int32, error) {
	userId, ok := request.Context().Value(userCtxKey).(int32)
	if !ok {
		return 0, errors.New("user id not found")
	}

	return userId, nil
}
