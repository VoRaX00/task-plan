package infrastructure

import (
	"gorm.io/gorm"
	"task-plan/internal/application"
)

type Repository struct {
	application.IAuthRepository
	application.ITaskRepository
	application.IGroupRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		IAuthRepository:  NewAuthRepository(db),
		ITaskRepository:  NewTaskRepository(db),
		IGroupRepository: NewGroupRepository(db),
	}
}
