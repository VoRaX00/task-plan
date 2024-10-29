package application

import (
	"github.com/google/uuid"
	"task-plan/internal/domain"
)

type IBaseRepository[T any, U any] interface {
	Create(model T) (U, error)
	GetById(id U) (T, error)
}

type IAuthRepository interface {
	//Create(user domain.User) (string, error)
	IBaseRepository[domain.User, uuid.UUID]
	GetByEmail(email string) (domain.User, error)
	//GetById(userId uuid.UUID) (domain.User, error)
}
