package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"
	"time"
)
import "gopkg.in/yaml.v2"

type ScheduleItem struct {
	Day        time.Weekday `yaml:"day"`
	StartTime  string       `yaml:"start_time"`
	Duration   string       `yaml:"duration"`
	Place      string       `yaml:"place"`
	MaxPlayers int          `yaml:"max_players"`
}

type AppConfig struct {
	ChatID       int    `envconfig:"CHAT_ID"`
	Token        string `envconfig:"TOKEN"`
	DatabasePath string `envconfig:"DATABASE_PATH"`

	ScheduleItems []ScheduleItem `yaml:"schedule"`
}

func NewAppConfig() (*AppConfig, error) {
	f, err := os.Open("./config/config.yml")
	if err != nil {
		return nil, fmt.Errorf("error on reading config file: %v", err)
	}

	var appConfig *AppConfig

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&appConfig)
	if err != nil {
		return nil, fmt.Errorf("error on parse config file: %v", err)
	}

	err = envconfig.Process("", appConfig)
	if err != nil {
		return nil, fmt.Errorf("error on parse env variables: %v", err)
	}

	return appConfig, nil
}

func (c *AppConfig) GetChatID() int {
	return c.ChatID
}

func (c *AppConfig) GetToken() string {
	return c.Token
}

func (c *AppConfig) GetDatabasePath() string {
	return c.DatabasePath
}

func (c *AppConfig) GetSchedule() []ScheduleItem {
	return c.ScheduleItems
}

func processError(err error) {
	log.Println(err)
	os.Exit(2)
}
