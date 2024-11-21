package domain

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(100)"`
	Description string `json:"description" gorm:"size:255"`
	Level       string `json:"level" gorm:"type:varchar(50)"`
	GroupID     int
}
