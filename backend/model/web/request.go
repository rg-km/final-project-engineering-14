package web

import (
	"errors"
	"net"
	"net/mail"
	"strings"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

var (
	ErrRequiredUsername = errors.New("required username")
	ErrRequiredEmail    = errors.New("required email")
	ErrRequiredPassword = errors.New("required password")
	ErrRequiredPhone    = errors.New("required phone")
	ErrInvalidEmail     = errors.New("invalid email")
	ErrDomainNotFound   = errors.New("domain not found")
	ErrInvalidAction    = errors.New("invalid action")
)

func (r *RegisterRequest) PrepareRegister() {
	r.Username = strings.TrimSpace(r.Username)
	r.Email = strings.TrimSpace(r.Email)
	r.Password = strings.TrimSpace(r.Password)
	r.Phone = strings.TrimSpace(r.Phone)
}

func (r *RegisterRequest) ValidateRegister(action string) error {
	switch strings.ToLower(action) {
	case "create":
		if r.Username == "" {
			return ErrRequiredUsername
		}
		if err := validateEmail(r.Email); err != nil {
			return err
		}
		if r.Password == "" {
			return ErrRequiredPassword
		}
		if r.Phone == "" {
			return ErrRequiredPhone
		}
		return nil
	}

	return ErrInvalidAction
}

func validateEmail(email string) error {
	if email == "" {
		return ErrRequiredEmail
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return ErrInvalidEmail
	}

	parts := strings.Split(email, "@")

	_, err = net.LookupMX(parts[1])
	if err != nil {
		return ErrDomainNotFound
	}

	return nil
}
