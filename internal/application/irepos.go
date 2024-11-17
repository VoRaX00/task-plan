package application

import (
	"github.com/google/uuid"
	"task-plan/internal/domain"
)

type (
	IBaseRepository[T any, U any] interface {
		Create(T) (U, error)
		GetById(U) (T, error)
	}

	IAuthRepository interface {
		IBaseRepository[domain.User, uuid.UUID]
		GetByEmail(email string) (domain.User, error)
	}

	IGroupRepository interface {
	}

	ITaskRepository interface {
	}
)
