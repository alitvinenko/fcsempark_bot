package handlers

import (
	"context"
	"github.com/alitvinenko/fcsempark_bot/internal/service"
	tele "gopkg.in/telebot.v3"
)

type PollHandler struct {
	service service.PollService
}

func NewPollHandler(service service.PollService) *PollHandler {
	return &PollHandler{service: service}
}

func (h *PollHandler) Handle(c tele.Context) error {
	if c.Poll().Closed {
		return nil
	}

	return h.service.CheckAndClose(context.Background(), c.Poll().ID, c.Poll().Options[0].VoterCount)
}
