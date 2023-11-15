// Выбирает дату ближайшей игры из настроек и создает опрос в ТГ.
// Если опрос на эту дату уже создан, ничего не делает.
//
// Запускать нужно по расписанию в момент когда нужно создавать опрос на следующую игру
package main

import (
	"github.com/alitvinenko/fcsempark_bot/internal/polls/managers"
	"github.com/alitvinenko/fcsempark_bot/internal/polls/repository"
	"github.com/alitvinenko/fcsempark_bot/internal/settings"
	"github.com/alitvinenko/fcsempark_bot/internal/settings/storage/json"
	"github.com/boltdb/bolt"
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
	"strconv"
)

const settingsPath = "data/settings.json"
const dbFilePath = "data/db.json"

func main() {
	db := initDB(dbFilePath)
	defer func() { _ = db.Close() }()

	gameSettings := loadGamesSettings(settings.NewLoader(json.New(settingsPath)))

	chatID, err := strconv.Atoi(os.Getenv("CHAT_ID"))
	if err != nil {
		log.Fatal("Invalid chatID")
	}
	log.Println("settings initialized")

	bot := initBot()

	pollManager := managers.NewCreatePollManager(chatID, gameSettings.GamesSettings, bot, repository.NewPollRepository(db))
	err = pollManager.CreateAndPin()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func initBot() *tele.Bot {
	pref := tele.Settings{
		Token: os.Getenv("TOKEN"),
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("bot initialized")

	return b
}

func loadGamesSettings(loader *settings.Loader) *settings.Settings {
	s, err := loader.Load()
	if err != nil {
		log.Fatal(err)
	}

	return s
}

func initDB(path string) *bolt.DB {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.PollsBucket))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
