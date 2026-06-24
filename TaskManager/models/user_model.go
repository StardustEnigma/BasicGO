package models

import "time"

type User struct {
	UserId    int    `json:"user_id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}