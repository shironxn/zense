package domain

import "time"

type Forum struct {
	ID        uint
	UserID    uint
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
	Comments  []Comment
	Topics    []Topic `gorm:"many2many:forum_topics"`
}
