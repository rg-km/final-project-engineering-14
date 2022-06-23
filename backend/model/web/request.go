package web

import (
	"errors"
	"net"
	"net/mail"
	"strings"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type QuestionRequest struct {
	ProgrammingLanguange string `json:"programming_languange"`
	Question             string `json:"question"`
}

type AnswersAttemptRequest struct {
	AnswerOne   string `json:"1"`
	AnswerTwo   string `json:"2"`
	AnswerThree string `json:"3"`
	AnswerFour  string `json:"4"`
	AnswerFive  string `json:"5"`
	AnswerSix   string `json:"6"`
	AnswerSeven string `json:"7"`
	AnswerEight string `json:"8"`
	AnswerNine  string `json:"9"`
	AnswerTen   string `json:"10"`
}

var (
	ErrRequiredUsername         = errors.New("required username")
	ErrRequiredEmail            = errors.New("required email")
	ErrRequiredPassword         = errors.New("required password")
	ErrRequiredPhone            = errors.New("required phone")
	ErrInvalidEmail             = errors.New("invalid email")
	ErrDomainNotFound           = errors.New("domain not found")
	ErrInvalidPhone             = errors.New("invalid phone")
	ErrInvalidUsername          = errors.New("invalid username")
	ErrInvalidPassword          = errors.New("invalid password")
	ErrInvalidTitle             = errors.New("invalid title")
	ErrInvalidAction            = errors.New("invalid action")
	ErrRequiredProglang         = errors.New("programming language is required")
	ErrRequiredLevel            = errors.New("level is required")
	ErrRequiredQuestion         = errors.New("question is required")
	ErrRequiredQPL              = errors.New("question and programming language are required")
	ErrAllRequestAnswersAttempt = errors.New("all request answers attempt is required")
)

func (l *LoginRequest) ValidateLogin() error {
	if err := validateEmail(l.Email); err != nil {
		return err
	}
	if l.Password == "" {
		return ErrInvalidPassword
	}

	return nil
}

func (r *RegisterRequest) ValidateRegister() error {
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

func (q *QuestionRequest) ValidateQuestion(action string) error {
	switch strings.ToLower(action) {
	case "create":
		if q.Question == "" {
			return ErrRequiredQuestion
		}
		if q.ProgrammingLanguange == "" {
			return ErrRequiredProglang
		}
		return nil
	case "update":
		if q.Question == "" && q.ProgrammingLanguange == "" {
			return ErrRequiredQPL
		}
		return nil
	}

	return ErrInvalidAction
}

func (a *AnswersAttemptRequest) ValidateAnswersAttempt() error {
	if a.AnswerOne == "" && a.AnswerTwo == "" && a.AnswerThree == "" && a.AnswerFour == "" && a.AnswerFive == "" && a.AnswerSix == "" && a.AnswerSeven == "" && a.AnswerEight == "" && a.AnswerNine == "" && a.AnswerTen == "" {
		return ErrAllRequestAnswersAttempt
	}

	return nil
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
