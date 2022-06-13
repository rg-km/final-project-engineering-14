package service

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rg-km/final-project-engineering-14/backend/helper"
	"github.com/rg-km/final-project-engineering-14/backend/model/domain"
	"github.com/rg-km/final-project-engineering-14/backend/model/web"
	"github.com/rg-km/final-project-engineering-14/backend/pkg/repository"
)

const (
	signingKey   = "secret_key"
	TokenExpires = time.Minute * 20
)

type AuthServiceSQLite struct {
	repository repository.AuthRepository
}

type Claims struct {
	jwt.StandardClaims
	Email  string `json:"email"`
	Role   string `json:"role"`
	UserId uint32 `json:"user_id"`
}

func NewAuthService(repository repository.AuthRepository) *AuthServiceSQLite {
	return &AuthServiceSQLite{
		repository: repository,
	}
}

func (service *AuthServiceSQLite) Create(ctx context.Context, request web.RegisterRequest) (web.RegisterResponse, error) {
	timeLoc, _ := time.LoadLocation("Asia/Jakarta")

	user := domain.UserDomain{
		Username:  request.Username,
		Email:     request.Email,
		Password:  request.Password,
		Phone:     request.Phone,
		Role:      "guest",
		IsLogin:   false,
		CreatedAt: time.Now().In(timeLoc),
		UpdatedAt: time.Now().In(timeLoc),
	}

	userResult, err := service.repository.Save(ctx, user)
	helper.PanicIfError(err)

	return helper.ToRegisterResponse(userResult), nil
}

func (service *AuthServiceSQLite) Login(ctx context.Context, request web.LoginRequest) (web.LoginResponse, error) {
	user, err := service.repository.GetUser(ctx, request.Email, request.Password)
	helper.PanicIfError(err)

	if user.IsLogin {
		return web.LoginResponse{}, errors.New("user already login")
	}

	return helper.ToLoginResponse(user), nil
}
