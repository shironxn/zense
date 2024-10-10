package domain

import "time"

type JournalMood string
type JournalVisibility string

const (
	Happy  JournalMood = "happy"
	Good   JournalMood = "good"
	Normal JournalMood = "normal"
	Sad    JournalMood = "sad"
	Angry  JournalMood = "angry"
)

const (
	PrivateJournal JournalVisibility = "private"
	PublicJournal  JournalVisibility = "public"
)

type Journal struct {
	ID         uint
	UserID     uint
	Mood       JournalMood `gorm:"default:'normal'" sql:"type:mood"`
	Content    string
	Visibility JournalVisibility `gorm:"default:'private'" sql:"type:visibility"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
  User User
}
