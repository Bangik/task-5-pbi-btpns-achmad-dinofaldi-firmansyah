package model

import "time"

type User struct {
	Id       string `json:"id" gorm:"primaryKey"`
	Username string `json:"username" binding:"required,max=255"`
	Email    string `json:"email" binding:"required,max=255" gorm:"unique"`
	Password string `json:"password" binding:"required,min=6,max=255"`
}

type UserCredential struct {
	Email    string `json:"email" binding:"required,max=255"`
	Password string `json:"password" binding:"required,min=6,max=255"`
}

type UserResponse struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
