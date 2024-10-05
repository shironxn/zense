package web

import "time"

type UserResponse struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type UserAuth struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

type UserRegister struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type UserLogin struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type UserFindByID struct {
	ID uint `param:"id" validate:"required"`
}

type UserUpdate struct {
	ID       uint `param:"id" validate:"required"`
	Name     string
	Email    string `validate:"email,omitempty"`
	Password string
}

type UserDelete struct {
	ID uint `param:"id" validate:"required"`
}
