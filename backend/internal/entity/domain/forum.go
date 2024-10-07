package domain

import "time"

type ForumTopic string

const (
	Topic1 ForumTopic = "topic1"
	Topic2 ForumTopic = "topic2"
	Topic3 ForumTopic = "topic3"
)

type Forum struct {
	ID        uint
	UserID    uint
	Title     string
	Topic     ForumTopic `gorm:"default:'topic1'" sql:"type:topic"`
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Comments  []Comment
}
