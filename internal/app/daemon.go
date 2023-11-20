package app

import (
	"context"
	"github.com/alitvinenko/fcsempark_bot/internal/telegram/handlers/commands"
	tele "gopkg.in/telebot.v3"
	"log"
)

type Daemon struct {
	serviceProvider *serviceProvider
}

func NewDaemon(ctx context.Context) *Daemon {
	app := &Daemon{}

	err := app.initServiceProvider(ctx)
	if err != nil {
		log.Fatalf("error on init service provider: %v", err)
	}
	err = app.setBotHandlers()
	if err != nil {
		log.Fatalf("error on init bot commands handlers: %v", err)
	}

	return app
}

func (d Daemon) Run(_ context.Context) error {
	d.serviceProvider.getTgBot().Start()

	log.Println("bot started")

	return nil
}

func (a *Daemon) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *Daemon) setBotHandlers() error {

	b := a.serviceProvider.getTgBot()

	b.Handle(tele.OnPoll, a.serviceProvider.getPollHandler().Handle)

	b.Handle("/help", commands.HelpHandler)
	b.Handle("/stat", commands.StatHandler)
	b.Handle("/rules", commands.RulesHandler)

	// TODO: handle removing message with poll

	return nil
}
