package mapper

import (
	"task-plan/internal/application"
	"task-plan/internal/domain"
)

type UserMapper struct {
}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (m *UserMapper) UserAddToUser(user application.UserToAdd) domain.User {
	res := domain.User{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.Password,
	}
	return res
}
