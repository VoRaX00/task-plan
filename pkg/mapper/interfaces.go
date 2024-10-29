package mapper

import (
	"task-plan/internal/application/requestModels"
	"task-plan/internal/domain"
)

type IUserMapper interface {
	UserAddToUser(user requestModels.UserToAdd) domain.User
	UserToUserAdd(user domain.User) requestModels.UserToAdd
}
