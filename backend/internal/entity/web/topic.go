package web

import "time"

type TopicResponse struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type TopicCreate struct {
	UserID      uint   `json:"user_id"`
	Name        string `validate:"required"`
	Description string `validate:"required"`
}

type TopicFindByID struct {
	ID uint `param:"id"`
}

type TopicUpdate struct {
	ID          uint `param:"id"`
	Name        string
	Description string
}

type TopicDelete struct {
	ID uint `param:"id"`
}
