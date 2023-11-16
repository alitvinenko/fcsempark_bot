package main

import (
	"fmt"
	"github.com/alitvinenko/fcsempark_bot/internal/app/create_poll"
	"github.com/alitvinenko/fcsempark_bot/internal/app/daemon"
	"github.com/alitvinenko/fcsempark_bot/internal/app/showdb"
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

	switch command {
	case CommandDaemon:
		daemon.Run()
	case CommandCreatePoll:
		create_poll.Run()
	case CommandShowDB:
		showdb.Run()
	default:
		fmt.Printf("Invalid command. Use [%s|%s|%s]\n", CommandDaemon, CommandCreatePoll, CommandShowDB)
		os.Exit(1)
	}
}
