package application

import (
	"crypto/sha256"
	"fmt"
	"os"
	"task-plan/internal/application/requestModels"
	"task-plan/internal/domain"
	"task-plan/internal/infrastructure"
	"task-plan/pkg/mapper"
	"task-plan/pkg/tokenManager"
	"time"
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

const (
	accessTokenTTL  = time.Hour * 24
	refreshTokenTTL = time.Hour * 48
)

func (s *AuthService) GenerateEmailConfirmationToken(id string, manager *tokenManager.Manager) (string, error) {
	res, err := manager.GenerateEmailConfirmationToken(id)
	return res, err
}

func (s *AuthService) ParseToken(token string) (*domain.User, error) {
	return nil, nil
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}
