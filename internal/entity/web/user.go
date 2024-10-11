package web

import "time"

type UserResponse struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name,omitempty"`
	Email     string     `json:"email,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type UserAuth struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

type UserRegister struct {
	Name     string `validate:"required,min=4,max=16"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=32"`
}

type UserLogin struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type UserFindMe struct {
	ID uint
}

type UserFindByID struct {
	ID uint `param:"id"`
}

type UserUpdate struct {
	ID       uint   `param:"id"`
	UserID   uint   `json:"user_id"`
	Name     string `validate:"max=16"`
	Email    string `validate:"omitempty,email"`
	Password string `validate:"max=32"`
}

type UserDelete struct {
	ID     uint `param:"id"`
	UserID uint `json:"user_id"`
}
