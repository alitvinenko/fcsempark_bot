package app

import (
	"context"
	"log"
)

type CreatePollApp struct {
	serviceProvider *serviceProvider
}

func NewCreatePollApp(ctx context.Context) *CreatePollApp {
	app := &CreatePollApp{}

	err := app.initServiceProvider(ctx)
	if err != nil {
		log.Fatalf("error on init service provider: %v", err)
	}

	return app
}

func (a CreatePollApp) Run(ctx context.Context) error {
	return a.serviceProvider.getPollService().CreateAndPin(ctx)
}

func (a *CreatePollApp) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}
