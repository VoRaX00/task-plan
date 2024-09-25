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
		auth.POST("/sing-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.GET("/get", h.getUser)
			user.PUT("/update", h.updateUser)
		}

		task := api.Group("/task")
		{
			task.GET("/get:id", h.getTask)
			task.POST("/add", h.createTask)
		}

		group := api.Group("/group")
		{
			group.GET("/get:id", h.getGroup)
			group.POST("/add", h.createGroup)
		}
	}

	return router
}
