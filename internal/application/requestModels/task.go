package requestModels

type TaskToAdd struct {
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Level       string `json:"level" db:"level"`
}

type TaskToUpdate struct {
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Level       string `json:"level" db:"level"`
}

type TaskToDelete struct {
	Id int64 `json:"id" db:"id"`
}
