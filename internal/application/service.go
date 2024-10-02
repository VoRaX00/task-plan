package application

import "task-plan/internal/infrastructure"

type Service struct {
	IAuthService
}

func NewService(repos *infrastructure.Repository) *Service {
	return &Service{
		IAuthService: NewAuthService(repos.IAuthRepository),
	}
}
