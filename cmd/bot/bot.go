// Обработчик команд бота
// Реагирует на определенные команды и события связанные с ответами на опрос
package main

import (
	"github.com/alitvinenko/fcsempark_bot/internal/polls/managers"
	"github.com/alitvinenko/fcsempark_bot/internal/polls/repository"
	"github.com/alitvinenko/fcsempark_bot/internal/telegram/handlers"
	"github.com/alitvinenko/fcsempark_bot/internal/telegram/handlers/commands"
	"github.com/boltdb/bolt"
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
	"strconv"
	"time"
)

const dbFilePath = "data/db.json"

func main() {
	db := initDB(dbFilePath)
	defer func() { _ = db.Close() }()

	chatID, err := strconv.Atoi(os.Getenv("CHAT_ID"))
	if err != nil {
		log.Fatal("Invalid chatID")
	}
	log.Println("settings initialized")

	bot := initBot(chatID, repository.NewPollRepository(db))
	bot.Start()
}

func initBot(chatID int, r *repository.PollRepository) *tele.Bot {
	pref := tele.Settings{
		Token: os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{
			Timeout:        10 * time.Second,
			AllowedUpdates: []string{"message", "poll", "poll_answer", "callback_query"},
		},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	b.Handle(tele.OnPoll, handlers.NewPollHandler(managers.NewClosePollManager(chatID, b, r)).Handle)

	b.Handle("/help", commands.HelpHandler)
	b.Handle("/stat", commands.StatHandler)
	b.Handle("/rules", commands.RulesHandler)

	// TODO: handle removing message with poll

	log.Println("bot initialized")

	return b
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
