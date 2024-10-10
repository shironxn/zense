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
	User       *UserResponse            `json:"user,omitempty"`
}

type JournalCreate struct {
	UserID     uint               `json:"user_id"`
	Mood       domain.JournalMood `validate:"required"`
	Content    string             `validate:"required"`
	Visibility domain.JournalVisibility
}

type JournalFindByID struct {
	ID uint `param:"id"`
}

type JournalUpdate struct {
	ID         uint `param:"id"`
	UserID     uint `json:"user_id"`
	Mood       domain.JournalMood
	Content    string
	Visibility domain.JournalVisibility
}

type JournalDelete struct {
	ID     uint `param:"id"`
	UserID uint `json:"user_id"`
}
