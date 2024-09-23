package handler

import "task-plan/internal/application"

type Handler struct {
	services *application.Service
}

func NewHandler(services *application.Service) *Handler {
	return &Handler{
		services: services,
	}
}
