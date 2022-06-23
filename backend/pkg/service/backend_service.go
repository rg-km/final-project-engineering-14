package service

import (
	"time"

	"github.com/rg-km/final-project-engineering-14/backend/helper"
	"github.com/rg-km/final-project-engineering-14/backend/model/domain"
	"github.com/rg-km/final-project-engineering-14/backend/model/web"
	"github.com/rg-km/final-project-engineering-14/backend/pkg/repository"
)

type BackendSQLite struct {
	backendRepository  repository.BackendRepository
	frontendRepository repository.FrontendRepository
}

func NewBackendService(backendRepository repository.BackendRepository, frontendRepository repository.FrontendRepository) *BackendSQLite {
	return &BackendSQLite{
		backendRepository:  backendRepository,
		frontendRepository: frontendRepository,
	}
}

func (service *BackendSQLite) CreateAdminQuestion(userId int32, question web.QuestionRequest) (web.QuestionCreateResponse, error) {
	timeLoc, _ := time.LoadLocation("Asia/Jakarta")
	var resultQuestionDomain domain.QuestionDomain

	progLangDomain, err := service.frontendRepository.GetProgrammingLanguangeByName(
		question.ProgrammingLanguange,
	)
	helper.PanicIfError(err)

	questionDomain := domain.QuestionDomain{
		Question:   question.Question,
		ProglangId: progLangDomain.Id,
		CreatedAt:  time.Now().In(timeLoc),
		UpdatedAt:  time.Now().In(timeLoc),
	}

	answersDomain, err := service.backendRepository.GetAnswers()
	helper.PanicIfError(err)

	respQuestionDomain, err := service.backendRepository.SaveQuestion(
		questionDomain, progLangDomain, userId,
	)
	helper.PanicIfError(err)

	resultQuestionDomain = respQuestionDomain

	return helper.ToQuestionResponse(
		resultQuestionDomain, answersDomain, progLangDomain,
	), nil
}

func (service *BackendSQLite) GetAdminQuestions() ([]web.QuestionResponse, error) {
	questionsDomain, err := service.backendRepository.GetQuestions()
	helper.PanicIfError(err)

	return helper.ToQuestionResponses(questionsDomain), nil
}

func (service *BackendSQLite) GetCountQuestions() ([]web.CountQuestionResponse, error) {
	countQuestion, err := service.backendRepository.GetCountQuestions()
	helper.PanicIfError(err)

	return helper.ToCountQuestionDashboardAdmin(countQuestion), nil
}

func (service *BackendSQLite) UpdateAdminQuestion(question web.QuestionRequest, questionId int32) (bool, error) {
	var progLangDomain domain.ProgrammingLanguangeDomain

	if question.ProgrammingLanguange != "" {
		progLangDomainRepo, err := service.frontendRepository.GetProgrammingLanguangeByName(question.ProgrammingLanguange)
		helper.PanicIfError(err)

		progLangDomain = progLangDomainRepo
	}

	_, err := service.backendRepository.UpdateQuestion(
		progLangDomain, question, questionId,
	)
	helper.PanicIfError(err)

	return true, nil
}

func (service *BackendSQLite) DeleteAdminQuestion(questionId int32) (bool, error) {
	_, err := service.backendRepository.DeleteQuestion(questionId)
	helper.PanicIfError(err)

	return true, nil
}
