package handler

import (
	"fmt"

	"github.com/julienschmidt/httprouter"
	"github.com/rg-km/final-project-engineering-14/backend/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		service: services,
	}
}

func (handler *Handler) InitRoutes() *httprouter.Router {
	router := httprouter.New()

	fmt.Printf("\n\t[INFO] Running in localhost:%d\n", 3000)

	router.POST("/auth/sign-up", handler.signUp)

	router.POST("/auth/sign-in", handler.signIn)

	router.POST("/auth/sign-out", handler.signOut)

	router.GET("/admin/dashboard",
		handler.AuthMiddleWare(
			handler.AdminMiddleWare(
				handler.getAdminDashboard,
			),
		),
	)

	router.GET("/admin/questions",
		handler.AuthMiddleWare(
			handler.AdminMiddleWare(
				handler.getAdminQuestions,
			),
		),
	)

	router.POST("/admin/questions/create",
		handler.AuthMiddleWare(
			handler.AdminMiddleWare(
				handler.createAdminQuestion,
			),
		),
	)

	router.PUT("/admin/questions/update/:questionId",
		handler.AuthMiddleWare(
			handler.AdminMiddleWare(
				handler.updateAdminQuestion,
			),
		),
	)

	router.DELETE("/admin/questions/delete/:questionId",
		handler.AuthMiddleWare(
			handler.AdminMiddleWare(
				handler.deleteAdminQuestion,
			),
		),
	)

	router.GET("/home/start-ruang-arah",
		handler.AuthMiddleWare(
			handler.getProgrammingLanguanges,
		),
	)

	router.GET("/home/start-ruang-arah/:proglangId",
		handler.AuthMiddleWare(
			handler.getQuestionByProgrammingLanguage,
		),
	)

	router.POST("/home/process-and-result",
		handler.AuthMiddleWare(
			handler.postAnswersAttempt,
		),
	)

	return router
}
