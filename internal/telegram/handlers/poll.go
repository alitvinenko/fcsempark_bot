package handlers

import (
	"github.com/alitvinenko/fcsempark_bot/internal/polls/managers"
	tele "gopkg.in/telebot.v3"
)

type PollHandler struct {
	m *managers.ClosePollManager
}

func NewPollHandler(m *managers.ClosePollManager) *PollHandler {
	return &PollHandler{m: m}
}

func (h *PollHandler) Handle(c tele.Context) error {
	if c.Poll().Closed {
		return nil
	}

	return h.m.CheckAndClose(c.Poll().ID, c.Poll().Options[0].VoterCount)
}
