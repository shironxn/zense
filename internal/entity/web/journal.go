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
	UserID     uint                     `json:"user_id"`
	Mood       domain.JournalMood       `validate:"required,oneof=happy good normal sad angry"`
	Content    string                   `validate:"required"`
	Visibility domain.JournalVisibility `validate:"required,oneof=private public"`
}

type JournalFindByID struct {
	ID uint `param:"id"`
}

type JournalUpdate struct {
	ID         uint               `param:"id"`
	UserID     uint               `json:"user_id"`
	Mood       domain.JournalMood `validate:"omitempty,oneof=happy good normal sad angry"`
	Content    string
	Visibility domain.JournalVisibility `omitempty,validate:"oneof=private public"`
}

type JournalDelete struct {
	ID     uint `param:"id"`
	UserID uint `json:"user_id"`
}
