package di

import (
	"task-plan/internal/application"
	"task-plan/internal/infrastructure"
)

type Service struct {
	application.IAuthService
	application.IGroupService
	application.ITaskService
	application.IMessageEmailService
}

func NewService(repos *infrastructure.Repository) *Service {
	return &Service{
		IAuthService:         application.NewAuthService(repos.IAuthRepository),
		IGroupService:        application.NewGroupService(repos.IGroupRepository),
		ITaskService:         application.NewTaskService(repos.ITaskRepository),
		IMessageEmailService: application.NewMessageEmailService(),
	}
}
