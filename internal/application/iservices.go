package application

import (
	"github.com/google/uuid"
	"task-plan/internal/application/requestModels"
	"task-plan/internal/domain"
)

type IBaseServices[T any, U any] interface {
	Create(T) (U, error)
	GetById(U) (T, error)
}

type IAuthService interface {
	IBaseServices[requestModels.UserToAdd, uuid.UUID]
	GenerateTokens(ip string) (map[string]string, error)
	GenerateEmailConfirmationToken(id string) (string, error)
	CheckUser(user requestModels.UserLogin) error
	ParseToken(token string) (*domain.User, error)
}

type IMessageEmailService interface {
	Send(emailMessage string) error
	SendConfirmEmail(email, callbackUrl string) error
}

type ITaskService interface {
	Create()
}
