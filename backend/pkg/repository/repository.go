package repository

import (
	"database/sql"

	"github.com/rg-km/final-project-engineering-14/backend/model/domain"
	"github.com/rg-km/final-project-engineering-14/backend/model/web"
)

type AuthRepository interface {
	Save(user domain.UserDomain, email string) (domain.UserDomain, error)
	GetUser(email, password string) (domain.UserDomain, error)
	GetUserById(userId int32) (domain.UserDomain, error)
	Logout(userId int32) (bool, error)
}

type FrontendRepository interface {
	GetProgrammingLanguanges() ([]domain.ProgrammingLanguangeDomain, error)
	GetProgrammingLanguangeByName(proglangName string) (domain.ProgrammingLanguangeDomain, error)
	GetQuestionByProgrammingLanguange(proglangId, perPage, page int32) ([]web.QuestionRequest, error)
	SaveAnswersAttempt(userId int32, webReq []domain.AnswersDomain) (bool, error)
	CountAnswersAttempt(userId int32) (int32, error)
}

type BackendRepository interface {
	SaveQuestion(question domain.QuestionDomain, progLang domain.ProgrammingLanguangeDomain, userId int32) (domain.QuestionDomain, error)
	GetQuestions() ([]web.QuestionRequest, error)
	GetCountQuestions() ([]web.CountQuestionResponse, error)
	UpdateQuestion(proglang domain.ProgrammingLanguangeDomain, question web.QuestionRequest, questionId int32) (bool, error)
	DeleteQuestion(questionId int32) (bool, error)
	GetAnswers() ([]domain.AnswersDomain, error)
}

type Repository struct {
	AuthRepository     AuthRepository
	FrontendRepository FrontendRepository
	BackendRepository  BackendRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		AuthRepository:     NewAuthRepository(db),
		FrontendRepository: NewFrontendRepository(db),
		BackendRepository:  NewBackendRepository(db),
	}
}
