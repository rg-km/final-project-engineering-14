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
	lists, err := service.frontendRepository.GetQuestionByProgrammingLanguange(
		proglangId, perPage, page,
	)
	helper.PanicIfError(err)

	answers, err := service.backendRepository.GetAnswers()
	helper.PanicIfError(err)

	return helper.ToQuestionPageResponse(lists, answers), nil
}

func (service *FrontendServiceSQLite) PostAnswersAttempt(userId int32, answers web.AnswersAttemptRequest) (web.TotalAnswerAttemptResponse, error) {
	var level string
	var levelDomain domain.LevelDomain
	answersDomain := []domain.AnswersDomain{
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

	_, err := service.frontendRepository.SaveAnswersAttempt(userId, answersDomain)
	helper.PanicIfError(err)

	score, err := service.frontendRepository.CountAnswersAttempt(userId)
	helper.PanicIfError(err)

	if score >= 75 {
		level = "Intermediate"
	} else {
		level = "Beginner"
	}

	if level != "" {
		levelResp, err := service.frontendRepository.FindLevelByName(level)
		helper.PanicIfError(err)
		levelDomain = levelResp
	}

	user, err := service.authrepository.GetUserById(userId)
	helper.PanicIfError(err)

	return helper.ToAnswerAttemptResponse(user, score, levelDomain), nil
}

func (service *FrontendServiceSQLite) SeeRecommendationByLevelId(levelId int32) (web.RecommendationResponse, error) {
	lists, err := service.frontendRepository.FindRecommendationByLevelId(levelId)
	helper.PanicIfError(err)

	return helper.ToRecommendationResponses(lists), nil
}
