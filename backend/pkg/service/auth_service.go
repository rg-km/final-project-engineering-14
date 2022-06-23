package service

import (
	"context"
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
	TokenExpires = time.Minute * 30
)

type AuthServiceSQLite struct {
	repository repository.AuthRepository
}

type Claims struct {
	jwt.StandardClaims
	Email  string `json:"email"`
	Role   string `json:"role"`
	UserId int32  `json:"user_id"`
}

func NewAuthService(repository repository.AuthRepository) *AuthServiceSQLite {
	return &AuthServiceSQLite{
		repository: repository,
	}
}

func (service *AuthServiceSQLite) Create(request web.RegisterRequest, email string) (web.RegisterResponse, error) {
	hashed := security.GeneratePasswordHash(request.Password)
	timeLoc, _ := time.LoadLocation("Asia/Jakarta")

	userDomain := domain.UserDomain{
		Username:  request.Username,
		Email:     "",
		Password:  hashed,
		Phone:     request.Phone,
		Role:      "guest",
		IsLogin:   false,
		CreatedAt: time.Now().In(timeLoc),
		UpdatedAt: time.Now().In(timeLoc),
	}

	userResult, err := service.repository.Save(userDomain, email)
	if err != nil {
		return web.RegisterResponse{}, err
	}

	return helper.ToRegisterResponse(userResult, email), nil
}

func (service *AuthServiceSQLite) GenerateToken(email, password string) (web.LoginResponse, error) {
	user, err := service.repository.GetUser(
		email, security.GeneratePasswordHash(password),
	)
	helper.PanicIfError(err)

	// if user.IsLogin {
	// 	return web.LoginResponse{}, errors.New("already login")
	// }

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

func (service *AuthServiceSQLite) ParseToken(ctx context.Context, token string) (int32, string, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, "", err
	}

	claims, ok := tkn.Claims.(*Claims)
	if !ok || !tkn.Valid {
		return 0, "", fmt.Errorf("invalid token")
	}

	newCtx := context.WithValue(ctx, "email", claims.Email)
	ctx = context.WithValue(newCtx, "role", claims.Role)
	ctx = context.WithValue(ctx, "user_id", claims.UserId)
	ctx = context.WithValue(newCtx, "props", claims)

	return claims.UserId, claims.Role, nil
}

func (service *AuthServiceSQLite) Logout(userId int32) (bool, error) {
	return service.repository.Logout(userId)
}
