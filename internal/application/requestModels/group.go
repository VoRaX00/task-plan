package requestModels

import (
	"github.com/google/uuid"
	"task-plan/internal/domain"
)

type GroupToAdd struct {
	Name          string    `json:"name" db:"name"`
	AdminUserId   uuid.UUID `json:"adminUserId" db:"user_id"`
	LevelProgress []string  `json:"levelProgress" db:"level_progress"`
}

type GroupToUpdate struct {
	Name          string        `json:"name" db:"name"`
	LevelProgress []string      `json:"levelProgress" db:"level_progress"`
	Users         []domain.User `json:"users" db:"users"`
	Tasks         []domain.Task `json:"tasks" db:"tasks"`
}

type GroupToDelete struct {
	Id uuid.UUID `json:"id" db:"id"`
}
