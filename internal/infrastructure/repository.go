package infrastructure

import (
	"gorm.io/gorm"
	"task-plan/internal/application"
)

type Repository struct {
	application.IAuthRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		IAuthRepository: NewAuthRepository(db),
	}
}
