package service

import (
	"github.com/rg-km/final-project-engineering-14/backend/helper"
	"github.com/rg-km/final-project-engineering-14/backend/model/domain"
	"github.com/rg-km/final-project-engineering-14/backend/model/web"
	"github.com/rg-km/final-project-engineering-14/backend/pkg/repository"
)

type FrontendServiceSQLite struct {
	authrepository     repository.AuthRepository
	frontendRepository repository.FrontendRepository
	backendRepository  repository.BackendRepository
}

func NewFrontendService(authrepository repository.AuthRepository, frontendRepository repository.FrontendRepository, backendRepository repository.BackendRepository) *FrontendServiceSQLite {
	return &FrontendServiceSQLite{
		authrepository:     authrepository,
		frontendRepository: frontendRepository,
		backendRepository:  backendRepository,
	}
}

func (service *FrontendServiceSQLite) GetProgrammingLanguanges() ([]web.ProgrammingLanguageResponse, error) {
	lists, err := service.frontendRepository.GetProgrammingLanguanges()
	helper.PanicIfError(err)

	return helper.ToProgrammingLanguangeResponses(lists), nil
}

func (service *FrontendServiceSQLite) GetQuestionByProgrammingLanguage(proglangId, perPage, page int32) ([]web.QuestionCreateResponse, error) {
	lists, err := service.frontendRepository.GetQuestionByProgrammingLanguange(proglangId, perPage, page)
	helper.PanicIfError(err)

	answers, err := service.backendRepository.GetAnswers()
	helper.PanicIfError(err)

	return helper.ToQuestionPageResponse(lists, answers), nil
}

func (service *FrontendServiceSQLite) PostAnswersAttempt(userId int32, answers web.AnswersAttemptRequest) (web.TotalAnswerAttemptResponse, error) {
	var level string
	test1 := []domain.AnswersDomain{
		{Answer: answers.AnswerOne},
		{Answer: answers.AnswerTwo},
		{Answer: answers.AnswerThree},
		{Answer: answers.AnswerFour},
		{Answer: answers.AnswerFive},
		{Answer: answers.AnswerSix},
		{Answer: answers.AnswerSeven},
		{Answer: answers.AnswerEight},
		{Answer: answers.AnswerNine},
		{Answer: answers.AnswerTen},
	}

	_, err := service.frontendRepository.SaveAnswersAttempt(userId, test1)
	helper.PanicIfError(err)

	score, err := service.frontendRepository.CountAnswersAttempt(userId)

	if score >= 75 {
		level = "INTERMEDIATE"
	} else {
		level = "BEGINNER"
	}

	user, err := service.authrepository.GetUserById(userId)
	helper.PanicIfError(err)

	return helper.ToAnswerAttemptResponse(user, score, level), nil
}
