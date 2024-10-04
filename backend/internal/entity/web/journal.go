package web

import (
	"time"

	"github.com/aternity/zense/internal/entity/domain"
)

type JournalResponse struct {
	ID         uint                     `json:"id"`
	UserID     uint                     `json:"user_id"`
	Mood       domain.JournalMood       `json:"mood"`
	Content    string                   `json:"content"`
	Visibility domain.JournalVisibility `json:"visibility"`
	CreatedAt  time.Time                `json:"created_at"`
	UpdatedAt  time.Time                `json:"updated_at"`
}

type JournalCreate struct {
	UserID     uint                     `validate:"required"`
	Mood       domain.JournalMood       `validate:"required"`
	Content    string                   `validate:"required"`
	Visibility domain.JournalVisibility `validate:"required"`
}

type JournalUpdate struct {
	ID         uint `validate:"required"`
	UserID     uint `validate:"required"`
	Mood       domain.JournalMood
	Content    string
	Visibility domain.JournalVisibility
}

type JournalDelete struct {
	ID     uint `validate:"required"`
	UserID uint `validate:"required"`
}
