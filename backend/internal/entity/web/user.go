package web

import "time"

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserCreate struct {
	Name     uint   `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type UserUpdate struct {
	ID       uint `validate:"required"`
	Name     string
	Email    string `validate:"email,omitempty"`
	Password string
}

type UserDelete struct {
	ID uint `validate:"required"`
}
