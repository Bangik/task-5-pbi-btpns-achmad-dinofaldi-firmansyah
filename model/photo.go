package model

import "time"

type Photo struct {
	Id        string    `form:"id" json:"id" gorm:"primaryKey"`
	Title     string    `form:"title" json:"title" binding:"required,max=255"`
	Caption   string    `form:"caption" json:"caption" binding:"required,max=255"`
	Url       string    `json:"url" binding:"required,max=255"`
	UserId    string    `json:"user_id" binding:"required,max=255"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PhotoResponse struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Url       string    `json:"url"`
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
}
