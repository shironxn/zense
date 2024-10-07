package web

import (
	"time"

	"github.com/aternity/zense/internal/entity/domain"
)

type ForumResponse struct {
	ID        uint              `json:"id"`
	UserID    uint              `json:"user_id"`
	Title     string            `json:"title"`
	Topic     domain.ForumTopic `json:"topic"`
	Content   string            `json:"content"`
	CreatedAt *time.Time        `json:"created_at,omitempty"`
	UpdatedAt *time.Time        `json:"updated_at,omitempty"`
}

type ForumCreate struct {
	UserID  uint              `validate:"required"`
	Title   string            `validate:"required"`
	Topic   domain.ForumTopic `validate:"required"`
	Content string            `validate:"required"`
}

type ForumFindByID struct {
	ID uint `param:"id" validate:"required"`
}

type ForumUpdate struct {
	ID      uint `param:"id" validate:"required"`
	UserID  uint `validate:"required"`
	Title   string
	Topic   domain.ForumTopic
	Content string
}

type ForumDelete struct {
	ID     uint `param:"id" validate:"required"`
	UserID uint `validate:"required"`
}
