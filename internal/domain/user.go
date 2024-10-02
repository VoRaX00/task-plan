package domain

import "github.com/google/uuid"

type User struct {
	Id           uuid.UUID `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"password" db:"password_hash"`
	Groups       []Group   `json:"groups" db:"groups"`
}
