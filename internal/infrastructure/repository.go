package infrastructure

import "task-plan/internal/application"

type Repository struct {
	application.IAuthRepository
}

func NewRepository() *Repository {
	return &Repository{
		IAuthRepository: NewAuthRepository(),
	}
}
