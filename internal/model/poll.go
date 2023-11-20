package model

import "time"

type PollStatus string

const (
	PollStatusActive PollStatus = "active"
	PollStatusClosed PollStatus = "closed"
)

type Poll struct {
	ID                string     `json:"id"`
	TelegramMessageID int        `json:"telegram_message_id"`
	Date              time.Time  `json:"date"`
	Status            PollStatus `json:"status"`
	MaxPlayers        int        `json:"max_players"`
}
