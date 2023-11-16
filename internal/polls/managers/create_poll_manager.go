package managers

import (
	"errors"
	"fmt"
	"github.com/alitvinenko/fcsempark_bot/internal/config"
	"github.com/alitvinenko/fcsempark_bot/internal/lib/e"
	time2 "github.com/alitvinenko/fcsempark_bot/internal/lib/time"
	"github.com/alitvinenko/fcsempark_bot/internal/polls/repository"
	tele "gopkg.in/telebot.v3"
	"log"
	"time"
)

type CreatePollManager struct {
	chatID        int
	scheduleItems []config.ScheduleItem
	bot           *tele.Bot
	repository    *repository.PollRepository
}

func NewCreatePollManager(chatID int, scheduleItems []config.ScheduleItem, bot *tele.Bot, repository *repository.PollRepository) *CreatePollManager {
	return &CreatePollManager{chatID: chatID, scheduleItems: scheduleItems, bot: bot, repository: repository}
}

var notFoundNextPollSettingsError = errors.New("next poll settings not found")

func (p *CreatePollManager) CreateAndPin() (err error) {
	defer func() { err = e.WrapIfErr("error on create poll", err) }()

	nextPollSettings, err := p.getNextPollSettings()
	if err != nil {
		return err
	}
	if nextPollSettings == nil {
		return notFoundNextPollSettingsError
	}

	nextGameDay := time2.StartOfDay(time2.NextWeekday(time.Now(), nextPollSettings.Day))

	existsNextGamePoll, err := p.repository.IsExists(nextGameDay)
	if err != nil {
		return err
	}

	if existsNextGamePoll {
		log.Printf("poll on %s already exists", nextGameDay.String())
		return nil
	}

	message, err := p.bot.Send(&tele.Chat{ID: int64(p.chatID)}, p.buildPoll(nextPollSettings))
	if err != nil {
		return err
	}

	dbPoll := repository.Poll{
		ID:          message.Poll.ID,
		TgMessageID: message.ID,
		Day:         nextGameDay,
		Status:      repository.PollStatusActive,
		MaxPlayers:  nextPollSettings.MaxPlayers,
	}
	err = p.repository.Save(dbPoll)
	if err != nil {
		log.Println(err)
	}

	err = p.bot.Pin(message)
	if err != nil {
		log.Println(err)
	}

	log.Println("new poll created")

	return nil
}

func (p *CreatePollManager) getNextPollSettings() (*config.ScheduleItem, error) {
	now := int(time.Now().Weekday())
	index := 0
	minPeriod := 8

	for i, v := range p.scheduleItems {
		vWeekDay := int(v.Day)

		var period int
		if vWeekDay > now {
			period = vWeekDay - now
		} else if vWeekDay < now {
			period = 6 - now + vWeekDay
		} else {
			period = 7
		}

		if minPeriod > period {
			minPeriod = period
			index = i
		}
	}

	return &p.scheduleItems[index], nil
}

func (p *CreatePollManager) buildPoll(scheduleItem *config.ScheduleItem) *tele.Poll {
	question := p.buildPollQuestion(scheduleItem)

	return &tele.Poll{
		Type:     "regular",
		Question: question,
		Options: []tele.PollOption{tele.PollOption{
			Text: "Готов",
		}, tele.PollOption{
			Text: "Думаю",
		}, tele.PollOption{
			Text: "Пропущу",
		}},
		Anonymous:       false,
		MultipleAnswers: false,
	}
}

func (p *CreatePollManager) buildPollQuestion(s *config.ScheduleItem) string {
	question := fmt.Sprintf("%s, %s (%s), %s", time2.ToRus(s.Day), s.StartTime, s.Duration, s.Place)

	return question
}
