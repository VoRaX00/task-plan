package application

import (
	"github.com/google/uuid"
	"task-plan/internal/application/requestModels"
	"task-plan/internal/domain"
)

type IAuthService interface {
	Create(requestModels.UserRegister) (uuid.UUID, error)
	GetById(uuid.UUID) (requestModels.UserToGet, error)
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
	Create(requestModels.TaskToAdd) (int, error)
	GetById(id int) (requestModels.TaskToGet, error)
}

type IGroupService interface {
	Create(add requestModels.GroupToAdd) (int, error)
	GetById(id int) (requestModels.GroupToGet, error)
}
