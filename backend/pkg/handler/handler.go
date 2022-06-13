package handler

import (
	"fmt"
	"net/http"

	"github.com/rg-km/final-project-engineering-14/backend/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (handler *Handler) InitRoutes() *http.ServeMux {
	router := http.NewServeMux()

	fmt.Printf("\n\t[INFO] Running in localhost:%d\n", 8080)

	router.HandleFunc("/auth/signup", handler.signUp)

	return router
}
