package domain

type Task struct {
	Id          int64  `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Level       string `json:"level" db:"level"`
}
