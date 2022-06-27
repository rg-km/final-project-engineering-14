package service

import (
	"context"

	"github.com/rg-km/final-project-engineering-14/backend/model/web"
	"github.com/rg-km/final-project-engineering-14/backend/pkg/repository"
)

type AuthService interface {
	Create(request web.RegisterRequest, email string) (web.RegisterResponse, error)
	GenerateToken(email, password string) (web.LoginResponse, error)
	ParseToken(ctx context.Context, token string) (int32, string, error)
	Logout(userId int32) (bool, error)
}

type FrontendService interface {
	GetProgrammingLanguanges() ([]web.ProgrammingLanguageResponse, error)
	GetQuestionByProgrammingLanguage(proglangId, perPage, page int32) ([]web.QuestionCreateResponse, error)
	PostAnswersAttempt(userId int32, answers web.AnswersAttemptRequest) (web.TotalAnswerAttemptResponse, error)
	SeeRecommendationByLevelId(levelId int32) (web.RecommendationResponse, error)
}

type BackendService interface {
	CreateAdminQuestion(userId int32, question web.QuestionRequest) (web.QuestionCreateResponse, error)
	GetAdminQuestions() ([]web.QuestionResponse, error)
	GetCountQuestions() ([]web.CountQuestionResponse, error)
	UpdateAdminQuestion(question web.QuestionRequest, questionId int32) (web.QuestionResponse, error)
	DeleteAdminQuestion(questionId int32) (bool, error)
}

type Service struct {
	AuthService     AuthService
	FrontendService FrontendService
	BackendService  BackendService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		AuthService: NewAuthService(repository.AuthRepository),
		FrontendService: NewFrontendService(
			repository.AuthRepository,
			repository.FrontendRepository,
			repository.BackendRepository,
		),
		BackendService: NewBackendService(
			repository.BackendRepository,
			repository.FrontendRepository,
		),
	}
}
