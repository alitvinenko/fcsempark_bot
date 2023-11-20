package database

import (
	"github.com/alitvinenko/fcsempark_bot/internal/repository/poll/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func Init(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Fatalf("error on init DB: %v", err)
	}

	err = db.AutoMigrate(&model.Poll{})
	if err != nil {
		log.Fatalf("error on automigrate: %v", err)
	}

	return db
}
