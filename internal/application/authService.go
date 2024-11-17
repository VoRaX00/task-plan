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

const (
	accessTokenTTL  = time.Hour * 24
	refreshTokenTTL = time.Hour * 48
)

type AuthService struct {
	repo    IAuthRepository
	mapper  *mapper.Mapper
	manager *tokenManager.Manager
}

func NewAuthService(repo IAuthRepository) *AuthService {
	signingKey := os.Getenv("SIGNING_KEY")
	return &AuthService{
		repo:    repo,
		mapper:  mapper.NewMapper(),
		manager: tokenManager.NewManager(signingKey),
	}
}

func (s *AuthService) Create(userToAdd requestModels.UserRegister) (uuid.UUID, error) {
	user := s.mapper.UserRegisterToUser(userToAdd)
	return s.repo.Create(user)
}

func (s *AuthService) GetById(id uuid.UUID) (requestModels.UserToGet, error) {
	user, err := s.repo.GetById(id)
	if err != nil {
		return requestModels.UserToGet{}, err
	}
	req := s.mapper.UserToUserGet(user)
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

func (s *AuthService) GenerateTokens(ip string) (map[string]string, error) {
	tokens := map[string]string{}
	access, err := s.manager.NewAccessToken(ip, accessTokenTTL)
	if err != nil {
		return nil, err
	}

	refresh, err := s.manager.NewRefreshToken()
	if err != nil {
		return nil, err
	}
	tokens["accessToken"] = access
	tokens["refreshToken"] = refresh
	return tokens, nil
}

func (s *AuthService) GenerateEmailConfirmationToken(id string) (string, error) {
	res, err := s.manager.GenerateEmailConfirmationToken(id)
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
