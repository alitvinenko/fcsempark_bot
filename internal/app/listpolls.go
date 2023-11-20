package app

import (
	"context"
	"encoding/json"
	"log"
)

type ListPolls struct {
	serviceProvider *serviceProvider
}

func NewListPolls(ctx context.Context) *ListPolls {
	app := &ListPolls{}

	err := app.initServiceProvider(ctx)
	if err != nil {
		log.Fatalf("error on init service provider: %v", err)
	}

	return app
}

func (a ListPolls) Run(ctx context.Context) error {
	items, err := a.serviceProvider.getPollService().AllSavedItems(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Выводим результаты в консоль
	log.Println("All saved items:")
	for _, p := range items {
		str, err := json.Marshal(p)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Println(string(str))
	}

	return nil
}

func (a *ListPolls) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}
