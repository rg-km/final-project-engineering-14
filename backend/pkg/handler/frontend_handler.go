package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/rg-km/final-project-engineering-14/backend/helper"
	"github.com/rg-km/final-project-engineering-14/backend/model/web"
)

func (handler *Handler) getProgrammingLanguanges(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	proglangResponses, err := handler.service.FrontendService.GetProgrammingLanguanges()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    err.Error(),
		})
		return
	}

	webResponse := web.WebResponse{
		Status:  http.StatusOK,
		Message: "GET OK",
		Data:    proglangResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *Handler) getQuestionByProgrammingLanguage(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	proglangIdParam := params.ByName("proglangId")
	proglangId, err := strconv.Atoi(proglangIdParam)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    err.Error(),
		})
		return
	}

	pageStr := request.URL.Query().Get("page")
	perPageStr := request.URL.Query().Get("per_page")

	page, _ := strconv.Atoi(pageStr)
	perPage, _ := strconv.Atoi(perPageStr)

	questionResponses, err := handler.service.FrontendService.GetQuestionByProgrammingLanguage(int32(proglangId), int32(perPage), int32(page))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    err.Error(),
		})
		return
	}

	webResponse := web.WebResponse{
		Status:  http.StatusOK,
		Message: "GET OK",
		Data:    questionResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *Handler) postAnswersAttempt(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId, _ := getUserId(request)
	var answerAttempt web.AnswersAttemptRequest
	err := json.NewDecoder(request.Body).Decode(&answerAttempt)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    err.Error(),
		})
		return
	}

	err = answerAttempt.ValidateAnswersAttempt()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    err.Error(),
		})
		return
	}

	answersAttemptResp, err := handler.service.FrontendService.PostAnswersAttempt(
		userId, answerAttempt,
	)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    err.Error(),
		})
		return
	}

	webResponse := web.WebResponse{
		Status:  http.StatusOK,
		Message: "POST OK",
		Data:    answersAttemptResp,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *Handler) seeRecommendationByLevelId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	levelIdParam := params.ByName("levelId")
	levelId, err := strconv.Atoi(levelIdParam)
	helper.PanicIfError(err)

	recommendationResponse, err := handler.service.FrontendService.SeeRecommendationByLevelId(
		int32(levelId),
	)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    err.Error(),
		})
		return
	}

	webResponse := web.WebResponse{
		Status:  http.StatusOK,
		Message: "GET OK",
		Data:    recommendationResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
