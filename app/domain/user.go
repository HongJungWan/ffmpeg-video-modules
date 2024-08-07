package domain

import (
	"time"
)

type User struct {
	ID           int       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Username     string    `gorm:"type:varchar(255);not null" json:"username"`
	Email        string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	PasswordHash string    `gorm:"type:varchar(255);not null" json:"-"`
	CreatedAt    time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

func NewUser(username, email, passwordHash string) *User {
	return &User{
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (u *User) UpdatePassword(newPasswordHash string) {
	u.PasswordHash = newPasswordHash
	u.UpdatedAt = time.Now()
}
