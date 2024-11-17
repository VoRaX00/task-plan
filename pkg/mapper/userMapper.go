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

func (m *UserMapper) UserRegisterToUser(user requestModels.UserRegister) domain.User {
	res := domain.User{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.Password,
	}
	return res
}

func (m *UserMapper) UserToUserGet(user domain.User) requestModels.UserToGet {
	return requestModels.UserToGet{
		Name:  user.Name,
		Email: user.Email,
	}
}
