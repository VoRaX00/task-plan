package requestModels

type UserLogin struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password_hash"`
}

type UserToAdd struct {
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
