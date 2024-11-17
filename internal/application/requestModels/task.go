package requestModels

import "task-plan/internal/domain"

type (
	TaskToAdd struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Level       string `json:"level"`
	}

	TaskToUpdate struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Level       string `json:"level"`
	}

	TaskToDelete struct {
		Id int64 `json:"id"`
	}

	TaskToGet struct {
		domain.Task
	}
)
