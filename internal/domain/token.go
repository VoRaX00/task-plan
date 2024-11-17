package domain

import "time"

type Token struct {
	RefreshTokenHash string
	Ip               string
	UserToken        User `gorm:"foreignkey:UserID"`
	CreatedAt        time.Time
	ExpiresAt        time.Time
}
