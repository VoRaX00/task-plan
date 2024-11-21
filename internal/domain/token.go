package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Token struct {
	gorm.Model
	RefreshTokenHash string
	Ip               string
	UserID           uuid.UUID `gorm:"foreignKey:ID"`
	CreatedAt        time.Time
	ExpiresAt        time.Time
}
