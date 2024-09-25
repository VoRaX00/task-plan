package domain

type User struct {
	Name         string  `json:"name" db:"name"`
	Email        string  `json:"email" db:"email"`
	PasswordHash string  `json:"password" db:"password_hash"`
	Groups       []Group `json:"groups" db:"groups"`
}
