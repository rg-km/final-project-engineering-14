package helper

import (
	"github.com/rg-km/final-project-engineering-14/backend/model/domain"
	"github.com/rg-km/final-project-engineering-14/backend/model/web"
)

func ToRegisterResponse(user domain.UserDomain, email string) web.RegisterResponse {
	return web.RegisterResponse{
		Username: user.Username,
		Email:    email,
		Phone:    user.Phone,
	}
}

func ToLoginResponse(user domain.UserDomain, token string) web.LoginResponse {
	return web.LoginResponse{
		Token: token,
	}
}

func ToProgrammingLanguangeResponses(todoList []domain.ProgrammingLanguangeDomain) []web.ProgrammingLanguageResponse {
	var todoListResponse []web.ProgrammingLanguageResponse
	for _, todo := range todoList {
		todoListResponse = append(todoListResponse, web.ProgrammingLanguageResponse{
			Name: todo.Name,
		})
	}
	return todoListResponse
}

func ToQuestionResponse(question domain.QuestionDomain, answer []domain.AnswersDomain, progLangDomain domain.ProgrammingLanguangeDomain) web.QuestionCreateResponse {
	return web.QuestionCreateResponse{
		Question:             question.Question,
		Answer:               ToAnswerResponses(answer),
		ProgrammingLanguange: progLangDomain.Name,
	}
}

func ToQuestionPageResponse(quesReq []web.QuestionRequest, answer []domain.AnswersDomain) []web.QuestionCreateResponse {
	var questionListResponse []web.QuestionCreateResponse
	for _, question := range quesReq {
		questionListResponse = append(questionListResponse, web.QuestionCreateResponse{
			Question:             question.Question,
			Answer:               ToAnswerResponses(answer),
			ProgrammingLanguange: question.ProgrammingLanguange,
		})
	}
	return questionListResponse
}

func ToAnswerResponses(answerList []domain.AnswersDomain) []web.AnswerResponse {
	var answerListResponse []web.AnswerResponse
	for _, answer := range answerList {
		answerListResponse = append(answerListResponse, web.AnswerResponse{
			Answer: answer.Answer,
		})
	}
	return answerListResponse
}

func ToQuestionResponses(questionList []web.QuestionRequest) []web.QuestionResponse {
	var questionListResponse []web.QuestionResponse
	for _, question := range questionList {
		questionListResponse = append(questionListResponse, web.QuestionResponse{
			Question:             question.Question,
			ProgrammingLanguange: question.ProgrammingLanguange,
		})
	}
	return questionListResponse
}

func ToCountQuestionDashboardAdmin(countQuestion []web.CountQuestionResponse) []web.CountQuestionResponse {
	var countQuestionResponse []web.CountQuestionResponse
	for _, count := range countQuestion {
		countQuestionResponse = append(countQuestionResponse, web.CountQuestionResponse{
			ProgrammingLanguage: count.ProgrammingLanguage,
			TotalQuestion:       count.TotalQuestion,
		})
	}
	return countQuestionResponse
}

func ToAnswerAttemptResponse(user domain.UserDomain, score int32, level string) web.TotalAnswerAttemptResponse {
	return web.TotalAnswerAttemptResponse{
		Username: user.Username,
		Level:    level,
		Score:    score,
	}
}
