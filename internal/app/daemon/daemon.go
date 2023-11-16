package daemon

import (
	"github.com/alitvinenko/fcsempark_bot/internal/config"
	"github.com/alitvinenko/fcsempark_bot/internal/database"
	"github.com/alitvinenko/fcsempark_bot/internal/polls/managers"
	"github.com/alitvinenko/fcsempark_bot/internal/polls/repository"
	"github.com/alitvinenko/fcsempark_bot/internal/telegram/handlers"
	"github.com/alitvinenko/fcsempark_bot/internal/telegram/handlers/commands"
	tele "gopkg.in/telebot.v3"
	"log"
	"time"
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
	setBotHandlers(bot, config.GetChatID(), repository.NewPollRepository(db))
	bot.Start()
}

func initBot(token string) *tele.Bot {
	pref := tele.Settings{
		Token: token,
		Poller: &tele.LongPoller{
			Timeout:        10 * time.Second,
			AllowedUpdates: []string{"message", "poll", "poll_answer", "callback_query"},
		},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("bot initialized")

	return b
}

func setBotHandlers(b *tele.Bot, chatID int, r *repository.PollRepository) {
	b.Handle(tele.OnPoll, handlers.NewPollHandler(managers.NewClosePollManager(chatID, b, r)).Handle)

	b.Handle("/help", commands.HelpHandler)
	b.Handle("/stat", commands.StatHandler)
	b.Handle("/rules", commands.RulesHandler)

	// TODO: handle removing message with poll
}
