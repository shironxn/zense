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
	Visibility CommentVisibility `gorm:"type:enum('review', 'public', 'private')"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
