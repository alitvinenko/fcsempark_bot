package handlers

import (
	"encoding/json"
	"github.com/alitvinenko/fcsempark_bot/internal/polls/managers"
	tele "gopkg.in/telebot.v3"
	"log"
)

type PollAnswerHandler struct {
	m *managers.ClosePollManager
}

func NewPollAnswerHandler(m *managers.ClosePollManager) *PollAnswerHandler {
	return &PollAnswerHandler{m: m}
}

func (h *PollAnswerHandler) Handle(c tele.Context) error {

	//total := c.Message().Poll.VoterCount

	log.Println("It's poll answer handler")
	log.SetPrefix("[POLL ANSWER HANDLER] ")

	pollID := c.PollAnswer().PollID

	err := h.m.CheckAndClose(pollID, 0)
	if err != nil {
		log.Println(err)
	}

	answer, _ := json.Marshal(c.PollAnswer())
	log.Printf("%v", string(answer))
	log.Printf("%v", c.PollAnswer().Options)
	//return c.Send(fmt.Sprintf("Poll #%s. Total voters count : %d", c.Message().Poll.ID, total))

	return nil
}
