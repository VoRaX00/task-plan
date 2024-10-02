package infrastructure

import (
	"github.com/jmoiron/sqlx"
	"task-plan/internal/domain"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

func (r *AuthRepository) Create(user domain.User) (string, error) {
	return "", nil
}
func (r *AuthRepository) Get(user domain.User) (domain.User, error) {
	return domain.User{}, nil
}
func (r *AuthRepository) GetById(userId string) (domain.User, error) {
	return domain.User{}, nil
}
