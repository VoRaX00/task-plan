package application

import (
	"task-plan/internal/application/requestModels"
	"task-plan/internal/domain"
)

type IAuthService interface {
	Create(userToAdd requestModels.UserToAdd) (string, error)
	GenerateToken(user domain.User) (string, error)
	GenerateEmailConfirmationToken(id string) (string, error)
	ParseToken(token string) (*domain.User, error)
}

type IMessageEmailService interface {
	Send(emailMessage string) error
	SendConfirmEmail(email, callbackUrl string) error
}
