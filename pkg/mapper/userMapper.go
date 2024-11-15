package mapper

import (
	"task-plan/internal/application/requestModels"
	"task-plan/internal/domain"
)

type UserMapper struct {
}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (m *UserMapper) UserAddToUser(user requestModels.UserToAdd) domain.User {
	res := domain.User{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.Password,
	}
	return res
}

func (m *UserMapper) UserToUserAdd(user domain.User) requestModels.UserToAdd {
	return requestModels.UserToAdd{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.PasswordHash,
	}
}
