package application

import "task-plan/internal/infrastructure"

type Service struct {
	IAuthService
	IMessageEmailService
}

func NewService(repos *infrastructure.Repository) *Service {
	return &Service{
		IAuthService:         NewAuthService(repos.IAuthRepository),
		IMessageEmailService: NewMessageEmailService(),
	}
}
