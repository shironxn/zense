package web

import (
	"time"

	"github.com/aternity/zense/internal/entity/domain"
)

type JournalResponse struct {
	ID         uint                     `json:"id"`
	UserID     uint                     `json:"user_id,omitempty"`
	Mood       domain.JournalMood       `json:"mood,omitempty"`
	Content    string                   `json:"content,omitempty"`
	Visibility domain.JournalVisibility `json:"visibility,omitempty"`
	CreatedAt  *time.Time               `json:"created_at,omitempty"`
	UpdatedAt  *time.Time               `json:"updated_at,omitempty"`
}

type JournalCreate struct {
	UserID     uint               `validate:"required"`
	Mood       domain.JournalMood `validate:"required"`
	Content    string             `validate:"required"`
	Visibility domain.JournalVisibility
}

type JournalFindByID struct {
	ID uint `param:"id" validate:"required"`
}

type JournalUpdate struct {
	ID         uint `param:"id" validate:"required"`
	UserID     uint `validate:"required"`
	Mood       domain.JournalMood
	Content    string
	Visibility domain.JournalVisibility
}

type JournalDelete struct {
	ID     uint `param:"id" validate:"required"`
	UserID uint `validate:"required"`
}
