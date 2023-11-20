package main

import (
	"context"
	"fmt"
	"github.com/alitvinenko/fcsempark_bot/internal/app"
	"log"
	"os"
)

const (
	CommandDaemon     = "daemon"
	CommandCreatePoll = "createpoll"
	CommandShowDB     = "showdb"
)

func main() {
	var command string

	if len(os.Args) == 1 {
		// Если команда не указана, используем значение по умолчанию "daemon"
		command = CommandDaemon
	} else if len(os.Args) == 2 {
		command = os.Args[1]
	} else {
		fmt.Printf("Usage: go run main.go [%s|%s|%s]\n", CommandDaemon, CommandCreatePoll, CommandShowDB)
		os.Exit(1)
	}

	ctx := context.Background()
	var application app.App

	switch command {
	case CommandDaemon:
		application = app.NewDaemon(ctx)
	case CommandCreatePoll:
		application = app.NewCreatePollApp(ctx)
	case CommandShowDB:
		application = app.NewListPolls(ctx)
	default:
		fmt.Printf("Invalid command. Use [%s|%s|%s]\n", CommandDaemon, CommandCreatePoll, CommandShowDB)
		os.Exit(1)
	}

	err := application.Run(ctx)
	if err != nil {
		log.Fatalf("error on run application: %v", err)
	}
}
