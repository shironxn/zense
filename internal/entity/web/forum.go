package web

import (
	"time"
)

type ForumResponse struct {
	ID        uint            `json:"id"`
	UserID    uint            `json:"user_id,omitempty"`
	Title     string          `json:"title,omitempty"`
	Content   string          `json:"content,omitempty"`
	CreatedAt *time.Time      `json:"created_at,omitempty"`
	UpdatedAt *time.Time      `json:"updated_at,omitempty"`
	Topics    []TopicResponse `json:"topics,omitempty"`
	User      *UserResponse   `json:"user,omitempty"`
}

type ForumCreate struct {
	UserID  uint   `json:"user_id"`
	Title   string `validate:"required"`
	Topics  []uint `validate:"required"`
	Content string `validate:"required"`
}

type ForumFindByID struct {
	ID uint `param:"id"`
}

type ForumUpdate struct {
	ID      uint `param:"id"`
	UserID  uint `json:"user_id"`
	Title   string
	Topics  []uint
	Content string
}

type ForumDelete struct {
	ID     uint `param:"id"`
	UserID uint `json:"user_id"`
}

type ForumRemoveTopic struct {
	ID      uint `param:"id"`
	UserID  uint `json:"user_id"`
	TopicID uint `json:"topic_id" validate:"required"`
}
