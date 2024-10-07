package domain

import "time"

type CommentVisibility string

const (
	ReviewComment  CommentVisibility = "review"
	PublicComment  CommentVisibility = "public"
	PrivateComment CommentVisibility = "private"
)

type Comment struct {
	ID         uint
	UserID     uint
	ForumID    uint
	Content    string
	Visibility CommentVisibility `gorm:"default:'review'" sql:"type:visibility"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
