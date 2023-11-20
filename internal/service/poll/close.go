package poll

import (
	"context"
	"github.com/alitvinenko/fcsempark_bot/internal/model"
	tele "gopkg.in/telebot.v3"
	"log"
	"strconv"
)

func (s *service) CheckAndClose(ctx context.Context, pollID string, votedCount int) error {
	poll, err := s.pollRepository.Get(ctx, pollID)
	if err != nil {
		return err
	}
	if poll == nil {
		log.Printf("poll %s not found. skipped", pollID)
		return nil
	}

	if votedCount < poll.MaxPlayers {
		return nil
	}

	_, err = s.tgBot.StopPoll(&tele.StoredMessage{
		MessageID: strconv.Itoa(poll.TelegramMessageID),
		ChatID:    int64(s.chatID),
	})
	if err != nil {
		return err
	}

	poll.Status = model.PollStatusClosed
	err = s.pollRepository.Save(ctx, poll)
	if err != nil {
		return err
	}

	recipient := &tele.Chat{ID: int64(s.chatID)}

	_ = s.tgBot.Unpin(recipient, poll.TelegramMessageID)

	options := &tele.SendOptions{ReplyTo: &tele.Message{ID: poll.TelegramMessageID}}
	_, err = s.tgBot.Send(recipient, "Отлично! Мы собрали достаточно игроков на следующий матч. Давайте делиться!", options)

	return nil
}
