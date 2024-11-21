package domain

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name          string          `json:"name" gorm:"size:255"`
	Users         []User          `json:"users" gorm:"many2many:user_groups;"`
	Tasks         []Task          `json:"tasks" gorm:"foreignKey:GroupID"`
	LevelProgress []LevelProgress `json:"levelProgress" gorm:"many2many:group_level_progress;"`
}
