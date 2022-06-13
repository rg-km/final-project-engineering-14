package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rg-km/final-project-engineering-14/backend/helper"
	"github.com/rg-km/final-project-engineering-14/backend/model/domain"
	"github.com/rg-km/final-project-engineering-14/backend/model/web"
	"github.com/rg-km/final-project-engineering-14/backend/pkg/repository"
	"github.com/rg-km/final-project-engineering-14/backend/security"
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
	hashed := security.GeneratePasswordHash(request.Password)
	timeLoc, _ := time.LoadLocation("Asia/Jakarta")

	user := domain.UserDomain{
		Username:  request.Username,
		Email:     request.Email,
		Password:  hashed,
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

// func (service *AuthServiceSQLite) Login(ctx context.Context, request web.LoginRequest) (web.LoginResponse, error) {
// 	user, err := service.repository.GetUser(ctx, request.Email, request.Password)
// 	helper.PanicIfError(err)

// 	if user.IsLogin {
// 		return web.LoginResponse{}, errors.New("user already login")
// 	}

// 	return helper.ToLoginResponse(user), nil
// }

func (service *AuthServiceSQLite) GenerateToken(ctx context.Context, email, password string) (web.LoginResponse, error) {
	user, err := service.repository.GetUser(
		ctx, email, security.GeneratePasswordHash(password))
	if err != nil {
		return web.LoginResponse{}, errors.New("login failed")
	}

	if user.IsLogin {
		return web.LoginResponse{}, errors.New("already login")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpires).Unix(),
			Issuer:    "todo-app",
			IssuedAt:  time.Now().Unix(),
		},
		Email:  user.Email,
		Role:   user.Role,
		UserId: user.Id,
	})

	tokenString, err := token.SignedString([]byte(signingKey))
	helper.PanicIfError(err)

	return helper.ToLoginResponse(user, tokenString), nil
}

func (service *AuthServiceSQLite) ParseToken(ctx context.Context, token string) (uint32, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := tkn.Claims.(*Claims)
	if !ok || !tkn.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	return claims.UserId, nil
}

func (service *AuthServiceSQLite) Logout(ctx context.Context, userId uint32) (bool, error) {
	return service.repository.Logout(ctx, userId)
}
