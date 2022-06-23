package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/rg-km/final-project-engineering-14/backend/helper"
	"github.com/rg-km/final-project-engineering-14/backend/model/web"
)

func (handler *Handler) createAdminQuestion(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId, _ := getUserId(request)
	var questionRequest web.QuestionRequest
	err := json.NewDecoder(request.Body).Decode(&questionRequest)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}

	err = questionRequest.ValidateQuestion("create")
	if err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}

	questionResponse, err := handler.service.BackendService.CreateAdminQuestion(
		userId, questionRequest,
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
		Status:  http.StatusCreated,
		Message: "CREATE OK",
		Data:    questionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *Handler) getAdminQuestions(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	questionResponses, err := handler.service.BackendService.GetAdminQuestions()
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
		Message: http.StatusText(http.StatusOK),
		Data:    questionResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *Handler) getAdminDashboard(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	dashboardResponse, err := handler.service.BackendService.GetCountQuestions()
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
		Message: http.StatusText(http.StatusOK),
		Data:    dashboardResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *Handler) updateAdminQuestion(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var questionUpdateRequest web.QuestionRequest
	err := json.NewDecoder(request.Body).Decode(&questionUpdateRequest)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}

	questionIdParam := params.ByName("questionId")
	questionId, err := strconv.Atoi(questionIdParam)
	helper.PanicIfError(err)

	err = questionUpdateRequest.ValidateQuestion("update")
	if err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(writer).Encode(web.WebResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}

	_, err = handler.service.BackendService.UpdateAdminQuestion(
		questionUpdateRequest, int32(questionId),
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
		Message: "UPDATE OK",
		Data:    "Successfully updated question",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *Handler) deleteAdminQuestion(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	questionIdParam := params.ByName("questionId")
	questionId, err := strconv.Atoi(questionIdParam)
	helper.PanicIfError(err)

	_, err = handler.service.BackendService.DeleteAdminQuestion(int32(questionId))
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
		Message: "DELETE OK",
		Data:    "Successfully deleted the question",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
