package handler

import (
	"github.com/gin-gonic/gin"
	"task-plan/internal/application"
)

type Handler struct {
	services *application.Service
}

func NewHandler(services *application.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.SignUp)
		auth.POST("sign-in", h.SignIn)
	}

	api := router.Group("/api")
	{
		api.GET("/tasks", h.GetTasks)
		api.POST("/create-task", h.CreateTask)
	}

	return router
}
