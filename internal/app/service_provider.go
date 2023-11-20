package app

import (
	"github.com/alitvinenko/fcsempark_bot/internal/config"
	"github.com/alitvinenko/fcsempark_bot/internal/database"
	"github.com/alitvinenko/fcsempark_bot/internal/repository"
	"github.com/alitvinenko/fcsempark_bot/internal/repository/poll"
	"github.com/alitvinenko/fcsempark_bot/internal/service"
	pService "github.com/alitvinenko/fcsempark_bot/internal/service/poll"
	"github.com/alitvinenko/fcsempark_bot/internal/telegram/handlers"
	tele "gopkg.in/telebot.v3"
	"gorm.io/gorm"
	"log"
	"time"
)

type serviceProvider struct {
	config *config.AppConfig

	tgBot *tele.Bot

	db *gorm.DB

	pollRepository repository.PollRepository

	pollService service.PollService

	// handlers
	pollHandler *handlers.PollHandler
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) getConfig() *config.AppConfig {
	if s.config != nil {
		return s.config
	}

	appConfig, err := config.NewAppConfig()
	if err != nil {
		log.Fatalf("error on load appConfig: %s", err)
	}

	s.config = appConfig

	return s.config
}

func (s *serviceProvider) getTgBot() *tele.Bot {
	if s.tgBot != nil {
		return s.tgBot
	}

	pref := tele.Settings{
		Token: s.getConfig().GetToken(),
		Poller: &tele.LongPoller{
			Timeout:        10 * time.Second,
			AllowedUpdates: []string{"message", "poll", "poll_answer", "callback_query"},
		},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	s.tgBot = b

	return s.tgBot
}

func (s *serviceProvider) getDB() *gorm.DB {
	if s.db != nil {
		return s.db
	}

	s.db = database.Init(s.getConfig().GetDatabasePath())

	return s.db
}

func (s *serviceProvider) getPollRepository() repository.PollRepository {
	if s.pollRepository != nil {
		return s.pollRepository
	}

	s.pollRepository = poll.NewRepository(s.getDB())

	return s.pollRepository
}

func (s *serviceProvider) getPollService() service.PollService {
	if s.pollService != nil {
		return s.pollService
	}

	s.pollService = pService.NewService(s.getPollRepository(), s.getTgBot(), s.getConfig().GetChatID(), s.getConfig().GetSchedule())

	return s.pollService
}

func (s *serviceProvider) getPollHandler() *handlers.PollHandler {
	if s.pollHandler != nil {
		return s.pollHandler
	}

	s.pollHandler = handlers.NewPollHandler(s.getPollService())

	return s.pollHandler
}
