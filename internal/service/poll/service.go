package poll

import (
	"github.com/alitvinenko/fcsempark_bot/internal/config"
	"github.com/alitvinenko/fcsempark_bot/internal/repository"
	def "github.com/alitvinenko/fcsempark_bot/internal/service"
	tele "gopkg.in/telebot.v3"
)

var _ def.PollService = (*service)(nil)

type service struct {
	pollRepository repository.PollRepository
	tgBot          *tele.Bot

	chatID        int
	scheduleItems []config.ScheduleItem
}

func NewService(pRepo repository.PollRepository, tgBot *tele.Bot, chatID int, scheduleItems []config.ScheduleItem) *service {
	return &service{pollRepository: pRepo, tgBot: tgBot, chatID: chatID, scheduleItems: scheduleItems}
}
