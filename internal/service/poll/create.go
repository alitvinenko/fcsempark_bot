package poll

import (
	"context"
	"errors"
	"fmt"
	"github.com/alitvinenko/fcsempark_bot/internal/config"
	"github.com/alitvinenko/fcsempark_bot/internal/lib/e"
	time2 "github.com/alitvinenko/fcsempark_bot/internal/lib/time"
	"github.com/alitvinenko/fcsempark_bot/internal/model"
	tele "gopkg.in/telebot.v3"
	"log"
	"time"
)

var notFoundNextPollSettingsError = errors.New("next poll settings not found")

func (s *service) CreateAndPin(ctx context.Context) (err error) {
	defer func() { err = e.WrapIfErr("error on create poll", err) }()

	nextPollSettings, err := s.getNextPollSettings()
	if err != nil {
		return err
	}
	if nextPollSettings == nil {
		return notFoundNextPollSettingsError
	}

	nextGameDay := time2.StartOfDay(time2.NextWeekday(time.Now(), nextPollSettings.Day))

	existsNextGamePoll, err := s.isExistsNextGamePoll(ctx, nextGameDay)
	if err != nil {
		return err
	}

	if existsNextGamePoll {
		log.Printf("poll on %s already exists", nextGameDay.String())
		return nil
	}

	message, err := s.tgBot.Send(&tele.Chat{ID: int64(s.chatID)}, buildPoll(nextPollSettings))
	if err != nil {
		return err
	}

	err = s.pollRepository.Save(ctx, &model.Poll{
		ID:                message.Poll.ID,
		TelegramMessageID: message.ID,
		Date:              nextGameDay,
		Status:            model.PollStatusActive,
		MaxPlayers:        nextPollSettings.MaxPlayers,
	})
	if err != nil {
		log.Println(err)
	}

	err = s.tgBot.Pin(message, &tele.SendOptions{
		DisableNotification: false,
	})
	if err != nil {
		log.Println(err)
	}

	log.Println("new poll created")

	return nil
}

func (s *service) getNextPollSettings() (*config.ScheduleItem, error) {
	now := int(time.Now().Weekday())
	index := 0
	minPeriod := 8

	for i, v := range s.scheduleItems {
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

	return &s.scheduleItems[index], nil
}

func (s *service) isExistsNextGamePoll(ctx context.Context, t time.Time) (bool, error) {
	poll, err := s.pollRepository.GetByDate(ctx, t)
	if err != nil {
		return false, err
	}

	return poll != nil, nil
}

func buildPoll(scheduleItem *config.ScheduleItem) *tele.Poll {
	question := buildPollQuestion(scheduleItem)

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

func buildPollQuestion(s *config.ScheduleItem) string {
	question := fmt.Sprintf("%s, %s (%s), %s", time2.ToRus(s.Day), s.StartTime, s.Duration, s.Place)

	return question
}
