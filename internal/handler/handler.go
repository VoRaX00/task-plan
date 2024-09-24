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

	return router
}
