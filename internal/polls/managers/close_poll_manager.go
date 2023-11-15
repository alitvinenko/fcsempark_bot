package managers

import (
	"encoding/json"
	"github.com/alitvinenko/fcsempark_bot/internal/polls/repository"
	tele "gopkg.in/telebot.v3"
	"log"
	"strconv"
)

type ClosePollManager struct {
	chatID     int
	bot        *tele.Bot
	repository *repository.PollRepository
}

func NewClosePollManager(chatID int, bot *tele.Bot, repository *repository.PollRepository) *ClosePollManager {
	return &ClosePollManager{chatID: chatID, bot: bot, repository: repository}
}

// CheckAndClose Проверяет нужно ли закрыть опрос и закрывает если нужно
// Вместе с закрытием, сообщение с опросом открепляется и меняется статус в БД
func (m *ClosePollManager) CheckAndClose(pollID string, rightVotedCount int) error {
	poll, err := m.repository.ByID(pollID)
	if err != nil {
		return err
	}
	if poll.ID == "" {
		log.Printf("poll %s not found in db. skipped", pollID)
		return nil
	}

	s, _ := json.Marshal(poll)
	log.Println(string(s))

	if rightVotedCount < poll.MaxPlayers {
		return nil
	}

	_, err = m.bot.StopPoll(&tele.StoredMessage{MessageID: strconv.Itoa(poll.TgMessageID), ChatID: int64(m.chatID)})
	if err != nil {
		return err
	}

	poll.Status = repository.PollStatusClosed

	err = m.repository.Save(poll)
	if err != nil {
		return err
	}

	recipient := &tele.Chat{ID: int64(m.chatID)}

	_ = m.bot.Unpin(recipient, poll.TgMessageID)

	options := &tele.SendOptions{ReplyTo: &tele.Message{ID: poll.TgMessageID}}
	_, err = m.bot.Send(recipient, "Отлично! Мы собрали достаточно игроков на следующий матч. Давайте делиться!", options)

	return nil
}
