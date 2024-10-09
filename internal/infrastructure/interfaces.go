package infrastructure

import (
	"github.com/google/uuid"
	"task-plan/internal/domain"
)

type IAuthRepository interface {
	Create(user domain.User) (string, error)
	GetByEmail(email string) (domain.User, error)
	GetById(userId uuid.UUID) (domain.User, error)
}
