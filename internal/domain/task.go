package domain

type Task struct {
	Name        string        `json:"name" db:"name"`
	Description string        `json:"description" db:"description"`
	Level       LevelProgress `json:"level" db:"level"`
}
