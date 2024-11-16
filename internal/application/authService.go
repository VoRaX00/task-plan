package application

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"os"
	"task-plan/internal/application/requestModels"
	"task-plan/internal/domain"
	"task-plan/pkg/mapper"
	"task-plan/pkg/tokenManager"
	"time"
)

type AuthService struct {
	mapper mapper.Mapper
	repo   IAuthRepository
}

func NewAuthService(repo IAuthRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Create(userToAdd requestModels.UserToAdd) (uuid.UUID, error) {
	user := s.mapper.UserAddToUser(userToAdd)
	return s.repo.Create(user)
}

func (s *AuthService) GetById(id uuid.UUID) (requestModels.UserToAdd, error) {
	user, err := s.repo.GetById(id)
	if err != nil {
		return requestModels.UserToAdd{}, err
	}
	req := s.mapper.UserToUserAdd(user)
	return req, nil
}

func (s *AuthService) CheckUser(user requestModels.UserLogin) error {
	getUser, err := s.repo.GetByEmail(user.Email)
	if err != nil {
		return err
	}
	passwordHash := s.generatePasswordHash(user.Password)
	if passwordHash == getUser.PasswordHash {
		return nil
	}
	return errors.New("wrong password")
}

func (s *AuthService) GenerateTokens(user requestModels.UserLogin) (map[string]string, error) {
	return nil, nil
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
