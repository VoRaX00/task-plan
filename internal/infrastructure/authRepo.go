package infrastructure

import (
	"github.com/google/uuid"
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
	var id uuid.UUID
	query := "INSERT INTO users (id, email, name, password_hash) VALUES ($1, $2, $3, $4) RETURNING id"
	row := r.db.QueryRow(query, user.Id, user.Email, user.Name, user.PasswordHash)
	err := row.Scan(&id)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (r *AuthRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	query := "SELECT id, email, name FROM users WHERE email = $1"
	row := r.db.QueryRow(query, email)

	err := row.Scan(&user.Id, &user.Email, &user.Name)
	return domain.User{}, err
}

func (r *AuthRepository) GetById(userId uuid.UUID) (domain.User, error) {
	var user domain.User
	query := "SELECT id, email, name FROM users WHERE id = $1"
	row := r.db.QueryRow(query, userId)

	err := row.Scan(&user.Id, &user.Email, &user.Name)
	return domain.User{}, err
}
