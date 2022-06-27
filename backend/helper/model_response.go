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
		Id:       user.Id,
		Username: user.Username,
		Role:     user.Role,
		Token:    token,
	}
}

func ToProgrammingLanguangeResponses(todoList []domain.ProgrammingLanguangeDomain) []web.ProgrammingLanguageResponse {
	var todoListResponse []web.ProgrammingLanguageResponse
	for _, todo := range todoList {
		todoListResponse = append(todoListResponse, web.ProgrammingLanguageResponse{
			Id:        todo.Id,
			Name:      todo.Name,
			UrlImages: todo.UrlImages,
		})
	}
	return todoListResponse
}

func ToQuestionResponse(question domain.QuestionDomain, answer []domain.AnswersDomain, progLangDomain domain.ProgrammingLanguangeDomain) web.QuestionCreateResponse {
	return web.QuestionCreateResponse{
		Id:                   question.Id,
		Question:             question.Question,
		Answer:               ToAnswerResponses(answer),
		ProgrammingLanguange: progLangDomain.Name,
	}
}

func ToUpdateQuestionResponse(questionResp web.QuestionResponse) web.QuestionResponse {
	return web.QuestionResponse{
		Id:                   questionResp.Id,
		Question:             questionResp.Question,
		ProgrammingLanguange: questionResp.ProgrammingLanguange,
	}
}

func ToQuestionPageResponse(quesReq []web.QuestionRequest, answer []domain.AnswersDomain) []web.QuestionCreateResponse {
	var questionListResponse []web.QuestionCreateResponse
	for _, question := range quesReq {
		questionListResponse = append(questionListResponse, web.QuestionCreateResponse{
			Id:                   question.Id,
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
			Id:                   question.Id,
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

func ToAnswerAttemptResponse(user domain.UserDomain, score int32, levelDomain domain.LevelDomain) web.TotalAnswerAttemptResponse {
	return web.TotalAnswerAttemptResponse{
		LevelId:  levelDomain.Id,
		Username: user.Username,
		Level:    levelDomain.Level,
		Score:    score,
	}
}

func ToRecommendationResponses(recommendation domain.RecommendationDomain) web.RecommendationResponse {
	return web.RecommendationResponse{
		ImageUrl: recommendation.ImageUrl,
	}
}
