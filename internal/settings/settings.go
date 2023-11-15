package settings

import (
	"time"
)

type GameSettings struct {
	Weekday    time.Weekday
	StartTime  string
	Duration   string
	Place      string
	MaxPlayers int
}

type Settings struct {
	GamesSettings []GameSettings
}

type Storage interface {
	Load() (*Settings, error)
}

type Loader struct {
	storage Storage
}

func NewLoader(client Storage) *Loader {
	return &Loader{storage: client}
}

func New(gamesSettings []GameSettings) *Settings {
	return &Settings{GamesSettings: gamesSettings}
}

func (l *Loader) Load() (*Settings, error) {
	return l.storage.Load()
}
