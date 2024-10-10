package web

import (
	"time"

	"github.com/aternity/zense/internal/entity/domain"
)

type CommentResponse struct {
	ID         uint                     `json:"id"`
	ForumID    uint                     `json:"forum_id"`
	UserID     uint                     `json:"user_id,omitempty"`
	Content    string                   `json:"content,omitempty"`
	Visibility domain.CommentVisibility `json:"comment,omitempty"`
	CreatedAt  *time.Time               `json:"created_at,omitempty"`
	UpdatedAt  *time.Time               `json:"updated_at,omitempty"`
	User       *UserResponse            `json:"user,omitempty"`
}

type CommentFindByID struct {
	ID uint `param:"id" validate:"required"`
}

type CommentCreate struct {
	UserID     uint   `validate:"required"`
	ForumID    uint   `validate:"required"`
	Content    string `validate:"required"`
	Visibility domain.CommentVisibility
}

type CommentUpdate struct {
	ID         uint `param:"id" validate:"required"`
	UserID     uint `validate:"required"`
	Content    string
	Visibility domain.CommentVisibility
}

type CommentDelete struct {
	ID     uint `param:"id" validate:"required"`
	UserID uint `validate:"required"`
}
