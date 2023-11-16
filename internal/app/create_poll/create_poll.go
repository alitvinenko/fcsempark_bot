package create_poll

import (
	"github.com/alitvinenko/fcsempark_bot/internal/config"
	"github.com/alitvinenko/fcsempark_bot/internal/database"
	"github.com/alitvinenko/fcsempark_bot/internal/polls/managers"
	"github.com/alitvinenko/fcsempark_bot/internal/polls/repository"
	tele "gopkg.in/telebot.v3"
	"log"
)

func Run() {
	if err := config.InitConfig(); err != nil {
		log.Fatalf("Error on init config: %v", err)
	}

	log.Printf("TOKEN=%s", config.GetToken())
	log.Printf("CHAT_ID=%s", config.GetChatID())
	log.Printf("DATABASE_PATH=%s", config.GetDatabasePath())
	log.Printf("SCHEDULE=%+v", config.GetSchedule())

	db := database.InitDB(config.GetDatabasePath())
	defer func() { _ = db.Close() }()

	bot := initBot(config.GetToken())

	pollManager := managers.NewCreatePollManager(config.GetChatID(), config.GetSchedule(), bot, repository.NewPollRepository(db))
	err := pollManager.CreateAndPin()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func initBot(token string) *tele.Bot {
	b, err := tele.NewBot(tele.Settings{Token: token})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("bot initialized")

	return b
}
