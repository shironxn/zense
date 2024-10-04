package web

import "time"

type ForumResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Title     string    `json:"title"`
	Topic     string    `json:"topic"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ForumCreate struct {
	UserID  uint   `validate:"required"`
	Title   string `validate:"required"`
	Topic   string `validate:"required"`
	Content string `validate:"content"`
}

type ForumUpdate struct {
	ID      uint `validate:"required"`
	UserID  uint `validate:"required"`
	Title   string
	Topic   string
	Content string
}

type ForumDelete struct {
	ID     uint `validate:"required"`
	UserID uint `validate:"required"`
}
