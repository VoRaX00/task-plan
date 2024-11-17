package infrastructure

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"task-plan/internal/domain"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) Create(user domain.User) (uuid.UUID, error) {
	user.ID = uuid.New()
	result := r.db.Create(&user)
	if result.Error != nil {
		return uuid.Nil, result.Error
	}
	return user.ID, nil
}

func (r *AuthRepository) CreateToken(token domain.Token) error {
	tx := r.db.Create(&token)
	return tx.Error
}

func (r *AuthRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	tx := r.db.First(&user, "email = ?", email)
	if tx.Error != nil {
		return domain.User{}, tx.Error
	}

	return user, nil
}

func (r *AuthRepository) GetById(userId uuid.UUID) (domain.User, error) {
	var user domain.User
	tx := r.db.First(&user, "id = ?", userId)
	if tx.Error != nil {
		return domain.User{}, tx.Error
	}
	return user, nil
}
