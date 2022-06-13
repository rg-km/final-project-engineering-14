package helper

import (
	"github.com/rg-km/final-project-engineering-14/backend/model/domain"
	"github.com/rg-km/final-project-engineering-14/backend/model/web"
)

func ToRegisterResponse(user domain.UserDomain) web.RegisterResponse {
	return web.RegisterResponse{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Phone:    user.Phone,
	}
}
