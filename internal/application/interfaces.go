package application

import "task-plan/internal/domain"

type IAuthService interface {
	Create(user domain.User) (string, error)
	GenerateToken(user domain.User) (string, error)
	ParseToken(token string) (*domain.User, error)
}
