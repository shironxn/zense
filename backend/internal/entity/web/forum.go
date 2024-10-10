package web

import (
	"time"

	"github.com/aternity/zense/internal/entity/domain"
)

type ForumResponse struct {
	ID        uint              `json:"id"`
	UserID    uint              `json:"user_id,omitempty"`
	Title     string            `json:"title,omitempty"`
	Topic     domain.ForumTopic `json:"topic,omitempty"`
	Content   string            `json:"content,omitempty"`
	CreatedAt *time.Time        `json:"created_at,omitempty"`
	UpdatedAt *time.Time        `json:"updated_at,omitempty"`
	User      *UserResponse     `json:"user,omitempty"`
}

type ForumCreate struct {
	UserID  uint              `json:"user_id"`
	Title   string            `validate:"required"`
	Topic   domain.ForumTopic `validate:"required"`
	Content string            `validate:"required"`
}

type ForumFindByID struct {
	ID uint `param:"id"`
}

type ForumUpdate struct {
	ID      uint `param:"id"`
	UserID  uint `json:"user_id"`
	Title   string
	Topic   domain.ForumTopic
	Content string
}

type ForumDelete struct {
	ID     uint `param:"id"`
	UserID uint `json:"user_id"`
}
