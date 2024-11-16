package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	Name         string    `json:"name" gorm:"type:varchar(50)"`
	Email        string    `json:"email" gorm:"type:text;unique"`
	PasswordHash string    `json:"password" gorm:"type:text"`
	Groups       []Group   `json:"groups" gorm:"foreignkey:UserID"`
}
