package service

import (
	"context"

	"github.com/rg-km/final-project-engineering-14/backend/model/web"
	"github.com/rg-km/final-project-engineering-14/backend/pkg/repository"
)

type AuthService interface {
	Create(ctx context.Context, request web.RegisterRequest) (web.RegisterResponse, error)
}

type Service struct {
	AuthService AuthService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		AuthService: NewAuthService(repository.AuthRepository),
	}
}
