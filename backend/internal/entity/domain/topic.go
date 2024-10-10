package domain

import "time"

type Topic struct {
	ID          uint
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Forums      []Forum `gorm:"many2many:forum_topics"`
}
