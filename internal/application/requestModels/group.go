package requestModels

import (
	"github.com/google/uuid"
	"task-plan/internal/domain"
)

type (
	GroupToAdd struct {
		Name          string    `json:"name"`
		AdminUserId   uuid.UUID `json:"adminUserId"`
		LevelProgress []string  `json:"levelProgress"`
	}

	GroupToUpdate struct {
		Name          string        `json:"name"`
		LevelProgress []string      `json:"levelProgress"`
		Users         []domain.User `json:"users"`
		Tasks         []domain.Task `json:"tasks"`
	}

	GroupToDelete struct {
		Id uuid.UUID `json:"id"`
	}

	GroupToGet struct {
		domain.Group
	}
)
