package json

import (
	"encoding/json"
	"errors"
	"github.com/alitvinenko/fcsempark_bot/internal/lib/e"
	"github.com/alitvinenko/fcsempark_bot/internal/lib/time"
	"github.com/alitvinenko/fcsempark_bot/internal/settings"
	"io"
	"os"
)

type GameSettings struct {
	Day        string `json:"day"`
	StartTime  string `json:"start_time"`
	Duration   string `json:"duration"`
	Place      string `json:"place"`
	MaxPlayers int    `json:"max_players"`
}

type Settings struct {
	Schedule []GameSettings `json:"schedule"`
}

type Storage struct {
	path string
}

func New(path string) *Storage {
	return &Storage{path: path}
}

func (c *Storage) Load() (s *settings.Settings, err error) {
	defer func() { err = e.WrapIfErr("invalid load games settings", err) }()

	file, err := os.Open(c.path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()

	byteValues, _ := io.ReadAll(file)

	var loadedSettings Settings
	if err := json.Unmarshal(byteValues, &loadedSettings); err != nil {
		return nil, err
	}

	if len(loadedSettings.Schedule) == 0 {
		return nil, errors.New("games settings not found")
	}

	var gamesSettings []settings.GameSettings
	for i := 0; i < len(loadedSettings.Schedule); i++ {
		weekDay, err := time.ParseWeekday(loadedSettings.Schedule[i].Day)
		if err != nil {
			return nil, errors.New("invalid Day value")
		}

		oneGameSettings := settings.GameSettings{
			Weekday:    weekDay,
			StartTime:  loadedSettings.Schedule[i].StartTime,
			Duration:   loadedSettings.Schedule[i].Duration,
			Place:      loadedSettings.Schedule[i].Place,
			MaxPlayers: loadedSettings.Schedule[i].MaxPlayers,
		}

		gamesSettings = append(gamesSettings, oneGameSettings)
	}

	return settings.New(gamesSettings), nil
}
