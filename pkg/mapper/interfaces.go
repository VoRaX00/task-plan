package mapper

import (
	"task-plan/internal/application/requestModels"
	"task-plan/internal/domain"
)

type IUserMapper interface {
	UserRegisterToUser(user requestModels.UserRegister) domain.User
	UserToUserGet(user domain.User) requestModels.UserToGet
}
