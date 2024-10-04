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
	Mood       JournalMood `gorm:"type:enum('happy','good','normal','sad','angry')"`
	Content    string
	Visibility JournalVisibility `gorm:"type:enum('private','public')"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
