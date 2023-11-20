package model

import "time"

type PollStatus string

const (
	PollStatusActive PollStatus = "active"
	PollStatusClosed PollStatus = "closed"
)

type Poll struct {
	ID                string `gorm:"primaryKey"`
	TelegramMessageID int
	Date              time.Time
	Status            PollStatus
	MaxPlayers        int
}
