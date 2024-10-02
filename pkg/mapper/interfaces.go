package mapper

import (
	"task-plan/internal/application"
	"task-plan/internal/domain"
)

type IUserMapper interface {
	UserAddToUser(user application.UserToAdd) domain.User
}
