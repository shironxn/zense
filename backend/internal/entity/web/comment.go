package web

import (
	"time"

	"github.com/aternity/zense/internal/entity/domain"
)

type CommentResponse struct {
	ID         uint                     `json:"id"`
	UserID     uint                     `json:"user_id"`
	ForumID    uint                     `json:"forum_id"`
	Content    string                   `json:"content"`
	Visibility domain.CommentVisibility `json:"comment"`
	CreatedAt  time.Time                `json:"created_at"`
	UpdatedAt  time.Time                `json:"updated_at"`
}

type CommentCreate struct {
	UserID     uint                     `validate:"required"`
	ForumID    uint                     `validate:"required"`
	Content    string                   `validate:"required"`
	Visibility domain.CommentVisibility `validate:"required"`
}

type CommentUpdate struct {
	ID         uint `validate:"required"`
	UserID     uint `validate:"required"`
	Content    string
	Visibility domain.CommentVisibility
}

type CommentDelete struct {
	ID     uint `validate:"required"`
	UserID uint `validate:"required"`
}
