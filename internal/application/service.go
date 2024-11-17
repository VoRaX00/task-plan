package application

import "task-plan/internal/infrastructure"

type Service struct {
	IAuthService
	IGroupService
	ITaskService
	IMessageEmailService
}

func NewService(repos *infrastructure.Repository) *Service {
	return &Service{
		IAuthService:         NewAuthService(repos.IAuthRepository),
		IGroupService:        NewGroupService(repos.IGroupRepository),
		ITaskService:         NewTaskService(repos.ITaskRepository),
		IMessageEmailService: NewMessageEmailService(),
	}
}
