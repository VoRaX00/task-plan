package application

import (
	"task-plan/internal/application/requestModels"
	"task-plan/internal/domain"
	"task-plan/internal/infrastructure"
	"task-plan/pkg/mapper"
)

type AuthService struct {
	mapper mapper.Mapper
	repo   infrastructure.IAuthRepository
}

func NewAuthService(repo infrastructure.IAuthRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Create(userToAdd requestModels.UserToAdd) (string, error) {
	user := s.mapper.UserAddToUser(userToAdd)
	return s.repo.Create(user)
}

func (s *AuthService) GenerateToken(user domain.User) (string, error) {
	return "", nil
}

func (s *AuthService) GenerateEmailConfirmationToken(id string) (string, error) {
	return "", nil
}

func (s *AuthService) ParseToken(token string) (*domain.User, error) {
	return nil, nil
}
