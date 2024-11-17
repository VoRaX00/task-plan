package requestModels

import (
	"task-plan/internal/domain"
)

type (
	GroupToAdd struct {
		Name          string      `json:"name"`
		AdminUser     domain.User `json:"adminUser" gorm:"forgeinkey:UserID"`
		LevelProgress []string    `json:"levelProgress"`
	}

	GroupToUpdate struct {
		Name          string        `json:"name"`
		LevelProgress []string      `json:"levelProgress"`
		Users         []domain.User `json:"users"`
		Tasks         []domain.Task `json:"tasks"`
	}

	GroupToDelete struct {
		Id int `json:"id"`
	}

	GroupToGet struct {
		domain.Group
	}
)
