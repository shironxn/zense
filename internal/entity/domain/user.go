package domain

import "time"

type User struct {
	ID        uint
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Journals  []Journal
	Forums    []Forum
	Comments  []Comment
}
