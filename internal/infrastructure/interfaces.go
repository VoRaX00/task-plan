package infrastructure

import "task-plan/internal/domain"

type IAuthRepository interface {
	Create(user domain.User) (string, error)
	Get(user domain.User) (domain.User, error)
	GetById(userId string) (domain.User, error)
}
