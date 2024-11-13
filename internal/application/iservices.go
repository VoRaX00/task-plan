package application

import (
	"github.com/google/uuid"
	"task-plan/internal/application/requestModels"
	"task-plan/internal/domain"
	"task-plan/pkg/tokenManager"
)

type IBaseServices[T any, U any] interface {
	Create(T) (U, error)
	GetById(U) (T, error)
}

type IAuthService interface {
	IBaseServices[requestModels.UserToAdd, uuid.UUID]
	GenerateToken(user domain.User) (string, error)
	GenerateEmailConfirmationToken(id string, manager *tokenManager.Manager) (string, error)
	ParseToken(token string) (*domain.User, error)
}

type IMessageEmailService interface {
	Send(emailMessage string) error
	SendConfirmEmail(email, callbackUrl string) error
}
