package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type ScheduleItem struct {
	Day        time.Weekday `mapstructure:"day"`
	StartTime  string       `mapstructure:"start_time"`
	Duration   string       `mapstructure:"duration"`
	Place      string       `mapstructure:"place"`
	MaxPlayers int          `mapstructure:"max_players"`
}

type AppConfig struct {
	ChatID       int    `mapstructure:"chat_id"`
	Token        string `mapstructure:"token"`
	DatabasePath string `mapstructure:"database_path"`

	ScheduleItems []ScheduleItem `mapstructure:"schedule"`
}

var appConfig AppConfig

func InitConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.AutomaticEnv()
	viper.BindEnv("TOKEN")
	viper.BindEnv("CHAT_ID")
	viper.BindEnv("DATABASE_PATH")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error on reading config file: %v", err)
	}
	if err := viper.Unmarshal(&appConfig); err != nil {
		return fmt.Errorf("error on parse config file: %v", err)
	}

	return nil
}

func GetChatID() int {
	return appConfig.ChatID
}

func GetToken() string {
	return appConfig.Token
}

func GetDatabasePath() string {
	return appConfig.DatabasePath
}

func GetSchedule() []ScheduleItem {
	return appConfig.ScheduleItems
}
